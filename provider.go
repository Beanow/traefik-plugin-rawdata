package traefik_plugin_rawdata

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/traefik/genconf/dynamic"
)

// Config the plugin configuration.
type Config struct {
	PollInterval   string `json:"pollInterval,omitempty"`
	Endpoint       string `json:"endpoint,omitempty"`
	Namespace      string `json:"namespace,omitempty"`
	StripNamespace bool   `json:"stripNamespace,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		PollInterval:   "5s",                                // 5 * time.Second
		Endpoint:       "http://127.0.0.1:8080/api/rawdata", // default unsafe api endpoint
		Namespace:      "file",
		StripNamespace: true,
	}
}

// Provider a simple provider plugin.
type Provider struct {
	name           string
	pollInterval   time.Duration
	endpoint       url.URL
	namespace      string
	stripNamespace bool
	suffix         string

	cancel func()
	client *http.Client
}

type RunTimeRepresentation struct {
	Routers        map[string]*dynamic.Router        `json:"routers,omitempty"`
	Middlewares    map[string]*dynamic.Middleware    `json:"middlewares,omitempty"`
	Services       map[string]*dynamic.Service       `json:"services,omitempty"`
	TCPRouters     map[string]*dynamic.TCPRouter     `json:"tcpRouters,omitempty"`
	TCPMiddlewares map[string]*dynamic.TCPMiddleware `json:"tcpMiddlewares,omitempty"`
	TCPServices    map[string]*dynamic.TCPService    `json:"tcpServices,omitempty"`
	UDPRouters     map[string]*dynamic.UDPRouter     `json:"udpRouters,omitempty"`
	UDPServices    map[string]*dynamic.UDPService    `json:"udpServices,omitempty"`
}

// New creates a new Provider plugin.
func New(ctx context.Context, config *Config, name string) (*Provider, error) {
	pi, err := time.ParseDuration(config.PollInterval)
	if err != nil {
		return nil, err
	}

	end, err := url.ParseRequestURI(config.Endpoint)
	if err != nil {
		return nil, err
	}

	return &Provider{
		name:           name,
		pollInterval:   pi,
		endpoint:       *end,
		namespace:      config.Namespace,
		stripNamespace: config.StripNamespace,
		suffix:         fmt.Sprintf("@%s", config.Namespace),
		client:         &http.Client{Timeout: 5 * time.Second},
	}, nil
}

// Init the provider.
func (p *Provider) Init() error {
	if p.pollInterval <= 0 {
		return fmt.Errorf("poll interval must be greater than 0")
	}

	return nil
}

// Provide creates and send dynamic configuration.
func (p *Provider) Provide(cfgChan chan<- json.Marshaler) error {
	ctx, cancel := context.WithCancel(context.Background())
	p.cancel = cancel

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Print(err)
			}
		}()

		p.loadConfiguration(ctx, cfgChan)
	}()

	return nil
}

func (p *Provider) loadConfiguration(ctx context.Context, cfgChan chan<- json.Marshaler) {
	ticker := time.NewTicker(p.pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			configuration, err := p.pollEndpoint()
			if err != nil {
				log.Print(err)
			} else {
				cfgChan <- &dynamic.JSONPayload{Configuration: configuration}
			}

		case <-ctx.Done():
			return
		}
	}
}

// Stop to stop the provider and the related go routines.
func (p *Provider) Stop() error {
	p.cancel()
	return nil
}

func (p *Provider) pollEndpoint() (*dynamic.Configuration, error) {
	resp, err := p.client.Get(p.endpoint.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoint: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()
	data, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: %d: %s", resp.StatusCode, string(data))
	}

	var runtime *RunTimeRepresentation
	err = json.Unmarshal(data, &runtime)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshall response: %w", err)
	}

	return p.filteredConfig(runtime)
}

func (p *Provider) strippedKey(key string) string {
	if !strings.HasSuffix(key, p.suffix) {
		return ""
	}
	if p.stripNamespace {
		return strings.TrimSuffix(key, p.suffix)
	}
	return key
}

func (p *Provider) filteredConfig(runtime *RunTimeRepresentation) (*dynamic.Configuration, error) {
	conf := &dynamic.Configuration{
		HTTP: &dynamic.HTTPConfiguration{
			Routers:     make(map[string]*dynamic.Router),
			Middlewares: make(map[string]*dynamic.Middleware),
			Services:    make(map[string]*dynamic.Service),
		},
		TCP: &dynamic.TCPConfiguration{
			Routers:     make(map[string]*dynamic.TCPRouter),
			Middlewares: make(map[string]*dynamic.TCPMiddleware),
			Services:    make(map[string]*dynamic.TCPService),
		},
		UDP: &dynamic.UDPConfiguration{
			Routers:  make(map[string]*dynamic.UDPRouter),
			Services: make(map[string]*dynamic.UDPService),
		},
	}

	// Yay for generics not being available yet.

	for name, spec := range runtime.Routers {
		if name = p.strippedKey(name); name != "" {
			conf.HTTP.Routers[name] = spec
		}
	}

	for name, spec := range runtime.Middlewares {
		if name = p.strippedKey(name); name != "" {
			conf.HTTP.Middlewares[name] = spec
		}
	}
	for name, spec := range runtime.Services {
		if name = p.strippedKey(name); name != "" {
			conf.HTTP.Services[name] = spec
		}
	}

	for name, spec := range runtime.TCPRouters {
		if name = p.strippedKey(name); name != "" {
			conf.TCP.Routers[name] = spec
		}
	}
	for name, spec := range runtime.TCPMiddlewares {
		if name = p.strippedKey(name); name != "" {
			conf.TCP.Middlewares[name] = spec
		}
	}
	for name, spec := range runtime.TCPServices {
		if name = p.strippedKey(name); name != "" {
			conf.TCP.Services[name] = spec
		}
	}

	for name, spec := range runtime.UDPRouters {
		if name = p.strippedKey(name); name != "" {
			conf.UDP.Routers[name] = spec
		}
	}
	for name, spec := range runtime.UDPServices {
		if name = p.strippedKey(name); name != "" {
			conf.UDP.Services[name] = spec
		}
	}

	return conf, nil
}
