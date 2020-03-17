// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.PEH
//////////////////////////////////////////////
package peh

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_Expression struct
type MSFT_Expression struct {
	cim.WmiInstance

	//
	SourceInfo string

	//
	SourceLines []string
}

// SetSourceInfo sets the value of SourceInfo for the instance
func (instance *MSFT_Expression) SetPropertySourceInfo(value string) (err error) {
	return instance.SetProperty("SourceInfo", value)
}

// GetSourceInfo gets the value of SourceInfo for the instance
func (instance *MSFT_Expression) GetPropertySourceInfo() (value string, err error) {
	retValue, err := instance.GetProperty("SourceInfo")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSourceLines sets the value of SourceLines for the instance
func (instance *MSFT_Expression) SetPropertySourceLines(value []string) (err error) {
	return instance.SetProperty("SourceLines", value)
}

// GetSourceLines gets the value of SourceLines for the instance
func (instance *MSFT_Expression) GetPropertySourceLines() (value []string, err error) {
	retValue, err := instance.GetProperty("SourceLines")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
