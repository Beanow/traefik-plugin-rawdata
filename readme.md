# Traefik Plugin Rawdata

> Traefik Provider plugin to import `/api/rawdata` from another Traefik instance.

**Highly experimental**, you might use this as a stopgap solution while migrating services for example. But in production I would avoiding using this plugin as a permanent solution.

## Usecase

Traefik's built-in provider configuration has some limits.

For example, https://github.com/traefik/traefik/issues/6063
while Docker Swarm and Docker as-is are supported. You can't specify them both.
Nor can you use multiple HTTP endpoints https://doc.traefik.io/traefik/providers/http/#endpoint.

This plugin offers a poor man workaround, where you add another Traefik instance *purely for discovery* and import this to your main Traefik instance.

## Gotchas

This plugin **does not attempt** to rewrite service IP's, networking information, or anything like hat. So whatever your secondary Traefik instance discovers, needs to be accessible by your primary Traefik instance in the same way.
