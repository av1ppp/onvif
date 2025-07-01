package errors

import "github.com/joomcode/errorx"

var (
	Root = errorx.NewNamespace("onvif")

	Common = Root.NewType("common")

	PropMethod = errorx.RegisterProperty("method")
)
