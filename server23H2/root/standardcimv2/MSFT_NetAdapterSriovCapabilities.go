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

// MSFT_NetAdapterSriovCapabilities struct
type MSFT_NetAdapterSriovCapabilities struct {
	*cim.WmiInstance

	//
	AsymmetricQueuePairsForNonDefaultVPortsSupported bool

	//
	MaxNumMacAddresses uint32

	//
	MaxNumQueuePairs uint32

	//
	MaxNumQueuePairsPerNonDefaultVPort uint32

	//
	MaxNumSwitches uint32

	//
	MaxNumVFs uint32

	//
	MaxNumVPorts uint32

	//
	PerVportInterruptModerationSupported bool

	//
	SingleVportPoolSupported bool

	//
	VfRssSupported bool

	//
	VlanSupported bool
}

func NewMSFT_NetAdapterSriovCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterSriovCapabilities, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterSriovCapabilities{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_NetAdapterSriovCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapterSriovCapabilities, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterSriovCapabilities{
		WmiInstance: tmp,
	}
	return
}

// SetAsymmetricQueuePairsForNonDefaultVPortsSupported sets the value of AsymmetricQueuePairsForNonDefaultVPortsSupported for the instance
func (instance *MSFT_NetAdapterSriovCapabilities) SetPropertyAsymmetricQueuePairsForNonDefaultVPortsSupported(value bool) (err error) {
	return instance.SetProperty("AsymmetricQueuePairsForNonDefaultVPortsSupported", (value))
}

// GetAsymmetricQueuePairsForNonDefaultVPortsSupported gets the value of AsymmetricQueuePairsForNonDefaultVPortsSupported for the instance
func (instance *MSFT_NetAdapterSriovCapabilities) GetPropertyAsymmetricQueuePairsForNonDefaultVPortsSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("AsymmetricQueuePairsForNonDefaultVPortsSupported")
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

// SetMaxNumMacAddresses sets the value of MaxNumMacAddresses for the instance
func (instance *MSFT_NetAdapterSriovCapabilities) SetPropertyMaxNumMacAddresses(value uint32) (err error) {
	return instance.SetProperty("MaxNumMacAddresses", (value))
}

// GetMaxNumMacAddresses gets the value of MaxNumMacAddresses for the instance
func (instance *MSFT_NetAdapterSriovCapabilities) GetPropertyMaxNumMacAddresses() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxNumMacAddresses")
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

// SetMaxNumQueuePairs sets the value of MaxNumQueuePairs for the instance
func (instance *MSFT_NetAdapterSriovCapabilities) SetPropertyMaxNumQueuePairs(value uint32) (err error) {
	return instance.SetProperty("MaxNumQueuePairs", (value))
}

// GetMaxNumQueuePairs gets the value of MaxNumQueuePairs for the instance
func (instance *MSFT_NetAdapterSriovCapabilities) GetPropertyMaxNumQueuePairs() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxNumQueuePairs")
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

// SetMaxNumQueuePairsPerNonDefaultVPort sets the value of MaxNumQueuePairsPerNonDefaultVPort for the instance
func (instance *MSFT_NetAdapterSriovCapabilities) SetPropertyMaxNumQueuePairsPerNonDefaultVPort(value uint32) (err error) {
	return instance.SetProperty("MaxNumQueuePairsPerNonDefaultVPort", (value))
}

// GetMaxNumQueuePairsPerNonDefaultVPort gets the value of MaxNumQueuePairsPerNonDefaultVPort for the instance
func (instance *MSFT_NetAdapterSriovCapabilities) GetPropertyMaxNumQueuePairsPerNonDefaultVPort() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxNumQueuePairsPerNonDefaultVPort")
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

// SetMaxNumSwitches sets the value of MaxNumSwitches for the instance
func (instance *MSFT_NetAdapterSriovCapabilities) SetPropertyMaxNumSwitches(value uint32) (err error) {
	return instance.SetProperty("MaxNumSwitches", (value))
}

// GetMaxNumSwitches gets the value of MaxNumSwitches for the instance
func (instance *MSFT_NetAdapterSriovCapabilities) GetPropertyMaxNumSwitches() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxNumSwitches")
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

// SetMaxNumVFs sets the value of MaxNumVFs for the instance
func (instance *MSFT_NetAdapterSriovCapabilities) SetPropertyMaxNumVFs(value uint32) (err error) {
	return instance.SetProperty("MaxNumVFs", (value))
}

// GetMaxNumVFs gets the value of MaxNumVFs for the instance
func (instance *MSFT_NetAdapterSriovCapabilities) GetPropertyMaxNumVFs() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxNumVFs")
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

// SetMaxNumVPorts sets the value of MaxNumVPorts for the instance
func (instance *MSFT_NetAdapterSriovCapabilities) SetPropertyMaxNumVPorts(value uint32) (err error) {
	return instance.SetProperty("MaxNumVPorts", (value))
}

// GetMaxNumVPorts gets the value of MaxNumVPorts for the instance
func (instance *MSFT_NetAdapterSriovCapabilities) GetPropertyMaxNumVPorts() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxNumVPorts")
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

// SetPerVportInterruptModerationSupported sets the value of PerVportInterruptModerationSupported for the instance
func (instance *MSFT_NetAdapterSriovCapabilities) SetPropertyPerVportInterruptModerationSupported(value bool) (err error) {
	return instance.SetProperty("PerVportInterruptModerationSupported", (value))
}

// GetPerVportInterruptModerationSupported gets the value of PerVportInterruptModerationSupported for the instance
func (instance *MSFT_NetAdapterSriovCapabilities) GetPropertyPerVportInterruptModerationSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("PerVportInterruptModerationSupported")
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

// SetSingleVportPoolSupported sets the value of SingleVportPoolSupported for the instance
func (instance *MSFT_NetAdapterSriovCapabilities) SetPropertySingleVportPoolSupported(value bool) (err error) {
	return instance.SetProperty("SingleVportPoolSupported", (value))
}

// GetSingleVportPoolSupported gets the value of SingleVportPoolSupported for the instance
func (instance *MSFT_NetAdapterSriovCapabilities) GetPropertySingleVportPoolSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("SingleVportPoolSupported")
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

// SetVfRssSupported sets the value of VfRssSupported for the instance
func (instance *MSFT_NetAdapterSriovCapabilities) SetPropertyVfRssSupported(value bool) (err error) {
	return instance.SetProperty("VfRssSupported", (value))
}

// GetVfRssSupported gets the value of VfRssSupported for the instance
func (instance *MSFT_NetAdapterSriovCapabilities) GetPropertyVfRssSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("VfRssSupported")
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

// SetVlanSupported sets the value of VlanSupported for the instance
func (instance *MSFT_NetAdapterSriovCapabilities) SetPropertyVlanSupported(value bool) (err error) {
	return instance.SetProperty("VlanSupported", (value))
}

// GetVlanSupported gets the value of VlanSupported for the instance
func (instance *MSFT_NetAdapterSriovCapabilities) GetPropertyVlanSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("VlanSupported")
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
