// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_VirtualSystemReferencePointSettingData struct
type Msvm_VirtualSystemReferencePointSettingData struct {
	*CIM_SettingData

	//
	ConsistencyLevel uint8
}

func NewMsvm_VirtualSystemReferencePointSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_VirtualSystemReferencePointSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualSystemReferencePointSettingData{
		CIM_SettingData: tmp,
	}
	return
}

func NewMsvm_VirtualSystemReferencePointSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VirtualSystemReferencePointSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualSystemReferencePointSettingData{
		CIM_SettingData: tmp,
	}
	return
}

// SetConsistencyLevel sets the value of ConsistencyLevel for the instance
func (instance *Msvm_VirtualSystemReferencePointSettingData) SetPropertyConsistencyLevel(value uint8) (err error) {
	return instance.SetProperty("ConsistencyLevel", value)
}

// GetConsistencyLevel gets the value of ConsistencyLevel for the instance
func (instance *Msvm_VirtualSystemReferencePointSettingData) GetPropertyConsistencyLevel() (value uint8, err error) {
	retValue, err := instance.GetProperty("ConsistencyLevel")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}
