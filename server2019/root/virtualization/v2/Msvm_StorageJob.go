// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// Msvm_StorageJob struct
type Msvm_StorageJob struct {
	CIM_ConcreteJob

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

// SetCancellable sets the value of Cancellable for the instance
func (instance *Msvm_StorageJob) SetPropertyCancellable(value bool) (err error) {
	return instance.SetProperty("Cancellable", value)
}

// GetCancellable gets the value of Cancellable for the instance
func (instance *Msvm_StorageJob) GetPropertyCancellable() (value bool, err error) {
	retValue, err := instance.GetProperty("Cancellable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetChild sets the value of Child for the instance
func (instance *Msvm_StorageJob) SetPropertyChild(value string) (err error) {
	return instance.SetProperty("Child", value)
}

// GetChild gets the value of Child for the instance
func (instance *Msvm_StorageJob) GetPropertyChild() (value string, err error) {
	retValue, err := instance.GetProperty("Child")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetErrorSummaryDescription sets the value of ErrorSummaryDescription for the instance
func (instance *Msvm_StorageJob) SetPropertyErrorSummaryDescription(value string) (err error) {
	return instance.SetProperty("ErrorSummaryDescription", value)
}

// GetErrorSummaryDescription gets the value of ErrorSummaryDescription for the instance
func (instance *Msvm_StorageJob) GetPropertyErrorSummaryDescription() (value string, err error) {
	retValue, err := instance.GetProperty("ErrorSummaryDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetJobCompletionStatusCode sets the value of JobCompletionStatusCode for the instance
func (instance *Msvm_StorageJob) SetPropertyJobCompletionStatusCode(value uint32) (err error) {
	return instance.SetProperty("JobCompletionStatusCode", value)
}

// GetJobCompletionStatusCode gets the value of JobCompletionStatusCode for the instance
func (instance *Msvm_StorageJob) GetPropertyJobCompletionStatusCode() (value uint32, err error) {
	retValue, err := instance.GetProperty("JobCompletionStatusCode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetJobType sets the value of JobType for the instance
func (instance *Msvm_StorageJob) SetPropertyJobType(value uint16) (err error) {
	return instance.SetProperty("JobType", value)
}

// GetJobType gets the value of JobType for the instance
func (instance *Msvm_StorageJob) GetPropertyJobType() (value uint16, err error) {
	retValue, err := instance.GetProperty("JobType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParent sets the value of Parent for the instance
func (instance *Msvm_StorageJob) SetPropertyParent(value string) (err error) {
	return instance.SetProperty("Parent", value)
}

// GetParent gets the value of Parent for the instance
func (instance *Msvm_StorageJob) GetPropertyParent() (value string, err error) {
	retValue, err := instance.GetProperty("Parent")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
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
