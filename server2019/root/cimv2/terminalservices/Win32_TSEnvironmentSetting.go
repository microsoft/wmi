// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.TerminalServices
//////////////////////////////////////////////
package terminalservices

// Win32_TSEnvironmentSetting struct
type Win32_TSEnvironmentSetting struct {
	Win32_TerminalSetting

	//
	ClientWallPaper uint32

	//
	InitialProgramPath string

	//
	InitialProgramPolicy uint32

	//
	PolicySourceClientWallPaper uint32

	//
	PolicySourceInitialProgramPath uint32

	//
	PolicySourceStartIn uint32

	//
	Startin string
}

// SetClientWallPaper sets the value of ClientWallPaper for the instance
func (instance *Win32_TSEnvironmentSetting) SetPropertyClientWallPaper(value uint32) (err error) {
	return instance.SetProperty("ClientWallPaper", value)
}

// GetClientWallPaper gets the value of ClientWallPaper for the instance
func (instance *Win32_TSEnvironmentSetting) GetPropertyClientWallPaper() (value uint32, err error) {
	retValue, err := instance.GetProperty("ClientWallPaper")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInitialProgramPath sets the value of InitialProgramPath for the instance
func (instance *Win32_TSEnvironmentSetting) SetPropertyInitialProgramPath(value string) (err error) {
	return instance.SetProperty("InitialProgramPath", value)
}

// GetInitialProgramPath gets the value of InitialProgramPath for the instance
func (instance *Win32_TSEnvironmentSetting) GetPropertyInitialProgramPath() (value string, err error) {
	retValue, err := instance.GetProperty("InitialProgramPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInitialProgramPolicy sets the value of InitialProgramPolicy for the instance
func (instance *Win32_TSEnvironmentSetting) SetPropertyInitialProgramPolicy(value uint32) (err error) {
	return instance.SetProperty("InitialProgramPolicy", value)
}

// GetInitialProgramPolicy gets the value of InitialProgramPolicy for the instance
func (instance *Win32_TSEnvironmentSetting) GetPropertyInitialProgramPolicy() (value uint32, err error) {
	retValue, err := instance.GetProperty("InitialProgramPolicy")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicySourceClientWallPaper sets the value of PolicySourceClientWallPaper for the instance
func (instance *Win32_TSEnvironmentSetting) SetPropertyPolicySourceClientWallPaper(value uint32) (err error) {
	return instance.SetProperty("PolicySourceClientWallPaper", value)
}

// GetPolicySourceClientWallPaper gets the value of PolicySourceClientWallPaper for the instance
func (instance *Win32_TSEnvironmentSetting) GetPropertyPolicySourceClientWallPaper() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceClientWallPaper")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicySourceInitialProgramPath sets the value of PolicySourceInitialProgramPath for the instance
func (instance *Win32_TSEnvironmentSetting) SetPropertyPolicySourceInitialProgramPath(value uint32) (err error) {
	return instance.SetProperty("PolicySourceInitialProgramPath", value)
}

// GetPolicySourceInitialProgramPath gets the value of PolicySourceInitialProgramPath for the instance
func (instance *Win32_TSEnvironmentSetting) GetPropertyPolicySourceInitialProgramPath() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceInitialProgramPath")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicySourceStartIn sets the value of PolicySourceStartIn for the instance
func (instance *Win32_TSEnvironmentSetting) SetPropertyPolicySourceStartIn(value uint32) (err error) {
	return instance.SetProperty("PolicySourceStartIn", value)
}

// GetPolicySourceStartIn gets the value of PolicySourceStartIn for the instance
func (instance *Win32_TSEnvironmentSetting) GetPropertyPolicySourceStartIn() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceStartIn")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStartin sets the value of Startin for the instance
func (instance *Win32_TSEnvironmentSetting) SetPropertyStartin(value string) (err error) {
	return instance.SetProperty("Startin", value)
}

// GetStartin gets the value of Startin for the instance
func (instance *Win32_TSEnvironmentSetting) GetPropertyStartin() (value string, err error) {
	retValue, err := instance.GetProperty("Startin")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="InitialProgramPath" type="string "></param>
// <param name="Startin" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSEnvironmentSetting) InitialProgram( /* IN */ InitialProgramPath string,
	/* IN */ Startin string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("InitialProgram", InitialProgramPath, Startin)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ClientWallPaper" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSEnvironmentSetting) SetClientWallPaper( /* IN */ ClientWallPaper uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetClientWallPaper", ClientWallPaper)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
