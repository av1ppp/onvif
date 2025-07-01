package api

import (
	"github.com/av1ppp/onvif/errors"
)

var (
	root = errors.Root.NewSubNamespace("api")

	commonErrors = root.NewType("common")
)
