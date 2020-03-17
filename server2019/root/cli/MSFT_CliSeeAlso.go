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

// MSFT_CliSeeAlso struct
type MSFT_CliSeeAlso struct {
	cim.WmiInstance

	//
	Description string

	//
	Original MSFT_CliAlias

	//
	Related MSFT_CliAlias
}

// SetDescription sets the value of Description for the instance
func (instance *MSFT_CliSeeAlso) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", value)
}

// GetDescription gets the value of Description for the instance
func (instance *MSFT_CliSeeAlso) GetPropertyDescription() (value string, err error) {
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

// SetOriginal sets the value of Original for the instance
func (instance *MSFT_CliSeeAlso) SetPropertyOriginal(value MSFT_CliAlias) (err error) {
	return instance.SetProperty("Original", value)
}

// GetOriginal gets the value of Original for the instance
func (instance *MSFT_CliSeeAlso) GetPropertyOriginal() (value MSFT_CliAlias, err error) {
	retValue, err := instance.GetProperty("Original")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_CliAlias)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRelated sets the value of Related for the instance
func (instance *MSFT_CliSeeAlso) SetPropertyRelated(value MSFT_CliAlias) (err error) {
	return instance.SetProperty("Related", value)
}

// GetRelated gets the value of Related for the instance
func (instance *MSFT_CliSeeAlso) GetPropertyRelated() (value MSFT_CliAlias, err error) {
	retValue, err := instance.GetProperty("Related")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_CliAlias)
	if !ok {
		// TODO: Set an error
	}
	return
}
