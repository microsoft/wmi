// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
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

// MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilities struct
type MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilities struct {
	*cim.WmiInstance

	//
	LsoV2Supported uint32

	//
	ReceiveChecksumOffloadSupported uint32

	//
	RssSupported uint32

	//
	TransmitChecksumOffloadSupported uint32

	//
	VmqSupported uint32
}

func NewMSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilities, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilities{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilities, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilities{
		WmiInstance: tmp,
	}
	return
}

// SetLsoV2Supported sets the value of LsoV2Supported for the instance
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilities) SetPropertyLsoV2Supported(value uint32) (err error) {
	return instance.SetProperty("LsoV2Supported", (value))
}

// GetLsoV2Supported gets the value of LsoV2Supported for the instance
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilities) GetPropertyLsoV2Supported() (value uint32, err error) {
	retValue, err := instance.GetProperty("LsoV2Supported")
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

// SetReceiveChecksumOffloadSupported sets the value of ReceiveChecksumOffloadSupported for the instance
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilities) SetPropertyReceiveChecksumOffloadSupported(value uint32) (err error) {
	return instance.SetProperty("ReceiveChecksumOffloadSupported", (value))
}

// GetReceiveChecksumOffloadSupported gets the value of ReceiveChecksumOffloadSupported for the instance
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilities) GetPropertyReceiveChecksumOffloadSupported() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReceiveChecksumOffloadSupported")
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

// SetRssSupported sets the value of RssSupported for the instance
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilities) SetPropertyRssSupported(value uint32) (err error) {
	return instance.SetProperty("RssSupported", (value))
}

// GetRssSupported gets the value of RssSupported for the instance
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilities) GetPropertyRssSupported() (value uint32, err error) {
	retValue, err := instance.GetProperty("RssSupported")
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

// SetTransmitChecksumOffloadSupported sets the value of TransmitChecksumOffloadSupported for the instance
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilities) SetPropertyTransmitChecksumOffloadSupported(value uint32) (err error) {
	return instance.SetProperty("TransmitChecksumOffloadSupported", (value))
}

// GetTransmitChecksumOffloadSupported gets the value of TransmitChecksumOffloadSupported for the instance
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilities) GetPropertyTransmitChecksumOffloadSupported() (value uint32, err error) {
	retValue, err := instance.GetProperty("TransmitChecksumOffloadSupported")
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

// SetVmqSupported sets the value of VmqSupported for the instance
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilities) SetPropertyVmqSupported(value uint32) (err error) {
	return instance.SetProperty("VmqSupported", (value))
}

// GetVmqSupported gets the value of VmqSupported for the instance
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilities) GetPropertyVmqSupported() (value uint32, err error) {
	retValue, err := instance.GetProperty("VmqSupported")
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
