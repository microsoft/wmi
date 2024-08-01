// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_FileShareAccessControlEntry struct
type MSFT_FileShareAccessControlEntry struct {
	*cim.WmiInstance

	// Denotes the access type ( Allow, Deny ).
	AccessControlType FileShareAccessControlEntry_AccessControlType

	// Denotes the access right.
	AccessRight FileShareAccessControlEntry_AccessRight

	// The name of the account to which the access right is granted.
	AccountName string
}

func NewMSFT_FileShareAccessControlEntryEx1(instance *cim.WmiInstance) (newInstance *MSFT_FileShareAccessControlEntry, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_FileShareAccessControlEntry{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_FileShareAccessControlEntryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_FileShareAccessControlEntry, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_FileShareAccessControlEntry{
		WmiInstance: tmp,
	}
	return
}

// SetAccessControlType sets the value of AccessControlType for the instance
func (instance *MSFT_FileShareAccessControlEntry) SetPropertyAccessControlType(value FileShareAccessControlEntry_AccessControlType) (err error) {
	return instance.SetProperty("AccessControlType", (value))
}

// GetAccessControlType gets the value of AccessControlType for the instance
func (instance *MSFT_FileShareAccessControlEntry) GetPropertyAccessControlType() (value FileShareAccessControlEntry_AccessControlType, err error) {
	retValue, err := instance.GetProperty("AccessControlType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FileShareAccessControlEntry_AccessControlType(valuetmp)

	return
}

// SetAccessRight sets the value of AccessRight for the instance
func (instance *MSFT_FileShareAccessControlEntry) SetPropertyAccessRight(value FileShareAccessControlEntry_AccessRight) (err error) {
	return instance.SetProperty("AccessRight", (value))
}

// GetAccessRight gets the value of AccessRight for the instance
func (instance *MSFT_FileShareAccessControlEntry) GetPropertyAccessRight() (value FileShareAccessControlEntry_AccessRight, err error) {
	retValue, err := instance.GetProperty("AccessRight")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FileShareAccessControlEntry_AccessRight(valuetmp)

	return
}

// SetAccountName sets the value of AccountName for the instance
func (instance *MSFT_FileShareAccessControlEntry) SetPropertyAccountName(value string) (err error) {
	return instance.SetProperty("AccountName", (value))
}

// GetAccountName gets the value of AccountName for the instance
func (instance *MSFT_FileShareAccessControlEntry) GetPropertyAccountName() (value string, err error) {
	retValue, err := instance.GetProperty("AccountName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}
