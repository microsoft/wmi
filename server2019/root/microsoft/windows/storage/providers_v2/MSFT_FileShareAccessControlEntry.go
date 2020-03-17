// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_FileShareAccessControlEntry struct
type MSFT_FileShareAccessControlEntry struct {
	cim.WmiInstance

	// Denotes the access type ( Allow, Deny ).
	AccessControlType FileShareAccessControlEntry_AccessControlType

	// Denotes the access right.
	AccessRight FileShareAccessControlEntry_AccessRight

	// The name of the account to which the access right is granted.
	AccountName string
}

// SetAccessControlType sets the value of AccessControlType for the instance
func (instance *MSFT_FileShareAccessControlEntry) SetPropertyAccessControlType(value FileShareAccessControlEntry_AccessControlType) (err error) {
	return instance.SetProperty("AccessControlType", value)
}

// GetAccessControlType gets the value of AccessControlType for the instance
func (instance *MSFT_FileShareAccessControlEntry) GetPropertyAccessControlType() (value FileShareAccessControlEntry_AccessControlType, err error) {
	retValue, err := instance.GetProperty("AccessControlType")
	if err != nil {
		return
	}
	value, ok := retValue.(FileShareAccessControlEntry_AccessControlType)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAccessRight sets the value of AccessRight for the instance
func (instance *MSFT_FileShareAccessControlEntry) SetPropertyAccessRight(value FileShareAccessControlEntry_AccessRight) (err error) {
	return instance.SetProperty("AccessRight", value)
}

// GetAccessRight gets the value of AccessRight for the instance
func (instance *MSFT_FileShareAccessControlEntry) GetPropertyAccessRight() (value FileShareAccessControlEntry_AccessRight, err error) {
	retValue, err := instance.GetProperty("AccessRight")
	if err != nil {
		return
	}
	value, ok := retValue.(FileShareAccessControlEntry_AccessRight)
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
