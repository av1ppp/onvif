package analytics

import (
	"github.com/av1ppp/onvif/xsd"
	"github.com/av1ppp/onvif/xsd/onvif"
)

// CreateAnalyticsModules

type CreateAnalyticsModules struct {
	XMLName            string                `xml:"tev:CreateAnalyticsModules"`
	ConfigurationToken onvif.ReferenceToken  `xml:"tan:ConfigurationToken"`
	AnalyticsModule    []onvif.ConfigRequest `xml:"tan:AnalyticsModule"`
}

type CreateAnalyticsModulesResponse struct{}

//
// DeleteAnalyticsModules

type DeleteAnalyticsModules struct {
	XMLName             string               `xml:"tan:DeleteAnalyticsModules"`
	ConfigurationToken  onvif.ReferenceToken `xml:"tan:ConfigurationToken"`
	AnalyticsModuleName []xsd.String         `xml:"tan:AnalyticsModuleName"`
}

type DeleteAnalyticsModulesResponse struct{}

//
// GetAnalyticsModuleOptions

type GetAnalyticsModuleOptions struct {
	XMLName            string               `xml:"tan:GetAnalyticsModuleOptions"`
	Type               xsd.QName            `xml:"tan:Type,omitempty"`
	ConfigurationToken onvif.ReferenceToken `xml:"tan:ConfigurationToken"`
}

type GetAnalyticsModuleOptionsResponse struct {
	Options []AnalyticsModuleOptions
}

//
// GetAnalyticsModules

type GetAnalyticsModules struct {
	XMLName            string               `xml:"tan:GetAnalyticsModules"`
	ConfigurationToken onvif.ReferenceToken `xml:"tan:ConfigurationToken"`
}

type GetAnalyticsModulesResponse struct {
	AnalyticsModule []onvif.Config
}

//
// GetServiceCapabilities

type GetServiceCapabilities struct {
	XMLName string `xml:"tan:GetServiceCapabilities"`
}

type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities
}

//
// GetSupportedAnalyticsModules

type GetSupportedAnalyticsModules struct {
	XMLName            string               `xml:"tan:GetSupportedAnalyticsModules"`
	ConfigurationToken onvif.ReferenceToken `xml:"tan:ConfigurationToken"`
}

type GetSupportedAnalyticsModulesResponse struct {
	SupportedAnalyticsModules SupportedAnalyticsModules
}

//
// GetSupportedMetadata

type GetSupportedMetadata struct {
	XMLName string    `xml:"tan:GetSupportedMetadata"`
	Type    xsd.QName `xml:"tan:Type"`
}

type GetSupportedMetadataResponse struct {
	AnalyticsModule MetadataInfo `xml:"tan:AnalyticsModule"` // optional, unbounded
}

//
// ModifyAnalyticsModules

type ModifyAnalyticsModules struct {
	XMLName            string                `xml:"tan:ModifyAnalyticsModules"`
	ConfigurationToken onvif.ReferenceToken  `xml:"tan:ConfigurationToken"`
	AnalyticsModule    []onvif.ConfigRequest `xml:"tan:AnalyticsModule"`
}

type ModifyAnalyticsModulesResponse struct{}

//
// CreateRules

type CreateRules struct {
	XMLName            string                `xml:"tan:CreateRules"`
	ConfigurationToken onvif.ReferenceToken  `xml:"tan:ConfigurationToken"`
	Rule               []onvif.ConfigRequest `xml:"tan:Rule"`
}

type CreateRulesResponse struct{}

//
// DeleteRules

type DeleteRules struct {
	XMLName            string               `xml:"tan:DeleteRules"`
	ConfigurationToken onvif.ReferenceToken `xml:"tan:ConfigurationToken"`
	RuleName           []xsd.String         `xml:"tan:RuleName"`
}

type DeleteRulesResponse struct{}

//
// GetRuleOptions

type GetRuleOptions struct {
	XMLName            string               `xml:"tan:GetRuleOptions"`
	RuleType           xsd.QName            `xml:"tan:RuleType"`
	ConfigurationToken onvif.ReferenceToken `xml:"tan:ConfigurationToken"`
}

type GetRuleOptionsResponse struct {
	RuleOptions []RuleOptions
}

//
// GetRules

type GetRules struct {
	XMLName            string               `xml:"tan:GetRules"`
	ConfigurationToken onvif.ReferenceToken `xml:"tan:ConfigurationToken"`
}

type GetRulesResponse struct {
	Rule []onvif.Config
}

//
// GetSupportedRules

type GetSupportedRules struct {
	XMLName            string               `xml:"tan:GetSupportedRules"`
	ConfigurationToken onvif.ReferenceToken `xml:"tan:ConfigurationToken"`
}

type GetSupportedRulesResponse struct {
	SupportedRules SupportedRules
}

//
// ModifyRules

type ModifyRules struct {
	XMLName            string                `xml:"tan:ModifyRules"`
	ConfigurationToken onvif.ReferenceToken  `xml:"tan:ConfigurationToken"`
	Rule               []onvif.ConfigRequest `xml:"tan:Rule"`
}

type ModifyRulesResponse struct{}
