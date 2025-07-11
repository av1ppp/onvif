package sdkmedia

//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetServiceCapabilities
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetVideoSources
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetAudioSources
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetAudioOutputs
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media CreateProfile
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetProfile
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetProfiles
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media AddVideoEncoderConfiguration
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media RemoveVideoEncoderConfiguration
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media AddVideoSourceConfiguration
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media RemoveVideoSourceConfiguration
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media AddAudioEncoderConfiguration
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media RemoveAudioEncoderConfiguration
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media AddAudioSourceConfiguration
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media RemoveAudioSourceConfiguration
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media AddPTZConfiguration
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media RemovePTZConfiguration
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media AddVideoAnalyticsConfiguration
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media RemoveVideoAnalyticsConfiguration
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media AddMetadataConfiguration
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media RemoveMetadataConfiguration
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media AddAudioOutputConfiguration
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media RemoveAudioOutputConfiguration
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media AddAudioDecoderConfiguration
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media RemoveAudioDecoderConfiguration
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media DeleteProfile
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetVideoSourceConfigurations
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetVideoEncoderConfigurations
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetAudioSourceConfigurations
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetAudioEncoderConfigurations
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetVideoAnalyticsConfigurations
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetMetadataConfigurations
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetAudioOutputConfigurations
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetAudioDecoderConfigurations
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetVideoSourceConfiguration
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetVideoEncoderConfiguration
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetAudioSourceConfiguration
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetAudioEncoderConfiguration
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetVideoAnalyticsConfiguration
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetMetadataConfiguration
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetAudioOutputConfiguration
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetAudioDecoderConfiguration
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetCompatibleVideoEncoderConfigurations
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetCompatibleVideoSourceConfigurations
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetCompatibleAudioEncoderConfigurations
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetCompatibleAudioSourceConfigurations
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetCompatibleVideoAnalyticsConfigurations
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetCompatibleMetadataConfigurations
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetCompatibleAudioOutputConfigurations
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetCompatibleAudioDecoderConfigurations
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media SetVideoSourceConfiguration
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media SetVideoEncoderConfiguration
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media SetAudioSourceConfiguration
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media SetAudioEncoderConfiguration
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media SetVideoAnalyticsConfiguration
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media SetMetadataConfiguration
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media SetAudioOutputConfiguration
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media SetAudioDecoderConfiguration
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetVideoSourceConfigurationOptions
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetVideoEncoderConfigurationOptions
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetAudioSourceConfigurationOptions
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetAudioEncoderConfigurationOptions
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetMetadataConfigurationOptions
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetAudioOutputConfigurationOptions
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetAudioDecoderConfigurationOptions
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetGuaranteedNumberOfVideoEncoderInstances
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetStreamUri
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media StartMulticastStreaming
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media StopMulticastStreaming
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media SetSynchronizationPoint
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetSnapshotUri
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetVideoSourceModes
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media SetVideoSourceMode
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetOSDs
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetOSD
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media GetOSDOptions
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media SetOSD
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media CreateOSD
//go:generate go run github.com/av1ppp/onvif/sdk/codegen media media DeleteOSD
