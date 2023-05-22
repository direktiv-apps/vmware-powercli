// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostParamsBodyCommandsItems post params body commands items
//
// swagger:model postParamsBodyCommandsItems
type PostParamsBodyCommandsItems struct {

	// Command to run
	// Example: Get-VM -Name jq(.vm) | ConvertTo-Json -Depth 1 -AsArray
	Command string `json:"command,omitempty"`

	// Stops excecution if command fails, otherwise proceeds with next command
	Continue *bool `json:"continue,omitempty"`

	// Environment variables set for each command.
	// Example: [{"name":"MYVALUE","value":"hello"}]
	Envs []*PostParamsBodyCommandsItemsEnvsItems `json:"envs"`

	// If set to false the command will not print the full command with arguments to logs.
	Print *bool `json:"print,omitempty"`

	// If set to false the command will not print output to logs.
	Silent *bool `json:"silent,omitempty"`
}

// Validate validates this post params body commands items
func (m *PostParamsBodyCommandsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostParamsBodyCommandsItems) validateEnvs(formats strfmt.Registry) error {
	if swag.IsZero(m.Envs) { // not required
		return nil
	}

	for i := 0; i < len(m.Envs); i++ {
		if swag.IsZero(m.Envs[i]) { // not required
			continue
		}

		if m.Envs[i] != nil {
			if err := m.Envs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("envs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("envs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this post params body commands items based on the context it is used
func (m *PostParamsBodyCommandsItems) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEnvs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostParamsBodyCommandsItems) contextValidateEnvs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Envs); i++ {

		if m.Envs[i] != nil {
			if err := m.Envs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("envs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("envs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostParamsBodyCommandsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostParamsBodyCommandsItems) UnmarshalBinary(b []byte) error {
	var res PostParamsBodyCommandsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
