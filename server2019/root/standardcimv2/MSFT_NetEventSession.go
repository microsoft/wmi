// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSFT_NetEventSession struct
type MSFT_NetEventSession struct { 
	*CIM_LogicalElement

	// 
	CaptureMode uint8

	// 
	Guid string

	// 
	LocalFilePath string

	// 
	MaxFileSize uint32

	// 
	MaxNumberOfBuffers uint8

	// 
	SessionStatus uint8

	// 
	TraceBufferSize uint32
}

	func NewMSFT_NetEventSessionEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetEventSession, err error) {tmp, err := NewCIM_LogicalElementEx1(instance)
		
	if err != nil { return }
	newInstance = &MSFT_NetEventSession {
	CIM_LogicalElement: tmp,
	}
	return
	}
	

	func NewMSFT_NetEventSessionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_NetEventSession, err error) {tmp, err := NewCIM_LogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_NetEventSession {
	CIM_LogicalElement: tmp,
	}
	return
	}
	

// SetCaptureMode sets the value of CaptureMode for the instance
func (instance *MSFT_NetEventSession) SetPropertyCaptureMode(value uint8) (err error) { 
    return instance.SetProperty("CaptureMode", (value))
}

// GetCaptureMode gets the value of CaptureMode for the instance
func (instance *MSFT_NetEventSession) GetPropertyCaptureMode()(value uint8, err error) { 
    retValue, err := instance.GetProperty("CaptureMode")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// SetGuid sets the value of Guid for the instance
func (instance *MSFT_NetEventSession) SetPropertyGuid(value string) (err error) { 
    return instance.SetProperty("Guid", (value))
}

// GetGuid gets the value of Guid for the instance
func (instance *MSFT_NetEventSession) GetPropertyGuid()(value string, err error) { 
    retValue, err := instance.GetProperty("Guid")
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

// SetLocalFilePath sets the value of LocalFilePath for the instance
func (instance *MSFT_NetEventSession) SetPropertyLocalFilePath(value string) (err error) { 
    return instance.SetProperty("LocalFilePath", (value))
}

// GetLocalFilePath gets the value of LocalFilePath for the instance
func (instance *MSFT_NetEventSession) GetPropertyLocalFilePath()(value string, err error) { 
    retValue, err := instance.GetProperty("LocalFilePath")
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

// SetMaxFileSize sets the value of MaxFileSize for the instance
func (instance *MSFT_NetEventSession) SetPropertyMaxFileSize(value uint32) (err error) { 
    return instance.SetProperty("MaxFileSize", (value))
}

// GetMaxFileSize gets the value of MaxFileSize for the instance
func (instance *MSFT_NetEventSession) GetPropertyMaxFileSize()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MaxFileSize")
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

// SetMaxNumberOfBuffers sets the value of MaxNumberOfBuffers for the instance
func (instance *MSFT_NetEventSession) SetPropertyMaxNumberOfBuffers(value uint8) (err error) { 
    return instance.SetProperty("MaxNumberOfBuffers", (value))
}

// GetMaxNumberOfBuffers gets the value of MaxNumberOfBuffers for the instance
func (instance *MSFT_NetEventSession) GetPropertyMaxNumberOfBuffers()(value uint8, err error) { 
    retValue, err := instance.GetProperty("MaxNumberOfBuffers")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// SetSessionStatus sets the value of SessionStatus for the instance
func (instance *MSFT_NetEventSession) SetPropertySessionStatus(value uint8) (err error) { 
    return instance.SetProperty("SessionStatus", (value))
}

// GetSessionStatus gets the value of SessionStatus for the instance
func (instance *MSFT_NetEventSession) GetPropertySessionStatus()(value uint8, err error) { 
    retValue, err := instance.GetProperty("SessionStatus")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// SetTraceBufferSize sets the value of TraceBufferSize for the instance
func (instance *MSFT_NetEventSession) SetPropertyTraceBufferSize(value uint32) (err error) { 
    return instance.SetProperty("TraceBufferSize", (value))
}

// GetTraceBufferSize gets the value of TraceBufferSize for the instance
func (instance *MSFT_NetEventSession) GetPropertyTraceBufferSize()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TraceBufferSize")
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

// 

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetEventSession) Start() (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("Start" );
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetEventSession) Stop() (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("Stop" );
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


