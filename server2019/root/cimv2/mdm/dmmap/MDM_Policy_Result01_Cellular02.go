// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_Policy_Result01_Cellular02 struct
type MDM_Policy_Result01_Cellular02 struct {
	cim.WmiInstance

	//
	InstanceID string

	//
	LetAppsAccessCellularData int32

	//
	LetAppsAccessCellularData_ForceAllowTheseApps string

	//
	LetAppsAccessCellularData_ForceDenyTheseApps string

	//
	LetAppsAccessCellularData_UserInControlOfTheseApps string

	//
	ParentID string

	//
	ShowAppCellularAccessUI string
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_Cellular02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_Cellular02) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLetAppsAccessCellularData sets the value of LetAppsAccessCellularData for the instance
func (instance *MDM_Policy_Result01_Cellular02) SetPropertyLetAppsAccessCellularData(value int32) (err error) {
	return instance.SetProperty("LetAppsAccessCellularData", value)
}

// GetLetAppsAccessCellularData gets the value of LetAppsAccessCellularData for the instance
func (instance *MDM_Policy_Result01_Cellular02) GetPropertyLetAppsAccessCellularData() (value int32, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessCellularData")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLetAppsAccessCellularData_ForceAllowTheseApps sets the value of LetAppsAccessCellularData_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Cellular02) SetPropertyLetAppsAccessCellularData_ForceAllowTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessCellularData_ForceAllowTheseApps", value)
}

// GetLetAppsAccessCellularData_ForceAllowTheseApps gets the value of LetAppsAccessCellularData_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Cellular02) GetPropertyLetAppsAccessCellularData_ForceAllowTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessCellularData_ForceAllowTheseApps")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLetAppsAccessCellularData_ForceDenyTheseApps sets the value of LetAppsAccessCellularData_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Cellular02) SetPropertyLetAppsAccessCellularData_ForceDenyTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessCellularData_ForceDenyTheseApps", value)
}

// GetLetAppsAccessCellularData_ForceDenyTheseApps gets the value of LetAppsAccessCellularData_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Cellular02) GetPropertyLetAppsAccessCellularData_ForceDenyTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessCellularData_ForceDenyTheseApps")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLetAppsAccessCellularData_UserInControlOfTheseApps sets the value of LetAppsAccessCellularData_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Cellular02) SetPropertyLetAppsAccessCellularData_UserInControlOfTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessCellularData_UserInControlOfTheseApps", value)
}

// GetLetAppsAccessCellularData_UserInControlOfTheseApps gets the value of LetAppsAccessCellularData_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Cellular02) GetPropertyLetAppsAccessCellularData_UserInControlOfTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessCellularData_UserInControlOfTheseApps")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_Cellular02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_Cellular02) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetShowAppCellularAccessUI sets the value of ShowAppCellularAccessUI for the instance
func (instance *MDM_Policy_Result01_Cellular02) SetPropertyShowAppCellularAccessUI(value string) (err error) {
	return instance.SetProperty("ShowAppCellularAccessUI", value)
}

// GetShowAppCellularAccessUI gets the value of ShowAppCellularAccessUI for the instance
func (instance *MDM_Policy_Result01_Cellular02) GetPropertyShowAppCellularAccessUI() (value string, err error) {
	retValue, err := instance.GetProperty("ShowAppCellularAccessUI")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
