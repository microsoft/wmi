// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ISCSI_TargetMapping struct
type ISCSI_TargetMapping struct {
	*cim.WmiInstance

	//
	FromPersistentLogin bool

	//
	LUNCount uint32

	//
	LUNList []ISCSI_LUNList

	//
	OSBus uint32

	//
	OSTarget uint32

	//
	Reserved uint64

	//
	TargetName string

	//
	UniqueSessionId uint64
}

func NewISCSI_TargetMappingEx1(instance *cim.WmiInstance) (newInstance *ISCSI_TargetMapping, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &ISCSI_TargetMapping{
		WmiInstance: tmp,
	}
	return
}

func NewISCSI_TargetMappingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ISCSI_TargetMapping, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ISCSI_TargetMapping{
		WmiInstance: tmp,
	}
	return
}

// SetFromPersistentLogin sets the value of FromPersistentLogin for the instance
func (instance *ISCSI_TargetMapping) SetPropertyFromPersistentLogin(value bool) (err error) {
	return instance.SetProperty("FromPersistentLogin", (value))
}

// GetFromPersistentLogin gets the value of FromPersistentLogin for the instance
func (instance *ISCSI_TargetMapping) GetPropertyFromPersistentLogin() (value bool, err error) {
	retValue, err := instance.GetProperty("FromPersistentLogin")
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

// SetLUNCount sets the value of LUNCount for the instance
func (instance *ISCSI_TargetMapping) SetPropertyLUNCount(value uint32) (err error) {
	return instance.SetProperty("LUNCount", (value))
}

// GetLUNCount gets the value of LUNCount for the instance
func (instance *ISCSI_TargetMapping) GetPropertyLUNCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("LUNCount")
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

// SetLUNList sets the value of LUNList for the instance
func (instance *ISCSI_TargetMapping) SetPropertyLUNList(value []ISCSI_LUNList) (err error) {
	return instance.SetProperty("LUNList", (value))
}

// GetLUNList gets the value of LUNList for the instance
func (instance *ISCSI_TargetMapping) GetPropertyLUNList() (value []ISCSI_LUNList, err error) {
	retValue, err := instance.GetProperty("LUNList")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(ISCSI_LUNList)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " ISCSI_LUNList is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, ISCSI_LUNList(valuetmp))
	}

	return
}

// SetOSBus sets the value of OSBus for the instance
func (instance *ISCSI_TargetMapping) SetPropertyOSBus(value uint32) (err error) {
	return instance.SetProperty("OSBus", (value))
}

// GetOSBus gets the value of OSBus for the instance
func (instance *ISCSI_TargetMapping) GetPropertyOSBus() (value uint32, err error) {
	retValue, err := instance.GetProperty("OSBus")
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

// SetOSTarget sets the value of OSTarget for the instance
func (instance *ISCSI_TargetMapping) SetPropertyOSTarget(value uint32) (err error) {
	return instance.SetProperty("OSTarget", (value))
}

// GetOSTarget gets the value of OSTarget for the instance
func (instance *ISCSI_TargetMapping) GetPropertyOSTarget() (value uint32, err error) {
	retValue, err := instance.GetProperty("OSTarget")
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
func (instance *ISCSI_TargetMapping) SetPropertyReserved(value uint64) (err error) {
	return instance.SetProperty("Reserved", (value))
}

// GetReserved gets the value of Reserved for the instance
func (instance *ISCSI_TargetMapping) GetPropertyReserved() (value uint64, err error) {
	retValue, err := instance.GetProperty("Reserved")
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

// SetTargetName sets the value of TargetName for the instance
func (instance *ISCSI_TargetMapping) SetPropertyTargetName(value string) (err error) {
	return instance.SetProperty("TargetName", (value))
}

// GetTargetName gets the value of TargetName for the instance
func (instance *ISCSI_TargetMapping) GetPropertyTargetName() (value string, err error) {
	retValue, err := instance.GetProperty("TargetName")
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

// SetUniqueSessionId sets the value of UniqueSessionId for the instance
func (instance *ISCSI_TargetMapping) SetPropertyUniqueSessionId(value uint64) (err error) {
	return instance.SetProperty("UniqueSessionId", (value))
}

// GetUniqueSessionId gets the value of UniqueSessionId for the instance
func (instance *ISCSI_TargetMapping) GetPropertyUniqueSessionId() (value uint64, err error) {
	retValue, err := instance.GetProperty("UniqueSessionId")
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
