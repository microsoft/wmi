// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.CIMV2.TerminalServices
//////////////////////////////////////////////
package terminalservices

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_TSSessionSetting struct
type Win32_TSSessionSetting struct {
	*Win32_TerminalSetting

	//
	ActiveSessionLimit uint32

	//
	BrokenConnectionAction uint32

	//
	BrokenConnectionPolicy uint32

	//
	DisconnectedSessionLimit uint32

	//
	EnableTimeoutWarning uint32

	//
	IdleSessionLimit uint32

	//
	PolicySourceActiveSessionLimit uint32

	//
	PolicySourceBrokenConnectionAction uint32

	//
	PolicySourceDisconnectedSessionLimit uint32

	//
	PolicySourceIdleSessionLimit uint32

	//
	PolicySourceReconnectionPolicy uint32

	//
	ReconnectionPolicy uint32

	//
	TimeLimitPolicy uint32
}

func NewWin32_TSSessionSettingEx1(instance *cim.WmiInstance) (newInstance *Win32_TSSessionSetting, err error) {
	tmp, err := NewWin32_TerminalSettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_TSSessionSetting{
		Win32_TerminalSetting: tmp,
	}
	return
}

func NewWin32_TSSessionSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_TSSessionSetting, err error) {
	tmp, err := NewWin32_TerminalSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_TSSessionSetting{
		Win32_TerminalSetting: tmp,
	}
	return
}

// SetActiveSessionLimit sets the value of ActiveSessionLimit for the instance
func (instance *Win32_TSSessionSetting) SetPropertyActiveSessionLimit(value uint32) (err error) {
	return instance.SetProperty("ActiveSessionLimit", (value))
}

// GetActiveSessionLimit gets the value of ActiveSessionLimit for the instance
func (instance *Win32_TSSessionSetting) GetPropertyActiveSessionLimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("ActiveSessionLimit")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetBrokenConnectionAction sets the value of BrokenConnectionAction for the instance
func (instance *Win32_TSSessionSetting) SetPropertyBrokenConnectionAction(value uint32) (err error) {
	return instance.SetProperty("BrokenConnectionAction", (value))
}

// GetBrokenConnectionAction gets the value of BrokenConnectionAction for the instance
func (instance *Win32_TSSessionSetting) GetPropertyBrokenConnectionAction() (value uint32, err error) {
	retValue, err := instance.GetProperty("BrokenConnectionAction")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetBrokenConnectionPolicy sets the value of BrokenConnectionPolicy for the instance
func (instance *Win32_TSSessionSetting) SetPropertyBrokenConnectionPolicy(value uint32) (err error) {
	return instance.SetProperty("BrokenConnectionPolicy", (value))
}

// GetBrokenConnectionPolicy gets the value of BrokenConnectionPolicy for the instance
func (instance *Win32_TSSessionSetting) GetPropertyBrokenConnectionPolicy() (value uint32, err error) {
	retValue, err := instance.GetProperty("BrokenConnectionPolicy")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetDisconnectedSessionLimit sets the value of DisconnectedSessionLimit for the instance
func (instance *Win32_TSSessionSetting) SetPropertyDisconnectedSessionLimit(value uint32) (err error) {
	return instance.SetProperty("DisconnectedSessionLimit", (value))
}

// GetDisconnectedSessionLimit gets the value of DisconnectedSessionLimit for the instance
func (instance *Win32_TSSessionSetting) GetPropertyDisconnectedSessionLimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("DisconnectedSessionLimit")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetEnableTimeoutWarning sets the value of EnableTimeoutWarning for the instance
func (instance *Win32_TSSessionSetting) SetPropertyEnableTimeoutWarning(value uint32) (err error) {
	return instance.SetProperty("EnableTimeoutWarning", (value))
}

// GetEnableTimeoutWarning gets the value of EnableTimeoutWarning for the instance
func (instance *Win32_TSSessionSetting) GetPropertyEnableTimeoutWarning() (value uint32, err error) {
	retValue, err := instance.GetProperty("EnableTimeoutWarning")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetIdleSessionLimit sets the value of IdleSessionLimit for the instance
func (instance *Win32_TSSessionSetting) SetPropertyIdleSessionLimit(value uint32) (err error) {
	return instance.SetProperty("IdleSessionLimit", (value))
}

// GetIdleSessionLimit gets the value of IdleSessionLimit for the instance
func (instance *Win32_TSSessionSetting) GetPropertyIdleSessionLimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("IdleSessionLimit")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetPolicySourceActiveSessionLimit sets the value of PolicySourceActiveSessionLimit for the instance
func (instance *Win32_TSSessionSetting) SetPropertyPolicySourceActiveSessionLimit(value uint32) (err error) {
	return instance.SetProperty("PolicySourceActiveSessionLimit", (value))
}

// GetPolicySourceActiveSessionLimit gets the value of PolicySourceActiveSessionLimit for the instance
func (instance *Win32_TSSessionSetting) GetPropertyPolicySourceActiveSessionLimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceActiveSessionLimit")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetPolicySourceBrokenConnectionAction sets the value of PolicySourceBrokenConnectionAction for the instance
func (instance *Win32_TSSessionSetting) SetPropertyPolicySourceBrokenConnectionAction(value uint32) (err error) {
	return instance.SetProperty("PolicySourceBrokenConnectionAction", (value))
}

// GetPolicySourceBrokenConnectionAction gets the value of PolicySourceBrokenConnectionAction for the instance
func (instance *Win32_TSSessionSetting) GetPropertyPolicySourceBrokenConnectionAction() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceBrokenConnectionAction")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetPolicySourceDisconnectedSessionLimit sets the value of PolicySourceDisconnectedSessionLimit for the instance
func (instance *Win32_TSSessionSetting) SetPropertyPolicySourceDisconnectedSessionLimit(value uint32) (err error) {
	return instance.SetProperty("PolicySourceDisconnectedSessionLimit", (value))
}

// GetPolicySourceDisconnectedSessionLimit gets the value of PolicySourceDisconnectedSessionLimit for the instance
func (instance *Win32_TSSessionSetting) GetPropertyPolicySourceDisconnectedSessionLimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceDisconnectedSessionLimit")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetPolicySourceIdleSessionLimit sets the value of PolicySourceIdleSessionLimit for the instance
func (instance *Win32_TSSessionSetting) SetPropertyPolicySourceIdleSessionLimit(value uint32) (err error) {
	return instance.SetProperty("PolicySourceIdleSessionLimit", (value))
}

// GetPolicySourceIdleSessionLimit gets the value of PolicySourceIdleSessionLimit for the instance
func (instance *Win32_TSSessionSetting) GetPropertyPolicySourceIdleSessionLimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceIdleSessionLimit")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetPolicySourceReconnectionPolicy sets the value of PolicySourceReconnectionPolicy for the instance
func (instance *Win32_TSSessionSetting) SetPropertyPolicySourceReconnectionPolicy(value uint32) (err error) {
	return instance.SetProperty("PolicySourceReconnectionPolicy", (value))
}

// GetPolicySourceReconnectionPolicy gets the value of PolicySourceReconnectionPolicy for the instance
func (instance *Win32_TSSessionSetting) GetPropertyPolicySourceReconnectionPolicy() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceReconnectionPolicy")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetReconnectionPolicy sets the value of ReconnectionPolicy for the instance
func (instance *Win32_TSSessionSetting) SetPropertyReconnectionPolicy(value uint32) (err error) {
	return instance.SetProperty("ReconnectionPolicy", (value))
}

// GetReconnectionPolicy gets the value of ReconnectionPolicy for the instance
func (instance *Win32_TSSessionSetting) GetPropertyReconnectionPolicy() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReconnectionPolicy")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetTimeLimitPolicy sets the value of TimeLimitPolicy for the instance
func (instance *Win32_TSSessionSetting) SetPropertyTimeLimitPolicy(value uint32) (err error) {
	return instance.SetProperty("TimeLimitPolicy", (value))
}

// GetTimeLimitPolicy gets the value of TimeLimitPolicy for the instance
func (instance *Win32_TSSessionSetting) GetPropertyTimeLimitPolicy() (value uint32, err error) {
	retValue, err := instance.GetProperty("TimeLimitPolicy")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

//

// <param name="SessionLimitType" type="string "></param>
// <param name="ValueLimit" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSSessionSetting) TimeLimit( /* IN */ SessionLimitType string,
	/* IN */ ValueLimit uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("TimeLimit", SessionLimitType, ValueLimit)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="BrokenConnectionAction" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSSessionSetting) BrokenConnection( /* IN */ BrokenConnectionAction uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("BrokenConnection", BrokenConnectionAction)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
