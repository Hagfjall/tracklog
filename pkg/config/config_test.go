package config

import (
	"bytes"
	"testing"
)

var data = []byte(`[server]
development = true
listen_address = ":8080"
csrf_auth_key = "secr"
signing_key = "secret"
mapbox_access_token = "abc"
reverse_proxy_auth = true
reverse_proxy_auth_header = "X-Username"
reverse_proxy_auth_auto_register = false

[db]
driver = "postgres"
dsn = "dbname=tracklog"
`)

func TestRead(t *testing.T) {
	c, err := Read(bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	if expected := true; c.Server.Development != expected {
		t.Errorf("expected Server.Development = %v; got %v", expected, c.Server.Development)
	}
	if expected := ":8080"; c.Server.ListenAddress != expected {
		t.Errorf("expected Server.ListenAddress = %q; got %q", expected, c.Server.ListenAddress)
	}
	if expected := "secr"; c.Server.CSRFAuthKey != expected {
		t.Errorf("expected Server.CSRFAuthKey = %q; got %q", expected, c.Server.CSRFAuthKey)
	}
	if expected := "secret"; c.Server.SigningKey != expected {
		t.Errorf("expected Server.SigningKey = %q; got %q", expected, c.Server.SigningKey)
	}
	if expected := "abc"; c.Server.MapboxAccessToken != expected {
		t.Errorf("expected Server.MapboxAccessToken = %q; got %q", expected, c.Server.MapboxAccessToken)
	}
	if expected := true; c.Server.ReverseProxyAuth != expected {
		t.Errorf("expected Server.ReverseProxyAuth = %v; got %v", expected, c.Server.ReverseProxyAuth)
	}
	if expected := "X-Username"; c.Server.ReverseProxyAuthHeader != expected {
		t.Errorf("expected Server.ReverseProxyAuthHeader = %q; got %q", expected, c.Server.ReverseProxyAuthHeader)
	}
	if expected := false; c.Server.ReverseProxyAuthAutoRegister != expected {
		t.Errorf("expected Server.ReverseProxyAuthAutoRegister = %v; got %v", expected, c.Server.ReverseProxyAuthAutoRegister)
	}
	if expected := "postgres"; c.DB.Driver != expected {
		t.Errorf("expected DB.Driver = %q; got %q", expected, c.DB.Driver)
	}
	if expected := "dbname=tracklog"; c.DB.DSN != expected {
		t.Errorf("expected DB.DSN = %q; got %q", expected, c.DB.DSN)
	}
}
