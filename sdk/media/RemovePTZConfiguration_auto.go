// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/logx"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// Call_RemovePTZConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a RemovePTZConfigurationResponse.
func Call_RemovePTZConfiguration(ctx context.Context, dev *onvif.Device, request media.RemovePTZConfiguration) (media.RemovePTZConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemovePTZConfigurationResponse media.RemovePTZConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.RemovePTZConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "RemovePTZConfiguration")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.RemovePTZConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "RemovePTZConfiguration")
	}
	return reply.Body.RemovePTZConfigurationResponse, nil
}

// CallWithLogging_RemovePTZConfiguration works like Call_RemovePTZConfiguration but also logs the response body.
func CallWithLogging_RemovePTZConfiguration(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.RemovePTZConfiguration) (media.RemovePTZConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemovePTZConfigurationResponse media.RemovePTZConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.RemovePTZConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "RemovePTZConfiguration")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "RemovePTZConfiguration")
	if err != nil {
		return reply.Body.RemovePTZConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "RemovePTZConfiguration")
	}
	return reply.Body.RemovePTZConfigurationResponse, nil
}
