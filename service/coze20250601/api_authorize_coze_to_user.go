// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package coze20250601

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opAuthorizeCozeToUserCommon = "AuthorizeCozeToUser"

// AuthorizeCozeToUserCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the AuthorizeCozeToUserCommon operation. The "output" return
// value will be populated with the AuthorizeCozeToUserCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AuthorizeCozeToUserCommon Request to send the API call to the service.
// the "output" return value is not valid until after AuthorizeCozeToUserCommon Send returns without error.
//
// See AuthorizeCozeToUserCommon for more information on using the AuthorizeCozeToUserCommon
// API call, and error handling.
//
//    // Example sending a request using the AuthorizeCozeToUserCommonRequest method.
//    req, resp := client.AuthorizeCozeToUserCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *COZE20250601) AuthorizeCozeToUserCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAuthorizeCozeToUserCommon,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// AuthorizeCozeToUserCommon API operation for COZE20250601.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for COZE20250601's
// API operation AuthorizeCozeToUserCommon for usage and error information.
func (c *COZE20250601) AuthorizeCozeToUserCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AuthorizeCozeToUserCommonRequest(input)
	return out, req.Send()
}

// AuthorizeCozeToUserCommonWithContext is the same as AuthorizeCozeToUserCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AuthorizeCozeToUserCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *COZE20250601) AuthorizeCozeToUserCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AuthorizeCozeToUserCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAuthorizeCozeToUser = "AuthorizeCozeToUser"

// AuthorizeCozeToUserRequest generates a "volcengine/request.Request" representing the
// client's request for the AuthorizeCozeToUser operation. The "output" return
// value will be populated with the AuthorizeCozeToUserCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AuthorizeCozeToUserCommon Request to send the API call to the service.
// the "output" return value is not valid until after AuthorizeCozeToUserCommon Send returns without error.
//
// See AuthorizeCozeToUser for more information on using the AuthorizeCozeToUser
// API call, and error handling.
//
//    // Example sending a request using the AuthorizeCozeToUserRequest method.
//    req, resp := client.AuthorizeCozeToUserRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *COZE20250601) AuthorizeCozeToUserRequest(input *AuthorizeCozeToUserInput) (req *request.Request, output *AuthorizeCozeToUserOutput) {
	op := &request.Operation{
		Name:       opAuthorizeCozeToUser,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AuthorizeCozeToUserInput{}
	}

	output = &AuthorizeCozeToUserOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// AuthorizeCozeToUser API operation for COZE20250601.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for COZE20250601's
// API operation AuthorizeCozeToUser for usage and error information.
func (c *COZE20250601) AuthorizeCozeToUser(input *AuthorizeCozeToUserInput) (*AuthorizeCozeToUserOutput, error) {
	req, out := c.AuthorizeCozeToUserRequest(input)
	return out, req.Send()
}

// AuthorizeCozeToUserWithContext is the same as AuthorizeCozeToUser with the addition of
// the ability to pass a context and additional request options.
//
// See AuthorizeCozeToUser for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *COZE20250601) AuthorizeCozeToUserWithContext(ctx volcengine.Context, input *AuthorizeCozeToUserInput, opts ...request.Option) (*AuthorizeCozeToUserOutput, error) {
	req, out := c.AuthorizeCozeToUserRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AuthorizeCozeToUserInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// UserId is a required field
	UserId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s AuthorizeCozeToUserInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AuthorizeCozeToUserInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AuthorizeCozeToUserInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AuthorizeCozeToUserInput"}
	if s.UserId == nil {
		invalidParams.Add(request.NewErrParamRequired("UserId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetUserId sets the UserId field's value.
func (s *AuthorizeCozeToUserInput) SetUserId(v string) *AuthorizeCozeToUserInput {
	s.UserId = &v
	return s
}

type AuthorizeCozeToUserOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s AuthorizeCozeToUserOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AuthorizeCozeToUserOutput) GoString() string {
	return s.String()
}
