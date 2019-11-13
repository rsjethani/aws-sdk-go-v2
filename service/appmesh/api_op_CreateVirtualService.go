// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appmesh

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateVirtualServiceInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `locationName:"clientToken" type:"string" idempotencyToken:"true"`

	// MeshName is a required field
	MeshName *string `location:"uri" locationName:"meshName" min:"1" type:"string" required:"true"`

	// An object that represents the specification of a virtual service.
	//
	// Spec is a required field
	Spec *VirtualServiceSpec `locationName:"spec" type:"structure" required:"true"`

	Tags []TagRef `locationName:"tags" type:"list"`

	// VirtualServiceName is a required field
	VirtualServiceName *string `locationName:"virtualServiceName" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateVirtualServiceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateVirtualServiceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateVirtualServiceInput"}

	if s.MeshName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MeshName"))
	}
	if s.MeshName != nil && len(*s.MeshName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MeshName", 1))
	}

	if s.Spec == nil {
		invalidParams.Add(aws.NewErrParamRequired("Spec"))
	}

	if s.VirtualServiceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VirtualServiceName"))
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
func (s CreateVirtualServiceInput) MarshalFields(e protocol.FieldEncoder) error {
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
	if s.VirtualServiceName != nil {
		v := *s.VirtualServiceName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "virtualServiceName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MeshName != nil {
		v := *s.MeshName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "meshName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateVirtualServiceOutput struct {
	_ struct{} `type:"structure" payload:"VirtualService"`

	// An object that represents a virtual service returned by a describe operation.
	//
	// VirtualService is a required field
	VirtualService *VirtualServiceData `locationName:"virtualService" type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateVirtualServiceOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateVirtualServiceOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.VirtualService != nil {
		v := s.VirtualService

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "virtualService", v, metadata)
	}
	return nil
}

const opCreateVirtualService = "CreateVirtualService"

// CreateVirtualServiceRequest returns a request value for making API operation for
// AWS App Mesh.
//
// Creates a virtual service within a service mesh.
//
// A virtual service is an abstraction of a real service that is provided by
// a virtual node directly or indirectly by means of a virtual router. Dependent
// services call your virtual service by its virtualServiceName, and those requests
// are routed to the virtual node or virtual router that is specified as the
// provider for the virtual service.
//
//    // Example sending a request using CreateVirtualServiceRequest.
//    req := client.CreateVirtualServiceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/CreateVirtualService
func (c *Client) CreateVirtualServiceRequest(input *CreateVirtualServiceInput) CreateVirtualServiceRequest {
	op := &aws.Operation{
		Name:       opCreateVirtualService,
		HTTPMethod: "PUT",
		HTTPPath:   "/v20190125/meshes/{meshName}/virtualServices",
	}

	if input == nil {
		input = &CreateVirtualServiceInput{}
	}

	req := c.newRequest(op, input, &CreateVirtualServiceOutput{})
	return CreateVirtualServiceRequest{Request: req, Input: input, Copy: c.CreateVirtualServiceRequest}
}

// CreateVirtualServiceRequest is the request type for the
// CreateVirtualService API operation.
type CreateVirtualServiceRequest struct {
	*aws.Request
	Input *CreateVirtualServiceInput
	Copy  func(*CreateVirtualServiceInput) CreateVirtualServiceRequest
}

// Send marshals and sends the CreateVirtualService API request.
func (r CreateVirtualServiceRequest) Send(ctx context.Context) (*CreateVirtualServiceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateVirtualServiceResponse{
		CreateVirtualServiceOutput: r.Request.Data.(*CreateVirtualServiceOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateVirtualServiceResponse is the response type for the
// CreateVirtualService API operation.
type CreateVirtualServiceResponse struct {
	*CreateVirtualServiceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateVirtualService request.
func (r *CreateVirtualServiceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
