package errors

import "github.com/joomcode/errorx"

var (
	Root = errorx.NewNamespace("onvif")

	Common = Root.NewType("common")

	traitSoap = errorx.RegisterTrait("soap")
	Soap      = Root.NewType("soap", traitSoap)

	PropMethod     = errorx.RegisterProperty("method")
	PropStatusCode = errorx.RegisterProperty("status_code")
	PropSoapCode   = errorx.RegisterProperty("soap_code")
)

func IsSoapError(err error) bool {
	return errorx.HasTrait(err, traitSoap)
}
