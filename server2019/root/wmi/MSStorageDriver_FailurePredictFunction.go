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

// MSStorageDriver_FailurePredictFunction struct
type MSStorageDriver_FailurePredictFunction struct { 
	*MSStorageDriver

	// 
	Active bool

	// 
	InstanceName string
}

	func NewMSStorageDriver_FailurePredictFunctionEx1(instance *cim.WmiInstance) (newInstance *MSStorageDriver_FailurePredictFunction, err error) {tmp, err := NewMSStorageDriverEx1(instance)
		
	if err != nil { return }
	newInstance = &MSStorageDriver_FailurePredictFunction {
	MSStorageDriver: tmp,
	}
	return
	}
	

	func NewMSStorageDriver_FailurePredictFunctionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSStorageDriver_FailurePredictFunction, err error) {tmp, err := NewMSStorageDriverEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSStorageDriver_FailurePredictFunction {
	MSStorageDriver: tmp,
	}
	return
	}
	

// SetActive sets the value of Active for the instance
func (instance *MSStorageDriver_FailurePredictFunction) SetPropertyActive(value bool) (err error) { 
    return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSStorageDriver_FailurePredictFunction) GetPropertyActive()(value bool, err error) { 
    retValue, err := instance.GetProperty("Active")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSStorageDriver_FailurePredictFunction) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSStorageDriver_FailurePredictFunction) GetPropertyInstanceName()(value string, err error) { 
    retValue, err := instance.GetProperty("InstanceName")
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

// 

// <param name="Allow" type="bool "></param>
func (instance *MSStorageDriver_FailurePredictFunction) AllowPerformanceHit( /* IN */ Allow bool) (err error) {_, err = instance.InvokeMethodWithReturn("AllowPerformanceHit" , Allow);
	if err != nil { return }
	return
	
}


// 

// <param name="Enable" type="bool "></param>
func (instance *MSStorageDriver_FailurePredictFunction) EnableDisableHardwareFailurePrediction( /* IN */ Enable bool) (err error) {_, err = instance.InvokeMethodWithReturn("EnableDisableHardwareFailurePrediction" , Enable);
	if err != nil { return }
	return
	
}


// 

// <param name="Enable" type="bool "></param>
// <param name="Period" type="uint32 "></param>
func (instance *MSStorageDriver_FailurePredictFunction) EnableDisableFailurePredictionPolling( /* IN */ Period uint32,
 /* IN */ Enable bool) (err error) {_, err = instance.InvokeMethodWithReturn("EnableDisableFailurePredictionPolling" , Period, Enable);
	if err != nil { return }
	return
	
}


// 

// <param name="Capability" type="uint32 "></param>
func (instance *MSStorageDriver_FailurePredictFunction) GetFailurePredictionCapability( /* OUT */ Capability uint32) (err error) {_, err = instance.InvokeMethod("GetFailurePredictionCapability" )
	if err != nil { return }
	return
	
}


// 

// <param name="Success" type="bool "></param>
func (instance *MSStorageDriver_FailurePredictFunction) EnableOfflineDiags( /* OUT */ Success bool) (err error) {_, err = instance.InvokeMethod("EnableOfflineDiags" )
	if err != nil { return }
	return
	
}


// 

// <param name="LogAddress" type="uint8 "></param>
// <param name="SectorCount" type="uint8 "></param>

// <param name="Length" type="uint32 "></param>
// <param name="LogSectors" type="uint8 []"></param>
func (instance *MSStorageDriver_FailurePredictFunction) ReadLogSectors( /* IN */ LogAddress uint8,
 /* IN */ SectorCount uint8,
 /* OUT */ Length uint32,
 /* OUT */ LogSectors []uint8) (err error) {_, err = instance.InvokeMethod("ReadLogSectors" , LogAddress, SectorCount)
	if err != nil { return }
	return
	
}


// 

// <param name="Length" type="uint32 "></param>
// <param name="LogAddress" type="uint8 "></param>
// <param name="LogSectors" type="uint8 []"></param>
// <param name="SectorCount" type="uint8 "></param>

// <param name="Success" type="bool "></param>
func (instance *MSStorageDriver_FailurePredictFunction) WriteLogSectors( /* IN */ LogAddress uint8,
 /* IN */ SectorCount uint8,
 /* IN */ Length uint32,
 /* IN */ LogSectors []uint8,
 /* OUT */ Success bool) (err error) {_, err = instance.InvokeMethod("WriteLogSectors" , LogAddress, SectorCount, Length, LogSectors)
	if err != nil { return }
	return
	
}


// 

// <param name="Subcommand" type="uint8 "></param>

// <param name="ReturnCode" type="uint32 "></param>
func (instance *MSStorageDriver_FailurePredictFunction) ExecuteSelfTest( /* IN */ Subcommand uint8,
 /* OUT */ ReturnCode uint32) (err error) {_, err = instance.InvokeMethod("ExecuteSelfTest" , Subcommand)
	if err != nil { return }
	return
	
}


