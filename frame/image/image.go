package image

import (
	"github.com/cloudwego/kitex/server"
	"github.com/ishumei/krpc/frame/kservice"
	"github.com/ishumei/krpc/frame/sconfig"
	"github.com/ishumei/krpc/frame/ssuite"
	"github.com/ishumei/krpc/protocols/image/kitex_gen/shumei/strategy/re"
	"github.com/ishumei/krpc/protocols/image/kitex_gen/shumei/strategy/re/imagepredictor"
	"github.com/samber/do"
)

type ImageService struct {
	*kservice.Kservice
}

func NewImageService(i *do.Injector) (*ImageService, error) {
	opts := do.MustInvoke[*ssuite.ServerOptions](sconfig.Injector)

	predictor := do.MustInvoke[re.ImagePredictor](sconfig.Injector)

	return &ImageService{
		Kservice: kservice.MustNewKservice(i, imagepredictor.NewServer(predictor, server.WithSuite(opts))),
	}, nil
}

func init() {
	do.Provide(sconfig.Injector, NewImageService)
}
