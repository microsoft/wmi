// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetAdapterUroSettingData struct
type MSFT_NetAdapterUroSettingData struct {
	*MSFT_NetAdapterSettingData

	//
	Enabled bool

	//
	FailureReason uint32

	//
	Operational bool

	//
	UroHardwareCapabilities MSFT_NetAdapterUroCapabilities
}

func NewMSFT_NetAdapterUroSettingDataEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterUroSettingData, err error) {
	tmp, err := NewMSFT_NetAdapterSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterUroSettingData{
		MSFT_NetAdapterSettingData: tmp,
	}
	return
}

func NewMSFT_NetAdapterUroSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapterUroSettingData, err error) {
	tmp, err := NewMSFT_NetAdapterSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterUroSettingData{
		MSFT_NetAdapterSettingData: tmp,
	}
	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *MSFT_NetAdapterUroSettingData) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *MSFT_NetAdapterUroSettingData) GetPropertyEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("Enabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetFailureReason sets the value of FailureReason for the instance
func (instance *MSFT_NetAdapterUroSettingData) SetPropertyFailureReason(value uint32) (err error) {
	return instance.SetProperty("FailureReason", (value))
}

// GetFailureReason gets the value of FailureReason for the instance
func (instance *MSFT_NetAdapterUroSettingData) GetPropertyFailureReason() (value uint32, err error) {
	retValue, err := instance.GetProperty("FailureReason")
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

// SetOperational sets the value of Operational for the instance
func (instance *MSFT_NetAdapterUroSettingData) SetPropertyOperational(value bool) (err error) {
	return instance.SetProperty("Operational", (value))
}

// GetOperational gets the value of Operational for the instance
func (instance *MSFT_NetAdapterUroSettingData) GetPropertyOperational() (value bool, err error) {
	retValue, err := instance.GetProperty("Operational")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetUroHardwareCapabilities sets the value of UroHardwareCapabilities for the instance
func (instance *MSFT_NetAdapterUroSettingData) SetPropertyUroHardwareCapabilities(value MSFT_NetAdapterUroCapabilities) (err error) {
	return instance.SetProperty("UroHardwareCapabilities", (value))
}

// GetUroHardwareCapabilities gets the value of UroHardwareCapabilities for the instance
func (instance *MSFT_NetAdapterUroSettingData) GetPropertyUroHardwareCapabilities() (value MSFT_NetAdapterUroCapabilities, err error) {
	retValue, err := instance.GetProperty("UroHardwareCapabilities")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_NetAdapterUroCapabilities)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_NetAdapterUroCapabilities is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_NetAdapterUroCapabilities(valuetmp)

	return
}

//

// <param name="Enable" type="bool "></param>

// <param name="cmdletOutput" type="MSFT_NetAdapterUroSettingData "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapterUroSettingData) Enable( /* IN */ Enable bool,
	/* OUT */ cmdletOutput MSFT_NetAdapterUroSettingData) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Enable", Enable)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Disable" type="bool "></param>

// <param name="cmdletOutput" type="MSFT_NetAdapterUroSettingData "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapterUroSettingData) Disable( /* IN */ Disable bool,
	/* OUT */ cmdletOutput MSFT_NetAdapterUroSettingData) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Disable", Disable)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
