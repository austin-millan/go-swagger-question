// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResponseMeta response meta
//
// swagger:model ResponseMeta
type ResponseMeta struct {

	// meta
	Meta string `json:"meta,omitempty"`
}

// Validate validates this response meta
func (m *ResponseMeta) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this response meta based on context it is used
func (m *ResponseMeta) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResponseMeta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponseMeta) UnmarshalBinary(b []byte) error {
	var res ResponseMeta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}