// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSiSCSI_PortalInfoClass struct
type MSiSCSI_PortalInfoClass struct {
	*cim.WmiInstance

	//
	Active bool

	//
	InstanceName string

	// Number of elements in iScsiPortalInfo array
	PortalInfoCount uint32

	// Variable length array of iScsiPortalInfo.  PortalInfoCount specifies the number of elements in the array. NOTE: this is a variable length array.
	PortalInformation []ISCSI_PortalInfo
}

func NewMSiSCSI_PortalInfoClassEx1(instance *cim.WmiInstance) (newInstance *MSiSCSI_PortalInfoClass, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_PortalInfoClass{
		WmiInstance: tmp,
	}
	return
}

func NewMSiSCSI_PortalInfoClassEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSiSCSI_PortalInfoClass, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_PortalInfoClass{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSiSCSI_PortalInfoClass) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSiSCSI_PortalInfoClass) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSiSCSI_PortalInfoClass) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSiSCSI_PortalInfoClass) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
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

// SetPortalInfoCount sets the value of PortalInfoCount for the instance
func (instance *MSiSCSI_PortalInfoClass) SetPropertyPortalInfoCount(value uint32) (err error) {
	return instance.SetProperty("PortalInfoCount", (value))
}

// GetPortalInfoCount gets the value of PortalInfoCount for the instance
func (instance *MSiSCSI_PortalInfoClass) GetPropertyPortalInfoCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("PortalInfoCount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetPortalInformation sets the value of PortalInformation for the instance
func (instance *MSiSCSI_PortalInfoClass) SetPropertyPortalInformation(value []ISCSI_PortalInfo) (err error) {
	return instance.SetProperty("PortalInformation", (value))
}

// GetPortalInformation gets the value of PortalInformation for the instance
func (instance *MSiSCSI_PortalInfoClass) GetPropertyPortalInformation() (value []ISCSI_PortalInfo, err error) {
	retValue, err := instance.GetProperty("PortalInformation")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(ISCSI_PortalInfo)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " ISCSI_PortalInfo is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, ISCSI_PortalInfo(valuetmp))
	}

	return
}
