// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetAdapterLsoSettingData struct
type MSFT_NetAdapterLsoSettingData struct {
	*MSFT_NetAdapterSettingData

	//
	IPv4Enabled bool

	//
	IPv6Enabled bool

	//
	LargeSendOffloadV1HardwareCapabilities MSFT_NetAdapterLargeSendOffloadV1Capabilities

	//
	LargeSendOffloadV2HardwareCapabilities MSFT_NetAdapterLargeSendOffloadV2Capabilities

	//
	MaximumLsoVersionSupported uint32

	//
	V1IPv4Enabled bool
}

func NewMSFT_NetAdapterLsoSettingDataEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterLsoSettingData, err error) {
	tmp, err := NewMSFT_NetAdapterSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterLsoSettingData{
		MSFT_NetAdapterSettingData: tmp,
	}
	return
}

func NewMSFT_NetAdapterLsoSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapterLsoSettingData, err error) {
	tmp, err := NewMSFT_NetAdapterSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterLsoSettingData{
		MSFT_NetAdapterSettingData: tmp,
	}
	return
}

// SetIPv4Enabled sets the value of IPv4Enabled for the instance
func (instance *MSFT_NetAdapterLsoSettingData) SetPropertyIPv4Enabled(value bool) (err error) {
	return instance.SetProperty("IPv4Enabled", (value))
}

// GetIPv4Enabled gets the value of IPv4Enabled for the instance
func (instance *MSFT_NetAdapterLsoSettingData) GetPropertyIPv4Enabled() (value bool, err error) {
	retValue, err := instance.GetProperty("IPv4Enabled")
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

// SetIPv6Enabled sets the value of IPv6Enabled for the instance
func (instance *MSFT_NetAdapterLsoSettingData) SetPropertyIPv6Enabled(value bool) (err error) {
	return instance.SetProperty("IPv6Enabled", (value))
}

// GetIPv6Enabled gets the value of IPv6Enabled for the instance
func (instance *MSFT_NetAdapterLsoSettingData) GetPropertyIPv6Enabled() (value bool, err error) {
	retValue, err := instance.GetProperty("IPv6Enabled")
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

// SetLargeSendOffloadV1HardwareCapabilities sets the value of LargeSendOffloadV1HardwareCapabilities for the instance
func (instance *MSFT_NetAdapterLsoSettingData) SetPropertyLargeSendOffloadV1HardwareCapabilities(value MSFT_NetAdapterLargeSendOffloadV1Capabilities) (err error) {
	return instance.SetProperty("LargeSendOffloadV1HardwareCapabilities", (value))
}

// GetLargeSendOffloadV1HardwareCapabilities gets the value of LargeSendOffloadV1HardwareCapabilities for the instance
func (instance *MSFT_NetAdapterLsoSettingData) GetPropertyLargeSendOffloadV1HardwareCapabilities() (value MSFT_NetAdapterLargeSendOffloadV1Capabilities, err error) {
	retValue, err := instance.GetProperty("LargeSendOffloadV1HardwareCapabilities")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_NetAdapterLargeSendOffloadV1Capabilities)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_NetAdapterLargeSendOffloadV1Capabilities is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_NetAdapterLargeSendOffloadV1Capabilities(valuetmp)

	return
}

// SetLargeSendOffloadV2HardwareCapabilities sets the value of LargeSendOffloadV2HardwareCapabilities for the instance
func (instance *MSFT_NetAdapterLsoSettingData) SetPropertyLargeSendOffloadV2HardwareCapabilities(value MSFT_NetAdapterLargeSendOffloadV2Capabilities) (err error) {
	return instance.SetProperty("LargeSendOffloadV2HardwareCapabilities", (value))
}

// GetLargeSendOffloadV2HardwareCapabilities gets the value of LargeSendOffloadV2HardwareCapabilities for the instance
func (instance *MSFT_NetAdapterLsoSettingData) GetPropertyLargeSendOffloadV2HardwareCapabilities() (value MSFT_NetAdapterLargeSendOffloadV2Capabilities, err error) {
	retValue, err := instance.GetProperty("LargeSendOffloadV2HardwareCapabilities")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_NetAdapterLargeSendOffloadV2Capabilities)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_NetAdapterLargeSendOffloadV2Capabilities is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_NetAdapterLargeSendOffloadV2Capabilities(valuetmp)

	return
}

// SetMaximumLsoVersionSupported sets the value of MaximumLsoVersionSupported for the instance
func (instance *MSFT_NetAdapterLsoSettingData) SetPropertyMaximumLsoVersionSupported(value uint32) (err error) {
	return instance.SetProperty("MaximumLsoVersionSupported", (value))
}

// GetMaximumLsoVersionSupported gets the value of MaximumLsoVersionSupported for the instance
func (instance *MSFT_NetAdapterLsoSettingData) GetPropertyMaximumLsoVersionSupported() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumLsoVersionSupported")
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

// SetV1IPv4Enabled sets the value of V1IPv4Enabled for the instance
func (instance *MSFT_NetAdapterLsoSettingData) SetPropertyV1IPv4Enabled(value bool) (err error) {
	return instance.SetProperty("V1IPv4Enabled", (value))
}

// GetV1IPv4Enabled gets the value of V1IPv4Enabled for the instance
func (instance *MSFT_NetAdapterLsoSettingData) GetPropertyV1IPv4Enabled() (value bool, err error) {
	retValue, err := instance.GetProperty("V1IPv4Enabled")
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

//

// <param name="IPv4" type="bool "></param>
// <param name="IPv6" type="bool "></param>

// <param name="cmdletOutput" type="MSFT_NetAdapterLsoSettingData "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapterLsoSettingData) Enable( /* IN */ IPv4 bool,
	/* IN */ IPv6 bool,
	/* OUT */ cmdletOutput MSFT_NetAdapterLsoSettingData) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Enable", IPv4, IPv6)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="IPv4" type="bool "></param>
// <param name="IPv6" type="bool "></param>

// <param name="cmdletOutput" type="MSFT_NetAdapterLsoSettingData "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapterLsoSettingData) Disable( /* IN */ IPv4 bool,
	/* IN */ IPv6 bool,
	/* OUT */ cmdletOutput MSFT_NetAdapterLsoSettingData) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Disable", IPv4, IPv6)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
