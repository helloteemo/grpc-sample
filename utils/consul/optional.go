package consul

import (
	"github.com/hashicorp/consul/api"
	"net/http"
	"time"
)

// Option optional
type Option func(cfg *api.Config)

// Token set token param
func Token(token string) Option {
	return func(cfg *api.Config) {
		cfg.Token = token
	}
}

// TokenFile set token file param
func TokenFile(tokenFile string) Option {
	return func(cfg *api.Config) {
		cfg.TokenFile = tokenFile
	}
}

// Scheme set scheme param
func Scheme(scheme string) Option {
	return func(cfg *api.Config) {
		cfg.Scheme = scheme
	}
}

// Datacenter set datacenter param
func Datacenter(datacenter string) Option {
	return func(cfg *api.Config) {
		cfg.Datacenter = datacenter
	}
}

// HttpClient set httpClient param
func HttpClient(httpClient *http.Client) Option {
	return func(cfg *api.Config) {
		cfg.HttpClient = httpClient
	}
}

// Transport set transport param
func Transport(transport *http.Transport) Option {
	return func(cfg *api.Config) {
		cfg.Transport = transport
	}
}

// HttpAuth set HttpAuth param
func HttpAuth(baseAuth *api.HttpBasicAuth) Option {
	return func(cfg *api.Config) {
		cfg.HttpAuth = baseAuth
	}
}

// WaitTime set waitTime param
func WaitTime(waitTime time.Duration) Option {
	return func(cfg *api.Config) {
		cfg.WaitTime = waitTime
	}
}

// TLSConfig set tls config param
func TLSConfig(tlsConfig api.TLSConfig) Option {
	return func(cfg *api.Config) {
		cfg.TLSConfig = tlsConfig
	}
}
