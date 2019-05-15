// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixCreatePasswordResetResponse openpitrix create password reset response
// swagger:model openpitrixCreatePasswordResetResponse
type OpenpitrixCreatePasswordResetResponse struct {

	// reset id, used to change password
	ResetID string `json:"reset_id,omitempty"`

	// id of user that reset password
	UserID string `json:"user_id,omitempty"`
}

// Validate validates this openpitrix create password reset response
func (m *OpenpitrixCreatePasswordResetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixCreatePasswordResetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixCreatePasswordResetResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixCreatePasswordResetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
