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

// MSFT_CliAlias struct
type MSFT_CliAlias struct {
	cim.WmiInstance

	//
	Connection MSFT_CliConnection

	//
	Description string

	//
	Formats []MSFT_CliFormat

	//
	FriendlyName string

	//
	PWhere string

	//
	Qualifiers []MSFT_CliQualifier

	//
	Target string

	//
	Verbs []MSFT_CliVerb
}

// SetConnection sets the value of Connection for the instance
func (instance *MSFT_CliAlias) SetPropertyConnection(value MSFT_CliConnection) (err error) {
	return instance.SetProperty("Connection", value)
}

// GetConnection gets the value of Connection for the instance
func (instance *MSFT_CliAlias) GetPropertyConnection() (value MSFT_CliConnection, err error) {
	retValue, err := instance.GetProperty("Connection")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_CliConnection)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDescription sets the value of Description for the instance
func (instance *MSFT_CliAlias) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", value)
}

// GetDescription gets the value of Description for the instance
func (instance *MSFT_CliAlias) GetPropertyDescription() (value string, err error) {
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

// SetFormats sets the value of Formats for the instance
func (instance *MSFT_CliAlias) SetPropertyFormats(value []MSFT_CliFormat) (err error) {
	return instance.SetProperty("Formats", value)
}

// GetFormats gets the value of Formats for the instance
func (instance *MSFT_CliAlias) GetPropertyFormats() (value []MSFT_CliFormat, err error) {
	retValue, err := instance.GetProperty("Formats")
	if err != nil {
		return
	}
	value, ok := retValue.([]MSFT_CliFormat)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFriendlyName sets the value of FriendlyName for the instance
func (instance *MSFT_CliAlias) SetPropertyFriendlyName(value string) (err error) {
	return instance.SetProperty("FriendlyName", value)
}

// GetFriendlyName gets the value of FriendlyName for the instance
func (instance *MSFT_CliAlias) GetPropertyFriendlyName() (value string, err error) {
	retValue, err := instance.GetProperty("FriendlyName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPWhere sets the value of PWhere for the instance
func (instance *MSFT_CliAlias) SetPropertyPWhere(value string) (err error) {
	return instance.SetProperty("PWhere", value)
}

// GetPWhere gets the value of PWhere for the instance
func (instance *MSFT_CliAlias) GetPropertyPWhere() (value string, err error) {
	retValue, err := instance.GetProperty("PWhere")
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
func (instance *MSFT_CliAlias) SetPropertyQualifiers(value []MSFT_CliQualifier) (err error) {
	return instance.SetProperty("Qualifiers", value)
}

// GetQualifiers gets the value of Qualifiers for the instance
func (instance *MSFT_CliAlias) GetPropertyQualifiers() (value []MSFT_CliQualifier, err error) {
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

// SetTarget sets the value of Target for the instance
func (instance *MSFT_CliAlias) SetPropertyTarget(value string) (err error) {
	return instance.SetProperty("Target", value)
}

// GetTarget gets the value of Target for the instance
func (instance *MSFT_CliAlias) GetPropertyTarget() (value string, err error) {
	retValue, err := instance.GetProperty("Target")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVerbs sets the value of Verbs for the instance
func (instance *MSFT_CliAlias) SetPropertyVerbs(value []MSFT_CliVerb) (err error) {
	return instance.SetProperty("Verbs", value)
}

// GetVerbs gets the value of Verbs for the instance
func (instance *MSFT_CliAlias) GetPropertyVerbs() (value []MSFT_CliVerb, err error) {
	retValue, err := instance.GetProperty("Verbs")
	if err != nil {
		return
	}
	value, ok := retValue.([]MSFT_CliVerb)
	if !ok {
		// TODO: Set an error
	}
	return
}
