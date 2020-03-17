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

// MDM_EnterpriseModernAppManagement_AppManagement01_02 struct
type MDM_EnterpriseModernAppManagement_AppManagement01_02 struct {
	cim.WmiInstance

	//
	DoNotUpdate int32

	//
	InstanceID string

	//
	MaintainProcessorArchitectureOnUpdate int32

	//
	ParentID string
}

// SetDoNotUpdate sets the value of DoNotUpdate for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_02) SetPropertyDoNotUpdate(value int32) (err error) {
	return instance.SetProperty("DoNotUpdate", value)
}

// GetDoNotUpdate gets the value of DoNotUpdate for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_02) GetPropertyDoNotUpdate() (value int32, err error) {
	retValue, err := instance.GetProperty("DoNotUpdate")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_02) GetPropertyInstanceID() (value string, err error) {
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

// SetMaintainProcessorArchitectureOnUpdate sets the value of MaintainProcessorArchitectureOnUpdate for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_02) SetPropertyMaintainProcessorArchitectureOnUpdate(value int32) (err error) {
	return instance.SetProperty("MaintainProcessorArchitectureOnUpdate", value)
}

// GetMaintainProcessorArchitectureOnUpdate gets the value of MaintainProcessorArchitectureOnUpdate for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_02) GetPropertyMaintainProcessorArchitectureOnUpdate() (value int32, err error) {
	retValue, err := instance.GetProperty("MaintainProcessorArchitectureOnUpdate")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_02) GetPropertyParentID() (value string, err error) {
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
