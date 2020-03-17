// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.subscription
//////////////////////////////////////////////
package subscription

// LogFileEventConsumer struct
type LogFileEventConsumer struct {
	__EventConsumer

	//
	Filename string

	//
	IsUnicode bool

	//
	MaximumFileSize uint64

	//
	Name string

	//
	Text string
}

// SetFilename sets the value of Filename for the instance
func (instance *LogFileEventConsumer) SetPropertyFilename(value string) (err error) {
	return instance.SetProperty("Filename", value)
}

// GetFilename gets the value of Filename for the instance
func (instance *LogFileEventConsumer) GetPropertyFilename() (value string, err error) {
	retValue, err := instance.GetProperty("Filename")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsUnicode sets the value of IsUnicode for the instance
func (instance *LogFileEventConsumer) SetPropertyIsUnicode(value bool) (err error) {
	return instance.SetProperty("IsUnicode", value)
}

// GetIsUnicode gets the value of IsUnicode for the instance
func (instance *LogFileEventConsumer) GetPropertyIsUnicode() (value bool, err error) {
	retValue, err := instance.GetProperty("IsUnicode")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaximumFileSize sets the value of MaximumFileSize for the instance
func (instance *LogFileEventConsumer) SetPropertyMaximumFileSize(value uint64) (err error) {
	return instance.SetProperty("MaximumFileSize", value)
}

// GetMaximumFileSize gets the value of MaximumFileSize for the instance
func (instance *LogFileEventConsumer) GetPropertyMaximumFileSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaximumFileSize")
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
func (instance *LogFileEventConsumer) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *LogFileEventConsumer) GetPropertyName() (value string, err error) {
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

// SetText sets the value of Text for the instance
func (instance *LogFileEventConsumer) SetPropertyText(value string) (err error) {
	return instance.SetProperty("Text", value)
}

// GetText gets the value of Text for the instance
func (instance *LogFileEventConsumer) GetPropertyText() (value string, err error) {
	retValue, err := instance.GetProperty("Text")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
