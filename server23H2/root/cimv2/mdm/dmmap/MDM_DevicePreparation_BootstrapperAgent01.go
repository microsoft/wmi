// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
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

// MDM_DevicePreparation_BootstrapperAgent01 struct
type MDM_DevicePreparation_BootstrapperAgent01 struct {
	*cim.WmiInstance

	//
	ClassID string

	//
	ExecutionContext string

	//
	InstallationStatusUri string

	//
	InstanceID string

	//
	ParentID string
}

func NewMDM_DevicePreparation_BootstrapperAgent01Ex1(instance *cim.WmiInstance) (newInstance *MDM_DevicePreparation_BootstrapperAgent01, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_DevicePreparation_BootstrapperAgent01{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_DevicePreparation_BootstrapperAgent01Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_DevicePreparation_BootstrapperAgent01, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_DevicePreparation_BootstrapperAgent01{
		WmiInstance: tmp,
	}
	return
}

// SetClassID sets the value of ClassID for the instance
func (instance *MDM_DevicePreparation_BootstrapperAgent01) SetPropertyClassID(value string) (err error) {
	return instance.SetProperty("ClassID", (value))
}

// GetClassID gets the value of ClassID for the instance
func (instance *MDM_DevicePreparation_BootstrapperAgent01) GetPropertyClassID() (value string, err error) {
	retValue, err := instance.GetProperty("ClassID")
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

// SetExecutionContext sets the value of ExecutionContext for the instance
func (instance *MDM_DevicePreparation_BootstrapperAgent01) SetPropertyExecutionContext(value string) (err error) {
	return instance.SetProperty("ExecutionContext", (value))
}

// GetExecutionContext gets the value of ExecutionContext for the instance
func (instance *MDM_DevicePreparation_BootstrapperAgent01) GetPropertyExecutionContext() (value string, err error) {
	retValue, err := instance.GetProperty("ExecutionContext")
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

// SetInstallationStatusUri sets the value of InstallationStatusUri for the instance
func (instance *MDM_DevicePreparation_BootstrapperAgent01) SetPropertyInstallationStatusUri(value string) (err error) {
	return instance.SetProperty("InstallationStatusUri", (value))
}

// GetInstallationStatusUri gets the value of InstallationStatusUri for the instance
func (instance *MDM_DevicePreparation_BootstrapperAgent01) GetPropertyInstallationStatusUri() (value string, err error) {
	retValue, err := instance.GetProperty("InstallationStatusUri")
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

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_DevicePreparation_BootstrapperAgent01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_DevicePreparation_BootstrapperAgent01) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_DevicePreparation_BootstrapperAgent01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_DevicePreparation_BootstrapperAgent01) GetPropertyParentID() (value string, err error) {
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
