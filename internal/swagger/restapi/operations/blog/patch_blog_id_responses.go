// Code generated by go-swagger; DO NOT EDIT.

package blog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/romanyx/swagger_blog_example/internal/swagger/models"
)

// PatchBlogIDOKCode is the HTTP code returned for type PatchBlogIDOK
const PatchBlogIDOKCode int = 200

/*PatchBlogIDOK Success

swagger:response patchBlogIdOK
*/
type PatchBlogIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.MessageResponse `json:"body,omitempty"`
}

// NewPatchBlogIDOK creates PatchBlogIDOK with default headers values
func NewPatchBlogIDOK() *PatchBlogIDOK {

	return &PatchBlogIDOK{}
}

// WithPayload adds the payload to the patch blog Id o k response
func (o *PatchBlogIDOK) WithPayload(payload *models.MessageResponse) *PatchBlogIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch blog Id o k response
func (o *PatchBlogIDOK) SetPayload(payload *models.MessageResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchBlogIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PatchBlogIDDefault Error

swagger:response patchBlogIdDefault
*/
type PatchBlogIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewPatchBlogIDDefault creates PatchBlogIDDefault with default headers values
func NewPatchBlogIDDefault(code int) *PatchBlogIDDefault {
	if code <= 0 {
		code = 500
	}

	return &PatchBlogIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the patch blog ID default response
func (o *PatchBlogIDDefault) WithStatusCode(code int) *PatchBlogIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the patch blog ID default response
func (o *PatchBlogIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the patch blog ID default response
func (o *PatchBlogIDDefault) WithPayload(payload *models.ErrorResponse) *PatchBlogIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch blog ID default response
func (o *PatchBlogIDDefault) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchBlogIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
