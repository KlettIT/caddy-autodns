package autodns

import (
	"fmt"
	"testing"

	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/libdns/autodns"
)

func TestUnmarshalCaddyFile(t *testing.T) {
	tests := []string{
		`autodns {
			username theusername
			password secretpassword
		}`}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			// given
			dispenser := caddyfile.NewTestDispenser(tc)
			p := Provider{&autodns.Provider{}}
			// when
			err := p.UnmarshalCaddyfile(dispenser)
			// then
			if err != nil {
				t.Errorf("UnmarshalCaddyfile failed with %v", err)
				return
			}

			expectedUsername := "theusername"
			actualUsername := p.Provider.Username
			if expectedUsername != actualUsername {
				t.Errorf("Expected Username to be '%s' but got '%s'", expectedUsername, actualUsername)
			}

			expectePassword := "secretpassword"
			actualPassword := p.Provider.Password
			if expectePassword != actualPassword {
				t.Errorf("Expected SecretAccessKey to be '%s' but got '%s'", expectePassword, actualPassword)
			}
		})
	}
}