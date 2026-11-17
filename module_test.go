package neoserv

import (
	"testing"

	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	libdnsneoserv "github.com/libdns/neoserv"
)

func TestUnmarshalCaddyfile(t *testing.T) {
	p := Provider{new(libdnsneoserv.Provider)}
	err := p.UnmarshalCaddyfile(caddyfile.NewTestDispenser(`neoserv {
		username user@example.com
		password secret
		totp_secret ABCDEFGH
	}`))
	if err != nil || p.Username != "user@example.com" || p.Password != "secret" || p.TOTPSecret != "ABCDEFGH" {
		t.Fatalf("got %q %q %q, err %v", p.Username, p.Password, p.TOTPSecret, err)
	}

	p = Provider{new(libdnsneoserv.Provider)}
	if err := p.UnmarshalCaddyfile(caddyfile.NewTestDispenser(`neoserv {
		username user@example.com
		password secret
	}`)); err == nil {
		t.Fatal("expected error for missing totp_secret")
	}
}
