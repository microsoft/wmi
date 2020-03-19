// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetAdapterPowerManagement_WakePattern_Bitmap struct
type MSFT_NetAdapterPowerManagement_WakePattern_Bitmap struct {
	*MSFT_NetAdapterPowerManagement_WakePattern

	//
	Mask []uint8

	//
	Pattern []uint8
}

func NewMSFT_NetAdapterPowerManagement_WakePattern_BitmapEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterPowerManagement_WakePattern_Bitmap, err error) {
	tmp, err := NewMSFT_NetAdapterPowerManagement_WakePatternEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterPowerManagement_WakePattern_Bitmap{
		MSFT_NetAdapterPowerManagement_WakePattern: tmp,
	}
	return
}

func NewMSFT_NetAdapterPowerManagement_WakePattern_BitmapEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapterPowerManagement_WakePattern_Bitmap, err error) {
	tmp, err := NewMSFT_NetAdapterPowerManagement_WakePatternEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterPowerManagement_WakePattern_Bitmap{
		MSFT_NetAdapterPowerManagement_WakePattern: tmp,
	}
	return
}

// SetMask sets the value of Mask for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern_Bitmap) SetPropertyMask(value []uint8) (err error) {
	return instance.SetProperty("Mask", value)
}

// GetMask gets the value of Mask for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern_Bitmap) GetPropertyMask() (value []uint8, err error) {
	retValue, err := instance.GetProperty("Mask")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPattern sets the value of Pattern for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern_Bitmap) SetPropertyPattern(value []uint8) (err error) {
	return instance.SetProperty("Pattern", value)
}

// GetPattern gets the value of Pattern for the instance
func (instance *MSFT_NetAdapterPowerManagement_WakePattern_Bitmap) GetPropertyPattern() (value []uint8, err error) {
	retValue, err := instance.GetProperty("Pattern")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}
