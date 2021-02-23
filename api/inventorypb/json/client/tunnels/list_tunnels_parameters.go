// Code generated by go-swagger; DO NOT EDIT.

package tunnels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListTunnelsParams creates a new ListTunnelsParams object
// with the default values initialized.
func NewListTunnelsParams() *ListTunnelsParams {
	var ()
	return &ListTunnelsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListTunnelsParamsWithTimeout creates a new ListTunnelsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListTunnelsParamsWithTimeout(timeout time.Duration) *ListTunnelsParams {
	var ()
	return &ListTunnelsParams{

		timeout: timeout,
	}
}

// NewListTunnelsParamsWithContext creates a new ListTunnelsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListTunnelsParamsWithContext(ctx context.Context) *ListTunnelsParams {
	var ()
	return &ListTunnelsParams{

		Context: ctx,
	}
}

// NewListTunnelsParamsWithHTTPClient creates a new ListTunnelsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListTunnelsParamsWithHTTPClient(client *http.Client) *ListTunnelsParams {
	var ()
	return &ListTunnelsParams{
		HTTPClient: client,
	}
}

/*ListTunnelsParams contains all the parameters to send to the API endpoint
for the list tunnels operation typically these are written to a http.Request
*/
type ListTunnelsParams struct {

	/*Body*/
	Body ListTunnelsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list tunnels params
func (o *ListTunnelsParams) WithTimeout(timeout time.Duration) *ListTunnelsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list tunnels params
func (o *ListTunnelsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list tunnels params
func (o *ListTunnelsParams) WithContext(ctx context.Context) *ListTunnelsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list tunnels params
func (o *ListTunnelsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list tunnels params
func (o *ListTunnelsParams) WithHTTPClient(client *http.Client) *ListTunnelsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list tunnels params
func (o *ListTunnelsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the list tunnels params
func (o *ListTunnelsParams) WithBody(body ListTunnelsBody) *ListTunnelsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the list tunnels params
func (o *ListTunnelsParams) SetBody(body ListTunnelsBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ListTunnelsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}