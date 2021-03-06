// Code generated by go-swagger; DO NOT EDIT.

package xtra_db_cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RestartXtraDBClusterReader is a Reader for the RestartXtraDBCluster structure.
type RestartXtraDBClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestartXtraDBClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRestartXtraDBClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRestartXtraDBClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRestartXtraDBClusterOK creates a RestartXtraDBClusterOK with default headers values
func NewRestartXtraDBClusterOK() *RestartXtraDBClusterOK {
	return &RestartXtraDBClusterOK{}
}

/*RestartXtraDBClusterOK handles this case with default header values.

A successful response.
*/
type RestartXtraDBClusterOK struct {
	Payload interface{}
}

func (o *RestartXtraDBClusterOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/XtraDBCluster/Restart][%d] restartXtraDbClusterOk  %+v", 200, o.Payload)
}

func (o *RestartXtraDBClusterOK) GetPayload() interface{} {
	return o.Payload
}

func (o *RestartXtraDBClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestartXtraDBClusterDefault creates a RestartXtraDBClusterDefault with default headers values
func NewRestartXtraDBClusterDefault(code int) *RestartXtraDBClusterDefault {
	return &RestartXtraDBClusterDefault{
		_statusCode: code,
	}
}

/*RestartXtraDBClusterDefault handles this case with default header values.

An unexpected error response.
*/
type RestartXtraDBClusterDefault struct {
	_statusCode int

	Payload *RestartXtraDBClusterDefaultBody
}

// Code gets the status code for the restart xtra DB cluster default response
func (o *RestartXtraDBClusterDefault) Code() int {
	return o._statusCode
}

func (o *RestartXtraDBClusterDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/XtraDBCluster/Restart][%d] RestartXtraDBCluster default  %+v", o._statusCode, o.Payload)
}

func (o *RestartXtraDBClusterDefault) GetPayload() *RestartXtraDBClusterDefaultBody {
	return o.Payload
}

func (o *RestartXtraDBClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(RestartXtraDBClusterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*RestartXtraDBClusterBody restart xtra DB cluster body
swagger:model RestartXtraDBClusterBody
*/
type RestartXtraDBClusterBody struct {

	// Kubernetes cluster name.
	KubernetesClusterName string `json:"kubernetes_cluster_name,omitempty"`

	// XtraDB cluster name.
	Name string `json:"name,omitempty"`
}

// Validate validates this restart xtra DB cluster body
func (o *RestartXtraDBClusterBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RestartXtraDBClusterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RestartXtraDBClusterBody) UnmarshalBinary(b []byte) error {
	var res RestartXtraDBClusterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RestartXtraDBClusterDefaultBody restart xtra DB cluster default body
swagger:model RestartXtraDBClusterDefaultBody
*/
type RestartXtraDBClusterDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this restart xtra DB cluster default body
func (o *RestartXtraDBClusterDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RestartXtraDBClusterDefaultBody) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("RestartXtraDBCluster default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *RestartXtraDBClusterDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RestartXtraDBClusterDefaultBody) UnmarshalBinary(b []byte) error {
	var res RestartXtraDBClusterDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
