package logging

import (
	"os"

	"github.com/av1ppp/logx"
	"github.com/av1ppp/logx/handlercolor1"
)

func New() *logx.Logger {
	return logx.New(handlercolor1.New(os.Stdout, handlercolor1.DefaultOptions)).With(
		logx.Module("onvif"),
	)
}
