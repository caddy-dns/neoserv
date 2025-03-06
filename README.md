Neoserv (NEOSERV.si) module for Caddy
===========================

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with [Neoserv](https://neoserv.si) (using [moj.neoserv.si](https://moj.neoserv.si) account).

## Caddy module name

```
dns.providers.neoserv
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuers/acme/) like so:

```json
{
	"module": "acme",
	"challenges": {
		"dns": {
			"provider": {
				"name": "neoserv",
				"username": "NEOSERV_USERNAME",
				"password": "NEOSERV_PASSWORD"
			}
		}
	}
}
```

or with the Caddyfile:

```
# globally
{
	acme_dns neoserv {
		username {env.NEOSERV_USERNAME}
		password {env.NEOSERV_PASSWORD}
	}
}
```

```
# one site
tls {
	dns neoserv {
		username {env.NEOSERV_USERNAME}
		password {env.NEOSERV_PASSWORD}
	}
}
```

You can replace `{env.NEOSERV_USERNAME}` and `{env.NEOSERV_PASSWORD}` with your actual Neoserv username and password if you prefer to put it directly in your config insead of an environment variable.
