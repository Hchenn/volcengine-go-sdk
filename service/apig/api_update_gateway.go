// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package apig

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdateGatewayCommon = "UpdateGateway"

// UpdateGatewayCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateGatewayCommon operation. The "output" return
// value will be populated with the UpdateGatewayCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateGatewayCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateGatewayCommon Send returns without error.
//
// See UpdateGatewayCommon for more information on using the UpdateGatewayCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdateGatewayCommonRequest method.
//    req, resp := client.UpdateGatewayCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *APIG) UpdateGatewayCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateGatewayCommon,
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

// UpdateGatewayCommon API operation for APIG.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for APIG's
// API operation UpdateGatewayCommon for usage and error information.
func (c *APIG) UpdateGatewayCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateGatewayCommonRequest(input)
	return out, req.Send()
}

// UpdateGatewayCommonWithContext is the same as UpdateGatewayCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateGatewayCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *APIG) UpdateGatewayCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateGatewayCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateGateway = "UpdateGateway"

// UpdateGatewayRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateGateway operation. The "output" return
// value will be populated with the UpdateGatewayCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateGatewayCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateGatewayCommon Send returns without error.
//
// See UpdateGateway for more information on using the UpdateGateway
// API call, and error handling.
//
//    // Example sending a request using the UpdateGatewayRequest method.
//    req, resp := client.UpdateGatewayRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *APIG) UpdateGatewayRequest(input *UpdateGatewayInput) (req *request.Request, output *UpdateGatewayOutput) {
	op := &request.Operation{
		Name:       opUpdateGateway,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateGatewayInput{}
	}

	output = &UpdateGatewayOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdateGateway API operation for APIG.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for APIG's
// API operation UpdateGateway for usage and error information.
func (c *APIG) UpdateGateway(input *UpdateGatewayInput) (*UpdateGatewayOutput, error) {
	req, out := c.UpdateGatewayRequest(input)
	return out, req.Send()
}

// UpdateGatewayWithContext is the same as UpdateGateway with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateGateway for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *APIG) UpdateGatewayWithContext(ctx volcengine.Context, input *UpdateGatewayInput, opts ...request.Option) (*UpdateGatewayOutput, error) {
	req, out := c.UpdateGatewayRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type MonitorSpecForUpdateGatewayInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Enable *bool `type:"boolean" json:",omitempty"`

	WorkspaceID *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s MonitorSpecForUpdateGatewayInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s MonitorSpecForUpdateGatewayInput) GoString() string {
	return s.String()
}

// SetEnable sets the Enable field's value.
func (s *MonitorSpecForUpdateGatewayInput) SetEnable(v bool) *MonitorSpecForUpdateGatewayInput {
	s.Enable = &v
	return s
}

// SetWorkspaceID sets the WorkspaceID field's value.
func (s *MonitorSpecForUpdateGatewayInput) SetWorkspaceID(v string) *MonitorSpecForUpdateGatewayInput {
	s.WorkspaceID = &v
	return s
}

type UpdateGatewayInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Comments *string `type:"string" json:",omitempty"`

	// Id is a required field
	Id *string `type:"string" json:",omitempty" required:"true"`

	MonitorSpec *MonitorSpecForUpdateGatewayInput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s UpdateGatewayInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateGatewayInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateGatewayInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateGatewayInput"}
	if s.Id == nil {
		invalidParams.Add(request.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetComments sets the Comments field's value.
func (s *UpdateGatewayInput) SetComments(v string) *UpdateGatewayInput {
	s.Comments = &v
	return s
}

// SetId sets the Id field's value.
func (s *UpdateGatewayInput) SetId(v string) *UpdateGatewayInput {
	s.Id = &v
	return s
}

// SetMonitorSpec sets the MonitorSpec field's value.
func (s *UpdateGatewayInput) SetMonitorSpec(v *MonitorSpecForUpdateGatewayInput) *UpdateGatewayInput {
	s.MonitorSpec = v
	return s
}

type UpdateGatewayOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Id *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s UpdateGatewayOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateGatewayOutput) GoString() string {
	return s.String()
}

// SetId sets the Id field's value.
func (s *UpdateGatewayOutput) SetId(v string) *UpdateGatewayOutput {
	s.Id = &v
	return s
}
