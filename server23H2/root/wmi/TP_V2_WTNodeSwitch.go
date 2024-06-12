// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// TP_V2_WTNodeSwitch struct
type TP_V2_WTNodeSwitch struct {
	*ThreadPoolTrace_V2

	//
	CurrentGroup uint16

	//
	CurrentNode uint32

	//
	CurrentWorkerCount uint32

	//
	NextGroup uint16

	//
	NextNode uint32

	//
	NextWorkerCount uint32

	//
	PoolId uint32
}

func NewTP_V2_WTNodeSwitchEx1(instance *cim.WmiInstance) (newInstance *TP_V2_WTNodeSwitch, err error) {
	tmp, err := NewThreadPoolTrace_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &TP_V2_WTNodeSwitch{
		ThreadPoolTrace_V2: tmp,
	}
	return
}

func NewTP_V2_WTNodeSwitchEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TP_V2_WTNodeSwitch, err error) {
	tmp, err := NewThreadPoolTrace_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TP_V2_WTNodeSwitch{
		ThreadPoolTrace_V2: tmp,
	}
	return
}

// SetCurrentGroup sets the value of CurrentGroup for the instance
func (instance *TP_V2_WTNodeSwitch) SetPropertyCurrentGroup(value uint16) (err error) {
	return instance.SetProperty("CurrentGroup", (value))
}

// GetCurrentGroup gets the value of CurrentGroup for the instance
func (instance *TP_V2_WTNodeSwitch) GetPropertyCurrentGroup() (value uint16, err error) {
	retValue, err := instance.GetProperty("CurrentGroup")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetCurrentNode sets the value of CurrentNode for the instance
func (instance *TP_V2_WTNodeSwitch) SetPropertyCurrentNode(value uint32) (err error) {
	return instance.SetProperty("CurrentNode", (value))
}

// GetCurrentNode gets the value of CurrentNode for the instance
func (instance *TP_V2_WTNodeSwitch) GetPropertyCurrentNode() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentNode")
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

// SetCurrentWorkerCount sets the value of CurrentWorkerCount for the instance
func (instance *TP_V2_WTNodeSwitch) SetPropertyCurrentWorkerCount(value uint32) (err error) {
	return instance.SetProperty("CurrentWorkerCount", (value))
}

// GetCurrentWorkerCount gets the value of CurrentWorkerCount for the instance
func (instance *TP_V2_WTNodeSwitch) GetPropertyCurrentWorkerCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentWorkerCount")
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

// SetNextGroup sets the value of NextGroup for the instance
func (instance *TP_V2_WTNodeSwitch) SetPropertyNextGroup(value uint16) (err error) {
	return instance.SetProperty("NextGroup", (value))
}

// GetNextGroup gets the value of NextGroup for the instance
func (instance *TP_V2_WTNodeSwitch) GetPropertyNextGroup() (value uint16, err error) {
	retValue, err := instance.GetProperty("NextGroup")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetNextNode sets the value of NextNode for the instance
func (instance *TP_V2_WTNodeSwitch) SetPropertyNextNode(value uint32) (err error) {
	return instance.SetProperty("NextNode", (value))
}

// GetNextNode gets the value of NextNode for the instance
func (instance *TP_V2_WTNodeSwitch) GetPropertyNextNode() (value uint32, err error) {
	retValue, err := instance.GetProperty("NextNode")
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

// SetNextWorkerCount sets the value of NextWorkerCount for the instance
func (instance *TP_V2_WTNodeSwitch) SetPropertyNextWorkerCount(value uint32) (err error) {
	return instance.SetProperty("NextWorkerCount", (value))
}

// GetNextWorkerCount gets the value of NextWorkerCount for the instance
func (instance *TP_V2_WTNodeSwitch) GetPropertyNextWorkerCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("NextWorkerCount")
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

// SetPoolId sets the value of PoolId for the instance
func (instance *TP_V2_WTNodeSwitch) SetPropertyPoolId(value uint32) (err error) {
	return instance.SetProperty("PoolId", (value))
}

// GetPoolId gets the value of PoolId for the instance
func (instance *TP_V2_WTNodeSwitch) GetPropertyPoolId() (value uint32, err error) {
	retValue, err := instance.GetProperty("PoolId")
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
