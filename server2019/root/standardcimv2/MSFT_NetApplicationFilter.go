// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetApplicationFilter struct
type MSFT_NetApplicationFilter struct {
	CIM_FilterEntryBase

	//
	AppPath string

	//
	Package string
}

// SetAppPath sets the value of AppPath for the instance
func (instance *MSFT_NetApplicationFilter) SetPropertyAppPath(value string) (err error) {
	return instance.SetProperty("AppPath", value)
}

// GetAppPath gets the value of AppPath for the instance
func (instance *MSFT_NetApplicationFilter) GetPropertyAppPath() (value string, err error) {
	retValue, err := instance.GetProperty("AppPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPackage sets the value of Package for the instance
func (instance *MSFT_NetApplicationFilter) SetPropertyPackage(value string) (err error) {
	return instance.SetProperty("Package", value)
}

// GetPackage gets the value of Package for the instance
func (instance *MSFT_NetApplicationFilter) GetPropertyPackage() (value string, err error) {
	retValue, err := instance.GetProperty("Package")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
