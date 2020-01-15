// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type EnableKeyInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for the customer master key (CMK).
	//
	// Specify the key ID or the Amazon Resource Name (ARN) of the CMK.
	//
	// For example:
	//
	//    * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//    * Key ARN: arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// To get the key ID and key ARN for a CMK, use ListKeys or DescribeKey.
	//
	// KeyId is a required field
	KeyId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s EnableKeyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EnableKeyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "EnableKeyInput"}

	if s.KeyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("KeyId"))
	}
	if s.KeyId != nil && len(*s.KeyId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("KeyId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type EnableKeyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s EnableKeyOutput) String() string {
	return awsutil.Prettify(s)
}

const opEnableKey = "EnableKey"

// EnableKeyRequest returns a request value for making API operation for
// AWS Key Management Service.
//
// Sets the key state of a customer master key (CMK) to enabled. This allows
// you to use the CMK for cryptographic operations. You cannot perform this
// operation on a CMK in a different AWS account.
//
// The CMK that you use for this operation must be in a compatible key state.
// For details, see How Key State Affects Use of a Customer Master Key (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html)
// in the AWS Key Management Service Developer Guide.
//
//    // Example sending a request using EnableKeyRequest.
//    req := client.EnableKeyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/EnableKey
func (c *Client) EnableKeyRequest(input *EnableKeyInput) EnableKeyRequest {
	op := &aws.Operation{
		Name:       opEnableKey,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &EnableKeyInput{}
	}

	req := c.newRequest(op, input, &EnableKeyOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return EnableKeyRequest{Request: req, Input: input, Copy: c.EnableKeyRequest}
}

// EnableKeyRequest is the request type for the
// EnableKey API operation.
type EnableKeyRequest struct {
	*aws.Request
	Input *EnableKeyInput
	Copy  func(*EnableKeyInput) EnableKeyRequest
}

// Send marshals and sends the EnableKey API request.
func (r EnableKeyRequest) Send(ctx context.Context) (*EnableKeyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &EnableKeyResponse{
		EnableKeyOutput: r.Request.Data.(*EnableKeyOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// EnableKeyResponse is the response type for the
// EnableKey API operation.
type EnableKeyResponse struct {
	*EnableKeyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// EnableKey request.
func (r *EnableKeyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}