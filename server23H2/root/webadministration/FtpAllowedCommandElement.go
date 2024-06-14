// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// FtpAllowedCommandElement struct
type FtpAllowedCommandElement struct {
	*CollectionElement

	//
	Allowed bool

	//
	Command string
}

func NewFtpAllowedCommandElementEx1(instance *cim.WmiInstance) (newInstance *FtpAllowedCommandElement, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FtpAllowedCommandElement{
		CollectionElement: tmp,
	}
	return
}

func NewFtpAllowedCommandElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FtpAllowedCommandElement, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FtpAllowedCommandElement{
		CollectionElement: tmp,
	}
	return
}

// SetAllowed sets the value of Allowed for the instance
func (instance *FtpAllowedCommandElement) SetPropertyAllowed(value bool) (err error) {
	return instance.SetProperty("Allowed", (value))
}

// GetAllowed gets the value of Allowed for the instance
func (instance *FtpAllowedCommandElement) GetPropertyAllowed() (value bool, err error) {
	retValue, err := instance.GetProperty("Allowed")
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

// SetCommand sets the value of Command for the instance
func (instance *FtpAllowedCommandElement) SetPropertyCommand(value string) (err error) {
	return instance.SetProperty("Command", (value))
}

// GetCommand gets the value of Command for the instance
func (instance *FtpAllowedCommandElement) GetPropertyCommand() (value string, err error) {
	retValue, err := instance.GetProperty("Command")
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
