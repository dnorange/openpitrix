// Code generated by go-swagger; DO NOT EDIT.

package access_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetRoleParams creates a new GetRoleParams object
// with the default values initialized.
func NewGetRoleParams() *GetRoleParams {
	var ()
	return &GetRoleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRoleParamsWithTimeout creates a new GetRoleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRoleParamsWithTimeout(timeout time.Duration) *GetRoleParams {
	var ()
	return &GetRoleParams{

		timeout: timeout,
	}
}

// NewGetRoleParamsWithContext creates a new GetRoleParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRoleParamsWithContext(ctx context.Context) *GetRoleParams {
	var ()
	return &GetRoleParams{

		Context: ctx,
	}
}

// NewGetRoleParamsWithHTTPClient creates a new GetRoleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRoleParamsWithHTTPClient(client *http.Client) *GetRoleParams {
	var ()
	return &GetRoleParams{
		HTTPClient: client,
	}
}

/*GetRoleParams contains all the parameters to send to the API endpoint
for the get role operation typically these are written to a http.Request
*/
type GetRoleParams struct {

	/*RoleID
	  required, use role id to get role info.

	*/
	RoleID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get role params
func (o *GetRoleParams) WithTimeout(timeout time.Duration) *GetRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get role params
func (o *GetRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get role params
func (o *GetRoleParams) WithContext(ctx context.Context) *GetRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get role params
func (o *GetRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get role params
func (o *GetRoleParams) WithHTTPClient(client *http.Client) *GetRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get role params
func (o *GetRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRoleID adds the roleID to the get role params
func (o *GetRoleParams) WithRoleID(roleID *string) *GetRoleParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the get role params
func (o *GetRoleParams) SetRoleID(roleID *string) {
	o.RoleID = roleID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.RoleID != nil {

		// query param role_id
		var qrRoleID string
		if o.RoleID != nil {
			qrRoleID = *o.RoleID
		}
		qRoleID := qrRoleID
		if qRoleID != "" {
			if err := r.SetQueryParam("role_id", qRoleID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
