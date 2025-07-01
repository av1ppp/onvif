package logging

import (
	"os"
	"time"

	"github.com/av1ppp/logx"
	"github.com/av1ppp/logx/handlercolor1"
	"github.com/fatih/color"
)

func New() *logx.Logger {
	return logx.New(handlercolor1.New(os.Stdout, &handlercolor1.Options{
		Level:         logx.LevelDebug,
		TimeFormat:    time.DateTime,
		SrcFileMode:   handlercolor1.ShortFile,
		SrcFileLength: 0,
		MsgPrefix:     color.HiWhiteString("| "),
		MsgLength:     0,
		MsgColor:      color.New(),
		NoColor:       false,
	})).With(
		logx.Module("onvif"),
	)
}
