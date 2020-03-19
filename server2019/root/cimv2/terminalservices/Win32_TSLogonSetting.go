// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2.TerminalServices
//////////////////////////////////////////////
package terminalservices

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_TSLogonSetting struct
type Win32_TSLogonSetting struct {
	*Win32_TerminalSetting

	//
	ClientLogonInfoPolicy uint32

	//
	Domain string

	//
	PolicySourceDomain uint32

	//
	PolicySourcePromptForPassword uint32

	//
	PolicySourceUserName uint32

	//
	PromptForPassword uint32

	//
	UserName string
}

func NewWin32_TSLogonSettingEx1(instance *cim.WmiInstance) (newInstance *Win32_TSLogonSetting, err error) {
	tmp, err := NewWin32_TerminalSettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_TSLogonSetting{
		Win32_TerminalSetting: tmp,
	}
	return
}

func NewWin32_TSLogonSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_TSLogonSetting, err error) {
	tmp, err := NewWin32_TerminalSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_TSLogonSetting{
		Win32_TerminalSetting: tmp,
	}
	return
}

// SetClientLogonInfoPolicy sets the value of ClientLogonInfoPolicy for the instance
func (instance *Win32_TSLogonSetting) SetPropertyClientLogonInfoPolicy(value uint32) (err error) {
	return instance.SetProperty("ClientLogonInfoPolicy", value)
}

// GetClientLogonInfoPolicy gets the value of ClientLogonInfoPolicy for the instance
func (instance *Win32_TSLogonSetting) GetPropertyClientLogonInfoPolicy() (value uint32, err error) {
	retValue, err := instance.GetProperty("ClientLogonInfoPolicy")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDomain sets the value of Domain for the instance
func (instance *Win32_TSLogonSetting) SetPropertyDomain(value string) (err error) {
	return instance.SetProperty("Domain", value)
}

// GetDomain gets the value of Domain for the instance
func (instance *Win32_TSLogonSetting) GetPropertyDomain() (value string, err error) {
	retValue, err := instance.GetProperty("Domain")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicySourceDomain sets the value of PolicySourceDomain for the instance
func (instance *Win32_TSLogonSetting) SetPropertyPolicySourceDomain(value uint32) (err error) {
	return instance.SetProperty("PolicySourceDomain", value)
}

// GetPolicySourceDomain gets the value of PolicySourceDomain for the instance
func (instance *Win32_TSLogonSetting) GetPropertyPolicySourceDomain() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceDomain")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicySourcePromptForPassword sets the value of PolicySourcePromptForPassword for the instance
func (instance *Win32_TSLogonSetting) SetPropertyPolicySourcePromptForPassword(value uint32) (err error) {
	return instance.SetProperty("PolicySourcePromptForPassword", value)
}

// GetPolicySourcePromptForPassword gets the value of PolicySourcePromptForPassword for the instance
func (instance *Win32_TSLogonSetting) GetPropertyPolicySourcePromptForPassword() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourcePromptForPassword")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicySourceUserName sets the value of PolicySourceUserName for the instance
func (instance *Win32_TSLogonSetting) SetPropertyPolicySourceUserName(value uint32) (err error) {
	return instance.SetProperty("PolicySourceUserName", value)
}

// GetPolicySourceUserName gets the value of PolicySourceUserName for the instance
func (instance *Win32_TSLogonSetting) GetPropertyPolicySourceUserName() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceUserName")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPromptForPassword sets the value of PromptForPassword for the instance
func (instance *Win32_TSLogonSetting) SetPropertyPromptForPassword(value uint32) (err error) {
	return instance.SetProperty("PromptForPassword", value)
}

// GetPromptForPassword gets the value of PromptForPassword for the instance
func (instance *Win32_TSLogonSetting) GetPropertyPromptForPassword() (value uint32, err error) {
	retValue, err := instance.GetProperty("PromptForPassword")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUserName sets the value of UserName for the instance
func (instance *Win32_TSLogonSetting) SetPropertyUserName(value string) (err error) {
	return instance.SetProperty("UserName", value)
}

// GetUserName gets the value of UserName for the instance
func (instance *Win32_TSLogonSetting) GetPropertyUserName() (value string, err error) {
	retValue, err := instance.GetProperty("UserName")
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

// <param name="Domain" type="string "></param>
// <param name="Password" type="string "></param>
// <param name="UserName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSLogonSetting) ExplicitLogon( /* IN */ UserName string,
	/* IN */ Domain string,
	/* IN */ Password string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ExplicitLogon", UserName, Domain, Password)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="PromptForPassword" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSLogonSetting) SetPromptForPassword( /* IN */ PromptForPassword uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetPromptForPassword", PromptForPassword)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
