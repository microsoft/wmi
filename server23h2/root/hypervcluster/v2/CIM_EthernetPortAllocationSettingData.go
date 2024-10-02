// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_EthernetPortAllocationSettingData struct
type CIM_EthernetPortAllocationSettingData struct {
	*CIM_ResourceAllocationSettingData

	// The desired VLAN mode that is requested for use. This property is usedto set the initial OperationalEndpointMode property value in theinstance of CIM_VLANEndpoint associated with the targeted Ethernet Port.Refer to the description for the property OperationalEndpointMode inCIM_VLANEndpoint for a description of the values
	DesiredVLANEndpointMode EthernetPortAllocationSettingData_DesiredVLANEndpointMode

	// A string describing the type of VLAN endpoint model that is supported by this VLANEndpoint, when the value of the mode property is set to 1 (i.e., "Other"). This property should be set to NULL when the mode property is any value other than 1.
	OtherEndpointMode string
}

func NewCIM_EthernetPortAllocationSettingDataEx1(instance *cim.WmiInstance) (newInstance *CIM_EthernetPortAllocationSettingData, err error) {
	tmp, err := NewCIM_ResourceAllocationSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_EthernetPortAllocationSettingData{
		CIM_ResourceAllocationSettingData: tmp,
	}
	return
}

func NewCIM_EthernetPortAllocationSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_EthernetPortAllocationSettingData, err error) {
	tmp, err := NewCIM_ResourceAllocationSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_EthernetPortAllocationSettingData{
		CIM_ResourceAllocationSettingData: tmp,
	}
	return
}

// SetDesiredVLANEndpointMode sets the value of DesiredVLANEndpointMode for the instance
func (instance *CIM_EthernetPortAllocationSettingData) SetPropertyDesiredVLANEndpointMode(value EthernetPortAllocationSettingData_DesiredVLANEndpointMode) (err error) {
	return instance.SetProperty("DesiredVLANEndpointMode", (value))
}

// GetDesiredVLANEndpointMode gets the value of DesiredVLANEndpointMode for the instance
func (instance *CIM_EthernetPortAllocationSettingData) GetPropertyDesiredVLANEndpointMode() (value EthernetPortAllocationSettingData_DesiredVLANEndpointMode, err error) {
	retValue, err := instance.GetProperty("DesiredVLANEndpointMode")
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

	value = EthernetPortAllocationSettingData_DesiredVLANEndpointMode(valuetmp)

	return
}

// SetOtherEndpointMode sets the value of OtherEndpointMode for the instance
func (instance *CIM_EthernetPortAllocationSettingData) SetPropertyOtherEndpointMode(value string) (err error) {
	return instance.SetProperty("OtherEndpointMode", (value))
}

// GetOtherEndpointMode gets the value of OtherEndpointMode for the instance
func (instance *CIM_EthernetPortAllocationSettingData) GetPropertyOtherEndpointMode() (value string, err error) {
	retValue, err := instance.GetProperty("OtherEndpointMode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}
