// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// Msvm_ConcreteJob struct
type Msvm_ConcreteJob struct { 
	*CIM_ConcreteJob

	// 
	Cancellable bool

	// 
	ErrorSummaryDescription string

	// 
	JobType ConcreteJob_JobType
}

	func NewMsvm_ConcreteJobEx1(instance *cim.WmiInstance) (newInstance *Msvm_ConcreteJob, err error) {tmp, err := NewCIM_ConcreteJobEx1(instance)
		
	if err != nil { return }
	newInstance = &Msvm_ConcreteJob {
	CIM_ConcreteJob: tmp,
	}
	return
	}
	

	func NewMsvm_ConcreteJobEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Msvm_ConcreteJob, err error) {tmp, err := NewCIM_ConcreteJobEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Msvm_ConcreteJob {
	CIM_ConcreteJob: tmp,
	}
	return
	}
	

// SetCancellable sets the value of Cancellable for the instance
func (instance *Msvm_ConcreteJob) SetPropertyCancellable(value bool) (err error) { 
    return instance.SetProperty("Cancellable", (value))
}

// GetCancellable gets the value of Cancellable for the instance
func (instance *Msvm_ConcreteJob) GetPropertyCancellable()(value bool, err error) { 
    retValue, err := instance.GetProperty("Cancellable")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetErrorSummaryDescription sets the value of ErrorSummaryDescription for the instance
func (instance *Msvm_ConcreteJob) SetPropertyErrorSummaryDescription(value string) (err error) { 
    return instance.SetProperty("ErrorSummaryDescription", (value))
}

// GetErrorSummaryDescription gets the value of ErrorSummaryDescription for the instance
func (instance *Msvm_ConcreteJob) GetPropertyErrorSummaryDescription()(value string, err error) { 
    retValue, err := instance.GetProperty("ErrorSummaryDescription")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetJobType sets the value of JobType for the instance
func (instance *Msvm_ConcreteJob) SetPropertyJobType(value ConcreteJob_JobType) (err error) { 
    return instance.SetProperty("JobType", (value))
}

// GetJobType gets the value of JobType for the instance
func (instance *Msvm_ConcreteJob) GetPropertyJobType()(value ConcreteJob_JobType, err error) { 
    retValue, err := instance.GetProperty("JobType")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = ConcreteJob_JobType(valuetmp)

    return
}

// 

// <param name="Errors" type="string []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ConcreteJob) GetErrorEx( /* OUT */ Errors []string) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetErrorEx" )
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


