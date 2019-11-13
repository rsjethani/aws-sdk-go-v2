// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package connect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListHoursOfOperationsInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the Amazon Connect instance.
	//
	// InstanceId is a required field
	InstanceId *string `location:"uri" locationName:"InstanceId" min:"1" type:"string" required:"true"`

	// The maximimum number of results to return per page.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListHoursOfOperationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListHoursOfOperationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListHoursOfOperationsInput"}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}
	if s.InstanceId != nil && len(*s.InstanceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InstanceId", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListHoursOfOperationsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.InstanceId != nil {
		v := *s.InstanceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "InstanceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListHoursOfOperationsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the hours of operation.
	HoursOfOperationSummaryList []HoursOfOperationSummary `type:"list"`

	// If there are additional results, this is the token for the next set of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListHoursOfOperationsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListHoursOfOperationsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.HoursOfOperationSummaryList != nil {
		v := s.HoursOfOperationSummaryList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "HoursOfOperationSummaryList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListHoursOfOperations = "ListHoursOfOperations"

// ListHoursOfOperationsRequest returns a request value for making API operation for
// Amazon Connect Service.
//
// Provides information about the hours of operation for the specified Amazon
// Connect instance.
//
//    // Example sending a request using ListHoursOfOperationsRequest.
//    req := client.ListHoursOfOperationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/connect-2017-08-08/ListHoursOfOperations
func (c *Client) ListHoursOfOperationsRequest(input *ListHoursOfOperationsInput) ListHoursOfOperationsRequest {
	op := &aws.Operation{
		Name:       opListHoursOfOperations,
		HTTPMethod: "GET",
		HTTPPath:   "/hours-of-operations-summary/{InstanceId}",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListHoursOfOperationsInput{}
	}

	req := c.newRequest(op, input, &ListHoursOfOperationsOutput{})
	return ListHoursOfOperationsRequest{Request: req, Input: input, Copy: c.ListHoursOfOperationsRequest}
}

// ListHoursOfOperationsRequest is the request type for the
// ListHoursOfOperations API operation.
type ListHoursOfOperationsRequest struct {
	*aws.Request
	Input *ListHoursOfOperationsInput
	Copy  func(*ListHoursOfOperationsInput) ListHoursOfOperationsRequest
}

// Send marshals and sends the ListHoursOfOperations API request.
func (r ListHoursOfOperationsRequest) Send(ctx context.Context) (*ListHoursOfOperationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListHoursOfOperationsResponse{
		ListHoursOfOperationsOutput: r.Request.Data.(*ListHoursOfOperationsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListHoursOfOperationsRequestPaginator returns a paginator for ListHoursOfOperations.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListHoursOfOperationsRequest(input)
//   p := connect.NewListHoursOfOperationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListHoursOfOperationsPaginator(req ListHoursOfOperationsRequest) ListHoursOfOperationsPaginator {
	return ListHoursOfOperationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListHoursOfOperationsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListHoursOfOperationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListHoursOfOperationsPaginator struct {
	aws.Pager
}

func (p *ListHoursOfOperationsPaginator) CurrentPage() *ListHoursOfOperationsOutput {
	return p.Pager.CurrentPage().(*ListHoursOfOperationsOutput)
}

// ListHoursOfOperationsResponse is the response type for the
// ListHoursOfOperations API operation.
type ListHoursOfOperationsResponse struct {
	*ListHoursOfOperationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListHoursOfOperations request.
func (r *ListHoursOfOperationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
