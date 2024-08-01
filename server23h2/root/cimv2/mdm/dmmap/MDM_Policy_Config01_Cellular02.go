// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_Policy_Config01_Cellular02 struct
type MDM_Policy_Config01_Cellular02 struct {
	*cim.WmiInstance

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

func NewMDM_Policy_Config01_Cellular02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Config01_Cellular02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_Cellular02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Config01_Cellular02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Config01_Cellular02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_Cellular02{
		WmiInstance: tmp,
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_Cellular02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_Cellular02) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetLetAppsAccessCellularData sets the value of LetAppsAccessCellularData for the instance
func (instance *MDM_Policy_Config01_Cellular02) SetPropertyLetAppsAccessCellularData(value int32) (err error) {
	return instance.SetProperty("LetAppsAccessCellularData", (value))
}

// GetLetAppsAccessCellularData gets the value of LetAppsAccessCellularData for the instance
func (instance *MDM_Policy_Config01_Cellular02) GetPropertyLetAppsAccessCellularData() (value int32, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessCellularData")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetLetAppsAccessCellularData_ForceAllowTheseApps sets the value of LetAppsAccessCellularData_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Config01_Cellular02) SetPropertyLetAppsAccessCellularData_ForceAllowTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessCellularData_ForceAllowTheseApps", (value))
}

// GetLetAppsAccessCellularData_ForceAllowTheseApps gets the value of LetAppsAccessCellularData_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Config01_Cellular02) GetPropertyLetAppsAccessCellularData_ForceAllowTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessCellularData_ForceAllowTheseApps")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetLetAppsAccessCellularData_ForceDenyTheseApps sets the value of LetAppsAccessCellularData_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Config01_Cellular02) SetPropertyLetAppsAccessCellularData_ForceDenyTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessCellularData_ForceDenyTheseApps", (value))
}

// GetLetAppsAccessCellularData_ForceDenyTheseApps gets the value of LetAppsAccessCellularData_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Config01_Cellular02) GetPropertyLetAppsAccessCellularData_ForceDenyTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessCellularData_ForceDenyTheseApps")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetLetAppsAccessCellularData_UserInControlOfTheseApps sets the value of LetAppsAccessCellularData_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Config01_Cellular02) SetPropertyLetAppsAccessCellularData_UserInControlOfTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessCellularData_UserInControlOfTheseApps", (value))
}

// GetLetAppsAccessCellularData_UserInControlOfTheseApps gets the value of LetAppsAccessCellularData_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Config01_Cellular02) GetPropertyLetAppsAccessCellularData_UserInControlOfTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessCellularData_UserInControlOfTheseApps")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_Cellular02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_Cellular02) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetShowAppCellularAccessUI sets the value of ShowAppCellularAccessUI for the instance
func (instance *MDM_Policy_Config01_Cellular02) SetPropertyShowAppCellularAccessUI(value string) (err error) {
	return instance.SetProperty("ShowAppCellularAccessUI", (value))
}

// GetShowAppCellularAccessUI gets the value of ShowAppCellularAccessUI for the instance
func (instance *MDM_Policy_Config01_Cellular02) GetPropertyShowAppCellularAccessUI() (value string, err error) {
	retValue, err := instance.GetProperty("ShowAppCellularAccessUI")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}
