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

// MSFT_StorageSetting struct
type MSFT_StorageSetting struct {
	cim.WmiInstance

	//
	NewDiskPolicy uint16

	//
	ScrubPolicy uint32
}

// SetNewDiskPolicy sets the value of NewDiskPolicy for the instance
func (instance *MSFT_StorageSetting) SetPropertyNewDiskPolicy(value uint16) (err error) {
	return instance.SetProperty("NewDiskPolicy", value)
}

// GetNewDiskPolicy gets the value of NewDiskPolicy for the instance
func (instance *MSFT_StorageSetting) GetPropertyNewDiskPolicy() (value uint16, err error) {
	retValue, err := instance.GetProperty("NewDiskPolicy")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetScrubPolicy sets the value of ScrubPolicy for the instance
func (instance *MSFT_StorageSetting) SetPropertyScrubPolicy(value uint32) (err error) {
	return instance.SetProperty("ScrubPolicy", value)
}

// GetScrubPolicy gets the value of ScrubPolicy for the instance
func (instance *MSFT_StorageSetting) GetPropertyScrubPolicy() (value uint32, err error) {
	retValue, err := instance.GetProperty("ScrubPolicy")
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

// <param name="ReturnValue" type="uint32 "></param>
// <param name="StorageSetting" type="MSFT_StorageSetting "></param>
func (instance *MSFT_StorageSetting) Get( /* OUT */ StorageSetting MSFT_StorageSetting) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="NewDiskPolicy" type="uint16 "></param>
// <param name="ScrubPolicy" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageSetting) Set( /* IN */ NewDiskPolicy uint16,
	/* IN */ ScrubPolicy uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Set", NewDiskPolicy, ScrubPolicy)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageSetting) UpdateHostStorageCache() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("UpdateHostStorageCache")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
