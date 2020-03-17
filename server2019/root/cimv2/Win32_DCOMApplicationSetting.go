// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_DCOMApplicationSetting struct
type Win32_DCOMApplicationSetting struct {
	Win32_COMSetting

	//
	AppID string

	//
	AuthenticationLevel uint32

	//
	CustomSurrogate string

	//
	EnableAtStorageActivation bool

	//
	LocalService string

	//
	RemoteServerName string

	//
	RunAsUser string

	//
	ServiceParameters string

	//
	UseSurrogate bool
}

// SetAppID sets the value of AppID for the instance
func (instance *Win32_DCOMApplicationSetting) SetPropertyAppID(value string) (err error) {
	return instance.SetProperty("AppID", value)
}

// GetAppID gets the value of AppID for the instance
func (instance *Win32_DCOMApplicationSetting) GetPropertyAppID() (value string, err error) {
	retValue, err := instance.GetProperty("AppID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAuthenticationLevel sets the value of AuthenticationLevel for the instance
func (instance *Win32_DCOMApplicationSetting) SetPropertyAuthenticationLevel(value uint32) (err error) {
	return instance.SetProperty("AuthenticationLevel", value)
}

// GetAuthenticationLevel gets the value of AuthenticationLevel for the instance
func (instance *Win32_DCOMApplicationSetting) GetPropertyAuthenticationLevel() (value uint32, err error) {
	retValue, err := instance.GetProperty("AuthenticationLevel")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCustomSurrogate sets the value of CustomSurrogate for the instance
func (instance *Win32_DCOMApplicationSetting) SetPropertyCustomSurrogate(value string) (err error) {
	return instance.SetProperty("CustomSurrogate", value)
}

// GetCustomSurrogate gets the value of CustomSurrogate for the instance
func (instance *Win32_DCOMApplicationSetting) GetPropertyCustomSurrogate() (value string, err error) {
	retValue, err := instance.GetProperty("CustomSurrogate")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnableAtStorageActivation sets the value of EnableAtStorageActivation for the instance
func (instance *Win32_DCOMApplicationSetting) SetPropertyEnableAtStorageActivation(value bool) (err error) {
	return instance.SetProperty("EnableAtStorageActivation", value)
}

// GetEnableAtStorageActivation gets the value of EnableAtStorageActivation for the instance
func (instance *Win32_DCOMApplicationSetting) GetPropertyEnableAtStorageActivation() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableAtStorageActivation")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalService sets the value of LocalService for the instance
func (instance *Win32_DCOMApplicationSetting) SetPropertyLocalService(value string) (err error) {
	return instance.SetProperty("LocalService", value)
}

// GetLocalService gets the value of LocalService for the instance
func (instance *Win32_DCOMApplicationSetting) GetPropertyLocalService() (value string, err error) {
	retValue, err := instance.GetProperty("LocalService")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemoteServerName sets the value of RemoteServerName for the instance
func (instance *Win32_DCOMApplicationSetting) SetPropertyRemoteServerName(value string) (err error) {
	return instance.SetProperty("RemoteServerName", value)
}

// GetRemoteServerName gets the value of RemoteServerName for the instance
func (instance *Win32_DCOMApplicationSetting) GetPropertyRemoteServerName() (value string, err error) {
	retValue, err := instance.GetProperty("RemoteServerName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRunAsUser sets the value of RunAsUser for the instance
func (instance *Win32_DCOMApplicationSetting) SetPropertyRunAsUser(value string) (err error) {
	return instance.SetProperty("RunAsUser", value)
}

// GetRunAsUser gets the value of RunAsUser for the instance
func (instance *Win32_DCOMApplicationSetting) GetPropertyRunAsUser() (value string, err error) {
	retValue, err := instance.GetProperty("RunAsUser")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetServiceParameters sets the value of ServiceParameters for the instance
func (instance *Win32_DCOMApplicationSetting) SetPropertyServiceParameters(value string) (err error) {
	return instance.SetProperty("ServiceParameters", value)
}

// GetServiceParameters gets the value of ServiceParameters for the instance
func (instance *Win32_DCOMApplicationSetting) GetPropertyServiceParameters() (value string, err error) {
	retValue, err := instance.GetProperty("ServiceParameters")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUseSurrogate sets the value of UseSurrogate for the instance
func (instance *Win32_DCOMApplicationSetting) SetPropertyUseSurrogate(value bool) (err error) {
	return instance.SetProperty("UseSurrogate", value)
}

// GetUseSurrogate gets the value of UseSurrogate for the instance
func (instance *Win32_DCOMApplicationSetting) GetPropertyUseSurrogate() (value bool, err error) {
	retValue, err := instance.GetProperty("UseSurrogate")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="Descriptor" type="Win32_SecurityDescriptor "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_DCOMApplicationSetting) GetLaunchSecurityDescriptor( /* OUT */ Descriptor Win32_SecurityDescriptor) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetLaunchSecurityDescriptor")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Descriptor" type="Win32_SecurityDescriptor "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_DCOMApplicationSetting) SetLaunchSecurityDescriptor( /* IN */ Descriptor Win32_SecurityDescriptor) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetLaunchSecurityDescriptor", Descriptor)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Descriptor" type="Win32_SecurityDescriptor "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_DCOMApplicationSetting) GetAccessSecurityDescriptor( /* OUT */ Descriptor Win32_SecurityDescriptor) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetAccessSecurityDescriptor")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Descriptor" type="Win32_SecurityDescriptor "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_DCOMApplicationSetting) SetAccessSecurityDescriptor( /* IN */ Descriptor Win32_SecurityDescriptor) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetAccessSecurityDescriptor", Descriptor)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Descriptor" type="Win32_SecurityDescriptor "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_DCOMApplicationSetting) GetConfigurationSecurityDescriptor( /* OUT */ Descriptor Win32_SecurityDescriptor) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetConfigurationSecurityDescriptor")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Descriptor" type="Win32_SecurityDescriptor "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_DCOMApplicationSetting) SetConfigurationSecurityDescriptor( /* IN */ Descriptor Win32_SecurityDescriptor) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetConfigurationSecurityDescriptor", Descriptor)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
