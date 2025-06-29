// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysqlv2

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGrantDBAccountPrivilegeCommon = "GrantDBAccountPrivilege"

// GrantDBAccountPrivilegeCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GrantDBAccountPrivilegeCommon operation. The "output" return
// value will be populated with the GrantDBAccountPrivilegeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GrantDBAccountPrivilegeCommon Request to send the API call to the service.
// the "output" return value is not valid until after GrantDBAccountPrivilegeCommon Send returns without error.
//
// See GrantDBAccountPrivilegeCommon for more information on using the GrantDBAccountPrivilegeCommon
// API call, and error handling.
//
//    // Example sending a request using the GrantDBAccountPrivilegeCommonRequest method.
//    req, resp := client.GrantDBAccountPrivilegeCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) GrantDBAccountPrivilegeCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGrantDBAccountPrivilegeCommon,
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

// GrantDBAccountPrivilegeCommon API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation GrantDBAccountPrivilegeCommon for usage and error information.
func (c *RDSMYSQLV2) GrantDBAccountPrivilegeCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GrantDBAccountPrivilegeCommonRequest(input)
	return out, req.Send()
}

// GrantDBAccountPrivilegeCommonWithContext is the same as GrantDBAccountPrivilegeCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GrantDBAccountPrivilegeCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) GrantDBAccountPrivilegeCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GrantDBAccountPrivilegeCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGrantDBAccountPrivilege = "GrantDBAccountPrivilege"

// GrantDBAccountPrivilegeRequest generates a "volcengine/request.Request" representing the
// client's request for the GrantDBAccountPrivilege operation. The "output" return
// value will be populated with the GrantDBAccountPrivilegeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GrantDBAccountPrivilegeCommon Request to send the API call to the service.
// the "output" return value is not valid until after GrantDBAccountPrivilegeCommon Send returns without error.
//
// See GrantDBAccountPrivilege for more information on using the GrantDBAccountPrivilege
// API call, and error handling.
//
//    // Example sending a request using the GrantDBAccountPrivilegeRequest method.
//    req, resp := client.GrantDBAccountPrivilegeRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) GrantDBAccountPrivilegeRequest(input *GrantDBAccountPrivilegeInput) (req *request.Request, output *GrantDBAccountPrivilegeOutput) {
	op := &request.Operation{
		Name:       opGrantDBAccountPrivilege,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GrantDBAccountPrivilegeInput{}
	}

	output = &GrantDBAccountPrivilegeOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GrantDBAccountPrivilege API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation GrantDBAccountPrivilege for usage and error information.
func (c *RDSMYSQLV2) GrantDBAccountPrivilege(input *GrantDBAccountPrivilegeInput) (*GrantDBAccountPrivilegeOutput, error) {
	req, out := c.GrantDBAccountPrivilegeRequest(input)
	return out, req.Send()
}

// GrantDBAccountPrivilegeWithContext is the same as GrantDBAccountPrivilege with the addition of
// the ability to pass a context and additional request options.
//
// See GrantDBAccountPrivilege for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) GrantDBAccountPrivilegeWithContext(ctx volcengine.Context, input *GrantDBAccountPrivilegeInput, opts ...request.Option) (*GrantDBAccountPrivilegeOutput, error) {
	req, out := c.GrantDBAccountPrivilegeRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AccountPrivilegeForGrantDBAccountPrivilegeInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AccountPrivilege *string `type:"string" json:",omitempty"`

	AccountPrivilegeDetail *string `type:"string" json:",omitempty"`

	DBName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s AccountPrivilegeForGrantDBAccountPrivilegeInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AccountPrivilegeForGrantDBAccountPrivilegeInput) GoString() string {
	return s.String()
}

// SetAccountPrivilege sets the AccountPrivilege field's value.
func (s *AccountPrivilegeForGrantDBAccountPrivilegeInput) SetAccountPrivilege(v string) *AccountPrivilegeForGrantDBAccountPrivilegeInput {
	s.AccountPrivilege = &v
	return s
}

// SetAccountPrivilegeDetail sets the AccountPrivilegeDetail field's value.
func (s *AccountPrivilegeForGrantDBAccountPrivilegeInput) SetAccountPrivilegeDetail(v string) *AccountPrivilegeForGrantDBAccountPrivilegeInput {
	s.AccountPrivilegeDetail = &v
	return s
}

// SetDBName sets the DBName field's value.
func (s *AccountPrivilegeForGrantDBAccountPrivilegeInput) SetDBName(v string) *AccountPrivilegeForGrantDBAccountPrivilegeInput {
	s.DBName = &v
	return s
}

type ColumnPrivilegeForGrantDBAccountPrivilegeInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AccountPrivilegeDetail *string `type:"string" json:",omitempty"`

	ColumnName *string `type:"string" json:",omitempty"`

	TableName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ColumnPrivilegeForGrantDBAccountPrivilegeInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ColumnPrivilegeForGrantDBAccountPrivilegeInput) GoString() string {
	return s.String()
}

// SetAccountPrivilegeDetail sets the AccountPrivilegeDetail field's value.
func (s *ColumnPrivilegeForGrantDBAccountPrivilegeInput) SetAccountPrivilegeDetail(v string) *ColumnPrivilegeForGrantDBAccountPrivilegeInput {
	s.AccountPrivilegeDetail = &v
	return s
}

// SetColumnName sets the ColumnName field's value.
func (s *ColumnPrivilegeForGrantDBAccountPrivilegeInput) SetColumnName(v string) *ColumnPrivilegeForGrantDBAccountPrivilegeInput {
	s.ColumnName = &v
	return s
}

// SetTableName sets the TableName field's value.
func (s *ColumnPrivilegeForGrantDBAccountPrivilegeInput) SetTableName(v string) *ColumnPrivilegeForGrantDBAccountPrivilegeInput {
	s.TableName = &v
	return s
}

type GrantDBAccountPrivilegeInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// AccountName is a required field
	AccountName *string `type:"string" json:",omitempty" required:"true"`

	AccountPrivileges []*AccountPrivilegeForGrantDBAccountPrivilegeInput `type:"list" json:",omitempty"`

	DryRun *bool `type:"boolean" json:",omitempty"`

	Host *string `type:"string" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`

	TableColumnPrivileges []*TableColumnPrivilegeForGrantDBAccountPrivilegeInput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s GrantDBAccountPrivilegeInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GrantDBAccountPrivilegeInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GrantDBAccountPrivilegeInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GrantDBAccountPrivilegeInput"}
	if s.AccountName == nil {
		invalidParams.Add(request.NewErrParamRequired("AccountName"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAccountName sets the AccountName field's value.
func (s *GrantDBAccountPrivilegeInput) SetAccountName(v string) *GrantDBAccountPrivilegeInput {
	s.AccountName = &v
	return s
}

// SetAccountPrivileges sets the AccountPrivileges field's value.
func (s *GrantDBAccountPrivilegeInput) SetAccountPrivileges(v []*AccountPrivilegeForGrantDBAccountPrivilegeInput) *GrantDBAccountPrivilegeInput {
	s.AccountPrivileges = v
	return s
}

// SetDryRun sets the DryRun field's value.
func (s *GrantDBAccountPrivilegeInput) SetDryRun(v bool) *GrantDBAccountPrivilegeInput {
	s.DryRun = &v
	return s
}

// SetHost sets the Host field's value.
func (s *GrantDBAccountPrivilegeInput) SetHost(v string) *GrantDBAccountPrivilegeInput {
	s.Host = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *GrantDBAccountPrivilegeInput) SetInstanceId(v string) *GrantDBAccountPrivilegeInput {
	s.InstanceId = &v
	return s
}

// SetTableColumnPrivileges sets the TableColumnPrivileges field's value.
func (s *GrantDBAccountPrivilegeInput) SetTableColumnPrivileges(v []*TableColumnPrivilegeForGrantDBAccountPrivilegeInput) *GrantDBAccountPrivilegeInput {
	s.TableColumnPrivileges = v
	return s
}

type GrantDBAccountPrivilegeOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	GrantPrivilegeSQL []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s GrantDBAccountPrivilegeOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GrantDBAccountPrivilegeOutput) GoString() string {
	return s.String()
}

// SetGrantPrivilegeSQL sets the GrantPrivilegeSQL field's value.
func (s *GrantDBAccountPrivilegeOutput) SetGrantPrivilegeSQL(v []*string) *GrantDBAccountPrivilegeOutput {
	s.GrantPrivilegeSQL = v
	return s
}

type TableColumnPrivilegeForGrantDBAccountPrivilegeInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ColumnPrivileges []*ColumnPrivilegeForGrantDBAccountPrivilegeInput `type:"list" json:",omitempty"`

	DBName *string `type:"string" json:",omitempty"`

	TablePrivileges []*TablePrivilegeForGrantDBAccountPrivilegeInput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s TableColumnPrivilegeForGrantDBAccountPrivilegeInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TableColumnPrivilegeForGrantDBAccountPrivilegeInput) GoString() string {
	return s.String()
}

// SetColumnPrivileges sets the ColumnPrivileges field's value.
func (s *TableColumnPrivilegeForGrantDBAccountPrivilegeInput) SetColumnPrivileges(v []*ColumnPrivilegeForGrantDBAccountPrivilegeInput) *TableColumnPrivilegeForGrantDBAccountPrivilegeInput {
	s.ColumnPrivileges = v
	return s
}

// SetDBName sets the DBName field's value.
func (s *TableColumnPrivilegeForGrantDBAccountPrivilegeInput) SetDBName(v string) *TableColumnPrivilegeForGrantDBAccountPrivilegeInput {
	s.DBName = &v
	return s
}

// SetTablePrivileges sets the TablePrivileges field's value.
func (s *TableColumnPrivilegeForGrantDBAccountPrivilegeInput) SetTablePrivileges(v []*TablePrivilegeForGrantDBAccountPrivilegeInput) *TableColumnPrivilegeForGrantDBAccountPrivilegeInput {
	s.TablePrivileges = v
	return s
}

type TablePrivilegeForGrantDBAccountPrivilegeInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AccountPrivilegeDetail *string `type:"string" json:",omitempty"`

	TableName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s TablePrivilegeForGrantDBAccountPrivilegeInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TablePrivilegeForGrantDBAccountPrivilegeInput) GoString() string {
	return s.String()
}

// SetAccountPrivilegeDetail sets the AccountPrivilegeDetail field's value.
func (s *TablePrivilegeForGrantDBAccountPrivilegeInput) SetAccountPrivilegeDetail(v string) *TablePrivilegeForGrantDBAccountPrivilegeInput {
	s.AccountPrivilegeDetail = &v
	return s
}

// SetTableName sets the TableName field's value.
func (s *TablePrivilegeForGrantDBAccountPrivilegeInput) SetTableName(v string) *TablePrivilegeForGrantDBAccountPrivilegeInput {
	s.TableName = &v
	return s
}
