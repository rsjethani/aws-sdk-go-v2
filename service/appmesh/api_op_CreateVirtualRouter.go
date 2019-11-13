// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appmesh

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateVirtualRouterInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `locationName:"clientToken" type:"string" idempotencyToken:"true"`

	// MeshName is a required field
	MeshName *string `location:"uri" locationName:"meshName" min:"1" type:"string" required:"true"`

	// An object that represents the specification of a virtual router.
	//
	// Spec is a required field
	Spec *VirtualRouterSpec `locationName:"spec" type:"structure" required:"true"`

	Tags []TagRef `locationName:"tags" type:"list"`

	// VirtualRouterName is a required field
	VirtualRouterName *string `locationName:"virtualRouterName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateVirtualRouterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateVirtualRouterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateVirtualRouterInput"}

	if s.MeshName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MeshName"))
	}
	if s.MeshName != nil && len(*s.MeshName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MeshName", 1))
	}

	if s.Spec == nil {
		invalidParams.Add(aws.NewErrParamRequired("Spec"))
	}

	if s.VirtualRouterName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VirtualRouterName"))
	}
	if s.VirtualRouterName != nil && len(*s.VirtualRouterName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VirtualRouterName", 1))
	}
	if s.Spec != nil {
		if err := s.Spec.Validate(); err != nil {
			invalidParams.AddNested("Spec", err.(aws.ErrInvalidParams))
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateVirtualRouterInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	var ClientToken string
	if s.ClientToken != nil {
		ClientToken = *s.ClientToken
	} else {
		ClientToken = protocol.GetIdempotencyToken()
	}
	{
		v := ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Spec != nil {
		v := s.Spec

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "spec", v, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "tags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.VirtualRouterName != nil {
		v := *s.VirtualRouterName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "virtualRouterName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MeshName != nil {
		v := *s.MeshName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "meshName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateVirtualRouterOutput struct {
	_ struct{} `type:"structure" payload:"VirtualRouter"`

	// An object that represents a virtual router returned by a describe operation.
	//
	// VirtualRouter is a required field
	VirtualRouter *VirtualRouterData `locationName:"virtualRouter" type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateVirtualRouterOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateVirtualRouterOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.VirtualRouter != nil {
		v := s.VirtualRouter

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "virtualRouter", v, metadata)
	}
	return nil
}

const opCreateVirtualRouter = "CreateVirtualRouter"

// CreateVirtualRouterRequest returns a request value for making API operation for
// AWS App Mesh.
//
// Creates a virtual router within a service mesh.
//
// Any inbound traffic that your virtual router expects should be specified
// as a listener.
//
// Virtual routers handle traffic for one or more virtual services within your
// mesh. After you create your virtual router, create and associate routes for
// your virtual router that direct incoming requests to different virtual nodes.
//
//    // Example sending a request using CreateVirtualRouterRequest.
//    req := client.CreateVirtualRouterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/CreateVirtualRouter
func (c *Client) CreateVirtualRouterRequest(input *CreateVirtualRouterInput) CreateVirtualRouterRequest {
	op := &aws.Operation{
		Name:       opCreateVirtualRouter,
		HTTPMethod: "PUT",
		HTTPPath:   "/v20190125/meshes/{meshName}/virtualRouters",
	}

	if input == nil {
		input = &CreateVirtualRouterInput{}
	}

	req := c.newRequest(op, input, &CreateVirtualRouterOutput{})
	return CreateVirtualRouterRequest{Request: req, Input: input, Copy: c.CreateVirtualRouterRequest}
}

// CreateVirtualRouterRequest is the request type for the
// CreateVirtualRouter API operation.
type CreateVirtualRouterRequest struct {
	*aws.Request
	Input *CreateVirtualRouterInput
	Copy  func(*CreateVirtualRouterInput) CreateVirtualRouterRequest
}

// Send marshals and sends the CreateVirtualRouter API request.
func (r CreateVirtualRouterRequest) Send(ctx context.Context) (*CreateVirtualRouterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateVirtualRouterResponse{
		CreateVirtualRouterOutput: r.Request.Data.(*CreateVirtualRouterOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateVirtualRouterResponse is the response type for the
// CreateVirtualRouter API operation.
type CreateVirtualRouterResponse struct {
	*CreateVirtualRouterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateVirtualRouter request.
func (r *CreateVirtualRouterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
