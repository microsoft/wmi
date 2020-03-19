// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_System struct
type CIM_System struct {
	*CIM_EnabledLogicalElement

	//
	CreationClassName string

	//
	IdentifyingDescriptions []string

	//
	NameFormat string

	//
	OtherIdentifyingInfo []string

	//
	PrimaryOwnerContact string

	//
	PrimaryOwnerName string

	//
	Roles []string
}

func NewCIM_SystemEx1(instance *cim.WmiInstance) (newInstance *CIM_System, err error) {
	tmp, err := NewCIM_EnabledLogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_System{
		CIM_EnabledLogicalElement: tmp,
	}
	return
}

func NewCIM_SystemEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_System, err error) {
	tmp, err := NewCIM_EnabledLogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_System{
		CIM_EnabledLogicalElement: tmp,
	}
	return
}

// SetCreationClassName sets the value of CreationClassName for the instance
func (instance *CIM_System) SetPropertyCreationClassName(value string) (err error) {
	return instance.SetProperty("CreationClassName", value)
}

// GetCreationClassName gets the value of CreationClassName for the instance
func (instance *CIM_System) GetPropertyCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("CreationClassName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIdentifyingDescriptions sets the value of IdentifyingDescriptions for the instance
func (instance *CIM_System) SetPropertyIdentifyingDescriptions(value []string) (err error) {
	return instance.SetProperty("IdentifyingDescriptions", value)
}

// GetIdentifyingDescriptions gets the value of IdentifyingDescriptions for the instance
func (instance *CIM_System) GetPropertyIdentifyingDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("IdentifyingDescriptions")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNameFormat sets the value of NameFormat for the instance
func (instance *CIM_System) SetPropertyNameFormat(value string) (err error) {
	return instance.SetProperty("NameFormat", value)
}

// GetNameFormat gets the value of NameFormat for the instance
func (instance *CIM_System) GetPropertyNameFormat() (value string, err error) {
	retValue, err := instance.GetProperty("NameFormat")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherIdentifyingInfo sets the value of OtherIdentifyingInfo for the instance
func (instance *CIM_System) SetPropertyOtherIdentifyingInfo(value []string) (err error) {
	return instance.SetProperty("OtherIdentifyingInfo", value)
}

// GetOtherIdentifyingInfo gets the value of OtherIdentifyingInfo for the instance
func (instance *CIM_System) GetPropertyOtherIdentifyingInfo() (value []string, err error) {
	retValue, err := instance.GetProperty("OtherIdentifyingInfo")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPrimaryOwnerContact sets the value of PrimaryOwnerContact for the instance
func (instance *CIM_System) SetPropertyPrimaryOwnerContact(value string) (err error) {
	return instance.SetProperty("PrimaryOwnerContact", value)
}

// GetPrimaryOwnerContact gets the value of PrimaryOwnerContact for the instance
func (instance *CIM_System) GetPropertyPrimaryOwnerContact() (value string, err error) {
	retValue, err := instance.GetProperty("PrimaryOwnerContact")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPrimaryOwnerName sets the value of PrimaryOwnerName for the instance
func (instance *CIM_System) SetPropertyPrimaryOwnerName(value string) (err error) {
	return instance.SetProperty("PrimaryOwnerName", value)
}

// GetPrimaryOwnerName gets the value of PrimaryOwnerName for the instance
func (instance *CIM_System) GetPropertyPrimaryOwnerName() (value string, err error) {
	retValue, err := instance.GetProperty("PrimaryOwnerName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRoles sets the value of Roles for the instance
func (instance *CIM_System) SetPropertyRoles(value []string) (err error) {
	return instance.SetProperty("Roles", value)
}

// GetRoles gets the value of Roles for the instance
func (instance *CIM_System) GetPropertyRoles() (value []string, err error) {
	retValue, err := instance.GetProperty("Roles")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
