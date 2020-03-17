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

// MSFT_CliProperty struct
type MSFT_CliProperty struct {
	cim.WmiInstance

	//
	Derivation string

	//
	Description string

	//
	Name string

	//
	Qualifiers []MSFT_CliQualifier
}

// SetDerivation sets the value of Derivation for the instance
func (instance *MSFT_CliProperty) SetPropertyDerivation(value string) (err error) {
	return instance.SetProperty("Derivation", value)
}

// GetDerivation gets the value of Derivation for the instance
func (instance *MSFT_CliProperty) GetPropertyDerivation() (value string, err error) {
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
func (instance *MSFT_CliProperty) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", value)
}

// GetDescription gets the value of Description for the instance
func (instance *MSFT_CliProperty) GetPropertyDescription() (value string, err error) {
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
func (instance *MSFT_CliProperty) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_CliProperty) GetPropertyName() (value string, err error) {
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

// SetQualifiers sets the value of Qualifiers for the instance
func (instance *MSFT_CliProperty) SetPropertyQualifiers(value []MSFT_CliQualifier) (err error) {
	return instance.SetProperty("Qualifiers", value)
}

// GetQualifiers gets the value of Qualifiers for the instance
func (instance *MSFT_CliProperty) GetPropertyQualifiers() (value []MSFT_CliQualifier, err error) {
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
