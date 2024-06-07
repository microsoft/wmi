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

// Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDReplicationLogreader struct
type Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDReplicationLogreader struct { 
	*Win32_PerfRawData

	// 
	LogreaderDeliveredCmdsPersec uint64

	// 
	LogreaderDeliveredTransPersec uint64

	// 
	LogreaderDeliveryLatency uint64
}

	func NewWin32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDReplicationLogreaderEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDReplicationLogreader, err error) {tmp, err := NewWin32_PerfRawDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDReplicationLogreader {
	Win32_PerfRawData: tmp,
	}
	return
	}
	

	func NewWin32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDReplicationLogreaderEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDReplicationLogreader, err error) {tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDReplicationLogreader {
	Win32_PerfRawData: tmp,
	}
	return
	}
	

// SetLogreaderDeliveredCmdsPersec sets the value of LogreaderDeliveredCmdsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDReplicationLogreader) SetPropertyLogreaderDeliveredCmdsPersec(value uint64) (err error) { 
    return instance.SetProperty("LogreaderDeliveredCmdsPersec", (value))
}

// GetLogreaderDeliveredCmdsPersec gets the value of LogreaderDeliveredCmdsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDReplicationLogreader) GetPropertyLogreaderDeliveredCmdsPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("LogreaderDeliveredCmdsPersec")
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

// SetLogreaderDeliveredTransPersec sets the value of LogreaderDeliveredTransPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDReplicationLogreader) SetPropertyLogreaderDeliveredTransPersec(value uint64) (err error) { 
    return instance.SetProperty("LogreaderDeliveredTransPersec", (value))
}

// GetLogreaderDeliveredTransPersec gets the value of LogreaderDeliveredTransPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDReplicationLogreader) GetPropertyLogreaderDeliveredTransPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("LogreaderDeliveredTransPersec")
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

// SetLogreaderDeliveryLatency sets the value of LogreaderDeliveryLatency for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDReplicationLogreader) SetPropertyLogreaderDeliveryLatency(value uint64) (err error) { 
    return instance.SetProperty("LogreaderDeliveryLatency", (value))
}

// GetLogreaderDeliveryLatency gets the value of LogreaderDeliveryLatency for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDReplicationLogreader) GetPropertyLogreaderDeliveryLatency()(value uint64, err error) { 
    retValue, err := instance.GetProperty("LogreaderDeliveryLatency")
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

