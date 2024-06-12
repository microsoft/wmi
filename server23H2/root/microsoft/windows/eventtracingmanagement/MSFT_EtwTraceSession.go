// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.Microsoft.Windows.EventTracingManagement
//////////////////////////////////////////////
package eventtracingmanagement

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_EtwTraceSession struct
type MSFT_EtwTraceSession struct {
	*CIM_LogicalElement

	//
	BufferSize uint32

	//
	ClockType uint32

	//
	FlushTimer uint32

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
}

func NewMSFT_EtwTraceSessionEx1(instance *cim.WmiInstance) (newInstance *MSFT_EtwTraceSession, err error) {
	tmp, err := NewCIM_LogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_EtwTraceSession{
		CIM_LogicalElement: tmp,
	}
	return
}

func NewMSFT_EtwTraceSessionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_EtwTraceSession, err error) {
	tmp, err := NewCIM_LogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_EtwTraceSession{
		CIM_LogicalElement: tmp,
	}
	return
}

// SetBufferSize sets the value of BufferSize for the instance
func (instance *MSFT_EtwTraceSession) SetPropertyBufferSize(value uint32) (err error) {
	return instance.SetProperty("BufferSize", (value))
}

// GetBufferSize gets the value of BufferSize for the instance
func (instance *MSFT_EtwTraceSession) GetPropertyBufferSize() (value uint32, err error) {
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
func (instance *MSFT_EtwTraceSession) SetPropertyClockType(value uint32) (err error) {
	return instance.SetProperty("ClockType", (value))
}

// GetClockType gets the value of ClockType for the instance
func (instance *MSFT_EtwTraceSession) GetPropertyClockType() (value uint32, err error) {
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

// SetFlushTimer sets the value of FlushTimer for the instance
func (instance *MSFT_EtwTraceSession) SetPropertyFlushTimer(value uint32) (err error) {
	return instance.SetProperty("FlushTimer", (value))
}

// GetFlushTimer gets the value of FlushTimer for the instance
func (instance *MSFT_EtwTraceSession) GetPropertyFlushTimer() (value uint32, err error) {
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

// SetLocalFilePath sets the value of LocalFilePath for the instance
func (instance *MSFT_EtwTraceSession) SetPropertyLocalFilePath(value string) (err error) {
	return instance.SetProperty("LocalFilePath", (value))
}

// GetLocalFilePath gets the value of LocalFilePath for the instance
func (instance *MSFT_EtwTraceSession) GetPropertyLocalFilePath() (value string, err error) {
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
func (instance *MSFT_EtwTraceSession) SetPropertyLogFileMode(value uint32) (err error) {
	return instance.SetProperty("LogFileMode", (value))
}

// GetLogFileMode gets the value of LogFileMode for the instance
func (instance *MSFT_EtwTraceSession) GetPropertyLogFileMode() (value uint32, err error) {
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
func (instance *MSFT_EtwTraceSession) SetPropertyMaximumBuffers(value uint32) (err error) {
	return instance.SetProperty("MaximumBuffers", (value))
}

// GetMaximumBuffers gets the value of MaximumBuffers for the instance
func (instance *MSFT_EtwTraceSession) GetPropertyMaximumBuffers() (value uint32, err error) {
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
func (instance *MSFT_EtwTraceSession) SetPropertyMaximumFileSize(value uint32) (err error) {
	return instance.SetProperty("MaximumFileSize", (value))
}

// GetMaximumFileSize gets the value of MaximumFileSize for the instance
func (instance *MSFT_EtwTraceSession) GetPropertyMaximumFileSize() (value uint32, err error) {
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
func (instance *MSFT_EtwTraceSession) SetPropertyMinimumBuffers(value uint32) (err error) {
	return instance.SetProperty("MinimumBuffers", (value))
}

// GetMinimumBuffers gets the value of MinimumBuffers for the instance
func (instance *MSFT_EtwTraceSession) GetPropertyMinimumBuffers() (value uint32, err error) {
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

//

// <param name="DeleteFile" type="bool "></param>
// <param name="DestinationFolder" type="string "></param>

// <param name="ErrorCode" type="uint32 "></param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="SourceFilePath" type="string "></param>
func (instance *MSFT_EtwTraceSession) Send( /* IN */ DestinationFolder string,
	/* IN */ DeleteFile bool,
	/* OUT */ SourceFilePath string,
	/* OUT */ ErrorCode uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Send", DestinationFolder, DeleteFile)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
