// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright 2021 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/coreos/go-oidc/v3/oidc"
)

// SigningCertHandlerFunc turns a function with the right signature into a signing cert handler
type SigningCertHandlerFunc func(SigningCertParams, *oidc.IDToken) middleware.Responder

// Handle executing the request and returning a response
func (fn SigningCertHandlerFunc) Handle(params SigningCertParams, principal *oidc.IDToken) middleware.Responder {
	return fn(params, principal)
}

// SigningCertHandler interface for that can handle valid signing cert params
type SigningCertHandler interface {
	Handle(SigningCertParams, *oidc.IDToken) middleware.Responder
}

// NewSigningCert creates a new http.Handler for the signing cert operation
func NewSigningCert(ctx *middleware.Context, handler SigningCertHandler) *SigningCert {
	return &SigningCert{Context: ctx, Handler: handler}
}

/* SigningCert swagger:route POST /signingCert signingCert

create a cert, return content with a location header (with URL to CTL entry)

*/
type SigningCert struct {
	Context *middleware.Context
	Handler SigningCertHandler
}

func (o *SigningCert) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewSigningCertParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *oidc.IDToken
	if uprinc != nil {
		principal = uprinc.(*oidc.IDToken) // this is really a oidc.IDToken, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
