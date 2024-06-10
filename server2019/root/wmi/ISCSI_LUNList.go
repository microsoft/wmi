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

// ISCSI_LUNList struct
type ISCSI_LUNList struct {
	*cim.WmiInstance

	//
	OSLUN uint32

	//
	Reserved uint32

	//
	TargetLUN uint64
}

func NewISCSI_LUNListEx1(instance *cim.WmiInstance) (newInstance *ISCSI_LUNList, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &ISCSI_LUNList{
		WmiInstance: tmp,
	}
	return
}

func NewISCSI_LUNListEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ISCSI_LUNList, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ISCSI_LUNList{
		WmiInstance: tmp,
	}
	return
}

// SetOSLUN sets the value of OSLUN for the instance
func (instance *ISCSI_LUNList) SetPropertyOSLUN(value uint32) (err error) {
	return instance.SetProperty("OSLUN", (value))
}

// GetOSLUN gets the value of OSLUN for the instance
func (instance *ISCSI_LUNList) GetPropertyOSLUN() (value uint32, err error) {
	retValue, err := instance.GetProperty("OSLUN")
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

// SetReserved sets the value of Reserved for the instance
func (instance *ISCSI_LUNList) SetPropertyReserved(value uint32) (err error) {
	return instance.SetProperty("Reserved", (value))
}

// GetReserved gets the value of Reserved for the instance
func (instance *ISCSI_LUNList) GetPropertyReserved() (value uint32, err error) {
	retValue, err := instance.GetProperty("Reserved")
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

// SetTargetLUN sets the value of TargetLUN for the instance
func (instance *ISCSI_LUNList) SetPropertyTargetLUN(value uint64) (err error) {
	return instance.SetProperty("TargetLUN", (value))
}

// GetTargetLUN gets the value of TargetLUN for the instance
func (instance *ISCSI_LUNList) GetPropertyTargetLUN() (value uint64, err error) {
	retValue, err := instance.GetProperty("TargetLUN")
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
