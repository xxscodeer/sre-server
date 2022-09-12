package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	LOGGER *zap.SugaredLogger
	CONFIG *viper.Viper
)
