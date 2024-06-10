// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.CIMV2.mdm.dmmap
//
// ////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_Policy_User_Result01_FileExplorer02 struct
type MDM_Policy_User_Result01_FileExplorer02 struct {
	*cim.WmiInstance

	//
	AllowOptionToShowNetwork int32

	//
	AllowOptionToShowThisPC int32

	//
	InstanceID string

	//
	ParentID string

	//
	SetAllowedFolderLocations int32

	//
	SetAllowedStorageLocations int32
}

func NewMDM_Policy_User_Result01_FileExplorer02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_User_Result01_FileExplorer02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_User_Result01_FileExplorer02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_User_Result01_FileExplorer02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_User_Result01_FileExplorer02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_User_Result01_FileExplorer02{
		WmiInstance: tmp,
	}
	return
}

// SetAllowOptionToShowNetwork sets the value of AllowOptionToShowNetwork for the instance
func (instance *MDM_Policy_User_Result01_FileExplorer02) SetPropertyAllowOptionToShowNetwork(value int32) (err error) {
	return instance.SetProperty("AllowOptionToShowNetwork", (value))
}

// GetAllowOptionToShowNetwork gets the value of AllowOptionToShowNetwork for the instance
func (instance *MDM_Policy_User_Result01_FileExplorer02) GetPropertyAllowOptionToShowNetwork() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowOptionToShowNetwork")
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

// SetAllowOptionToShowThisPC sets the value of AllowOptionToShowThisPC for the instance
func (instance *MDM_Policy_User_Result01_FileExplorer02) SetPropertyAllowOptionToShowThisPC(value int32) (err error) {
	return instance.SetProperty("AllowOptionToShowThisPC", (value))
}

// GetAllowOptionToShowThisPC gets the value of AllowOptionToShowThisPC for the instance
func (instance *MDM_Policy_User_Result01_FileExplorer02) GetPropertyAllowOptionToShowThisPC() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowOptionToShowThisPC")
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

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_User_Result01_FileExplorer02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_User_Result01_FileExplorer02) GetPropertyInstanceID() (value string, err error) {
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Policy_User_Result01_FileExplorer02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_User_Result01_FileExplorer02) GetPropertyParentID() (value string, err error) {
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

// SetSetAllowedFolderLocations sets the value of SetAllowedFolderLocations for the instance
func (instance *MDM_Policy_User_Result01_FileExplorer02) SetPropertySetAllowedFolderLocations(value int32) (err error) {
	return instance.SetProperty("SetAllowedFolderLocations", (value))
}

// GetSetAllowedFolderLocations gets the value of SetAllowedFolderLocations for the instance
func (instance *MDM_Policy_User_Result01_FileExplorer02) GetPropertySetAllowedFolderLocations() (value int32, err error) {
	retValue, err := instance.GetProperty("SetAllowedFolderLocations")
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

// SetSetAllowedStorageLocations sets the value of SetAllowedStorageLocations for the instance
func (instance *MDM_Policy_User_Result01_FileExplorer02) SetPropertySetAllowedStorageLocations(value int32) (err error) {
	return instance.SetProperty("SetAllowedStorageLocations", (value))
}

// GetSetAllowedStorageLocations gets the value of SetAllowedStorageLocations for the instance
func (instance *MDM_Policy_User_Result01_FileExplorer02) GetPropertySetAllowedStorageLocations() (value int32, err error) {
	retValue, err := instance.GetProperty("SetAllowedStorageLocations")
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
