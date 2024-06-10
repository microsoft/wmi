// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSNdis_PMCapabilitiesParam struct
type MSNdis_PMCapabilitiesParam struct {
	*MSNdis

	//
	DeviceSleepOnDisconnect MSNdis_PMCapabilityState

	//
	Header MSNdis_ObjectHeader

	//
	PMARPOffload MSNdis_PMCapabilityState

	//
	PMNDOffload MSNdis_PMCapabilityState

	//
	PMWiFiRekeyOffload MSNdis_PMCapabilityState

	//
	WakeOnMagicPacket MSNdis_PMCapabilityState

	//
	WakeOnPattern MSNdis_PMCapabilityState
}

func NewMSNdis_PMCapabilitiesParamEx1(instance *cim.WmiInstance) (newInstance *MSNdis_PMCapabilitiesParam, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_PMCapabilitiesParam{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_PMCapabilitiesParamEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_PMCapabilitiesParam, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_PMCapabilitiesParam{
		MSNdis: tmp,
	}
	return
}

// SetDeviceSleepOnDisconnect sets the value of DeviceSleepOnDisconnect for the instance
func (instance *MSNdis_PMCapabilitiesParam) SetPropertyDeviceSleepOnDisconnect(value MSNdis_PMCapabilityState) (err error) {
	return instance.SetProperty("DeviceSleepOnDisconnect", (value))
}

// GetDeviceSleepOnDisconnect gets the value of DeviceSleepOnDisconnect for the instance
func (instance *MSNdis_PMCapabilitiesParam) GetPropertyDeviceSleepOnDisconnect() (value MSNdis_PMCapabilityState, err error) {
	retValue, err := instance.GetProperty("DeviceSleepOnDisconnect")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_PMCapabilityState)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_PMCapabilityState is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_PMCapabilityState(valuetmp)

	return
}

// SetHeader sets the value of Header for the instance
func (instance *MSNdis_PMCapabilitiesParam) SetPropertyHeader(value MSNdis_ObjectHeader) (err error) {
	return instance.SetProperty("Header", (value))
}

// GetHeader gets the value of Header for the instance
func (instance *MSNdis_PMCapabilitiesParam) GetPropertyHeader() (value MSNdis_ObjectHeader, err error) {
	retValue, err := instance.GetProperty("Header")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_ObjectHeader)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_ObjectHeader is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_ObjectHeader(valuetmp)

	return
}

// SetPMARPOffload sets the value of PMARPOffload for the instance
func (instance *MSNdis_PMCapabilitiesParam) SetPropertyPMARPOffload(value MSNdis_PMCapabilityState) (err error) {
	return instance.SetProperty("PMARPOffload", (value))
}

// GetPMARPOffload gets the value of PMARPOffload for the instance
func (instance *MSNdis_PMCapabilitiesParam) GetPropertyPMARPOffload() (value MSNdis_PMCapabilityState, err error) {
	retValue, err := instance.GetProperty("PMARPOffload")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_PMCapabilityState)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_PMCapabilityState is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_PMCapabilityState(valuetmp)

	return
}

// SetPMNDOffload sets the value of PMNDOffload for the instance
func (instance *MSNdis_PMCapabilitiesParam) SetPropertyPMNDOffload(value MSNdis_PMCapabilityState) (err error) {
	return instance.SetProperty("PMNDOffload", (value))
}

// GetPMNDOffload gets the value of PMNDOffload for the instance
func (instance *MSNdis_PMCapabilitiesParam) GetPropertyPMNDOffload() (value MSNdis_PMCapabilityState, err error) {
	retValue, err := instance.GetProperty("PMNDOffload")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_PMCapabilityState)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_PMCapabilityState is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_PMCapabilityState(valuetmp)

	return
}

// SetPMWiFiRekeyOffload sets the value of PMWiFiRekeyOffload for the instance
func (instance *MSNdis_PMCapabilitiesParam) SetPropertyPMWiFiRekeyOffload(value MSNdis_PMCapabilityState) (err error) {
	return instance.SetProperty("PMWiFiRekeyOffload", (value))
}

// GetPMWiFiRekeyOffload gets the value of PMWiFiRekeyOffload for the instance
func (instance *MSNdis_PMCapabilitiesParam) GetPropertyPMWiFiRekeyOffload() (value MSNdis_PMCapabilityState, err error) {
	retValue, err := instance.GetProperty("PMWiFiRekeyOffload")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_PMCapabilityState)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_PMCapabilityState is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_PMCapabilityState(valuetmp)

	return
}

// SetWakeOnMagicPacket sets the value of WakeOnMagicPacket for the instance
func (instance *MSNdis_PMCapabilitiesParam) SetPropertyWakeOnMagicPacket(value MSNdis_PMCapabilityState) (err error) {
	return instance.SetProperty("WakeOnMagicPacket", (value))
}

// GetWakeOnMagicPacket gets the value of WakeOnMagicPacket for the instance
func (instance *MSNdis_PMCapabilitiesParam) GetPropertyWakeOnMagicPacket() (value MSNdis_PMCapabilityState, err error) {
	retValue, err := instance.GetProperty("WakeOnMagicPacket")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_PMCapabilityState)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_PMCapabilityState is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_PMCapabilityState(valuetmp)

	return
}

// SetWakeOnPattern sets the value of WakeOnPattern for the instance
func (instance *MSNdis_PMCapabilitiesParam) SetPropertyWakeOnPattern(value MSNdis_PMCapabilityState) (err error) {
	return instance.SetProperty("WakeOnPattern", (value))
}

// GetWakeOnPattern gets the value of WakeOnPattern for the instance
func (instance *MSNdis_PMCapabilitiesParam) GetPropertyWakeOnPattern() (value MSNdis_PMCapabilityState, err error) {
	retValue, err := instance.GetProperty("WakeOnPattern")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_PMCapabilityState)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_PMCapabilityState is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_PMCapabilityState(valuetmp)

	return
}
