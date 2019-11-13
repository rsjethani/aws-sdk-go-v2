// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codestarnotifications

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateNotificationRuleInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the notification rule.
	//
	// Arn is a required field
	Arn *string `type:"string" required:"true"`

	// The level of detail to include in the notifications for this resource. BASIC
	// will include only the contents of the event as it would appear in AWS CloudWatch.
	// FULL will include any supplemental information provided by AWS CodeStar Notifications
	// and/or the service for the resource for which the notification is created.
	DetailType DetailType `type:"string" enum:"true"`

	// A list of event types associated with this notification rule.
	EventTypeIds []string `type:"list"`

	// The name of the notification rule.
	Name *string `min:"1" type:"string" sensitive:"true"`

	// The status of the notification rule. Valid statuses include enabled (sending
	// notifications) or disabled (not sending notifications).
	Status NotificationRuleStatus `type:"string" enum:"true"`

	// The address and type of the targets to receive notifications from this notification
	// rule.
	Targets []Target `type:"list"`
}

// String returns the string representation
func (s UpdateNotificationRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateNotificationRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateNotificationRuleInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.Targets != nil {
		for i, v := range s.Targets {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Targets", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateNotificationRuleInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.DetailType) > 0 {
		v := s.DetailType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DetailType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.EventTypeIds != nil {
		v := s.EventTypeIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "EventTypeIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Targets != nil {
		v := s.Targets

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Targets", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

type UpdateNotificationRuleOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateNotificationRuleOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateNotificationRuleOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateNotificationRule = "UpdateNotificationRule"

// UpdateNotificationRuleRequest returns a request value for making API operation for
// AWS CodeStar Notifications.
//
// Updates a notification rule for a resource. You can change the events that
// trigger the notification rule, the status of the rule, and the targets that
// receive the notifications.
//
// To add or remove tags for a notification rule, you must use TagResource and
// UntagResource.
//
//    // Example sending a request using UpdateNotificationRuleRequest.
//    req := client.UpdateNotificationRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codestar-notifications-2019-10-15/UpdateNotificationRule
func (c *Client) UpdateNotificationRuleRequest(input *UpdateNotificationRuleInput) UpdateNotificationRuleRequest {
	op := &aws.Operation{
		Name:       opUpdateNotificationRule,
		HTTPMethod: "POST",
		HTTPPath:   "/updateNotificationRule",
	}

	if input == nil {
		input = &UpdateNotificationRuleInput{}
	}

	req := c.newRequest(op, input, &UpdateNotificationRuleOutput{})
	return UpdateNotificationRuleRequest{Request: req, Input: input, Copy: c.UpdateNotificationRuleRequest}
}

// UpdateNotificationRuleRequest is the request type for the
// UpdateNotificationRule API operation.
type UpdateNotificationRuleRequest struct {
	*aws.Request
	Input *UpdateNotificationRuleInput
	Copy  func(*UpdateNotificationRuleInput) UpdateNotificationRuleRequest
}

// Send marshals and sends the UpdateNotificationRule API request.
func (r UpdateNotificationRuleRequest) Send(ctx context.Context) (*UpdateNotificationRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateNotificationRuleResponse{
		UpdateNotificationRuleOutput: r.Request.Data.(*UpdateNotificationRuleOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateNotificationRuleResponse is the response type for the
// UpdateNotificationRule API operation.
type UpdateNotificationRuleResponse struct {
	*UpdateNotificationRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateNotificationRule request.
func (r *UpdateNotificationRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
