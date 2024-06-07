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

// Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage struct
type Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage struct { 
	*Win32_PerfRawData

	// 
	AvgBytesPerRead uint64

	// 
	AvgBytesPerRead_Base uint32

	// 
	AvgBytesPerTransfer uint64

	// 
	AvgBytesPerTransfer_Base uint32

	// 
	AvgBytesPerWrite uint64

	// 
	AvgBytesPerWrite_Base uint32

	// 
	AvgmicrosecPerRead uint64

	// 
	AvgmicrosecPerRead_Base uint32

	// 
	AvgmicrosecPerTransfer uint64

	// 
	AvgmicrosecPerTransfer_Base uint32

	// 
	AvgmicrosecPerWrite uint64

	// 
	AvgmicrosecPerWrite_Base uint32

	// 
	HTTPStorageIOretryPersec uint64

	// 
	OutstandingHTTPStorageIO uint64

	// 
	ReadBytesPerSec uint64

	// 
	ReadsPerSec uint64

	// 
	TotalBytesPerSec uint64

	// 
	TransfersPerSec uint64

	// 
	WriteBytesPerSec uint64

	// 
	WritesPerSec uint64
}

	func NewWin32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorageEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage, err error) {tmp, err := NewWin32_PerfRawDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage {
	Win32_PerfRawData: tmp,
	}
	return
	}
	

	func NewWin32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorageEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage, err error) {tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage {
	Win32_PerfRawData: tmp,
	}
	return
	}
	

// SetAvgBytesPerRead sets the value of AvgBytesPerRead for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) SetPropertyAvgBytesPerRead(value uint64) (err error) { 
    return instance.SetProperty("AvgBytesPerRead", (value))
}

// GetAvgBytesPerRead gets the value of AvgBytesPerRead for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) GetPropertyAvgBytesPerRead()(value uint64, err error) { 
    retValue, err := instance.GetProperty("AvgBytesPerRead")
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

// SetAvgBytesPerRead_Base sets the value of AvgBytesPerRead_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) SetPropertyAvgBytesPerRead_Base(value uint32) (err error) { 
    return instance.SetProperty("AvgBytesPerRead_Base", (value))
}

// GetAvgBytesPerRead_Base gets the value of AvgBytesPerRead_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) GetPropertyAvgBytesPerRead_Base()(value uint32, err error) { 
    retValue, err := instance.GetProperty("AvgBytesPerRead_Base")
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

// SetAvgBytesPerTransfer sets the value of AvgBytesPerTransfer for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) SetPropertyAvgBytesPerTransfer(value uint64) (err error) { 
    return instance.SetProperty("AvgBytesPerTransfer", (value))
}

// GetAvgBytesPerTransfer gets the value of AvgBytesPerTransfer for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) GetPropertyAvgBytesPerTransfer()(value uint64, err error) { 
    retValue, err := instance.GetProperty("AvgBytesPerTransfer")
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

// SetAvgBytesPerTransfer_Base sets the value of AvgBytesPerTransfer_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) SetPropertyAvgBytesPerTransfer_Base(value uint32) (err error) { 
    return instance.SetProperty("AvgBytesPerTransfer_Base", (value))
}

// GetAvgBytesPerTransfer_Base gets the value of AvgBytesPerTransfer_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) GetPropertyAvgBytesPerTransfer_Base()(value uint32, err error) { 
    retValue, err := instance.GetProperty("AvgBytesPerTransfer_Base")
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

// SetAvgBytesPerWrite sets the value of AvgBytesPerWrite for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) SetPropertyAvgBytesPerWrite(value uint64) (err error) { 
    return instance.SetProperty("AvgBytesPerWrite", (value))
}

// GetAvgBytesPerWrite gets the value of AvgBytesPerWrite for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) GetPropertyAvgBytesPerWrite()(value uint64, err error) { 
    retValue, err := instance.GetProperty("AvgBytesPerWrite")
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

// SetAvgBytesPerWrite_Base sets the value of AvgBytesPerWrite_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) SetPropertyAvgBytesPerWrite_Base(value uint32) (err error) { 
    return instance.SetProperty("AvgBytesPerWrite_Base", (value))
}

// GetAvgBytesPerWrite_Base gets the value of AvgBytesPerWrite_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) GetPropertyAvgBytesPerWrite_Base()(value uint32, err error) { 
    retValue, err := instance.GetProperty("AvgBytesPerWrite_Base")
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

// SetAvgmicrosecPerRead sets the value of AvgmicrosecPerRead for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) SetPropertyAvgmicrosecPerRead(value uint64) (err error) { 
    return instance.SetProperty("AvgmicrosecPerRead", (value))
}

// GetAvgmicrosecPerRead gets the value of AvgmicrosecPerRead for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) GetPropertyAvgmicrosecPerRead()(value uint64, err error) { 
    retValue, err := instance.GetProperty("AvgmicrosecPerRead")
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

// SetAvgmicrosecPerRead_Base sets the value of AvgmicrosecPerRead_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) SetPropertyAvgmicrosecPerRead_Base(value uint32) (err error) { 
    return instance.SetProperty("AvgmicrosecPerRead_Base", (value))
}

// GetAvgmicrosecPerRead_Base gets the value of AvgmicrosecPerRead_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) GetPropertyAvgmicrosecPerRead_Base()(value uint32, err error) { 
    retValue, err := instance.GetProperty("AvgmicrosecPerRead_Base")
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

// SetAvgmicrosecPerTransfer sets the value of AvgmicrosecPerTransfer for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) SetPropertyAvgmicrosecPerTransfer(value uint64) (err error) { 
    return instance.SetProperty("AvgmicrosecPerTransfer", (value))
}

// GetAvgmicrosecPerTransfer gets the value of AvgmicrosecPerTransfer for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) GetPropertyAvgmicrosecPerTransfer()(value uint64, err error) { 
    retValue, err := instance.GetProperty("AvgmicrosecPerTransfer")
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

// SetAvgmicrosecPerTransfer_Base sets the value of AvgmicrosecPerTransfer_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) SetPropertyAvgmicrosecPerTransfer_Base(value uint32) (err error) { 
    return instance.SetProperty("AvgmicrosecPerTransfer_Base", (value))
}

// GetAvgmicrosecPerTransfer_Base gets the value of AvgmicrosecPerTransfer_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) GetPropertyAvgmicrosecPerTransfer_Base()(value uint32, err error) { 
    retValue, err := instance.GetProperty("AvgmicrosecPerTransfer_Base")
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

// SetAvgmicrosecPerWrite sets the value of AvgmicrosecPerWrite for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) SetPropertyAvgmicrosecPerWrite(value uint64) (err error) { 
    return instance.SetProperty("AvgmicrosecPerWrite", (value))
}

// GetAvgmicrosecPerWrite gets the value of AvgmicrosecPerWrite for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) GetPropertyAvgmicrosecPerWrite()(value uint64, err error) { 
    retValue, err := instance.GetProperty("AvgmicrosecPerWrite")
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

// SetAvgmicrosecPerWrite_Base sets the value of AvgmicrosecPerWrite_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) SetPropertyAvgmicrosecPerWrite_Base(value uint32) (err error) { 
    return instance.SetProperty("AvgmicrosecPerWrite_Base", (value))
}

// GetAvgmicrosecPerWrite_Base gets the value of AvgmicrosecPerWrite_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) GetPropertyAvgmicrosecPerWrite_Base()(value uint32, err error) { 
    retValue, err := instance.GetProperty("AvgmicrosecPerWrite_Base")
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

// SetHTTPStorageIOretryPersec sets the value of HTTPStorageIOretryPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) SetPropertyHTTPStorageIOretryPersec(value uint64) (err error) { 
    return instance.SetProperty("HTTPStorageIOretryPersec", (value))
}

// GetHTTPStorageIOretryPersec gets the value of HTTPStorageIOretryPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) GetPropertyHTTPStorageIOretryPersec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("HTTPStorageIOretryPersec")
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

// SetOutstandingHTTPStorageIO sets the value of OutstandingHTTPStorageIO for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) SetPropertyOutstandingHTTPStorageIO(value uint64) (err error) { 
    return instance.SetProperty("OutstandingHTTPStorageIO", (value))
}

// GetOutstandingHTTPStorageIO gets the value of OutstandingHTTPStorageIO for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) GetPropertyOutstandingHTTPStorageIO()(value uint64, err error) { 
    retValue, err := instance.GetProperty("OutstandingHTTPStorageIO")
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

// SetReadBytesPerSec sets the value of ReadBytesPerSec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) SetPropertyReadBytesPerSec(value uint64) (err error) { 
    return instance.SetProperty("ReadBytesPerSec", (value))
}

// GetReadBytesPerSec gets the value of ReadBytesPerSec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) GetPropertyReadBytesPerSec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("ReadBytesPerSec")
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

// SetReadsPerSec sets the value of ReadsPerSec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) SetPropertyReadsPerSec(value uint64) (err error) { 
    return instance.SetProperty("ReadsPerSec", (value))
}

// GetReadsPerSec gets the value of ReadsPerSec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) GetPropertyReadsPerSec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("ReadsPerSec")
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

// SetTotalBytesPerSec sets the value of TotalBytesPerSec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) SetPropertyTotalBytesPerSec(value uint64) (err error) { 
    return instance.SetProperty("TotalBytesPerSec", (value))
}

// GetTotalBytesPerSec gets the value of TotalBytesPerSec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) GetPropertyTotalBytesPerSec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("TotalBytesPerSec")
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

// SetTransfersPerSec sets the value of TransfersPerSec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) SetPropertyTransfersPerSec(value uint64) (err error) { 
    return instance.SetProperty("TransfersPerSec", (value))
}

// GetTransfersPerSec gets the value of TransfersPerSec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) GetPropertyTransfersPerSec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("TransfersPerSec")
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

// SetWriteBytesPerSec sets the value of WriteBytesPerSec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) SetPropertyWriteBytesPerSec(value uint64) (err error) { 
    return instance.SetProperty("WriteBytesPerSec", (value))
}

// GetWriteBytesPerSec gets the value of WriteBytesPerSec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) GetPropertyWriteBytesPerSec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("WriteBytesPerSec")
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

// SetWritesPerSec sets the value of WritesPerSec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) SetPropertyWritesPerSec(value uint64) (err error) { 
    return instance.SetProperty("WritesPerSec", (value))
}

// GetWritesPerSec gets the value of WritesPerSec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDHTTPStorage) GetPropertyWritesPerSec()(value uint64, err error) { 
    retValue, err := instance.GetProperty("WritesPerSec")
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

