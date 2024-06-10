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

// MSiSCSI_RedirectPortalInfoClass struct
type MSiSCSI_RedirectPortalInfoClass struct {
	*cim.WmiInstance

	//
	Active bool

	//
	InstanceName string

	// Variable length array of ISCSI_RedirectSessionInfo. SessionCount specifies the number of elements in the array. NOTE: this is a variable length array.
	RedirectSessionList []ISCSI_RedirectSessionInfo

	// Number of elements in RedirectSessionInfo array
	SessionCount uint32

	// Id that is globally unique for all instances of iSCSI initiators.
	UniqueAdapterId uint64
}

func NewMSiSCSI_RedirectPortalInfoClassEx1(instance *cim.WmiInstance) (newInstance *MSiSCSI_RedirectPortalInfoClass, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_RedirectPortalInfoClass{
		WmiInstance: tmp,
	}
	return
}

func NewMSiSCSI_RedirectPortalInfoClassEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSiSCSI_RedirectPortalInfoClass, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_RedirectPortalInfoClass{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSiSCSI_RedirectPortalInfoClass) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSiSCSI_RedirectPortalInfoClass) GetPropertyActive() (value bool, err error) {
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
func (instance *MSiSCSI_RedirectPortalInfoClass) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSiSCSI_RedirectPortalInfoClass) GetPropertyInstanceName() (value string, err error) {
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

// SetRedirectSessionList sets the value of RedirectSessionList for the instance
func (instance *MSiSCSI_RedirectPortalInfoClass) SetPropertyRedirectSessionList(value []ISCSI_RedirectSessionInfo) (err error) {
	return instance.SetProperty("RedirectSessionList", (value))
}

// GetRedirectSessionList gets the value of RedirectSessionList for the instance
func (instance *MSiSCSI_RedirectPortalInfoClass) GetPropertyRedirectSessionList() (value []ISCSI_RedirectSessionInfo, err error) {
	retValue, err := instance.GetProperty("RedirectSessionList")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(ISCSI_RedirectSessionInfo)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " ISCSI_RedirectSessionInfo is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, ISCSI_RedirectSessionInfo(valuetmp))
	}

	return
}

// SetSessionCount sets the value of SessionCount for the instance
func (instance *MSiSCSI_RedirectPortalInfoClass) SetPropertySessionCount(value uint32) (err error) {
	return instance.SetProperty("SessionCount", (value))
}

// GetSessionCount gets the value of SessionCount for the instance
func (instance *MSiSCSI_RedirectPortalInfoClass) GetPropertySessionCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("SessionCount")
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

// SetUniqueAdapterId sets the value of UniqueAdapterId for the instance
func (instance *MSiSCSI_RedirectPortalInfoClass) SetPropertyUniqueAdapterId(value uint64) (err error) {
	return instance.SetProperty("UniqueAdapterId", (value))
}

// GetUniqueAdapterId gets the value of UniqueAdapterId for the instance
func (instance *MSiSCSI_RedirectPortalInfoClass) GetPropertyUniqueAdapterId() (value uint64, err error) {
	retValue, err := instance.GetProperty("UniqueAdapterId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}
