// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_EthernetSwitchPortIsolationSettingData struct
type Msvm_EthernetSwitchPortIsolationSettingData struct {
	*Msvm_EthernetSwitchPortFeatureSettingData

	//
	AllowUntaggedTraffic bool

	//
	DefaultIsolationId uint32

	//
	EnableMultiTenantStack bool

	//
	IsolationMode uint32
}

func NewMsvm_EthernetSwitchPortIsolationSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_EthernetSwitchPortIsolationSettingData, err error) {
	tmp, err := NewMsvm_EthernetSwitchPortFeatureSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetSwitchPortIsolationSettingData{
		Msvm_EthernetSwitchPortFeatureSettingData: tmp,
	}
	return
}

func NewMsvm_EthernetSwitchPortIsolationSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_EthernetSwitchPortIsolationSettingData, err error) {
	tmp, err := NewMsvm_EthernetSwitchPortFeatureSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetSwitchPortIsolationSettingData{
		Msvm_EthernetSwitchPortFeatureSettingData: tmp,
	}
	return
}

// SetAllowUntaggedTraffic sets the value of AllowUntaggedTraffic for the instance
func (instance *Msvm_EthernetSwitchPortIsolationSettingData) SetPropertyAllowUntaggedTraffic(value bool) (err error) {
	return instance.SetProperty("AllowUntaggedTraffic", (value))
}

// GetAllowUntaggedTraffic gets the value of AllowUntaggedTraffic for the instance
func (instance *Msvm_EthernetSwitchPortIsolationSettingData) GetPropertyAllowUntaggedTraffic() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowUntaggedTraffic")
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

// SetDefaultIsolationId sets the value of DefaultIsolationId for the instance
func (instance *Msvm_EthernetSwitchPortIsolationSettingData) SetPropertyDefaultIsolationId(value uint32) (err error) {
	return instance.SetProperty("DefaultIsolationId", (value))
}

// GetDefaultIsolationId gets the value of DefaultIsolationId for the instance
func (instance *Msvm_EthernetSwitchPortIsolationSettingData) GetPropertyDefaultIsolationId() (value uint32, err error) {
	retValue, err := instance.GetProperty("DefaultIsolationId")
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

// SetEnableMultiTenantStack sets the value of EnableMultiTenantStack for the instance
func (instance *Msvm_EthernetSwitchPortIsolationSettingData) SetPropertyEnableMultiTenantStack(value bool) (err error) {
	return instance.SetProperty("EnableMultiTenantStack", (value))
}

// GetEnableMultiTenantStack gets the value of EnableMultiTenantStack for the instance
func (instance *Msvm_EthernetSwitchPortIsolationSettingData) GetPropertyEnableMultiTenantStack() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableMultiTenantStack")
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

// SetIsolationMode sets the value of IsolationMode for the instance
func (instance *Msvm_EthernetSwitchPortIsolationSettingData) SetPropertyIsolationMode(value uint32) (err error) {
	return instance.SetProperty("IsolationMode", (value))
}

// GetIsolationMode gets the value of IsolationMode for the instance
func (instance *Msvm_EthernetSwitchPortIsolationSettingData) GetPropertyIsolationMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("IsolationMode")
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
func (instance *Msvm_EthernetSwitchPortIsolationSettingData) GetRelatedEthernetPortAllocationSettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_EthernetPortAllocationSettingData")
}
