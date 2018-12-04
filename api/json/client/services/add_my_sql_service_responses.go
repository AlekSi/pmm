// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/percona/pmm/api/json/models"
)

// AddMySQLServiceReader is a Reader for the AddMySQLService structure.
type AddMySQLServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddMySQLServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddMySQLServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddMySQLServiceOK creates a AddMySQLServiceOK with default headers values
func NewAddMySQLServiceOK() *AddMySQLServiceOK {
	return &AddMySQLServiceOK{}
}

/*AddMySQLServiceOK handles this case with default header values.

(empty)
*/
type AddMySQLServiceOK struct {
	Payload *models.InventoryAddMySQLServiceResponse
}

func (o *AddMySQLServiceOK) Error() string {
	return fmt.Sprintf("[POST /v0/inventory/Services/AddMySQL][%d] addMySqlServiceOK  %+v", 200, o.Payload)
}

func (o *AddMySQLServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryAddMySQLServiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}