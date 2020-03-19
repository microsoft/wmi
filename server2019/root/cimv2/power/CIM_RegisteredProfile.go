// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2.power
//////////////////////////////////////////////
package power

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_RegisteredProfile struct
type CIM_RegisteredProfile struct {
	*CIM_ManagedElement

	//
	AdvertiseTypeDescriptions []string

	//
	AdvertiseTypes []uint16

	//
	OtherRegisteredOrganization string

	//
	RegisteredName string

	//
	RegisteredOrganization uint16

	//
	RegisteredVersion string
}

func NewCIM_RegisteredProfileEx1(instance *cim.WmiInstance) (newInstance *CIM_RegisteredProfile, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_RegisteredProfile{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewCIM_RegisteredProfileEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_RegisteredProfile, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_RegisteredProfile{
		CIM_ManagedElement: tmp,
	}
	return
}

// SetAdvertiseTypeDescriptions sets the value of AdvertiseTypeDescriptions for the instance
func (instance *CIM_RegisteredProfile) SetPropertyAdvertiseTypeDescriptions(value []string) (err error) {
	return instance.SetProperty("AdvertiseTypeDescriptions", value)
}

// GetAdvertiseTypeDescriptions gets the value of AdvertiseTypeDescriptions for the instance
func (instance *CIM_RegisteredProfile) GetPropertyAdvertiseTypeDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("AdvertiseTypeDescriptions")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAdvertiseTypes sets the value of AdvertiseTypes for the instance
func (instance *CIM_RegisteredProfile) SetPropertyAdvertiseTypes(value []uint16) (err error) {
	return instance.SetProperty("AdvertiseTypes", value)
}

// GetAdvertiseTypes gets the value of AdvertiseTypes for the instance
func (instance *CIM_RegisteredProfile) GetPropertyAdvertiseTypes() (value []uint16, err error) {
	retValue, err := instance.GetProperty("AdvertiseTypes")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherRegisteredOrganization sets the value of OtherRegisteredOrganization for the instance
func (instance *CIM_RegisteredProfile) SetPropertyOtherRegisteredOrganization(value string) (err error) {
	return instance.SetProperty("OtherRegisteredOrganization", value)
}

// GetOtherRegisteredOrganization gets the value of OtherRegisteredOrganization for the instance
func (instance *CIM_RegisteredProfile) GetPropertyOtherRegisteredOrganization() (value string, err error) {
	retValue, err := instance.GetProperty("OtherRegisteredOrganization")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRegisteredName sets the value of RegisteredName for the instance
func (instance *CIM_RegisteredProfile) SetPropertyRegisteredName(value string) (err error) {
	return instance.SetProperty("RegisteredName", value)
}

// GetRegisteredName gets the value of RegisteredName for the instance
func (instance *CIM_RegisteredProfile) GetPropertyRegisteredName() (value string, err error) {
	retValue, err := instance.GetProperty("RegisteredName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRegisteredOrganization sets the value of RegisteredOrganization for the instance
func (instance *CIM_RegisteredProfile) SetPropertyRegisteredOrganization(value uint16) (err error) {
	return instance.SetProperty("RegisteredOrganization", value)
}

// GetRegisteredOrganization gets the value of RegisteredOrganization for the instance
func (instance *CIM_RegisteredProfile) GetPropertyRegisteredOrganization() (value uint16, err error) {
	retValue, err := instance.GetProperty("RegisteredOrganization")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRegisteredVersion sets the value of RegisteredVersion for the instance
func (instance *CIM_RegisteredProfile) SetPropertyRegisteredVersion(value string) (err error) {
	return instance.SetProperty("RegisteredVersion", value)
}

// GetRegisteredVersion gets the value of RegisteredVersion for the instance
func (instance *CIM_RegisteredProfile) GetPropertyRegisteredVersion() (value string, err error) {
	retValue, err := instance.GetProperty("RegisteredVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
