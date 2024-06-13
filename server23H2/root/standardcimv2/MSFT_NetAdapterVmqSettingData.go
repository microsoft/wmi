// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetAdapterVmqSettingData struct
type MSFT_NetAdapterVmqSettingData struct {
	*MSFT_NetAdapterSettingData

	//
	AnyVlanSupported bool

	//
	BaseProcessorGroup uint16

	//
	BaseProcessorNumber uint8

	//
	DynamicProcessorAffinityChangeSupported bool

	//
	Enabled bool

	//
	InterruptVectorCoalescingSupported bool

	//
	LookaheadSplitSupported bool

	//
	MaxLookaheadSplitSize uint32

	//
	MaxProcessorNumber uint8

	//
	MaxProcessors uint32

	//
	MinLookaheadSplitSize uint32

	//
	NumaNode uint16

	//
	NumberOfReceiveQueues uint32

	//
	NumMacAddressesPerPort uint32

	//
	NumVlansPerPort uint32

	//
	TotalNumberOfMacAddresses uint32

	//
	VlanFilteringSupported bool
}

func NewMSFT_NetAdapterVmqSettingDataEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterVmqSettingData, err error) {
	tmp, err := NewMSFT_NetAdapterSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterVmqSettingData{
		MSFT_NetAdapterSettingData: tmp,
	}
	return
}

func NewMSFT_NetAdapterVmqSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapterVmqSettingData, err error) {
	tmp, err := NewMSFT_NetAdapterSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterVmqSettingData{
		MSFT_NetAdapterSettingData: tmp,
	}
	return
}

// SetAnyVlanSupported sets the value of AnyVlanSupported for the instance
func (instance *MSFT_NetAdapterVmqSettingData) SetPropertyAnyVlanSupported(value bool) (err error) {
	return instance.SetProperty("AnyVlanSupported", (value))
}

// GetAnyVlanSupported gets the value of AnyVlanSupported for the instance
func (instance *MSFT_NetAdapterVmqSettingData) GetPropertyAnyVlanSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("AnyVlanSupported")
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

// SetBaseProcessorGroup sets the value of BaseProcessorGroup for the instance
func (instance *MSFT_NetAdapterVmqSettingData) SetPropertyBaseProcessorGroup(value uint16) (err error) {
	return instance.SetProperty("BaseProcessorGroup", (value))
}

// GetBaseProcessorGroup gets the value of BaseProcessorGroup for the instance
func (instance *MSFT_NetAdapterVmqSettingData) GetPropertyBaseProcessorGroup() (value uint16, err error) {
	retValue, err := instance.GetProperty("BaseProcessorGroup")
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

// SetBaseProcessorNumber sets the value of BaseProcessorNumber for the instance
func (instance *MSFT_NetAdapterVmqSettingData) SetPropertyBaseProcessorNumber(value uint8) (err error) {
	return instance.SetProperty("BaseProcessorNumber", (value))
}

// GetBaseProcessorNumber gets the value of BaseProcessorNumber for the instance
func (instance *MSFT_NetAdapterVmqSettingData) GetPropertyBaseProcessorNumber() (value uint8, err error) {
	retValue, err := instance.GetProperty("BaseProcessorNumber")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetDynamicProcessorAffinityChangeSupported sets the value of DynamicProcessorAffinityChangeSupported for the instance
func (instance *MSFT_NetAdapterVmqSettingData) SetPropertyDynamicProcessorAffinityChangeSupported(value bool) (err error) {
	return instance.SetProperty("DynamicProcessorAffinityChangeSupported", (value))
}

// GetDynamicProcessorAffinityChangeSupported gets the value of DynamicProcessorAffinityChangeSupported for the instance
func (instance *MSFT_NetAdapterVmqSettingData) GetPropertyDynamicProcessorAffinityChangeSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("DynamicProcessorAffinityChangeSupported")
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

// SetEnabled sets the value of Enabled for the instance
func (instance *MSFT_NetAdapterVmqSettingData) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *MSFT_NetAdapterVmqSettingData) GetPropertyEnabled() (value bool, err error) {
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

// SetInterruptVectorCoalescingSupported sets the value of InterruptVectorCoalescingSupported for the instance
func (instance *MSFT_NetAdapterVmqSettingData) SetPropertyInterruptVectorCoalescingSupported(value bool) (err error) {
	return instance.SetProperty("InterruptVectorCoalescingSupported", (value))
}

// GetInterruptVectorCoalescingSupported gets the value of InterruptVectorCoalescingSupported for the instance
func (instance *MSFT_NetAdapterVmqSettingData) GetPropertyInterruptVectorCoalescingSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("InterruptVectorCoalescingSupported")
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

// SetLookaheadSplitSupported sets the value of LookaheadSplitSupported for the instance
func (instance *MSFT_NetAdapterVmqSettingData) SetPropertyLookaheadSplitSupported(value bool) (err error) {
	return instance.SetProperty("LookaheadSplitSupported", (value))
}

// GetLookaheadSplitSupported gets the value of LookaheadSplitSupported for the instance
func (instance *MSFT_NetAdapterVmqSettingData) GetPropertyLookaheadSplitSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("LookaheadSplitSupported")
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

// SetMaxLookaheadSplitSize sets the value of MaxLookaheadSplitSize for the instance
func (instance *MSFT_NetAdapterVmqSettingData) SetPropertyMaxLookaheadSplitSize(value uint32) (err error) {
	return instance.SetProperty("MaxLookaheadSplitSize", (value))
}

// GetMaxLookaheadSplitSize gets the value of MaxLookaheadSplitSize for the instance
func (instance *MSFT_NetAdapterVmqSettingData) GetPropertyMaxLookaheadSplitSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxLookaheadSplitSize")
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

// SetMaxProcessorNumber sets the value of MaxProcessorNumber for the instance
func (instance *MSFT_NetAdapterVmqSettingData) SetPropertyMaxProcessorNumber(value uint8) (err error) {
	return instance.SetProperty("MaxProcessorNumber", (value))
}

// GetMaxProcessorNumber gets the value of MaxProcessorNumber for the instance
func (instance *MSFT_NetAdapterVmqSettingData) GetPropertyMaxProcessorNumber() (value uint8, err error) {
	retValue, err := instance.GetProperty("MaxProcessorNumber")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetMaxProcessors sets the value of MaxProcessors for the instance
func (instance *MSFT_NetAdapterVmqSettingData) SetPropertyMaxProcessors(value uint32) (err error) {
	return instance.SetProperty("MaxProcessors", (value))
}

// GetMaxProcessors gets the value of MaxProcessors for the instance
func (instance *MSFT_NetAdapterVmqSettingData) GetPropertyMaxProcessors() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxProcessors")
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

// SetMinLookaheadSplitSize sets the value of MinLookaheadSplitSize for the instance
func (instance *MSFT_NetAdapterVmqSettingData) SetPropertyMinLookaheadSplitSize(value uint32) (err error) {
	return instance.SetProperty("MinLookaheadSplitSize", (value))
}

// GetMinLookaheadSplitSize gets the value of MinLookaheadSplitSize for the instance
func (instance *MSFT_NetAdapterVmqSettingData) GetPropertyMinLookaheadSplitSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinLookaheadSplitSize")
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

// SetNumaNode sets the value of NumaNode for the instance
func (instance *MSFT_NetAdapterVmqSettingData) SetPropertyNumaNode(value uint16) (err error) {
	return instance.SetProperty("NumaNode", (value))
}

// GetNumaNode gets the value of NumaNode for the instance
func (instance *MSFT_NetAdapterVmqSettingData) GetPropertyNumaNode() (value uint16, err error) {
	retValue, err := instance.GetProperty("NumaNode")
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

// SetNumberOfReceiveQueues sets the value of NumberOfReceiveQueues for the instance
func (instance *MSFT_NetAdapterVmqSettingData) SetPropertyNumberOfReceiveQueues(value uint32) (err error) {
	return instance.SetProperty("NumberOfReceiveQueues", (value))
}

// GetNumberOfReceiveQueues gets the value of NumberOfReceiveQueues for the instance
func (instance *MSFT_NetAdapterVmqSettingData) GetPropertyNumberOfReceiveQueues() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfReceiveQueues")
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

// SetNumMacAddressesPerPort sets the value of NumMacAddressesPerPort for the instance
func (instance *MSFT_NetAdapterVmqSettingData) SetPropertyNumMacAddressesPerPort(value uint32) (err error) {
	return instance.SetProperty("NumMacAddressesPerPort", (value))
}

// GetNumMacAddressesPerPort gets the value of NumMacAddressesPerPort for the instance
func (instance *MSFT_NetAdapterVmqSettingData) GetPropertyNumMacAddressesPerPort() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumMacAddressesPerPort")
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

// SetNumVlansPerPort sets the value of NumVlansPerPort for the instance
func (instance *MSFT_NetAdapterVmqSettingData) SetPropertyNumVlansPerPort(value uint32) (err error) {
	return instance.SetProperty("NumVlansPerPort", (value))
}

// GetNumVlansPerPort gets the value of NumVlansPerPort for the instance
func (instance *MSFT_NetAdapterVmqSettingData) GetPropertyNumVlansPerPort() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumVlansPerPort")
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

// SetTotalNumberOfMacAddresses sets the value of TotalNumberOfMacAddresses for the instance
func (instance *MSFT_NetAdapterVmqSettingData) SetPropertyTotalNumberOfMacAddresses(value uint32) (err error) {
	return instance.SetProperty("TotalNumberOfMacAddresses", (value))
}

// GetTotalNumberOfMacAddresses gets the value of TotalNumberOfMacAddresses for the instance
func (instance *MSFT_NetAdapterVmqSettingData) GetPropertyTotalNumberOfMacAddresses() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalNumberOfMacAddresses")
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

// SetVlanFilteringSupported sets the value of VlanFilteringSupported for the instance
func (instance *MSFT_NetAdapterVmqSettingData) SetPropertyVlanFilteringSupported(value bool) (err error) {
	return instance.SetProperty("VlanFilteringSupported", (value))
}

// GetVlanFilteringSupported gets the value of VlanFilteringSupported for the instance
func (instance *MSFT_NetAdapterVmqSettingData) GetPropertyVlanFilteringSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("VlanFilteringSupported")
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

// <param name="cmdletOutput" type="MSFT_NetAdapterVmqSettingData "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapterVmqSettingData) Enable( /* OUT */ cmdletOutput MSFT_NetAdapterVmqSettingData) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Enable")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="cmdletOutput" type="MSFT_NetAdapterVmqSettingData "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapterVmqSettingData) Disable( /* OUT */ cmdletOutput MSFT_NetAdapterVmqSettingData) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Disable")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
