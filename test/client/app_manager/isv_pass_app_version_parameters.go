// Code generated by go-swagger; DO NOT EDIT.

package app_manager

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

	"openpitrix.io/openpitrix/test/models"
)

// NewIsvPassAppVersionParams creates a new IsvPassAppVersionParams object
// with the default values initialized.
func NewIsvPassAppVersionParams() *IsvPassAppVersionParams {
	var ()
	return &IsvPassAppVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIsvPassAppVersionParamsWithTimeout creates a new IsvPassAppVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIsvPassAppVersionParamsWithTimeout(timeout time.Duration) *IsvPassAppVersionParams {
	var ()
	return &IsvPassAppVersionParams{

		timeout: timeout,
	}
}

// NewIsvPassAppVersionParamsWithContext creates a new IsvPassAppVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewIsvPassAppVersionParamsWithContext(ctx context.Context) *IsvPassAppVersionParams {
	var ()
	return &IsvPassAppVersionParams{

		Context: ctx,
	}
}

// NewIsvPassAppVersionParamsWithHTTPClient creates a new IsvPassAppVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIsvPassAppVersionParamsWithHTTPClient(client *http.Client) *IsvPassAppVersionParams {
	var ()
	return &IsvPassAppVersionParams{
		HTTPClient: client,
	}
}

/*IsvPassAppVersionParams contains all the parameters to send to the API endpoint
for the isv pass app version operation typically these are written to a http.Request
*/
type IsvPassAppVersionParams struct {

	/*Body*/
	Body *models.OpenpitrixPassAppVersionRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the isv pass app version params
func (o *IsvPassAppVersionParams) WithTimeout(timeout time.Duration) *IsvPassAppVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the isv pass app version params
func (o *IsvPassAppVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the isv pass app version params
func (o *IsvPassAppVersionParams) WithContext(ctx context.Context) *IsvPassAppVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the isv pass app version params
func (o *IsvPassAppVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the isv pass app version params
func (o *IsvPassAppVersionParams) WithHTTPClient(client *http.Client) *IsvPassAppVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the isv pass app version params
func (o *IsvPassAppVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the isv pass app version params
func (o *IsvPassAppVersionParams) WithBody(body *models.OpenpitrixPassAppVersionRequest) *IsvPassAppVersionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the isv pass app version params
func (o *IsvPassAppVersionParams) SetBody(body *models.OpenpitrixPassAppVersionRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *IsvPassAppVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
