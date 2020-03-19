// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess.Client
//////////////////////////////////////////////
package client

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// PS_VpnConnection struct
type PS_VpnConnection struct {
	*cim.WmiInstance

	//
	AllUserConnection bool

	//
	ApplicationID []string

	//
	AuthenticationMethod []string

	//
	ConnectionStatus string

	//
	CustomConfiguration string

	//
	DnsConfig []PS_VpnConnectionTriggerDnsConfiguration

	//
	DnsSuffix string

	//
	DnsSuffixSearchList []string

	//
	EapConfigXmlStream string

	//
	EncryptionLevel string

	//
	Guid string

	//
	IdleDisconnectSeconds uint32

	//
	IsAutoTriggerEnabled bool

	//
	L2tpIPsecAuth string

	//
	L2tpPsk string

	//
	MachineCertificateEKUFilter []string

	//
	MachineCertificateIssuerFilter string

	//
	Name string

	//
	NapState string

	//
	PlugInApplicationID string

	//
	ProfileType string

	//
	ProvisioningAuthority string

	//
	Proxy PS_VpnConnectionProxy

	//
	RememberCredential bool

	//
	Routes []PS_VpnConnectionRoute

	//
	ServerAddress string

	//
	ServerList []PS_VpnServerAddress

	//
	SplitTunneling bool

	//
	TrustedNetwork []string

	//
	TunnelType string

	//
	UseWinlogonCredential bool

	//
	VpnConfigurationXml string
}

func NewPS_VpnConnectionEx1(instance *cim.WmiInstance) (newInstance *PS_VpnConnection, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_VpnConnection{
		WmiInstance: tmp,
	}
	return
}

func NewPS_VpnConnectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_VpnConnection, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_VpnConnection{
		WmiInstance: tmp,
	}
	return
}

// SetAllUserConnection sets the value of AllUserConnection for the instance
func (instance *PS_VpnConnection) SetPropertyAllUserConnection(value bool) (err error) {
	return instance.SetProperty("AllUserConnection", value)
}

// GetAllUserConnection gets the value of AllUserConnection for the instance
func (instance *PS_VpnConnection) GetPropertyAllUserConnection() (value bool, err error) {
	retValue, err := instance.GetProperty("AllUserConnection")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetApplicationID sets the value of ApplicationID for the instance
func (instance *PS_VpnConnection) SetPropertyApplicationID(value []string) (err error) {
	return instance.SetProperty("ApplicationID", value)
}

// GetApplicationID gets the value of ApplicationID for the instance
func (instance *PS_VpnConnection) GetPropertyApplicationID() (value []string, err error) {
	retValue, err := instance.GetProperty("ApplicationID")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAuthenticationMethod sets the value of AuthenticationMethod for the instance
func (instance *PS_VpnConnection) SetPropertyAuthenticationMethod(value []string) (err error) {
	return instance.SetProperty("AuthenticationMethod", value)
}

// GetAuthenticationMethod gets the value of AuthenticationMethod for the instance
func (instance *PS_VpnConnection) GetPropertyAuthenticationMethod() (value []string, err error) {
	retValue, err := instance.GetProperty("AuthenticationMethod")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConnectionStatus sets the value of ConnectionStatus for the instance
func (instance *PS_VpnConnection) SetPropertyConnectionStatus(value string) (err error) {
	return instance.SetProperty("ConnectionStatus", value)
}

// GetConnectionStatus gets the value of ConnectionStatus for the instance
func (instance *PS_VpnConnection) GetPropertyConnectionStatus() (value string, err error) {
	retValue, err := instance.GetProperty("ConnectionStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCustomConfiguration sets the value of CustomConfiguration for the instance
func (instance *PS_VpnConnection) SetPropertyCustomConfiguration(value string) (err error) {
	return instance.SetProperty("CustomConfiguration", value)
}

// GetCustomConfiguration gets the value of CustomConfiguration for the instance
func (instance *PS_VpnConnection) GetPropertyCustomConfiguration() (value string, err error) {
	retValue, err := instance.GetProperty("CustomConfiguration")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDnsConfig sets the value of DnsConfig for the instance
func (instance *PS_VpnConnection) SetPropertyDnsConfig(value []PS_VpnConnectionTriggerDnsConfiguration) (err error) {
	return instance.SetProperty("DnsConfig", value)
}

// GetDnsConfig gets the value of DnsConfig for the instance
func (instance *PS_VpnConnection) GetPropertyDnsConfig() (value []PS_VpnConnectionTriggerDnsConfiguration, err error) {
	retValue, err := instance.GetProperty("DnsConfig")
	if err != nil {
		return
	}
	value, ok := retValue.([]PS_VpnConnectionTriggerDnsConfiguration)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDnsSuffix sets the value of DnsSuffix for the instance
func (instance *PS_VpnConnection) SetPropertyDnsSuffix(value string) (err error) {
	return instance.SetProperty("DnsSuffix", value)
}

// GetDnsSuffix gets the value of DnsSuffix for the instance
func (instance *PS_VpnConnection) GetPropertyDnsSuffix() (value string, err error) {
	retValue, err := instance.GetProperty("DnsSuffix")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDnsSuffixSearchList sets the value of DnsSuffixSearchList for the instance
func (instance *PS_VpnConnection) SetPropertyDnsSuffixSearchList(value []string) (err error) {
	return instance.SetProperty("DnsSuffixSearchList", value)
}

// GetDnsSuffixSearchList gets the value of DnsSuffixSearchList for the instance
func (instance *PS_VpnConnection) GetPropertyDnsSuffixSearchList() (value []string, err error) {
	retValue, err := instance.GetProperty("DnsSuffixSearchList")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEapConfigXmlStream sets the value of EapConfigXmlStream for the instance
func (instance *PS_VpnConnection) SetPropertyEapConfigXmlStream(value string) (err error) {
	return instance.SetProperty("EapConfigXmlStream", value)
}

// GetEapConfigXmlStream gets the value of EapConfigXmlStream for the instance
func (instance *PS_VpnConnection) GetPropertyEapConfigXmlStream() (value string, err error) {
	retValue, err := instance.GetProperty("EapConfigXmlStream")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEncryptionLevel sets the value of EncryptionLevel for the instance
func (instance *PS_VpnConnection) SetPropertyEncryptionLevel(value string) (err error) {
	return instance.SetProperty("EncryptionLevel", value)
}

// GetEncryptionLevel gets the value of EncryptionLevel for the instance
func (instance *PS_VpnConnection) GetPropertyEncryptionLevel() (value string, err error) {
	retValue, err := instance.GetProperty("EncryptionLevel")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGuid sets the value of Guid for the instance
func (instance *PS_VpnConnection) SetPropertyGuid(value string) (err error) {
	return instance.SetProperty("Guid", value)
}

// GetGuid gets the value of Guid for the instance
func (instance *PS_VpnConnection) GetPropertyGuid() (value string, err error) {
	retValue, err := instance.GetProperty("Guid")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIdleDisconnectSeconds sets the value of IdleDisconnectSeconds for the instance
func (instance *PS_VpnConnection) SetPropertyIdleDisconnectSeconds(value uint32) (err error) {
	return instance.SetProperty("IdleDisconnectSeconds", value)
}

// GetIdleDisconnectSeconds gets the value of IdleDisconnectSeconds for the instance
func (instance *PS_VpnConnection) GetPropertyIdleDisconnectSeconds() (value uint32, err error) {
	retValue, err := instance.GetProperty("IdleDisconnectSeconds")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsAutoTriggerEnabled sets the value of IsAutoTriggerEnabled for the instance
func (instance *PS_VpnConnection) SetPropertyIsAutoTriggerEnabled(value bool) (err error) {
	return instance.SetProperty("IsAutoTriggerEnabled", value)
}

// GetIsAutoTriggerEnabled gets the value of IsAutoTriggerEnabled for the instance
func (instance *PS_VpnConnection) GetPropertyIsAutoTriggerEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("IsAutoTriggerEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetL2tpIPsecAuth sets the value of L2tpIPsecAuth for the instance
func (instance *PS_VpnConnection) SetPropertyL2tpIPsecAuth(value string) (err error) {
	return instance.SetProperty("L2tpIPsecAuth", value)
}

// GetL2tpIPsecAuth gets the value of L2tpIPsecAuth for the instance
func (instance *PS_VpnConnection) GetPropertyL2tpIPsecAuth() (value string, err error) {
	retValue, err := instance.GetProperty("L2tpIPsecAuth")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetL2tpPsk sets the value of L2tpPsk for the instance
func (instance *PS_VpnConnection) SetPropertyL2tpPsk(value string) (err error) {
	return instance.SetProperty("L2tpPsk", value)
}

// GetL2tpPsk gets the value of L2tpPsk for the instance
func (instance *PS_VpnConnection) GetPropertyL2tpPsk() (value string, err error) {
	retValue, err := instance.GetProperty("L2tpPsk")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMachineCertificateEKUFilter sets the value of MachineCertificateEKUFilter for the instance
func (instance *PS_VpnConnection) SetPropertyMachineCertificateEKUFilter(value []string) (err error) {
	return instance.SetProperty("MachineCertificateEKUFilter", value)
}

// GetMachineCertificateEKUFilter gets the value of MachineCertificateEKUFilter for the instance
func (instance *PS_VpnConnection) GetPropertyMachineCertificateEKUFilter() (value []string, err error) {
	retValue, err := instance.GetProperty("MachineCertificateEKUFilter")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMachineCertificateIssuerFilter sets the value of MachineCertificateIssuerFilter for the instance
func (instance *PS_VpnConnection) SetPropertyMachineCertificateIssuerFilter(value string) (err error) {
	return instance.SetProperty("MachineCertificateIssuerFilter", value)
}

// GetMachineCertificateIssuerFilter gets the value of MachineCertificateIssuerFilter for the instance
func (instance *PS_VpnConnection) GetPropertyMachineCertificateIssuerFilter() (value string, err error) {
	retValue, err := instance.GetProperty("MachineCertificateIssuerFilter")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *PS_VpnConnection) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *PS_VpnConnection) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNapState sets the value of NapState for the instance
func (instance *PS_VpnConnection) SetPropertyNapState(value string) (err error) {
	return instance.SetProperty("NapState", value)
}

// GetNapState gets the value of NapState for the instance
func (instance *PS_VpnConnection) GetPropertyNapState() (value string, err error) {
	retValue, err := instance.GetProperty("NapState")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPlugInApplicationID sets the value of PlugInApplicationID for the instance
func (instance *PS_VpnConnection) SetPropertyPlugInApplicationID(value string) (err error) {
	return instance.SetProperty("PlugInApplicationID", value)
}

// GetPlugInApplicationID gets the value of PlugInApplicationID for the instance
func (instance *PS_VpnConnection) GetPropertyPlugInApplicationID() (value string, err error) {
	retValue, err := instance.GetProperty("PlugInApplicationID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProfileType sets the value of ProfileType for the instance
func (instance *PS_VpnConnection) SetPropertyProfileType(value string) (err error) {
	return instance.SetProperty("ProfileType", value)
}

// GetProfileType gets the value of ProfileType for the instance
func (instance *PS_VpnConnection) GetPropertyProfileType() (value string, err error) {
	retValue, err := instance.GetProperty("ProfileType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProvisioningAuthority sets the value of ProvisioningAuthority for the instance
func (instance *PS_VpnConnection) SetPropertyProvisioningAuthority(value string) (err error) {
	return instance.SetProperty("ProvisioningAuthority", value)
}

// GetProvisioningAuthority gets the value of ProvisioningAuthority for the instance
func (instance *PS_VpnConnection) GetPropertyProvisioningAuthority() (value string, err error) {
	retValue, err := instance.GetProperty("ProvisioningAuthority")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProxy sets the value of Proxy for the instance
func (instance *PS_VpnConnection) SetPropertyProxy(value PS_VpnConnectionProxy) (err error) {
	return instance.SetProperty("Proxy", value)
}

// GetProxy gets the value of Proxy for the instance
func (instance *PS_VpnConnection) GetPropertyProxy() (value PS_VpnConnectionProxy, err error) {
	retValue, err := instance.GetProperty("Proxy")
	if err != nil {
		return
	}
	value, ok := retValue.(PS_VpnConnectionProxy)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRememberCredential sets the value of RememberCredential for the instance
func (instance *PS_VpnConnection) SetPropertyRememberCredential(value bool) (err error) {
	return instance.SetProperty("RememberCredential", value)
}

// GetRememberCredential gets the value of RememberCredential for the instance
func (instance *PS_VpnConnection) GetPropertyRememberCredential() (value bool, err error) {
	retValue, err := instance.GetProperty("RememberCredential")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRoutes sets the value of Routes for the instance
func (instance *PS_VpnConnection) SetPropertyRoutes(value []PS_VpnConnectionRoute) (err error) {
	return instance.SetProperty("Routes", value)
}

// GetRoutes gets the value of Routes for the instance
func (instance *PS_VpnConnection) GetPropertyRoutes() (value []PS_VpnConnectionRoute, err error) {
	retValue, err := instance.GetProperty("Routes")
	if err != nil {
		return
	}
	value, ok := retValue.([]PS_VpnConnectionRoute)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetServerAddress sets the value of ServerAddress for the instance
func (instance *PS_VpnConnection) SetPropertyServerAddress(value string) (err error) {
	return instance.SetProperty("ServerAddress", value)
}

// GetServerAddress gets the value of ServerAddress for the instance
func (instance *PS_VpnConnection) GetPropertyServerAddress() (value string, err error) {
	retValue, err := instance.GetProperty("ServerAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetServerList sets the value of ServerList for the instance
func (instance *PS_VpnConnection) SetPropertyServerList(value []PS_VpnServerAddress) (err error) {
	return instance.SetProperty("ServerList", value)
}

// GetServerList gets the value of ServerList for the instance
func (instance *PS_VpnConnection) GetPropertyServerList() (value []PS_VpnServerAddress, err error) {
	retValue, err := instance.GetProperty("ServerList")
	if err != nil {
		return
	}
	value, ok := retValue.([]PS_VpnServerAddress)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSplitTunneling sets the value of SplitTunneling for the instance
func (instance *PS_VpnConnection) SetPropertySplitTunneling(value bool) (err error) {
	return instance.SetProperty("SplitTunneling", value)
}

// GetSplitTunneling gets the value of SplitTunneling for the instance
func (instance *PS_VpnConnection) GetPropertySplitTunneling() (value bool, err error) {
	retValue, err := instance.GetProperty("SplitTunneling")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTrustedNetwork sets the value of TrustedNetwork for the instance
func (instance *PS_VpnConnection) SetPropertyTrustedNetwork(value []string) (err error) {
	return instance.SetProperty("TrustedNetwork", value)
}

// GetTrustedNetwork gets the value of TrustedNetwork for the instance
func (instance *PS_VpnConnection) GetPropertyTrustedNetwork() (value []string, err error) {
	retValue, err := instance.GetProperty("TrustedNetwork")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTunnelType sets the value of TunnelType for the instance
func (instance *PS_VpnConnection) SetPropertyTunnelType(value string) (err error) {
	return instance.SetProperty("TunnelType", value)
}

// GetTunnelType gets the value of TunnelType for the instance
func (instance *PS_VpnConnection) GetPropertyTunnelType() (value string, err error) {
	retValue, err := instance.GetProperty("TunnelType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUseWinlogonCredential sets the value of UseWinlogonCredential for the instance
func (instance *PS_VpnConnection) SetPropertyUseWinlogonCredential(value bool) (err error) {
	return instance.SetProperty("UseWinlogonCredential", value)
}

// GetUseWinlogonCredential gets the value of UseWinlogonCredential for the instance
func (instance *PS_VpnConnection) GetPropertyUseWinlogonCredential() (value bool, err error) {
	retValue, err := instance.GetProperty("UseWinlogonCredential")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVpnConfigurationXml sets the value of VpnConfigurationXml for the instance
func (instance *PS_VpnConnection) SetPropertyVpnConfigurationXml(value string) (err error) {
	return instance.SetProperty("VpnConfigurationXml", value)
}

// GetVpnConfigurationXml gets the value of VpnConfigurationXml for the instance
func (instance *PS_VpnConnection) GetPropertyVpnConfigurationXml() (value string, err error) {
	retValue, err := instance.GetProperty("VpnConfigurationXml")
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

// <param name="AllUserConnection" type="bool "></param>
// <param name="Force" type="bool "></param>
// <param name="Name" type="string []"></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="VpnCommonConfig []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnConnection) Remove( /* IN */ Name []string,
	/* IN */ Force bool,
	/* IN */ PassThru bool,
	/* IN */ AllUserConnection bool,
	/* OUT */ cmdletOutput []VpnCommonConfig) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Remove", Name, Force, PassThru, AllUserConnection)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AllUserConnection" type="bool "></param>
// <param name="Name" type="string []"></param>

// <param name="cmdletOutput" type="VpnCommonConfig []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnConnection) Get( /* IN */ Name []string,
	/* IN */ AllUserConnection bool,
	/* OUT */ cmdletOutput []VpnCommonConfig) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", Name, AllUserConnection)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AllUserConnection" type="bool "></param>
// <param name="AuthenticationMethod" type="string []"></param>
// <param name="DnsSuffix" type="string "></param>
// <param name="EapConfigXmlStream" type="string "></param>
// <param name="EncryptionLevel" type="string "></param>
// <param name="Force" type="bool "></param>
// <param name="IdleDisconnectSeconds" type="uint32 "></param>
// <param name="L2tpPsk" type="string "></param>
// <param name="MachineCertificateEKUFilter" type="string []"></param>
// <param name="MachineCertificateIssuerFilter" type="uint8 []"></param>
// <param name="Name" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="RememberCredential" type="bool "></param>
// <param name="ServerAddress" type="string "></param>
// <param name="ServerList" type="VpnServerAddress []"></param>
// <param name="SplitTunneling" type="bool "></param>
// <param name="TunnelType" type="string "></param>
// <param name="UseWinlogonCredential" type="bool "></param>

// <param name="cmdletOutput" type="VpnConnection "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnConnection) Set( /* IN */ ServerAddress string,
	/* IN */ AllUserConnection bool,
	/* IN */ SplitTunneling bool,
	/* IN */ RememberCredential bool,
	/* IN */ TunnelType string,
	/* IN */ PassThru bool,
	/* IN */ Force bool,
	/* IN */ L2tpPsk string,
	/* IN */ AuthenticationMethod []string,
	/* IN */ EapConfigXmlStream string,
	/* IN */ Name string,
	/* IN */ UseWinlogonCredential bool,
	/* IN */ EncryptionLevel string,
	/* IN */ MachineCertificateEKUFilter []string,
	/* IN */ MachineCertificateIssuerFilter []uint8,
	/* IN */ ServerList []VpnServerAddress,
	/* IN */ IdleDisconnectSeconds uint32,
	/* IN */ DnsSuffix string,
	/* OUT */ cmdletOutput VpnConnection) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Set", ServerAddress, AllUserConnection, SplitTunneling, RememberCredential, TunnelType, PassThru, Force, L2tpPsk, AuthenticationMethod, EapConfigXmlStream, Name, UseWinlogonCredential, EncryptionLevel, MachineCertificateEKUFilter, MachineCertificateIssuerFilter, ServerList, IdleDisconnectSeconds, DnsSuffix)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="CustomConfiguration" type="string "></param>
// <param name="DnsSuffix" type="string "></param>
// <param name="Force" type="bool "></param>
// <param name="IdleDisconnectSeconds" type="uint32 "></param>
// <param name="Name" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="PlugInApplicationID" type="string "></param>
// <param name="RememberCredential" type="bool "></param>
// <param name="ServerAddress" type="string "></param>
// <param name="ServerList" type="VpnServerAddress []"></param>
// <param name="SplitTunneling" type="bool "></param>
// <param name="ThirdPartyVpn" type="bool "></param>

// <param name="cmdletOutput" type="ThirdPartyVpnConnection "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnConnection) SetByThirdParty( /* IN */ Name string,
	/* IN */ ServerAddress string,
	/* IN */ RememberCredential bool,
	/* IN */ SplitTunneling bool,
	/* IN */ PassThru bool,
	/* IN */ Force bool,
	/* IN */ ServerList []VpnServerAddress,
	/* IN */ DnsSuffix string,
	/* IN */ IdleDisconnectSeconds uint32,
	/* IN */ PlugInApplicationID string,
	/* IN */ CustomConfiguration string,
	/* IN */ ThirdPartyVpn bool,
	/* OUT */ cmdletOutput ThirdPartyVpnConnection) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetByThirdParty", Name, ServerAddress, RememberCredential, SplitTunneling, PassThru, Force, ServerList, DnsSuffix, IdleDisconnectSeconds, PlugInApplicationID, CustomConfiguration, ThirdPartyVpn)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AllUserConnection" type="bool "></param>
// <param name="AuthenticationMethod" type="string []"></param>
// <param name="DnsSuffix" type="string "></param>
// <param name="EapConfigXmlStream" type="string "></param>
// <param name="EncryptionLevel" type="string "></param>
// <param name="Force" type="bool "></param>
// <param name="IdleDisconnectSeconds" type="uint32 "></param>
// <param name="L2tpPsk" type="string "></param>
// <param name="MachineCertificateEKUFilter" type="string []"></param>
// <param name="MachineCertificateIssuerFilter" type="uint8 []"></param>
// <param name="Name" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="RememberCredential" type="bool "></param>
// <param name="ServerAddress" type="string "></param>
// <param name="ServerList" type="VpnServerAddress []"></param>
// <param name="SplitTunneling" type="bool "></param>
// <param name="TunnelType" type="string "></param>
// <param name="UseWinlogonCredential" type="bool "></param>

// <param name="cmdletOutput" type="VpnConnection "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnConnection) Add( /* IN */ Name string,
	/* IN */ ServerAddress string,
	/* IN */ TunnelType string,
	/* IN */ AllUserConnection bool,
	/* IN */ RememberCredential bool,
	/* IN */ IdleDisconnectSeconds uint32,
	/* IN */ DnsSuffix string,
	/* IN */ SplitTunneling bool,
	/* IN */ Force bool,
	/* IN */ PassThru bool,
	/* IN */ L2tpPsk string,
	/* IN */ UseWinlogonCredential bool,
	/* IN */ EapConfigXmlStream string,
	/* IN */ AuthenticationMethod []string,
	/* IN */ EncryptionLevel string,
	/* IN */ MachineCertificateIssuerFilter []uint8,
	/* IN */ MachineCertificateEKUFilter []string,
	/* IN */ ServerList []VpnServerAddress,
	/* OUT */ cmdletOutput VpnConnection) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Add", Name, ServerAddress, TunnelType, AllUserConnection, RememberCredential, IdleDisconnectSeconds, DnsSuffix, SplitTunneling, Force, PassThru, L2tpPsk, UseWinlogonCredential, EapConfigXmlStream, AuthenticationMethod, EncryptionLevel, MachineCertificateIssuerFilter, MachineCertificateEKUFilter, ServerList)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="CustomConfiguration" type="string "></param>
// <param name="DnsSuffix" type="string "></param>
// <param name="Force" type="bool "></param>
// <param name="IdleDisconnectSeconds" type="uint32 "></param>
// <param name="Name" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="PlugInApplicationID" type="string "></param>
// <param name="RememberCredential" type="bool "></param>
// <param name="ServerAddress" type="string "></param>
// <param name="ServerList" type="VpnServerAddress []"></param>
// <param name="SplitTunneling" type="bool "></param>

// <param name="cmdletOutput" type="ThirdPartyVpnConnection "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnConnection) NewByThirdParty( /* IN */ Name string,
	/* IN */ ServerAddress string,
	/* IN */ RememberCredential bool,
	/* IN */ SplitTunneling bool,
	/* IN */ PassThru bool,
	/* IN */ Force bool,
	/* IN */ ServerList []VpnServerAddress,
	/* IN */ DnsSuffix string,
	/* IN */ IdleDisconnectSeconds uint32,
	/* IN */ PlugInApplicationID string,
	/* IN */ CustomConfiguration string,
	/* OUT */ cmdletOutput ThirdPartyVpnConnection) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("NewByThirdParty", Name, ServerAddress, RememberCredential, SplitTunneling, PassThru, Force, ServerList, DnsSuffix, IdleDisconnectSeconds, PlugInApplicationID, CustomConfiguration)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
