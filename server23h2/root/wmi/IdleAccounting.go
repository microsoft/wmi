// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// IdleAccounting struct
type IdleAccounting struct {
	*WMIEvent

	//
	Active bool

	//
	InstanceName string

	//
	ResetCount uint32

	//
	StartTime uint64

	//
	State []IdleStateAccounting

	//
	StateCount uint32

	//
	TotalTransitions uint32
}

func NewIdleAccountingEx1(instance *cim.WmiInstance) (newInstance *IdleAccounting, err error) {
	tmp, err := NewWMIEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &IdleAccounting{
		WMIEvent: tmp,
	}
	return
}

func NewIdleAccountingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *IdleAccounting, err error) {
	tmp, err := NewWMIEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &IdleAccounting{
		WMIEvent: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *IdleAccounting) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *IdleAccounting) GetPropertyActive() (value bool, err error) {
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
func (instance *IdleAccounting) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *IdleAccounting) GetPropertyInstanceName() (value string, err error) {
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

// SetResetCount sets the value of ResetCount for the instance
func (instance *IdleAccounting) SetPropertyResetCount(value uint32) (err error) {
	return instance.SetProperty("ResetCount", (value))
}

// GetResetCount gets the value of ResetCount for the instance
func (instance *IdleAccounting) GetPropertyResetCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ResetCount")
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

// SetStartTime sets the value of StartTime for the instance
func (instance *IdleAccounting) SetPropertyStartTime(value uint64) (err error) {
	return instance.SetProperty("StartTime", (value))
}

// GetStartTime gets the value of StartTime for the instance
func (instance *IdleAccounting) GetPropertyStartTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("StartTime")
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

// SetState sets the value of State for the instance
func (instance *IdleAccounting) SetPropertyState(value []IdleStateAccounting) (err error) {
	return instance.SetProperty("State", (value))
}

// GetState gets the value of State for the instance
func (instance *IdleAccounting) GetPropertyState() (value []IdleStateAccounting, err error) {
	retValue, err := instance.GetProperty("State")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(IdleStateAccounting)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " IdleStateAccounting is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, IdleStateAccounting(valuetmp))
	}

	return
}

// SetStateCount sets the value of StateCount for the instance
func (instance *IdleAccounting) SetPropertyStateCount(value uint32) (err error) {
	return instance.SetProperty("StateCount", (value))
}

// GetStateCount gets the value of StateCount for the instance
func (instance *IdleAccounting) GetPropertyStateCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("StateCount")
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

// SetTotalTransitions sets the value of TotalTransitions for the instance
func (instance *IdleAccounting) SetPropertyTotalTransitions(value uint32) (err error) {
	return instance.SetProperty("TotalTransitions", (value))
}

// GetTotalTransitions gets the value of TotalTransitions for the instance
func (instance *IdleAccounting) GetPropertyTotalTransitions() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalTransitions")
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
