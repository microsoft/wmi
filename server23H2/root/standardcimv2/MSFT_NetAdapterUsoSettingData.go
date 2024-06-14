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

// MSFT_NetAdapterUsoSettingData struct
type MSFT_NetAdapterUsoSettingData struct {
	*MSFT_NetAdapterSettingData

	//
	IPv4Enabled bool

	//
	IPv6Enabled bool

	//
	UdpSegmentationOffloadHardwareCapabilities MSFT_NetAdapterUdpSegmentationOffloadCapabilities
}

func NewMSFT_NetAdapterUsoSettingDataEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterUsoSettingData, err error) {
	tmp, err := NewMSFT_NetAdapterSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterUsoSettingData{
		MSFT_NetAdapterSettingData: tmp,
	}
	return
}

func NewMSFT_NetAdapterUsoSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapterUsoSettingData, err error) {
	tmp, err := NewMSFT_NetAdapterSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterUsoSettingData{
		MSFT_NetAdapterSettingData: tmp,
	}
	return
}

// SetIPv4Enabled sets the value of IPv4Enabled for the instance
func (instance *MSFT_NetAdapterUsoSettingData) SetPropertyIPv4Enabled(value bool) (err error) {
	return instance.SetProperty("IPv4Enabled", (value))
}

// GetIPv4Enabled gets the value of IPv4Enabled for the instance
func (instance *MSFT_NetAdapterUsoSettingData) GetPropertyIPv4Enabled() (value bool, err error) {
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
func (instance *MSFT_NetAdapterUsoSettingData) SetPropertyIPv6Enabled(value bool) (err error) {
	return instance.SetProperty("IPv6Enabled", (value))
}

// GetIPv6Enabled gets the value of IPv6Enabled for the instance
func (instance *MSFT_NetAdapterUsoSettingData) GetPropertyIPv6Enabled() (value bool, err error) {
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

// SetUdpSegmentationOffloadHardwareCapabilities sets the value of UdpSegmentationOffloadHardwareCapabilities for the instance
func (instance *MSFT_NetAdapterUsoSettingData) SetPropertyUdpSegmentationOffloadHardwareCapabilities(value MSFT_NetAdapterUdpSegmentationOffloadCapabilities) (err error) {
	return instance.SetProperty("UdpSegmentationOffloadHardwareCapabilities", (value))
}

// GetUdpSegmentationOffloadHardwareCapabilities gets the value of UdpSegmentationOffloadHardwareCapabilities for the instance
func (instance *MSFT_NetAdapterUsoSettingData) GetPropertyUdpSegmentationOffloadHardwareCapabilities() (value MSFT_NetAdapterUdpSegmentationOffloadCapabilities, err error) {
	retValue, err := instance.GetProperty("UdpSegmentationOffloadHardwareCapabilities")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_NetAdapterUdpSegmentationOffloadCapabilities)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_NetAdapterUdpSegmentationOffloadCapabilities is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_NetAdapterUdpSegmentationOffloadCapabilities(valuetmp)

	return
}

//

// <param name="IPv4" type="bool "></param>
// <param name="IPv6" type="bool "></param>

// <param name="cmdletOutput" type="MSFT_NetAdapterUsoSettingData "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapterUsoSettingData) Enable( /* IN */ IPv4 bool,
	/* IN */ IPv6 bool,
	/* OUT */ cmdletOutput MSFT_NetAdapterUsoSettingData) (result uint32, err error) {
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

// <param name="cmdletOutput" type="MSFT_NetAdapterUsoSettingData "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapterUsoSettingData) Disable( /* IN */ IPv4 bool,
	/* IN */ IPv6 bool,
	/* OUT */ cmdletOutput MSFT_NetAdapterUsoSettingData) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Disable", IPv4, IPv6)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
