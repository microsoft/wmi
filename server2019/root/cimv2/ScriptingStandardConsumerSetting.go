// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// ScriptingStandardConsumerSetting struct
type ScriptingStandardConsumerSetting struct {
	CIM_Setting

	//
	MaximumScripts uint32

	//
	Timeout uint32
}

// SetMaximumScripts sets the value of MaximumScripts for the instance
func (instance *ScriptingStandardConsumerSetting) SetPropertyMaximumScripts(value uint32) (err error) {
	return instance.SetProperty("MaximumScripts", value)
}

// GetMaximumScripts gets the value of MaximumScripts for the instance
func (instance *ScriptingStandardConsumerSetting) GetPropertyMaximumScripts() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumScripts")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimeout sets the value of Timeout for the instance
func (instance *ScriptingStandardConsumerSetting) SetPropertyTimeout(value uint32) (err error) {
	return instance.SetProperty("Timeout", value)
}

// GetTimeout gets the value of Timeout for the instance
func (instance *ScriptingStandardConsumerSetting) GetPropertyTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("Timeout")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
