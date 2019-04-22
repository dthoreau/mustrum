// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Colour colour
// swagger:model colour
type Colour struct {

	// blue
	// Maximum: 255
	// Minimum: 0
	Blue *int64 `json:"blue,omitempty"`

	// green
	// Maximum: 255
	// Minimum: 0
	Green *int64 `json:"green,omitempty"`

	// red
	// Maximum: 255
	// Minimum: 0
	Red *int64 `json:"red,omitempty"`
}

// Validate validates this colour
func (m *Colour) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGreen(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Colour) validateBlue(formats strfmt.Registry) error {

	if swag.IsZero(m.Blue) { // not required
		return nil
	}

	if err := validate.MinimumInt("blue", "body", int64(*m.Blue), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("blue", "body", int64(*m.Blue), 255, false); err != nil {
		return err
	}

	return nil
}

func (m *Colour) validateGreen(formats strfmt.Registry) error {

	if swag.IsZero(m.Green) { // not required
		return nil
	}

	if err := validate.MinimumInt("green", "body", int64(*m.Green), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("green", "body", int64(*m.Green), 255, false); err != nil {
		return err
	}

	return nil
}

func (m *Colour) validateRed(formats strfmt.Registry) error {

	if swag.IsZero(m.Red) { // not required
		return nil
	}

	if err := validate.MinimumInt("red", "body", int64(*m.Red), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("red", "body", int64(*m.Red), 255, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Colour) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Colour) UnmarshalBinary(b []byte) error {
	var res Colour
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
