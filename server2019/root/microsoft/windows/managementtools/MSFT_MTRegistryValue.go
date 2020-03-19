// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.ManagementTools
//////////////////////////////////////////////
package managementtools

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_MTRegistryValue struct
type MSFT_MTRegistryValue struct {
	*MSFT_MTRegistryObject

	//
	Status uint16

	//
	Type uint32
}

func NewMSFT_MTRegistryValueEx1(instance *cim.WmiInstance) (newInstance *MSFT_MTRegistryValue, err error) {
	tmp, err := NewMSFT_MTRegistryObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_MTRegistryValue{
		MSFT_MTRegistryObject: tmp,
	}
	return
}

func NewMSFT_MTRegistryValueEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_MTRegistryValue, err error) {
	tmp, err := NewMSFT_MTRegistryObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_MTRegistryValue{
		MSFT_MTRegistryObject: tmp,
	}
	return
}

// SetStatus sets the value of Status for the instance
func (instance *MSFT_MTRegistryValue) SetPropertyStatus(value uint16) (err error) {
	return instance.SetProperty("Status", value)
}

// GetStatus gets the value of Status for the instance
func (instance *MSFT_MTRegistryValue) GetPropertyStatus() (value uint16, err error) {
	retValue, err := instance.GetProperty("Status")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetType sets the value of Type for the instance
func (instance *MSFT_MTRegistryValue) SetPropertyType(value uint32) (err error) {
	return instance.SetProperty("Type", value)
}

// GetType gets the value of Type for the instance
func (instance *MSFT_MTRegistryValue) GetPropertyType() (value uint32, err error) {
	retValue, err := instance.GetProperty("Type")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="NewName" type="string "></param>

// <param name="Result" type="MSFT_MTRegistryValue "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MTRegistryValue) Rename( /* IN */ NewName string,
	/* OUT */ Result MSFT_MTRegistryValue) (result uint32, err error) {
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

// <param name="Result" type="MSFT_MTRegistryValue "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MTRegistryValue) GetValue( /* IN */ Name string,
	/* OUT */ Result MSFT_MTRegistryValue) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetValue", Name)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
