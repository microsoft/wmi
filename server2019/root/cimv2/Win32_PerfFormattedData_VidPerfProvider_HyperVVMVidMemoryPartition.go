// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// Win32_PerfFormattedData_VidPerfProvider_HyperVVMVidMemoryPartition struct
type Win32_PerfFormattedData_VidPerfProvider_HyperVVMVidMemoryPartition struct { 
	*Win32_PerfFormattedData

	// 
	Index uint64
}

	func NewWin32_PerfFormattedData_VidPerfProvider_HyperVVMVidMemoryPartitionEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_VidPerfProvider_HyperVVMVidMemoryPartition, err error) {tmp, err := NewWin32_PerfFormattedDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Win32_PerfFormattedData_VidPerfProvider_HyperVVMVidMemoryPartition {
	Win32_PerfFormattedData: tmp,
	}
	return
	}
	

	func NewWin32_PerfFormattedData_VidPerfProvider_HyperVVMVidMemoryPartitionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Win32_PerfFormattedData_VidPerfProvider_HyperVVMVidMemoryPartition, err error) {tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Win32_PerfFormattedData_VidPerfProvider_HyperVVMVidMemoryPartition {
	Win32_PerfFormattedData: tmp,
	}
	return
	}
	

// SetIndex sets the value of Index for the instance
func (instance *Win32_PerfFormattedData_VidPerfProvider_HyperVVMVidMemoryPartition) SetPropertyIndex(value uint64) (err error) { 
    return instance.SetProperty("Index", (value))
}

// GetIndex gets the value of Index for the instance
func (instance *Win32_PerfFormattedData_VidPerfProvider_HyperVVMVidMemoryPartition) GetPropertyIndex()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Index")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

