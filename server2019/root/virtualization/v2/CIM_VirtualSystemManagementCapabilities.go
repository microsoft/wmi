// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// CIM_VirtualSystemManagementCapabilities struct
type CIM_VirtualSystemManagementCapabilities struct {
	CIM_EnabledLogicalElementCapabilities

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

// SetAsynchronousMethodsSupported sets the value of AsynchronousMethodsSupported for the instance
func (instance *CIM_VirtualSystemManagementCapabilities) SetPropertyAsynchronousMethodsSupported(value []VirtualSystemManagementCapabilities_AsynchronousMethodsSupported) (err error) {
	return instance.SetProperty("AsynchronousMethodsSupported", value)
}

// GetAsynchronousMethodsSupported gets the value of AsynchronousMethodsSupported for the instance
func (instance *CIM_VirtualSystemManagementCapabilities) GetPropertyAsynchronousMethodsSupported() (value []VirtualSystemManagementCapabilities_AsynchronousMethodsSupported, err error) {
	retValue, err := instance.GetProperty("AsynchronousMethodsSupported")
	if err != nil {
		return
	}
	value, ok := retValue.([]VirtualSystemManagementCapabilities_AsynchronousMethodsSupported)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIndicationsSupported sets the value of IndicationsSupported for the instance
func (instance *CIM_VirtualSystemManagementCapabilities) SetPropertyIndicationsSupported(value []VirtualSystemManagementCapabilities_IndicationsSupported) (err error) {
	return instance.SetProperty("IndicationsSupported", value)
}

// GetIndicationsSupported gets the value of IndicationsSupported for the instance
func (instance *CIM_VirtualSystemManagementCapabilities) GetPropertyIndicationsSupported() (value []VirtualSystemManagementCapabilities_IndicationsSupported, err error) {
	retValue, err := instance.GetProperty("IndicationsSupported")
	if err != nil {
		return
	}
	value, ok := retValue.([]VirtualSystemManagementCapabilities_IndicationsSupported)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSynchronousMethodsSupported sets the value of SynchronousMethodsSupported for the instance
func (instance *CIM_VirtualSystemManagementCapabilities) SetPropertySynchronousMethodsSupported(value []VirtualSystemManagementCapabilities_SynchronousMethodsSupported) (err error) {
	return instance.SetProperty("SynchronousMethodsSupported", value)
}

// GetSynchronousMethodsSupported gets the value of SynchronousMethodsSupported for the instance
func (instance *CIM_VirtualSystemManagementCapabilities) GetPropertySynchronousMethodsSupported() (value []VirtualSystemManagementCapabilities_SynchronousMethodsSupported, err error) {
	retValue, err := instance.GetProperty("SynchronousMethodsSupported")
	if err != nil {
		return
	}
	value, ok := retValue.([]VirtualSystemManagementCapabilities_SynchronousMethodsSupported)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualSystemTypesSupported sets the value of VirtualSystemTypesSupported for the instance
func (instance *CIM_VirtualSystemManagementCapabilities) SetPropertyVirtualSystemTypesSupported(value []string) (err error) {
	return instance.SetProperty("VirtualSystemTypesSupported", value)
}

// GetVirtualSystemTypesSupported gets the value of VirtualSystemTypesSupported for the instance
func (instance *CIM_VirtualSystemManagementCapabilities) GetPropertyVirtualSystemTypesSupported() (value []string, err error) {
	retValue, err := instance.GetProperty("VirtualSystemTypesSupported")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
