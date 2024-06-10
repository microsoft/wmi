// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetConSecRule struct
type MSFT_NetConSecRule struct {
	*MSFT_NetSARule

	//
	AllowSetKey bool

	//
	AllowWatchKey bool

	//
	BypassTunnelIfEncrypted bool

	//
	InboundSecurity uint16

	//
	KeyModule uint16

	//
	LocalTunnelEndpoint []string

	//
	Machines string

	//
	MaxReturnPathLifetimeSeconds uint32

	//
	Mode uint16

	//
	OutboundSecurity uint16

	//
	RemoteTunnelEndpoint []string

	//
	RemoteTunnelEndpointDNSName string

	//
	RequireAuthorization bool

	//
	Users string
}

func NewMSFT_NetConSecRuleEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetConSecRule, err error) {
	tmp, err := NewMSFT_NetSARuleEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetConSecRule{
		MSFT_NetSARule: tmp,
	}
	return
}

func NewMSFT_NetConSecRuleEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetConSecRule, err error) {
	tmp, err := NewMSFT_NetSARuleEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetConSecRule{
		MSFT_NetSARule: tmp,
	}
	return
}

// SetAllowSetKey sets the value of AllowSetKey for the instance
func (instance *MSFT_NetConSecRule) SetPropertyAllowSetKey(value bool) (err error) {
	return instance.SetProperty("AllowSetKey", (value))
}

// GetAllowSetKey gets the value of AllowSetKey for the instance
func (instance *MSFT_NetConSecRule) GetPropertyAllowSetKey() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowSetKey")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetAllowWatchKey sets the value of AllowWatchKey for the instance
func (instance *MSFT_NetConSecRule) SetPropertyAllowWatchKey(value bool) (err error) {
	return instance.SetProperty("AllowWatchKey", (value))
}

// GetAllowWatchKey gets the value of AllowWatchKey for the instance
func (instance *MSFT_NetConSecRule) GetPropertyAllowWatchKey() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowWatchKey")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetBypassTunnelIfEncrypted sets the value of BypassTunnelIfEncrypted for the instance
func (instance *MSFT_NetConSecRule) SetPropertyBypassTunnelIfEncrypted(value bool) (err error) {
	return instance.SetProperty("BypassTunnelIfEncrypted", (value))
}

// GetBypassTunnelIfEncrypted gets the value of BypassTunnelIfEncrypted for the instance
func (instance *MSFT_NetConSecRule) GetPropertyBypassTunnelIfEncrypted() (value bool, err error) {
	retValue, err := instance.GetProperty("BypassTunnelIfEncrypted")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetInboundSecurity sets the value of InboundSecurity for the instance
func (instance *MSFT_NetConSecRule) SetPropertyInboundSecurity(value uint16) (err error) {
	return instance.SetProperty("InboundSecurity", (value))
}

// GetInboundSecurity gets the value of InboundSecurity for the instance
func (instance *MSFT_NetConSecRule) GetPropertyInboundSecurity() (value uint16, err error) {
	retValue, err := instance.GetProperty("InboundSecurity")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetKeyModule sets the value of KeyModule for the instance
func (instance *MSFT_NetConSecRule) SetPropertyKeyModule(value uint16) (err error) {
	return instance.SetProperty("KeyModule", (value))
}

// GetKeyModule gets the value of KeyModule for the instance
func (instance *MSFT_NetConSecRule) GetPropertyKeyModule() (value uint16, err error) {
	retValue, err := instance.GetProperty("KeyModule")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetLocalTunnelEndpoint sets the value of LocalTunnelEndpoint for the instance
func (instance *MSFT_NetConSecRule) SetPropertyLocalTunnelEndpoint(value []string) (err error) {
	return instance.SetProperty("LocalTunnelEndpoint", (value))
}

// GetLocalTunnelEndpoint gets the value of LocalTunnelEndpoint for the instance
func (instance *MSFT_NetConSecRule) GetPropertyLocalTunnelEndpoint() (value []string, err error) {
	retValue, err := instance.GetProperty("LocalTunnelEndpoint")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetMachines sets the value of Machines for the instance
func (instance *MSFT_NetConSecRule) SetPropertyMachines(value string) (err error) {
	return instance.SetProperty("Machines", (value))
}

// GetMachines gets the value of Machines for the instance
func (instance *MSFT_NetConSecRule) GetPropertyMachines() (value string, err error) {
	retValue, err := instance.GetProperty("Machines")
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

// SetMaxReturnPathLifetimeSeconds sets the value of MaxReturnPathLifetimeSeconds for the instance
func (instance *MSFT_NetConSecRule) SetPropertyMaxReturnPathLifetimeSeconds(value uint32) (err error) {
	return instance.SetProperty("MaxReturnPathLifetimeSeconds", (value))
}

// GetMaxReturnPathLifetimeSeconds gets the value of MaxReturnPathLifetimeSeconds for the instance
func (instance *MSFT_NetConSecRule) GetPropertyMaxReturnPathLifetimeSeconds() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxReturnPathLifetimeSeconds")
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

// SetMode sets the value of Mode for the instance
func (instance *MSFT_NetConSecRule) SetPropertyMode(value uint16) (err error) {
	return instance.SetProperty("Mode", (value))
}

// GetMode gets the value of Mode for the instance
func (instance *MSFT_NetConSecRule) GetPropertyMode() (value uint16, err error) {
	retValue, err := instance.GetProperty("Mode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetOutboundSecurity sets the value of OutboundSecurity for the instance
func (instance *MSFT_NetConSecRule) SetPropertyOutboundSecurity(value uint16) (err error) {
	return instance.SetProperty("OutboundSecurity", (value))
}

// GetOutboundSecurity gets the value of OutboundSecurity for the instance
func (instance *MSFT_NetConSecRule) GetPropertyOutboundSecurity() (value uint16, err error) {
	retValue, err := instance.GetProperty("OutboundSecurity")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetRemoteTunnelEndpoint sets the value of RemoteTunnelEndpoint for the instance
func (instance *MSFT_NetConSecRule) SetPropertyRemoteTunnelEndpoint(value []string) (err error) {
	return instance.SetProperty("RemoteTunnelEndpoint", (value))
}

// GetRemoteTunnelEndpoint gets the value of RemoteTunnelEndpoint for the instance
func (instance *MSFT_NetConSecRule) GetPropertyRemoteTunnelEndpoint() (value []string, err error) {
	retValue, err := instance.GetProperty("RemoteTunnelEndpoint")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetRemoteTunnelEndpointDNSName sets the value of RemoteTunnelEndpointDNSName for the instance
func (instance *MSFT_NetConSecRule) SetPropertyRemoteTunnelEndpointDNSName(value string) (err error) {
	return instance.SetProperty("RemoteTunnelEndpointDNSName", (value))
}

// GetRemoteTunnelEndpointDNSName gets the value of RemoteTunnelEndpointDNSName for the instance
func (instance *MSFT_NetConSecRule) GetPropertyRemoteTunnelEndpointDNSName() (value string, err error) {
	retValue, err := instance.GetProperty("RemoteTunnelEndpointDNSName")
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

// SetRequireAuthorization sets the value of RequireAuthorization for the instance
func (instance *MSFT_NetConSecRule) SetPropertyRequireAuthorization(value bool) (err error) {
	return instance.SetProperty("RequireAuthorization", (value))
}

// GetRequireAuthorization gets the value of RequireAuthorization for the instance
func (instance *MSFT_NetConSecRule) GetPropertyRequireAuthorization() (value bool, err error) {
	retValue, err := instance.GetProperty("RequireAuthorization")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetUsers sets the value of Users for the instance
func (instance *MSFT_NetConSecRule) SetPropertyUsers(value string) (err error) {
	return instance.SetProperty("Users", (value))
}

// GetUsers gets the value of Users for the instance
func (instance *MSFT_NetConSecRule) GetPropertyUsers() (value string, err error) {
	retValue, err := instance.GetProperty("Users")
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

// <param name="AddressType" type="uint16 "></param>
// <param name="DnsServers" type="string []"></param>
// <param name="Domains" type="string []"></param>
// <param name="EndpointType" type="uint16 "></param>
// <param name="Servers" type="string []"></param>

// <param name="Output" type="MSFT_NetSecDeltaCollection []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetConSecRule) SyncPolicyDelta( /* IN */ Servers []string,
	/* IN */ Domains []string,
	/* IN */ EndpointType uint16,
	/* IN */ AddressType uint16,
	/* IN */ DnsServers []string,
	/* OUT */ Output []MSFT_NetSecDeltaCollection) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SyncPolicyDelta", Servers, Domains, EndpointType, AddressType, DnsServers)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Action" type="uint16 "></param>
// <param name="EndpointType" type="uint16 "></param>
// <param name="IPv4Addresses" type="string []"></param>
// <param name="IPv6Addresses" type="string []"></param>
// <param name="PassThru" type="bool "></param>

// <param name="Output" type="MSFT_NetConSecRule []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetConSecRule) SetPolicyDelta( /* IN */ Action uint16,
	/* IN */ IPv6Addresses []string,
	/* IN */ IPv4Addresses []string,
	/* IN */ EndpointType uint16,
	/* IN */ PassThru bool,
	/* OUT */ Output []MSFT_NetConSecRule) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetPolicyDelta", Action, IPv6Addresses, IPv4Addresses, EndpointType, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Dependents" type="CIM_ManagedSystemElement []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetConSecRule) EnumerateFull( /* OUT */ Dependents []CIM_ManagedSystemElement) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("EnumerateFull")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="LocalAddress" type="string "></param>
// <param name="LocalPort" type="uint16 "></param>
// <param name="Protocol" type="string "></param>
// <param name="RemoteAddress" type="string "></param>
// <param name="RemotePort" type="uint16 "></param>

// <param name="CmdletOutput" type="MSFT_NetConSecRule []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetConSecRule) Find( /* IN */ LocalAddress string,
	/* IN */ RemoteAddress string,
	/* IN */ Protocol string,
	/* IN */ LocalPort uint16,
	/* IN */ RemotePort uint16,
	/* OUT */ CmdletOutput []MSFT_NetConSecRule) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Find", LocalAddress, RemoteAddress, Protocol, LocalPort, RemotePort)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetConSecRule) Enable() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Enable")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetConSecRule) Disable() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Disable")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="NewName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetConSecRule) Rename( /* IN */ NewName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Rename", NewName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="NewGPOSession" type="string "></param>
// <param name="NewName" type="string "></param>
// <param name="NewPolicyStore" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetConSecRule) CloneObject( /* IN */ NewName string,
	/* IN */ NewPolicyStore string,
	/* IN */ NewGPOSession string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("CloneObject", NewName, NewPolicyStore, NewGPOSession)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
