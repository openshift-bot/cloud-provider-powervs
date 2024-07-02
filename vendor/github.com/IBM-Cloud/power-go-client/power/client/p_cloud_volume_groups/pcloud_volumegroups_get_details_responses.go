// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volume_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudVolumegroupsGetDetailsReader is a Reader for the PcloudVolumegroupsGetDetails structure.
type PcloudVolumegroupsGetDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudVolumegroupsGetDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudVolumegroupsGetDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudVolumegroupsGetDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudVolumegroupsGetDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudVolumegroupsGetDetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudVolumegroupsGetDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudVolumegroupsGetDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}/details] pcloud.volumegroups.getDetails", response, response.Code())
	}
}

// NewPcloudVolumegroupsGetDetailsOK creates a PcloudVolumegroupsGetDetailsOK with default headers values
func NewPcloudVolumegroupsGetDetailsOK() *PcloudVolumegroupsGetDetailsOK {
	return &PcloudVolumegroupsGetDetailsOK{}
}

/*
PcloudVolumegroupsGetDetailsOK describes a response with status code 200, with default header values.

OK
*/
type PcloudVolumegroupsGetDetailsOK struct {
	Payload *models.VolumeGroupDetails
}

// IsSuccess returns true when this pcloud volumegroups get details o k response has a 2xx status code
func (o *PcloudVolumegroupsGetDetailsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud volumegroups get details o k response has a 3xx status code
func (o *PcloudVolumegroupsGetDetailsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volumegroups get details o k response has a 4xx status code
func (o *PcloudVolumegroupsGetDetailsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud volumegroups get details o k response has a 5xx status code
func (o *PcloudVolumegroupsGetDetailsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud volumegroups get details o k response a status code equal to that given
func (o *PcloudVolumegroupsGetDetailsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud volumegroups get details o k response
func (o *PcloudVolumegroupsGetDetailsOK) Code() int {
	return 200
}

func (o *PcloudVolumegroupsGetDetailsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}/details][%d] pcloudVolumegroupsGetDetailsOK %s", 200, payload)
}

func (o *PcloudVolumegroupsGetDetailsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}/details][%d] pcloudVolumegroupsGetDetailsOK %s", 200, payload)
}

func (o *PcloudVolumegroupsGetDetailsOK) GetPayload() *models.VolumeGroupDetails {
	return o.Payload
}

func (o *PcloudVolumegroupsGetDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumeGroupDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumegroupsGetDetailsBadRequest creates a PcloudVolumegroupsGetDetailsBadRequest with default headers values
func NewPcloudVolumegroupsGetDetailsBadRequest() *PcloudVolumegroupsGetDetailsBadRequest {
	return &PcloudVolumegroupsGetDetailsBadRequest{}
}

/*
PcloudVolumegroupsGetDetailsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudVolumegroupsGetDetailsBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud volumegroups get details bad request response has a 2xx status code
func (o *PcloudVolumegroupsGetDetailsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud volumegroups get details bad request response has a 3xx status code
func (o *PcloudVolumegroupsGetDetailsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volumegroups get details bad request response has a 4xx status code
func (o *PcloudVolumegroupsGetDetailsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud volumegroups get details bad request response has a 5xx status code
func (o *PcloudVolumegroupsGetDetailsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud volumegroups get details bad request response a status code equal to that given
func (o *PcloudVolumegroupsGetDetailsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud volumegroups get details bad request response
func (o *PcloudVolumegroupsGetDetailsBadRequest) Code() int {
	return 400
}

func (o *PcloudVolumegroupsGetDetailsBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}/details][%d] pcloudVolumegroupsGetDetailsBadRequest %s", 400, payload)
}

func (o *PcloudVolumegroupsGetDetailsBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}/details][%d] pcloudVolumegroupsGetDetailsBadRequest %s", 400, payload)
}

func (o *PcloudVolumegroupsGetDetailsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumegroupsGetDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumegroupsGetDetailsUnauthorized creates a PcloudVolumegroupsGetDetailsUnauthorized with default headers values
func NewPcloudVolumegroupsGetDetailsUnauthorized() *PcloudVolumegroupsGetDetailsUnauthorized {
	return &PcloudVolumegroupsGetDetailsUnauthorized{}
}

/*
PcloudVolumegroupsGetDetailsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudVolumegroupsGetDetailsUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud volumegroups get details unauthorized response has a 2xx status code
func (o *PcloudVolumegroupsGetDetailsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud volumegroups get details unauthorized response has a 3xx status code
func (o *PcloudVolumegroupsGetDetailsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volumegroups get details unauthorized response has a 4xx status code
func (o *PcloudVolumegroupsGetDetailsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud volumegroups get details unauthorized response has a 5xx status code
func (o *PcloudVolumegroupsGetDetailsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud volumegroups get details unauthorized response a status code equal to that given
func (o *PcloudVolumegroupsGetDetailsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud volumegroups get details unauthorized response
func (o *PcloudVolumegroupsGetDetailsUnauthorized) Code() int {
	return 401
}

func (o *PcloudVolumegroupsGetDetailsUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}/details][%d] pcloudVolumegroupsGetDetailsUnauthorized %s", 401, payload)
}

func (o *PcloudVolumegroupsGetDetailsUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}/details][%d] pcloudVolumegroupsGetDetailsUnauthorized %s", 401, payload)
}

func (o *PcloudVolumegroupsGetDetailsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumegroupsGetDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumegroupsGetDetailsForbidden creates a PcloudVolumegroupsGetDetailsForbidden with default headers values
func NewPcloudVolumegroupsGetDetailsForbidden() *PcloudVolumegroupsGetDetailsForbidden {
	return &PcloudVolumegroupsGetDetailsForbidden{}
}

/*
PcloudVolumegroupsGetDetailsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudVolumegroupsGetDetailsForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud volumegroups get details forbidden response has a 2xx status code
func (o *PcloudVolumegroupsGetDetailsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud volumegroups get details forbidden response has a 3xx status code
func (o *PcloudVolumegroupsGetDetailsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volumegroups get details forbidden response has a 4xx status code
func (o *PcloudVolumegroupsGetDetailsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud volumegroups get details forbidden response has a 5xx status code
func (o *PcloudVolumegroupsGetDetailsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud volumegroups get details forbidden response a status code equal to that given
func (o *PcloudVolumegroupsGetDetailsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud volumegroups get details forbidden response
func (o *PcloudVolumegroupsGetDetailsForbidden) Code() int {
	return 403
}

func (o *PcloudVolumegroupsGetDetailsForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}/details][%d] pcloudVolumegroupsGetDetailsForbidden %s", 403, payload)
}

func (o *PcloudVolumegroupsGetDetailsForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}/details][%d] pcloudVolumegroupsGetDetailsForbidden %s", 403, payload)
}

func (o *PcloudVolumegroupsGetDetailsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumegroupsGetDetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumegroupsGetDetailsNotFound creates a PcloudVolumegroupsGetDetailsNotFound with default headers values
func NewPcloudVolumegroupsGetDetailsNotFound() *PcloudVolumegroupsGetDetailsNotFound {
	return &PcloudVolumegroupsGetDetailsNotFound{}
}

/*
PcloudVolumegroupsGetDetailsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudVolumegroupsGetDetailsNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud volumegroups get details not found response has a 2xx status code
func (o *PcloudVolumegroupsGetDetailsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud volumegroups get details not found response has a 3xx status code
func (o *PcloudVolumegroupsGetDetailsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volumegroups get details not found response has a 4xx status code
func (o *PcloudVolumegroupsGetDetailsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud volumegroups get details not found response has a 5xx status code
func (o *PcloudVolumegroupsGetDetailsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud volumegroups get details not found response a status code equal to that given
func (o *PcloudVolumegroupsGetDetailsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud volumegroups get details not found response
func (o *PcloudVolumegroupsGetDetailsNotFound) Code() int {
	return 404
}

func (o *PcloudVolumegroupsGetDetailsNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}/details][%d] pcloudVolumegroupsGetDetailsNotFound %s", 404, payload)
}

func (o *PcloudVolumegroupsGetDetailsNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}/details][%d] pcloudVolumegroupsGetDetailsNotFound %s", 404, payload)
}

func (o *PcloudVolumegroupsGetDetailsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumegroupsGetDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumegroupsGetDetailsInternalServerError creates a PcloudVolumegroupsGetDetailsInternalServerError with default headers values
func NewPcloudVolumegroupsGetDetailsInternalServerError() *PcloudVolumegroupsGetDetailsInternalServerError {
	return &PcloudVolumegroupsGetDetailsInternalServerError{}
}

/*
PcloudVolumegroupsGetDetailsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudVolumegroupsGetDetailsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud volumegroups get details internal server error response has a 2xx status code
func (o *PcloudVolumegroupsGetDetailsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud volumegroups get details internal server error response has a 3xx status code
func (o *PcloudVolumegroupsGetDetailsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volumegroups get details internal server error response has a 4xx status code
func (o *PcloudVolumegroupsGetDetailsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud volumegroups get details internal server error response has a 5xx status code
func (o *PcloudVolumegroupsGetDetailsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud volumegroups get details internal server error response a status code equal to that given
func (o *PcloudVolumegroupsGetDetailsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud volumegroups get details internal server error response
func (o *PcloudVolumegroupsGetDetailsInternalServerError) Code() int {
	return 500
}

func (o *PcloudVolumegroupsGetDetailsInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}/details][%d] pcloudVolumegroupsGetDetailsInternalServerError %s", 500, payload)
}

func (o *PcloudVolumegroupsGetDetailsInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}/details][%d] pcloudVolumegroupsGetDetailsInternalServerError %s", 500, payload)
}

func (o *PcloudVolumegroupsGetDetailsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumegroupsGetDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
