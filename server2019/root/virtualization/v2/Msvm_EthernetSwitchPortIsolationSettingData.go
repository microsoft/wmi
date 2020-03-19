// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
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
	return instance.SetProperty("AllowUntaggedTraffic", value)
}

// GetAllowUntaggedTraffic gets the value of AllowUntaggedTraffic for the instance
func (instance *Msvm_EthernetSwitchPortIsolationSettingData) GetPropertyAllowUntaggedTraffic() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowUntaggedTraffic")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDefaultIsolationId sets the value of DefaultIsolationId for the instance
func (instance *Msvm_EthernetSwitchPortIsolationSettingData) SetPropertyDefaultIsolationId(value uint32) (err error) {
	return instance.SetProperty("DefaultIsolationId", value)
}

// GetDefaultIsolationId gets the value of DefaultIsolationId for the instance
func (instance *Msvm_EthernetSwitchPortIsolationSettingData) GetPropertyDefaultIsolationId() (value uint32, err error) {
	retValue, err := instance.GetProperty("DefaultIsolationId")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnableMultiTenantStack sets the value of EnableMultiTenantStack for the instance
func (instance *Msvm_EthernetSwitchPortIsolationSettingData) SetPropertyEnableMultiTenantStack(value bool) (err error) {
	return instance.SetProperty("EnableMultiTenantStack", value)
}

// GetEnableMultiTenantStack gets the value of EnableMultiTenantStack for the instance
func (instance *Msvm_EthernetSwitchPortIsolationSettingData) GetPropertyEnableMultiTenantStack() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableMultiTenantStack")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsolationMode sets the value of IsolationMode for the instance
func (instance *Msvm_EthernetSwitchPortIsolationSettingData) SetPropertyIsolationMode(value uint32) (err error) {
	return instance.SetProperty("IsolationMode", value)
}

// GetIsolationMode gets the value of IsolationMode for the instance
func (instance *Msvm_EthernetSwitchPortIsolationSettingData) GetPropertyIsolationMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("IsolationMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_EthernetSwitchPortIsolationSettingData) GetRelatedEthernetSwitchFeatureCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_EthernetSwitchFeatureCapabilities")
}
