// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-operator/pkg/scyllaclient/internal/agent/models"
)

// NewJobStopParams creates a new JobStopParams object
// with the default values initialized.
func NewJobStopParams() *JobStopParams {
	var ()
	return &JobStopParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJobStopParamsWithTimeout creates a new JobStopParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJobStopParamsWithTimeout(timeout time.Duration) *JobStopParams {
	var ()
	return &JobStopParams{

		timeout: timeout,
	}
}

// NewJobStopParamsWithContext creates a new JobStopParams object
// with the default values initialized, and the ability to set a context for a request
func NewJobStopParamsWithContext(ctx context.Context) *JobStopParams {
	var ()
	return &JobStopParams{

		Context: ctx,
	}
}

// NewJobStopParamsWithHTTPClient creates a new JobStopParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewJobStopParamsWithHTTPClient(client *http.Client) *JobStopParams {
	var ()
	return &JobStopParams{
		HTTPClient: client,
	}
}

/*JobStopParams contains all the parameters to send to the API endpoint
for the job stop operation typically these are written to a http.Request
*/
type JobStopParams struct {

	/*Jobid
	  jobid

	*/
	Jobid *models.Jobid

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the job stop params
func (o *JobStopParams) WithTimeout(timeout time.Duration) *JobStopParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the job stop params
func (o *JobStopParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the job stop params
func (o *JobStopParams) WithContext(ctx context.Context) *JobStopParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the job stop params
func (o *JobStopParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the job stop params
func (o *JobStopParams) WithHTTPClient(client *http.Client) *JobStopParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the job stop params
func (o *JobStopParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJobid adds the jobid to the job stop params
func (o *JobStopParams) WithJobid(jobid *models.Jobid) *JobStopParams {
	o.SetJobid(jobid)
	return o
}

// SetJobid adds the jobid to the job stop params
func (o *JobStopParams) SetJobid(jobid *models.Jobid) {
	o.Jobid = jobid
}

// WriteToRequest writes these params to a swagger request
func (o *JobStopParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Jobid != nil {
		if err := r.SetBodyParam(o.Jobid); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}