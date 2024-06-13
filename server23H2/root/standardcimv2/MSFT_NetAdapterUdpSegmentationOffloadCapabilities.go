// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetAdapterUdpSegmentationOffloadCapabilities struct
type MSFT_NetAdapterUdpSegmentationOffloadCapabilities struct {
	*cim.WmiInstance

	//
	IPv4Encapsulation MSFT_NetAdapterUsoEncapsulationTypes

	//
	IPv4MaxOffloadSizeSupported uint32

	//
	IPv4MinSegmentCountSupported uint32

	//
	IPv4SubMssFinalSegmentSupported bool

	//
	IPv6Encapsulation MSFT_NetAdapterUsoEncapsulationTypes

	//
	IPv6IpExtensionHeadersSupported bool

	//
	IPv6MaxOffLoadSizeSupported uint32

	//
	IPv6MinSegmentCountSupported uint32

	//
	IPv6SubMssFinalSegmentSupported bool
}

func NewMSFT_NetAdapterUdpSegmentationOffloadCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterUdpSegmentationOffloadCapabilities, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterUdpSegmentationOffloadCapabilities{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_NetAdapterUdpSegmentationOffloadCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapterUdpSegmentationOffloadCapabilities, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterUdpSegmentationOffloadCapabilities{
		WmiInstance: tmp,
	}
	return
}

// SetIPv4Encapsulation sets the value of IPv4Encapsulation for the instance
func (instance *MSFT_NetAdapterUdpSegmentationOffloadCapabilities) SetPropertyIPv4Encapsulation(value MSFT_NetAdapterUsoEncapsulationTypes) (err error) {
	return instance.SetProperty("IPv4Encapsulation", (value))
}

// GetIPv4Encapsulation gets the value of IPv4Encapsulation for the instance
func (instance *MSFT_NetAdapterUdpSegmentationOffloadCapabilities) GetPropertyIPv4Encapsulation() (value MSFT_NetAdapterUsoEncapsulationTypes, err error) {
	retValue, err := instance.GetProperty("IPv4Encapsulation")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_NetAdapterUsoEncapsulationTypes)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_NetAdapterUsoEncapsulationTypes is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_NetAdapterUsoEncapsulationTypes(valuetmp)

	return
}

// SetIPv4MaxOffloadSizeSupported sets the value of IPv4MaxOffloadSizeSupported for the instance
func (instance *MSFT_NetAdapterUdpSegmentationOffloadCapabilities) SetPropertyIPv4MaxOffloadSizeSupported(value uint32) (err error) {
	return instance.SetProperty("IPv4MaxOffloadSizeSupported", (value))
}

// GetIPv4MaxOffloadSizeSupported gets the value of IPv4MaxOffloadSizeSupported for the instance
func (instance *MSFT_NetAdapterUdpSegmentationOffloadCapabilities) GetPropertyIPv4MaxOffloadSizeSupported() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv4MaxOffloadSizeSupported")
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

// SetIPv4MinSegmentCountSupported sets the value of IPv4MinSegmentCountSupported for the instance
func (instance *MSFT_NetAdapterUdpSegmentationOffloadCapabilities) SetPropertyIPv4MinSegmentCountSupported(value uint32) (err error) {
	return instance.SetProperty("IPv4MinSegmentCountSupported", (value))
}

// GetIPv4MinSegmentCountSupported gets the value of IPv4MinSegmentCountSupported for the instance
func (instance *MSFT_NetAdapterUdpSegmentationOffloadCapabilities) GetPropertyIPv4MinSegmentCountSupported() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv4MinSegmentCountSupported")
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

// SetIPv4SubMssFinalSegmentSupported sets the value of IPv4SubMssFinalSegmentSupported for the instance
func (instance *MSFT_NetAdapterUdpSegmentationOffloadCapabilities) SetPropertyIPv4SubMssFinalSegmentSupported(value bool) (err error) {
	return instance.SetProperty("IPv4SubMssFinalSegmentSupported", (value))
}

// GetIPv4SubMssFinalSegmentSupported gets the value of IPv4SubMssFinalSegmentSupported for the instance
func (instance *MSFT_NetAdapterUdpSegmentationOffloadCapabilities) GetPropertyIPv4SubMssFinalSegmentSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("IPv4SubMssFinalSegmentSupported")
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

// SetIPv6Encapsulation sets the value of IPv6Encapsulation for the instance
func (instance *MSFT_NetAdapterUdpSegmentationOffloadCapabilities) SetPropertyIPv6Encapsulation(value MSFT_NetAdapterUsoEncapsulationTypes) (err error) {
	return instance.SetProperty("IPv6Encapsulation", (value))
}

// GetIPv6Encapsulation gets the value of IPv6Encapsulation for the instance
func (instance *MSFT_NetAdapterUdpSegmentationOffloadCapabilities) GetPropertyIPv6Encapsulation() (value MSFT_NetAdapterUsoEncapsulationTypes, err error) {
	retValue, err := instance.GetProperty("IPv6Encapsulation")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_NetAdapterUsoEncapsulationTypes)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_NetAdapterUsoEncapsulationTypes is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_NetAdapterUsoEncapsulationTypes(valuetmp)

	return
}

// SetIPv6IpExtensionHeadersSupported sets the value of IPv6IpExtensionHeadersSupported for the instance
func (instance *MSFT_NetAdapterUdpSegmentationOffloadCapabilities) SetPropertyIPv6IpExtensionHeadersSupported(value bool) (err error) {
	return instance.SetProperty("IPv6IpExtensionHeadersSupported", (value))
}

// GetIPv6IpExtensionHeadersSupported gets the value of IPv6IpExtensionHeadersSupported for the instance
func (instance *MSFT_NetAdapterUdpSegmentationOffloadCapabilities) GetPropertyIPv6IpExtensionHeadersSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("IPv6IpExtensionHeadersSupported")
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

// SetIPv6MaxOffLoadSizeSupported sets the value of IPv6MaxOffLoadSizeSupported for the instance
func (instance *MSFT_NetAdapterUdpSegmentationOffloadCapabilities) SetPropertyIPv6MaxOffLoadSizeSupported(value uint32) (err error) {
	return instance.SetProperty("IPv6MaxOffLoadSizeSupported", (value))
}

// GetIPv6MaxOffLoadSizeSupported gets the value of IPv6MaxOffLoadSizeSupported for the instance
func (instance *MSFT_NetAdapterUdpSegmentationOffloadCapabilities) GetPropertyIPv6MaxOffLoadSizeSupported() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv6MaxOffLoadSizeSupported")
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

// SetIPv6MinSegmentCountSupported sets the value of IPv6MinSegmentCountSupported for the instance
func (instance *MSFT_NetAdapterUdpSegmentationOffloadCapabilities) SetPropertyIPv6MinSegmentCountSupported(value uint32) (err error) {
	return instance.SetProperty("IPv6MinSegmentCountSupported", (value))
}

// GetIPv6MinSegmentCountSupported gets the value of IPv6MinSegmentCountSupported for the instance
func (instance *MSFT_NetAdapterUdpSegmentationOffloadCapabilities) GetPropertyIPv6MinSegmentCountSupported() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv6MinSegmentCountSupported")
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

// SetIPv6SubMssFinalSegmentSupported sets the value of IPv6SubMssFinalSegmentSupported for the instance
func (instance *MSFT_NetAdapterUdpSegmentationOffloadCapabilities) SetPropertyIPv6SubMssFinalSegmentSupported(value bool) (err error) {
	return instance.SetProperty("IPv6SubMssFinalSegmentSupported", (value))
}

// GetIPv6SubMssFinalSegmentSupported gets the value of IPv6SubMssFinalSegmentSupported for the instance
func (instance *MSFT_NetAdapterUdpSegmentationOffloadCapabilities) GetPropertyIPv6SubMssFinalSegmentSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("IPv6SubMssFinalSegmentSupported")
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
