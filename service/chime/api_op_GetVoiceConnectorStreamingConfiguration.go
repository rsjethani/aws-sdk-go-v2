// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetVoiceConnectorStreamingConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime Voice Connector ID.
	//
	// VoiceConnectorId is a required field
	VoiceConnectorId *string `location:"uri" locationName:"voiceConnectorId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetVoiceConnectorStreamingConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetVoiceConnectorStreamingConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetVoiceConnectorStreamingConfigurationInput"}

	if s.VoiceConnectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VoiceConnectorId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetVoiceConnectorStreamingConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.VoiceConnectorId != nil {
		v := *s.VoiceConnectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "voiceConnectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetVoiceConnectorStreamingConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// The streaming configuration details.
	StreamingConfiguration *StreamingConfiguration `type:"structure"`
}

// String returns the string representation
func (s GetVoiceConnectorStreamingConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetVoiceConnectorStreamingConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.StreamingConfiguration != nil {
		v := s.StreamingConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "StreamingConfiguration", v, metadata)
	}
	return nil
}

const opGetVoiceConnectorStreamingConfiguration = "GetVoiceConnectorStreamingConfiguration"

// GetVoiceConnectorStreamingConfigurationRequest returns a request value for making API operation for
// Amazon Chime.
//
// Retrieves the streaming configuration details for the specified Amazon Chime
// Voice Connector. Shows whether media streaming is enabled for sending to
// Amazon Kinesis, and shows the retention period for the Amazon Kinesis data,
// in hours.
//
//    // Example sending a request using GetVoiceConnectorStreamingConfigurationRequest.
//    req := client.GetVoiceConnectorStreamingConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/GetVoiceConnectorStreamingConfiguration
func (c *Client) GetVoiceConnectorStreamingConfigurationRequest(input *GetVoiceConnectorStreamingConfigurationInput) GetVoiceConnectorStreamingConfigurationRequest {
	op := &aws.Operation{
		Name:       opGetVoiceConnectorStreamingConfiguration,
		HTTPMethod: "GET",
		HTTPPath:   "/voice-connectors/{voiceConnectorId}/streaming-configuration",
	}

	if input == nil {
		input = &GetVoiceConnectorStreamingConfigurationInput{}
	}

	req := c.newRequest(op, input, &GetVoiceConnectorStreamingConfigurationOutput{})
	return GetVoiceConnectorStreamingConfigurationRequest{Request: req, Input: input, Copy: c.GetVoiceConnectorStreamingConfigurationRequest}
}

// GetVoiceConnectorStreamingConfigurationRequest is the request type for the
// GetVoiceConnectorStreamingConfiguration API operation.
type GetVoiceConnectorStreamingConfigurationRequest struct {
	*aws.Request
	Input *GetVoiceConnectorStreamingConfigurationInput
	Copy  func(*GetVoiceConnectorStreamingConfigurationInput) GetVoiceConnectorStreamingConfigurationRequest
}

// Send marshals and sends the GetVoiceConnectorStreamingConfiguration API request.
func (r GetVoiceConnectorStreamingConfigurationRequest) Send(ctx context.Context) (*GetVoiceConnectorStreamingConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetVoiceConnectorStreamingConfigurationResponse{
		GetVoiceConnectorStreamingConfigurationOutput: r.Request.Data.(*GetVoiceConnectorStreamingConfigurationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetVoiceConnectorStreamingConfigurationResponse is the response type for the
// GetVoiceConnectorStreamingConfiguration API operation.
type GetVoiceConnectorStreamingConfigurationResponse struct {
	*GetVoiceConnectorStreamingConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetVoiceConnectorStreamingConfiguration request.
func (r *GetVoiceConnectorStreamingConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
