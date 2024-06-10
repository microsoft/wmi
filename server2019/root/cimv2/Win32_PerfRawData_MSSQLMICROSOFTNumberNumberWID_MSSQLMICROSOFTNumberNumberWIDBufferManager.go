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

// Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager struct
type Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager struct {
	*Win32_PerfRawData

	//
	BackgroundwriterpagesPersec uint64

	//
	Buffercachehitratio uint64

	//
	Buffercachehitratio_Base uint64

	//
	CheckpointpagesPersec uint64

	//
	Databasepages uint64

	//
	Extensionallocatedpages uint64

	//
	Extensionfreepages uint64

	//
	Extensioninuseaspercentage uint64

	//
	ExtensionoutstandingIOcounter uint64

	//
	ExtensionpageevictionsPersec uint64

	//
	ExtensionpagereadsPersec uint64

	//
	Extensionpageunreferencedtime uint64

	//
	ExtensionpagewritesPersec uint64

	//
	FreeliststallsPersec uint64

	//
	IntegralControllerSlope uint64

	//
	LazywritesPersec uint64

	//
	Pagelifeexpectancy uint64

	//
	PagelookupsPersec uint64

	//
	PagereadsPersec uint64

	//
	PagewritesPersec uint64

	//
	ReadaheadpagesPersec uint64

	//
	ReadaheadtimePersec uint64

	//
	Targetpages uint64
}

func NewWin32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManagerEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManagerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetBackgroundwriterpagesPersec sets the value of BackgroundwriterpagesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) SetPropertyBackgroundwriterpagesPersec(value uint64) (err error) {
	return instance.SetProperty("BackgroundwriterpagesPersec", (value))
}

// GetBackgroundwriterpagesPersec gets the value of BackgroundwriterpagesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) GetPropertyBackgroundwriterpagesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("BackgroundwriterpagesPersec")
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

// SetBuffercachehitratio sets the value of Buffercachehitratio for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) SetPropertyBuffercachehitratio(value uint64) (err error) {
	return instance.SetProperty("Buffercachehitratio", (value))
}

// GetBuffercachehitratio gets the value of Buffercachehitratio for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) GetPropertyBuffercachehitratio() (value uint64, err error) {
	retValue, err := instance.GetProperty("Buffercachehitratio")
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

// SetBuffercachehitratio_Base sets the value of Buffercachehitratio_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) SetPropertyBuffercachehitratio_Base(value uint64) (err error) {
	return instance.SetProperty("Buffercachehitratio_Base", (value))
}

// GetBuffercachehitratio_Base gets the value of Buffercachehitratio_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) GetPropertyBuffercachehitratio_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("Buffercachehitratio_Base")
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

// SetCheckpointpagesPersec sets the value of CheckpointpagesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) SetPropertyCheckpointpagesPersec(value uint64) (err error) {
	return instance.SetProperty("CheckpointpagesPersec", (value))
}

// GetCheckpointpagesPersec gets the value of CheckpointpagesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) GetPropertyCheckpointpagesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("CheckpointpagesPersec")
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

// SetDatabasepages sets the value of Databasepages for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) SetPropertyDatabasepages(value uint64) (err error) {
	return instance.SetProperty("Databasepages", (value))
}

// GetDatabasepages gets the value of Databasepages for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) GetPropertyDatabasepages() (value uint64, err error) {
	retValue, err := instance.GetProperty("Databasepages")
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

// SetExtensionallocatedpages sets the value of Extensionallocatedpages for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) SetPropertyExtensionallocatedpages(value uint64) (err error) {
	return instance.SetProperty("Extensionallocatedpages", (value))
}

// GetExtensionallocatedpages gets the value of Extensionallocatedpages for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) GetPropertyExtensionallocatedpages() (value uint64, err error) {
	retValue, err := instance.GetProperty("Extensionallocatedpages")
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

// SetExtensionfreepages sets the value of Extensionfreepages for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) SetPropertyExtensionfreepages(value uint64) (err error) {
	return instance.SetProperty("Extensionfreepages", (value))
}

// GetExtensionfreepages gets the value of Extensionfreepages for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) GetPropertyExtensionfreepages() (value uint64, err error) {
	retValue, err := instance.GetProperty("Extensionfreepages")
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

// SetExtensioninuseaspercentage sets the value of Extensioninuseaspercentage for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) SetPropertyExtensioninuseaspercentage(value uint64) (err error) {
	return instance.SetProperty("Extensioninuseaspercentage", (value))
}

// GetExtensioninuseaspercentage gets the value of Extensioninuseaspercentage for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) GetPropertyExtensioninuseaspercentage() (value uint64, err error) {
	retValue, err := instance.GetProperty("Extensioninuseaspercentage")
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

// SetExtensionoutstandingIOcounter sets the value of ExtensionoutstandingIOcounter for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) SetPropertyExtensionoutstandingIOcounter(value uint64) (err error) {
	return instance.SetProperty("ExtensionoutstandingIOcounter", (value))
}

// GetExtensionoutstandingIOcounter gets the value of ExtensionoutstandingIOcounter for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) GetPropertyExtensionoutstandingIOcounter() (value uint64, err error) {
	retValue, err := instance.GetProperty("ExtensionoutstandingIOcounter")
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

// SetExtensionpageevictionsPersec sets the value of ExtensionpageevictionsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) SetPropertyExtensionpageevictionsPersec(value uint64) (err error) {
	return instance.SetProperty("ExtensionpageevictionsPersec", (value))
}

// GetExtensionpageevictionsPersec gets the value of ExtensionpageevictionsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) GetPropertyExtensionpageevictionsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ExtensionpageevictionsPersec")
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

// SetExtensionpagereadsPersec sets the value of ExtensionpagereadsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) SetPropertyExtensionpagereadsPersec(value uint64) (err error) {
	return instance.SetProperty("ExtensionpagereadsPersec", (value))
}

// GetExtensionpagereadsPersec gets the value of ExtensionpagereadsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) GetPropertyExtensionpagereadsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ExtensionpagereadsPersec")
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

// SetExtensionpageunreferencedtime sets the value of Extensionpageunreferencedtime for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) SetPropertyExtensionpageunreferencedtime(value uint64) (err error) {
	return instance.SetProperty("Extensionpageunreferencedtime", (value))
}

// GetExtensionpageunreferencedtime gets the value of Extensionpageunreferencedtime for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) GetPropertyExtensionpageunreferencedtime() (value uint64, err error) {
	retValue, err := instance.GetProperty("Extensionpageunreferencedtime")
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

// SetExtensionpagewritesPersec sets the value of ExtensionpagewritesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) SetPropertyExtensionpagewritesPersec(value uint64) (err error) {
	return instance.SetProperty("ExtensionpagewritesPersec", (value))
}

// GetExtensionpagewritesPersec gets the value of ExtensionpagewritesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) GetPropertyExtensionpagewritesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ExtensionpagewritesPersec")
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

// SetFreeliststallsPersec sets the value of FreeliststallsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) SetPropertyFreeliststallsPersec(value uint64) (err error) {
	return instance.SetProperty("FreeliststallsPersec", (value))
}

// GetFreeliststallsPersec gets the value of FreeliststallsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) GetPropertyFreeliststallsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("FreeliststallsPersec")
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

// SetIntegralControllerSlope sets the value of IntegralControllerSlope for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) SetPropertyIntegralControllerSlope(value uint64) (err error) {
	return instance.SetProperty("IntegralControllerSlope", (value))
}

// GetIntegralControllerSlope gets the value of IntegralControllerSlope for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) GetPropertyIntegralControllerSlope() (value uint64, err error) {
	retValue, err := instance.GetProperty("IntegralControllerSlope")
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

// SetLazywritesPersec sets the value of LazywritesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) SetPropertyLazywritesPersec(value uint64) (err error) {
	return instance.SetProperty("LazywritesPersec", (value))
}

// GetLazywritesPersec gets the value of LazywritesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) GetPropertyLazywritesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("LazywritesPersec")
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

// SetPagelifeexpectancy sets the value of Pagelifeexpectancy for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) SetPropertyPagelifeexpectancy(value uint64) (err error) {
	return instance.SetProperty("Pagelifeexpectancy", (value))
}

// GetPagelifeexpectancy gets the value of Pagelifeexpectancy for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) GetPropertyPagelifeexpectancy() (value uint64, err error) {
	retValue, err := instance.GetProperty("Pagelifeexpectancy")
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

// SetPagelookupsPersec sets the value of PagelookupsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) SetPropertyPagelookupsPersec(value uint64) (err error) {
	return instance.SetProperty("PagelookupsPersec", (value))
}

// GetPagelookupsPersec gets the value of PagelookupsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) GetPropertyPagelookupsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("PagelookupsPersec")
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

// SetPagereadsPersec sets the value of PagereadsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) SetPropertyPagereadsPersec(value uint64) (err error) {
	return instance.SetProperty("PagereadsPersec", (value))
}

// GetPagereadsPersec gets the value of PagereadsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) GetPropertyPagereadsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("PagereadsPersec")
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

// SetPagewritesPersec sets the value of PagewritesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) SetPropertyPagewritesPersec(value uint64) (err error) {
	return instance.SetProperty("PagewritesPersec", (value))
}

// GetPagewritesPersec gets the value of PagewritesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) GetPropertyPagewritesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("PagewritesPersec")
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

// SetReadaheadpagesPersec sets the value of ReadaheadpagesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) SetPropertyReadaheadpagesPersec(value uint64) (err error) {
	return instance.SetProperty("ReadaheadpagesPersec", (value))
}

// GetReadaheadpagesPersec gets the value of ReadaheadpagesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) GetPropertyReadaheadpagesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadaheadpagesPersec")
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

// SetReadaheadtimePersec sets the value of ReadaheadtimePersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) SetPropertyReadaheadtimePersec(value uint64) (err error) {
	return instance.SetProperty("ReadaheadtimePersec", (value))
}

// GetReadaheadtimePersec gets the value of ReadaheadtimePersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) GetPropertyReadaheadtimePersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadaheadtimePersec")
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

// SetTargetpages sets the value of Targetpages for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) SetPropertyTargetpages(value uint64) (err error) {
	return instance.SetProperty("Targetpages", (value))
}

// GetTargetpages gets the value of Targetpages for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBufferManager) GetPropertyTargetpages() (value uint64, err error) {
	retValue, err := instance.GetProperty("Targetpages")
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
