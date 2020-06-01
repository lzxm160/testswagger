// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/lzxm160/testswagger/models"
)

// CreateOKCode is the HTTP code returned for type CreateOK
const CreateOKCode int = 200

/*CreateOK returns a hash

swagger:response createOK
*/
type CreateOK struct {

	/*did create hash
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewCreateOK creates CreateOK with default headers values
func NewCreateOK() *CreateOK {

	return &CreateOK{}
}

// WithPayload adds the payload to the create o k response
func (o *CreateOK) WithPayload(payload string) *CreateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create o k response
func (o *CreateOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*CreateDefault error

swagger:response createDefault
*/
type CreateDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateDefault creates CreateDefault with default headers values
func NewCreateDefault(code int) *CreateDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create default response
func (o *CreateDefault) WithStatusCode(code int) *CreateDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create default response
func (o *CreateDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create default response
func (o *CreateDefault) WithPayload(payload *models.Error) *CreateDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create default response
func (o *CreateDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
