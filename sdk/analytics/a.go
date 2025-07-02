package analytics

//go:generate go run github.com/av1ppp/onvif/sdk/codegen analytics analytics CreateAnalyticsModules
//go:generate go run github.com/av1ppp/onvif/sdk/codegen analytics analytics DeleteAnalyticsModules
//go:generate go run github.com/av1ppp/onvif/sdk/codegen analytics analytics GetAnalyticsModuleOptions
//go:generate go run github.com/av1ppp/onvif/sdk/codegen analytics analytics GetAnalyticsModules
//go:generate go run github.com/av1ppp/onvif/sdk/codegen analytics analytics GetServiceCapabilities
//go:generate go run github.com/av1ppp/onvif/sdk/codegen analytics analytics GetSupportedAnalyticsModules
//go:generate go run github.com/av1ppp/onvif/sdk/codegen analytics analytics GetSupportedMetadata
//go:generate go run github.com/av1ppp/onvif/sdk/codegen analytics analytics ModifyAnalyticsModules
//go:generate go run github.com/av1ppp/onvif/sdk/codegen analytics analytics CreateRules
//go:generate go run github.com/av1ppp/onvif/sdk/codegen analytics analytics DeleteRules
//go:generate go run github.com/av1ppp/onvif/sdk/codegen analytics analytics GetRuleOptions
//go:generate go run github.com/av1ppp/onvif/sdk/codegen analytics analytics GetRules
//go:generate go run github.com/av1ppp/onvif/sdk/codegen analytics analytics GetSupportedRules
//go:generate go run github.com/av1ppp/onvif/sdk/codegen analytics analytics ModifyRules
