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

// Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable struct
type Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable struct { 
	*Win32_PerfFormattedData

	// 
	AvgtimedeleteFileTableitem uint64

	// 
	AvgtimeFileTableenumeration uint64

	// 
	AvgtimeFileTablehandlekill uint64

	// 
	AvgtimemoveFileTableitem uint64

	// 
	AvgtimeperfileIOrequest uint64

	// 
	AvgtimeperfileIOresponse uint64

	// 
	AvgtimerenameFileTableitem uint64

	// 
	AvgtimetogetFileTableitem uint64

	// 
	AvgtimeupdateFileTableitem uint64

	// 
	FileTabledboperationsPersec uint64

	// 
	FileTableenumerationreqsPersec uint64

	// 
	FileTablefileIOrequestsPersec uint64

	// 
	FileTablefileIOresponsePersec uint64

	// 
	FileTableitemdeletereqsPersec uint64

	// 
	FileTableitemgetrequestsPersec uint64

	// 
	FileTableitemmovereqsPersec uint64

	// 
	FileTableitemrenamereqsPersec uint64

	// 
	FileTableitemupdatereqsPersec uint64

	// 
	FileTablekillhandleopsPersec uint64

	// 
	FileTabletableoperationsPersec uint64
}

	func NewWin32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTableEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable, err error) {tmp, err := NewWin32_PerfFormattedDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable {
	Win32_PerfFormattedData: tmp,
	}
	return
	}
	

	func NewWin32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTableEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable, err error) {tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable {
	Win32_PerfFormattedData: tmp,
	}
	return
	}
	

// SetAvgtimedeleteFileTableitem sets the value of AvgtimedeleteFileTableitem for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyAvgtimedeleteFileTableitem(value uint64) (err error) { 
    return instance.SetProperty("AvgtimedeleteFileTableitem", (value))
}

// GetAvgtimedeleteFileTableitem gets the value of AvgtimedeleteFileTableitem for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyAvgtimedeleteFileTableitem()(value uint64, err error) { 
    retValue, err := instance.GetProperty("AvgtimedeleteFileTableitem")
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

// SetAvgtimeFileTableenumeration sets the value of AvgtimeFileTableenumeration for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyAvgtimeFileTableenumeration(value uint64) (err error) { 
    return instance.SetProperty("AvgtimeFileTableenumeration", (value))
}

// GetAvgtimeFileTableenumeration gets the value of AvgtimeFileTableenumeration for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyAvgtimeFileTableenumeration()(value uint64, err error) { 
    retValue, err := instance.GetProperty("AvgtimeFileTableenumeration")
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

// SetAvgtimeFileTablehandlekill sets the value of AvgtimeFileTablehandlekill for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyAvgtimeFileTablehandlekill(value uint64) (err error) { 
    return instance.SetProperty("AvgtimeFileTablehandlekill", (value))
}

// GetAvgtimeFileTablehandlekill gets the value of AvgtimeFileTablehandlekill for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyAvgtimeFileTablehandlekill()(value uint64, err error) { 
    retValue, err := instance.GetProperty("AvgtimeFileTablehandlekill")
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

// SetAvgtimemoveFileTableitem sets the value of AvgtimemoveFileTableitem for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyAvgtimemoveFileTableitem(value uint64) (err error) { 
    return instance.SetProperty("AvgtimemoveFileTableitem", (value))
}

// GetAvgtimemoveFileTableitem gets the value of AvgtimemoveFileTableitem for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyAvgtimemoveFileTableitem()(value uint64, err error) { 
    retValue, err := instance.GetProperty("AvgtimemoveFileTableitem")
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

// SetAvgtimeperfileIOrequest sets the value of AvgtimeperfileIOrequest for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyAvgtimeperfileIOrequest(value uint64) (err error) { 
    return instance.SetProperty("AvgtimeperfileIOrequest", (value))
}

// GetAvgtimeperfileIOrequest gets the value of AvgtimeperfileIOrequest for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyAvgtimeperfileIOrequest()(value uint64, err error) { 
    retValue, err := instance.GetProperty("AvgtimeperfileIOrequest")
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

// SetAvgtimeperfileIOresponse sets the value of AvgtimeperfileIOresponse for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyAvgtimeperfileIOresponse(value uint64) (err error) { 
    return instance.SetProperty("AvgtimeperfileIOresponse", (value))
}

// GetAvgtimeperfileIOresponse gets the value of AvgtimeperfileIOresponse for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyAvgtimeperfileIOresponse()(value uint64, err error) { 
    retValue, err := instance.GetProperty("AvgtimeperfileIOresponse")
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

// SetAvgtimerenameFileTableitem sets the value of AvgtimerenameFileTableitem for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyAvgtimerenameFileTableitem(value uint64) (err error) { 
    return instance.SetProperty("AvgtimerenameFileTableitem", (value))
}

// GetAvgtimerenameFileTableitem gets the value of AvgtimerenameFileTableitem for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyAvgtimerenameFileTableitem()(value uint64, err error) { 
    retValue, err := instance.GetProperty("AvgtimerenameFileTableitem")
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

// SetAvgtimetogetFileTableitem sets the value of AvgtimetogetFileTableitem for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyAvgtimetogetFileTableitem(value uint64) (err error) { 
    return instance.SetProperty("AvgtimetogetFileTableitem", (value))
}

// GetAvgtimetogetFileTableitem gets the value of AvgtimetogetFileTableitem for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyAvgtimetogetFileTableitem()(value uint64, err error) { 
    retValue, err := instance.GetProperty("AvgtimetogetFileTableitem")
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

// SetAvgtimeupdateFileTableitem sets the value of AvgtimeupdateFileTableitem for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyAvgtimeupdateFileTableitem(value uint64) (err error) { 
    return instance.SetProperty("AvgtimeupdateFileTableitem", (value))
}

// GetAvgtimeupdateFileTableitem gets the value of AvgtimeupdateFileTableitem for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyAvgtimeupdateFileTableitem()(value uint64, err error) { 
    retValue, err := instance.GetProperty("AvgtimeupdateFileTableitem")
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

// SetFileTabledboperationsPersec sets the value of FileTabledboperationsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyFileTabledboperationsPersec(value uint64) (err error) { 
    return instance.SetProperty("FileTabledboperationsPersec", (value))
}

// GetFileTabledboperationsPersec gets the value of FileTabledboperationsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyFileTabledboperationsPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("FileTabledboperationsPersec")
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

// SetFileTableenumerationreqsPersec sets the value of FileTableenumerationreqsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyFileTableenumerationreqsPersec(value uint64) (err error) { 
    return instance.SetProperty("FileTableenumerationreqsPersec", (value))
}

// GetFileTableenumerationreqsPersec gets the value of FileTableenumerationreqsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyFileTableenumerationreqsPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("FileTableenumerationreqsPersec")
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

// SetFileTablefileIOrequestsPersec sets the value of FileTablefileIOrequestsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyFileTablefileIOrequestsPersec(value uint64) (err error) { 
    return instance.SetProperty("FileTablefileIOrequestsPersec", (value))
}

// GetFileTablefileIOrequestsPersec gets the value of FileTablefileIOrequestsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyFileTablefileIOrequestsPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("FileTablefileIOrequestsPersec")
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

// SetFileTablefileIOresponsePersec sets the value of FileTablefileIOresponsePersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyFileTablefileIOresponsePersec(value uint64) (err error) { 
    return instance.SetProperty("FileTablefileIOresponsePersec", (value))
}

// GetFileTablefileIOresponsePersec gets the value of FileTablefileIOresponsePersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyFileTablefileIOresponsePersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("FileTablefileIOresponsePersec")
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

// SetFileTableitemdeletereqsPersec sets the value of FileTableitemdeletereqsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyFileTableitemdeletereqsPersec(value uint64) (err error) { 
    return instance.SetProperty("FileTableitemdeletereqsPersec", (value))
}

// GetFileTableitemdeletereqsPersec gets the value of FileTableitemdeletereqsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyFileTableitemdeletereqsPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("FileTableitemdeletereqsPersec")
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

// SetFileTableitemgetrequestsPersec sets the value of FileTableitemgetrequestsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyFileTableitemgetrequestsPersec(value uint64) (err error) { 
    return instance.SetProperty("FileTableitemgetrequestsPersec", (value))
}

// GetFileTableitemgetrequestsPersec gets the value of FileTableitemgetrequestsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyFileTableitemgetrequestsPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("FileTableitemgetrequestsPersec")
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

// SetFileTableitemmovereqsPersec sets the value of FileTableitemmovereqsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyFileTableitemmovereqsPersec(value uint64) (err error) { 
    return instance.SetProperty("FileTableitemmovereqsPersec", (value))
}

// GetFileTableitemmovereqsPersec gets the value of FileTableitemmovereqsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyFileTableitemmovereqsPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("FileTableitemmovereqsPersec")
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

// SetFileTableitemrenamereqsPersec sets the value of FileTableitemrenamereqsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyFileTableitemrenamereqsPersec(value uint64) (err error) { 
    return instance.SetProperty("FileTableitemrenamereqsPersec", (value))
}

// GetFileTableitemrenamereqsPersec gets the value of FileTableitemrenamereqsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyFileTableitemrenamereqsPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("FileTableitemrenamereqsPersec")
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

// SetFileTableitemupdatereqsPersec sets the value of FileTableitemupdatereqsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyFileTableitemupdatereqsPersec(value uint64) (err error) { 
    return instance.SetProperty("FileTableitemupdatereqsPersec", (value))
}

// GetFileTableitemupdatereqsPersec gets the value of FileTableitemupdatereqsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyFileTableitemupdatereqsPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("FileTableitemupdatereqsPersec")
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

// SetFileTablekillhandleopsPersec sets the value of FileTablekillhandleopsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyFileTablekillhandleopsPersec(value uint64) (err error) { 
    return instance.SetProperty("FileTablekillhandleopsPersec", (value))
}

// GetFileTablekillhandleopsPersec gets the value of FileTablekillhandleopsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyFileTablekillhandleopsPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("FileTablekillhandleopsPersec")
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

// SetFileTabletableoperationsPersec sets the value of FileTabletableoperationsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyFileTabletableoperationsPersec(value uint64) (err error) { 
    return instance.SetProperty("FileTabletableoperationsPersec", (value))
}

// GetFileTabletableoperationsPersec gets the value of FileTabletableoperationsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyFileTabletableoperationsPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("FileTabletableoperationsPersec")
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

