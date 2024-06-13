// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.EventTracingManagement
//////////////////////////////////////////////
package eventtracingmanagement

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_AutologgerConfig struct
type MSFT_AutologgerConfig struct {
	*CIM_LogicalElement

	//
	BufferSize uint32

	//
	ClockType uint32

	//
	DisableRealtimePersistence uint32

	//
	FileCount uint32

	//
	FileMax uint32

	//
	FlushTimer uint32

	//
	Guid string

	//
	InitStatus uint32

	//
	LocalFilePath string

	//
	LogFileMode uint32

	//
	MaximumBuffers uint32

	//
	MaximumFileSize uint32

	//
	MinimumBuffers uint32

	//
	Start uint32
}

func NewMSFT_AutologgerConfigEx1(instance *cim.WmiInstance) (newInstance *MSFT_AutologgerConfig, err error) {
	tmp, err := NewCIM_LogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_AutologgerConfig{
		CIM_LogicalElement: tmp,
	}
	return
}

func NewMSFT_AutologgerConfigEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_AutologgerConfig, err error) {
	tmp, err := NewCIM_LogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_AutologgerConfig{
		CIM_LogicalElement: tmp,
	}
	return
}

// SetBufferSize sets the value of BufferSize for the instance
func (instance *MSFT_AutologgerConfig) SetPropertyBufferSize(value uint32) (err error) {
	return instance.SetProperty("BufferSize", (value))
}

// GetBufferSize gets the value of BufferSize for the instance
func (instance *MSFT_AutologgerConfig) GetPropertyBufferSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("BufferSize")
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

// SetClockType sets the value of ClockType for the instance
func (instance *MSFT_AutologgerConfig) SetPropertyClockType(value uint32) (err error) {
	return instance.SetProperty("ClockType", (value))
}

// GetClockType gets the value of ClockType for the instance
func (instance *MSFT_AutologgerConfig) GetPropertyClockType() (value uint32, err error) {
	retValue, err := instance.GetProperty("ClockType")
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

// SetDisableRealtimePersistence sets the value of DisableRealtimePersistence for the instance
func (instance *MSFT_AutologgerConfig) SetPropertyDisableRealtimePersistence(value uint32) (err error) {
	return instance.SetProperty("DisableRealtimePersistence", (value))
}

// GetDisableRealtimePersistence gets the value of DisableRealtimePersistence for the instance
func (instance *MSFT_AutologgerConfig) GetPropertyDisableRealtimePersistence() (value uint32, err error) {
	retValue, err := instance.GetProperty("DisableRealtimePersistence")
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

// SetFileCount sets the value of FileCount for the instance
func (instance *MSFT_AutologgerConfig) SetPropertyFileCount(value uint32) (err error) {
	return instance.SetProperty("FileCount", (value))
}

// GetFileCount gets the value of FileCount for the instance
func (instance *MSFT_AutologgerConfig) GetPropertyFileCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("FileCount")
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

// SetFileMax sets the value of FileMax for the instance
func (instance *MSFT_AutologgerConfig) SetPropertyFileMax(value uint32) (err error) {
	return instance.SetProperty("FileMax", (value))
}

// GetFileMax gets the value of FileMax for the instance
func (instance *MSFT_AutologgerConfig) GetPropertyFileMax() (value uint32, err error) {
	retValue, err := instance.GetProperty("FileMax")
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

// SetFlushTimer sets the value of FlushTimer for the instance
func (instance *MSFT_AutologgerConfig) SetPropertyFlushTimer(value uint32) (err error) {
	return instance.SetProperty("FlushTimer", (value))
}

// GetFlushTimer gets the value of FlushTimer for the instance
func (instance *MSFT_AutologgerConfig) GetPropertyFlushTimer() (value uint32, err error) {
	retValue, err := instance.GetProperty("FlushTimer")
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

// SetGuid sets the value of Guid for the instance
func (instance *MSFT_AutologgerConfig) SetPropertyGuid(value string) (err error) {
	return instance.SetProperty("Guid", (value))
}

// GetGuid gets the value of Guid for the instance
func (instance *MSFT_AutologgerConfig) GetPropertyGuid() (value string, err error) {
	retValue, err := instance.GetProperty("Guid")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetInitStatus sets the value of InitStatus for the instance
func (instance *MSFT_AutologgerConfig) SetPropertyInitStatus(value uint32) (err error) {
	return instance.SetProperty("InitStatus", (value))
}

// GetInitStatus gets the value of InitStatus for the instance
func (instance *MSFT_AutologgerConfig) GetPropertyInitStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("InitStatus")
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

// SetLocalFilePath sets the value of LocalFilePath for the instance
func (instance *MSFT_AutologgerConfig) SetPropertyLocalFilePath(value string) (err error) {
	return instance.SetProperty("LocalFilePath", (value))
}

// GetLocalFilePath gets the value of LocalFilePath for the instance
func (instance *MSFT_AutologgerConfig) GetPropertyLocalFilePath() (value string, err error) {
	retValue, err := instance.GetProperty("LocalFilePath")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetLogFileMode sets the value of LogFileMode for the instance
func (instance *MSFT_AutologgerConfig) SetPropertyLogFileMode(value uint32) (err error) {
	return instance.SetProperty("LogFileMode", (value))
}

// GetLogFileMode gets the value of LogFileMode for the instance
func (instance *MSFT_AutologgerConfig) GetPropertyLogFileMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogFileMode")
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

// SetMaximumBuffers sets the value of MaximumBuffers for the instance
func (instance *MSFT_AutologgerConfig) SetPropertyMaximumBuffers(value uint32) (err error) {
	return instance.SetProperty("MaximumBuffers", (value))
}

// GetMaximumBuffers gets the value of MaximumBuffers for the instance
func (instance *MSFT_AutologgerConfig) GetPropertyMaximumBuffers() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumBuffers")
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

// SetMaximumFileSize sets the value of MaximumFileSize for the instance
func (instance *MSFT_AutologgerConfig) SetPropertyMaximumFileSize(value uint32) (err error) {
	return instance.SetProperty("MaximumFileSize", (value))
}

// GetMaximumFileSize gets the value of MaximumFileSize for the instance
func (instance *MSFT_AutologgerConfig) GetPropertyMaximumFileSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumFileSize")
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

// SetMinimumBuffers sets the value of MinimumBuffers for the instance
func (instance *MSFT_AutologgerConfig) SetPropertyMinimumBuffers(value uint32) (err error) {
	return instance.SetProperty("MinimumBuffers", (value))
}

// GetMinimumBuffers gets the value of MinimumBuffers for the instance
func (instance *MSFT_AutologgerConfig) GetPropertyMinimumBuffers() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinimumBuffers")
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

// SetStart sets the value of Start for the instance
func (instance *MSFT_AutologgerConfig) SetPropertyStart(value uint32) (err error) {
	return instance.SetProperty("Start", (value))
}

// GetStart gets the value of Start for the instance
func (instance *MSFT_AutologgerConfig) GetPropertyStart() (value uint32, err error) {
	retValue, err := instance.GetProperty("Start")
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
