// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_ProtocolControllerForUnit struct
type CIM_ProtocolControllerForUnit struct {
	*CIM_ProtocolControllerForDevice

	// The access rights granted to the referenced logical unit as exposed through referenced ProtocolController. The 'No Access' value is used in implementations where the DeviceNumber is reserved, but no access is granted.
	///If the instrumentation exposes PrivilegeManagementService, this property MUST be synchronized with the Activities property of any Privilege instances associated with StorageHardwareIDs associated to the referenced ProtocolController and the referenced LogicalDevice. In particular, when this property is 'Read Write', Privilege.Activities MUST include entries for 'Read' and 'Write'. When this property is 'Read-Only', Privilege.Activities MUST include an entry for 'Read'. The corresponding entries for Privilege.ActivityQualifiers MUST be 'CDB=*' and the corresponding entries for Privilege.QualifierFormat MUST be 'SCSI Command'.
	DeviceAccess ProtocolControllerForUnit_DeviceAccess
}

func NewCIM_ProtocolControllerForUnitEx1(instance *cim.WmiInstance) (newInstance *CIM_ProtocolControllerForUnit, err error) {
	tmp, err := NewCIM_ProtocolControllerForDeviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_ProtocolControllerForUnit{
		CIM_ProtocolControllerForDevice: tmp,
	}
	return
}

func NewCIM_ProtocolControllerForUnitEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ProtocolControllerForUnit, err error) {
	tmp, err := NewCIM_ProtocolControllerForDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ProtocolControllerForUnit{
		CIM_ProtocolControllerForDevice: tmp,
	}
	return
}

// SetDeviceAccess sets the value of DeviceAccess for the instance
func (instance *CIM_ProtocolControllerForUnit) SetPropertyDeviceAccess(value ProtocolControllerForUnit_DeviceAccess) (err error) {
	return instance.SetProperty("DeviceAccess", (value))
}

// GetDeviceAccess gets the value of DeviceAccess for the instance
func (instance *CIM_ProtocolControllerForUnit) GetPropertyDeviceAccess() (value ProtocolControllerForUnit_DeviceAccess, err error) {
	retValue, err := instance.GetProperty("DeviceAccess")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = ProtocolControllerForUnit_DeviceAccess(valuetmp)

	return
}
