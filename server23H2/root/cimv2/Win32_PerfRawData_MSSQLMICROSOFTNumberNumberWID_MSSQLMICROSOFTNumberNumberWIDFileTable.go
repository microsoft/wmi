// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable struct
type Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable struct {
	*Win32_PerfRawData

	//
	AvgtimedeleteFileTableitem uint64

	//
	AvgtimedeleteFileTableitem_Base uint32

	//
	AvgtimeFileTableenumeration uint64

	//
	AvgtimeFileTableenumeration_Base uint32

	//
	AvgtimeFileTablehandlekill uint64

	//
	AvgtimeFileTablehandlekill_Base uint32

	//
	AvgtimemoveFileTableitem uint64

	//
	AvgtimemoveFileTableitem_Base uint32

	//
	AvgtimeperfileIOrequest uint64

	//
	AvgtimeperfileIOrequest_Base uint32

	//
	AvgtimeperfileIOresponse uint64

	//
	AvgtimeperfileIOresponse_Base uint32

	//
	AvgtimerenameFileTableitem uint64

	//
	AvgtimerenameFileTableitem_Base uint32

	//
	AvgtimetogetFileTableitem uint64

	//
	AvgtimetogetFileTableitem_Base uint32

	//
	AvgtimeupdateFileTableitem uint64

	//
	AvgtimeupdateFileTableitem_Base uint32

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

func NewWin32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTableEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTableEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetAvgtimedeleteFileTableitem sets the value of AvgtimedeleteFileTableitem for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyAvgtimedeleteFileTableitem(value uint64) (err error) {
	return instance.SetProperty("AvgtimedeleteFileTableitem", (value))
}

// GetAvgtimedeleteFileTableitem gets the value of AvgtimedeleteFileTableitem for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyAvgtimedeleteFileTableitem() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgtimedeleteFileTableitem")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetAvgtimedeleteFileTableitem_Base sets the value of AvgtimedeleteFileTableitem_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyAvgtimedeleteFileTableitem_Base(value uint32) (err error) {
	return instance.SetProperty("AvgtimedeleteFileTableitem_Base", (value))
}

// GetAvgtimedeleteFileTableitem_Base gets the value of AvgtimedeleteFileTableitem_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyAvgtimedeleteFileTableitem_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("AvgtimedeleteFileTableitem_Base")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetAvgtimeFileTableenumeration sets the value of AvgtimeFileTableenumeration for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyAvgtimeFileTableenumeration(value uint64) (err error) {
	return instance.SetProperty("AvgtimeFileTableenumeration", (value))
}

// GetAvgtimeFileTableenumeration gets the value of AvgtimeFileTableenumeration for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyAvgtimeFileTableenumeration() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgtimeFileTableenumeration")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetAvgtimeFileTableenumeration_Base sets the value of AvgtimeFileTableenumeration_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyAvgtimeFileTableenumeration_Base(value uint32) (err error) {
	return instance.SetProperty("AvgtimeFileTableenumeration_Base", (value))
}

// GetAvgtimeFileTableenumeration_Base gets the value of AvgtimeFileTableenumeration_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyAvgtimeFileTableenumeration_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("AvgtimeFileTableenumeration_Base")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetAvgtimeFileTablehandlekill sets the value of AvgtimeFileTablehandlekill for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyAvgtimeFileTablehandlekill(value uint64) (err error) {
	return instance.SetProperty("AvgtimeFileTablehandlekill", (value))
}

// GetAvgtimeFileTablehandlekill gets the value of AvgtimeFileTablehandlekill for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyAvgtimeFileTablehandlekill() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgtimeFileTablehandlekill")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetAvgtimeFileTablehandlekill_Base sets the value of AvgtimeFileTablehandlekill_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyAvgtimeFileTablehandlekill_Base(value uint32) (err error) {
	return instance.SetProperty("AvgtimeFileTablehandlekill_Base", (value))
}

// GetAvgtimeFileTablehandlekill_Base gets the value of AvgtimeFileTablehandlekill_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyAvgtimeFileTablehandlekill_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("AvgtimeFileTablehandlekill_Base")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetAvgtimemoveFileTableitem sets the value of AvgtimemoveFileTableitem for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyAvgtimemoveFileTableitem(value uint64) (err error) {
	return instance.SetProperty("AvgtimemoveFileTableitem", (value))
}

// GetAvgtimemoveFileTableitem gets the value of AvgtimemoveFileTableitem for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyAvgtimemoveFileTableitem() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgtimemoveFileTableitem")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetAvgtimemoveFileTableitem_Base sets the value of AvgtimemoveFileTableitem_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyAvgtimemoveFileTableitem_Base(value uint32) (err error) {
	return instance.SetProperty("AvgtimemoveFileTableitem_Base", (value))
}

// GetAvgtimemoveFileTableitem_Base gets the value of AvgtimemoveFileTableitem_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyAvgtimemoveFileTableitem_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("AvgtimemoveFileTableitem_Base")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetAvgtimeperfileIOrequest sets the value of AvgtimeperfileIOrequest for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyAvgtimeperfileIOrequest(value uint64) (err error) {
	return instance.SetProperty("AvgtimeperfileIOrequest", (value))
}

// GetAvgtimeperfileIOrequest gets the value of AvgtimeperfileIOrequest for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyAvgtimeperfileIOrequest() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgtimeperfileIOrequest")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetAvgtimeperfileIOrequest_Base sets the value of AvgtimeperfileIOrequest_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyAvgtimeperfileIOrequest_Base(value uint32) (err error) {
	return instance.SetProperty("AvgtimeperfileIOrequest_Base", (value))
}

// GetAvgtimeperfileIOrequest_Base gets the value of AvgtimeperfileIOrequest_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyAvgtimeperfileIOrequest_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("AvgtimeperfileIOrequest_Base")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetAvgtimeperfileIOresponse sets the value of AvgtimeperfileIOresponse for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyAvgtimeperfileIOresponse(value uint64) (err error) {
	return instance.SetProperty("AvgtimeperfileIOresponse", (value))
}

// GetAvgtimeperfileIOresponse gets the value of AvgtimeperfileIOresponse for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyAvgtimeperfileIOresponse() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgtimeperfileIOresponse")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetAvgtimeperfileIOresponse_Base sets the value of AvgtimeperfileIOresponse_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyAvgtimeperfileIOresponse_Base(value uint32) (err error) {
	return instance.SetProperty("AvgtimeperfileIOresponse_Base", (value))
}

// GetAvgtimeperfileIOresponse_Base gets the value of AvgtimeperfileIOresponse_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyAvgtimeperfileIOresponse_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("AvgtimeperfileIOresponse_Base")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetAvgtimerenameFileTableitem sets the value of AvgtimerenameFileTableitem for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyAvgtimerenameFileTableitem(value uint64) (err error) {
	return instance.SetProperty("AvgtimerenameFileTableitem", (value))
}

// GetAvgtimerenameFileTableitem gets the value of AvgtimerenameFileTableitem for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyAvgtimerenameFileTableitem() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgtimerenameFileTableitem")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetAvgtimerenameFileTableitem_Base sets the value of AvgtimerenameFileTableitem_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyAvgtimerenameFileTableitem_Base(value uint32) (err error) {
	return instance.SetProperty("AvgtimerenameFileTableitem_Base", (value))
}

// GetAvgtimerenameFileTableitem_Base gets the value of AvgtimerenameFileTableitem_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyAvgtimerenameFileTableitem_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("AvgtimerenameFileTableitem_Base")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetAvgtimetogetFileTableitem sets the value of AvgtimetogetFileTableitem for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyAvgtimetogetFileTableitem(value uint64) (err error) {
	return instance.SetProperty("AvgtimetogetFileTableitem", (value))
}

// GetAvgtimetogetFileTableitem gets the value of AvgtimetogetFileTableitem for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyAvgtimetogetFileTableitem() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgtimetogetFileTableitem")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetAvgtimetogetFileTableitem_Base sets the value of AvgtimetogetFileTableitem_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyAvgtimetogetFileTableitem_Base(value uint32) (err error) {
	return instance.SetProperty("AvgtimetogetFileTableitem_Base", (value))
}

// GetAvgtimetogetFileTableitem_Base gets the value of AvgtimetogetFileTableitem_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyAvgtimetogetFileTableitem_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("AvgtimetogetFileTableitem_Base")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetAvgtimeupdateFileTableitem sets the value of AvgtimeupdateFileTableitem for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyAvgtimeupdateFileTableitem(value uint64) (err error) {
	return instance.SetProperty("AvgtimeupdateFileTableitem", (value))
}

// GetAvgtimeupdateFileTableitem gets the value of AvgtimeupdateFileTableitem for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyAvgtimeupdateFileTableitem() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvgtimeupdateFileTableitem")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetAvgtimeupdateFileTableitem_Base sets the value of AvgtimeupdateFileTableitem_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyAvgtimeupdateFileTableitem_Base(value uint32) (err error) {
	return instance.SetProperty("AvgtimeupdateFileTableitem_Base", (value))
}

// GetAvgtimeupdateFileTableitem_Base gets the value of AvgtimeupdateFileTableitem_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyAvgtimeupdateFileTableitem_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("AvgtimeupdateFileTableitem_Base")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetFileTabledboperationsPersec sets the value of FileTabledboperationsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyFileTabledboperationsPersec(value uint64) (err error) {
	return instance.SetProperty("FileTabledboperationsPersec", (value))
}

// GetFileTabledboperationsPersec gets the value of FileTabledboperationsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyFileTabledboperationsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("FileTabledboperationsPersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetFileTableenumerationreqsPersec sets the value of FileTableenumerationreqsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyFileTableenumerationreqsPersec(value uint64) (err error) {
	return instance.SetProperty("FileTableenumerationreqsPersec", (value))
}

// GetFileTableenumerationreqsPersec gets the value of FileTableenumerationreqsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyFileTableenumerationreqsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("FileTableenumerationreqsPersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetFileTablefileIOrequestsPersec sets the value of FileTablefileIOrequestsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyFileTablefileIOrequestsPersec(value uint64) (err error) {
	return instance.SetProperty("FileTablefileIOrequestsPersec", (value))
}

// GetFileTablefileIOrequestsPersec gets the value of FileTablefileIOrequestsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyFileTablefileIOrequestsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("FileTablefileIOrequestsPersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetFileTablefileIOresponsePersec sets the value of FileTablefileIOresponsePersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyFileTablefileIOresponsePersec(value uint64) (err error) {
	return instance.SetProperty("FileTablefileIOresponsePersec", (value))
}

// GetFileTablefileIOresponsePersec gets the value of FileTablefileIOresponsePersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyFileTablefileIOresponsePersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("FileTablefileIOresponsePersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetFileTableitemdeletereqsPersec sets the value of FileTableitemdeletereqsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyFileTableitemdeletereqsPersec(value uint64) (err error) {
	return instance.SetProperty("FileTableitemdeletereqsPersec", (value))
}

// GetFileTableitemdeletereqsPersec gets the value of FileTableitemdeletereqsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyFileTableitemdeletereqsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("FileTableitemdeletereqsPersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetFileTableitemgetrequestsPersec sets the value of FileTableitemgetrequestsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyFileTableitemgetrequestsPersec(value uint64) (err error) {
	return instance.SetProperty("FileTableitemgetrequestsPersec", (value))
}

// GetFileTableitemgetrequestsPersec gets the value of FileTableitemgetrequestsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyFileTableitemgetrequestsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("FileTableitemgetrequestsPersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetFileTableitemmovereqsPersec sets the value of FileTableitemmovereqsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyFileTableitemmovereqsPersec(value uint64) (err error) {
	return instance.SetProperty("FileTableitemmovereqsPersec", (value))
}

// GetFileTableitemmovereqsPersec gets the value of FileTableitemmovereqsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyFileTableitemmovereqsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("FileTableitemmovereqsPersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetFileTableitemrenamereqsPersec sets the value of FileTableitemrenamereqsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyFileTableitemrenamereqsPersec(value uint64) (err error) {
	return instance.SetProperty("FileTableitemrenamereqsPersec", (value))
}

// GetFileTableitemrenamereqsPersec gets the value of FileTableitemrenamereqsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyFileTableitemrenamereqsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("FileTableitemrenamereqsPersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetFileTableitemupdatereqsPersec sets the value of FileTableitemupdatereqsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyFileTableitemupdatereqsPersec(value uint64) (err error) {
	return instance.SetProperty("FileTableitemupdatereqsPersec", (value))
}

// GetFileTableitemupdatereqsPersec gets the value of FileTableitemupdatereqsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyFileTableitemupdatereqsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("FileTableitemupdatereqsPersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetFileTablekillhandleopsPersec sets the value of FileTablekillhandleopsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyFileTablekillhandleopsPersec(value uint64) (err error) {
	return instance.SetProperty("FileTablekillhandleopsPersec", (value))
}

// GetFileTablekillhandleopsPersec gets the value of FileTablekillhandleopsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyFileTablekillhandleopsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("FileTablekillhandleopsPersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetFileTabletableoperationsPersec sets the value of FileTabletableoperationsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) SetPropertyFileTabletableoperationsPersec(value uint64) (err error) {
	return instance.SetProperty("FileTabletableoperationsPersec", (value))
}

// GetFileTabletableoperationsPersec gets the value of FileTabletableoperationsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDFileTable) GetPropertyFileTabletableoperationsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("FileTabletableoperationsPersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}
