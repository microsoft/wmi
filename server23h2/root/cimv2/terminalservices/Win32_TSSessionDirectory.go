// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.CIMV2.TerminalServices
//////////////////////////////////////////////
package terminalservices

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_TSSessionDirectory struct
type Win32_TSSessionDirectory struct {
	*CIM_Setting

	//
	GetLoadBalancingState uint32

	//
	GetServerWeight uint32

	//
	GetTSRedirectorMode uint32

	//
	PolicySourceLoadBalancing uint32

	//
	PolicySourceSessionDirectoryActive uint32

	//
	PolicySourceSessionDirectoryClusterName uint32

	//
	PolicySourceSessionDirectoryExposeServerIP uint32

	//
	PolicySourceSessionDirectoryLocation uint32

	//
	SessionDirectoryActive uint32

	//
	SessionDirectoryClusterName string

	//
	SessionDirectoryExposeServerIP uint32

	//
	SessionDirectoryIPAddress string

	//
	SessionDirectoryLocation string
}

func NewWin32_TSSessionDirectoryEx1(instance *cim.WmiInstance) (newInstance *Win32_TSSessionDirectory, err error) {
	tmp, err := NewCIM_SettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_TSSessionDirectory{
		CIM_Setting: tmp,
	}
	return
}

func NewWin32_TSSessionDirectoryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_TSSessionDirectory, err error) {
	tmp, err := NewCIM_SettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_TSSessionDirectory{
		CIM_Setting: tmp,
	}
	return
}

// SetGetLoadBalancingState sets the value of GetLoadBalancingState for the instance
func (instance *Win32_TSSessionDirectory) SetPropertyGetLoadBalancingState(value uint32) (err error) {
	return instance.SetProperty("GetLoadBalancingState", (value))
}

// GetGetLoadBalancingState gets the value of GetLoadBalancingState for the instance
func (instance *Win32_TSSessionDirectory) GetPropertyGetLoadBalancingState() (value uint32, err error) {
	retValue, err := instance.GetProperty("GetLoadBalancingState")
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

// SetGetServerWeight sets the value of GetServerWeight for the instance
func (instance *Win32_TSSessionDirectory) SetPropertyGetServerWeight(value uint32) (err error) {
	return instance.SetProperty("GetServerWeight", (value))
}

// GetGetServerWeight gets the value of GetServerWeight for the instance
func (instance *Win32_TSSessionDirectory) GetPropertyGetServerWeight() (value uint32, err error) {
	retValue, err := instance.GetProperty("GetServerWeight")
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

// SetGetTSRedirectorMode sets the value of GetTSRedirectorMode for the instance
func (instance *Win32_TSSessionDirectory) SetPropertyGetTSRedirectorMode(value uint32) (err error) {
	return instance.SetProperty("GetTSRedirectorMode", (value))
}

// GetGetTSRedirectorMode gets the value of GetTSRedirectorMode for the instance
func (instance *Win32_TSSessionDirectory) GetPropertyGetTSRedirectorMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("GetTSRedirectorMode")
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

// SetPolicySourceLoadBalancing sets the value of PolicySourceLoadBalancing for the instance
func (instance *Win32_TSSessionDirectory) SetPropertyPolicySourceLoadBalancing(value uint32) (err error) {
	return instance.SetProperty("PolicySourceLoadBalancing", (value))
}

// GetPolicySourceLoadBalancing gets the value of PolicySourceLoadBalancing for the instance
func (instance *Win32_TSSessionDirectory) GetPropertyPolicySourceLoadBalancing() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceLoadBalancing")
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

// SetPolicySourceSessionDirectoryActive sets the value of PolicySourceSessionDirectoryActive for the instance
func (instance *Win32_TSSessionDirectory) SetPropertyPolicySourceSessionDirectoryActive(value uint32) (err error) {
	return instance.SetProperty("PolicySourceSessionDirectoryActive", (value))
}

// GetPolicySourceSessionDirectoryActive gets the value of PolicySourceSessionDirectoryActive for the instance
func (instance *Win32_TSSessionDirectory) GetPropertyPolicySourceSessionDirectoryActive() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceSessionDirectoryActive")
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

// SetPolicySourceSessionDirectoryClusterName sets the value of PolicySourceSessionDirectoryClusterName for the instance
func (instance *Win32_TSSessionDirectory) SetPropertyPolicySourceSessionDirectoryClusterName(value uint32) (err error) {
	return instance.SetProperty("PolicySourceSessionDirectoryClusterName", (value))
}

// GetPolicySourceSessionDirectoryClusterName gets the value of PolicySourceSessionDirectoryClusterName for the instance
func (instance *Win32_TSSessionDirectory) GetPropertyPolicySourceSessionDirectoryClusterName() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceSessionDirectoryClusterName")
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

// SetPolicySourceSessionDirectoryExposeServerIP sets the value of PolicySourceSessionDirectoryExposeServerIP for the instance
func (instance *Win32_TSSessionDirectory) SetPropertyPolicySourceSessionDirectoryExposeServerIP(value uint32) (err error) {
	return instance.SetProperty("PolicySourceSessionDirectoryExposeServerIP", (value))
}

// GetPolicySourceSessionDirectoryExposeServerIP gets the value of PolicySourceSessionDirectoryExposeServerIP for the instance
func (instance *Win32_TSSessionDirectory) GetPropertyPolicySourceSessionDirectoryExposeServerIP() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceSessionDirectoryExposeServerIP")
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

// SetPolicySourceSessionDirectoryLocation sets the value of PolicySourceSessionDirectoryLocation for the instance
func (instance *Win32_TSSessionDirectory) SetPropertyPolicySourceSessionDirectoryLocation(value uint32) (err error) {
	return instance.SetProperty("PolicySourceSessionDirectoryLocation", (value))
}

// GetPolicySourceSessionDirectoryLocation gets the value of PolicySourceSessionDirectoryLocation for the instance
func (instance *Win32_TSSessionDirectory) GetPropertyPolicySourceSessionDirectoryLocation() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicySourceSessionDirectoryLocation")
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

// SetSessionDirectoryActive sets the value of SessionDirectoryActive for the instance
func (instance *Win32_TSSessionDirectory) SetPropertySessionDirectoryActive(value uint32) (err error) {
	return instance.SetProperty("SessionDirectoryActive", (value))
}

// GetSessionDirectoryActive gets the value of SessionDirectoryActive for the instance
func (instance *Win32_TSSessionDirectory) GetPropertySessionDirectoryActive() (value uint32, err error) {
	retValue, err := instance.GetProperty("SessionDirectoryActive")
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

// SetSessionDirectoryClusterName sets the value of SessionDirectoryClusterName for the instance
func (instance *Win32_TSSessionDirectory) SetPropertySessionDirectoryClusterName(value string) (err error) {
	return instance.SetProperty("SessionDirectoryClusterName", (value))
}

// GetSessionDirectoryClusterName gets the value of SessionDirectoryClusterName for the instance
func (instance *Win32_TSSessionDirectory) GetPropertySessionDirectoryClusterName() (value string, err error) {
	retValue, err := instance.GetProperty("SessionDirectoryClusterName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetSessionDirectoryExposeServerIP sets the value of SessionDirectoryExposeServerIP for the instance
func (instance *Win32_TSSessionDirectory) SetPropertySessionDirectoryExposeServerIP(value uint32) (err error) {
	return instance.SetProperty("SessionDirectoryExposeServerIP", (value))
}

// GetSessionDirectoryExposeServerIP gets the value of SessionDirectoryExposeServerIP for the instance
func (instance *Win32_TSSessionDirectory) GetPropertySessionDirectoryExposeServerIP() (value uint32, err error) {
	retValue, err := instance.GetProperty("SessionDirectoryExposeServerIP")
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

// SetSessionDirectoryIPAddress sets the value of SessionDirectoryIPAddress for the instance
func (instance *Win32_TSSessionDirectory) SetPropertySessionDirectoryIPAddress(value string) (err error) {
	return instance.SetProperty("SessionDirectoryIPAddress", (value))
}

// GetSessionDirectoryIPAddress gets the value of SessionDirectoryIPAddress for the instance
func (instance *Win32_TSSessionDirectory) GetPropertySessionDirectoryIPAddress() (value string, err error) {
	retValue, err := instance.GetProperty("SessionDirectoryIPAddress")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetSessionDirectoryLocation sets the value of SessionDirectoryLocation for the instance
func (instance *Win32_TSSessionDirectory) SetPropertySessionDirectoryLocation(value string) (err error) {
	return instance.SetProperty("SessionDirectoryLocation", (value))
}

// GetSessionDirectoryLocation gets the value of SessionDirectoryLocation for the instance
func (instance *Win32_TSSessionDirectory) GetPropertySessionDirectoryLocation() (value string, err error) {
	retValue, err := instance.GetProperty("SessionDirectoryLocation")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

//

// <param name="PropertyName" type="string "></param>
// <param name="Value" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSSessionDirectory) SetSessionDirectoryProperty( /* IN */ PropertyName string,
	/* IN */ Value string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetSessionDirectoryProperty", PropertyName, Value)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="SessionDirectoryActive" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSSessionDirectory) SetSessionDirectoryActive( /* IN */ SessionDirectoryActive uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetSessionDirectoryActive", SessionDirectoryActive)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="SessionDirectoryExposeServerIP" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSSessionDirectory) SetSessionDirectoryExposeServerIP( /* IN */ SessionDirectoryExposeServerIP uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetSessionDirectoryExposeServerIP", SessionDirectoryExposeServerIP)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="fTokenRedirection" type="uint32 "></param>

// <param name="AdapterNames" type="string []"></param>
// <param name="IPAddresses" type="string []"></param>
// <param name="NetConName" type="string []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSSessionDirectory) GetRedirectableAddresses( /* IN */ fTokenRedirection uint32,
	/* OUT */ IPAddresses []string,
	/* OUT */ AdapterNames []string,
	/* OUT */ NetConName []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetRedirectableAddresses", fTokenRedirection)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="fTokenRedirection" type="uint32 "></param>
// <param name="IPAddresses" type="string []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSSessionDirectory) GetCurrentRedirectableAddresses( /* OUT */ fTokenRedirection uint32,
	/* OUT */ IPAddresses []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetCurrentRedirectableAddresses")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="fTokenRedirection" type="uint32 "></param>
// <param name="IPAddresses" type="string []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSSessionDirectory) SetCurrentRedirectableAddresses( /* IN */ fTokenRedirection uint32,
	/* IN */ IPAddresses []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetCurrentRedirectableAddresses", fTokenRedirection, IPAddresses)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ServerName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSSessionDirectory) PingSessionDirectory( /* IN */ ServerName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("PingSessionDirectory", ServerName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="StateValue" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSSessionDirectory) SetLoadBalancingState( /* IN */ StateValue uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetLoadBalancingState", StateValue)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ServerWeightValue" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSSessionDirectory) SetServerWeight( /* IN */ ServerWeightValue uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetServerWeight", ServerWeightValue)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="UserDiskMaxSizeInGB" type="uint32 "></param>
// <param name="UserDisksStorageUrl" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSSessionDirectory) CreateUserDiskTemplate( /* IN */ UserDisksStorageUrl string,
	/* IN */ UserDiskMaxSizeInGB uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("CreateUserDiskTemplate", UserDisksStorageUrl, UserDiskMaxSizeInGB)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="UvhdRoamingPolicyXml" type="string "></param>
// <param name="UvhdShareUrl" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSSessionDirectory) EnableUserVhd( /* IN */ UvhdShareUrl string,
	/* IN */ UvhdRoamingPolicyXml string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("EnableUserVhd", UvhdShareUrl, UvhdRoamingPolicyXml)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_TSSessionDirectory) DisableUserVhd() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("DisableUserVhd")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
