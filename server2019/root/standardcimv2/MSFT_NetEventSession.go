// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetEventSession struct
type MSFT_NetEventSession struct {
	CIM_LogicalElement

	//
	CaptureMode uint8

	//
	Guid string

	//
	LocalFilePath string

	//
	MaxFileSize uint32

	//
	MaxNumberOfBuffers uint8

	//
	SessionStatus uint8

	//
	TraceBufferSize uint32
}

// SetCaptureMode sets the value of CaptureMode for the instance
func (instance *MSFT_NetEventSession) SetPropertyCaptureMode(value uint8) (err error) {
	return instance.SetProperty("CaptureMode", value)
}

// GetCaptureMode gets the value of CaptureMode for the instance
func (instance *MSFT_NetEventSession) GetPropertyCaptureMode() (value uint8, err error) {
	retValue, err := instance.GetProperty("CaptureMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGuid sets the value of Guid for the instance
func (instance *MSFT_NetEventSession) SetPropertyGuid(value string) (err error) {
	return instance.SetProperty("Guid", value)
}

// GetGuid gets the value of Guid for the instance
func (instance *MSFT_NetEventSession) GetPropertyGuid() (value string, err error) {
	retValue, err := instance.GetProperty("Guid")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalFilePath sets the value of LocalFilePath for the instance
func (instance *MSFT_NetEventSession) SetPropertyLocalFilePath(value string) (err error) {
	return instance.SetProperty("LocalFilePath", value)
}

// GetLocalFilePath gets the value of LocalFilePath for the instance
func (instance *MSFT_NetEventSession) GetPropertyLocalFilePath() (value string, err error) {
	retValue, err := instance.GetProperty("LocalFilePath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxFileSize sets the value of MaxFileSize for the instance
func (instance *MSFT_NetEventSession) SetPropertyMaxFileSize(value uint32) (err error) {
	return instance.SetProperty("MaxFileSize", value)
}

// GetMaxFileSize gets the value of MaxFileSize for the instance
func (instance *MSFT_NetEventSession) GetPropertyMaxFileSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxFileSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxNumberOfBuffers sets the value of MaxNumberOfBuffers for the instance
func (instance *MSFT_NetEventSession) SetPropertyMaxNumberOfBuffers(value uint8) (err error) {
	return instance.SetProperty("MaxNumberOfBuffers", value)
}

// GetMaxNumberOfBuffers gets the value of MaxNumberOfBuffers for the instance
func (instance *MSFT_NetEventSession) GetPropertyMaxNumberOfBuffers() (value uint8, err error) {
	retValue, err := instance.GetProperty("MaxNumberOfBuffers")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSessionStatus sets the value of SessionStatus for the instance
func (instance *MSFT_NetEventSession) SetPropertySessionStatus(value uint8) (err error) {
	return instance.SetProperty("SessionStatus", value)
}

// GetSessionStatus gets the value of SessionStatus for the instance
func (instance *MSFT_NetEventSession) GetPropertySessionStatus() (value uint8, err error) {
	retValue, err := instance.GetProperty("SessionStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTraceBufferSize sets the value of TraceBufferSize for the instance
func (instance *MSFT_NetEventSession) SetPropertyTraceBufferSize(value uint32) (err error) {
	return instance.SetProperty("TraceBufferSize", value)
}

// GetTraceBufferSize gets the value of TraceBufferSize for the instance
func (instance *MSFT_NetEventSession) GetPropertyTraceBufferSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("TraceBufferSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetEventSession) Start() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Start")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetEventSession) Stop() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Stop")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
