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

// CIM_ProtocolControllerForDevice struct
type CIM_ProtocolControllerForDevice struct {
	*CIM_Dependency

	// This property describes the priority given to accesses of the device through this Controller. The highest priority path will have the lowest value for this parameter.
	AccessPriority uint16

	// The AccessState property describes the accessibility of the LogicalDevice through the ProtocolController.
	///Unknown (0) indicates the instrumentation does not know whether access is or is not functioning.
	///Active (2) indicates normal access.
	///Inactive (3) indicates the instrumentation knows this path is not active, and one of the other values (below) does not apply.
	///Replication in Progress (4) indicates that the path is temporarily inactive due to a replication activity.
	///Mapping Inconsistency (5) indicates the instrumentation has detected that this path is inactive due to an inconsistency in the DeviceNumber/DeviceAccess configuration.
	AccessState ProtocolControllerForDevice_AccessState

	// Address of the associated Device in the context of the Antecedent Controller.
	DeviceNumber string
}

func NewCIM_ProtocolControllerForDeviceEx1(instance *cim.WmiInstance) (newInstance *CIM_ProtocolControllerForDevice, err error) {
	tmp, err := NewCIM_DependencyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_ProtocolControllerForDevice{
		CIM_Dependency: tmp,
	}
	return
}

func NewCIM_ProtocolControllerForDeviceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ProtocolControllerForDevice, err error) {
	tmp, err := NewCIM_DependencyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ProtocolControllerForDevice{
		CIM_Dependency: tmp,
	}
	return
}

// SetAccessPriority sets the value of AccessPriority for the instance
func (instance *CIM_ProtocolControllerForDevice) SetPropertyAccessPriority(value uint16) (err error) {
	return instance.SetProperty("AccessPriority", value)
}

// GetAccessPriority gets the value of AccessPriority for the instance
func (instance *CIM_ProtocolControllerForDevice) GetPropertyAccessPriority() (value uint16, err error) {
	retValue, err := instance.GetProperty("AccessPriority")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAccessState sets the value of AccessState for the instance
func (instance *CIM_ProtocolControllerForDevice) SetPropertyAccessState(value ProtocolControllerForDevice_AccessState) (err error) {
	return instance.SetProperty("AccessState", value)
}

// GetAccessState gets the value of AccessState for the instance
func (instance *CIM_ProtocolControllerForDevice) GetPropertyAccessState() (value ProtocolControllerForDevice_AccessState, err error) {
	retValue, err := instance.GetProperty("AccessState")
	if err != nil {
		return
	}
	value, ok := retValue.(ProtocolControllerForDevice_AccessState)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDeviceNumber sets the value of DeviceNumber for the instance
func (instance *CIM_ProtocolControllerForDevice) SetPropertyDeviceNumber(value string) (err error) {
	return instance.SetProperty("DeviceNumber", value)
}

// GetDeviceNumber gets the value of DeviceNumber for the instance
func (instance *CIM_ProtocolControllerForDevice) GetPropertyDeviceNumber() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
