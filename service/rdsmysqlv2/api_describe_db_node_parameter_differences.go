// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysqlv2

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeDBNodeParameterDifferencesCommon = "DescribeDBNodeParameterDifferences"

// DescribeDBNodeParameterDifferencesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDBNodeParameterDifferencesCommon operation. The "output" return
// value will be populated with the DescribeDBNodeParameterDifferencesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDBNodeParameterDifferencesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDBNodeParameterDifferencesCommon Send returns without error.
//
// See DescribeDBNodeParameterDifferencesCommon for more information on using the DescribeDBNodeParameterDifferencesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeDBNodeParameterDifferencesCommonRequest method.
//    req, resp := client.DescribeDBNodeParameterDifferencesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) DescribeDBNodeParameterDifferencesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeDBNodeParameterDifferencesCommon,
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

// DescribeDBNodeParameterDifferencesCommon API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation DescribeDBNodeParameterDifferencesCommon for usage and error information.
func (c *RDSMYSQLV2) DescribeDBNodeParameterDifferencesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeDBNodeParameterDifferencesCommonRequest(input)
	return out, req.Send()
}

// DescribeDBNodeParameterDifferencesCommonWithContext is the same as DescribeDBNodeParameterDifferencesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDBNodeParameterDifferencesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) DescribeDBNodeParameterDifferencesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeDBNodeParameterDifferencesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeDBNodeParameterDifferences = "DescribeDBNodeParameterDifferences"

// DescribeDBNodeParameterDifferencesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDBNodeParameterDifferences operation. The "output" return
// value will be populated with the DescribeDBNodeParameterDifferencesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDBNodeParameterDifferencesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDBNodeParameterDifferencesCommon Send returns without error.
//
// See DescribeDBNodeParameterDifferences for more information on using the DescribeDBNodeParameterDifferences
// API call, and error handling.
//
//    // Example sending a request using the DescribeDBNodeParameterDifferencesRequest method.
//    req, resp := client.DescribeDBNodeParameterDifferencesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) DescribeDBNodeParameterDifferencesRequest(input *DescribeDBNodeParameterDifferencesInput) (req *request.Request, output *DescribeDBNodeParameterDifferencesOutput) {
	op := &request.Operation{
		Name:       opDescribeDBNodeParameterDifferences,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDBNodeParameterDifferencesInput{}
	}

	output = &DescribeDBNodeParameterDifferencesOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeDBNodeParameterDifferences API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation DescribeDBNodeParameterDifferences for usage and error information.
func (c *RDSMYSQLV2) DescribeDBNodeParameterDifferences(input *DescribeDBNodeParameterDifferencesInput) (*DescribeDBNodeParameterDifferencesOutput, error) {
	req, out := c.DescribeDBNodeParameterDifferencesRequest(input)
	return out, req.Send()
}

// DescribeDBNodeParameterDifferencesWithContext is the same as DescribeDBNodeParameterDifferences with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDBNodeParameterDifferences for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) DescribeDBNodeParameterDifferencesWithContext(ctx volcengine.Context, input *DescribeDBNodeParameterDifferencesInput, opts ...request.Option) (*DescribeDBNodeParameterDifferencesOutput, error) {
	req, out := c.DescribeDBNodeParameterDifferencesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeDBNodeParameterDifferencesInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DescribeDBNodeParameterDifferencesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDBNodeParameterDifferencesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDBNodeParameterDifferencesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeDBNodeParameterDifferencesInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeDBNodeParameterDifferencesInput) SetInstanceId(v string) *DescribeDBNodeParameterDifferencesInput {
	s.InstanceId = &v
	return s
}

type DescribeDBNodeParameterDifferencesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	NodeDifferentParameters []*NodeDifferentParameterForDescribeDBNodeParameterDifferencesOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DescribeDBNodeParameterDifferencesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDBNodeParameterDifferencesOutput) GoString() string {
	return s.String()
}

// SetNodeDifferentParameters sets the NodeDifferentParameters field's value.
func (s *DescribeDBNodeParameterDifferencesOutput) SetNodeDifferentParameters(v []*NodeDifferentParameterForDescribeDBNodeParameterDifferencesOutput) *DescribeDBNodeParameterDifferencesOutput {
	s.NodeDifferentParameters = v
	return s
}

type NodeDifferentParameterForDescribeDBNodeParameterDifferencesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	NodeParameterValueDetails []*NodeParameterValueDetailForDescribeDBNodeParameterDifferencesOutput `type:"list" json:",omitempty"`

	ParameterName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s NodeDifferentParameterForDescribeDBNodeParameterDifferencesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s NodeDifferentParameterForDescribeDBNodeParameterDifferencesOutput) GoString() string {
	return s.String()
}

// SetNodeParameterValueDetails sets the NodeParameterValueDetails field's value.
func (s *NodeDifferentParameterForDescribeDBNodeParameterDifferencesOutput) SetNodeParameterValueDetails(v []*NodeParameterValueDetailForDescribeDBNodeParameterDifferencesOutput) *NodeDifferentParameterForDescribeDBNodeParameterDifferencesOutput {
	s.NodeParameterValueDetails = v
	return s
}

// SetParameterName sets the ParameterName field's value.
func (s *NodeDifferentParameterForDescribeDBNodeParameterDifferencesOutput) SetParameterName(v string) *NodeDifferentParameterForDescribeDBNodeParameterDifferencesOutput {
	s.ParameterName = &v
	return s
}

type NodeParameterValueDetailForDescribeDBNodeParameterDifferencesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Expression *string `type:"string" json:",omitempty"`

	NodeId *string `type:"string" json:",omitempty"`

	NodeSpec *string `type:"string" json:",omitempty"`

	NodeType *string `type:"string" json:",omitempty"`

	ParameterValue *string `type:"string" json:",omitempty"`

	ZoneId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s NodeParameterValueDetailForDescribeDBNodeParameterDifferencesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s NodeParameterValueDetailForDescribeDBNodeParameterDifferencesOutput) GoString() string {
	return s.String()
}

// SetExpression sets the Expression field's value.
func (s *NodeParameterValueDetailForDescribeDBNodeParameterDifferencesOutput) SetExpression(v string) *NodeParameterValueDetailForDescribeDBNodeParameterDifferencesOutput {
	s.Expression = &v
	return s
}

// SetNodeId sets the NodeId field's value.
func (s *NodeParameterValueDetailForDescribeDBNodeParameterDifferencesOutput) SetNodeId(v string) *NodeParameterValueDetailForDescribeDBNodeParameterDifferencesOutput {
	s.NodeId = &v
	return s
}

// SetNodeSpec sets the NodeSpec field's value.
func (s *NodeParameterValueDetailForDescribeDBNodeParameterDifferencesOutput) SetNodeSpec(v string) *NodeParameterValueDetailForDescribeDBNodeParameterDifferencesOutput {
	s.NodeSpec = &v
	return s
}

// SetNodeType sets the NodeType field's value.
func (s *NodeParameterValueDetailForDescribeDBNodeParameterDifferencesOutput) SetNodeType(v string) *NodeParameterValueDetailForDescribeDBNodeParameterDifferencesOutput {
	s.NodeType = &v
	return s
}

// SetParameterValue sets the ParameterValue field's value.
func (s *NodeParameterValueDetailForDescribeDBNodeParameterDifferencesOutput) SetParameterValue(v string) *NodeParameterValueDetailForDescribeDBNodeParameterDifferencesOutput {
	s.ParameterValue = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *NodeParameterValueDetailForDescribeDBNodeParameterDifferencesOutput) SetZoneId(v string) *NodeParameterValueDetailForDescribeDBNodeParameterDifferencesOutput {
	s.ZoneId = &v
	return s
}
