// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.RemoteAccess
//
// ////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// PS_VpnS2SInterface struct
type PS_VpnS2SInterface struct {
	*cim.WmiInstance
}

func NewPS_VpnS2SInterfaceEx1(instance *cim.WmiInstance) (newInstance *PS_VpnS2SInterface, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_VpnS2SInterface{
		WmiInstance: tmp,
	}
	return
}

func NewPS_VpnS2SInterfaceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_VpnS2SInterface, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_VpnS2SInterface{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="AdminStatus" type="bool "></param>
// <param name="AuthenticationMethod" type="string "></param>
// <param name="AuthenticationTransformConstants" type="uint32 "></param>
// <param name="AutoConnectEnabled" type="bool "></param>
// <param name="Certificate" type="uint8 []"></param>
// <param name="CipherTransformConstants" type="uint32 "></param>
// <param name="CustomPolicy" type="bool "></param>
// <param name="Destination" type="string []"></param>
// <param name="DHGroup" type="uint32 "></param>
// <param name="EapMethod" type="string "></param>
// <param name="EnableQoS" type="uint32 "></param>
// <param name="EncryptionMethod" type="uint32 "></param>
// <param name="IdleDisconnectSeconds" type="uint32 "></param>
// <param name="InitiateConfigPayload" type="bool "></param>
// <param name="IntegrityCheckMethod" type="uint32 "></param>
// <param name="InternalIPv4" type="bool "></param>
// <param name="InternalIPv6" type="bool "></param>
// <param name="IPv4Subnet" type="string []"></param>
// <param name="IPv4TriggerFilter" type="string []"></param>
// <param name="IPv4TriggerFilterAction" type="uint32 "></param>
// <param name="IPv6Subnet" type="string []"></param>
// <param name="IPv6TriggerFilter" type="string []"></param>
// <param name="IPv6TriggerFilterAction" type="uint32 "></param>
// <param name="LocalVpnTrafficSelector" type="VpnTrafficSelector []"></param>
// <param name="MMSALifeTimeSeconds" type="uint32 "></param>
// <param name="Name" type="string "></param>
// <param name="NetworkOutageTimeSeconds" type="uint32 "></param>
// <param name="NumberOfTries" type="uint32 "></param>
// <param name="PassThru" type="bool "></param>
// <param name="Password" type="string "></param>
// <param name="Persistent" type="bool "></param>
// <param name="PfsGroup" type="uint32 "></param>
// <param name="PostConnectionIPv4Subnet" type="string []"></param>
// <param name="PostConnectionIPv6Subnet" type="string []"></param>
// <param name="PromoteAlternate" type="bool "></param>
// <param name="Protocol" type="string "></param>
// <param name="RadiusAttributeClass" type="string "></param>
// <param name="RemoteVpnTrafficSelector" type="VpnTrafficSelector []"></param>
// <param name="ResponderAuthenticationMethod" type="string "></param>
// <param name="RetryIntervalSeconds" type="uint32 "></param>
// <param name="RoutingDomain" type="string "></param>
// <param name="RxBandwidthKbps" type="uint64 "></param>
// <param name="SADataSizeForRenegotiationKilobytes" type="uint32 "></param>
// <param name="SALifeTimeSeconds" type="uint32 "></param>
// <param name="SharedSecret" type="string "></param>
// <param name="SourceIpAddress" type="string "></param>
// <param name="TxBandwidthKbps" type="uint64 "></param>
// <param name="UserName" type="string "></param>

// <param name="cmdletOutput" type="VpnS2SInterface "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnS2SInterface) AddByCustomPolicy( /* IN */ Name string,
	/* IN */ Protocol string,
	/* IN */ Destination []string,
	/* IN */ AdminStatus bool,
	/* IN */ PromoteAlternate bool,
	/* IN */ AuthenticationMethod string,
	/* IN */ PostConnectionIPv4Subnet []string,
	/* IN */ PostConnectionIPv6Subnet []string,
	/* IN */ InitiateConfigPayload bool,
	/* IN */ RadiusAttributeClass string,
	/* IN */ EnableQoS uint32,
	/* IN */ TxBandwidthKbps uint64,
	/* IN */ RxBandwidthKbps uint64,
	/* IN */ IPv4TriggerFilter []string,
	/* IN */ IPv6TriggerFilter []string,
	/* IN */ Persistent bool,
	/* IN */ IPv4TriggerFilterAction uint32,
	/* IN */ IPv6TriggerFilterAction uint32,
	/* IN */ SADataSizeForRenegotiationKilobytes uint32,
	/* IN */ IPv4Subnet []string,
	/* IN */ IPv6Subnet []string,
	/* IN */ ResponderAuthenticationMethod string,
	/* IN */ PassThru bool,
	/* IN */ RoutingDomain string,
	/* IN */ Certificate []uint8,
	/* IN */ SharedSecret string,
	/* IN */ NetworkOutageTimeSeconds uint32,
	/* IN */ NumberOfTries uint32,
	/* IN */ RetryIntervalSeconds uint32,
	/* IN */ SALifeTimeSeconds uint32,
	/* IN */ MMSALifeTimeSeconds uint32,
	/* IN */ EapMethod string,
	/* IN */ InternalIPv4 bool,
	/* IN */ InternalIPv6 bool,
	/* IN */ IdleDisconnectSeconds uint32,
	/* IN */ UserName string,
	/* IN */ Password string,
	/* IN */ CustomPolicy bool,
	/* IN */ EncryptionMethod uint32,
	/* IN */ IntegrityCheckMethod uint32,
	/* IN */ CipherTransformConstants uint32,
	/* IN */ AuthenticationTransformConstants uint32,
	/* IN */ PfsGroup uint32,
	/* IN */ DHGroup uint32,
	/* IN */ SourceIpAddress string,
	/* IN */ LocalVpnTrafficSelector []VpnTrafficSelector,
	/* IN */ RemoteVpnTrafficSelector []VpnTrafficSelector,
	/* IN */ AutoConnectEnabled bool,
	/* OUT */ cmdletOutput VpnS2SInterface) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("AddByCustomPolicy", Name, Protocol, Destination, AdminStatus, PromoteAlternate, AuthenticationMethod, PostConnectionIPv4Subnet, PostConnectionIPv6Subnet, InitiateConfigPayload, RadiusAttributeClass, EnableQoS, TxBandwidthKbps, RxBandwidthKbps, IPv4TriggerFilter, IPv6TriggerFilter, Persistent, IPv4TriggerFilterAction, IPv6TriggerFilterAction, SADataSizeForRenegotiationKilobytes, IPv4Subnet, IPv6Subnet, ResponderAuthenticationMethod, PassThru, RoutingDomain, Certificate, SharedSecret, NetworkOutageTimeSeconds, NumberOfTries, RetryIntervalSeconds, SALifeTimeSeconds, MMSALifeTimeSeconds, EapMethod, InternalIPv4, InternalIPv6, IdleDisconnectSeconds, UserName, Password, CustomPolicy, EncryptionMethod, IntegrityCheckMethod, CipherTransformConstants, AuthenticationTransformConstants, PfsGroup, DHGroup, SourceIpAddress, LocalVpnTrafficSelector, RemoteVpnTrafficSelector, AutoConnectEnabled)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AdminStatus" type="bool "></param>
// <param name="AuthenticationMethod" type="string "></param>
// <param name="AutoConnectEnabled" type="bool "></param>
// <param name="Certificate" type="uint8 []"></param>
// <param name="Destination" type="string []"></param>
// <param name="EapMethod" type="string "></param>
// <param name="EnableQoS" type="uint32 "></param>
// <param name="EncryptionType" type="string "></param>
// <param name="IdleDisconnectSeconds" type="uint32 "></param>
// <param name="InitiateConfigPayload" type="bool "></param>
// <param name="InternalIPv4" type="bool "></param>
// <param name="InternalIPv6" type="bool "></param>
// <param name="IPv4Subnet" type="string []"></param>
// <param name="IPv4TriggerFilter" type="string []"></param>
// <param name="IPv4TriggerFilterAction" type="uint32 "></param>
// <param name="IPv6Subnet" type="string []"></param>
// <param name="IPv6TriggerFilter" type="string []"></param>
// <param name="IPv6TriggerFilterAction" type="uint32 "></param>
// <param name="LocalVpnTrafficSelector" type="VpnTrafficSelector []"></param>
// <param name="MMSALifeTimeSeconds" type="uint32 "></param>
// <param name="Name" type="string "></param>
// <param name="NetworkOutageTimeSeconds" type="uint32 "></param>
// <param name="NumberOfTries" type="uint32 "></param>
// <param name="PassThru" type="bool "></param>
// <param name="Password" type="string "></param>
// <param name="Persistent" type="bool "></param>
// <param name="PostConnectionIPv4Subnet" type="string []"></param>
// <param name="PostConnectionIPv6Subnet" type="string []"></param>
// <param name="PromoteAlternate" type="bool "></param>
// <param name="Protocol" type="string "></param>
// <param name="RadiusAttributeClass" type="string "></param>
// <param name="RemoteVpnTrafficSelector" type="VpnTrafficSelector []"></param>
// <param name="ResponderAuthenticationMethod" type="string "></param>
// <param name="RetryIntervalSeconds" type="uint32 "></param>
// <param name="RoutingDomain" type="string "></param>
// <param name="RxBandwidthKbps" type="uint64 "></param>
// <param name="SADataSizeForRenegotiationKilobytes" type="uint32 "></param>
// <param name="SALifeTimeSeconds" type="uint32 "></param>
// <param name="SharedSecret" type="string "></param>
// <param name="SourceIpAddress" type="string "></param>
// <param name="TxBandwidthKbps" type="uint64 "></param>
// <param name="UserName" type="string "></param>

// <param name="cmdletOutput" type="VpnS2SInterface "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnS2SInterface) AddByEncryptionType( /* IN */ Name string,
	/* IN */ Protocol string,
	/* IN */ Destination []string,
	/* IN */ AdminStatus bool,
	/* IN */ PromoteAlternate bool,
	/* IN */ AuthenticationMethod string,
	/* IN */ PostConnectionIPv4Subnet []string,
	/* IN */ PostConnectionIPv6Subnet []string,
	/* IN */ InitiateConfigPayload bool,
	/* IN */ RadiusAttributeClass string,
	/* IN */ EnableQoS uint32,
	/* IN */ TxBandwidthKbps uint64,
	/* IN */ RxBandwidthKbps uint64,
	/* IN */ IPv4TriggerFilter []string,
	/* IN */ IPv6TriggerFilter []string,
	/* IN */ Persistent bool,
	/* IN */ IPv4TriggerFilterAction uint32,
	/* IN */ IPv6TriggerFilterAction uint32,
	/* IN */ SADataSizeForRenegotiationKilobytes uint32,
	/* IN */ IPv4Subnet []string,
	/* IN */ IPv6Subnet []string,
	/* IN */ ResponderAuthenticationMethod string,
	/* IN */ PassThru bool,
	/* IN */ RoutingDomain string,
	/* IN */ Certificate []uint8,
	/* IN */ SharedSecret string,
	/* IN */ NetworkOutageTimeSeconds uint32,
	/* IN */ NumberOfTries uint32,
	/* IN */ RetryIntervalSeconds uint32,
	/* IN */ SALifeTimeSeconds uint32,
	/* IN */ MMSALifeTimeSeconds uint32,
	/* IN */ EapMethod string,
	/* IN */ InternalIPv4 bool,
	/* IN */ InternalIPv6 bool,
	/* IN */ IdleDisconnectSeconds uint32,
	/* IN */ UserName string,
	/* IN */ Password string,
	/* IN */ EncryptionType string,
	/* IN */ SourceIpAddress string,
	/* IN */ LocalVpnTrafficSelector []VpnTrafficSelector,
	/* IN */ RemoteVpnTrafficSelector []VpnTrafficSelector,
	/* IN */ AutoConnectEnabled bool,
	/* OUT */ cmdletOutput VpnS2SInterface) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("AddByEncryptionType", Name, Protocol, Destination, AdminStatus, PromoteAlternate, AuthenticationMethod, PostConnectionIPv4Subnet, PostConnectionIPv6Subnet, InitiateConfigPayload, RadiusAttributeClass, EnableQoS, TxBandwidthKbps, RxBandwidthKbps, IPv4TriggerFilter, IPv6TriggerFilter, Persistent, IPv4TriggerFilterAction, IPv6TriggerFilterAction, SADataSizeForRenegotiationKilobytes, IPv4Subnet, IPv6Subnet, ResponderAuthenticationMethod, PassThru, RoutingDomain, Certificate, SharedSecret, NetworkOutageTimeSeconds, NumberOfTries, RetryIntervalSeconds, SALifeTimeSeconds, MMSALifeTimeSeconds, EapMethod, InternalIPv4, InternalIPv6, IdleDisconnectSeconds, UserName, Password, EncryptionType, SourceIpAddress, LocalVpnTrafficSelector, RemoteVpnTrafficSelector, AutoConnectEnabled)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AdminStatus" type="bool "></param>
// <param name="Destination" type="string []"></param>
// <param name="EnableQoS" type="uint32 "></param>
// <param name="GreKey" type="uint32 "></param>
// <param name="GreTunnel" type="bool "></param>
// <param name="InternalIPv4" type="bool "></param>
// <param name="InternalIPv6" type="bool "></param>
// <param name="IPv4Address" type="string "></param>
// <param name="IPv4Subnet" type="string []"></param>
// <param name="IPv6Address" type="string "></param>
// <param name="IPv6Subnet" type="string []"></param>
// <param name="Name" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="RoutingDomain" type="string "></param>
// <param name="RxBandwidthKbps" type="uint64 "></param>
// <param name="SourceIpAddress" type="string "></param>
// <param name="TxBandwidthKbps" type="uint64 "></param>

// <param name="cmdletOutput" type="VpnS2SInterface "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnS2SInterface) AddGreInterface( /* IN */ Name string,
	/* IN */ GreKey uint32,
	/* IN */ Destination []string,
	/* IN */ AdminStatus bool,
	/* IN */ EnableQoS uint32,
	/* IN */ TxBandwidthKbps uint64,
	/* IN */ RxBandwidthKbps uint64,
	/* IN */ IPv4Subnet []string,
	/* IN */ IPv6Subnet []string,
	/* IN */ PassThru bool,
	/* IN */ RoutingDomain string,
	/* IN */ InternalIPv4 bool,
	/* IN */ InternalIPv6 bool,
	/* IN */ IPv4Address string,
	/* IN */ IPv6Address string,
	/* IN */ SourceIpAddress string,
	/* IN */ GreTunnel bool,
	/* OUT */ cmdletOutput VpnS2SInterface) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("AddGreInterface", Name, GreKey, Destination, AdminStatus, EnableQoS, TxBandwidthKbps, RxBandwidthKbps, IPv4Subnet, IPv6Subnet, PassThru, RoutingDomain, InternalIPv4, InternalIPv6, IPv4Address, IPv6Address, SourceIpAddress, GreTunnel)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Force" type="bool "></param>
// <param name="Name" type="string []"></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="VpnS2SInterface []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnS2SInterface) Remove( /* IN */ Name []string,
	/* IN */ PassThru bool,
	/* IN */ Force bool,
	/* OUT */ cmdletOutput []VpnS2SInterface) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Remove", Name, PassThru, Force)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Name" type="string "></param>
// <param name="RoutingDomain" type="string "></param>

// <param name="cmdletOutput" type="VpnS2SInterface []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnS2SInterface) Get( /* IN */ Name string,
	/* IN */ RoutingDomain string,
	/* OUT */ cmdletOutput []VpnS2SInterface) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", Name, RoutingDomain)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Name" type="string "></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="VpnS2SInterface "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnS2SInterface) Connect( /* IN */ Name string,
	/* IN */ PassThru bool,
	/* OUT */ cmdletOutput VpnS2SInterface) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Connect", Name, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Force" type="bool "></param>
// <param name="Name" type="string "></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="VpnS2SInterface "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnS2SInterface) Disconnect( /* IN */ Name string,
	/* IN */ PassThru bool,
	/* IN */ Force bool,
	/* OUT */ cmdletOutput VpnS2SInterface) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Disconnect", Name, PassThru, Force)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AdminStatus" type="bool "></param>
// <param name="AuthenticationMethod" type="string "></param>
// <param name="AuthenticationTransformConstants" type="uint32 "></param>
// <param name="AutoConnectEnabled" type="bool "></param>
// <param name="Certificate" type="uint8 []"></param>
// <param name="CipherTransformConstants" type="uint32 "></param>
// <param name="CustomPolicy" type="bool "></param>
// <param name="Destination" type="string []"></param>
// <param name="DHGroup" type="uint32 "></param>
// <param name="EapMethod" type="string "></param>
// <param name="EnableQoS" type="uint32 "></param>
// <param name="EncryptionMethod" type="uint32 "></param>
// <param name="Force" type="bool "></param>
// <param name="IdleDisconnectSeconds" type="uint32 "></param>
// <param name="InitiateConfigPayload" type="bool "></param>
// <param name="IntegrityCheckMethod" type="uint32 "></param>
// <param name="InternalIPv4" type="bool "></param>
// <param name="InternalIPv6" type="bool "></param>
// <param name="IPv4Subnet" type="string []"></param>
// <param name="IPv4TriggerFilter" type="string []"></param>
// <param name="IPv4TriggerFilterAction" type="uint32 "></param>
// <param name="IPv6Subnet" type="string []"></param>
// <param name="IPv6TriggerFilter" type="string []"></param>
// <param name="IPv6TriggerFilterAction" type="uint32 "></param>
// <param name="LocalVpnTrafficSelector" type="VpnTrafficSelector []"></param>
// <param name="MMSALifeTimeSeconds" type="uint32 "></param>
// <param name="Name" type="string "></param>
// <param name="NetworkOutageTimeSeconds" type="uint32 "></param>
// <param name="NumberOfTries" type="uint32 "></param>
// <param name="PassThru" type="bool "></param>
// <param name="Password" type="string "></param>
// <param name="Persistent" type="bool "></param>
// <param name="PfsGroup" type="uint32 "></param>
// <param name="PostConnectionIPv4Subnet" type="string []"></param>
// <param name="PostConnectionIPv6Subnet" type="string []"></param>
// <param name="PromoteAlternate" type="bool "></param>
// <param name="RadiusAttributeClass" type="string "></param>
// <param name="RemoteVpnTrafficSelector" type="VpnTrafficSelector []"></param>
// <param name="ResponderAuthenticationMethod" type="string "></param>
// <param name="RetryIntervalSeconds" type="uint32 "></param>
// <param name="RxBandwidthKbps" type="uint64 "></param>
// <param name="SADataSizeForRenegotiationKilobytes" type="uint32 "></param>
// <param name="SALifeTimeSeconds" type="uint32 "></param>
// <param name="SharedSecret" type="string "></param>
// <param name="SourceIpAddress" type="string "></param>
// <param name="TxBandwidthKbps" type="uint64 "></param>
// <param name="UserName" type="string "></param>

// <param name="cmdletOutput" type="VpnS2SInterface "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnS2SInterface) SetByCustomPolicy( /* IN */ Destination []string,
	/* IN */ AdminStatus bool,
	/* IN */ PromoteAlternate bool,
	/* IN */ AuthenticationMethod string,
	/* IN */ EapMethod string,
	/* IN */ InternalIPv4 bool,
	/* IN */ Force bool,
	/* IN */ PostConnectionIPv4Subnet []string,
	/* IN */ IPv4TriggerFilter []string,
	/* IN */ IPv4TriggerFilterAction uint32,
	/* IN */ PostConnectionIPv6Subnet []string,
	/* IN */ IPv6TriggerFilter []string,
	/* IN */ IPv6TriggerFilterAction uint32,
	/* IN */ IPv4Subnet []string,
	/* IN */ Name string,
	/* IN */ ResponderAuthenticationMethod string,
	/* IN */ PassThru bool,
	/* IN */ Persistent bool,
	/* IN */ InitiateConfigPayload bool,
	/* IN */ RadiusAttributeClass string,
	/* IN */ EnableQoS uint32,
	/* IN */ TxBandwidthKbps uint64,
	/* IN */ RxBandwidthKbps uint64,
	/* IN */ NetworkOutageTimeSeconds uint32,
	/* IN */ NumberOfTries uint32,
	/* IN */ RetryIntervalSeconds uint32,
	/* IN */ SADataSizeForRenegotiationKilobytes uint32,
	/* IN */ SALifeTimeSeconds uint32,
	/* IN */ MMSALifeTimeSeconds uint32,
	/* IN */ IPv6Subnet []string,
	/* IN */ InternalIPv6 bool,
	/* IN */ IdleDisconnectSeconds uint32,
	/* IN */ UserName string,
	/* IN */ Password string,
	/* IN */ Certificate []uint8,
	/* IN */ SharedSecret string,
	/* IN */ EncryptionMethod uint32,
	/* IN */ IntegrityCheckMethod uint32,
	/* IN */ CipherTransformConstants uint32,
	/* IN */ DHGroup uint32,
	/* IN */ AuthenticationTransformConstants uint32,
	/* IN */ PfsGroup uint32,
	/* IN */ CustomPolicy bool,
	/* IN */ SourceIpAddress string,
	/* IN */ LocalVpnTrafficSelector []VpnTrafficSelector,
	/* IN */ RemoteVpnTrafficSelector []VpnTrafficSelector,
	/* IN */ AutoConnectEnabled bool,
	/* OUT */ cmdletOutput VpnS2SInterface) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetByCustomPolicy", Destination, AdminStatus, PromoteAlternate, AuthenticationMethod, EapMethod, InternalIPv4, Force, PostConnectionIPv4Subnet, IPv4TriggerFilter, IPv4TriggerFilterAction, PostConnectionIPv6Subnet, IPv6TriggerFilter, IPv6TriggerFilterAction, IPv4Subnet, Name, ResponderAuthenticationMethod, PassThru, Persistent, InitiateConfigPayload, RadiusAttributeClass, EnableQoS, TxBandwidthKbps, RxBandwidthKbps, NetworkOutageTimeSeconds, NumberOfTries, RetryIntervalSeconds, SADataSizeForRenegotiationKilobytes, SALifeTimeSeconds, MMSALifeTimeSeconds, IPv6Subnet, InternalIPv6, IdleDisconnectSeconds, UserName, Password, Certificate, SharedSecret, EncryptionMethod, IntegrityCheckMethod, CipherTransformConstants, DHGroup, AuthenticationTransformConstants, PfsGroup, CustomPolicy, SourceIpAddress, LocalVpnTrafficSelector, RemoteVpnTrafficSelector, AutoConnectEnabled)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AdminStatus" type="bool "></param>
// <param name="AuthenticationMethod" type="string "></param>
// <param name="AutoConnectEnabled" type="bool "></param>
// <param name="Certificate" type="uint8 []"></param>
// <param name="Destination" type="string []"></param>
// <param name="EapMethod" type="string "></param>
// <param name="EnableQoS" type="uint32 "></param>
// <param name="EncryptionType" type="string "></param>
// <param name="Force" type="bool "></param>
// <param name="IdleDisconnectSeconds" type="uint32 "></param>
// <param name="InitiateConfigPayload" type="bool "></param>
// <param name="InternalIPv4" type="bool "></param>
// <param name="InternalIPv6" type="bool "></param>
// <param name="IPv4Subnet" type="string []"></param>
// <param name="IPv4TriggerFilter" type="string []"></param>
// <param name="IPv4TriggerFilterAction" type="uint32 "></param>
// <param name="IPv6Subnet" type="string []"></param>
// <param name="IPv6TriggerFilter" type="string []"></param>
// <param name="IPv6TriggerFilterAction" type="uint32 "></param>
// <param name="LocalVpnTrafficSelector" type="VpnTrafficSelector []"></param>
// <param name="MMSALifeTimeSeconds" type="uint32 "></param>
// <param name="Name" type="string "></param>
// <param name="NetworkOutageTimeSeconds" type="uint32 "></param>
// <param name="NumberOfTries" type="uint32 "></param>
// <param name="PassThru" type="bool "></param>
// <param name="Password" type="string "></param>
// <param name="Persistent" type="bool "></param>
// <param name="PostConnectionIPv4Subnet" type="string []"></param>
// <param name="PostConnectionIPv6Subnet" type="string []"></param>
// <param name="PromoteAlternate" type="bool "></param>
// <param name="RadiusAttributeClass" type="string "></param>
// <param name="RemoteVpnTrafficSelector" type="VpnTrafficSelector []"></param>
// <param name="ResponderAuthenticationMethod" type="string "></param>
// <param name="RetryIntervalSeconds" type="uint32 "></param>
// <param name="RxBandwidthKbps" type="uint64 "></param>
// <param name="SADataSizeForRenegotiationKilobytes" type="uint32 "></param>
// <param name="SALifeTimeSeconds" type="uint32 "></param>
// <param name="SharedSecret" type="string "></param>
// <param name="SourceIpAddress" type="string "></param>
// <param name="TxBandwidthKbps" type="uint64 "></param>
// <param name="UserName" type="string "></param>

// <param name="cmdletOutput" type="VpnS2SInterface "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnS2SInterface) SetByEncryptionType( /* IN */ Destination []string,
	/* IN */ AdminStatus bool,
	/* IN */ PromoteAlternate bool,
	/* IN */ AuthenticationMethod string,
	/* IN */ EapMethod string,
	/* IN */ InternalIPv4 bool,
	/* IN */ Force bool,
	/* IN */ PostConnectionIPv4Subnet []string,
	/* IN */ IPv4TriggerFilter []string,
	/* IN */ IPv4TriggerFilterAction uint32,
	/* IN */ PostConnectionIPv6Subnet []string,
	/* IN */ IPv6TriggerFilter []string,
	/* IN */ IPv6TriggerFilterAction uint32,
	/* IN */ IPv4Subnet []string,
	/* IN */ Name string,
	/* IN */ ResponderAuthenticationMethod string,
	/* IN */ PassThru bool,
	/* IN */ Persistent bool,
	/* IN */ InitiateConfigPayload bool,
	/* IN */ RadiusAttributeClass string,
	/* IN */ EnableQoS uint32,
	/* IN */ TxBandwidthKbps uint64,
	/* IN */ RxBandwidthKbps uint64,
	/* IN */ NetworkOutageTimeSeconds uint32,
	/* IN */ NumberOfTries uint32,
	/* IN */ RetryIntervalSeconds uint32,
	/* IN */ SADataSizeForRenegotiationKilobytes uint32,
	/* IN */ SALifeTimeSeconds uint32,
	/* IN */ MMSALifeTimeSeconds uint32,
	/* IN */ IPv6Subnet []string,
	/* IN */ InternalIPv6 bool,
	/* IN */ IdleDisconnectSeconds uint32,
	/* IN */ UserName string,
	/* IN */ Password string,
	/* IN */ Certificate []uint8,
	/* IN */ SharedSecret string,
	/* IN */ EncryptionType string,
	/* IN */ SourceIpAddress string,
	/* IN */ LocalVpnTrafficSelector []VpnTrafficSelector,
	/* IN */ RemoteVpnTrafficSelector []VpnTrafficSelector,
	/* IN */ AutoConnectEnabled bool,
	/* OUT */ cmdletOutput VpnS2SInterface) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetByEncryptionType", Destination, AdminStatus, PromoteAlternate, AuthenticationMethod, EapMethod, InternalIPv4, Force, PostConnectionIPv4Subnet, IPv4TriggerFilter, IPv4TriggerFilterAction, PostConnectionIPv6Subnet, IPv6TriggerFilter, IPv6TriggerFilterAction, IPv4Subnet, Name, ResponderAuthenticationMethod, PassThru, Persistent, InitiateConfigPayload, RadiusAttributeClass, EnableQoS, TxBandwidthKbps, RxBandwidthKbps, NetworkOutageTimeSeconds, NumberOfTries, RetryIntervalSeconds, SADataSizeForRenegotiationKilobytes, SALifeTimeSeconds, MMSALifeTimeSeconds, IPv6Subnet, InternalIPv6, IdleDisconnectSeconds, UserName, Password, Certificate, SharedSecret, EncryptionType, SourceIpAddress, LocalVpnTrafficSelector, RemoteVpnTrafficSelector, AutoConnectEnabled)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AdminStatus" type="bool "></param>
// <param name="Destination" type="string []"></param>
// <param name="EnableQoS" type="uint32 "></param>
// <param name="Force" type="bool "></param>
// <param name="GreKey" type="uint32 "></param>
// <param name="GreTunnel" type="bool "></param>
// <param name="InternalIPv4" type="bool "></param>
// <param name="InternalIPv6" type="bool "></param>
// <param name="IPv4Address" type="string "></param>
// <param name="IPv4Subnet" type="string []"></param>
// <param name="IPv6Address" type="string "></param>
// <param name="IPv6Subnet" type="string []"></param>
// <param name="Name" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="RxBandwidthKbps" type="uint64 "></param>
// <param name="SourceIpAddress" type="string "></param>
// <param name="TxBandwidthKbps" type="uint64 "></param>

// <param name="cmdletOutput" type="VpnS2SInterface "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnS2SInterface) SetGreInterface( /* IN */ Name string,
	/* IN */ GreKey uint32,
	/* IN */ Destination []string,
	/* IN */ AdminStatus bool,
	/* IN */ EnableQoS uint32,
	/* IN */ TxBandwidthKbps uint64,
	/* IN */ RxBandwidthKbps uint64,
	/* IN */ IPv4Subnet []string,
	/* IN */ IPv6Subnet []string,
	/* IN */ PassThru bool,
	/* IN */ InternalIPv4 bool,
	/* IN */ InternalIPv6 bool,
	/* IN */ IPv4Address string,
	/* IN */ IPv6Address string,
	/* IN */ SourceIpAddress string,
	/* IN */ Force bool,
	/* IN */ GreTunnel bool,
	/* OUT */ cmdletOutput VpnS2SInterface) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetGreInterface", Name, GreKey, Destination, AdminStatus, EnableQoS, TxBandwidthKbps, RxBandwidthKbps, IPv4Subnet, IPv6Subnet, PassThru, InternalIPv4, InternalIPv6, IPv4Address, IPv6Address, SourceIpAddress, Force, GreTunnel)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
