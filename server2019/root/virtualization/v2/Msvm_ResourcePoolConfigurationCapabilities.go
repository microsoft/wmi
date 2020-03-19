// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_ResourcePoolConfigurationCapabilities struct
type Msvm_ResourcePoolConfigurationCapabilities struct {
	*CIM_Capabilities

	// This property reflects the methods of the associated service class that are supported that may return a Job.
	AsynchronousMethodsSupported []ResourcePoolConfigurationCapabilities_AsynchronousMethodsSupported

	// This property reflects the methods of the associated service class that are supported and block until completed (no Job is returned.)
	SynchronousMethodsSupported []ResourcePoolConfigurationCapabilities_SynchronousMethodsSupported
}

func NewMsvm_ResourcePoolConfigurationCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *Msvm_ResourcePoolConfigurationCapabilities, err error) {
	tmp, err := NewCIM_CapabilitiesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_ResourcePoolConfigurationCapabilities{
		CIM_Capabilities: tmp,
	}
	return
}

func NewMsvm_ResourcePoolConfigurationCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_ResourcePoolConfigurationCapabilities, err error) {
	tmp, err := NewCIM_CapabilitiesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_ResourcePoolConfigurationCapabilities{
		CIM_Capabilities: tmp,
	}
	return
}

// SetAsynchronousMethodsSupported sets the value of AsynchronousMethodsSupported for the instance
func (instance *Msvm_ResourcePoolConfigurationCapabilities) SetPropertyAsynchronousMethodsSupported(value []ResourcePoolConfigurationCapabilities_AsynchronousMethodsSupported) (err error) {
	return instance.SetProperty("AsynchronousMethodsSupported", value)
}

// GetAsynchronousMethodsSupported gets the value of AsynchronousMethodsSupported for the instance
func (instance *Msvm_ResourcePoolConfigurationCapabilities) GetPropertyAsynchronousMethodsSupported() (value []ResourcePoolConfigurationCapabilities_AsynchronousMethodsSupported, err error) {
	retValue, err := instance.GetProperty("AsynchronousMethodsSupported")
	if err != nil {
		return
	}
	value, ok := retValue.([]ResourcePoolConfigurationCapabilities_AsynchronousMethodsSupported)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSynchronousMethodsSupported sets the value of SynchronousMethodsSupported for the instance
func (instance *Msvm_ResourcePoolConfigurationCapabilities) SetPropertySynchronousMethodsSupported(value []ResourcePoolConfigurationCapabilities_SynchronousMethodsSupported) (err error) {
	return instance.SetProperty("SynchronousMethodsSupported", value)
}

// GetSynchronousMethodsSupported gets the value of SynchronousMethodsSupported for the instance
func (instance *Msvm_ResourcePoolConfigurationCapabilities) GetPropertySynchronousMethodsSupported() (value []ResourcePoolConfigurationCapabilities_SynchronousMethodsSupported, err error) {
	retValue, err := instance.GetProperty("SynchronousMethodsSupported")
	if err != nil {
		return
	}
	value, ok := retValue.([]ResourcePoolConfigurationCapabilities_SynchronousMethodsSupported)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_ResourcePoolConfigurationCapabilities) GetRelatedResourcePoolConfigurationService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ResourcePoolConfigurationService")
}
