// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// FltIoCompletion struct
type FltIoCompletion struct { 
	*FileIo

	// 
	CallbackDataPtr uint32

	// 
	FileContext uint32

	// 
	FileObject uint32

	// 
	InitialTime interface{}

	// 
	IrpPtr uint32

	// 
	MajorFunction uint32

	// 
	RoutineAddr uint32
}

	func NewFltIoCompletionEx1(instance *cim.WmiInstance) (newInstance *FltIoCompletion, err error) {tmp, err := NewFileIoEx1(instance)
		
	if err != nil { return }
	newInstance = &FltIoCompletion {
	FileIo: tmp,
	}
	return
	}
	

	func NewFltIoCompletionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *FltIoCompletion, err error) {tmp, err := NewFileIoEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &FltIoCompletion {
	FileIo: tmp,
	}
	return
	}
	

// SetCallbackDataPtr sets the value of CallbackDataPtr for the instance
func (instance *FltIoCompletion) SetPropertyCallbackDataPtr(value uint32) (err error) { 
    return instance.SetProperty("CallbackDataPtr", (value))
}

// GetCallbackDataPtr gets the value of CallbackDataPtr for the instance
func (instance *FltIoCompletion) GetPropertyCallbackDataPtr()(value uint32, err error) { 
    retValue, err := instance.GetProperty("CallbackDataPtr")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetFileContext sets the value of FileContext for the instance
func (instance *FltIoCompletion) SetPropertyFileContext(value uint32) (err error) { 
    return instance.SetProperty("FileContext", (value))
}

// GetFileContext gets the value of FileContext for the instance
func (instance *FltIoCompletion) GetPropertyFileContext()(value uint32, err error) { 
    retValue, err := instance.GetProperty("FileContext")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetFileObject sets the value of FileObject for the instance
func (instance *FltIoCompletion) SetPropertyFileObject(value uint32) (err error) { 
    return instance.SetProperty("FileObject", (value))
}

// GetFileObject gets the value of FileObject for the instance
func (instance *FltIoCompletion) GetPropertyFileObject()(value uint32, err error) { 
    retValue, err := instance.GetProperty("FileObject")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetInitialTime sets the value of InitialTime for the instance
func (instance *FltIoCompletion) SetPropertyInitialTime(value interface{}) (err error) { 
    return instance.SetProperty("InitialTime", (value))
}

// GetInitialTime gets the value of InitialTime for the instance
func (instance *FltIoCompletion) GetPropertyInitialTime()(value interface{}, err error) { 
    retValue, err := instance.GetProperty("InitialTime")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(interface{}); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = interface{}(valuetmp)

    return
}

// SetIrpPtr sets the value of IrpPtr for the instance
func (instance *FltIoCompletion) SetPropertyIrpPtr(value uint32) (err error) { 
    return instance.SetProperty("IrpPtr", (value))
}

// GetIrpPtr gets the value of IrpPtr for the instance
func (instance *FltIoCompletion) GetPropertyIrpPtr()(value uint32, err error) { 
    retValue, err := instance.GetProperty("IrpPtr")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetMajorFunction sets the value of MajorFunction for the instance
func (instance *FltIoCompletion) SetPropertyMajorFunction(value uint32) (err error) { 
    return instance.SetProperty("MajorFunction", (value))
}

// GetMajorFunction gets the value of MajorFunction for the instance
func (instance *FltIoCompletion) GetPropertyMajorFunction()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MajorFunction")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetRoutineAddr sets the value of RoutineAddr for the instance
func (instance *FltIoCompletion) SetPropertyRoutineAddr(value uint32) (err error) { 
    return instance.SetProperty("RoutineAddr", (value))
}

// GetRoutineAddr gets the value of RoutineAddr for the instance
func (instance *FltIoCompletion) GetPropertyRoutineAddr()(value uint32, err error) { 
    retValue, err := instance.GetProperty("RoutineAddr")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

