// Code generated by go-swagger; DO NOT EDIT.

package alerts

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

// NewToggleAlertParams creates a new ToggleAlertParams object
// with the default values initialized.
func NewToggleAlertParams() *ToggleAlertParams {
	var ()
	return &ToggleAlertParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewToggleAlertParamsWithTimeout creates a new ToggleAlertParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewToggleAlertParamsWithTimeout(timeout time.Duration) *ToggleAlertParams {
	var ()
	return &ToggleAlertParams{

		timeout: timeout,
	}
}

// NewToggleAlertParamsWithContext creates a new ToggleAlertParams object
// with the default values initialized, and the ability to set a context for a request
func NewToggleAlertParamsWithContext(ctx context.Context) *ToggleAlertParams {
	var ()
	return &ToggleAlertParams{

		Context: ctx,
	}
}

// NewToggleAlertParamsWithHTTPClient creates a new ToggleAlertParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewToggleAlertParamsWithHTTPClient(client *http.Client) *ToggleAlertParams {
	var ()
	return &ToggleAlertParams{
		HTTPClient: client,
	}
}

/*ToggleAlertParams contains all the parameters to send to the API endpoint
for the toggle alert operation typically these are written to a http.Request
*/
type ToggleAlertParams struct {

	/*Body*/
	Body ToggleAlertBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the toggle alert params
func (o *ToggleAlertParams) WithTimeout(timeout time.Duration) *ToggleAlertParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the toggle alert params
func (o *ToggleAlertParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the toggle alert params
func (o *ToggleAlertParams) WithContext(ctx context.Context) *ToggleAlertParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the toggle alert params
func (o *ToggleAlertParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the toggle alert params
func (o *ToggleAlertParams) WithHTTPClient(client *http.Client) *ToggleAlertParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the toggle alert params
func (o *ToggleAlertParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the toggle alert params
func (o *ToggleAlertParams) WithBody(body ToggleAlertBody) *ToggleAlertParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the toggle alert params
func (o *ToggleAlertParams) SetBody(body ToggleAlertBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ToggleAlertParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
