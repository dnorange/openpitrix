// Code generated by go-swagger; DO NOT EDIT.

package isv_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// DescribeVendorVerifyInfosReader is a Reader for the DescribeVendorVerifyInfos structure.
type DescribeVendorVerifyInfosReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeVendorVerifyInfosReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDescribeVendorVerifyInfosOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDescribeVendorVerifyInfosOK creates a DescribeVendorVerifyInfosOK with default headers values
func NewDescribeVendorVerifyInfosOK() *DescribeVendorVerifyInfosOK {
	return &DescribeVendorVerifyInfosOK{}
}

/*DescribeVendorVerifyInfosOK handles this case with default header values.

A successful response.
*/
type DescribeVendorVerifyInfosOK struct {
	Payload *models.OpenpitrixDescribeVendorVerifyInfosResponse
}

func (o *DescribeVendorVerifyInfosOK) Error() string {
	return fmt.Sprintf("[GET /v1/app_vendors][%d] describeVendorVerifyInfosOK  %+v", 200, o.Payload)
}

func (o *DescribeVendorVerifyInfosOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixDescribeVendorVerifyInfosResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
