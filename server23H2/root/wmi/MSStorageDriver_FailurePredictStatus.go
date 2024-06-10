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
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSStorageDriver_FailurePredictStatus struct
type MSStorageDriver_FailurePredictStatus struct {
	*MSStorageDriver

	//
	Active bool

	//
	InstanceName string

	//
	PredictFailure bool

	//
	Reason uint32
}

func NewMSStorageDriver_FailurePredictStatusEx1(instance *cim.WmiInstance) (newInstance *MSStorageDriver_FailurePredictStatus, err error) {
	tmp, err := NewMSStorageDriverEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSStorageDriver_FailurePredictStatus{
		MSStorageDriver: tmp,
	}
	return
}

func NewMSStorageDriver_FailurePredictStatusEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSStorageDriver_FailurePredictStatus, err error) {
	tmp, err := NewMSStorageDriverEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSStorageDriver_FailurePredictStatus{
		MSStorageDriver: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSStorageDriver_FailurePredictStatus) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSStorageDriver_FailurePredictStatus) GetPropertyActive() (value bool, err error) {
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
func (instance *MSStorageDriver_FailurePredictStatus) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSStorageDriver_FailurePredictStatus) GetPropertyInstanceName() (value string, err error) {
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

// SetPredictFailure sets the value of PredictFailure for the instance
func (instance *MSStorageDriver_FailurePredictStatus) SetPropertyPredictFailure(value bool) (err error) {
	return instance.SetProperty("PredictFailure", (value))
}

// GetPredictFailure gets the value of PredictFailure for the instance
func (instance *MSStorageDriver_FailurePredictStatus) GetPropertyPredictFailure() (value bool, err error) {
	retValue, err := instance.GetProperty("PredictFailure")
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

// SetReason sets the value of Reason for the instance
func (instance *MSStorageDriver_FailurePredictStatus) SetPropertyReason(value uint32) (err error) {
	return instance.SetProperty("Reason", (value))
}

// GetReason gets the value of Reason for the instance
func (instance *MSStorageDriver_FailurePredictStatus) GetPropertyReason() (value uint32, err error) {
	retValue, err := instance.GetProperty("Reason")
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
