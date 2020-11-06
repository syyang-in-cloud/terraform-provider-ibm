// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

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

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// NewPcloudPvminstancesDeleteParams creates a new PcloudPvminstancesDeleteParams object
// with the default values initialized.
func NewPcloudPvminstancesDeleteParams() *PcloudPvminstancesDeleteParams {
	var ()
	return &PcloudPvminstancesDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudPvminstancesDeleteParamsWithTimeout creates a new PcloudPvminstancesDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPcloudPvminstancesDeleteParamsWithTimeout(timeout time.Duration) *PcloudPvminstancesDeleteParams {
	var ()
	return &PcloudPvminstancesDeleteParams{

		timeout: timeout,
	}
}

// NewPcloudPvminstancesDeleteParamsWithContext creates a new PcloudPvminstancesDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewPcloudPvminstancesDeleteParamsWithContext(ctx context.Context) *PcloudPvminstancesDeleteParams {
	var ()
	return &PcloudPvminstancesDeleteParams{

		Context: ctx,
	}
}

// NewPcloudPvminstancesDeleteParamsWithHTTPClient creates a new PcloudPvminstancesDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPcloudPvminstancesDeleteParamsWithHTTPClient(client *http.Client) *PcloudPvminstancesDeleteParams {
	var ()
	return &PcloudPvminstancesDeleteParams{
		HTTPClient: client,
	}
}

/*PcloudPvminstancesDeleteParams contains all the parameters to send to the API endpoint
for the pcloud pvminstances delete operation typically these are written to a http.Request
*/
type PcloudPvminstancesDeleteParams struct {

	/*Body
	  Parameters used when deleting a PCloud PVM Instance

	*/
	Body *models.PVMInstanceDelete
	/*CloudInstanceID
	  Cloud Instance ID of a PCloud Instance

	*/
	CloudInstanceID string
	/*PvmInstanceID
	  PCloud PVM Instance ID

	*/
	PvmInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pcloud pvminstances delete params
func (o *PcloudPvminstancesDeleteParams) WithTimeout(timeout time.Duration) *PcloudPvminstancesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud pvminstances delete params
func (o *PcloudPvminstancesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud pvminstances delete params
func (o *PcloudPvminstancesDeleteParams) WithContext(ctx context.Context) *PcloudPvminstancesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud pvminstances delete params
func (o *PcloudPvminstancesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud pvminstances delete params
func (o *PcloudPvminstancesDeleteParams) WithHTTPClient(client *http.Client) *PcloudPvminstancesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud pvminstances delete params
func (o *PcloudPvminstancesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud pvminstances delete params
func (o *PcloudPvminstancesDeleteParams) WithBody(body *models.PVMInstanceDelete) *PcloudPvminstancesDeleteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud pvminstances delete params
func (o *PcloudPvminstancesDeleteParams) SetBody(body *models.PVMInstanceDelete) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud pvminstances delete params
func (o *PcloudPvminstancesDeleteParams) WithCloudInstanceID(cloudInstanceID string) *PcloudPvminstancesDeleteParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud pvminstances delete params
func (o *PcloudPvminstancesDeleteParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithPvmInstanceID adds the pvmInstanceID to the pcloud pvminstances delete params
func (o *PcloudPvminstancesDeleteParams) WithPvmInstanceID(pvmInstanceID string) *PcloudPvminstancesDeleteParams {
	o.SetPvmInstanceID(pvmInstanceID)
	return o
}

// SetPvmInstanceID adds the pvmInstanceId to the pcloud pvminstances delete params
func (o *PcloudPvminstancesDeleteParams) SetPvmInstanceID(pvmInstanceID string) {
	o.PvmInstanceID = pvmInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudPvminstancesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	// path param pvm_instance_id
	if err := r.SetPathParam("pvm_instance_id", o.PvmInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
