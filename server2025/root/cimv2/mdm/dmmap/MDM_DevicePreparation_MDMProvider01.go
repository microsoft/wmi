// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
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

// MDM_DevicePreparation_MDMProvider01 struct
type MDM_DevicePreparation_MDMProvider01 struct {
	*cim.WmiInstance

	//
	InstanceID string

	//
	MdmAgentInstalled bool

	//
	ParentID string

	//
	Progress string

	//
	RebootRequired bool
}

func NewMDM_DevicePreparation_MDMProvider01Ex1(instance *cim.WmiInstance) (newInstance *MDM_DevicePreparation_MDMProvider01, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_DevicePreparation_MDMProvider01{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_DevicePreparation_MDMProvider01Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_DevicePreparation_MDMProvider01, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_DevicePreparation_MDMProvider01{
		WmiInstance: tmp,
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_DevicePreparation_MDMProvider01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_DevicePreparation_MDMProvider01) GetPropertyInstanceID() (value string, err error) {
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

// SetMdmAgentInstalled sets the value of MdmAgentInstalled for the instance
func (instance *MDM_DevicePreparation_MDMProvider01) SetPropertyMdmAgentInstalled(value bool) (err error) {
	return instance.SetProperty("MdmAgentInstalled", (value))
}

// GetMdmAgentInstalled gets the value of MdmAgentInstalled for the instance
func (instance *MDM_DevicePreparation_MDMProvider01) GetPropertyMdmAgentInstalled() (value bool, err error) {
	retValue, err := instance.GetProperty("MdmAgentInstalled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_DevicePreparation_MDMProvider01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_DevicePreparation_MDMProvider01) GetPropertyParentID() (value string, err error) {
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

// SetProgress sets the value of Progress for the instance
func (instance *MDM_DevicePreparation_MDMProvider01) SetPropertyProgress(value string) (err error) {
	return instance.SetProperty("Progress", (value))
}

// GetProgress gets the value of Progress for the instance
func (instance *MDM_DevicePreparation_MDMProvider01) GetPropertyProgress() (value string, err error) {
	retValue, err := instance.GetProperty("Progress")
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

// SetRebootRequired sets the value of RebootRequired for the instance
func (instance *MDM_DevicePreparation_MDMProvider01) SetPropertyRebootRequired(value bool) (err error) {
	return instance.SetProperty("RebootRequired", (value))
}

// GetRebootRequired gets the value of RebootRequired for the instance
func (instance *MDM_DevicePreparation_MDMProvider01) GetPropertyRebootRequired() (value bool, err error) {
	retValue, err := instance.GetProperty("RebootRequired")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}
