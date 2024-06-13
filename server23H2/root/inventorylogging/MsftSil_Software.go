// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.InventoryLogging
//////////////////////////////////////////////
package inventorylogging

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MsftSil_Software struct
type MsftSil_Software struct {
	*MsftSil_Data

	//
	ID string

	//
	InstallDate string

	//
	Name string

	//
	Publisher string

	//
	Version string
}

func NewMsftSil_SoftwareEx1(instance *cim.WmiInstance) (newInstance *MsftSil_Software, err error) {
	tmp, err := NewMsftSil_DataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MsftSil_Software{
		MsftSil_Data: tmp,
	}
	return
}

func NewMsftSil_SoftwareEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MsftSil_Software, err error) {
	tmp, err := NewMsftSil_DataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MsftSil_Software{
		MsftSil_Data: tmp,
	}
	return
}

// SetID sets the value of ID for the instance
func (instance *MsftSil_Software) SetPropertyID(value string) (err error) {
	return instance.SetProperty("ID", (value))
}

// GetID gets the value of ID for the instance
func (instance *MsftSil_Software) GetPropertyID() (value string, err error) {
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

// SetInstallDate sets the value of InstallDate for the instance
func (instance *MsftSil_Software) SetPropertyInstallDate(value string) (err error) {
	return instance.SetProperty("InstallDate", (value))
}

// GetInstallDate gets the value of InstallDate for the instance
func (instance *MsftSil_Software) GetPropertyInstallDate() (value string, err error) {
	retValue, err := instance.GetProperty("InstallDate")
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
func (instance *MsftSil_Software) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MsftSil_Software) GetPropertyName() (value string, err error) {
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

// SetPublisher sets the value of Publisher for the instance
func (instance *MsftSil_Software) SetPropertyPublisher(value string) (err error) {
	return instance.SetProperty("Publisher", (value))
}

// GetPublisher gets the value of Publisher for the instance
func (instance *MsftSil_Software) GetPropertyPublisher() (value string, err error) {
	retValue, err := instance.GetProperty("Publisher")
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

// SetVersion sets the value of Version for the instance
func (instance *MsftSil_Software) SetPropertyVersion(value string) (err error) {
	return instance.SetProperty("Version", (value))
}

// GetVersion gets the value of Version for the instance
func (instance *MsftSil_Software) GetPropertyVersion() (value string, err error) {
	retValue, err := instance.GetProperty("Version")
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
