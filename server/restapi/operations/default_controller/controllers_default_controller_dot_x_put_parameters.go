// Code generated by go-swagger; DO NOT EDIT.

package default_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "mustrum/server/models"
)

// NewControllersDefaultControllerDotXPutParams creates a new ControllersDefaultControllerDotXPutParams object
// no default values defined in spec.
func NewControllersDefaultControllerDotXPutParams() ControllersDefaultControllerDotXPutParams {

	return ControllersDefaultControllerDotXPutParams{}
}

// ControllersDefaultControllerDotXPutParams contains all the bound params for the controllers default controller dot x put operation
// typically these are obtained from a http.Request
//
// swagger:parameters controllers.default_controller.dot_x_put
type ControllersDefaultControllerDotXPutParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	Colour *models.Colour
	/*x component of pixel
	  Required: true
	  In: path
	*/
	X float64
	/*y component of pixel
	  Required: true
	  In: query
	*/
	Y float64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewControllersDefaultControllerDotXPutParams() beforehand.
func (o *ControllersDefaultControllerDotXPutParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Colour
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("colour", "body"))
			} else {
				res = append(res, errors.NewParseError("colour", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Colour = &body
			}
		}
	} else {
		res = append(res, errors.Required("colour", "body"))
	}
	rX, rhkX, _ := route.Params.GetOK("x")
	if err := o.bindX(rX, rhkX, route.Formats); err != nil {
		res = append(res, err)
	}

	qY, qhkY, _ := qs.GetOK("y")
	if err := o.bindY(qY, qhkY, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindX binds and validates parameter X from path.
func (o *ControllersDefaultControllerDotXPutParams) bindX(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertFloat64(raw)
	if err != nil {
		return errors.InvalidType("x", "path", "float64", raw)
	}
	o.X = value

	return nil
}

// bindY binds and validates parameter Y from query.
func (o *ControllersDefaultControllerDotXPutParams) bindY(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("y", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("y", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertFloat64(raw)
	if err != nil {
		return errors.InvalidType("y", "query", "float64", raw)
	}
	o.Y = value

	return nil
}
