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

// MDM_EnterpriseModernAppManagement_AppManagement01_02 struct
type MDM_EnterpriseModernAppManagement_AppManagement01_02 struct {
	*cim.WmiInstance

	//
	DoNotUpdate int32

	//
	InstanceID string

	//
	MaintainProcessorArchitectureOnUpdate int32

	//
	ParentID string
}

func NewMDM_EnterpriseModernAppManagement_AppManagement01_02Ex1(instance *cim.WmiInstance) (newInstance *MDM_EnterpriseModernAppManagement_AppManagement01_02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_EnterpriseModernAppManagement_AppManagement01_02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_EnterpriseModernAppManagement_AppManagement01_02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_EnterpriseModernAppManagement_AppManagement01_02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_EnterpriseModernAppManagement_AppManagement01_02{
		WmiInstance: tmp,
	}
	return
}

// SetDoNotUpdate sets the value of DoNotUpdate for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_02) SetPropertyDoNotUpdate(value int32) (err error) {
	return instance.SetProperty("DoNotUpdate", (value))
}

// GetDoNotUpdate gets the value of DoNotUpdate for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_02) GetPropertyDoNotUpdate() (value int32, err error) {
	retValue, err := instance.GetProperty("DoNotUpdate")
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
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_02) GetPropertyInstanceID() (value string, err error) {
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

// SetMaintainProcessorArchitectureOnUpdate sets the value of MaintainProcessorArchitectureOnUpdate for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_02) SetPropertyMaintainProcessorArchitectureOnUpdate(value int32) (err error) {
	return instance.SetProperty("MaintainProcessorArchitectureOnUpdate", (value))
}

// GetMaintainProcessorArchitectureOnUpdate gets the value of MaintainProcessorArchitectureOnUpdate for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_02) GetPropertyMaintainProcessorArchitectureOnUpdate() (value int32, err error) {
	retValue, err := instance.GetProperty("MaintainProcessorArchitectureOnUpdate")
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_02) GetPropertyParentID() (value string, err error) {
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
