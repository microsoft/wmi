// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage
//////////////////////////////////////////////
package storage

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_FileShareAccessControlEntry struct
type MSFT_FileShareAccessControlEntry struct {
	cim.WmiInstance

	//
	AccessControlType uint16

	//
	AccessRight uint16

	//
	AccountName string
}

// SetAccessControlType sets the value of AccessControlType for the instance
func (instance *MSFT_FileShareAccessControlEntry) SetPropertyAccessControlType(value uint16) (err error) {
	return instance.SetProperty("AccessControlType", value)
}

// GetAccessControlType gets the value of AccessControlType for the instance
func (instance *MSFT_FileShareAccessControlEntry) GetPropertyAccessControlType() (value uint16, err error) {
	retValue, err := instance.GetProperty("AccessControlType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAccessRight sets the value of AccessRight for the instance
func (instance *MSFT_FileShareAccessControlEntry) SetPropertyAccessRight(value uint16) (err error) {
	return instance.SetProperty("AccessRight", value)
}

// GetAccessRight gets the value of AccessRight for the instance
func (instance *MSFT_FileShareAccessControlEntry) GetPropertyAccessRight() (value uint16, err error) {
	retValue, err := instance.GetProperty("AccessRight")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAccountName sets the value of AccountName for the instance
func (instance *MSFT_FileShareAccessControlEntry) SetPropertyAccountName(value string) (err error) {
	return instance.SetProperty("AccountName", value)
}

// GetAccountName gets the value of AccountName for the instance
func (instance *MSFT_FileShareAccessControlEntry) GetPropertyAccountName() (value string, err error) {
	retValue, err := instance.GetProperty("AccountName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
