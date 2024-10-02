// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.InventoryLogging
//////////////////////////////////////////////
package inventorylogging

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MsftSil_ComputerIdentity struct
type MsftSil_ComputerIdentity struct {
	*MsftSil_Data

	//
	HypervisorHostName string

	//
	ID string

	//
	Name string

	//
	UUID string

	//
	VMGUID string
}

func NewMsftSil_ComputerIdentityEx1(instance *cim.WmiInstance) (newInstance *MsftSil_ComputerIdentity, err error) {
	tmp, err := NewMsftSil_DataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MsftSil_ComputerIdentity{
		MsftSil_Data: tmp,
	}
	return
}

func NewMsftSil_ComputerIdentityEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MsftSil_ComputerIdentity, err error) {
	tmp, err := NewMsftSil_DataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MsftSil_ComputerIdentity{
		MsftSil_Data: tmp,
	}
	return
}

// SetHypervisorHostName sets the value of HypervisorHostName for the instance
func (instance *MsftSil_ComputerIdentity) SetPropertyHypervisorHostName(value string) (err error) {
	return instance.SetProperty("HypervisorHostName", (value))
}

// GetHypervisorHostName gets the value of HypervisorHostName for the instance
func (instance *MsftSil_ComputerIdentity) GetPropertyHypervisorHostName() (value string, err error) {
	retValue, err := instance.GetProperty("HypervisorHostName")
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

// SetID sets the value of ID for the instance
func (instance *MsftSil_ComputerIdentity) SetPropertyID(value string) (err error) {
	return instance.SetProperty("ID", (value))
}

// GetID gets the value of ID for the instance
func (instance *MsftSil_ComputerIdentity) GetPropertyID() (value string, err error) {
	retValue, err := instance.GetProperty("ID")
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

// SetName sets the value of Name for the instance
func (instance *MsftSil_ComputerIdentity) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MsftSil_ComputerIdentity) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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

// SetUUID sets the value of UUID for the instance
func (instance *MsftSil_ComputerIdentity) SetPropertyUUID(value string) (err error) {
	return instance.SetProperty("UUID", (value))
}

// GetUUID gets the value of UUID for the instance
func (instance *MsftSil_ComputerIdentity) GetPropertyUUID() (value string, err error) {
	retValue, err := instance.GetProperty("UUID")
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

// SetVMGUID sets the value of VMGUID for the instance
func (instance *MsftSil_ComputerIdentity) SetPropertyVMGUID(value string) (err error) {
	return instance.SetProperty("VMGUID", (value))
}

// GetVMGUID gets the value of VMGUID for the instance
func (instance *MsftSil_ComputerIdentity) GetPropertyVMGUID() (value string, err error) {
	retValue, err := instance.GetProperty("VMGUID")
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
