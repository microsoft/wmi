// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.StandardCimv2
//
// ////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_IPAssignmentSettingData struct
type CIM_IPAssignmentSettingData struct {
	*CIM_SettingData

	// 689
	AddressOrigin IPAssignmentSettingData_AddressOrigin

	// 693
	ProtocolIFType IPAssignmentSettingData_ProtocolIFType
}

func NewCIM_IPAssignmentSettingDataEx1(instance *cim.WmiInstance) (newInstance *CIM_IPAssignmentSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_IPAssignmentSettingData{
		CIM_SettingData: tmp,
	}
	return
}

func NewCIM_IPAssignmentSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_IPAssignmentSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_IPAssignmentSettingData{
		CIM_SettingData: tmp,
	}
	return
}

// SetAddressOrigin sets the value of AddressOrigin for the instance
func (instance *CIM_IPAssignmentSettingData) SetPropertyAddressOrigin(value IPAssignmentSettingData_AddressOrigin) (err error) {
	return instance.SetProperty("AddressOrigin", (value))
}

// GetAddressOrigin gets the value of AddressOrigin for the instance
func (instance *CIM_IPAssignmentSettingData) GetPropertyAddressOrigin() (value IPAssignmentSettingData_AddressOrigin, err error) {
	retValue, err := instance.GetProperty("AddressOrigin")
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

	value = IPAssignmentSettingData_AddressOrigin(valuetmp)

	return
}

// SetProtocolIFType sets the value of ProtocolIFType for the instance
func (instance *CIM_IPAssignmentSettingData) SetPropertyProtocolIFType(value IPAssignmentSettingData_ProtocolIFType) (err error) {
	return instance.SetProperty("ProtocolIFType", (value))
}

// GetProtocolIFType gets the value of ProtocolIFType for the instance
func (instance *CIM_IPAssignmentSettingData) GetPropertyProtocolIFType() (value IPAssignmentSettingData_ProtocolIFType, err error) {
	retValue, err := instance.GetProperty("ProtocolIFType")
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

	value = IPAssignmentSettingData_ProtocolIFType(valuetmp)

	return
}
