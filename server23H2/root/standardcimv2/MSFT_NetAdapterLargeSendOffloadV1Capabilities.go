// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
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

// MSFT_NetAdapterLargeSendOffloadV1Capabilities struct
type MSFT_NetAdapterLargeSendOffloadV1Capabilities struct {
	*cim.WmiInstance

	//
	IPv4Encapsulation MSFT_NetAdapterLsoEncapsulationTypes

	//
	IPv4IpOptionsSupported bool

	//
	IPv4MaxOffloadSizeSupported uint32

	//
	IPv4MinSegmentCountSupported uint32

	//
	IPv4TcpOptionsSupported bool
}

func NewMSFT_NetAdapterLargeSendOffloadV1CapabilitiesEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterLargeSendOffloadV1Capabilities, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterLargeSendOffloadV1Capabilities{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_NetAdapterLargeSendOffloadV1CapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapterLargeSendOffloadV1Capabilities, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterLargeSendOffloadV1Capabilities{
		WmiInstance: tmp,
	}
	return
}

// SetIPv4Encapsulation sets the value of IPv4Encapsulation for the instance
func (instance *MSFT_NetAdapterLargeSendOffloadV1Capabilities) SetPropertyIPv4Encapsulation(value MSFT_NetAdapterLsoEncapsulationTypes) (err error) {
	return instance.SetProperty("IPv4Encapsulation", (value))
}

// GetIPv4Encapsulation gets the value of IPv4Encapsulation for the instance
func (instance *MSFT_NetAdapterLargeSendOffloadV1Capabilities) GetPropertyIPv4Encapsulation() (value MSFT_NetAdapterLsoEncapsulationTypes, err error) {
	retValue, err := instance.GetProperty("IPv4Encapsulation")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_NetAdapterLsoEncapsulationTypes)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_NetAdapterLsoEncapsulationTypes is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_NetAdapterLsoEncapsulationTypes(valuetmp)

	return
}

// SetIPv4IpOptionsSupported sets the value of IPv4IpOptionsSupported for the instance
func (instance *MSFT_NetAdapterLargeSendOffloadV1Capabilities) SetPropertyIPv4IpOptionsSupported(value bool) (err error) {
	return instance.SetProperty("IPv4IpOptionsSupported", (value))
}

// GetIPv4IpOptionsSupported gets the value of IPv4IpOptionsSupported for the instance
func (instance *MSFT_NetAdapterLargeSendOffloadV1Capabilities) GetPropertyIPv4IpOptionsSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("IPv4IpOptionsSupported")
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

// SetIPv4MaxOffloadSizeSupported sets the value of IPv4MaxOffloadSizeSupported for the instance
func (instance *MSFT_NetAdapterLargeSendOffloadV1Capabilities) SetPropertyIPv4MaxOffloadSizeSupported(value uint32) (err error) {
	return instance.SetProperty("IPv4MaxOffloadSizeSupported", (value))
}

// GetIPv4MaxOffloadSizeSupported gets the value of IPv4MaxOffloadSizeSupported for the instance
func (instance *MSFT_NetAdapterLargeSendOffloadV1Capabilities) GetPropertyIPv4MaxOffloadSizeSupported() (value uint32, err error) {
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
func (instance *MSFT_NetAdapterLargeSendOffloadV1Capabilities) SetPropertyIPv4MinSegmentCountSupported(value uint32) (err error) {
	return instance.SetProperty("IPv4MinSegmentCountSupported", (value))
}

// GetIPv4MinSegmentCountSupported gets the value of IPv4MinSegmentCountSupported for the instance
func (instance *MSFT_NetAdapterLargeSendOffloadV1Capabilities) GetPropertyIPv4MinSegmentCountSupported() (value uint32, err error) {
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

// SetIPv4TcpOptionsSupported sets the value of IPv4TcpOptionsSupported for the instance
func (instance *MSFT_NetAdapterLargeSendOffloadV1Capabilities) SetPropertyIPv4TcpOptionsSupported(value bool) (err error) {
	return instance.SetProperty("IPv4TcpOptionsSupported", (value))
}

// GetIPv4TcpOptionsSupported gets the value of IPv4TcpOptionsSupported for the instance
func (instance *MSFT_NetAdapterLargeSendOffloadV1Capabilities) GetPropertyIPv4TcpOptionsSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("IPv4TcpOptionsSupported")
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
