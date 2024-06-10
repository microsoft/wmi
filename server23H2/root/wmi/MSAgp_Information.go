// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSAgp_Information struct
type MSAgp_Information struct {
	*MSAgp

	//
	Active bool

	//
	AgpCommand uint32

	//
	AgpStatus uint32

	//
	ApertureBase uint64

	//
	ApertureLength uint32

	//
	InstanceName string
}

func NewMSAgp_InformationEx1(instance *cim.WmiInstance) (newInstance *MSAgp_Information, err error) {
	tmp, err := NewMSAgpEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSAgp_Information{
		MSAgp: tmp,
	}
	return
}

func NewMSAgp_InformationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSAgp_Information, err error) {
	tmp, err := NewMSAgpEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSAgp_Information{
		MSAgp: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSAgp_Information) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSAgp_Information) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
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

// SetAgpCommand sets the value of AgpCommand for the instance
func (instance *MSAgp_Information) SetPropertyAgpCommand(value uint32) (err error) {
	return instance.SetProperty("AgpCommand", (value))
}

// GetAgpCommand gets the value of AgpCommand for the instance
func (instance *MSAgp_Information) GetPropertyAgpCommand() (value uint32, err error) {
	retValue, err := instance.GetProperty("AgpCommand")
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

// SetAgpStatus sets the value of AgpStatus for the instance
func (instance *MSAgp_Information) SetPropertyAgpStatus(value uint32) (err error) {
	return instance.SetProperty("AgpStatus", (value))
}

// GetAgpStatus gets the value of AgpStatus for the instance
func (instance *MSAgp_Information) GetPropertyAgpStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("AgpStatus")
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

// SetApertureBase sets the value of ApertureBase for the instance
func (instance *MSAgp_Information) SetPropertyApertureBase(value uint64) (err error) {
	return instance.SetProperty("ApertureBase", (value))
}

// GetApertureBase gets the value of ApertureBase for the instance
func (instance *MSAgp_Information) GetPropertyApertureBase() (value uint64, err error) {
	retValue, err := instance.GetProperty("ApertureBase")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetApertureLength sets the value of ApertureLength for the instance
func (instance *MSAgp_Information) SetPropertyApertureLength(value uint32) (err error) {
	return instance.SetProperty("ApertureLength", (value))
}

// GetApertureLength gets the value of ApertureLength for the instance
func (instance *MSAgp_Information) GetPropertyApertureLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("ApertureLength")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSAgp_Information) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSAgp_Information) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
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
