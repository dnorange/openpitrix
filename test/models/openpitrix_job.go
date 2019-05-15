// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixJob openpitrix job
// swagger:model openpitrixJob
type OpenpitrixJob struct {

	// id of app deployed in cluster
	AppID string `json:"app_id,omitempty"`

	// id of cluster run job
	ClusterID string `json:"cluster_id,omitempty"`

	// the time job create
	CreateTime strfmt.DateTime `json:"create_time,omitempty"`

	// directive, a json string, describe the info of running the job action
	Directive string `json:"directive,omitempty"`

	// error code, if job run failed will return a error code
	ErrorCode int64 `json:"error_code,omitempty"`

	// host name of server
	Executor string `json:"executor,omitempty"`

	// describe job's action eg:[CreateCluster|StartClusters|...]
	JobAction string `json:"job_action,omitempty"`

	// job id
	JobID string `json:"job_id,omitempty"`

	// owner
	Owner string `json:"owner,omitempty"`

	// own path, concat string group_path:user_id
	OwnerPath string `json:"owner_path,omitempty"`

	// runtime provider eg:[qingcloud|aliyun|aws|kubernetes]
	Provider string `json:"provider,omitempty"`

	// id of runtime of cluster
	RuntimeID string `json:"runtime_id,omitempty"`

	// status eg.[successful|failed|running|pending]
	Status string `json:"status,omitempty"`

	// record the status changed time
	StatusTime strfmt.DateTime `json:"status_time,omitempty"`

	// total count of task in job, a job contain one more task
	TaskCount int64 `json:"task_count,omitempty"`

	// id of specific app version
	VersionID string `json:"version_id,omitempty"`
}

// Validate validates this openpitrix job
func (m *OpenpitrixJob) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixJob) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixJob) UnmarshalBinary(b []byte) error {
	var res OpenpitrixJob
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
