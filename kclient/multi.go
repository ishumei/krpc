package kclient

import (
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/day253/krpc/conf"
	"github.com/day253/krpc/logging"
	"github.com/day253/krpc/objects"
	"github.com/day253/krpc/zookeeper"
	"github.com/samber/do"
)

type MultiClientConf struct {
	ClientName   string
	ResolverConf `mapstructure:",squash"`
	Models       map[string]ClientConf
}

func MustNewMultiClientConf(path, file, suffix string) {
	c := &MultiClientConf{}
	conf.MustLoadConf(c, path, file, suffix)
	klog.Info("load: ", objects.String(c))
	InjectClientFromMultiClientConf(c)
}

func InjectClientFromMultiClientConf(c *MultiClientConf) {
	do.Override(Injector, func(i *do.Injector) (discovery.Resolver, error) {
		logger, err := do.Invoke[*logging.Logger](logging.Injector)
		if err == nil {
			return zookeeper.NewZookeeperResolverWithConf(
				c.ResolverConf.Resolver,
				zookeeper.WithLogger(logger),
			)
		} else {
			return zookeeper.NewZookeeperResolverWithConf(
				c.ResolverConf.Resolver,
			)
		}
	})
	for name, opt := range c.Models {
		switch opt.Type {
		case ClientTypeAudio:
			do.OverrideNamedValue(Injector, name, MustNewAudioClientWithInjector(opt))
		case ClientTypeEvent:
			do.OverrideNamedValue(Injector, name, MustNewEventClientWithInjector(opt))
		case ClientTypeImage:
			do.OverrideNamedValue(Injector, name, MustNewImageClientWithInjector(opt))
		case ClientTypeText:
			do.OverrideNamedValue(Injector, name, MustNewTextClientWithInjector(opt))
		case ClientTypeArbiter:
			fallthrough
		default:
			do.OverrideNamedValue(Injector, name, MustNewArbiterClientWithInjector(opt))
		}
	}
}
