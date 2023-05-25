package api_client

import (
	"net/url"
	"sync"

	kratos "github.com/ory/kratos-client-go"
)

var (
	publicClientOnce     sync.Once
	publicClientInstance *kratos.APIClient
)

func InitPublicClient(url url.URL) *kratos.APIClient {

	publicClientOnce.Do(func() { // <-- atomic, does not allow repeating
		// publicClientInstance = kratos.NewHTTPClientWithConfig(
		// 	nil,
		// 	&kratos.TransportConfig{
		// 		Schemes:  []string{url.Scheme},
		// 		Host:     url.Host,
		// 		BasePath: url.Path})

		// publicClientInstance =

		cfg := kratos.NewConfiguration()
		cfg.Host = url.Host
		cfg.Scheme = url.Scheme
		cfg.Servers = kratos.ServerConfigurations{
			{URL: url.Path},
		}

		publicClientInstance = kratos.NewAPIClient(cfg)
	})

	return publicClientInstance
}

func PublicClient() *kratos.APIClient {
	return publicClientInstance
}

var (
	adminClientOnce     sync.Once
	adminClientInstance *kratos.APIClient
)

func InitAdminClient(url url.URL) *kratos.APIClient {
	adminClientOnce.Do(func() {
		cfg := kratos.NewConfiguration()
		cfg.Host = url.Host
		cfg.Scheme = url.Scheme
		cfg.Servers = kratos.ServerConfigurations{
			{URL: url.Path},
		}

		adminClientInstance = kratos.NewAPIClient(cfg)
	})

	return adminClientInstance
}

func AdminClient() *kratos.APIClient {
	return adminClientInstance
}
