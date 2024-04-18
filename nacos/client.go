package nacos

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func newConfigClient(host, tenant string) (config_client.IConfigClient, error) {
	sc := []constant.ServerConfig{
		{
			IpAddr: host,
			Port:   80,
		},
	}
	cc := constant.ClientConfig{
		NamespaceId:         tenant,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		CacheDir:            "/tmp/sdk-nacos/",
		LogLevel:            "error",
	}
	tmp, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		})

	if err != nil {
		return nil, err
	}

	return tmp, nil
}

// GetContent 通过dataID获取数据
func (nc *NConfig) GetContent(dataID string) (content string, err error) {
	content, err = nc.Client.GetConfig(vo.ConfigParam{
		DataId: dataID,
		Group:  c.Group,
	})
	return content, err
}

// DataIDMD5 获取指定dataid的md5值
func (nc *NConfig) DataIDMD5(dataID string) (md5 string, err error) {
	content, err := nc.GetContent(dataID)
	if err != nil {
		return "", err
	}
	if content == "" {
		return "", fmt.Errorf("data id: %d, content is empty", dataID)
	}
	return hashx.Md5(content), nil
}

// PublishConfig 数据写入nacos
func (nc *NConfig) PublishConfig(dataID, content string) (b bool, err error) {
	b, err = nc.Client.PublishConfig(vo.ConfigParam{
		DataId:  dataID,
		Group:   c.Group,
		Content: content,
		Type:    vo.JSON,
	})
	return b, err
}
