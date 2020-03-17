// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_MIMEInfoAction struct
type Win32_MIMEInfoAction struct {
	CIM_Action

	//
	CLSID string

	//
	ContentType string

	//
	Extension string
}

// SetCLSID sets the value of CLSID for the instance
func (instance *Win32_MIMEInfoAction) SetPropertyCLSID(value string) (err error) {
	return instance.SetProperty("CLSID", value)
}

// GetCLSID gets the value of CLSID for the instance
func (instance *Win32_MIMEInfoAction) GetPropertyCLSID() (value string, err error) {
	retValue, err := instance.GetProperty("CLSID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetContentType sets the value of ContentType for the instance
func (instance *Win32_MIMEInfoAction) SetPropertyContentType(value string) (err error) {
	return instance.SetProperty("ContentType", value)
}

// GetContentType gets the value of ContentType for the instance
func (instance *Win32_MIMEInfoAction) GetPropertyContentType() (value string, err error) {
	retValue, err := instance.GetProperty("ContentType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExtension sets the value of Extension for the instance
func (instance *Win32_MIMEInfoAction) SetPropertyExtension(value string) (err error) {
	return instance.SetProperty("Extension", value)
}

// GetExtension gets the value of Extension for the instance
func (instance *Win32_MIMEInfoAction) GetPropertyExtension() (value string, err error) {
	retValue, err := instance.GetProperty("Extension")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
