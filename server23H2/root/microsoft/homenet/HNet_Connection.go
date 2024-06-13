// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.HomeNet
//////////////////////////////////////////////
package homenet

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// HNet_Connection struct
type HNet_Connection struct {
	*cim.WmiInstance

	//
	Guid string

	//
	IsLanConnection bool

	//
	Name string

	//
	PhonebookPath string
}

func NewHNet_ConnectionEx1(instance *cim.WmiInstance) (newInstance *HNet_Connection, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &HNet_Connection{
		WmiInstance: tmp,
	}
	return
}

func NewHNet_ConnectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HNet_Connection, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HNet_Connection{
		WmiInstance: tmp,
	}
	return
}

// SetGuid sets the value of Guid for the instance
func (instance *HNet_Connection) SetPropertyGuid(value string) (err error) {
	return instance.SetProperty("Guid", (value))
}

// GetGuid gets the value of Guid for the instance
func (instance *HNet_Connection) GetPropertyGuid() (value string, err error) {
	retValue, err := instance.GetProperty("Guid")
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

// SetIsLanConnection sets the value of IsLanConnection for the instance
func (instance *HNet_Connection) SetPropertyIsLanConnection(value bool) (err error) {
	return instance.SetProperty("IsLanConnection", (value))
}

// GetIsLanConnection gets the value of IsLanConnection for the instance
func (instance *HNet_Connection) GetPropertyIsLanConnection() (value bool, err error) {
	retValue, err := instance.GetProperty("IsLanConnection")
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

// SetName sets the value of Name for the instance
func (instance *HNet_Connection) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *HNet_Connection) GetPropertyName() (value string, err error) {
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

// SetPhonebookPath sets the value of PhonebookPath for the instance
func (instance *HNet_Connection) SetPropertyPhonebookPath(value string) (err error) {
	return instance.SetProperty("PhonebookPath", (value))
}

// GetPhonebookPath gets the value of PhonebookPath for the instance
func (instance *HNet_Connection) GetPropertyPhonebookPath() (value string, err error) {
	retValue, err := instance.GetProperty("PhonebookPath")
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
