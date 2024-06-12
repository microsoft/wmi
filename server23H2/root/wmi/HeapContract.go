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

// HeapContract struct
type HeapContract struct {
	*HeapTrace

	//
	AllocatedSpace interface{}

	//
	CommittedSpace interface{}

	//
	DeCommitAddress uint32

	//
	DeCommittedSize interface{}

	//
	FreeSpace interface{}

	//
	HeapHandle uint32

	//
	NoOfUCRs uint32

	//
	ReservedSpace interface{}
}

func NewHeapContractEx1(instance *cim.WmiInstance) (newInstance *HeapContract, err error) {
	tmp, err := NewHeapTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &HeapContract{
		HeapTrace: tmp,
	}
	return
}

func NewHeapContractEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HeapContract, err error) {
	tmp, err := NewHeapTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HeapContract{
		HeapTrace: tmp,
	}
	return
}

// SetAllocatedSpace sets the value of AllocatedSpace for the instance
func (instance *HeapContract) SetPropertyAllocatedSpace(value interface{}) (err error) {
	return instance.SetProperty("AllocatedSpace", (value))
}

// GetAllocatedSpace gets the value of AllocatedSpace for the instance
func (instance *HeapContract) GetPropertyAllocatedSpace() (value interface{}, err error) {
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

// SetCommittedSpace sets the value of CommittedSpace for the instance
func (instance *HeapContract) SetPropertyCommittedSpace(value interface{}) (err error) {
	return instance.SetProperty("CommittedSpace", (value))
}

// GetCommittedSpace gets the value of CommittedSpace for the instance
func (instance *HeapContract) GetPropertyCommittedSpace() (value interface{}, err error) {
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

// SetDeCommitAddress sets the value of DeCommitAddress for the instance
func (instance *HeapContract) SetPropertyDeCommitAddress(value uint32) (err error) {
	return instance.SetProperty("DeCommitAddress", (value))
}

// GetDeCommitAddress gets the value of DeCommitAddress for the instance
func (instance *HeapContract) GetPropertyDeCommitAddress() (value uint32, err error) {
	retValue, err := instance.GetProperty("DeCommitAddress")
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

// SetDeCommittedSize sets the value of DeCommittedSize for the instance
func (instance *HeapContract) SetPropertyDeCommittedSize(value interface{}) (err error) {
	return instance.SetProperty("DeCommittedSize", (value))
}

// GetDeCommittedSize gets the value of DeCommittedSize for the instance
func (instance *HeapContract) GetPropertyDeCommittedSize() (value interface{}, err error) {
	retValue, err := instance.GetProperty("DeCommittedSize")
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
func (instance *HeapContract) SetPropertyFreeSpace(value interface{}) (err error) {
	return instance.SetProperty("FreeSpace", (value))
}

// GetFreeSpace gets the value of FreeSpace for the instance
func (instance *HeapContract) GetPropertyFreeSpace() (value interface{}, err error) {
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
func (instance *HeapContract) SetPropertyHeapHandle(value uint32) (err error) {
	return instance.SetProperty("HeapHandle", (value))
}

// GetHeapHandle gets the value of HeapHandle for the instance
func (instance *HeapContract) GetPropertyHeapHandle() (value uint32, err error) {
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
func (instance *HeapContract) SetPropertyNoOfUCRs(value uint32) (err error) {
	return instance.SetProperty("NoOfUCRs", (value))
}

// GetNoOfUCRs gets the value of NoOfUCRs for the instance
func (instance *HeapContract) GetPropertyNoOfUCRs() (value uint32, err error) {
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
func (instance *HeapContract) SetPropertyReservedSpace(value interface{}) (err error) {
	return instance.SetProperty("ReservedSpace", (value))
}

// GetReservedSpace gets the value of ReservedSpace for the instance
func (instance *HeapContract) GetPropertyReservedSpace() (value interface{}, err error) {
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
