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

// Msvm_VirtualSystemReferencePointExportSettingData struct
type Msvm_VirtualSystemReferencePointExportSettingData struct {
	*CIM_SettingData

	//
	BaseReferencePoint string

	//
	DisksToExport []string
}

func NewMsvm_VirtualSystemReferencePointExportSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_VirtualSystemReferencePointExportSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualSystemReferencePointExportSettingData{
		CIM_SettingData: tmp,
	}
	return
}

func NewMsvm_VirtualSystemReferencePointExportSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VirtualSystemReferencePointExportSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualSystemReferencePointExportSettingData{
		CIM_SettingData: tmp,
	}
	return
}

// SetBaseReferencePoint sets the value of BaseReferencePoint for the instance
func (instance *Msvm_VirtualSystemReferencePointExportSettingData) SetPropertyBaseReferencePoint(value string) (err error) {
	return instance.SetProperty("BaseReferencePoint", value)
}

// GetBaseReferencePoint gets the value of BaseReferencePoint for the instance
func (instance *Msvm_VirtualSystemReferencePointExportSettingData) GetPropertyBaseReferencePoint() (value string, err error) {
	retValue, err := instance.GetProperty("BaseReferencePoint")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisksToExport sets the value of DisksToExport for the instance
func (instance *Msvm_VirtualSystemReferencePointExportSettingData) SetPropertyDisksToExport(value []string) (err error) {
	return instance.SetProperty("DisksToExport", value)
}

// GetDisksToExport gets the value of DisksToExport for the instance
func (instance *Msvm_VirtualSystemReferencePointExportSettingData) GetPropertyDisksToExport() (value []string, err error) {
	retValue, err := instance.GetProperty("DisksToExport")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
