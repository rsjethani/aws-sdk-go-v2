// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type UpdatePhoneNumberSettingsInput struct {
	_ struct{} `type:"structure"`

	// The default outbound calling name for the account.
	//
	// CallingName is a required field
	CallingName *string `type:"string" required:"true" sensitive:"true"`
}

// String returns the string representation
func (s UpdatePhoneNumberSettingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdatePhoneNumberSettingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdatePhoneNumberSettingsInput"}

	if s.CallingName == nil {
		invalidParams.Add(aws.NewErrParamRequired("CallingName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdatePhoneNumberSettingsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.CallingName != nil {
		v := *s.CallingName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CallingName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdatePhoneNumberSettingsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdatePhoneNumberSettingsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdatePhoneNumberSettingsOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdatePhoneNumberSettings = "UpdatePhoneNumberSettings"

// UpdatePhoneNumberSettingsRequest returns a request value for making API operation for
// Amazon Chime.
//
// Updates the phone number settings for the administrator's AWS account, such
// as the default outbound calling name. You can update the default outbound
// calling name once every seven days. Outbound calling names can take up to
// 72 hours to be updated.
//
//    // Example sending a request using UpdatePhoneNumberSettingsRequest.
//    req := client.UpdatePhoneNumberSettingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/UpdatePhoneNumberSettings
func (c *Client) UpdatePhoneNumberSettingsRequest(input *UpdatePhoneNumberSettingsInput) UpdatePhoneNumberSettingsRequest {
	op := &aws.Operation{
		Name:       opUpdatePhoneNumberSettings,
		HTTPMethod: "PUT",
		HTTPPath:   "/settings/phone-number",
	}

	if input == nil {
		input = &UpdatePhoneNumberSettingsInput{}
	}

	req := c.newRequest(op, input, &UpdatePhoneNumberSettingsOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UpdatePhoneNumberSettingsRequest{Request: req, Input: input, Copy: c.UpdatePhoneNumberSettingsRequest}
}

// UpdatePhoneNumberSettingsRequest is the request type for the
// UpdatePhoneNumberSettings API operation.
type UpdatePhoneNumberSettingsRequest struct {
	*aws.Request
	Input *UpdatePhoneNumberSettingsInput
	Copy  func(*UpdatePhoneNumberSettingsInput) UpdatePhoneNumberSettingsRequest
}

// Send marshals and sends the UpdatePhoneNumberSettings API request.
func (r UpdatePhoneNumberSettingsRequest) Send(ctx context.Context) (*UpdatePhoneNumberSettingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdatePhoneNumberSettingsResponse{
		UpdatePhoneNumberSettingsOutput: r.Request.Data.(*UpdatePhoneNumberSettingsOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdatePhoneNumberSettingsResponse is the response type for the
// UpdatePhoneNumberSettings API operation.
type UpdatePhoneNumberSettingsResponse struct {
	*UpdatePhoneNumberSettingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdatePhoneNumberSettings request.
func (r *UpdatePhoneNumberSettingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
