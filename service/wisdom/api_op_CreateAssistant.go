// Code generated by smithy-go-codegen DO NOT EDIT.

package wisdom

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wisdom/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Amazon Connect Wisdom assistant.
func (c *Client) CreateAssistant(ctx context.Context, params *CreateAssistantInput, optFns ...func(*Options)) (*CreateAssistantOutput, error) {
	if params == nil {
		params = &CreateAssistantInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAssistant", params, optFns, c.addOperationCreateAssistantMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAssistantOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAssistantInput struct {

	// The name of the assistant.
	//
	// This member is required.
	Name *string

	// The type of assistant.
	//
	// This member is required.
	Type types.AssistantType

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string

	// The description of the assistant.
	Description *string

	// The KMS key used for encryption.
	ServerSideEncryptionConfiguration *types.ServerSideEncryptionConfiguration

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateAssistantOutput struct {

	// Information about the assistant.
	Assistant *types.AssistantData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAssistantMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAssistant{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAssistant{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateAssistantValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAssistant(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateAssistant(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wisdom",
		OperationName: "CreateAssistant",
	}
}
