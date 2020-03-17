// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

// Msvm_VirtualSystemReferencePointExportSettingData struct
type Msvm_VirtualSystemReferencePointExportSettingData struct {
	CIM_SettingData

	//
	BaseReferencePoint string

	//
	DisksToExport []string
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
