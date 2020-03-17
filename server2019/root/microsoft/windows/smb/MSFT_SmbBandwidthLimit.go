// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.SMB
//////////////////////////////////////////////
package smb

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_SmbBandwidthLimit struct
type MSFT_SmbBandwidthLimit struct {
	cim.WmiInstance

	//
	BytesPerSecond uint64

	//
	Category SmbBandwidthLimit_Category
}

// SetBytesPerSecond sets the value of BytesPerSecond for the instance
func (instance *MSFT_SmbBandwidthLimit) SetPropertyBytesPerSecond(value uint64) (err error) {
	return instance.SetProperty("BytesPerSecond", value)
}

// GetBytesPerSecond gets the value of BytesPerSecond for the instance
func (instance *MSFT_SmbBandwidthLimit) GetPropertyBytesPerSecond() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesPerSecond")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCategory sets the value of Category for the instance
func (instance *MSFT_SmbBandwidthLimit) SetPropertyCategory(value SmbBandwidthLimit_Category) (err error) {
	return instance.SetProperty("Category", value)
}

// GetCategory gets the value of Category for the instance
func (instance *MSFT_SmbBandwidthLimit) GetPropertyCategory() (value SmbBandwidthLimit_Category, err error) {
	retValue, err := instance.GetProperty("Category")
	if err != nil {
		return
	}
	value, ok := retValue.(SmbBandwidthLimit_Category)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="BytesPerSecond" type="uint64 "></param>
// <param name="Category" type="uint32 "></param>

// <param name="Output" type="MSFT_SmbBandwidthLimit []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbBandwidthLimit) Set( /* IN */ Category uint32,
	/* IN */ BytesPerSecond uint64,
	/* OUT */ Output []MSFT_SmbBandwidthLimit) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Set", Category, BytesPerSecond)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
