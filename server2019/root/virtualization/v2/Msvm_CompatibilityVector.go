// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_CompatibilityVector struct
type Msvm_CompatibilityVector struct {
	cim.WmiInstance

	//
	CompareOperation uint32

	//
	CompatibilityInfo uint64

	//
	VectorId uint32
}

// SetCompareOperation sets the value of CompareOperation for the instance
func (instance *Msvm_CompatibilityVector) SetPropertyCompareOperation(value uint32) (err error) {
	return instance.SetProperty("CompareOperation", value)
}

// GetCompareOperation gets the value of CompareOperation for the instance
func (instance *Msvm_CompatibilityVector) GetPropertyCompareOperation() (value uint32, err error) {
	retValue, err := instance.GetProperty("CompareOperation")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCompatibilityInfo sets the value of CompatibilityInfo for the instance
func (instance *Msvm_CompatibilityVector) SetPropertyCompatibilityInfo(value uint64) (err error) {
	return instance.SetProperty("CompatibilityInfo", value)
}

// GetCompatibilityInfo gets the value of CompatibilityInfo for the instance
func (instance *Msvm_CompatibilityVector) GetPropertyCompatibilityInfo() (value uint64, err error) {
	retValue, err := instance.GetProperty("CompatibilityInfo")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVectorId sets the value of VectorId for the instance
func (instance *Msvm_CompatibilityVector) SetPropertyVectorId(value uint32) (err error) {
	return instance.SetProperty("VectorId", value)
}

// GetVectorId gets the value of VectorId for the instance
func (instance *Msvm_CompatibilityVector) GetPropertyVectorId() (value uint32, err error) {
	retValue, err := instance.GetProperty("VectorId")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
