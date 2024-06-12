// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetAddressFilter struct
type MSFT_NetAddressFilter struct {
	*CIM_FilterEntryBase

	//
	LocalAddress []string

	//
	RemoteAddress []string
}

func NewMSFT_NetAddressFilterEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAddressFilter, err error) {
	tmp, err := NewCIM_FilterEntryBaseEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAddressFilter{
		CIM_FilterEntryBase: tmp,
	}
	return
}

func NewMSFT_NetAddressFilterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAddressFilter, err error) {
	tmp, err := NewCIM_FilterEntryBaseEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAddressFilter{
		CIM_FilterEntryBase: tmp,
	}
	return
}

// SetLocalAddress sets the value of LocalAddress for the instance
func (instance *MSFT_NetAddressFilter) SetPropertyLocalAddress(value []string) (err error) {
	return instance.SetProperty("LocalAddress", (value))
}

// GetLocalAddress gets the value of LocalAddress for the instance
func (instance *MSFT_NetAddressFilter) GetPropertyLocalAddress() (value []string, err error) {
	retValue, err := instance.GetProperty("LocalAddress")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetRemoteAddress sets the value of RemoteAddress for the instance
func (instance *MSFT_NetAddressFilter) SetPropertyRemoteAddress(value []string) (err error) {
	return instance.SetProperty("RemoteAddress", (value))
}

// GetRemoteAddress gets the value of RemoteAddress for the instance
func (instance *MSFT_NetAddressFilter) GetPropertyRemoteAddress() (value []string, err error) {
	retValue, err := instance.GetProperty("RemoteAddress")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

//

// <param name="InterfaceIndex" type="uint32 "></param>
// <param name="RemoteAddress" type="string "></param>

// <param name="IsolationType" type="uint32 "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAddressFilter) QueryIsolationType( /* IN */ InterfaceIndex uint32,
	/* IN */ RemoteAddress string,
	/* OUT */ IsolationType uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("QueryIsolationType", InterfaceIndex, RemoteAddress)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
