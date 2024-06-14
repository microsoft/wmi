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

// HeapExpand struct
type HeapExpand struct {
	*HeapTrace

	//
	AllocatedSpace interface{}

	//
	CommitAddress uint32

	//
	CommittedSize interface{}

	//
	CommittedSpace interface{}

	//
	FreeSpace interface{}

	//
	HeapHandle uint32

	//
	NoOfUCRs uint32

	//
	ReservedSpace interface{}
}

func NewHeapExpandEx1(instance *cim.WmiInstance) (newInstance *HeapExpand, err error) {
	tmp, err := NewHeapTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &HeapExpand{
		HeapTrace: tmp,
	}
	return
}

func NewHeapExpandEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HeapExpand, err error) {
	tmp, err := NewHeapTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HeapExpand{
		HeapTrace: tmp,
	}
	return
}

// SetAllocatedSpace sets the value of AllocatedSpace for the instance
func (instance *HeapExpand) SetPropertyAllocatedSpace(value interface{}) (err error) {
	return instance.SetProperty("AllocatedSpace", (value))
}

// GetAllocatedSpace gets the value of AllocatedSpace for the instance
func (instance *HeapExpand) GetPropertyAllocatedSpace() (value interface{}, err error) {
	retValue, err := instance.GetProperty("AllocatedSpace")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}

// SetCommitAddress sets the value of CommitAddress for the instance
func (instance *HeapExpand) SetPropertyCommitAddress(value uint32) (err error) {
	return instance.SetProperty("CommitAddress", (value))
}

// GetCommitAddress gets the value of CommitAddress for the instance
func (instance *HeapExpand) GetPropertyCommitAddress() (value uint32, err error) {
	retValue, err := instance.GetProperty("CommitAddress")
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

// SetCommittedSize sets the value of CommittedSize for the instance
func (instance *HeapExpand) SetPropertyCommittedSize(value interface{}) (err error) {
	return instance.SetProperty("CommittedSize", (value))
}

// GetCommittedSize gets the value of CommittedSize for the instance
func (instance *HeapExpand) GetPropertyCommittedSize() (value interface{}, err error) {
	retValue, err := instance.GetProperty("CommittedSize")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}

// SetCommittedSpace sets the value of CommittedSpace for the instance
func (instance *HeapExpand) SetPropertyCommittedSpace(value interface{}) (err error) {
	return instance.SetProperty("CommittedSpace", (value))
}

// GetCommittedSpace gets the value of CommittedSpace for the instance
func (instance *HeapExpand) GetPropertyCommittedSpace() (value interface{}, err error) {
	retValue, err := instance.GetProperty("CommittedSpace")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}

// SetFreeSpace sets the value of FreeSpace for the instance
func (instance *HeapExpand) SetPropertyFreeSpace(value interface{}) (err error) {
	return instance.SetProperty("FreeSpace", (value))
}

// GetFreeSpace gets the value of FreeSpace for the instance
func (instance *HeapExpand) GetPropertyFreeSpace() (value interface{}, err error) {
	retValue, err := instance.GetProperty("FreeSpace")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}

// SetHeapHandle sets the value of HeapHandle for the instance
func (instance *HeapExpand) SetPropertyHeapHandle(value uint32) (err error) {
	return instance.SetProperty("HeapHandle", (value))
}

// GetHeapHandle gets the value of HeapHandle for the instance
func (instance *HeapExpand) GetPropertyHeapHandle() (value uint32, err error) {
	retValue, err := instance.GetProperty("HeapHandle")
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

// SetNoOfUCRs sets the value of NoOfUCRs for the instance
func (instance *HeapExpand) SetPropertyNoOfUCRs(value uint32) (err error) {
	return instance.SetProperty("NoOfUCRs", (value))
}

// GetNoOfUCRs gets the value of NoOfUCRs for the instance
func (instance *HeapExpand) GetPropertyNoOfUCRs() (value uint32, err error) {
	retValue, err := instance.GetProperty("NoOfUCRs")
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

// SetReservedSpace sets the value of ReservedSpace for the instance
func (instance *HeapExpand) SetPropertyReservedSpace(value interface{}) (err error) {
	return instance.SetProperty("ReservedSpace", (value))
}

// GetReservedSpace gets the value of ReservedSpace for the instance
func (instance *HeapExpand) GetPropertyReservedSpace() (value interface{}, err error) {
	retValue, err := instance.GetProperty("ReservedSpace")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}
