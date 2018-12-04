// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixModifyAppRequest openpitrix modify app request
// swagger:model openpitrixModifyAppRequest
type OpenpitrixModifyAppRequest struct {

	// app id
	AppID string `json:"app_id,omitempty"`

	// category id
	CategoryID string `json:"category_id,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// home
	Home string `json:"home,omitempty"`

	// keywords
	Keywords string `json:"keywords,omitempty"`

	// maintainers
	Maintainers string `json:"maintainers,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// readme
	Readme string `json:"readme,omitempty"`

	// sources
	Sources string `json:"sources,omitempty"`
}

// Validate validates this openpitrix modify app request
func (m *OpenpitrixModifyAppRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixModifyAppRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixModifyAppRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixModifyAppRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
