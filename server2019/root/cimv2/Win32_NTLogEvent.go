// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_NTLogEvent struct
type Win32_NTLogEvent struct {
	cim.WmiInstance

	//
	Category uint16

	//
	CategoryString string

	//
	ComputerName string

	//
	Data []uint8

	//
	EventCode uint16

	//
	EventIdentifier uint32

	//
	EventType uint8

	//
	InsertionStrings []string

	//
	Logfile string

	//
	Message string

	//
	RecordNumber uint32

	//
	SourceName string

	//
	TimeGenerated string

	//
	TimeWritten string

	//
	Type string

	//
	User string
}

// SetCategory sets the value of Category for the instance
func (instance *Win32_NTLogEvent) SetPropertyCategory(value uint16) (err error) {
	return instance.SetProperty("Category", value)
}

// GetCategory gets the value of Category for the instance
func (instance *Win32_NTLogEvent) GetPropertyCategory() (value uint16, err error) {
	retValue, err := instance.GetProperty("Category")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCategoryString sets the value of CategoryString for the instance
func (instance *Win32_NTLogEvent) SetPropertyCategoryString(value string) (err error) {
	return instance.SetProperty("CategoryString", value)
}

// GetCategoryString gets the value of CategoryString for the instance
func (instance *Win32_NTLogEvent) GetPropertyCategoryString() (value string, err error) {
	retValue, err := instance.GetProperty("CategoryString")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetComputerName sets the value of ComputerName for the instance
func (instance *Win32_NTLogEvent) SetPropertyComputerName(value string) (err error) {
	return instance.SetProperty("ComputerName", value)
}

// GetComputerName gets the value of ComputerName for the instance
func (instance *Win32_NTLogEvent) GetPropertyComputerName() (value string, err error) {
	retValue, err := instance.GetProperty("ComputerName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetData sets the value of Data for the instance
func (instance *Win32_NTLogEvent) SetPropertyData(value []uint8) (err error) {
	return instance.SetProperty("Data", value)
}

// GetData gets the value of Data for the instance
func (instance *Win32_NTLogEvent) GetPropertyData() (value []uint8, err error) {
	retValue, err := instance.GetProperty("Data")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEventCode sets the value of EventCode for the instance
func (instance *Win32_NTLogEvent) SetPropertyEventCode(value uint16) (err error) {
	return instance.SetProperty("EventCode", value)
}

// GetEventCode gets the value of EventCode for the instance
func (instance *Win32_NTLogEvent) GetPropertyEventCode() (value uint16, err error) {
	retValue, err := instance.GetProperty("EventCode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEventIdentifier sets the value of EventIdentifier for the instance
func (instance *Win32_NTLogEvent) SetPropertyEventIdentifier(value uint32) (err error) {
	return instance.SetProperty("EventIdentifier", value)
}

// GetEventIdentifier gets the value of EventIdentifier for the instance
func (instance *Win32_NTLogEvent) GetPropertyEventIdentifier() (value uint32, err error) {
	retValue, err := instance.GetProperty("EventIdentifier")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEventType sets the value of EventType for the instance
func (instance *Win32_NTLogEvent) SetPropertyEventType(value uint8) (err error) {
	return instance.SetProperty("EventType", value)
}

// GetEventType gets the value of EventType for the instance
func (instance *Win32_NTLogEvent) GetPropertyEventType() (value uint8, err error) {
	retValue, err := instance.GetProperty("EventType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInsertionStrings sets the value of InsertionStrings for the instance
func (instance *Win32_NTLogEvent) SetPropertyInsertionStrings(value []string) (err error) {
	return instance.SetProperty("InsertionStrings", value)
}

// GetInsertionStrings gets the value of InsertionStrings for the instance
func (instance *Win32_NTLogEvent) GetPropertyInsertionStrings() (value []string, err error) {
	retValue, err := instance.GetProperty("InsertionStrings")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogfile sets the value of Logfile for the instance
func (instance *Win32_NTLogEvent) SetPropertyLogfile(value string) (err error) {
	return instance.SetProperty("Logfile", value)
}

// GetLogfile gets the value of Logfile for the instance
func (instance *Win32_NTLogEvent) GetPropertyLogfile() (value string, err error) {
	retValue, err := instance.GetProperty("Logfile")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMessage sets the value of Message for the instance
func (instance *Win32_NTLogEvent) SetPropertyMessage(value string) (err error) {
	return instance.SetProperty("Message", value)
}

// GetMessage gets the value of Message for the instance
func (instance *Win32_NTLogEvent) GetPropertyMessage() (value string, err error) {
	retValue, err := instance.GetProperty("Message")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRecordNumber sets the value of RecordNumber for the instance
func (instance *Win32_NTLogEvent) SetPropertyRecordNumber(value uint32) (err error) {
	return instance.SetProperty("RecordNumber", value)
}

// GetRecordNumber gets the value of RecordNumber for the instance
func (instance *Win32_NTLogEvent) GetPropertyRecordNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("RecordNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSourceName sets the value of SourceName for the instance
func (instance *Win32_NTLogEvent) SetPropertySourceName(value string) (err error) {
	return instance.SetProperty("SourceName", value)
}

// GetSourceName gets the value of SourceName for the instance
func (instance *Win32_NTLogEvent) GetPropertySourceName() (value string, err error) {
	retValue, err := instance.GetProperty("SourceName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimeGenerated sets the value of TimeGenerated for the instance
func (instance *Win32_NTLogEvent) SetPropertyTimeGenerated(value string) (err error) {
	return instance.SetProperty("TimeGenerated", value)
}

// GetTimeGenerated gets the value of TimeGenerated for the instance
func (instance *Win32_NTLogEvent) GetPropertyTimeGenerated() (value string, err error) {
	retValue, err := instance.GetProperty("TimeGenerated")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimeWritten sets the value of TimeWritten for the instance
func (instance *Win32_NTLogEvent) SetPropertyTimeWritten(value string) (err error) {
	return instance.SetProperty("TimeWritten", value)
}

// GetTimeWritten gets the value of TimeWritten for the instance
func (instance *Win32_NTLogEvent) GetPropertyTimeWritten() (value string, err error) {
	retValue, err := instance.GetProperty("TimeWritten")
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
func (instance *Win32_NTLogEvent) SetPropertyType(value string) (err error) {
	return instance.SetProperty("Type", value)
}

// GetType gets the value of Type for the instance
func (instance *Win32_NTLogEvent) GetPropertyType() (value string, err error) {
	retValue, err := instance.GetProperty("Type")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUser sets the value of User for the instance
func (instance *Win32_NTLogEvent) SetPropertyUser(value string) (err error) {
	return instance.SetProperty("User", value)
}

// GetUser gets the value of User for the instance
func (instance *Win32_NTLogEvent) GetPropertyUser() (value string, err error) {
	retValue, err := instance.GetProperty("User")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
