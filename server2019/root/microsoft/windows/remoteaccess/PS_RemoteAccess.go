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

// PS_RemoteAccess struct
type PS_RemoteAccess struct {
	*cim.WmiInstance
}

func NewPS_RemoteAccessEx1(instance *cim.WmiInstance) (newInstance *PS_RemoteAccess, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_RemoteAccess{
		WmiInstance: tmp,
	}
	return
}

func NewPS_RemoteAccessEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_RemoteAccess, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_RemoteAccess{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="ComputerName" type="string "></param>
// <param name="Prerequisite" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="Status" type="bool "></param>
func (instance *PS_RemoteAccess) InstallByDAPrerequisiteChecks( /* IN */ ComputerName string,
	/* IN */ Prerequisite bool,
	/* OUT */ Status bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("InstallByDAPrerequisiteChecks", ComputerName, Prerequisite)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ClientGpoName" type="string "></param>
// <param name="ComputerName" type="string "></param>
// <param name="ConnectToAddress" type="string "></param>
// <param name="DAInstallType" type="string "></param>
// <param name="DeployNat" type="bool "></param>
// <param name="Force" type="bool "></param>
// <param name="InternalInterface" type="string "></param>
// <param name="InternetInterface" type="string "></param>
// <param name="NlsCertificate" type="uint8 []"></param>
// <param name="NlsUrl" type="string "></param>
// <param name="NoPrerequisite" type="bool "></param>
// <param name="PassThru" type="bool "></param>
// <param name="ServerGpoName" type="string "></param>

// <param name="cmdletOutput" type="RemoteAccessCommon "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccess) InstallByDirectAccess( /* IN */ ComputerName string,
	/* IN */ DAInstallType string,
	/* IN */ ClientGpoName string,
	/* IN */ InternalInterface string,
	/* IN */ InternetInterface string,
	/* IN */ NlsCertificate []uint8,
	/* IN */ NlsUrl string,
	/* IN */ NoPrerequisite bool,
	/* IN */ ServerGpoName string,
	/* IN */ ConnectToAddress string,
	/* IN */ DeployNat bool,
	/* IN */ PassThru bool,
	/* IN */ Force bool,
	/* OUT */ cmdletOutput RemoteAccessCommon) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("InstallByDirectAccess", ComputerName, DAInstallType, ClientGpoName, InternalInterface, InternetInterface, NlsCertificate, NlsUrl, NoPrerequisite, ServerGpoName, ConnectToAddress, DeployNat, PassThru, Force)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="CapacityKbps" type="uint64 "></param>
// <param name="ComputerName" type="string "></param>
// <param name="MsgAuthenticator" type="string "></param>
// <param name="MultiTenancy" type="bool "></param>
// <param name="PassThru" type="bool "></param>
// <param name="RadiusPort" type="uint16 "></param>
// <param name="RadiusScore" type="uint8 "></param>
// <param name="RadiusServer" type="string "></param>
// <param name="RadiusTimeout" type="uint32 "></param>
// <param name="SharedSecret" type="string "></param>
// <param name="VpnType" type="string "></param>

// <param name="cmdletOutput" type="RemoteAccessCommon "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccess) InstallByMultiTenant( /* IN */ ComputerName string,
	/* IN */ MultiTenancy bool,
	/* IN */ VpnType string,
	/* IN */ MsgAuthenticator string,
	/* IN */ PassThru bool,
	/* IN */ RadiusPort uint16,
	/* IN */ RadiusScore uint8,
	/* IN */ RadiusServer string,
	/* IN */ RadiusTimeout uint32,
	/* IN */ SharedSecret string,
	/* IN */ CapacityKbps uint64,
	/* OUT */ cmdletOutput RemoteAccessCommon) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("InstallByMultiTenant", ComputerName, MultiTenancy, VpnType, MsgAuthenticator, PassThru, RadiusPort, RadiusScore, RadiusServer, RadiusTimeout, SharedSecret, CapacityKbps)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerName" type="string "></param>
// <param name="EntrypointName" type="string "></param>
// <param name="IPAddressRange" type="string []"></param>
// <param name="IPv6Prefix" type="string "></param>
// <param name="Legacy" type="bool "></param>
// <param name="MsgAuthenticator" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="RadiusPort" type="uint16 "></param>
// <param name="RadiusScore" type="uint8 "></param>
// <param name="RadiusServer" type="string "></param>
// <param name="RadiusTimeout" type="uint32 "></param>
// <param name="SharedSecret" type="string "></param>
// <param name="VpnType" type="string "></param>

// <param name="cmdletOutput" type="RemoteAccessCommon "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccess) InstallByVpn( /* IN */ ComputerName string,
	/* IN */ VpnType string,
	/* IN */ IPAddressRange []string,
	/* IN */ RadiusServer string,
	/* IN */ Legacy bool,
	/* IN */ SharedSecret string,
	/* IN */ RadiusTimeout uint32,
	/* IN */ RadiusScore uint8,
	/* IN */ RadiusPort uint16,
	/* IN */ MsgAuthenticator string,
	/* IN */ PassThru bool,
	/* IN */ IPv6Prefix string,
	/* IN */ EntrypointName string,
	/* OUT */ cmdletOutput RemoteAccessCommon) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("InstallByVpn", ComputerName, VpnType, IPAddressRange, RadiusServer, Legacy, SharedSecret, RadiusTimeout, RadiusScore, RadiusPort, MsgAuthenticator, PassThru, IPv6Prefix, EntrypointName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerName" type="string "></param>
// <param name="EntrypointName" type="string "></param>
// <param name="Force" type="bool "></param>
// <param name="VpnType" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccess) Uninstall( /* IN */ VpnType string,
	/* IN */ ComputerName string,
	/* IN */ EntrypointName string,
	/* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Uninstall", VpnType, ComputerName, EntrypointName, Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="CapacityKbps" type="uint64 "></param>
// <param name="ComputerName" type="string "></param>
// <param name="Force" type="bool "></param>
// <param name="InternalInterface" type="string "></param>
// <param name="InternetInterface" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="SslCertificate" type="uint8 []"></param>
// <param name="UseHttp" type="bool "></param>

// <param name="cmdletOutput" type="RemoteAccessCommon "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccess) Set( /* IN */ SslCertificate []uint8,
	/* IN */ ComputerName string,
	/* IN */ InternetInterface string,
	/* IN */ InternalInterface string,
	/* IN */ CapacityKbps uint64,
	/* IN */ UseHttp bool,
	/* IN */ Force bool,
	/* IN */ PassThru bool,
	/* OUT */ cmdletOutput RemoteAccessCommon) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Set", SslCertificate, ComputerName, InternetInterface, InternalInterface, CapacityKbps, UseHttp, Force, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerName" type="string "></param>
// <param name="EntrypointName" type="string "></param>

// <param name="cmdletOutput" type="RemoteAccessCommon "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccess) Get( /* IN */ ComputerName string,
	/* IN */ EntrypointName string,
	/* OUT */ cmdletOutput RemoteAccessCommon) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", ComputerName, EntrypointName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
