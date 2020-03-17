// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_FileSpecification struct
type Win32_FileSpecification struct {
	CIM_FileSpecification

	//
	Attributes uint16

	//
	FileID string

	//
	Language string

	//
	Sequence uint16
}

// SetAttributes sets the value of Attributes for the instance
func (instance *Win32_FileSpecification) SetPropertyAttributes(value uint16) (err error) {
	return instance.SetProperty("Attributes", value)
}

// GetAttributes gets the value of Attributes for the instance
func (instance *Win32_FileSpecification) GetPropertyAttributes() (value uint16, err error) {
	retValue, err := instance.GetProperty("Attributes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFileID sets the value of FileID for the instance
func (instance *Win32_FileSpecification) SetPropertyFileID(value string) (err error) {
	return instance.SetProperty("FileID", value)
}

// GetFileID gets the value of FileID for the instance
func (instance *Win32_FileSpecification) GetPropertyFileID() (value string, err error) {
	retValue, err := instance.GetProperty("FileID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLanguage sets the value of Language for the instance
func (instance *Win32_FileSpecification) SetPropertyLanguage(value string) (err error) {
	return instance.SetProperty("Language", value)
}

// GetLanguage gets the value of Language for the instance
func (instance *Win32_FileSpecification) GetPropertyLanguage() (value string, err error) {
	retValue, err := instance.GetProperty("Language")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSequence sets the value of Sequence for the instance
func (instance *Win32_FileSpecification) SetPropertySequence(value uint16) (err error) {
	return instance.SetProperty("Sequence", value)
}

// GetSequence gets the value of Sequence for the instance
func (instance *Win32_FileSpecification) GetPropertySequence() (value uint16, err error) {
	retValue, err := instance.GetProperty("Sequence")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
