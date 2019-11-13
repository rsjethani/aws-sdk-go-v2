// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type PutVoiceConnectorLoggingConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The logging configuration details to add.
	//
	// LoggingConfiguration is a required field
	LoggingConfiguration *LoggingConfiguration `type:"structure" required:"true"`

	// The Amazon Chime Voice Connector ID.
	//
	// VoiceConnectorId is a required field
	VoiceConnectorId *string `location:"uri" locationName:"voiceConnectorId" type:"string" required:"true"`
}

// String returns the string representation
func (s PutVoiceConnectorLoggingConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutVoiceConnectorLoggingConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutVoiceConnectorLoggingConfigurationInput"}

	if s.LoggingConfiguration == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoggingConfiguration"))
	}

	if s.VoiceConnectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VoiceConnectorId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutVoiceConnectorLoggingConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.LoggingConfiguration != nil {
		v := s.LoggingConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "LoggingConfiguration", v, metadata)
	}
	if s.VoiceConnectorId != nil {
		v := *s.VoiceConnectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "voiceConnectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type PutVoiceConnectorLoggingConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// The updated logging configuration details.
	LoggingConfiguration *LoggingConfiguration `type:"structure"`
}

// String returns the string representation
func (s PutVoiceConnectorLoggingConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutVoiceConnectorLoggingConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.LoggingConfiguration != nil {
		v := s.LoggingConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "LoggingConfiguration", v, metadata)
	}
	return nil
}

const opPutVoiceConnectorLoggingConfiguration = "PutVoiceConnectorLoggingConfiguration"

// PutVoiceConnectorLoggingConfigurationRequest returns a request value for making API operation for
// Amazon Chime.
//
// Adds a logging configuration for the specified Amazon Chime Voice Connector.
// The logging configuration specifies whether SIP message logs are enabled
// for sending to Amazon CloudWatch Logs.
//
//    // Example sending a request using PutVoiceConnectorLoggingConfigurationRequest.
//    req := client.PutVoiceConnectorLoggingConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/PutVoiceConnectorLoggingConfiguration
func (c *Client) PutVoiceConnectorLoggingConfigurationRequest(input *PutVoiceConnectorLoggingConfigurationInput) PutVoiceConnectorLoggingConfigurationRequest {
	op := &aws.Operation{
		Name:       opPutVoiceConnectorLoggingConfiguration,
		HTTPMethod: "PUT",
		HTTPPath:   "/voice-connectors/{voiceConnectorId}/logging-configuration",
	}

	if input == nil {
		input = &PutVoiceConnectorLoggingConfigurationInput{}
	}

	req := c.newRequest(op, input, &PutVoiceConnectorLoggingConfigurationOutput{})
	return PutVoiceConnectorLoggingConfigurationRequest{Request: req, Input: input, Copy: c.PutVoiceConnectorLoggingConfigurationRequest}
}

// PutVoiceConnectorLoggingConfigurationRequest is the request type for the
// PutVoiceConnectorLoggingConfiguration API operation.
type PutVoiceConnectorLoggingConfigurationRequest struct {
	*aws.Request
	Input *PutVoiceConnectorLoggingConfigurationInput
	Copy  func(*PutVoiceConnectorLoggingConfigurationInput) PutVoiceConnectorLoggingConfigurationRequest
}

// Send marshals and sends the PutVoiceConnectorLoggingConfiguration API request.
func (r PutVoiceConnectorLoggingConfigurationRequest) Send(ctx context.Context) (*PutVoiceConnectorLoggingConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutVoiceConnectorLoggingConfigurationResponse{
		PutVoiceConnectorLoggingConfigurationOutput: r.Request.Data.(*PutVoiceConnectorLoggingConfigurationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutVoiceConnectorLoggingConfigurationResponse is the response type for the
// PutVoiceConnectorLoggingConfiguration API operation.
type PutVoiceConnectorLoggingConfigurationResponse struct {
	*PutVoiceConnectorLoggingConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutVoiceConnectorLoggingConfiguration request.
func (r *PutVoiceConnectorLoggingConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
