// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixDeleteAppsRequest openpitrix delete apps request
// swagger:model openpitrixDeleteAppsRequest
type OpenpitrixDeleteAppsRequest struct {

	// app id
	AppID []string `json:"app_id"`
}

// Validate validates this openpitrix delete apps request
func (m *OpenpitrixDeleteAppsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixDeleteAppsRequest) validateAppID(formats strfmt.Registry) error {

	if swag.IsZero(m.AppID) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixDeleteAppsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixDeleteAppsRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixDeleteAppsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}