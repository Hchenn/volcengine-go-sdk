// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package volcobserve

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opEnablePresetAlertTemplateCommon = "EnablePresetAlertTemplate"

// EnablePresetAlertTemplateCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the EnablePresetAlertTemplateCommon operation. The "output" return
// value will be populated with the EnablePresetAlertTemplateCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned EnablePresetAlertTemplateCommon Request to send the API call to the service.
// the "output" return value is not valid until after EnablePresetAlertTemplateCommon Send returns without error.
//
// See EnablePresetAlertTemplateCommon for more information on using the EnablePresetAlertTemplateCommon
// API call, and error handling.
//
//    // Example sending a request using the EnablePresetAlertTemplateCommonRequest method.
//    req, resp := client.EnablePresetAlertTemplateCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VOLCOBSERVE) EnablePresetAlertTemplateCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opEnablePresetAlertTemplateCommon,
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

// EnablePresetAlertTemplateCommon API operation for VOLC_OBSERVE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VOLC_OBSERVE's
// API operation EnablePresetAlertTemplateCommon for usage and error information.
func (c *VOLCOBSERVE) EnablePresetAlertTemplateCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.EnablePresetAlertTemplateCommonRequest(input)
	return out, req.Send()
}

// EnablePresetAlertTemplateCommonWithContext is the same as EnablePresetAlertTemplateCommon with the addition of
// the ability to pass a context and additional request options.
//
// See EnablePresetAlertTemplateCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VOLCOBSERVE) EnablePresetAlertTemplateCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.EnablePresetAlertTemplateCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opEnablePresetAlertTemplate = "EnablePresetAlertTemplate"

// EnablePresetAlertTemplateRequest generates a "volcengine/request.Request" representing the
// client's request for the EnablePresetAlertTemplate operation. The "output" return
// value will be populated with the EnablePresetAlertTemplateCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned EnablePresetAlertTemplateCommon Request to send the API call to the service.
// the "output" return value is not valid until after EnablePresetAlertTemplateCommon Send returns without error.
//
// See EnablePresetAlertTemplate for more information on using the EnablePresetAlertTemplate
// API call, and error handling.
//
//    // Example sending a request using the EnablePresetAlertTemplateRequest method.
//    req, resp := client.EnablePresetAlertTemplateRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VOLCOBSERVE) EnablePresetAlertTemplateRequest(input *EnablePresetAlertTemplateInput) (req *request.Request, output *EnablePresetAlertTemplateOutput) {
	op := &request.Operation{
		Name:       opEnablePresetAlertTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &EnablePresetAlertTemplateInput{}
	}

	output = &EnablePresetAlertTemplateOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// EnablePresetAlertTemplate API operation for VOLC_OBSERVE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VOLC_OBSERVE's
// API operation EnablePresetAlertTemplate for usage and error information.
func (c *VOLCOBSERVE) EnablePresetAlertTemplate(input *EnablePresetAlertTemplateInput) (*EnablePresetAlertTemplateOutput, error) {
	req, out := c.EnablePresetAlertTemplateRequest(input)
	return out, req.Send()
}

// EnablePresetAlertTemplateWithContext is the same as EnablePresetAlertTemplate with the addition of
// the ability to pass a context and additional request options.
//
// See EnablePresetAlertTemplate for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VOLCOBSERVE) EnablePresetAlertTemplateWithContext(ctx volcengine.Context, input *EnablePresetAlertTemplateInput, opts ...request.Option) (*EnablePresetAlertTemplateOutput, error) {
	req, out := c.EnablePresetAlertTemplateRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type EnablePresetAlertTemplateInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AlertMethods []*string `type:"list" json:",omitempty"`

	ContactGroupIds []*string `type:"list" json:",omitempty"`

	// EffectEndAt is a required field
	EffectEndAt *string `type:"string" json:",omitempty" required:"true"`

	// EffectStartAt is a required field
	EffectStartAt *string `type:"string" json:",omitempty" required:"true"`

	NotificationId *string `type:"string" json:",omitempty"`

	// ProjectName is a required field
	ProjectName *string `type:"string" json:",omitempty" required:"true"`

	// TemplateId is a required field
	TemplateId *string `type:"string" json:",omitempty" required:"true"`

	UpgradePresetAlertTemplate *bool `type:"boolean" json:",omitempty"`

	Webhook *string `type:"string" json:",omitempty"`

	WebhookIds []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s EnablePresetAlertTemplateInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EnablePresetAlertTemplateInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EnablePresetAlertTemplateInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "EnablePresetAlertTemplateInput"}
	if s.EffectEndAt == nil {
		invalidParams.Add(request.NewErrParamRequired("EffectEndAt"))
	}
	if s.EffectStartAt == nil {
		invalidParams.Add(request.NewErrParamRequired("EffectStartAt"))
	}
	if s.ProjectName == nil {
		invalidParams.Add(request.NewErrParamRequired("ProjectName"))
	}
	if s.TemplateId == nil {
		invalidParams.Add(request.NewErrParamRequired("TemplateId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAlertMethods sets the AlertMethods field's value.
func (s *EnablePresetAlertTemplateInput) SetAlertMethods(v []*string) *EnablePresetAlertTemplateInput {
	s.AlertMethods = v
	return s
}

// SetContactGroupIds sets the ContactGroupIds field's value.
func (s *EnablePresetAlertTemplateInput) SetContactGroupIds(v []*string) *EnablePresetAlertTemplateInput {
	s.ContactGroupIds = v
	return s
}

// SetEffectEndAt sets the EffectEndAt field's value.
func (s *EnablePresetAlertTemplateInput) SetEffectEndAt(v string) *EnablePresetAlertTemplateInput {
	s.EffectEndAt = &v
	return s
}

// SetEffectStartAt sets the EffectStartAt field's value.
func (s *EnablePresetAlertTemplateInput) SetEffectStartAt(v string) *EnablePresetAlertTemplateInput {
	s.EffectStartAt = &v
	return s
}

// SetNotificationId sets the NotificationId field's value.
func (s *EnablePresetAlertTemplateInput) SetNotificationId(v string) *EnablePresetAlertTemplateInput {
	s.NotificationId = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *EnablePresetAlertTemplateInput) SetProjectName(v string) *EnablePresetAlertTemplateInput {
	s.ProjectName = &v
	return s
}

// SetTemplateId sets the TemplateId field's value.
func (s *EnablePresetAlertTemplateInput) SetTemplateId(v string) *EnablePresetAlertTemplateInput {
	s.TemplateId = &v
	return s
}

// SetUpgradePresetAlertTemplate sets the UpgradePresetAlertTemplate field's value.
func (s *EnablePresetAlertTemplateInput) SetUpgradePresetAlertTemplate(v bool) *EnablePresetAlertTemplateInput {
	s.UpgradePresetAlertTemplate = &v
	return s
}

// SetWebhook sets the Webhook field's value.
func (s *EnablePresetAlertTemplateInput) SetWebhook(v string) *EnablePresetAlertTemplateInput {
	s.Webhook = &v
	return s
}

// SetWebhookIds sets the WebhookIds field's value.
func (s *EnablePresetAlertTemplateInput) SetWebhookIds(v []*string) *EnablePresetAlertTemplateInput {
	s.WebhookIds = v
	return s
}

type EnablePresetAlertTemplateOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Data []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s EnablePresetAlertTemplateOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EnablePresetAlertTemplateOutput) GoString() string {
	return s.String()
}

// SetData sets the Data field's value.
func (s *EnablePresetAlertTemplateOutput) SetData(v []*string) *EnablePresetAlertTemplateOutput {
	s.Data = v
	return s
}
