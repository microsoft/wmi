// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_Synthetic3DServiceSettingData struct
type Msvm_Synthetic3DServiceSettingData struct {
	CIM_SettingData

	//
	GPUOvercommitEnabled bool
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
