// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.Gateway
//
// ////////////////////////////////////////////
package gateway

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// PS_GatewayTunnel struct
type PS_GatewayTunnel struct {
	*cim.WmiInstance

	//
	AdminStatus bool

	//
	AuthenticationMethod uint32

	//
	AutoConnectEnabled bool

	//
	ConnectivityState uint32

	//
	ConnectReason uint32

	//
	ConnectVersion uint64

	//
	Dhgroup uint32

	//
	DisconnectReason string

	//
	DpdTimeout uint32

	//
	EnableQoS uint32

	//
	EncryptionType uint32

	//
	IkeEncryption uint32

	//
	IkeIntegrity uint32

	//
	IpsecEncryption uint32

	//
	IpsecIntegrity uint32

	//
	IsConnectionNATed bool

	//
	LastConnectError uint32

	//
	LastConnectErrorString string

	//
	LocalExemptRoutes string

	//
	LocalIp string

	//
	LocalSubnets string

	//
	MMSALifeTimeInSeconds uint32

	//
	Pfsgroup uint32

	//
	ProtocolType uint32

	//
	RemoteExemptRoutes string

	//
	RemoteIp string

	//
	RemoteSubnets string

	//
	Routes string

	//
	RoutingDomainName string

	//
	RxBandwidthKbps uint64

	//
	SALifeKilloBytes uint32

	//
	SALifeTimeInSeconds uint32

	//
	Standby bool

	//
	TunnelId string

	//
	TunnelIdentity string

	//
	TunnelName string

	//
	TxBandwidthKbps uint64

	//
	UseNarrowTrafficSelectors bool
}

func NewPS_GatewayTunnelEx1(instance *cim.WmiInstance) (newInstance *PS_GatewayTunnel, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_GatewayTunnel{
		WmiInstance: tmp,
	}
	return
}

func NewPS_GatewayTunnelEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_GatewayTunnel, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_GatewayTunnel{
		WmiInstance: tmp,
	}
	return
}

// SetAdminStatus sets the value of AdminStatus for the instance
func (instance *PS_GatewayTunnel) SetPropertyAdminStatus(value bool) (err error) {
	return instance.SetProperty("AdminStatus", (value))
}

// GetAdminStatus gets the value of AdminStatus for the instance
func (instance *PS_GatewayTunnel) GetPropertyAdminStatus() (value bool, err error) {
	retValue, err := instance.GetProperty("AdminStatus")
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

// SetAuthenticationMethod sets the value of AuthenticationMethod for the instance
func (instance *PS_GatewayTunnel) SetPropertyAuthenticationMethod(value uint32) (err error) {
	return instance.SetProperty("AuthenticationMethod", (value))
}

// GetAuthenticationMethod gets the value of AuthenticationMethod for the instance
func (instance *PS_GatewayTunnel) GetPropertyAuthenticationMethod() (value uint32, err error) {
	retValue, err := instance.GetProperty("AuthenticationMethod")
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

// SetAutoConnectEnabled sets the value of AutoConnectEnabled for the instance
func (instance *PS_GatewayTunnel) SetPropertyAutoConnectEnabled(value bool) (err error) {
	return instance.SetProperty("AutoConnectEnabled", (value))
}

// GetAutoConnectEnabled gets the value of AutoConnectEnabled for the instance
func (instance *PS_GatewayTunnel) GetPropertyAutoConnectEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("AutoConnectEnabled")
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

// SetConnectivityState sets the value of ConnectivityState for the instance
func (instance *PS_GatewayTunnel) SetPropertyConnectivityState(value uint32) (err error) {
	return instance.SetProperty("ConnectivityState", (value))
}

// GetConnectivityState gets the value of ConnectivityState for the instance
func (instance *PS_GatewayTunnel) GetPropertyConnectivityState() (value uint32, err error) {
	retValue, err := instance.GetProperty("ConnectivityState")
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

// SetConnectReason sets the value of ConnectReason for the instance
func (instance *PS_GatewayTunnel) SetPropertyConnectReason(value uint32) (err error) {
	return instance.SetProperty("ConnectReason", (value))
}

// GetConnectReason gets the value of ConnectReason for the instance
func (instance *PS_GatewayTunnel) GetPropertyConnectReason() (value uint32, err error) {
	retValue, err := instance.GetProperty("ConnectReason")
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

// SetConnectVersion sets the value of ConnectVersion for the instance
func (instance *PS_GatewayTunnel) SetPropertyConnectVersion(value uint64) (err error) {
	return instance.SetProperty("ConnectVersion", (value))
}

// GetConnectVersion gets the value of ConnectVersion for the instance
func (instance *PS_GatewayTunnel) GetPropertyConnectVersion() (value uint64, err error) {
	retValue, err := instance.GetProperty("ConnectVersion")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetDhgroup sets the value of Dhgroup for the instance
func (instance *PS_GatewayTunnel) SetPropertyDhgroup(value uint32) (err error) {
	return instance.SetProperty("Dhgroup", (value))
}

// GetDhgroup gets the value of Dhgroup for the instance
func (instance *PS_GatewayTunnel) GetPropertyDhgroup() (value uint32, err error) {
	retValue, err := instance.GetProperty("Dhgroup")
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

// SetDisconnectReason sets the value of DisconnectReason for the instance
func (instance *PS_GatewayTunnel) SetPropertyDisconnectReason(value string) (err error) {
	return instance.SetProperty("DisconnectReason", (value))
}

// GetDisconnectReason gets the value of DisconnectReason for the instance
func (instance *PS_GatewayTunnel) GetPropertyDisconnectReason() (value string, err error) {
	retValue, err := instance.GetProperty("DisconnectReason")
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

// SetDpdTimeout sets the value of DpdTimeout for the instance
func (instance *PS_GatewayTunnel) SetPropertyDpdTimeout(value uint32) (err error) {
	return instance.SetProperty("DpdTimeout", (value))
}

// GetDpdTimeout gets the value of DpdTimeout for the instance
func (instance *PS_GatewayTunnel) GetPropertyDpdTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("DpdTimeout")
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

// SetEnableQoS sets the value of EnableQoS for the instance
func (instance *PS_GatewayTunnel) SetPropertyEnableQoS(value uint32) (err error) {
	return instance.SetProperty("EnableQoS", (value))
}

// GetEnableQoS gets the value of EnableQoS for the instance
func (instance *PS_GatewayTunnel) GetPropertyEnableQoS() (value uint32, err error) {
	retValue, err := instance.GetProperty("EnableQoS")
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

// SetEncryptionType sets the value of EncryptionType for the instance
func (instance *PS_GatewayTunnel) SetPropertyEncryptionType(value uint32) (err error) {
	return instance.SetProperty("EncryptionType", (value))
}

// GetEncryptionType gets the value of EncryptionType for the instance
func (instance *PS_GatewayTunnel) GetPropertyEncryptionType() (value uint32, err error) {
	retValue, err := instance.GetProperty("EncryptionType")
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

// SetIkeEncryption sets the value of IkeEncryption for the instance
func (instance *PS_GatewayTunnel) SetPropertyIkeEncryption(value uint32) (err error) {
	return instance.SetProperty("IkeEncryption", (value))
}

// GetIkeEncryption gets the value of IkeEncryption for the instance
func (instance *PS_GatewayTunnel) GetPropertyIkeEncryption() (value uint32, err error) {
	retValue, err := instance.GetProperty("IkeEncryption")
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

// SetIkeIntegrity sets the value of IkeIntegrity for the instance
func (instance *PS_GatewayTunnel) SetPropertyIkeIntegrity(value uint32) (err error) {
	return instance.SetProperty("IkeIntegrity", (value))
}

// GetIkeIntegrity gets the value of IkeIntegrity for the instance
func (instance *PS_GatewayTunnel) GetPropertyIkeIntegrity() (value uint32, err error) {
	retValue, err := instance.GetProperty("IkeIntegrity")
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

// SetIpsecEncryption sets the value of IpsecEncryption for the instance
func (instance *PS_GatewayTunnel) SetPropertyIpsecEncryption(value uint32) (err error) {
	return instance.SetProperty("IpsecEncryption", (value))
}

// GetIpsecEncryption gets the value of IpsecEncryption for the instance
func (instance *PS_GatewayTunnel) GetPropertyIpsecEncryption() (value uint32, err error) {
	retValue, err := instance.GetProperty("IpsecEncryption")
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

// SetIpsecIntegrity sets the value of IpsecIntegrity for the instance
func (instance *PS_GatewayTunnel) SetPropertyIpsecIntegrity(value uint32) (err error) {
	return instance.SetProperty("IpsecIntegrity", (value))
}

// GetIpsecIntegrity gets the value of IpsecIntegrity for the instance
func (instance *PS_GatewayTunnel) GetPropertyIpsecIntegrity() (value uint32, err error) {
	retValue, err := instance.GetProperty("IpsecIntegrity")
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

// SetIsConnectionNATed sets the value of IsConnectionNATed for the instance
func (instance *PS_GatewayTunnel) SetPropertyIsConnectionNATed(value bool) (err error) {
	return instance.SetProperty("IsConnectionNATed", (value))
}

// GetIsConnectionNATed gets the value of IsConnectionNATed for the instance
func (instance *PS_GatewayTunnel) GetPropertyIsConnectionNATed() (value bool, err error) {
	retValue, err := instance.GetProperty("IsConnectionNATed")
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

// SetLastConnectError sets the value of LastConnectError for the instance
func (instance *PS_GatewayTunnel) SetPropertyLastConnectError(value uint32) (err error) {
	return instance.SetProperty("LastConnectError", (value))
}

// GetLastConnectError gets the value of LastConnectError for the instance
func (instance *PS_GatewayTunnel) GetPropertyLastConnectError() (value uint32, err error) {
	retValue, err := instance.GetProperty("LastConnectError")
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

// SetLastConnectErrorString sets the value of LastConnectErrorString for the instance
func (instance *PS_GatewayTunnel) SetPropertyLastConnectErrorString(value string) (err error) {
	return instance.SetProperty("LastConnectErrorString", (value))
}

// GetLastConnectErrorString gets the value of LastConnectErrorString for the instance
func (instance *PS_GatewayTunnel) GetPropertyLastConnectErrorString() (value string, err error) {
	retValue, err := instance.GetProperty("LastConnectErrorString")
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

// SetLocalExemptRoutes sets the value of LocalExemptRoutes for the instance
func (instance *PS_GatewayTunnel) SetPropertyLocalExemptRoutes(value string) (err error) {
	return instance.SetProperty("LocalExemptRoutes", (value))
}

// GetLocalExemptRoutes gets the value of LocalExemptRoutes for the instance
func (instance *PS_GatewayTunnel) GetPropertyLocalExemptRoutes() (value string, err error) {
	retValue, err := instance.GetProperty("LocalExemptRoutes")
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

// SetLocalIp sets the value of LocalIp for the instance
func (instance *PS_GatewayTunnel) SetPropertyLocalIp(value string) (err error) {
	return instance.SetProperty("LocalIp", (value))
}

// GetLocalIp gets the value of LocalIp for the instance
func (instance *PS_GatewayTunnel) GetPropertyLocalIp() (value string, err error) {
	retValue, err := instance.GetProperty("LocalIp")
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

// SetLocalSubnets sets the value of LocalSubnets for the instance
func (instance *PS_GatewayTunnel) SetPropertyLocalSubnets(value string) (err error) {
	return instance.SetProperty("LocalSubnets", (value))
}

// GetLocalSubnets gets the value of LocalSubnets for the instance
func (instance *PS_GatewayTunnel) GetPropertyLocalSubnets() (value string, err error) {
	retValue, err := instance.GetProperty("LocalSubnets")
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

// SetMMSALifeTimeInSeconds sets the value of MMSALifeTimeInSeconds for the instance
func (instance *PS_GatewayTunnel) SetPropertyMMSALifeTimeInSeconds(value uint32) (err error) {
	return instance.SetProperty("MMSALifeTimeInSeconds", (value))
}

// GetMMSALifeTimeInSeconds gets the value of MMSALifeTimeInSeconds for the instance
func (instance *PS_GatewayTunnel) GetPropertyMMSALifeTimeInSeconds() (value uint32, err error) {
	retValue, err := instance.GetProperty("MMSALifeTimeInSeconds")
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

// SetPfsgroup sets the value of Pfsgroup for the instance
func (instance *PS_GatewayTunnel) SetPropertyPfsgroup(value uint32) (err error) {
	return instance.SetProperty("Pfsgroup", (value))
}

// GetPfsgroup gets the value of Pfsgroup for the instance
func (instance *PS_GatewayTunnel) GetPropertyPfsgroup() (value uint32, err error) {
	retValue, err := instance.GetProperty("Pfsgroup")
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

// SetProtocolType sets the value of ProtocolType for the instance
func (instance *PS_GatewayTunnel) SetPropertyProtocolType(value uint32) (err error) {
	return instance.SetProperty("ProtocolType", (value))
}

// GetProtocolType gets the value of ProtocolType for the instance
func (instance *PS_GatewayTunnel) GetPropertyProtocolType() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProtocolType")
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

// SetRemoteExemptRoutes sets the value of RemoteExemptRoutes for the instance
func (instance *PS_GatewayTunnel) SetPropertyRemoteExemptRoutes(value string) (err error) {
	return instance.SetProperty("RemoteExemptRoutes", (value))
}

// GetRemoteExemptRoutes gets the value of RemoteExemptRoutes for the instance
func (instance *PS_GatewayTunnel) GetPropertyRemoteExemptRoutes() (value string, err error) {
	retValue, err := instance.GetProperty("RemoteExemptRoutes")
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

// SetRemoteIp sets the value of RemoteIp for the instance
func (instance *PS_GatewayTunnel) SetPropertyRemoteIp(value string) (err error) {
	return instance.SetProperty("RemoteIp", (value))
}

// GetRemoteIp gets the value of RemoteIp for the instance
func (instance *PS_GatewayTunnel) GetPropertyRemoteIp() (value string, err error) {
	retValue, err := instance.GetProperty("RemoteIp")
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

// SetRemoteSubnets sets the value of RemoteSubnets for the instance
func (instance *PS_GatewayTunnel) SetPropertyRemoteSubnets(value string) (err error) {
	return instance.SetProperty("RemoteSubnets", (value))
}

// GetRemoteSubnets gets the value of RemoteSubnets for the instance
func (instance *PS_GatewayTunnel) GetPropertyRemoteSubnets() (value string, err error) {
	retValue, err := instance.GetProperty("RemoteSubnets")
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

// SetRoutes sets the value of Routes for the instance
func (instance *PS_GatewayTunnel) SetPropertyRoutes(value string) (err error) {
	return instance.SetProperty("Routes", (value))
}

// GetRoutes gets the value of Routes for the instance
func (instance *PS_GatewayTunnel) GetPropertyRoutes() (value string, err error) {
	retValue, err := instance.GetProperty("Routes")
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

// SetRoutingDomainName sets the value of RoutingDomainName for the instance
func (instance *PS_GatewayTunnel) SetPropertyRoutingDomainName(value string) (err error) {
	return instance.SetProperty("RoutingDomainName", (value))
}

// GetRoutingDomainName gets the value of RoutingDomainName for the instance
func (instance *PS_GatewayTunnel) GetPropertyRoutingDomainName() (value string, err error) {
	retValue, err := instance.GetProperty("RoutingDomainName")
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

// SetRxBandwidthKbps sets the value of RxBandwidthKbps for the instance
func (instance *PS_GatewayTunnel) SetPropertyRxBandwidthKbps(value uint64) (err error) {
	return instance.SetProperty("RxBandwidthKbps", (value))
}

// GetRxBandwidthKbps gets the value of RxBandwidthKbps for the instance
func (instance *PS_GatewayTunnel) GetPropertyRxBandwidthKbps() (value uint64, err error) {
	retValue, err := instance.GetProperty("RxBandwidthKbps")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetSALifeKilloBytes sets the value of SALifeKilloBytes for the instance
func (instance *PS_GatewayTunnel) SetPropertySALifeKilloBytes(value uint32) (err error) {
	return instance.SetProperty("SALifeKilloBytes", (value))
}

// GetSALifeKilloBytes gets the value of SALifeKilloBytes for the instance
func (instance *PS_GatewayTunnel) GetPropertySALifeKilloBytes() (value uint32, err error) {
	retValue, err := instance.GetProperty("SALifeKilloBytes")
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

// SetSALifeTimeInSeconds sets the value of SALifeTimeInSeconds for the instance
func (instance *PS_GatewayTunnel) SetPropertySALifeTimeInSeconds(value uint32) (err error) {
	return instance.SetProperty("SALifeTimeInSeconds", (value))
}

// GetSALifeTimeInSeconds gets the value of SALifeTimeInSeconds for the instance
func (instance *PS_GatewayTunnel) GetPropertySALifeTimeInSeconds() (value uint32, err error) {
	retValue, err := instance.GetProperty("SALifeTimeInSeconds")
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

// SetStandby sets the value of Standby for the instance
func (instance *PS_GatewayTunnel) SetPropertyStandby(value bool) (err error) {
	return instance.SetProperty("Standby", (value))
}

// GetStandby gets the value of Standby for the instance
func (instance *PS_GatewayTunnel) GetPropertyStandby() (value bool, err error) {
	retValue, err := instance.GetProperty("Standby")
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

// SetTunnelId sets the value of TunnelId for the instance
func (instance *PS_GatewayTunnel) SetPropertyTunnelId(value string) (err error) {
	return instance.SetProperty("TunnelId", (value))
}

// GetTunnelId gets the value of TunnelId for the instance
func (instance *PS_GatewayTunnel) GetPropertyTunnelId() (value string, err error) {
	retValue, err := instance.GetProperty("TunnelId")
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

// SetTunnelIdentity sets the value of TunnelIdentity for the instance
func (instance *PS_GatewayTunnel) SetPropertyTunnelIdentity(value string) (err error) {
	return instance.SetProperty("TunnelIdentity", (value))
}

// GetTunnelIdentity gets the value of TunnelIdentity for the instance
func (instance *PS_GatewayTunnel) GetPropertyTunnelIdentity() (value string, err error) {
	retValue, err := instance.GetProperty("TunnelIdentity")
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

// SetTunnelName sets the value of TunnelName for the instance
func (instance *PS_GatewayTunnel) SetPropertyTunnelName(value string) (err error) {
	return instance.SetProperty("TunnelName", (value))
}

// GetTunnelName gets the value of TunnelName for the instance
func (instance *PS_GatewayTunnel) GetPropertyTunnelName() (value string, err error) {
	retValue, err := instance.GetProperty("TunnelName")
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

// SetTxBandwidthKbps sets the value of TxBandwidthKbps for the instance
func (instance *PS_GatewayTunnel) SetPropertyTxBandwidthKbps(value uint64) (err error) {
	return instance.SetProperty("TxBandwidthKbps", (value))
}

// GetTxBandwidthKbps gets the value of TxBandwidthKbps for the instance
func (instance *PS_GatewayTunnel) GetPropertyTxBandwidthKbps() (value uint64, err error) {
	retValue, err := instance.GetProperty("TxBandwidthKbps")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetUseNarrowTrafficSelectors sets the value of UseNarrowTrafficSelectors for the instance
func (instance *PS_GatewayTunnel) SetPropertyUseNarrowTrafficSelectors(value bool) (err error) {
	return instance.SetProperty("UseNarrowTrafficSelectors", (value))
}

// GetUseNarrowTrafficSelectors gets the value of UseNarrowTrafficSelectors for the instance
func (instance *PS_GatewayTunnel) GetPropertyUseNarrowTrafficSelectors() (value bool, err error) {
	retValue, err := instance.GetProperty("UseNarrowTrafficSelectors")
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

//

// <param name="RoutingDomainName" type="string "></param>
// <param name="TunnelName" type="string "></param>

// <param name="cmdletOutput" type="PS_GatewayTunnel []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_GatewayTunnel) Get( /* IN */ TunnelName string,
	/* IN */ RoutingDomainName string,
	/* OUT */ cmdletOutput []PS_GatewayTunnel) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", TunnelName, RoutingDomainName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AdminStatus" type="bool "></param>
// <param name="AutoConnectEnabled" type="bool "></param>
// <param name="DhGroup" type="uint32 "></param>
// <param name="DpdTimeout" type="uint32 "></param>
// <param name="EncryptionType" type="uint32 "></param>
// <param name="IkeEncryption" type="uint32 "></param>
// <param name="IkeIntegrity" type="uint32 "></param>
// <param name="IpsecEncryption" type="uint32 "></param>
// <param name="IpsecIntegrity" type="uint32 "></param>
// <param name="LocalExemptRoutes" type="string "></param>
// <param name="LocalIp" type="string "></param>
// <param name="LocalSubnets" type="string "></param>
// <param name="MMSALifeTimeInSeconds" type="uint32 "></param>
// <param name="PfsGroup" type="uint32 "></param>
// <param name="ProtocolType" type="uint32 "></param>
// <param name="RemoteExemptRoutes" type="string "></param>
// <param name="RemoteIp" type="string "></param>
// <param name="RemoteSubnets" type="string "></param>
// <param name="Routes" type="string "></param>
// <param name="RoutingDomainName" type="string "></param>
// <param name="RxBandwidthKbps" type="uint64 "></param>
// <param name="SALifeKilloBytes" type="uint32 "></param>
// <param name="SALifeTimeInSeconds" type="uint32 "></param>
// <param name="SharedSecret" type="string "></param>
// <param name="Standby" type="bool "></param>
// <param name="TunnelId" type="string "></param>
// <param name="TunnelIdentity" type="string "></param>
// <param name="TunnelName" type="string "></param>
// <param name="TxBandwidthKbps" type="uint64 "></param>
// <param name="UseNarrowTrafficSelectors" type="bool "></param>

// <param name="cmdletOutput" type="PS_GatewayTunnel "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_GatewayTunnel) Set( /* IN */ TunnelName string,
	/* IN */ RoutingDomainName string,
	/* IN */ LocalIp string,
	/* IN */ RemoteIp string,
	/* IN */ TunnelIdentity string,
	/* IN */ Routes string,
	/* IN */ LocalExemptRoutes string,
	/* IN */ RemoteExemptRoutes string,
	/* IN */ LocalSubnets string,
	/* IN */ RemoteSubnets string,
	/* IN */ AdminStatus bool,
	/* IN */ Standby bool,
	/* IN */ AutoConnectEnabled bool,
	/* IN */ ProtocolType uint32,
	/* IN */ EncryptionType uint32,
	/* IN */ IpsecEncryption uint32,
	/* IN */ IpsecIntegrity uint32,
	/* IN */ PfsGroup uint32,
	/* IN */ IkeEncryption uint32,
	/* IN */ IkeIntegrity uint32,
	/* IN */ DhGroup uint32,
	/* IN */ MMSALifeTimeInSeconds uint32,
	/* IN */ SALifeTimeInSeconds uint32,
	/* IN */ SALifeKilloBytes uint32,
	/* IN */ DpdTimeout uint32,
	/* IN */ SharedSecret string,
	/* IN */ UseNarrowTrafficSelectors bool,
	/* IN */ TunnelId string,
	/* IN */ TxBandwidthKbps uint64,
	/* IN */ RxBandwidthKbps uint64,
	/* OUT */ cmdletOutput PS_GatewayTunnel) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Set", TunnelName, RoutingDomainName, LocalIp, RemoteIp, TunnelIdentity, Routes, LocalExemptRoutes, RemoteExemptRoutes, LocalSubnets, RemoteSubnets, AdminStatus, Standby, AutoConnectEnabled, ProtocolType, EncryptionType, IpsecEncryption, IpsecIntegrity, PfsGroup, IkeEncryption, IkeIntegrity, DhGroup, MMSALifeTimeInSeconds, SALifeTimeInSeconds, SALifeKilloBytes, DpdTimeout, SharedSecret, UseNarrowTrafficSelectors, TunnelId, TxBandwidthKbps, RxBandwidthKbps)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="RoutingDomainName" type="string "></param>
// <param name="TunnelName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_GatewayTunnel) Remove( /* IN */ TunnelName string,
	/* IN */ RoutingDomainName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Remove", TunnelName, RoutingDomainName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="RoutingDomainName" type="string "></param>
// <param name="TunnelName" type="string "></param>

// <param name="cmdletOutput" type="PS_GatewayTunnel "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_GatewayTunnel) Connect( /* IN */ TunnelName string,
	/* IN */ RoutingDomainName string,
	/* OUT */ cmdletOutput PS_GatewayTunnel) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Connect", TunnelName, RoutingDomainName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="RoutingDomainName" type="string "></param>
// <param name="TunnelName" type="string "></param>

// <param name="cmdletOutput" type="PS_GatewayTunnel "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_GatewayTunnel) Disconnect( /* IN */ TunnelName string,
	/* IN */ RoutingDomainName string,
	/* OUT */ cmdletOutput PS_GatewayTunnel) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Disconnect", TunnelName, RoutingDomainName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
