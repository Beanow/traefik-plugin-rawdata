package dynamic

// Code generated by centrifuge. DO NOT EDIT.

type AddPrefix struct {
	Prefix string `json:"prefix,omitempty"`
}

type BasicAuth struct {
	Users        Users  `json:"users,omitempty"`
	UsersFile    string `json:"usersFile,omitempty"`
	Realm        string `json:"realm,omitempty"`
	RemoveHeader bool   `json:"removeHeader,omitempty"`
	HeaderField  string `json:"headerField,omitempty"`
}

type Buffering struct {
	MaxRequestBodyBytes  int64  `json:"maxRequestBodyBytes,omitempty"`
	MemRequestBodyBytes  int64  `json:"memRequestBodyBytes,omitempty"`
	MaxResponseBodyBytes int64  `json:"maxResponseBodyBytes,omitempty"`
	MemResponseBodyBytes int64  `json:"memResponseBodyBytes,omitempty"`
	RetryExpression      string `json:"retryExpression,omitempty"`
}

type Chain struct {
	Middlewares []string `json:"middlewares,omitempty"`
}

type CircuitBreaker struct {
	Expression string `json:"expression,omitempty"`
}

type ClientTLS struct {
	CA                 string `json:"ca,omitempty"`
	CAOptional         bool   `json:"caOptional,omitempty"`
	Cert               string `json:"cert,omitempty"`
	Key                string `json:"key,omitempty"`
	InsecureSkipVerify bool   `json:"insecureSkipVerify,omitempty"`
}

type Compress struct {
	ExcludedContentTypes []string `json:"excludedContentTypes,omitempty"`
}

type ContentType struct {
	AutoDetect bool `json:"autoDetect,omitempty"`
}

type DigestAuth struct {
	Users        Users  `json:"users,omitempty"`
	UsersFile    string `json:"usersFile,omitempty"`
	RemoveHeader bool   `json:"removeHeader,omitempty"`
	Realm        string `json:"realm,omitempty"`
	HeaderField  string `json:"headerField,omitempty"`
}

type ErrorPage struct {
	Status  []string `json:"status,omitempty"`
	Service string   `json:"service,omitempty"`
	Query   string   `json:"query,omitempty"`
}

type ForwardAuth struct {
	Address                  string     `json:"address,omitempty"`
	TLS                      *ClientTLS `json:"tls,omitempty"`
	TrustForwardHeader       bool       `json:"trustForwardHeader,omitempty"`
	AuthResponseHeaders      []string   `json:"authResponseHeaders,omitempty"`
	AuthResponseHeadersRegex string     `json:"authResponseHeadersRegex,omitempty"`
	AuthRequestHeaders       []string   `json:"authRequestHeaders,omitempty"`
}

type Headers struct {
	CustomRequestHeaders              map[string]string `json:"customRequestHeaders,omitempty"`
	CustomResponseHeaders             map[string]string `json:"customResponseHeaders,omitempty"`
	AccessControlAllowCredentials     bool              `json:"accessControlAllowCredentials,omitempty"`
	AccessControlAllowHeaders         []string          `json:"accessControlAllowHeaders,omitempty"`
	AccessControlAllowMethods         []string          `json:"accessControlAllowMethods,omitempty"`
	AccessControlAllowOriginList      []string          `json:"accessControlAllowOriginList,omitempty"`
	AccessControlAllowOriginListRegex []string          `json:"accessControlAllowOriginListRegex,omitempty"`
	AccessControlExposeHeaders        []string          `json:"accessControlExposeHeaders,omitempty"`
	AccessControlMaxAge               int64             `json:"accessControlMaxAge,omitempty"`
	AddVaryHeader                     bool              `json:"addVaryHeader,omitempty"`
	AllowedHosts                      []string          `json:"allowedHosts,omitempty"`
	HostsProxyHeaders                 []string          `json:"hostsProxyHeaders,omitempty"`
	SSLRedirect                       bool              `json:"sslRedirect,omitempty"`
	SSLTemporaryRedirect              bool              `json:"sslTemporaryRedirect,omitempty"`
	SSLHost                           string            `json:"sslHost,omitempty"`
	SSLProxyHeaders                   map[string]string `json:"sslProxyHeaders,omitempty"`
	SSLForceHost                      bool              `json:"sslForceHost,omitempty"`
	STSSeconds                        int64             `json:"stsSeconds,omitempty"`
	STSIncludeSubdomains              bool              `json:"stsIncludeSubdomains,omitempty"`
	STSPreload                        bool              `json:"stsPreload,omitempty"`
	ForceSTSHeader                    bool              `json:"forceSTSHeader,omitempty"`
	FrameDeny                         bool              `json:"frameDeny,omitempty"`
	CustomFrameOptionsValue           string            `json:"customFrameOptionsValue,omitempty"`
	ContentTypeNosniff                bool              `json:"contentTypeNosniff,omitempty"`
	BrowserXSSFilter                  bool              `json:"browserXssFilter,omitempty"`
	CustomBrowserXSSValue             string            `json:"customBrowserXSSValue,omitempty"`
	ContentSecurityPolicy             string            `json:"contentSecurityPolicy,omitempty"`
	PublicKey                         string            `json:"publicKey,omitempty"`
	ReferrerPolicy                    string            `json:"referrerPolicy,omitempty"`
	FeaturePolicy                     string            `json:"featurePolicy,omitempty"`
	PermissionsPolicy                 string            `json:"permissionsPolicy,omitempty"`
	IsDevelopment                     bool              `json:"isDevelopment,omitempty"`
}

type IPStrategy struct {
	Depth       int      `json:"depth,omitempty"`
	ExcludedIPs []string `json:"excludedIPs,omitempty"`
}

type IPWhiteList struct {
	SourceRange []string    `json:"sourceRange,omitempty"`
	IPStrategy  *IPStrategy `json:"ipStrategy,omitempty"`
}

type InFlightReq struct {
	Amount          int64            `json:"amount,omitempty"`
	SourceCriterion *SourceCriterion `json:"sourceCriterion,omitempty"`
}

type Middleware struct {
	AddPrefix         *AddPrefix            `json:"addPrefix,omitempty"`
	StripPrefix       *StripPrefix          `json:"stripPrefix,omitempty"`
	StripPrefixRegex  *StripPrefixRegex     `json:"stripPrefixRegex,omitempty"`
	ReplacePath       *ReplacePath          `json:"replacePath,omitempty"`
	ReplacePathRegex  *ReplacePathRegex     `json:"replacePathRegex,omitempty"`
	Chain             *Chain                `json:"chain,omitempty"`
	IPWhiteList       *IPWhiteList          `json:"ipWhiteList,omitempty"`
	Headers           *Headers              `json:"headers,omitempty"`
	Errors            *ErrorPage            `json:"errors,omitempty"`
	RateLimit         *RateLimit            `json:"rateLimit,omitempty"`
	RedirectRegex     *RedirectRegex        `json:"redirectRegex,omitempty"`
	RedirectScheme    *RedirectScheme       `json:"redirectScheme,omitempty"`
	BasicAuth         *BasicAuth            `json:"basicAuth,omitempty"`
	DigestAuth        *DigestAuth           `json:"digestAuth,omitempty"`
	ForwardAuth       *ForwardAuth          `json:"forwardAuth,omitempty"`
	InFlightReq       *InFlightReq          `json:"inFlightReq,omitempty"`
	Buffering         *Buffering            `json:"buffering,omitempty"`
	CircuitBreaker    *CircuitBreaker       `json:"circuitBreaker,omitempty"`
	Compress          *Compress             `json:"compress,omitempty"`
	PassTLSClientCert *PassTLSClientCert    `json:"passTLSClientCert,omitempty"`
	Retry             *Retry                `json:"retry,omitempty"`
	ContentType       *ContentType          `json:"contentType,omitempty"`
	Plugin            map[string]PluginConf `json:"plugin,omitempty"`
}

type PassTLSClientCert struct {
	PEM  bool                      `json:"pem,omitempty"`
	Info *TLSClientCertificateInfo `json:"info,omitempty"`
}

type RateLimit struct {
	Average         int64            `json:"average,omitempty"`
	Period          string           `json:"period,omitempty"`
	Burst           int64            `json:"burst,omitempty"`
	SourceCriterion *SourceCriterion `json:"sourceCriterion,omitempty"`
}

type RedirectRegex struct {
	Regex       string `json:"regex,omitempty"`
	Replacement string `json:"replacement,omitempty"`
	Permanent   bool   `json:"permanent,omitempty"`
}

type RedirectScheme struct {
	Scheme    string `json:"scheme,omitempty"`
	Port      string `json:"port,omitempty"`
	Permanent bool   `json:"permanent,omitempty"`
}

type ReplacePath struct {
	Path string `json:"path,omitempty"`
}

type ReplacePathRegex struct {
	Regex       string `json:"regex,omitempty"`
	Replacement string `json:"replacement,omitempty"`
}

type Retry struct {
	Attempts        int    `json:"attempts,omitempty"`
	InitialInterval string `json:"initialInterval,omitempty"`
}

type SourceCriterion struct {
	IPStrategy        *IPStrategy `json:"ipStrategy,omitempty"`
	RequestHeaderName string      `json:"requestHeaderName,omitempty"`
	RequestHost       bool        `json:"requestHost,omitempty"`
}

type StripPrefix struct {
	Prefixes   []string `json:"prefixes,omitempty"`
	ForceSlash bool     `json:"forceSlash,omitempty"`
}

type StripPrefixRegex struct {
	Regex []string `json:"regex,omitempty"`
}

type TLSCLientCertificateDNInfo struct {
	Country         bool `json:"country,omitempty"`
	Province        bool `json:"province,omitempty"`
	Locality        bool `json:"locality,omitempty"`
	Organization    bool `json:"organization,omitempty"`
	CommonName      bool `json:"commonName,omitempty"`
	SerialNumber    bool `json:"serialNumber,omitempty"`
	DomainComponent bool `json:"domainComponent,omitempty"`
}

type TLSClientCertificateInfo struct {
	NotAfter     bool                        `json:"notAfter,omitempty"`
	NotBefore    bool                        `json:"notBefore,omitempty"`
	Sans         bool                        `json:"sans,omitempty"`
	Subject      *TLSCLientCertificateDNInfo `json:"subject,omitempty"`
	Issuer       *TLSCLientCertificateDNInfo `json:"issuer,omitempty"`
	SerialNumber bool                        `json:"serialNumber,omitempty"`
}

type Users []string