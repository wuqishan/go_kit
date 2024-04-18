package nacos

import "github.com/nacos-group/nacos-sdk-go/clients/config_client"

type NConfig struct {
	Host    string
	Uri     string
	Tenant  string
	Group   string
	Content string
	Client  config_client.IConfigClient
}

func NewNConfig(host, uri, tenant, group string) *NConfig {
	tmpClient, err := newConfigClient(host, tenant)
	if err != nil {
		panic(err)
	}
	return &NConfig{
		Host:   host,
		Uri:    uri,
		Tenant: tenant,
		Group:  group,
		Client: tmpClient,
	}
}
