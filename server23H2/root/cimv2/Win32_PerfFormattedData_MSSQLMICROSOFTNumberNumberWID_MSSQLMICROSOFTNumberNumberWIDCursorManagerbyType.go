// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.CIMV2
//
// ////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCursorManagerbyType struct
type Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCursorManagerbyType struct {
	*Win32_PerfFormattedData

	//
	Activecursors uint64

	//
	CachedCursorCounts uint64

	//
	CacheHitRatio uint64

	//
	CursorCacheUseCountsPersec uint64

	//
	Cursormemoryusage uint64

	//
	CursorRequestsPersec uint64

	//
	Cursorworktableusage uint64

	//
	Numberofactivecursorplans uint64
}

func NewWin32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCursorManagerbyTypeEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCursorManagerbyType, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCursorManagerbyType{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCursorManagerbyTypeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCursorManagerbyType, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCursorManagerbyType{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetActivecursors sets the value of Activecursors for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCursorManagerbyType) SetPropertyActivecursors(value uint64) (err error) {
	return instance.SetProperty("Activecursors", (value))
}

// GetActivecursors gets the value of Activecursors for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCursorManagerbyType) GetPropertyActivecursors() (value uint64, err error) {
	retValue, err := instance.GetProperty("Activecursors")
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

// SetCachedCursorCounts sets the value of CachedCursorCounts for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCursorManagerbyType) SetPropertyCachedCursorCounts(value uint64) (err error) {
	return instance.SetProperty("CachedCursorCounts", (value))
}

// GetCachedCursorCounts gets the value of CachedCursorCounts for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCursorManagerbyType) GetPropertyCachedCursorCounts() (value uint64, err error) {
	retValue, err := instance.GetProperty("CachedCursorCounts")
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

// SetCacheHitRatio sets the value of CacheHitRatio for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCursorManagerbyType) SetPropertyCacheHitRatio(value uint64) (err error) {
	return instance.SetProperty("CacheHitRatio", (value))
}

// GetCacheHitRatio gets the value of CacheHitRatio for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCursorManagerbyType) GetPropertyCacheHitRatio() (value uint64, err error) {
	retValue, err := instance.GetProperty("CacheHitRatio")
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

// SetCursorCacheUseCountsPersec sets the value of CursorCacheUseCountsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCursorManagerbyType) SetPropertyCursorCacheUseCountsPersec(value uint64) (err error) {
	return instance.SetProperty("CursorCacheUseCountsPersec", (value))
}

// GetCursorCacheUseCountsPersec gets the value of CursorCacheUseCountsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCursorManagerbyType) GetPropertyCursorCacheUseCountsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("CursorCacheUseCountsPersec")
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

// SetCursormemoryusage sets the value of Cursormemoryusage for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCursorManagerbyType) SetPropertyCursormemoryusage(value uint64) (err error) {
	return instance.SetProperty("Cursormemoryusage", (value))
}

// GetCursormemoryusage gets the value of Cursormemoryusage for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCursorManagerbyType) GetPropertyCursormemoryusage() (value uint64, err error) {
	retValue, err := instance.GetProperty("Cursormemoryusage")
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

// SetCursorRequestsPersec sets the value of CursorRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCursorManagerbyType) SetPropertyCursorRequestsPersec(value uint64) (err error) {
	return instance.SetProperty("CursorRequestsPersec", (value))
}

// GetCursorRequestsPersec gets the value of CursorRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCursorManagerbyType) GetPropertyCursorRequestsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("CursorRequestsPersec")
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

// SetCursorworktableusage sets the value of Cursorworktableusage for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCursorManagerbyType) SetPropertyCursorworktableusage(value uint64) (err error) {
	return instance.SetProperty("Cursorworktableusage", (value))
}

// GetCursorworktableusage gets the value of Cursorworktableusage for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCursorManagerbyType) GetPropertyCursorworktableusage() (value uint64, err error) {
	retValue, err := instance.GetProperty("Cursorworktableusage")
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

// SetNumberofactivecursorplans sets the value of Numberofactivecursorplans for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCursorManagerbyType) SetPropertyNumberofactivecursorplans(value uint64) (err error) {
	return instance.SetProperty("Numberofactivecursorplans", (value))
}

// GetNumberofactivecursorplans gets the value of Numberofactivecursorplans for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDCursorManagerbyType) GetPropertyNumberofactivecursorplans() (value uint64, err error) {
	retValue, err := instance.GetProperty("Numberofactivecursorplans")
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
