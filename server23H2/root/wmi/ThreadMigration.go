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

// ThreadMigration struct
type ThreadMigration struct {
	*Thread_V2

	//
	IdealProcessorAdjust bool

	//
	OldIdealProcessorIndex uint16

	//
	Priority uint8

	//
	SourceProcessorIndex uint16

	//
	TargetProcessorIndex uint16

	//
	ThreadId uint32
}

func NewThreadMigrationEx1(instance *cim.WmiInstance) (newInstance *ThreadMigration, err error) {
	tmp, err := NewThread_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &ThreadMigration{
		Thread_V2: tmp,
	}
	return
}

func NewThreadMigrationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ThreadMigration, err error) {
	tmp, err := NewThread_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ThreadMigration{
		Thread_V2: tmp,
	}
	return
}

// SetIdealProcessorAdjust sets the value of IdealProcessorAdjust for the instance
func (instance *ThreadMigration) SetPropertyIdealProcessorAdjust(value bool) (err error) {
	return instance.SetProperty("IdealProcessorAdjust", (value))
}

// GetIdealProcessorAdjust gets the value of IdealProcessorAdjust for the instance
func (instance *ThreadMigration) GetPropertyIdealProcessorAdjust() (value bool, err error) {
	retValue, err := instance.GetProperty("IdealProcessorAdjust")
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

// SetOldIdealProcessorIndex sets the value of OldIdealProcessorIndex for the instance
func (instance *ThreadMigration) SetPropertyOldIdealProcessorIndex(value uint16) (err error) {
	return instance.SetProperty("OldIdealProcessorIndex", (value))
}

// GetOldIdealProcessorIndex gets the value of OldIdealProcessorIndex for the instance
func (instance *ThreadMigration) GetPropertyOldIdealProcessorIndex() (value uint16, err error) {
	retValue, err := instance.GetProperty("OldIdealProcessorIndex")
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

// SetPriority sets the value of Priority for the instance
func (instance *ThreadMigration) SetPropertyPriority(value uint8) (err error) {
	return instance.SetProperty("Priority", (value))
}

// GetPriority gets the value of Priority for the instance
func (instance *ThreadMigration) GetPropertyPriority() (value uint8, err error) {
	retValue, err := instance.GetProperty("Priority")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetSourceProcessorIndex sets the value of SourceProcessorIndex for the instance
func (instance *ThreadMigration) SetPropertySourceProcessorIndex(value uint16) (err error) {
	return instance.SetProperty("SourceProcessorIndex", (value))
}

// GetSourceProcessorIndex gets the value of SourceProcessorIndex for the instance
func (instance *ThreadMigration) GetPropertySourceProcessorIndex() (value uint16, err error) {
	retValue, err := instance.GetProperty("SourceProcessorIndex")
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

// SetTargetProcessorIndex sets the value of TargetProcessorIndex for the instance
func (instance *ThreadMigration) SetPropertyTargetProcessorIndex(value uint16) (err error) {
	return instance.SetProperty("TargetProcessorIndex", (value))
}

// GetTargetProcessorIndex gets the value of TargetProcessorIndex for the instance
func (instance *ThreadMigration) GetPropertyTargetProcessorIndex() (value uint16, err error) {
	retValue, err := instance.GetProperty("TargetProcessorIndex")
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

// SetThreadId sets the value of ThreadId for the instance
func (instance *ThreadMigration) SetPropertyThreadId(value uint32) (err error) {
	return instance.SetProperty("ThreadId", (value))
}

// GetThreadId gets the value of ThreadId for the instance
func (instance *ThreadMigration) GetPropertyThreadId() (value uint32, err error) {
	retValue, err := instance.GetProperty("ThreadId")
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
