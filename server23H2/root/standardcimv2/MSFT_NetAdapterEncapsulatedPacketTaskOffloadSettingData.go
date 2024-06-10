// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.StandardCimv2
//
// ////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetAdapterEncapsulatedPacketTaskOffloadSettingData struct
type MSFT_NetAdapterEncapsulatedPacketTaskOffloadSettingData struct {
	*MSFT_NetAdapterSettingData

	//
	EncapsulatedPacketTaskOffloadHardwareCapabilitiesNvgre MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilities

	//
	EncapsulatedPacketTaskOffloadHardwareCapabilitiesVxlan MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilitiesEx

	//
	EncapsulationType uint16

	//
	IsVxlanUDPPortConfigurable bool

	//
	NvgreEncapsulatedPacketTaskOffloadEnabled bool

	//
	VxlanEncapsulatedPacketTaskOffloadEnabled bool

	//
	VxlanUDPPortNumber uint16
}

func NewMSFT_NetAdapterEncapsulatedPacketTaskOffloadSettingDataEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadSettingData, err error) {
	tmp, err := NewMSFT_NetAdapterSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterEncapsulatedPacketTaskOffloadSettingData{
		MSFT_NetAdapterSettingData: tmp,
	}
	return
}

func NewMSFT_NetAdapterEncapsulatedPacketTaskOffloadSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadSettingData, err error) {
	tmp, err := NewMSFT_NetAdapterSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterEncapsulatedPacketTaskOffloadSettingData{
		MSFT_NetAdapterSettingData: tmp,
	}
	return
}

// SetEncapsulatedPacketTaskOffloadHardwareCapabilitiesNvgre sets the value of EncapsulatedPacketTaskOffloadHardwareCapabilitiesNvgre for the instance
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadSettingData) SetPropertyEncapsulatedPacketTaskOffloadHardwareCapabilitiesNvgre(value MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilities) (err error) {
	return instance.SetProperty("EncapsulatedPacketTaskOffloadHardwareCapabilitiesNvgre", (value))
}

// GetEncapsulatedPacketTaskOffloadHardwareCapabilitiesNvgre gets the value of EncapsulatedPacketTaskOffloadHardwareCapabilitiesNvgre for the instance
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadSettingData) GetPropertyEncapsulatedPacketTaskOffloadHardwareCapabilitiesNvgre() (value MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilities, err error) {
	retValue, err := instance.GetProperty("EncapsulatedPacketTaskOffloadHardwareCapabilitiesNvgre")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilities)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilities is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilities(valuetmp)

	return
}

// SetEncapsulatedPacketTaskOffloadHardwareCapabilitiesVxlan sets the value of EncapsulatedPacketTaskOffloadHardwareCapabilitiesVxlan for the instance
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadSettingData) SetPropertyEncapsulatedPacketTaskOffloadHardwareCapabilitiesVxlan(value MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilitiesEx) (err error) {
	return instance.SetProperty("EncapsulatedPacketTaskOffloadHardwareCapabilitiesVxlan", (value))
}

// GetEncapsulatedPacketTaskOffloadHardwareCapabilitiesVxlan gets the value of EncapsulatedPacketTaskOffloadHardwareCapabilitiesVxlan for the instance
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadSettingData) GetPropertyEncapsulatedPacketTaskOffloadHardwareCapabilitiesVxlan() (value MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilitiesEx, err error) {
	retValue, err := instance.GetProperty("EncapsulatedPacketTaskOffloadHardwareCapabilitiesVxlan")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilitiesEx)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilitiesEx is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_NetAdapterEncapsulatedPacketTaskOffloadCapabilitiesEx(valuetmp)

	return
}

// SetEncapsulationType sets the value of EncapsulationType for the instance
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadSettingData) SetPropertyEncapsulationType(value uint16) (err error) {
	return instance.SetProperty("EncapsulationType", (value))
}

// GetEncapsulationType gets the value of EncapsulationType for the instance
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadSettingData) GetPropertyEncapsulationType() (value uint16, err error) {
	retValue, err := instance.GetProperty("EncapsulationType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetIsVxlanUDPPortConfigurable sets the value of IsVxlanUDPPortConfigurable for the instance
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadSettingData) SetPropertyIsVxlanUDPPortConfigurable(value bool) (err error) {
	return instance.SetProperty("IsVxlanUDPPortConfigurable", (value))
}

// GetIsVxlanUDPPortConfigurable gets the value of IsVxlanUDPPortConfigurable for the instance
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadSettingData) GetPropertyIsVxlanUDPPortConfigurable() (value bool, err error) {
	retValue, err := instance.GetProperty("IsVxlanUDPPortConfigurable")
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

// SetNvgreEncapsulatedPacketTaskOffloadEnabled sets the value of NvgreEncapsulatedPacketTaskOffloadEnabled for the instance
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadSettingData) SetPropertyNvgreEncapsulatedPacketTaskOffloadEnabled(value bool) (err error) {
	return instance.SetProperty("NvgreEncapsulatedPacketTaskOffloadEnabled", (value))
}

// GetNvgreEncapsulatedPacketTaskOffloadEnabled gets the value of NvgreEncapsulatedPacketTaskOffloadEnabled for the instance
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadSettingData) GetPropertyNvgreEncapsulatedPacketTaskOffloadEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("NvgreEncapsulatedPacketTaskOffloadEnabled")
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

// SetVxlanEncapsulatedPacketTaskOffloadEnabled sets the value of VxlanEncapsulatedPacketTaskOffloadEnabled for the instance
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadSettingData) SetPropertyVxlanEncapsulatedPacketTaskOffloadEnabled(value bool) (err error) {
	return instance.SetProperty("VxlanEncapsulatedPacketTaskOffloadEnabled", (value))
}

// GetVxlanEncapsulatedPacketTaskOffloadEnabled gets the value of VxlanEncapsulatedPacketTaskOffloadEnabled for the instance
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadSettingData) GetPropertyVxlanEncapsulatedPacketTaskOffloadEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("VxlanEncapsulatedPacketTaskOffloadEnabled")
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

// SetVxlanUDPPortNumber sets the value of VxlanUDPPortNumber for the instance
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadSettingData) SetPropertyVxlanUDPPortNumber(value uint16) (err error) {
	return instance.SetProperty("VxlanUDPPortNumber", (value))
}

// GetVxlanUDPPortNumber gets the value of VxlanUDPPortNumber for the instance
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadSettingData) GetPropertyVxlanUDPPortNumber() (value uint16, err error) {
	retValue, err := instance.GetProperty("VxlanUDPPortNumber")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

//

// <param name="EncapsulationType" type="uint16 "></param>

// <param name="cmdletOutput" type="MSFT_NetAdapterEncapsulatedPacketTaskOffloadSettingData "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadSettingData) Enable( /* IN */ EncapsulationType uint16,
	/* OUT */ cmdletOutput MSFT_NetAdapterEncapsulatedPacketTaskOffloadSettingData) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Enable", EncapsulationType)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="EncapsulationType" type="uint16 "></param>

// <param name="cmdletOutput" type="MSFT_NetAdapterEncapsulatedPacketTaskOffloadSettingData "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapterEncapsulatedPacketTaskOffloadSettingData) Disable( /* IN */ EncapsulationType uint16,
	/* OUT */ cmdletOutput MSFT_NetAdapterEncapsulatedPacketTaskOffloadSettingData) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Disable", EncapsulationType)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
