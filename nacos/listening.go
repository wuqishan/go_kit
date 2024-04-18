package nacos

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/nacos-group/nacos-sdk-go/common/logger"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

// ListeningSDKConfig ...
func (nc *NConfig) ListeningSDKConfig(key string) {
	ctx := gctx.New()

	// 后续启动配置监听
	err := nc.Client.ListenConfig(vo.ConfigParam{
		DataId: key,
		Group:  nc.Group,
		OnChange: func(namespace, group, dataId, data string) {
			logger.Info(ctx, fmt.Sprintf("ex_config_changed, %s > %s > %s, (%s)", namespace, group, dataId, data))
			// todo...
		},
	})
	if err != nil {
		panic("无法启动配置监听, " + err.Error())
	}

	logger.Info(ctx, "nacos listening start!")
	select {}
}
