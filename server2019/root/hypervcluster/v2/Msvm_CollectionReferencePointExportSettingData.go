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

// Msvm_CollectionReferencePointExportSettingData struct
type Msvm_CollectionReferencePointExportSettingData struct {
	*CIM_SettingData

	//
	BaseReferencePointCollection string

	//
	VirtualMachinesToDisksToExport []string
}

func NewMsvm_CollectionReferencePointExportSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_CollectionReferencePointExportSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_CollectionReferencePointExportSettingData{
		CIM_SettingData: tmp,
	}
	return
}

func NewMsvm_CollectionReferencePointExportSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_CollectionReferencePointExportSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_CollectionReferencePointExportSettingData{
		CIM_SettingData: tmp,
	}
	return
}

// SetBaseReferencePointCollection sets the value of BaseReferencePointCollection for the instance
func (instance *Msvm_CollectionReferencePointExportSettingData) SetPropertyBaseReferencePointCollection(value string) (err error) {
	return instance.SetProperty("BaseReferencePointCollection", value)
}

// GetBaseReferencePointCollection gets the value of BaseReferencePointCollection for the instance
func (instance *Msvm_CollectionReferencePointExportSettingData) GetPropertyBaseReferencePointCollection() (value string, err error) {
	retValue, err := instance.GetProperty("BaseReferencePointCollection")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualMachinesToDisksToExport sets the value of VirtualMachinesToDisksToExport for the instance
func (instance *Msvm_CollectionReferencePointExportSettingData) SetPropertyVirtualMachinesToDisksToExport(value []string) (err error) {
	return instance.SetProperty("VirtualMachinesToDisksToExport", value)
}

// GetVirtualMachinesToDisksToExport gets the value of VirtualMachinesToDisksToExport for the instance
func (instance *Msvm_CollectionReferencePointExportSettingData) GetPropertyVirtualMachinesToDisksToExport() (value []string, err error) {
	retValue, err := instance.GetProperty("VirtualMachinesToDisksToExport")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
