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

// Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLStatistics struct
type Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLStatistics struct {
	*Win32_PerfRawData

	//
	AutoParamAttemptsPersec uint64

	//
	BatchRequestsPersec uint64

	//
	FailedAutoParamsPersec uint64

	//
	ForcedParameterizationsPersec uint64

	//
	GuidedplanexecutionsPersec uint64

	//
	MisguidedplanexecutionsPersec uint64

	//
	SafeAutoParamsPersec uint64

	//
	SQLAttentionrate uint64

	//
	SQLCompilationsPersec uint64

	//
	SQLReCompilationsPersec uint64

	//
	UnsafeAutoParamsPersec uint64
}

func NewWin32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLStatisticsEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLStatistics, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLStatistics{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLStatisticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLStatistics, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLStatistics{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetAutoParamAttemptsPersec sets the value of AutoParamAttemptsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLStatistics) SetPropertyAutoParamAttemptsPersec(value uint64) (err error) {
	return instance.SetProperty("AutoParamAttemptsPersec", (value))
}

// GetAutoParamAttemptsPersec gets the value of AutoParamAttemptsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLStatistics) GetPropertyAutoParamAttemptsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("AutoParamAttemptsPersec")
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

// SetBatchRequestsPersec sets the value of BatchRequestsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLStatistics) SetPropertyBatchRequestsPersec(value uint64) (err error) {
	return instance.SetProperty("BatchRequestsPersec", (value))
}

// GetBatchRequestsPersec gets the value of BatchRequestsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLStatistics) GetPropertyBatchRequestsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("BatchRequestsPersec")
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

// SetFailedAutoParamsPersec sets the value of FailedAutoParamsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLStatistics) SetPropertyFailedAutoParamsPersec(value uint64) (err error) {
	return instance.SetProperty("FailedAutoParamsPersec", (value))
}

// GetFailedAutoParamsPersec gets the value of FailedAutoParamsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLStatistics) GetPropertyFailedAutoParamsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("FailedAutoParamsPersec")
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

// SetForcedParameterizationsPersec sets the value of ForcedParameterizationsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLStatistics) SetPropertyForcedParameterizationsPersec(value uint64) (err error) {
	return instance.SetProperty("ForcedParameterizationsPersec", (value))
}

// GetForcedParameterizationsPersec gets the value of ForcedParameterizationsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLStatistics) GetPropertyForcedParameterizationsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ForcedParameterizationsPersec")
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

// SetGuidedplanexecutionsPersec sets the value of GuidedplanexecutionsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLStatistics) SetPropertyGuidedplanexecutionsPersec(value uint64) (err error) {
	return instance.SetProperty("GuidedplanexecutionsPersec", (value))
}

// GetGuidedplanexecutionsPersec gets the value of GuidedplanexecutionsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLStatistics) GetPropertyGuidedplanexecutionsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("GuidedplanexecutionsPersec")
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

// SetMisguidedplanexecutionsPersec sets the value of MisguidedplanexecutionsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLStatistics) SetPropertyMisguidedplanexecutionsPersec(value uint64) (err error) {
	return instance.SetProperty("MisguidedplanexecutionsPersec", (value))
}

// GetMisguidedplanexecutionsPersec gets the value of MisguidedplanexecutionsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLStatistics) GetPropertyMisguidedplanexecutionsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("MisguidedplanexecutionsPersec")
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

// SetSafeAutoParamsPersec sets the value of SafeAutoParamsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLStatistics) SetPropertySafeAutoParamsPersec(value uint64) (err error) {
	return instance.SetProperty("SafeAutoParamsPersec", (value))
}

// GetSafeAutoParamsPersec gets the value of SafeAutoParamsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLStatistics) GetPropertySafeAutoParamsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("SafeAutoParamsPersec")
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

// SetSQLAttentionrate sets the value of SQLAttentionrate for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLStatistics) SetPropertySQLAttentionrate(value uint64) (err error) {
	return instance.SetProperty("SQLAttentionrate", (value))
}

// GetSQLAttentionrate gets the value of SQLAttentionrate for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLStatistics) GetPropertySQLAttentionrate() (value uint64, err error) {
	retValue, err := instance.GetProperty("SQLAttentionrate")
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

// SetSQLCompilationsPersec sets the value of SQLCompilationsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLStatistics) SetPropertySQLCompilationsPersec(value uint64) (err error) {
	return instance.SetProperty("SQLCompilationsPersec", (value))
}

// GetSQLCompilationsPersec gets the value of SQLCompilationsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLStatistics) GetPropertySQLCompilationsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("SQLCompilationsPersec")
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

// SetSQLReCompilationsPersec sets the value of SQLReCompilationsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLStatistics) SetPropertySQLReCompilationsPersec(value uint64) (err error) {
	return instance.SetProperty("SQLReCompilationsPersec", (value))
}

// GetSQLReCompilationsPersec gets the value of SQLReCompilationsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLStatistics) GetPropertySQLReCompilationsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("SQLReCompilationsPersec")
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

// SetUnsafeAutoParamsPersec sets the value of UnsafeAutoParamsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLStatistics) SetPropertyUnsafeAutoParamsPersec(value uint64) (err error) {
	return instance.SetProperty("UnsafeAutoParamsPersec", (value))
}

// GetUnsafeAutoParamsPersec gets the value of UnsafeAutoParamsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDSQLStatistics) GetPropertyUnsafeAutoParamsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("UnsafeAutoParamsPersec")
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
