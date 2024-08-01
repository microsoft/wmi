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

// MSFT_NetAdapterSriovSettingData struct
type MSFT_NetAdapterSriovSettingData struct {
	*MSFT_NetAdapterSettingData

	//
	CurrentCapabilities MSFT_NetAdapterSriovCapabilities

	//
	Enabled bool

	//
	HardwareCapabilities MSFT_NetAdapterSriovCapabilities

	//
	NumActiveDefaultVPortMacAddresses uint32

	//
	NumActiveDefaultVPortVlanIds uint32

	//
	NumActiveNonDefaultVPortMacAddresses uint32

	//
	NumActiveNonDefaultVPortVlanIds uint32

	//
	NumActiveVPorts uint32

	//
	NumAllocatedVFs uint32

	//
	NumQueuePairsForDefaultVPort uint32

	//
	NumQueuePairsForNonDefaultVPorts uint32

	//
	NumVFs uint32

	//
	NumVPorts uint32

	//
	SriovSupport uint32

	//
	SwitchName string

	//
	SwitchType uint16
}

func NewMSFT_NetAdapterSriovSettingDataEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterSriovSettingData, err error) {
	tmp, err := NewMSFT_NetAdapterSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterSriovSettingData{
		MSFT_NetAdapterSettingData: tmp,
	}
	return
}

func NewMSFT_NetAdapterSriovSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapterSriovSettingData, err error) {
	tmp, err := NewMSFT_NetAdapterSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterSriovSettingData{
		MSFT_NetAdapterSettingData: tmp,
	}
	return
}

// SetCurrentCapabilities sets the value of CurrentCapabilities for the instance
func (instance *MSFT_NetAdapterSriovSettingData) SetPropertyCurrentCapabilities(value MSFT_NetAdapterSriovCapabilities) (err error) {
	return instance.SetProperty("CurrentCapabilities", (value))
}

// GetCurrentCapabilities gets the value of CurrentCapabilities for the instance
func (instance *MSFT_NetAdapterSriovSettingData) GetPropertyCurrentCapabilities() (value MSFT_NetAdapterSriovCapabilities, err error) {
	retValue, err := instance.GetProperty("CurrentCapabilities")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_NetAdapterSriovCapabilities)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_NetAdapterSriovCapabilities is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_NetAdapterSriovCapabilities(valuetmp)

	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *MSFT_NetAdapterSriovSettingData) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *MSFT_NetAdapterSriovSettingData) GetPropertyEnabled() (value bool, err error) {
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

// SetHardwareCapabilities sets the value of HardwareCapabilities for the instance
func (instance *MSFT_NetAdapterSriovSettingData) SetPropertyHardwareCapabilities(value MSFT_NetAdapterSriovCapabilities) (err error) {
	return instance.SetProperty("HardwareCapabilities", (value))
}

// GetHardwareCapabilities gets the value of HardwareCapabilities for the instance
func (instance *MSFT_NetAdapterSriovSettingData) GetPropertyHardwareCapabilities() (value MSFT_NetAdapterSriovCapabilities, err error) {
	retValue, err := instance.GetProperty("HardwareCapabilities")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_NetAdapterSriovCapabilities)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_NetAdapterSriovCapabilities is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_NetAdapterSriovCapabilities(valuetmp)

	return
}

// SetNumActiveDefaultVPortMacAddresses sets the value of NumActiveDefaultVPortMacAddresses for the instance
func (instance *MSFT_NetAdapterSriovSettingData) SetPropertyNumActiveDefaultVPortMacAddresses(value uint32) (err error) {
	return instance.SetProperty("NumActiveDefaultVPortMacAddresses", (value))
}

// GetNumActiveDefaultVPortMacAddresses gets the value of NumActiveDefaultVPortMacAddresses for the instance
func (instance *MSFT_NetAdapterSriovSettingData) GetPropertyNumActiveDefaultVPortMacAddresses() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumActiveDefaultVPortMacAddresses")
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

// SetNumActiveDefaultVPortVlanIds sets the value of NumActiveDefaultVPortVlanIds for the instance
func (instance *MSFT_NetAdapterSriovSettingData) SetPropertyNumActiveDefaultVPortVlanIds(value uint32) (err error) {
	return instance.SetProperty("NumActiveDefaultVPortVlanIds", (value))
}

// GetNumActiveDefaultVPortVlanIds gets the value of NumActiveDefaultVPortVlanIds for the instance
func (instance *MSFT_NetAdapterSriovSettingData) GetPropertyNumActiveDefaultVPortVlanIds() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumActiveDefaultVPortVlanIds")
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

// SetNumActiveNonDefaultVPortMacAddresses sets the value of NumActiveNonDefaultVPortMacAddresses for the instance
func (instance *MSFT_NetAdapterSriovSettingData) SetPropertyNumActiveNonDefaultVPortMacAddresses(value uint32) (err error) {
	return instance.SetProperty("NumActiveNonDefaultVPortMacAddresses", (value))
}

// GetNumActiveNonDefaultVPortMacAddresses gets the value of NumActiveNonDefaultVPortMacAddresses for the instance
func (instance *MSFT_NetAdapterSriovSettingData) GetPropertyNumActiveNonDefaultVPortMacAddresses() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumActiveNonDefaultVPortMacAddresses")
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

// SetNumActiveNonDefaultVPortVlanIds sets the value of NumActiveNonDefaultVPortVlanIds for the instance
func (instance *MSFT_NetAdapterSriovSettingData) SetPropertyNumActiveNonDefaultVPortVlanIds(value uint32) (err error) {
	return instance.SetProperty("NumActiveNonDefaultVPortVlanIds", (value))
}

// GetNumActiveNonDefaultVPortVlanIds gets the value of NumActiveNonDefaultVPortVlanIds for the instance
func (instance *MSFT_NetAdapterSriovSettingData) GetPropertyNumActiveNonDefaultVPortVlanIds() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumActiveNonDefaultVPortVlanIds")
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

// SetNumActiveVPorts sets the value of NumActiveVPorts for the instance
func (instance *MSFT_NetAdapterSriovSettingData) SetPropertyNumActiveVPorts(value uint32) (err error) {
	return instance.SetProperty("NumActiveVPorts", (value))
}

// GetNumActiveVPorts gets the value of NumActiveVPorts for the instance
func (instance *MSFT_NetAdapterSriovSettingData) GetPropertyNumActiveVPorts() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumActiveVPorts")
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

// SetNumAllocatedVFs sets the value of NumAllocatedVFs for the instance
func (instance *MSFT_NetAdapterSriovSettingData) SetPropertyNumAllocatedVFs(value uint32) (err error) {
	return instance.SetProperty("NumAllocatedVFs", (value))
}

// GetNumAllocatedVFs gets the value of NumAllocatedVFs for the instance
func (instance *MSFT_NetAdapterSriovSettingData) GetPropertyNumAllocatedVFs() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumAllocatedVFs")
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

// SetNumQueuePairsForDefaultVPort sets the value of NumQueuePairsForDefaultVPort for the instance
func (instance *MSFT_NetAdapterSriovSettingData) SetPropertyNumQueuePairsForDefaultVPort(value uint32) (err error) {
	return instance.SetProperty("NumQueuePairsForDefaultVPort", (value))
}

// GetNumQueuePairsForDefaultVPort gets the value of NumQueuePairsForDefaultVPort for the instance
func (instance *MSFT_NetAdapterSriovSettingData) GetPropertyNumQueuePairsForDefaultVPort() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumQueuePairsForDefaultVPort")
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

// SetNumQueuePairsForNonDefaultVPorts sets the value of NumQueuePairsForNonDefaultVPorts for the instance
func (instance *MSFT_NetAdapterSriovSettingData) SetPropertyNumQueuePairsForNonDefaultVPorts(value uint32) (err error) {
	return instance.SetProperty("NumQueuePairsForNonDefaultVPorts", (value))
}

// GetNumQueuePairsForNonDefaultVPorts gets the value of NumQueuePairsForNonDefaultVPorts for the instance
func (instance *MSFT_NetAdapterSriovSettingData) GetPropertyNumQueuePairsForNonDefaultVPorts() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumQueuePairsForNonDefaultVPorts")
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

// SetNumVFs sets the value of NumVFs for the instance
func (instance *MSFT_NetAdapterSriovSettingData) SetPropertyNumVFs(value uint32) (err error) {
	return instance.SetProperty("NumVFs", (value))
}

// GetNumVFs gets the value of NumVFs for the instance
func (instance *MSFT_NetAdapterSriovSettingData) GetPropertyNumVFs() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumVFs")
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

// SetNumVPorts sets the value of NumVPorts for the instance
func (instance *MSFT_NetAdapterSriovSettingData) SetPropertyNumVPorts(value uint32) (err error) {
	return instance.SetProperty("NumVPorts", (value))
}

// GetNumVPorts gets the value of NumVPorts for the instance
func (instance *MSFT_NetAdapterSriovSettingData) GetPropertyNumVPorts() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumVPorts")
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

// SetSriovSupport sets the value of SriovSupport for the instance
func (instance *MSFT_NetAdapterSriovSettingData) SetPropertySriovSupport(value uint32) (err error) {
	return instance.SetProperty("SriovSupport", (value))
}

// GetSriovSupport gets the value of SriovSupport for the instance
func (instance *MSFT_NetAdapterSriovSettingData) GetPropertySriovSupport() (value uint32, err error) {
	retValue, err := instance.GetProperty("SriovSupport")
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

// SetSwitchName sets the value of SwitchName for the instance
func (instance *MSFT_NetAdapterSriovSettingData) SetPropertySwitchName(value string) (err error) {
	return instance.SetProperty("SwitchName", (value))
}

// GetSwitchName gets the value of SwitchName for the instance
func (instance *MSFT_NetAdapterSriovSettingData) GetPropertySwitchName() (value string, err error) {
	retValue, err := instance.GetProperty("SwitchName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetSwitchType sets the value of SwitchType for the instance
func (instance *MSFT_NetAdapterSriovSettingData) SetPropertySwitchType(value uint16) (err error) {
	return instance.SetProperty("SwitchType", (value))
}

// GetSwitchType gets the value of SwitchType for the instance
func (instance *MSFT_NetAdapterSriovSettingData) GetPropertySwitchType() (value uint16, err error) {
	retValue, err := instance.GetProperty("SwitchType")
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

// <param name="cmdletOutput" type="MSFT_NetAdapterSriovSettingData "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapterSriovSettingData) Enable( /* OUT */ cmdletOutput MSFT_NetAdapterSriovSettingData) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Enable")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="cmdletOutput" type="MSFT_NetAdapterSriovSettingData "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapterSriovSettingData) Disable( /* OUT */ cmdletOutput MSFT_NetAdapterSriovSettingData) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Disable")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
