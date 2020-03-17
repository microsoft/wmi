// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.TerminalServices
//////////////////////////////////////////////
package terminalservices

// Win32_TSRemoteControlSetting struct
type Win32_TSRemoteControlSetting struct {
	Win32_TerminalSetting

	//
	LevelOfControl uint32

	//
	PolicySourceLevelOfControl uint32

	//
	RemoteControlPolicy uint32
}

// SetLevelOfControl sets the value of LevelOfControl for the instance
func (instance *Win32_TSRemoteControlSetting) SetPropertyLevelOfControl(value uint32) (err error) {
	return instance.SetProperty("LevelOfControl", value)
}

// GetLevelOfControl gets the value of LevelOfControl for the instance
func (instance *Win32_TSRemoteControlSetting) GetPropertyLevelOfControl() (value uint32, err error) {
	retValue, err := instance.GetProperty("LevelOfControl")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicySourceLevelOfControl sets the value of PolicySourceLevelOfControl for the instance
func (instance *Win32_TSRemoteControlSetting) SetPropertyPolicySourceLevelOfControl(value uint32) (err error) {
	return instance.SetProperty("PolicySourceLevelOfControl", value)
}

// GetPolicySourceLevelOfControl gets the value of PolicySourceLevelOfControl for the instance
func (instance *Win32_TSRemoteControlSetting) GetPropertyPolicySourceLevelOfControl() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceLevelOfControl")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemoteControlPolicy sets the value of RemoteControlPolicy for the instance
func (instance *Win32_TSRemoteControlSetting) SetPropertyRemoteControlPolicy(value uint32) (err error) {
	return instance.SetProperty("RemoteControlPolicy", value)
}

// GetRemoteControlPolicy gets the value of RemoteControlPolicy for the instance
func (instance *Win32_TSRemoteControlSetting) GetPropertyRemoteControlPolicy() (value uint32, err error) {
	retValue, err := instance.GetProperty("RemoteControlPolicy")
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

// <param name="LevelOfControl" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSRemoteControlSetting) RemoteControl( /* IN */ LevelOfControl uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RemoteControl", LevelOfControl)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
