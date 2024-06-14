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

// FtpCommandFiltering struct
type FtpCommandFiltering struct {
	*EmbeddedObject

	//
	AllowUnlisted bool

	//
	CommandFiltering []FtpAllowedCommandElement

	//
	MaxCommandLine uint32
}

func NewFtpCommandFilteringEx1(instance *cim.WmiInstance) (newInstance *FtpCommandFiltering, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FtpCommandFiltering{
		EmbeddedObject: tmp,
	}
	return
}

func NewFtpCommandFilteringEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FtpCommandFiltering, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FtpCommandFiltering{
		EmbeddedObject: tmp,
	}
	return
}

// SetAllowUnlisted sets the value of AllowUnlisted for the instance
func (instance *FtpCommandFiltering) SetPropertyAllowUnlisted(value bool) (err error) {
	return instance.SetProperty("AllowUnlisted", (value))
}

// GetAllowUnlisted gets the value of AllowUnlisted for the instance
func (instance *FtpCommandFiltering) GetPropertyAllowUnlisted() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowUnlisted")
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

// SetCommandFiltering sets the value of CommandFiltering for the instance
func (instance *FtpCommandFiltering) SetPropertyCommandFiltering(value []FtpAllowedCommandElement) (err error) {
	return instance.SetProperty("CommandFiltering", (value))
}

// GetCommandFiltering gets the value of CommandFiltering for the instance
func (instance *FtpCommandFiltering) GetPropertyCommandFiltering() (value []FtpAllowedCommandElement, err error) {
	retValue, err := instance.GetProperty("CommandFiltering")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(FtpAllowedCommandElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " FtpAllowedCommandElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, FtpAllowedCommandElement(valuetmp))
	}

	return
}

// SetMaxCommandLine sets the value of MaxCommandLine for the instance
func (instance *FtpCommandFiltering) SetPropertyMaxCommandLine(value uint32) (err error) {
	return instance.SetProperty("MaxCommandLine", (value))
}

// GetMaxCommandLine gets the value of MaxCommandLine for the instance
func (instance *FtpCommandFiltering) GetPropertyMaxCommandLine() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxCommandLine")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}
