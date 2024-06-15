// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_StorageJob struct
type Msvm_StorageJob struct {
	*CIM_ConcreteJob

	//
	Cancellable bool

	//
	Child string

	//
	ErrorSummaryDescription string

	//
	JobCompletionStatusCode uint32

	//
	JobType uint16

	//
	Parent string
}

func NewMsvm_StorageJobEx1(instance *cim.WmiInstance) (newInstance *Msvm_StorageJob, err error) {
	tmp, err := NewCIM_ConcreteJobEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_StorageJob{
		CIM_ConcreteJob: tmp,
	}
	return
}

func NewMsvm_StorageJobEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_StorageJob, err error) {
	tmp, err := NewCIM_ConcreteJobEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_StorageJob{
		CIM_ConcreteJob: tmp,
	}
	return
}

// SetCancellable sets the value of Cancellable for the instance
func (instance *Msvm_StorageJob) SetPropertyCancellable(value bool) (err error) {
	return instance.SetProperty("Cancellable", (value))
}

// GetCancellable gets the value of Cancellable for the instance
func (instance *Msvm_StorageJob) GetPropertyCancellable() (value bool, err error) {
	retValue, err := instance.GetProperty("Cancellable")
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

// SetChild sets the value of Child for the instance
func (instance *Msvm_StorageJob) SetPropertyChild(value string) (err error) {
	return instance.SetProperty("Child", (value))
}

// GetChild gets the value of Child for the instance
func (instance *Msvm_StorageJob) GetPropertyChild() (value string, err error) {
	retValue, err := instance.GetProperty("Child")
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

// SetErrorSummaryDescription sets the value of ErrorSummaryDescription for the instance
func (instance *Msvm_StorageJob) SetPropertyErrorSummaryDescription(value string) (err error) {
	return instance.SetProperty("ErrorSummaryDescription", (value))
}

// GetErrorSummaryDescription gets the value of ErrorSummaryDescription for the instance
func (instance *Msvm_StorageJob) GetPropertyErrorSummaryDescription() (value string, err error) {
	retValue, err := instance.GetProperty("ErrorSummaryDescription")
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

// SetJobCompletionStatusCode sets the value of JobCompletionStatusCode for the instance
func (instance *Msvm_StorageJob) SetPropertyJobCompletionStatusCode(value uint32) (err error) {
	return instance.SetProperty("JobCompletionStatusCode", (value))
}

// GetJobCompletionStatusCode gets the value of JobCompletionStatusCode for the instance
func (instance *Msvm_StorageJob) GetPropertyJobCompletionStatusCode() (value uint32, err error) {
	retValue, err := instance.GetProperty("JobCompletionStatusCode")
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

// SetJobType sets the value of JobType for the instance
func (instance *Msvm_StorageJob) SetPropertyJobType(value uint16) (err error) {
	return instance.SetProperty("JobType", (value))
}

// GetJobType gets the value of JobType for the instance
func (instance *Msvm_StorageJob) GetPropertyJobType() (value uint16, err error) {
	retValue, err := instance.GetProperty("JobType")
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

// SetParent sets the value of Parent for the instance
func (instance *Msvm_StorageJob) SetPropertyParent(value string) (err error) {
	return instance.SetProperty("Parent", (value))
}

// GetParent gets the value of Parent for the instance
func (instance *Msvm_StorageJob) GetPropertyParent() (value string, err error) {
	retValue, err := instance.GetProperty("Parent")
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

//

// <param name="Errors" type="string []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_StorageJob) GetErrorEx( /* OUT */ Errors []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetErrorEx")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
