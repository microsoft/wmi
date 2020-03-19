// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.Powershellv3
//////////////////////////////////////////////
package powershellv3

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// PS_Module struct
type PS_Module struct {
	*CIM_ManagedElement

	//
	moduleManifestFileData []uint8

	//
	ModuleName string

	//
	moduleType uint16
}

func NewPS_ModuleEx1(instance *cim.WmiInstance) (newInstance *PS_Module, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &PS_Module{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewPS_ModuleEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_Module, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_Module{
		CIM_ManagedElement: tmp,
	}
	return
}

// SetmoduleManifestFileData sets the value of moduleManifestFileData for the instance
func (instance *PS_Module) SetPropertymoduleManifestFileData(value []uint8) (err error) {
	return instance.SetProperty("moduleManifestFileData", value)
}

// GetmoduleManifestFileData gets the value of moduleManifestFileData for the instance
func (instance *PS_Module) GetPropertymoduleManifestFileData() (value []uint8, err error) {
	retValue, err := instance.GetProperty("moduleManifestFileData")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetModuleName sets the value of ModuleName for the instance
func (instance *PS_Module) SetPropertyModuleName(value string) (err error) {
	return instance.SetProperty("ModuleName", value)
}

// GetModuleName gets the value of ModuleName for the instance
func (instance *PS_Module) GetPropertyModuleName() (value string, err error) {
	retValue, err := instance.GetProperty("ModuleName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetmoduleType sets the value of moduleType for the instance
func (instance *PS_Module) SetPropertymoduleType(value uint16) (err error) {
	return instance.SetProperty("moduleType", value)
}

// GetmoduleType gets the value of moduleType for the instance
func (instance *PS_Module) GetPropertymoduleType() (value uint16, err error) {
	retValue, err := instance.GetProperty("moduleType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
