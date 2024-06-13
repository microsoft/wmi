// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_CompatibilityVector struct
type Msvm_CompatibilityVector struct {
	*cim.WmiInstance

	//
	CompareOperation uint32

	//
	CompatibilityInfo uint64

	//
	VectorId uint32
}

func NewMsvm_CompatibilityVectorEx1(instance *cim.WmiInstance) (newInstance *Msvm_CompatibilityVector, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &Msvm_CompatibilityVector{
		WmiInstance: tmp,
	}
	return
}

func NewMsvm_CompatibilityVectorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_CompatibilityVector, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_CompatibilityVector{
		WmiInstance: tmp,
	}
	return
}

// SetCompareOperation sets the value of CompareOperation for the instance
func (instance *Msvm_CompatibilityVector) SetPropertyCompareOperation(value uint32) (err error) {
	return instance.SetProperty("CompareOperation", (value))
}

// GetCompareOperation gets the value of CompareOperation for the instance
func (instance *Msvm_CompatibilityVector) GetPropertyCompareOperation() (value uint32, err error) {
	retValue, err := instance.GetProperty("CompareOperation")
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

// SetCompatibilityInfo sets the value of CompatibilityInfo for the instance
func (instance *Msvm_CompatibilityVector) SetPropertyCompatibilityInfo(value uint64) (err error) {
	return instance.SetProperty("CompatibilityInfo", (value))
}

// GetCompatibilityInfo gets the value of CompatibilityInfo for the instance
func (instance *Msvm_CompatibilityVector) GetPropertyCompatibilityInfo() (value uint64, err error) {
	retValue, err := instance.GetProperty("CompatibilityInfo")
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

// SetVectorId sets the value of VectorId for the instance
func (instance *Msvm_CompatibilityVector) SetPropertyVectorId(value uint32) (err error) {
	return instance.SetProperty("VectorId", (value))
}

// GetVectorId gets the value of VectorId for the instance
func (instance *Msvm_CompatibilityVector) GetPropertyVectorId() (value uint32, err error) {
	retValue, err := instance.GetProperty("VectorId")
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
