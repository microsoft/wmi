// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// Msvm_CollectionReferencePointExportSettingData struct
type Msvm_CollectionReferencePointExportSettingData struct {
	CIM_SettingData

	//
	BaseReferencePointCollection string

	//
	VirtualMachinesToDisksToExport []string
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
