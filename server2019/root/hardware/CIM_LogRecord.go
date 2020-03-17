// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Hardware
//////////////////////////////////////////////
package hardware

// CIM_LogRecord struct
type CIM_LogRecord struct {
	CIM_ManagedElement

	//
	CreationClassName string

	//
	DataFormat string

	//
	LogCreationClassName string

	//
	LogName string

	//
	MessageTimestamp string

	//
	RecordData string

	//
	RecordFormat string

	//
	RecordID string
}

// SetCreationClassName sets the value of CreationClassName for the instance
func (instance *CIM_LogRecord) SetPropertyCreationClassName(value string) (err error) {
	return instance.SetProperty("CreationClassName", value)
}

// GetCreationClassName gets the value of CreationClassName for the instance
func (instance *CIM_LogRecord) GetPropertyCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("CreationClassName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDataFormat sets the value of DataFormat for the instance
func (instance *CIM_LogRecord) SetPropertyDataFormat(value string) (err error) {
	return instance.SetProperty("DataFormat", value)
}

// GetDataFormat gets the value of DataFormat for the instance
func (instance *CIM_LogRecord) GetPropertyDataFormat() (value string, err error) {
	retValue, err := instance.GetProperty("DataFormat")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogCreationClassName sets the value of LogCreationClassName for the instance
func (instance *CIM_LogRecord) SetPropertyLogCreationClassName(value string) (err error) {
	return instance.SetProperty("LogCreationClassName", value)
}

// GetLogCreationClassName gets the value of LogCreationClassName for the instance
func (instance *CIM_LogRecord) GetPropertyLogCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("LogCreationClassName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogName sets the value of LogName for the instance
func (instance *CIM_LogRecord) SetPropertyLogName(value string) (err error) {
	return instance.SetProperty("LogName", value)
}

// GetLogName gets the value of LogName for the instance
func (instance *CIM_LogRecord) GetPropertyLogName() (value string, err error) {
	retValue, err := instance.GetProperty("LogName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMessageTimestamp sets the value of MessageTimestamp for the instance
func (instance *CIM_LogRecord) SetPropertyMessageTimestamp(value string) (err error) {
	return instance.SetProperty("MessageTimestamp", value)
}

// GetMessageTimestamp gets the value of MessageTimestamp for the instance
func (instance *CIM_LogRecord) GetPropertyMessageTimestamp() (value string, err error) {
	retValue, err := instance.GetProperty("MessageTimestamp")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRecordData sets the value of RecordData for the instance
func (instance *CIM_LogRecord) SetPropertyRecordData(value string) (err error) {
	return instance.SetProperty("RecordData", value)
}

// GetRecordData gets the value of RecordData for the instance
func (instance *CIM_LogRecord) GetPropertyRecordData() (value string, err error) {
	retValue, err := instance.GetProperty("RecordData")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRecordFormat sets the value of RecordFormat for the instance
func (instance *CIM_LogRecord) SetPropertyRecordFormat(value string) (err error) {
	return instance.SetProperty("RecordFormat", value)
}

// GetRecordFormat gets the value of RecordFormat for the instance
func (instance *CIM_LogRecord) GetPropertyRecordFormat() (value string, err error) {
	retValue, err := instance.GetProperty("RecordFormat")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRecordID sets the value of RecordID for the instance
func (instance *CIM_LogRecord) SetPropertyRecordID(value string) (err error) {
	return instance.SetProperty("RecordID", value)
}

// GetRecordID gets the value of RecordID for the instance
func (instance *CIM_LogRecord) GetPropertyRecordID() (value string, err error) {
	retValue, err := instance.GetProperty("RecordID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
