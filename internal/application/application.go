package application

import (
	"github.com/saeedrm6/web-service-go/internal/config"
	"github.com/saeedrm6/web-service-go/pkg/translator/i18n"
)

func Run(cfg *config.Config) error  {
	translator, err := i18n.New(cfg.I18n.BundlePath)
	if err != nil {
		return err
	}

	_ = translator
	return err
}
