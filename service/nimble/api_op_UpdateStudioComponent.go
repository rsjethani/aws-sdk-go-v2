// Code generated by smithy-go-codegen DO NOT EDIT.

package nimble

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/nimble/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a studio component resource.
func (c *Client) UpdateStudioComponent(ctx context.Context, params *UpdateStudioComponentInput, optFns ...func(*Options)) (*UpdateStudioComponentOutput, error) {
	if params == nil {
		params = &UpdateStudioComponentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateStudioComponent", params, optFns, c.addOperationUpdateStudioComponentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateStudioComponentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type UpdateStudioComponentInput struct {

	// The studio component ID.
	//
	// This member is required.
	StudioComponentId *string

	// The studio ID.
	//
	// This member is required.
	StudioId *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. If you don’t specify a client token, the AWS SDK automatically
	// generates a client token and uses it for the request to ensure idempotency.
	ClientToken *string

	// The configuration of the studio component, based on component type.
	Configuration *types.StudioComponentConfiguration

	// The description.
	Description *string

	// The EC2 security groups that control access to the studio component.
	Ec2SecurityGroupIds []string

	// Initialization scripts for studio components.
	InitializationScripts []types.StudioComponentInitializationScript

	// The name for the studio component.
	Name *string

	// Parameters for the studio component scripts.
	ScriptParameters []types.ScriptParameterKeyValue

	// The specific subtype of a studio component.
	Subtype types.StudioComponentSubtype

	// The type of the studio component.
	Type types.StudioComponentType

	noSmithyDocumentSerde
}

//
type UpdateStudioComponentOutput struct {

	// Information about the studio component.
	StudioComponent *types.StudioComponent

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateStudioComponentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateStudioComponent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateStudioComponent{}, middleware.After)
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
	if err = addOpUpdateStudioComponentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateStudioComponent(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateStudioComponent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "nimble",
		OperationName: "UpdateStudioComponent",
	}
}
