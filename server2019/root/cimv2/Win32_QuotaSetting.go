// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_QuotaSetting struct
type Win32_QuotaSetting struct {
	CIM_Setting

	//
	DefaultLimit int64

	//
	DefaultWarningLimit int64

	//
	ExceededNotification bool

	//
	State uint32

	//
	VolumePath string

	//
	WarningExceededNotification bool
}

// SetDefaultLimit sets the value of DefaultLimit for the instance
func (instance *Win32_QuotaSetting) SetPropertyDefaultLimit(value int64) (err error) {
	return instance.SetProperty("DefaultLimit", value)
}

// GetDefaultLimit gets the value of DefaultLimit for the instance
func (instance *Win32_QuotaSetting) GetPropertyDefaultLimit() (value int64, err error) {
	retValue, err := instance.GetProperty("DefaultLimit")
	if err != nil {
		return
	}
	value, ok := retValue.(int64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDefaultWarningLimit sets the value of DefaultWarningLimit for the instance
func (instance *Win32_QuotaSetting) SetPropertyDefaultWarningLimit(value int64) (err error) {
	return instance.SetProperty("DefaultWarningLimit", value)
}

// GetDefaultWarningLimit gets the value of DefaultWarningLimit for the instance
func (instance *Win32_QuotaSetting) GetPropertyDefaultWarningLimit() (value int64, err error) {
	retValue, err := instance.GetProperty("DefaultWarningLimit")
	if err != nil {
		return
	}
	value, ok := retValue.(int64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExceededNotification sets the value of ExceededNotification for the instance
func (instance *Win32_QuotaSetting) SetPropertyExceededNotification(value bool) (err error) {
	return instance.SetProperty("ExceededNotification", value)
}

// GetExceededNotification gets the value of ExceededNotification for the instance
func (instance *Win32_QuotaSetting) GetPropertyExceededNotification() (value bool, err error) {
	retValue, err := instance.GetProperty("ExceededNotification")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetState sets the value of State for the instance
func (instance *Win32_QuotaSetting) SetPropertyState(value uint32) (err error) {
	return instance.SetProperty("State", value)
}

// GetState gets the value of State for the instance
func (instance *Win32_QuotaSetting) GetPropertyState() (value uint32, err error) {
	retValue, err := instance.GetProperty("State")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVolumePath sets the value of VolumePath for the instance
func (instance *Win32_QuotaSetting) SetPropertyVolumePath(value string) (err error) {
	return instance.SetProperty("VolumePath", value)
}

// GetVolumePath gets the value of VolumePath for the instance
func (instance *Win32_QuotaSetting) GetPropertyVolumePath() (value string, err error) {
	retValue, err := instance.GetProperty("VolumePath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWarningExceededNotification sets the value of WarningExceededNotification for the instance
func (instance *Win32_QuotaSetting) SetPropertyWarningExceededNotification(value bool) (err error) {
	return instance.SetProperty("WarningExceededNotification", value)
}

// GetWarningExceededNotification gets the value of WarningExceededNotification for the instance
func (instance *Win32_QuotaSetting) GetPropertyWarningExceededNotification() (value bool, err error) {
	retValue, err := instance.GetProperty("WarningExceededNotification")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
