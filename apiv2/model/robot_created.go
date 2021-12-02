// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RobotCreated The response for robot account creation.
//
// swagger:model RobotCreated
type RobotCreated struct {

	// The creation time of the robot.
	// Format: date-time
	CreationTime strfmt.DateTime `json:"creation_time,omitempty"`

	// The expiration data of the robot
	ExpiresAt int64 `json:"expires_at,omitempty"`

	// The ID of the robot
	ID int64 `json:"id,omitempty"`

	// The name of the tag
	Name string `json:"name,omitempty"`

	// The secret of the robot
	Secret string `json:"secret,omitempty"`
}

// Validate validates this robot created
func (m *RobotCreated) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RobotCreated) validateCreationTime(formats strfmt.Registry) error {

	if swag.IsZero(m.CreationTime) { // not required
		return nil
	}

	if err := validate.FormatOf("creation_time", "body", "date-time", m.CreationTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RobotCreated) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RobotCreated) UnmarshalBinary(b []byte) error {
	var res RobotCreated
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}