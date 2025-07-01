package errors

import (
	"encoding/xml"

	"github.com/av1ppp/onvif/xsd"
)

// https://www.onvif.org/specs/core/ONVIF-Core-Specification-v1612.pdf page 26
type FaultEnvelope struct {
	Header struct{} `xml:"Header"`
	Body   struct {
		Fault Fault `xml:"Fault"`
	} `xml:"Body"`
}

type Fault struct {
	XMLName xml.Name    `xml:"Fault"`
	Code    FaultCode   `xml:"Code"`
	Reason  FaultReason `xml:"Reason"`
	Detail  FaultDetail `xml:"Detail"`
}

type FaultCode struct {
	XMLName xml.Name          `xml:"Code"`
	Value   xsd.String        `xml:"Value"`
	Subcode *FaultCodeSubcode `xml:"Subcode"`
}

type FaultCodeSubcode struct {
	XMLName xml.Name          `xml:"Subcode"`
	Value   xsd.String        `xml:"Value"`
	Subcode *FaultCodeSubcode `xml:"Subcode"`
}

type FaultReason struct {
	XMLName xml.Name        `xml:"Reason"`
	Text    FaultReasonText `xml:"Text"`
}

type FaultReasonText struct {
	XMLName xml.Name     `xml:"Text"`
	Text    xsd.String   `xml:",chardata"`
	Lang    xsd.Language `xml:"lang,attr"`
}

type FaultDetail struct {
	XMLName xml.Name   `xml:"Detail"`
	Text    xsd.String `xml:"Text"`
}
