// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.ManagementTools
//////////////////////////////////////////////
package managementtools

// MSFT_MTEventChannel struct
type MSFT_MTEventChannel struct {
	CIM_ManagedElement

	//
	ClassicLog bool

	//
	DisplayName string

	//
	DisplayPath string

	//
	Enabled bool

	//
	EventsCount string

	//
	LogFilePath string

	//
	LogFileSize uint64

	//
	Name string

	//
	Type uint32
}

// SetClassicLog sets the value of ClassicLog for the instance
func (instance *MSFT_MTEventChannel) SetPropertyClassicLog(value bool) (err error) {
	return instance.SetProperty("ClassicLog", value)
}

// GetClassicLog gets the value of ClassicLog for the instance
func (instance *MSFT_MTEventChannel) GetPropertyClassicLog() (value bool, err error) {
	retValue, err := instance.GetProperty("ClassicLog")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisplayName sets the value of DisplayName for the instance
func (instance *MSFT_MTEventChannel) SetPropertyDisplayName(value string) (err error) {
	return instance.SetProperty("DisplayName", value)
}

// GetDisplayName gets the value of DisplayName for the instance
func (instance *MSFT_MTEventChannel) GetPropertyDisplayName() (value string, err error) {
	retValue, err := instance.GetProperty("DisplayName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisplayPath sets the value of DisplayPath for the instance
func (instance *MSFT_MTEventChannel) SetPropertyDisplayPath(value string) (err error) {
	return instance.SetProperty("DisplayPath", value)
}

// GetDisplayPath gets the value of DisplayPath for the instance
func (instance *MSFT_MTEventChannel) GetPropertyDisplayPath() (value string, err error) {
	retValue, err := instance.GetProperty("DisplayPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *MSFT_MTEventChannel) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", value)
}

// GetEnabled gets the value of Enabled for the instance
func (instance *MSFT_MTEventChannel) GetPropertyEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("Enabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEventsCount sets the value of EventsCount for the instance
func (instance *MSFT_MTEventChannel) SetPropertyEventsCount(value string) (err error) {
	return instance.SetProperty("EventsCount", value)
}

// GetEventsCount gets the value of EventsCount for the instance
func (instance *MSFT_MTEventChannel) GetPropertyEventsCount() (value string, err error) {
	retValue, err := instance.GetProperty("EventsCount")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogFilePath sets the value of LogFilePath for the instance
func (instance *MSFT_MTEventChannel) SetPropertyLogFilePath(value string) (err error) {
	return instance.SetProperty("LogFilePath", value)
}

// GetLogFilePath gets the value of LogFilePath for the instance
func (instance *MSFT_MTEventChannel) GetPropertyLogFilePath() (value string, err error) {
	retValue, err := instance.GetProperty("LogFilePath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogFileSize sets the value of LogFileSize for the instance
func (instance *MSFT_MTEventChannel) SetPropertyLogFileSize(value uint64) (err error) {
	return instance.SetProperty("LogFileSize", value)
}

// GetLogFileSize gets the value of LogFileSize for the instance
func (instance *MSFT_MTEventChannel) GetPropertyLogFileSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogFileSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_MTEventChannel) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_MTEventChannel) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetType sets the value of Type for the instance
func (instance *MSFT_MTEventChannel) SetPropertyType(value uint32) (err error) {
	return instance.SetProperty("Type", value)
}

// GetType gets the value of Type for the instance
func (instance *MSFT_MTEventChannel) GetPropertyType() (value uint32, err error) {
	retValue, err := instance.GetProperty("Type")
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
func (instance *MSFT_MTEventChannel) Enable() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Enable")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MTEventChannel) Disable() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Disable")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MTEventChannel) ClearLogFile() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ClearLogFile")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="BatchSize" type="uint32 "></param>
// <param name="FilterXml" type="string "></param>
// <param name="ReverseDirection" type="bool "></param>
// <param name="Skip" type="uint64 "></param>
// <param name="Top" type="uint64 "></param>

// <param name="Result" type="MSFT_MTEventRecord []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MTEventChannel) GetEventRecords( /* IN */ FilterXml string,
	/* IN */ Skip uint64,
	/* IN */ Top uint64,
	/* IN */ ReverseDirection bool,
	/* IN */ BatchSize uint32,
	/* OUT */ Result []MSFT_MTEventRecord) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetEventRecords", FilterXml, Skip, Top, ReverseDirection, BatchSize)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Result" type="MSFT_MTEventChannel []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MTEventChannel) GetWindowsEventChannels( /* OUT */ Result []MSFT_MTEventChannel) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetWindowsEventChannels")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
