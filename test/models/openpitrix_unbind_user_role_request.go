// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixUnbindUserRoleRequest openpitrix unbind user role request
// swagger:model openpitrixUnbindUserRoleRequest
type OpenpitrixUnbindUserRoleRequest struct {

	// ids of role for user to unbind with
	RoleID []string `json:"role_id"`

	// ids of user to unbind
	UserID []string `json:"user_id"`
}

// Validate validates this openpitrix unbind user role request
func (m *OpenpitrixUnbindUserRoleRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRoleID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixUnbindUserRoleRequest) validateRoleID(formats strfmt.Registry) error {

	if swag.IsZero(m.RoleID) { // not required
		return nil
	}

	return nil
}

func (m *OpenpitrixUnbindUserRoleRequest) validateUserID(formats strfmt.Registry) error {

	if swag.IsZero(m.UserID) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixUnbindUserRoleRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixUnbindUserRoleRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixUnbindUserRoleRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
