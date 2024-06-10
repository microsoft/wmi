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

// NtlmServerAccept_Info struct
type NtlmServerAccept_Info struct {
	*NtlmServerAccept

	// Client Domain Name
	DomainName string

	// Flags
	Flags uint32

	// In-Context
	InContext uint32

	// Out-Context
	OutContext uint32

	// Stage Hint
	StageHint uint32

	// Client User Name
	UserName string

	// Client Workstation
	Workstation string
}

func NewNtlmServerAccept_InfoEx1(instance *cim.WmiInstance) (newInstance *NtlmServerAccept_Info, err error) {
	tmp, err := NewNtlmServerAcceptEx1(instance)

	if err != nil {
		return
	}
	newInstance = &NtlmServerAccept_Info{
		NtlmServerAccept: tmp,
	}
	return
}

func NewNtlmServerAccept_InfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *NtlmServerAccept_Info, err error) {
	tmp, err := NewNtlmServerAcceptEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &NtlmServerAccept_Info{
		NtlmServerAccept: tmp,
	}
	return
}

// SetDomainName sets the value of DomainName for the instance
func (instance *NtlmServerAccept_Info) SetPropertyDomainName(value string) (err error) {
	return instance.SetProperty("DomainName", (value))
}

// GetDomainName gets the value of DomainName for the instance
func (instance *NtlmServerAccept_Info) GetPropertyDomainName() (value string, err error) {
	retValue, err := instance.GetProperty("DomainName")
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

// SetFlags sets the value of Flags for the instance
func (instance *NtlmServerAccept_Info) SetPropertyFlags(value uint32) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *NtlmServerAccept_Info) GetPropertyFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("Flags")
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

// SetInContext sets the value of InContext for the instance
func (instance *NtlmServerAccept_Info) SetPropertyInContext(value uint32) (err error) {
	return instance.SetProperty("InContext", (value))
}

// GetInContext gets the value of InContext for the instance
func (instance *NtlmServerAccept_Info) GetPropertyInContext() (value uint32, err error) {
	retValue, err := instance.GetProperty("InContext")
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

// SetOutContext sets the value of OutContext for the instance
func (instance *NtlmServerAccept_Info) SetPropertyOutContext(value uint32) (err error) {
	return instance.SetProperty("OutContext", (value))
}

// GetOutContext gets the value of OutContext for the instance
func (instance *NtlmServerAccept_Info) GetPropertyOutContext() (value uint32, err error) {
	retValue, err := instance.GetProperty("OutContext")
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

// SetStageHint sets the value of StageHint for the instance
func (instance *NtlmServerAccept_Info) SetPropertyStageHint(value uint32) (err error) {
	return instance.SetProperty("StageHint", (value))
}

// GetStageHint gets the value of StageHint for the instance
func (instance *NtlmServerAccept_Info) GetPropertyStageHint() (value uint32, err error) {
	retValue, err := instance.GetProperty("StageHint")
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

// SetUserName sets the value of UserName for the instance
func (instance *NtlmServerAccept_Info) SetPropertyUserName(value string) (err error) {
	return instance.SetProperty("UserName", (value))
}

// GetUserName gets the value of UserName for the instance
func (instance *NtlmServerAccept_Info) GetPropertyUserName() (value string, err error) {
	retValue, err := instance.GetProperty("UserName")
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

// SetWorkstation sets the value of Workstation for the instance
func (instance *NtlmServerAccept_Info) SetPropertyWorkstation(value string) (err error) {
	return instance.SetProperty("Workstation", (value))
}

// GetWorkstation gets the value of Workstation for the instance
func (instance *NtlmServerAccept_Info) GetPropertyWorkstation() (value string, err error) {
	retValue, err := instance.GetProperty("Workstation")
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
