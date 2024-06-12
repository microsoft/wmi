// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_VirtualSystemManagementCapabilities struct
type CIM_VirtualSystemManagementCapabilities struct {
	*CIM_EnabledLogicalElementCapabilities

	// Enumeration of method identifiers each identifying a method of class CIM_VirtualSystemManagementService that is supported asynchronously by the implementation.
	AsynchronousMethodsSupported []VirtualSystemManagementCapabilities_AsynchronousMethodsSupported

	// Enumeration of indication identifiers each identifying an indication that is supported by the implementation.
	///VirtualSystemStateChangeIndicationsSupported indicates whether or not the implementation supports notification on state changes of CIM_ComputerSystem instances representing virtual systems.
	///VirtualResourceStateChangeIndicationsSupported indicates whether or not the implementation supports notification on state changes of CIM_LogicalDevice instances representing resources of virtual systems.
	///ConcreteJobStateChangeIndicationsSupported indicates whether or not the implementation supports notification on state changes of CIM_ConcreteJob instances.
	IndicationsSupported []VirtualSystemManagementCapabilities_IndicationsSupported

	// Enumeration of method identifiers each identifying a method of class CIM_VirtualSystemManagementService that is supported synchronously by the implementation.
	SynchronousMethodsSupported []VirtualSystemManagementCapabilities_SynchronousMethodsSupported

	// Enumeration of strings each designating a type of virtual system that the implementation supports.
	///The value of each non-NULL array element shall conform to the format defined for the CIM_VirtualSystemSettingData.VirtualSystemType property.
	VirtualSystemTypesSupported []string
}

func NewCIM_VirtualSystemManagementCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *CIM_VirtualSystemManagementCapabilities, err error) {
	tmp, err := NewCIM_EnabledLogicalElementCapabilitiesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_VirtualSystemManagementCapabilities{
		CIM_EnabledLogicalElementCapabilities: tmp,
	}
	return
}

func NewCIM_VirtualSystemManagementCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_VirtualSystemManagementCapabilities, err error) {
	tmp, err := NewCIM_EnabledLogicalElementCapabilitiesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_VirtualSystemManagementCapabilities{
		CIM_EnabledLogicalElementCapabilities: tmp,
	}
	return
}

// SetAsynchronousMethodsSupported sets the value of AsynchronousMethodsSupported for the instance
func (instance *CIM_VirtualSystemManagementCapabilities) SetPropertyAsynchronousMethodsSupported(value []VirtualSystemManagementCapabilities_AsynchronousMethodsSupported) (err error) {
	return instance.SetProperty("AsynchronousMethodsSupported", (value))
}

// GetAsynchronousMethodsSupported gets the value of AsynchronousMethodsSupported for the instance
func (instance *CIM_VirtualSystemManagementCapabilities) GetPropertyAsynchronousMethodsSupported() (value []VirtualSystemManagementCapabilities_AsynchronousMethodsSupported, err error) {
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
		value = append(value, VirtualSystemManagementCapabilities_AsynchronousMethodsSupported(valuetmp))
	}

	return
}

// SetIndicationsSupported sets the value of IndicationsSupported for the instance
func (instance *CIM_VirtualSystemManagementCapabilities) SetPropertyIndicationsSupported(value []VirtualSystemManagementCapabilities_IndicationsSupported) (err error) {
	return instance.SetProperty("IndicationsSupported", (value))
}

// GetIndicationsSupported gets the value of IndicationsSupported for the instance
func (instance *CIM_VirtualSystemManagementCapabilities) GetPropertyIndicationsSupported() (value []VirtualSystemManagementCapabilities_IndicationsSupported, err error) {
	retValue, err := instance.GetProperty("IndicationsSupported")
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
		value = append(value, VirtualSystemManagementCapabilities_IndicationsSupported(valuetmp))
	}

	return
}

// SetSynchronousMethodsSupported sets the value of SynchronousMethodsSupported for the instance
func (instance *CIM_VirtualSystemManagementCapabilities) SetPropertySynchronousMethodsSupported(value []VirtualSystemManagementCapabilities_SynchronousMethodsSupported) (err error) {
	return instance.SetProperty("SynchronousMethodsSupported", (value))
}

// GetSynchronousMethodsSupported gets the value of SynchronousMethodsSupported for the instance
func (instance *CIM_VirtualSystemManagementCapabilities) GetPropertySynchronousMethodsSupported() (value []VirtualSystemManagementCapabilities_SynchronousMethodsSupported, err error) {
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
		value = append(value, VirtualSystemManagementCapabilities_SynchronousMethodsSupported(valuetmp))
	}

	return
}

// SetVirtualSystemTypesSupported sets the value of VirtualSystemTypesSupported for the instance
func (instance *CIM_VirtualSystemManagementCapabilities) SetPropertyVirtualSystemTypesSupported(value []string) (err error) {
	return instance.SetProperty("VirtualSystemTypesSupported", (value))
}

// GetVirtualSystemTypesSupported gets the value of VirtualSystemTypesSupported for the instance
func (instance *CIM_VirtualSystemManagementCapabilities) GetPropertyVirtualSystemTypesSupported() (value []string, err error) {
	retValue, err := instance.GetProperty("VirtualSystemTypesSupported")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}
