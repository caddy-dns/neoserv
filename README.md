Neoserv (NEOSERV.si) module for Caddy
===========================

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with [Neoserv](https://neoserv.si) (using [moj.neoserv.si](https://moj.neoserv.si) account).

## Two-factor login

Neoserv requires two-factor login (TOTP) for unattended access; without it, every
login from a new device asks for a code sent by email. Enable it in Moj NEOSERV
("Omogoči dvostopenjsko prijavo") and use the manual setup key ("Ključ, za ročno
nameščanje") as `totp_secret`. See the
[libdns/neoserv README](https://github.com/libdns/neoserv#two-factor-login) for
the step-by-step setup.

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
				"password": "NEOSERV_PASSWORD",
				"totp_secret": "NEOSERV_TOTP"
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
		totp_secret {env.NEOSERV_TOTP}
	}
}
```

```
# one site
tls {
	dns neoserv {
		username {env.NEOSERV_USERNAME}
		password {env.NEOSERV_PASSWORD}
		totp_secret {env.NEOSERV_TOTP}
	}
}
```

You can replace `{env.NEOSERV_USERNAME}`, `{env.NEOSERV_PASSWORD}` and `{env.NEOSERV_TOTP}` with your actual Neoserv username, password and two-factor key if you prefer to put it directly in your config insead of an environment variable.
