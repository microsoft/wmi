// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.SDDC.Management
//////////////////////////////////////////////
package management

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// SDDC_VmSnapshot struct
type SDDC_VmSnapshot struct {
	*cim.WmiInstance

	//
	CreationTime string

	//
	Id string

	//
	IsCurrentlyApplied bool

	//
	Name string

	//
	ParentId string
}

func NewSDDC_VmSnapshotEx1(instance *cim.WmiInstance) (newInstance *SDDC_VmSnapshot, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &SDDC_VmSnapshot{
		WmiInstance: tmp,
	}
	return
}

func NewSDDC_VmSnapshotEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SDDC_VmSnapshot, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SDDC_VmSnapshot{
		WmiInstance: tmp,
	}
	return
}

// SetCreationTime sets the value of CreationTime for the instance
func (instance *SDDC_VmSnapshot) SetPropertyCreationTime(value string) (err error) {
	return instance.SetProperty("CreationTime", value)
}

// GetCreationTime gets the value of CreationTime for the instance
func (instance *SDDC_VmSnapshot) GetPropertyCreationTime() (value string, err error) {
	retValue, err := instance.GetProperty("CreationTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetId sets the value of Id for the instance
func (instance *SDDC_VmSnapshot) SetPropertyId(value string) (err error) {
	return instance.SetProperty("Id", value)
}

// GetId gets the value of Id for the instance
func (instance *SDDC_VmSnapshot) GetPropertyId() (value string, err error) {
	retValue, err := instance.GetProperty("Id")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsCurrentlyApplied sets the value of IsCurrentlyApplied for the instance
func (instance *SDDC_VmSnapshot) SetPropertyIsCurrentlyApplied(value bool) (err error) {
	return instance.SetProperty("IsCurrentlyApplied", value)
}

// GetIsCurrentlyApplied gets the value of IsCurrentlyApplied for the instance
func (instance *SDDC_VmSnapshot) GetPropertyIsCurrentlyApplied() (value bool, err error) {
	retValue, err := instance.GetProperty("IsCurrentlyApplied")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *SDDC_VmSnapshot) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *SDDC_VmSnapshot) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentId sets the value of ParentId for the instance
func (instance *SDDC_VmSnapshot) SetPropertyParentId(value string) (err error) {
	return instance.SetProperty("ParentId", value)
}

// GetParentId gets the value of ParentId for the instance
func (instance *SDDC_VmSnapshot) GetPropertyParentId() (value string, err error) {
	retValue, err := instance.GetProperty("ParentId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
