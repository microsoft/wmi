// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Cli
//////////////////////////////////////////////
package cli

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_CliVerb struct
type MSFT_CliVerb struct {
	cim.WmiInstance

	//
	Derivation string

	//
	Description string

	//
	Name string

	//
	Parameters []MSFT_CliParam

	//
	Qualifiers []MSFT_CliQualifier

	//
	Usage string

	//
	VerbType uint32
}

// SetDerivation sets the value of Derivation for the instance
func (instance *MSFT_CliVerb) SetPropertyDerivation(value string) (err error) {
	return instance.SetProperty("Derivation", value)
}

// GetDerivation gets the value of Derivation for the instance
func (instance *MSFT_CliVerb) GetPropertyDerivation() (value string, err error) {
	retValue, err := instance.GetProperty("Derivation")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDescription sets the value of Description for the instance
func (instance *MSFT_CliVerb) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", value)
}

// GetDescription gets the value of Description for the instance
func (instance *MSFT_CliVerb) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_CliVerb) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_CliVerb) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParameters sets the value of Parameters for the instance
func (instance *MSFT_CliVerb) SetPropertyParameters(value []MSFT_CliParam) (err error) {
	return instance.SetProperty("Parameters", value)
}

// GetParameters gets the value of Parameters for the instance
func (instance *MSFT_CliVerb) GetPropertyParameters() (value []MSFT_CliParam, err error) {
	retValue, err := instance.GetProperty("Parameters")
	if err != nil {
		return
	}
	value, ok := retValue.([]MSFT_CliParam)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetQualifiers sets the value of Qualifiers for the instance
func (instance *MSFT_CliVerb) SetPropertyQualifiers(value []MSFT_CliQualifier) (err error) {
	return instance.SetProperty("Qualifiers", value)
}

// GetQualifiers gets the value of Qualifiers for the instance
func (instance *MSFT_CliVerb) GetPropertyQualifiers() (value []MSFT_CliQualifier, err error) {
	retValue, err := instance.GetProperty("Qualifiers")
	if err != nil {
		return
	}
	value, ok := retValue.([]MSFT_CliQualifier)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUsage sets the value of Usage for the instance
func (instance *MSFT_CliVerb) SetPropertyUsage(value string) (err error) {
	return instance.SetProperty("Usage", value)
}

// GetUsage gets the value of Usage for the instance
func (instance *MSFT_CliVerb) GetPropertyUsage() (value string, err error) {
	retValue, err := instance.GetProperty("Usage")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVerbType sets the value of VerbType for the instance
func (instance *MSFT_CliVerb) SetPropertyVerbType(value uint32) (err error) {
	return instance.SetProperty("VerbType", value)
}

// GetVerbType gets the value of VerbType for the instance
func (instance *MSFT_CliVerb) GetPropertyVerbType() (value uint32, err error) {
	retValue, err := instance.GetProperty("VerbType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
