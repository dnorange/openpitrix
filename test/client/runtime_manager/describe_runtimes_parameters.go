// Code generated by go-swagger; DO NOT EDIT.

package runtime_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDescribeRuntimesParams creates a new DescribeRuntimesParams object
// with the default values initialized.
func NewDescribeRuntimesParams() *DescribeRuntimesParams {
	var ()
	return &DescribeRuntimesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeRuntimesParamsWithTimeout creates a new DescribeRuntimesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDescribeRuntimesParamsWithTimeout(timeout time.Duration) *DescribeRuntimesParams {
	var ()
	return &DescribeRuntimesParams{

		timeout: timeout,
	}
}

// NewDescribeRuntimesParamsWithContext creates a new DescribeRuntimesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDescribeRuntimesParamsWithContext(ctx context.Context) *DescribeRuntimesParams {
	var ()
	return &DescribeRuntimesParams{

		Context: ctx,
	}
}

// NewDescribeRuntimesParamsWithHTTPClient creates a new DescribeRuntimesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDescribeRuntimesParamsWithHTTPClient(client *http.Client) *DescribeRuntimesParams {
	var ()
	return &DescribeRuntimesParams{
		HTTPClient: client,
	}
}

/*DescribeRuntimesParams contains all the parameters to send to the API endpoint
for the describe runtimes operation typically these are written to a http.Request
*/
type DescribeRuntimesParams struct {

	/*DisplayColumns
	  select columns to display.

	*/
	DisplayColumns []string
	/*Limit
	  data limit per page, default value 20, max value 200.

	*/
	Limit *int64
	/*Offset
	  data offset, default 0.

	*/
	Offset *int64
	/*Owner
	  owner.

	*/
	Owner []string
	/*Provider
	  runtime provider eg.[qingcloud|aliyun|aws|kubernetes].

	*/
	Provider []string
	/*Reverse
	  value = 0 sort ASC, value = 1 sort DESC.

	*/
	Reverse *bool
	/*RuntimeID
	  runtime ids.

	*/
	RuntimeID []string
	/*SearchWord
	  query key, support these fields(runtime_id, provider, zone, status, owner).

	*/
	SearchWord *string
	/*SortKey
	  sort key, order by sort_key, default create_time.

	*/
	SortKey *string
	/*Status
	  status eg.[active|deleted].

	*/
	Status []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the describe runtimes params
func (o *DescribeRuntimesParams) WithTimeout(timeout time.Duration) *DescribeRuntimesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe runtimes params
func (o *DescribeRuntimesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe runtimes params
func (o *DescribeRuntimesParams) WithContext(ctx context.Context) *DescribeRuntimesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe runtimes params
func (o *DescribeRuntimesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe runtimes params
func (o *DescribeRuntimesParams) WithHTTPClient(client *http.Client) *DescribeRuntimesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe runtimes params
func (o *DescribeRuntimesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDisplayColumns adds the displayColumns to the describe runtimes params
func (o *DescribeRuntimesParams) WithDisplayColumns(displayColumns []string) *DescribeRuntimesParams {
	o.SetDisplayColumns(displayColumns)
	return o
}

// SetDisplayColumns adds the displayColumns to the describe runtimes params
func (o *DescribeRuntimesParams) SetDisplayColumns(displayColumns []string) {
	o.DisplayColumns = displayColumns
}

// WithLimit adds the limit to the describe runtimes params
func (o *DescribeRuntimesParams) WithLimit(limit *int64) *DescribeRuntimesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the describe runtimes params
func (o *DescribeRuntimesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the describe runtimes params
func (o *DescribeRuntimesParams) WithOffset(offset *int64) *DescribeRuntimesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the describe runtimes params
func (o *DescribeRuntimesParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOwner adds the owner to the describe runtimes params
func (o *DescribeRuntimesParams) WithOwner(owner []string) *DescribeRuntimesParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the describe runtimes params
func (o *DescribeRuntimesParams) SetOwner(owner []string) {
	o.Owner = owner
}

// WithProvider adds the provider to the describe runtimes params
func (o *DescribeRuntimesParams) WithProvider(provider []string) *DescribeRuntimesParams {
	o.SetProvider(provider)
	return o
}

// SetProvider adds the provider to the describe runtimes params
func (o *DescribeRuntimesParams) SetProvider(provider []string) {
	o.Provider = provider
}

// WithReverse adds the reverse to the describe runtimes params
func (o *DescribeRuntimesParams) WithReverse(reverse *bool) *DescribeRuntimesParams {
	o.SetReverse(reverse)
	return o
}

// SetReverse adds the reverse to the describe runtimes params
func (o *DescribeRuntimesParams) SetReverse(reverse *bool) {
	o.Reverse = reverse
}

// WithRuntimeID adds the runtimeID to the describe runtimes params
func (o *DescribeRuntimesParams) WithRuntimeID(runtimeID []string) *DescribeRuntimesParams {
	o.SetRuntimeID(runtimeID)
	return o
}

// SetRuntimeID adds the runtimeId to the describe runtimes params
func (o *DescribeRuntimesParams) SetRuntimeID(runtimeID []string) {
	o.RuntimeID = runtimeID
}

// WithSearchWord adds the searchWord to the describe runtimes params
func (o *DescribeRuntimesParams) WithSearchWord(searchWord *string) *DescribeRuntimesParams {
	o.SetSearchWord(searchWord)
	return o
}

// SetSearchWord adds the searchWord to the describe runtimes params
func (o *DescribeRuntimesParams) SetSearchWord(searchWord *string) {
	o.SearchWord = searchWord
}

// WithSortKey adds the sortKey to the describe runtimes params
func (o *DescribeRuntimesParams) WithSortKey(sortKey *string) *DescribeRuntimesParams {
	o.SetSortKey(sortKey)
	return o
}

// SetSortKey adds the sortKey to the describe runtimes params
func (o *DescribeRuntimesParams) SetSortKey(sortKey *string) {
	o.SortKey = sortKey
}

// WithStatus adds the status to the describe runtimes params
func (o *DescribeRuntimesParams) WithStatus(status []string) *DescribeRuntimesParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the describe runtimes params
func (o *DescribeRuntimesParams) SetStatus(status []string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeRuntimesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesDisplayColumns := o.DisplayColumns

	joinedDisplayColumns := swag.JoinByFormat(valuesDisplayColumns, "multi")
	// query array param display_columns
	if err := r.SetQueryParam("display_columns", joinedDisplayColumns...); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	valuesOwner := o.Owner

	joinedOwner := swag.JoinByFormat(valuesOwner, "multi")
	// query array param owner
	if err := r.SetQueryParam("owner", joinedOwner...); err != nil {
		return err
	}

	valuesProvider := o.Provider

	joinedProvider := swag.JoinByFormat(valuesProvider, "multi")
	// query array param provider
	if err := r.SetQueryParam("provider", joinedProvider...); err != nil {
		return err
	}

	if o.Reverse != nil {

		// query param reverse
		var qrReverse bool
		if o.Reverse != nil {
			qrReverse = *o.Reverse
		}
		qReverse := swag.FormatBool(qrReverse)
		if qReverse != "" {
			if err := r.SetQueryParam("reverse", qReverse); err != nil {
				return err
			}
		}

	}

	valuesRuntimeID := o.RuntimeID

	joinedRuntimeID := swag.JoinByFormat(valuesRuntimeID, "multi")
	// query array param runtime_id
	if err := r.SetQueryParam("runtime_id", joinedRuntimeID...); err != nil {
		return err
	}

	if o.SearchWord != nil {

		// query param search_word
		var qrSearchWord string
		if o.SearchWord != nil {
			qrSearchWord = *o.SearchWord
		}
		qSearchWord := qrSearchWord
		if qSearchWord != "" {
			if err := r.SetQueryParam("search_word", qSearchWord); err != nil {
				return err
			}
		}

	}

	if o.SortKey != nil {

		// query param sort_key
		var qrSortKey string
		if o.SortKey != nil {
			qrSortKey = *o.SortKey
		}
		qSortKey := qrSortKey
		if qSortKey != "" {
			if err := r.SetQueryParam("sort_key", qSortKey); err != nil {
				return err
			}
		}

	}

	valuesStatus := o.Status

	joinedStatus := swag.JoinByFormat(valuesStatus, "multi")
	// query array param status
	if err := r.SetQueryParam("status", joinedStatus...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
