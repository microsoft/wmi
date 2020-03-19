// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_InstalledSoftwareIdentity struct
type CIM_InstalledSoftwareIdentity struct {
	*cim.WmiInstance

	//
	InstalledSoftware CIM_SoftwareIdentity

	//
	System CIM_System
}

func NewCIM_InstalledSoftwareIdentityEx1(instance *cim.WmiInstance) (newInstance *CIM_InstalledSoftwareIdentity, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &CIM_InstalledSoftwareIdentity{
		WmiInstance: tmp,
	}
	return
}

func NewCIM_InstalledSoftwareIdentityEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_InstalledSoftwareIdentity, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_InstalledSoftwareIdentity{
		WmiInstance: tmp,
	}
	return
}

// SetInstalledSoftware sets the value of InstalledSoftware for the instance
func (instance *CIM_InstalledSoftwareIdentity) SetPropertyInstalledSoftware(value CIM_SoftwareIdentity) (err error) {
	return instance.SetProperty("InstalledSoftware", value)
}

// GetInstalledSoftware gets the value of InstalledSoftware for the instance
func (instance *CIM_InstalledSoftwareIdentity) GetPropertyInstalledSoftware() (value CIM_SoftwareIdentity, err error) {
	retValue, err := instance.GetProperty("InstalledSoftware")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_SoftwareIdentity)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSystem sets the value of System for the instance
func (instance *CIM_InstalledSoftwareIdentity) SetPropertySystem(value CIM_System) (err error) {
	return instance.SetProperty("System", value)
}

// GetSystem gets the value of System for the instance
func (instance *CIM_InstalledSoftwareIdentity) GetPropertySystem() (value CIM_System, err error) {
	retValue, err := instance.GetProperty("System")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_System)
	if !ok {
		// TODO: Set an error
	}
	return
}
