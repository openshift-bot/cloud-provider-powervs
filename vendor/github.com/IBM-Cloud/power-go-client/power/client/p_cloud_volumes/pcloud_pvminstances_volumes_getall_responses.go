// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volumes

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

// PcloudPvminstancesVolumesGetallReader is a Reader for the PcloudPvminstancesVolumesGetall structure.
type PcloudPvminstancesVolumesGetallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPvminstancesVolumesGetallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudPvminstancesVolumesGetallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudPvminstancesVolumesGetallBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudPvminstancesVolumesGetallUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudPvminstancesVolumesGetallForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudPvminstancesVolumesGetallNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudPvminstancesVolumesGetallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes] pcloud.pvminstances.volumes.getall", response, response.Code())
	}
}

// NewPcloudPvminstancesVolumesGetallOK creates a PcloudPvminstancesVolumesGetallOK with default headers values
func NewPcloudPvminstancesVolumesGetallOK() *PcloudPvminstancesVolumesGetallOK {
	return &PcloudPvminstancesVolumesGetallOK{}
}

/*
PcloudPvminstancesVolumesGetallOK describes a response with status code 200, with default header values.

OK
*/
type PcloudPvminstancesVolumesGetallOK struct {
	Payload *models.Volumes
}

// IsSuccess returns true when this pcloud pvminstances volumes getall o k response has a 2xx status code
func (o *PcloudPvminstancesVolumesGetallOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud pvminstances volumes getall o k response has a 3xx status code
func (o *PcloudPvminstancesVolumesGetallOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances volumes getall o k response has a 4xx status code
func (o *PcloudPvminstancesVolumesGetallOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances volumes getall o k response has a 5xx status code
func (o *PcloudPvminstancesVolumesGetallOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances volumes getall o k response a status code equal to that given
func (o *PcloudPvminstancesVolumesGetallOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud pvminstances volumes getall o k response
func (o *PcloudPvminstancesVolumesGetallOK) Code() int {
	return 200
}

func (o *PcloudPvminstancesVolumesGetallOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes][%d] pcloudPvminstancesVolumesGetallOK %s", 200, payload)
}

func (o *PcloudPvminstancesVolumesGetallOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes][%d] pcloudPvminstancesVolumesGetallOK %s", 200, payload)
}

func (o *PcloudPvminstancesVolumesGetallOK) GetPayload() *models.Volumes {
	return o.Payload
}

func (o *PcloudPvminstancesVolumesGetallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Volumes)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVolumesGetallBadRequest creates a PcloudPvminstancesVolumesGetallBadRequest with default headers values
func NewPcloudPvminstancesVolumesGetallBadRequest() *PcloudPvminstancesVolumesGetallBadRequest {
	return &PcloudPvminstancesVolumesGetallBadRequest{}
}

/*
PcloudPvminstancesVolumesGetallBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudPvminstancesVolumesGetallBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances volumes getall bad request response has a 2xx status code
func (o *PcloudPvminstancesVolumesGetallBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances volumes getall bad request response has a 3xx status code
func (o *PcloudPvminstancesVolumesGetallBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances volumes getall bad request response has a 4xx status code
func (o *PcloudPvminstancesVolumesGetallBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances volumes getall bad request response has a 5xx status code
func (o *PcloudPvminstancesVolumesGetallBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances volumes getall bad request response a status code equal to that given
func (o *PcloudPvminstancesVolumesGetallBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud pvminstances volumes getall bad request response
func (o *PcloudPvminstancesVolumesGetallBadRequest) Code() int {
	return 400
}

func (o *PcloudPvminstancesVolumesGetallBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes][%d] pcloudPvminstancesVolumesGetallBadRequest %s", 400, payload)
}

func (o *PcloudPvminstancesVolumesGetallBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes][%d] pcloudPvminstancesVolumesGetallBadRequest %s", 400, payload)
}

func (o *PcloudPvminstancesVolumesGetallBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesVolumesGetallBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVolumesGetallUnauthorized creates a PcloudPvminstancesVolumesGetallUnauthorized with default headers values
func NewPcloudPvminstancesVolumesGetallUnauthorized() *PcloudPvminstancesVolumesGetallUnauthorized {
	return &PcloudPvminstancesVolumesGetallUnauthorized{}
}

/*
PcloudPvminstancesVolumesGetallUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudPvminstancesVolumesGetallUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances volumes getall unauthorized response has a 2xx status code
func (o *PcloudPvminstancesVolumesGetallUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances volumes getall unauthorized response has a 3xx status code
func (o *PcloudPvminstancesVolumesGetallUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances volumes getall unauthorized response has a 4xx status code
func (o *PcloudPvminstancesVolumesGetallUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances volumes getall unauthorized response has a 5xx status code
func (o *PcloudPvminstancesVolumesGetallUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances volumes getall unauthorized response a status code equal to that given
func (o *PcloudPvminstancesVolumesGetallUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud pvminstances volumes getall unauthorized response
func (o *PcloudPvminstancesVolumesGetallUnauthorized) Code() int {
	return 401
}

func (o *PcloudPvminstancesVolumesGetallUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes][%d] pcloudPvminstancesVolumesGetallUnauthorized %s", 401, payload)
}

func (o *PcloudPvminstancesVolumesGetallUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes][%d] pcloudPvminstancesVolumesGetallUnauthorized %s", 401, payload)
}

func (o *PcloudPvminstancesVolumesGetallUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesVolumesGetallUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVolumesGetallForbidden creates a PcloudPvminstancesVolumesGetallForbidden with default headers values
func NewPcloudPvminstancesVolumesGetallForbidden() *PcloudPvminstancesVolumesGetallForbidden {
	return &PcloudPvminstancesVolumesGetallForbidden{}
}

/*
PcloudPvminstancesVolumesGetallForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudPvminstancesVolumesGetallForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances volumes getall forbidden response has a 2xx status code
func (o *PcloudPvminstancesVolumesGetallForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances volumes getall forbidden response has a 3xx status code
func (o *PcloudPvminstancesVolumesGetallForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances volumes getall forbidden response has a 4xx status code
func (o *PcloudPvminstancesVolumesGetallForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances volumes getall forbidden response has a 5xx status code
func (o *PcloudPvminstancesVolumesGetallForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances volumes getall forbidden response a status code equal to that given
func (o *PcloudPvminstancesVolumesGetallForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud pvminstances volumes getall forbidden response
func (o *PcloudPvminstancesVolumesGetallForbidden) Code() int {
	return 403
}

func (o *PcloudPvminstancesVolumesGetallForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes][%d] pcloudPvminstancesVolumesGetallForbidden %s", 403, payload)
}

func (o *PcloudPvminstancesVolumesGetallForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes][%d] pcloudPvminstancesVolumesGetallForbidden %s", 403, payload)
}

func (o *PcloudPvminstancesVolumesGetallForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesVolumesGetallForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVolumesGetallNotFound creates a PcloudPvminstancesVolumesGetallNotFound with default headers values
func NewPcloudPvminstancesVolumesGetallNotFound() *PcloudPvminstancesVolumesGetallNotFound {
	return &PcloudPvminstancesVolumesGetallNotFound{}
}

/*
PcloudPvminstancesVolumesGetallNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudPvminstancesVolumesGetallNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances volumes getall not found response has a 2xx status code
func (o *PcloudPvminstancesVolumesGetallNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances volumes getall not found response has a 3xx status code
func (o *PcloudPvminstancesVolumesGetallNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances volumes getall not found response has a 4xx status code
func (o *PcloudPvminstancesVolumesGetallNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances volumes getall not found response has a 5xx status code
func (o *PcloudPvminstancesVolumesGetallNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances volumes getall not found response a status code equal to that given
func (o *PcloudPvminstancesVolumesGetallNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud pvminstances volumes getall not found response
func (o *PcloudPvminstancesVolumesGetallNotFound) Code() int {
	return 404
}

func (o *PcloudPvminstancesVolumesGetallNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes][%d] pcloudPvminstancesVolumesGetallNotFound %s", 404, payload)
}

func (o *PcloudPvminstancesVolumesGetallNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes][%d] pcloudPvminstancesVolumesGetallNotFound %s", 404, payload)
}

func (o *PcloudPvminstancesVolumesGetallNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesVolumesGetallNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVolumesGetallInternalServerError creates a PcloudPvminstancesVolumesGetallInternalServerError with default headers values
func NewPcloudPvminstancesVolumesGetallInternalServerError() *PcloudPvminstancesVolumesGetallInternalServerError {
	return &PcloudPvminstancesVolumesGetallInternalServerError{}
}

/*
PcloudPvminstancesVolumesGetallInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudPvminstancesVolumesGetallInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances volumes getall internal server error response has a 2xx status code
func (o *PcloudPvminstancesVolumesGetallInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances volumes getall internal server error response has a 3xx status code
func (o *PcloudPvminstancesVolumesGetallInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances volumes getall internal server error response has a 4xx status code
func (o *PcloudPvminstancesVolumesGetallInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances volumes getall internal server error response has a 5xx status code
func (o *PcloudPvminstancesVolumesGetallInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud pvminstances volumes getall internal server error response a status code equal to that given
func (o *PcloudPvminstancesVolumesGetallInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud pvminstances volumes getall internal server error response
func (o *PcloudPvminstancesVolumesGetallInternalServerError) Code() int {
	return 500
}

func (o *PcloudPvminstancesVolumesGetallInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes][%d] pcloudPvminstancesVolumesGetallInternalServerError %s", 500, payload)
}

func (o *PcloudPvminstancesVolumesGetallInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes][%d] pcloudPvminstancesVolumesGetallInternalServerError %s", 500, payload)
}

func (o *PcloudPvminstancesVolumesGetallInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesVolumesGetallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
