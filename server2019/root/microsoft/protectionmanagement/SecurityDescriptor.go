// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.protectionManagement
//////////////////////////////////////////////
package protectionmanagement

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __SecurityDescriptor struct
type __SecurityDescriptor struct {
	*__SecurityRelatedClass

	//
	ControlFlags uint32

	//
	DACL []__ACE

	//
	Group __ACE

	//
	Owner __ACE

	//
	SACL []__ACE

	//
	TIME_CREATED uint64
}

func New__SecurityDescriptorEx1(instance *cim.WmiInstance) (newInstance *__SecurityDescriptor, err error) {
	tmp, err := New__SecurityRelatedClassEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__SecurityDescriptor{
		__SecurityRelatedClass: tmp,
	}
	return
}

func New__SecurityDescriptorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__SecurityDescriptor, err error) {
	tmp, err := New__SecurityRelatedClassEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__SecurityDescriptor{
		__SecurityRelatedClass: tmp,
	}
	return
}

// SetControlFlags sets the value of ControlFlags for the instance
func (instance *__SecurityDescriptor) SetPropertyControlFlags(value uint32) (err error) {
	return instance.SetProperty("ControlFlags", value)
}

// GetControlFlags gets the value of ControlFlags for the instance
func (instance *__SecurityDescriptor) GetPropertyControlFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("ControlFlags")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDACL sets the value of DACL for the instance
func (instance *__SecurityDescriptor) SetPropertyDACL(value []__ACE) (err error) {
	return instance.SetProperty("DACL", value)
}

// GetDACL gets the value of DACL for the instance
func (instance *__SecurityDescriptor) GetPropertyDACL() (value []__ACE, err error) {
	retValue, err := instance.GetProperty("DACL")
	if err != nil {
		return
	}
	value, ok := retValue.([]__ACE)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGroup sets the value of Group for the instance
func (instance *__SecurityDescriptor) SetPropertyGroup(value __ACE) (err error) {
	return instance.SetProperty("Group", value)
}

// GetGroup gets the value of Group for the instance
func (instance *__SecurityDescriptor) GetPropertyGroup() (value __ACE, err error) {
	retValue, err := instance.GetProperty("Group")
	if err != nil {
		return
	}
	value, ok := retValue.(__ACE)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOwner sets the value of Owner for the instance
func (instance *__SecurityDescriptor) SetPropertyOwner(value __ACE) (err error) {
	return instance.SetProperty("Owner", value)
}

// GetOwner gets the value of Owner for the instance
func (instance *__SecurityDescriptor) GetPropertyOwner() (value __ACE, err error) {
	retValue, err := instance.GetProperty("Owner")
	if err != nil {
		return
	}
	value, ok := retValue.(__ACE)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSACL sets the value of SACL for the instance
func (instance *__SecurityDescriptor) SetPropertySACL(value []__ACE) (err error) {
	return instance.SetProperty("SACL", value)
}

// GetSACL gets the value of SACL for the instance
func (instance *__SecurityDescriptor) GetPropertySACL() (value []__ACE, err error) {
	retValue, err := instance.GetProperty("SACL")
	if err != nil {
		return
	}
	value, ok := retValue.([]__ACE)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTIME_CREATED sets the value of TIME_CREATED for the instance
func (instance *__SecurityDescriptor) SetPropertyTIME_CREATED(value uint64) (err error) {
	return instance.SetProperty("TIME_CREATED", value)
}

// GetTIME_CREATED gets the value of TIME_CREATED for the instance
func (instance *__SecurityDescriptor) GetPropertyTIME_CREATED() (value uint64, err error) {
	retValue, err := instance.GetProperty("TIME_CREATED")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
