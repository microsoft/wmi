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

// Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLatches struct
type Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLatches struct { 
	*Win32_PerfFormattedData

	// 
	AverageLatchWaitTimems uint64

	// 
	LatchWaitsPersec uint64

	// 
	NumberofSuperLatches uint64

	// 
	SuperLatchDemotionsPersec uint64

	// 
	SuperLatchPromotionsPersec uint64

	// 
	TotalLatchWaitTimems uint64
}

	func NewWin32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLatchesEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLatches, err error) {tmp, err := NewWin32_PerfFormattedDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLatches {
	Win32_PerfFormattedData: tmp,
	}
	return
	}
	

	func NewWin32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLatchesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLatches, err error) {tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLatches {
	Win32_PerfFormattedData: tmp,
	}
	return
	}
	

// SetAverageLatchWaitTimems sets the value of AverageLatchWaitTimems for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLatches) SetPropertyAverageLatchWaitTimems(value uint64) (err error) { 
    return instance.SetProperty("AverageLatchWaitTimems", (value))
}

// GetAverageLatchWaitTimems gets the value of AverageLatchWaitTimems for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLatches) GetPropertyAverageLatchWaitTimems()(value uint64, err error) { 
    retValue, err := instance.GetProperty("AverageLatchWaitTimems")
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

// SetLatchWaitsPersec sets the value of LatchWaitsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLatches) SetPropertyLatchWaitsPersec(value uint64) (err error) { 
    return instance.SetProperty("LatchWaitsPersec", (value))
}

// GetLatchWaitsPersec gets the value of LatchWaitsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLatches) GetPropertyLatchWaitsPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("LatchWaitsPersec")
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

// SetNumberofSuperLatches sets the value of NumberofSuperLatches for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLatches) SetPropertyNumberofSuperLatches(value uint64) (err error) { 
    return instance.SetProperty("NumberofSuperLatches", (value))
}

// GetNumberofSuperLatches gets the value of NumberofSuperLatches for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLatches) GetPropertyNumberofSuperLatches()(value uint64, err error) { 
    retValue, err := instance.GetProperty("NumberofSuperLatches")
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

// SetSuperLatchDemotionsPersec sets the value of SuperLatchDemotionsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLatches) SetPropertySuperLatchDemotionsPersec(value uint64) (err error) { 
    return instance.SetProperty("SuperLatchDemotionsPersec", (value))
}

// GetSuperLatchDemotionsPersec gets the value of SuperLatchDemotionsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLatches) GetPropertySuperLatchDemotionsPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("SuperLatchDemotionsPersec")
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

// SetSuperLatchPromotionsPersec sets the value of SuperLatchPromotionsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLatches) SetPropertySuperLatchPromotionsPersec(value uint64) (err error) { 
    return instance.SetProperty("SuperLatchPromotionsPersec", (value))
}

// GetSuperLatchPromotionsPersec gets the value of SuperLatchPromotionsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLatches) GetPropertySuperLatchPromotionsPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("SuperLatchPromotionsPersec")
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

// SetTotalLatchWaitTimems sets the value of TotalLatchWaitTimems for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLatches) SetPropertyTotalLatchWaitTimems(value uint64) (err error) { 
    return instance.SetProperty("TotalLatchWaitTimems", (value))
}

// GetTotalLatchWaitTimems gets the value of TotalLatchWaitTimems for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDLatches) GetPropertyTotalLatchWaitTimems()(value uint64, err error) { 
    retValue, err := instance.GetProperty("TotalLatchWaitTimems")
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

