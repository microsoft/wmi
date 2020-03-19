// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.SMB
//////////////////////////////////////////////
package smb

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_SmbShareAccessControlEntry struct
type MSFT_SmbShareAccessControlEntry struct {
	*cim.WmiInstance

	//
	AccessControlType SmbShareAccessControlEntry_AccessControlType

	//
	AccessRight SmbShareAccessControlEntry_AccessRight

	//
	AccountName string

	//
	Name string

	//
	ScopeName string
}

func NewMSFT_SmbShareAccessControlEntryEx1(instance *cim.WmiInstance) (newInstance *MSFT_SmbShareAccessControlEntry, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_SmbShareAccessControlEntry{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_SmbShareAccessControlEntryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_SmbShareAccessControlEntry, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_SmbShareAccessControlEntry{
		WmiInstance: tmp,
	}
	return
}

// SetAccessControlType sets the value of AccessControlType for the instance
func (instance *MSFT_SmbShareAccessControlEntry) SetPropertyAccessControlType(value SmbShareAccessControlEntry_AccessControlType) (err error) {
	return instance.SetProperty("AccessControlType", value)
}

// GetAccessControlType gets the value of AccessControlType for the instance
func (instance *MSFT_SmbShareAccessControlEntry) GetPropertyAccessControlType() (value SmbShareAccessControlEntry_AccessControlType, err error) {
	retValue, err := instance.GetProperty("AccessControlType")
	if err != nil {
		return
	}
	value, ok := retValue.(SmbShareAccessControlEntry_AccessControlType)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAccessRight sets the value of AccessRight for the instance
func (instance *MSFT_SmbShareAccessControlEntry) SetPropertyAccessRight(value SmbShareAccessControlEntry_AccessRight) (err error) {
	return instance.SetProperty("AccessRight", value)
}

// GetAccessRight gets the value of AccessRight for the instance
func (instance *MSFT_SmbShareAccessControlEntry) GetPropertyAccessRight() (value SmbShareAccessControlEntry_AccessRight, err error) {
	retValue, err := instance.GetProperty("AccessRight")
	if err != nil {
		return
	}
	value, ok := retValue.(SmbShareAccessControlEntry_AccessRight)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAccountName sets the value of AccountName for the instance
func (instance *MSFT_SmbShareAccessControlEntry) SetPropertyAccountName(value string) (err error) {
	return instance.SetProperty("AccountName", value)
}

// GetAccountName gets the value of AccountName for the instance
func (instance *MSFT_SmbShareAccessControlEntry) GetPropertyAccountName() (value string, err error) {
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

// SetName sets the value of Name for the instance
func (instance *MSFT_SmbShareAccessControlEntry) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_SmbShareAccessControlEntry) GetPropertyName() (value string, err error) {
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

// SetScopeName sets the value of ScopeName for the instance
func (instance *MSFT_SmbShareAccessControlEntry) SetPropertyScopeName(value string) (err error) {
	return instance.SetProperty("ScopeName", value)
}

// GetScopeName gets the value of ScopeName for the instance
func (instance *MSFT_SmbShareAccessControlEntry) GetPropertyScopeName() (value string, err error) {
	retValue, err := instance.GetProperty("ScopeName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
