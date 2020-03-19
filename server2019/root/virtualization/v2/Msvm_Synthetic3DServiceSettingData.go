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

// Msvm_Synthetic3DServiceSettingData struct
type Msvm_Synthetic3DServiceSettingData struct {
	*CIM_SettingData

	//
	GPUOvercommitEnabled bool
}

func NewMsvm_Synthetic3DServiceSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_Synthetic3DServiceSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_Synthetic3DServiceSettingData{
		CIM_SettingData: tmp,
	}
	return
}

func NewMsvm_Synthetic3DServiceSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_Synthetic3DServiceSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_Synthetic3DServiceSettingData{
		CIM_SettingData: tmp,
	}
	return
}

// SetGPUOvercommitEnabled sets the value of GPUOvercommitEnabled for the instance
func (instance *Msvm_Synthetic3DServiceSettingData) SetPropertyGPUOvercommitEnabled(value bool) (err error) {
	return instance.SetProperty("GPUOvercommitEnabled", value)
}

// GetGPUOvercommitEnabled gets the value of GPUOvercommitEnabled for the instance
func (instance *Msvm_Synthetic3DServiceSettingData) GetPropertyGPUOvercommitEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("GPUOvercommitEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_Synthetic3DServiceSettingData) GetRelatedSynthetic3DService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_Synthetic3DService")
}
