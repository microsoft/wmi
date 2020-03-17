// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

// Msvm_ConcreteJob struct
type Msvm_ConcreteJob struct {
	CIM_ConcreteJob

	//
	Cancellable bool

	//
	ErrorSummaryDescription string

	//
	JobType ConcreteJob_JobType
}

// SetCancellable sets the value of Cancellable for the instance
func (instance *Msvm_ConcreteJob) SetPropertyCancellable(value bool) (err error) {
	return instance.SetProperty("Cancellable", value)
}

// GetCancellable gets the value of Cancellable for the instance
func (instance *Msvm_ConcreteJob) GetPropertyCancellable() (value bool, err error) {
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

// SetErrorSummaryDescription sets the value of ErrorSummaryDescription for the instance
func (instance *Msvm_ConcreteJob) SetPropertyErrorSummaryDescription(value string) (err error) {
	return instance.SetProperty("ErrorSummaryDescription", value)
}

// GetErrorSummaryDescription gets the value of ErrorSummaryDescription for the instance
func (instance *Msvm_ConcreteJob) GetPropertyErrorSummaryDescription() (value string, err error) {
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

// SetJobType sets the value of JobType for the instance
func (instance *Msvm_ConcreteJob) SetPropertyJobType(value ConcreteJob_JobType) (err error) {
	return instance.SetProperty("JobType", value)
}

// GetJobType gets the value of JobType for the instance
func (instance *Msvm_ConcreteJob) GetPropertyJobType() (value ConcreteJob_JobType, err error) {
	retValue, err := instance.GetProperty("JobType")
	if err != nil {
		return
	}
	value, ok := retValue.(ConcreteJob_JobType)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="Errors" type="string []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ConcreteJob) GetErrorEx( /* OUT */ Errors []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetErrorEx")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
