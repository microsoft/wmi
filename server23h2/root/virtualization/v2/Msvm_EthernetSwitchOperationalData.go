// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_EthernetSwitchOperationalData struct
type Msvm_EthernetSwitchOperationalData struct {
	*Msvm_EthernetSwitchData

	//
	CurrentSwitchingMode uint32

	//
	SupportedSwitchingModes []uint32
}

func NewMsvm_EthernetSwitchOperationalDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_EthernetSwitchOperationalData, err error) {
	tmp, err := NewMsvm_EthernetSwitchDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetSwitchOperationalData{
		Msvm_EthernetSwitchData: tmp,
	}
	return
}

func NewMsvm_EthernetSwitchOperationalDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_EthernetSwitchOperationalData, err error) {
	tmp, err := NewMsvm_EthernetSwitchDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetSwitchOperationalData{
		Msvm_EthernetSwitchData: tmp,
	}
	return
}

// SetCurrentSwitchingMode sets the value of CurrentSwitchingMode for the instance
func (instance *Msvm_EthernetSwitchOperationalData) SetPropertyCurrentSwitchingMode(value uint32) (err error) {
	return instance.SetProperty("CurrentSwitchingMode", (value))
}

// GetCurrentSwitchingMode gets the value of CurrentSwitchingMode for the instance
func (instance *Msvm_EthernetSwitchOperationalData) GetPropertyCurrentSwitchingMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentSwitchingMode")
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

// SetSupportedSwitchingModes sets the value of SupportedSwitchingModes for the instance
func (instance *Msvm_EthernetSwitchOperationalData) SetPropertySupportedSwitchingModes(value []uint32) (err error) {
	return instance.SetProperty("SupportedSwitchingModes", (value))
}

// GetSupportedSwitchingModes gets the value of SupportedSwitchingModes for the instance
func (instance *Msvm_EthernetSwitchOperationalData) GetPropertySupportedSwitchingModes() (value []uint32, err error) {
	retValue, err := instance.GetProperty("SupportedSwitchingModes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint32(valuetmp))
	}

	return
}
func (instance *Msvm_EthernetSwitchOperationalData) GetRelatedVirtualEthernetSwitch() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualEthernetSwitch")
}
