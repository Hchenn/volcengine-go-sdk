// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package apig20221112iface provides an interface to enable mocking the APIG20221112 service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package apig20221112

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

// APIG20221112API provides an interface to enable mocking the
// apig20221112.APIG20221112 service client's API operation,
//
//    // volcengine sdk func uses an SDK service client to make a request to
//    // APIG20221112.
//    func myFunc(svc APIG20221112API) bool {
//        // Make svc.CreateRoute request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := apig20221112.New(sess)
//
//        myFunc(svc)
//    }
//
type APIG20221112API interface {
	CreateRouteCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateRouteCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateRouteCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateRoute(*CreateRouteInput) (*CreateRouteOutput, error)
	CreateRouteWithContext(volcengine.Context, *CreateRouteInput, ...request.Option) (*CreateRouteOutput, error)
	CreateRouteRequest(*CreateRouteInput) (*request.Request, *CreateRouteOutput)

	DeleteRouteCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteRouteCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteRouteCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteRoute(*DeleteRouteInput) (*DeleteRouteOutput, error)
	DeleteRouteWithContext(volcengine.Context, *DeleteRouteInput, ...request.Option) (*DeleteRouteOutput, error)
	DeleteRouteRequest(*DeleteRouteInput) (*request.Request, *DeleteRouteOutput)

	GetRouteCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GetRouteCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetRouteCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetRoute(*GetRouteInput) (*GetRouteOutput, error)
	GetRouteWithContext(volcengine.Context, *GetRouteInput, ...request.Option) (*GetRouteOutput, error)
	GetRouteRequest(*GetRouteInput) (*request.Request, *GetRouteOutput)

	ListRoutesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListRoutesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListRoutesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListRoutes(*ListRoutesInput) (*ListRoutesOutput, error)
	ListRoutesWithContext(volcengine.Context, *ListRoutesInput, ...request.Option) (*ListRoutesOutput, error)
	ListRoutesRequest(*ListRoutesInput) (*request.Request, *ListRoutesOutput)

	UpdateRouteCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateRouteCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateRouteCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateRoute(*UpdateRouteInput) (*UpdateRouteOutput, error)
	UpdateRouteWithContext(volcengine.Context, *UpdateRouteInput, ...request.Option) (*UpdateRouteOutput, error)
	UpdateRouteRequest(*UpdateRouteInput) (*request.Request, *UpdateRouteOutput)
}

var _ APIG20221112API = (*APIG20221112)(nil)
