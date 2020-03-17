// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_ODBCTranslatorSpecification struct
type Win32_ODBCTranslatorSpecification struct {
	CIM_Check

	//
	File string

	//
	SetupFile string

	//
	Translator string
}

// SetFile sets the value of File for the instance
func (instance *Win32_ODBCTranslatorSpecification) SetPropertyFile(value string) (err error) {
	return instance.SetProperty("File", value)
}

// GetFile gets the value of File for the instance
func (instance *Win32_ODBCTranslatorSpecification) GetPropertyFile() (value string, err error) {
	retValue, err := instance.GetProperty("File")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSetupFile sets the value of SetupFile for the instance
func (instance *Win32_ODBCTranslatorSpecification) SetPropertySetupFile(value string) (err error) {
	return instance.SetProperty("SetupFile", value)
}

// GetSetupFile gets the value of SetupFile for the instance
func (instance *Win32_ODBCTranslatorSpecification) GetPropertySetupFile() (value string, err error) {
	retValue, err := instance.GetProperty("SetupFile")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTranslator sets the value of Translator for the instance
func (instance *Win32_ODBCTranslatorSpecification) SetPropertyTranslator(value string) (err error) {
	return instance.SetProperty("Translator", value)
}

// GetTranslator gets the value of Translator for the instance
func (instance *Win32_ODBCTranslatorSpecification) GetPropertyTranslator() (value string, err error) {
	retValue, err := instance.GetProperty("Translator")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
