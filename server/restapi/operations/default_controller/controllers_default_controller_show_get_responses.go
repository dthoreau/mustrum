// Code generated by go-swagger; DO NOT EDIT.

package default_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ControllersDefaultControllerShowGetOKCode is the HTTP code returned for type ControllersDefaultControllerShowGetOK
const ControllersDefaultControllerShowGetOKCode int = 200

/*ControllersDefaultControllerShowGetOK Success

swagger:response controllersDefaultControllerShowGetOK
*/
type ControllersDefaultControllerShowGetOK struct {
}

// NewControllersDefaultControllerShowGetOK creates ControllersDefaultControllerShowGetOK with default headers values
func NewControllersDefaultControllerShowGetOK() *ControllersDefaultControllerShowGetOK {

	return &ControllersDefaultControllerShowGetOK{}
}

// WriteResponse to the client
func (o *ControllersDefaultControllerShowGetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
