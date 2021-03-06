// Code generated by go-swagger; DO NOT EDIT.

package account_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// GetPasswordResetReader is a Reader for the GetPasswordReset structure.
type GetPasswordResetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPasswordResetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPasswordResetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPasswordResetOK creates a GetPasswordResetOK with default headers values
func NewGetPasswordResetOK() *GetPasswordResetOK {
	return &GetPasswordResetOK{}
}

/*GetPasswordResetOK handles this case with default header values.

A successful response.
*/
type GetPasswordResetOK struct {
	Payload *models.OpenpitrixGetPasswordResetResponse
}

func (o *GetPasswordResetOK) Error() string {
	return fmt.Sprintf("[GET /v1/users/password/reset/{reset_id}][%d] getPasswordResetOK  %+v", 200, o.Payload)
}

func (o *GetPasswordResetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixGetPasswordResetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
