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

// Msvm_VirtualEthernetSwitchManagementCapabilities struct
type Msvm_VirtualEthernetSwitchManagementCapabilities struct {
	*CIM_VirtualSystemManagementCapabilities

	// A boolean value which indicates whether IOV is supported by the platform.If the value is TRUE, then IOV is supported by the platform and IOVSupportReasons will be empty. Otherwise the IOVSupportReasons property will have the reasons why IOV cannot be supported.
	IOVSupport bool

	// An array of strings that indicates the possible reasons why IOV is not supported. If the value of IOVSupport is TRUE this array will be empty.
	IOVSupportReasons []string
}

func NewMsvm_VirtualEthernetSwitchManagementCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *Msvm_VirtualEthernetSwitchManagementCapabilities, err error) {
	tmp, err := NewCIM_VirtualSystemManagementCapabilitiesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualEthernetSwitchManagementCapabilities{
		CIM_VirtualSystemManagementCapabilities: tmp,
	}
	return
}

func NewMsvm_VirtualEthernetSwitchManagementCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VirtualEthernetSwitchManagementCapabilities, err error) {
	tmp, err := NewCIM_VirtualSystemManagementCapabilitiesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualEthernetSwitchManagementCapabilities{
		CIM_VirtualSystemManagementCapabilities: tmp,
	}
	return
}

// SetIOVSupport sets the value of IOVSupport for the instance
func (instance *Msvm_VirtualEthernetSwitchManagementCapabilities) SetPropertyIOVSupport(value bool) (err error) {
	return instance.SetProperty("IOVSupport", (value))
}

// GetIOVSupport gets the value of IOVSupport for the instance
func (instance *Msvm_VirtualEthernetSwitchManagementCapabilities) GetPropertyIOVSupport() (value bool, err error) {
	retValue, err := instance.GetProperty("IOVSupport")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetIOVSupportReasons sets the value of IOVSupportReasons for the instance
func (instance *Msvm_VirtualEthernetSwitchManagementCapabilities) SetPropertyIOVSupportReasons(value []string) (err error) {
	return instance.SetProperty("IOVSupportReasons", (value))
}

// GetIOVSupportReasons gets the value of IOVSupportReasons for the instance
func (instance *Msvm_VirtualEthernetSwitchManagementCapabilities) GetPropertyIOVSupportReasons() (value []string, err error) {
	retValue, err := instance.GetProperty("IOVSupportReasons")
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
func (instance *Msvm_VirtualEthernetSwitchManagementCapabilities) GetRelatedVirtualEthernetSwitchManagementService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualEthernetSwitchManagementService")
}
