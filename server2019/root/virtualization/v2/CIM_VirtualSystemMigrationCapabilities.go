// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.virtualization.v2
//
// ////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_VirtualSystemMigrationCapabilities struct
type CIM_VirtualSystemMigrationCapabilities struct {
	*CIM_Capabilities

	// Enumeration of method identifiers whose implementation may be asynchronous; that is, the operation may not complete immediately and instead the method may return a job.
	AsynchronousMethodsSupported []VirtualSystemMigrationCapabilities_AsynchronousMethodsSupported

	// Array of format designators. Values indicate that the designated format is supported for input values of the DestinationHost parameter of the MigrateVirtualSystemToHost( ) method or the CheckVirtualSystemIsMigratableToHost( ) method of the associated instance of the CIM_VirtualSystemMigrationService class.
	///Format designators designate the following formats:
	///2 - Support of the Domain Name text format according to RFC 1035
	///3 - Support of the IPv4 dotted decimal format according to RFC 1208
	///4 - Support of the IPv6 text format according to RFC 4291
	DestinationHostFormatsSupported []VirtualSystemMigrationCapabilities_DestinationHostFormatsSupported

	// Enumeration of method identifiers whose implementation may be synchronous; that is, the operation may complete immediately and therefore the method may not return a job.
	SynchronousMethodsSupported []VirtualSystemMigrationCapabilities_SynchronousMethodsSupported
}

func NewCIM_VirtualSystemMigrationCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *CIM_VirtualSystemMigrationCapabilities, err error) {
	tmp, err := NewCIM_CapabilitiesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_VirtualSystemMigrationCapabilities{
		CIM_Capabilities: tmp,
	}
	return
}

func NewCIM_VirtualSystemMigrationCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_VirtualSystemMigrationCapabilities, err error) {
	tmp, err := NewCIM_CapabilitiesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_VirtualSystemMigrationCapabilities{
		CIM_Capabilities: tmp,
	}
	return
}

// SetAsynchronousMethodsSupported sets the value of AsynchronousMethodsSupported for the instance
func (instance *CIM_VirtualSystemMigrationCapabilities) SetPropertyAsynchronousMethodsSupported(value []VirtualSystemMigrationCapabilities_AsynchronousMethodsSupported) (err error) {
	return instance.SetProperty("AsynchronousMethodsSupported", (value))
}

// GetAsynchronousMethodsSupported gets the value of AsynchronousMethodsSupported for the instance
func (instance *CIM_VirtualSystemMigrationCapabilities) GetPropertyAsynchronousMethodsSupported() (value []VirtualSystemMigrationCapabilities_AsynchronousMethodsSupported, err error) {
	retValue, err := instance.GetProperty("AsynchronousMethodsSupported")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, VirtualSystemMigrationCapabilities_AsynchronousMethodsSupported(valuetmp))
	}

	return
}

// SetDestinationHostFormatsSupported sets the value of DestinationHostFormatsSupported for the instance
func (instance *CIM_VirtualSystemMigrationCapabilities) SetPropertyDestinationHostFormatsSupported(value []VirtualSystemMigrationCapabilities_DestinationHostFormatsSupported) (err error) {
	return instance.SetProperty("DestinationHostFormatsSupported", (value))
}

// GetDestinationHostFormatsSupported gets the value of DestinationHostFormatsSupported for the instance
func (instance *CIM_VirtualSystemMigrationCapabilities) GetPropertyDestinationHostFormatsSupported() (value []VirtualSystemMigrationCapabilities_DestinationHostFormatsSupported, err error) {
	retValue, err := instance.GetProperty("DestinationHostFormatsSupported")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, VirtualSystemMigrationCapabilities_DestinationHostFormatsSupported(valuetmp))
	}

	return
}

// SetSynchronousMethodsSupported sets the value of SynchronousMethodsSupported for the instance
func (instance *CIM_VirtualSystemMigrationCapabilities) SetPropertySynchronousMethodsSupported(value []VirtualSystemMigrationCapabilities_SynchronousMethodsSupported) (err error) {
	return instance.SetProperty("SynchronousMethodsSupported", (value))
}

// GetSynchronousMethodsSupported gets the value of SynchronousMethodsSupported for the instance
func (instance *CIM_VirtualSystemMigrationCapabilities) GetPropertySynchronousMethodsSupported() (value []VirtualSystemMigrationCapabilities_SynchronousMethodsSupported, err error) {
	retValue, err := instance.GetProperty("SynchronousMethodsSupported")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, VirtualSystemMigrationCapabilities_SynchronousMethodsSupported(valuetmp))
	}

	return
}
