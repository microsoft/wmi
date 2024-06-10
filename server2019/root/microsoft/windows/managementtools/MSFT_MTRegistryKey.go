// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.ManagementTools
//
// ////////////////////////////////////////////
package managementtools

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_MTRegistryKey struct
type MSFT_MTRegistryKey struct {
	*MSFT_MTRegistryObject

	//
	Modified string

	//
	SubKeyCount uint32

	//
	ValueCount uint32
}

func NewMSFT_MTRegistryKeyEx1(instance *cim.WmiInstance) (newInstance *MSFT_MTRegistryKey, err error) {
	tmp, err := NewMSFT_MTRegistryObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_MTRegistryKey{
		MSFT_MTRegistryObject: tmp,
	}
	return
}

func NewMSFT_MTRegistryKeyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_MTRegistryKey, err error) {
	tmp, err := NewMSFT_MTRegistryObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_MTRegistryKey{
		MSFT_MTRegistryObject: tmp,
	}
	return
}

// SetModified sets the value of Modified for the instance
func (instance *MSFT_MTRegistryKey) SetPropertyModified(value string) (err error) {
	return instance.SetProperty("Modified", (value))
}

// GetModified gets the value of Modified for the instance
func (instance *MSFT_MTRegistryKey) GetPropertyModified() (value string, err error) {
	retValue, err := instance.GetProperty("Modified")
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

// SetSubKeyCount sets the value of SubKeyCount for the instance
func (instance *MSFT_MTRegistryKey) SetPropertySubKeyCount(value uint32) (err error) {
	return instance.SetProperty("SubKeyCount", (value))
}

// GetSubKeyCount gets the value of SubKeyCount for the instance
func (instance *MSFT_MTRegistryKey) GetPropertySubKeyCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("SubKeyCount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetValueCount sets the value of ValueCount for the instance
func (instance *MSFT_MTRegistryKey) SetPropertyValueCount(value uint32) (err error) {
	return instance.SetProperty("ValueCount", (value))
}

// GetValueCount gets the value of ValueCount for the instance
func (instance *MSFT_MTRegistryKey) GetPropertyValueCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ValueCount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

//

// <param name="Results" type="MSFT_MTRegistryKey []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MTRegistryKey) GetSubKeys( /* OUT */ Results []MSFT_MTRegistryKey) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetSubKeys")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Results" type="MSFT_MTRegistryValue []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MTRegistryKey) GetValues( /* OUT */ Results []MSFT_MTRegistryValue) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetValues")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="NewName" type="string "></param>

// <param name="Result" type="MSFT_MTRegistryKey "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MTRegistryKey) Rename( /* IN */ NewName string,
	/* OUT */ Result MSFT_MTRegistryKey) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Rename", NewName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Name" type="string "></param>

// <param name="Result" type="MSFT_MTRegistryKey "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MTRegistryKey) GetKey( /* IN */ Name string,
	/* OUT */ Result MSFT_MTRegistryKey) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetKey", Name)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
