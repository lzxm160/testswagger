// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ResolveHandlerFunc turns a function with the right signature into a resolve handler
type ResolveHandlerFunc func(ResolveParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ResolveHandlerFunc) Handle(params ResolveParams) middleware.Responder {
	return fn(params)
}

// ResolveHandler interface for that can handle valid resolve params
type ResolveHandler interface {
	Handle(ResolveParams) middleware.Responder
}

// NewResolve creates a new http.Handler for the resolve operation
func NewResolve(ctx *middleware.Context, handler ResolveHandler) *Resolve {
	return &Resolve{Context: ctx, Handler: handler}
}

/*Resolve swagger:route GET /identifiers/{identifier} resolve

Resolve a DID or other identifier.

*/
type Resolve struct {
	Context *middleware.Context
	Handler ResolveHandler
}

func (o *Resolve) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewResolveParams()
	fmt.Println(r.RequestURI,r.Method)
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
