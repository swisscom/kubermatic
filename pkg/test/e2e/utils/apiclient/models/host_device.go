// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HostDevice host device
//
// swagger:model HostDevice
type HostDevice struct {

	// DeviceName is the resource name of the host device exposed by a device plugin
	DeviceName string `json:"deviceName,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this host device
func (m *HostDevice) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this host device based on context it is used
func (m *HostDevice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HostDevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostDevice) UnmarshalBinary(b []byte) error {
	var res HostDevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
