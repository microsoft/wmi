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

// PS_RemoteAccessRoutingDomain struct
type PS_RemoteAccessRoutingDomain struct {
	*cim.WmiInstance
}

func NewPS_RemoteAccessRoutingDomainEx1(instance *cim.WmiInstance) (newInstance *PS_RemoteAccessRoutingDomain, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_RemoteAccessRoutingDomain{
		WmiInstance: tmp,
	}
	return
}

func NewPS_RemoteAccessRoutingDomainEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_RemoteAccessRoutingDomain, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_RemoteAccessRoutingDomain{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="Name" type="string "></param>

// <param name="cmdletOutput" type="VpnRoutingDomainConfig []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessRoutingDomain) Get( /* IN */ Name string,
	/* OUT */ cmdletOutput []VpnRoutingDomainConfig) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", Name)
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
// <param name="Type" type="uint32 "></param>

// <param name="cmdletOutput" type="RoutingDomainConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessRoutingDomain) Enable( /* IN */ Name string,
	/* IN */ PassThru bool,
	/* IN */ Type uint32,
	/* OUT */ cmdletOutput RoutingDomainConfiguration) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Enable", Name, PassThru, Type)
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
// <param name="Type" type="uint32 "></param>

// <param name="cmdletOutput" type="RoutingDomainConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessRoutingDomain) Disable( /* IN */ Name string,
	/* IN */ Force bool,
	/* IN */ PassThru bool,
	/* IN */ Type uint32,
	/* OUT */ cmdletOutput RoutingDomainConfiguration) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Disable", Name, Force, PassThru, Type)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AuthenticationTransformConstant" type="uint32 "></param>
// <param name="CipherTransformConstant" type="uint32 "></param>
// <param name="CustomPolicy" type="bool "></param>
// <param name="DHGroup" type="uint32 "></param>
// <param name="DnsIPAddress" type="string []"></param>
// <param name="EnableQoS" type="uint32 "></param>
// <param name="EncryptionMethod" type="uint32 "></param>
// <param name="Force" type="bool "></param>
// <param name="IdleDisconnectSec" type="uint32 "></param>
// <param name="IntegrityCheckMethod" type="uint32 "></param>
// <param name="InterimAccountingPeriodSec" type="uint32 "></param>
// <param name="IPAddressRange" type="string []"></param>
// <param name="IPv6Prefix" type="string "></param>
// <param name="MaximumVpnConnections" type="uint32 "></param>
// <param name="MMSaLifeTimeSec" type="uint32 "></param>
// <param name="Name" type="string "></param>
// <param name="NetBiosIPAddress" type="string []"></param>
// <param name="PassThru" type="bool "></param>
// <param name="PfsGroup" type="uint32 "></param>
// <param name="RxBandwidthKbps" type="uint64 "></param>
// <param name="SaLifeTimeSec" type="uint32 "></param>
// <param name="SaRenegotiationDataSizeKB" type="uint32 "></param>
// <param name="TenantName" type="string []"></param>
// <param name="TxBandwidthKbps" type="uint64 "></param>

// <param name="cmdletOutput" type="VpnRoutingDomainConfig "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessRoutingDomain) SetByCustomPolicy( /* IN */ Name string,
	/* IN */ IdleDisconnectSec uint32,
	/* IN */ InterimAccountingPeriodSec uint32,
	/* IN */ IPAddressRange []string,
	/* IN */ IPv6Prefix string,
	/* IN */ SaLifeTimeSec uint32,
	/* IN */ MMSaLifeTimeSec uint32,
	/* IN */ NetBiosIPAddress []string,
	/* IN */ MaximumVpnConnections uint32,
	/* IN */ TenantName []string,
	/* IN */ PassThru bool,
	/* IN */ Force bool,
	/* IN */ EnableQoS uint32,
	/* IN */ TxBandwidthKbps uint64,
	/* IN */ RxBandwidthKbps uint64,
	/* IN */ DnsIPAddress []string,
	/* IN */ CustomPolicy bool,
	/* IN */ AuthenticationTransformConstant uint32,
	/* IN */ CipherTransformConstant uint32,
	/* IN */ EncryptionMethod uint32,
	/* IN */ IntegrityCheckMethod uint32,
	/* IN */ PfsGroup uint32,
	/* IN */ SaRenegotiationDataSizeKB uint32,
	/* IN */ DHGroup uint32,
	/* OUT */ cmdletOutput VpnRoutingDomainConfig) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetByCustomPolicy", Name, IdleDisconnectSec, InterimAccountingPeriodSec, IPAddressRange, IPv6Prefix, SaLifeTimeSec, MMSaLifeTimeSec, NetBiosIPAddress, MaximumVpnConnections, TenantName, PassThru, Force, EnableQoS, TxBandwidthKbps, RxBandwidthKbps, DnsIPAddress, CustomPolicy, AuthenticationTransformConstant, CipherTransformConstant, EncryptionMethod, IntegrityCheckMethod, PfsGroup, SaRenegotiationDataSizeKB, DHGroup)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="DnsIPAddress" type="string []"></param>
// <param name="EnableQoS" type="uint32 "></param>
// <param name="EncryptionType" type="string "></param>
// <param name="Force" type="bool "></param>
// <param name="InterimAccountingPeriodSec" type="uint32 "></param>
// <param name="IPAddressRange" type="string []"></param>
// <param name="IPv6Prefix" type="string "></param>
// <param name="MaximumVpnConnections" type="uint32 "></param>
// <param name="Name" type="string "></param>
// <param name="NetBiosIPAddress" type="string []"></param>
// <param name="PassThru" type="bool "></param>
// <param name="RxBandwidthKbps" type="uint64 "></param>
// <param name="TenantName" type="string []"></param>
// <param name="TxBandwidthKbps" type="uint64 "></param>

// <param name="cmdletOutput" type="VpnRoutingDomainConfig "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessRoutingDomain) SetByEncryptionType( /* IN */ Name string,
	/* IN */ InterimAccountingPeriodSec uint32,
	/* IN */ IPAddressRange []string,
	/* IN */ IPv6Prefix string,
	/* IN */ NetBiosIPAddress []string,
	/* IN */ MaximumVpnConnections uint32,
	/* IN */ TenantName []string,
	/* IN */ PassThru bool,
	/* IN */ Force bool,
	/* IN */ EnableQoS uint32,
	/* IN */ TxBandwidthKbps uint64,
	/* IN */ RxBandwidthKbps uint64,
	/* IN */ DnsIPAddress []string,
	/* IN */ EncryptionType string,
	/* OUT */ cmdletOutput VpnRoutingDomainConfig) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetByEncryptionType", Name, InterimAccountingPeriodSec, IPAddressRange, IPv6Prefix, NetBiosIPAddress, MaximumVpnConnections, TenantName, PassThru, Force, EnableQoS, TxBandwidthKbps, RxBandwidthKbps, DnsIPAddress, EncryptionType)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
