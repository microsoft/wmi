// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.RemoteAccess.Server
//
// ////////////////////////////////////////////
package server

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// PS_RemoteAccessLocal struct
type PS_RemoteAccessLocal struct {
	*cim.WmiInstance
}

func NewPS_RemoteAccessLocalEx1(instance *cim.WmiInstance) (newInstance *PS_RemoteAccessLocal, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_RemoteAccessLocal{
		WmiInstance: tmp,
	}
	return
}

func NewPS_RemoteAccessLocalEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_RemoteAccessLocal, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_RemoteAccessLocal{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="ReturnValue" type="uint32 "></param>
// <param name="Status" type="uint8 []"></param>
func (instance *PS_RemoteAccessLocal) CheckServerPreRequisites( /* OUT */ Status []uint8) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CheckServerPreRequisites")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Options" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessLocal) Sync( /* IN */ Options uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Sync", Options)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="BehindNat" type="bool "></param>
// <param name="ConnectTo" type="string "></param>
// <param name="InternetInterface" type="string "></param>
// <param name="IntranetInterface" type="string "></param>
// <param name="isNlbDeployed" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="Status" type="bool "></param>
func (instance *PS_RemoteAccessLocal) CheckInterfaces( /* IN */ IntranetInterface string,
	/* IN */ InternetInterface string,
	/* IN */ BehindNat bool,
	/* IN */ ConnectTo string,
	/* IN */ isNlbDeployed bool,
	/* OUT */ Status bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CheckInterfaces", IntranetInterface, InternetInterface, BehindNat, ConnectTo, isNlbDeployed)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ConnectTo" type="string "></param>
// <param name="DeploymentMode" type="uint32 "></param>
// <param name="isNlbDeployed" type="bool "></param>

// <param name="DeploymentMode" type="uint32 "></param>
// <param name="InternetInterfaces" type="string []"></param>
// <param name="IntranetInterfaces" type="string []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessLocal) GetInterfaces( /* IN/OUT */ DeploymentMode uint32,
	/* IN */ ConnectTo string,
	/* IN */ isNlbDeployed bool,
	/* OUT */ IntranetInterfaces []string,
	/* OUT */ InternetInterfaces []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetInterfaces", ConnectTo, isNlbDeployed)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="InternetInterface" type="string "></param>
// <param name="IntranetInterface" type="string "></param>

// <param name="Result" type="DACorpPrefixConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessLocal) GetCorpPrefix( /* IN */ IntranetInterface string,
	/* IN */ InternetInterface string,
	/* OUT */ Result DACorpPrefixConfiguration) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetCorpPrefix", IntranetInterface, InternetInterface)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="bEnable" type="bool "></param>
// <param name="FirewallGroupName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessLocal) EnableFirewallGroup( /* IN */ FirewallGroupName string,
	/* IN */ bEnable bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("EnableFirewallGroup", FirewallGroupName, bEnable)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="FriendlyName" type="string "></param>
// <param name="Purpose" type="uint32 "></param>
// <param name="SubjectName" type="string "></param>

// <param name="EncodedCertificate" type="uint8 []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessLocal) CreateSelfSignedCertificate( /* IN */ SubjectName string,
	/* IN */ FriendlyName string,
	/* IN */ Purpose uint32,
	/* OUT */ EncodedCertificate []uint8) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateSelfSignedCertificate", SubjectName, FriendlyName, Purpose)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="SubjectName" type="string "></param>
// <param name="TemplateName" type="string "></param>

// <param name="EncodedCertificate" type="uint8 []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessLocal) EnrollWebServerCertificate( /* IN */ TemplateName string,
	/* IN */ SubjectName string,
	/* OUT */ EncodedCertificate []uint8) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("EnrollWebServerCertificate", TemplateName, SubjectName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ConfigVersion" type="RemoteAccessConfigurationVersion "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessLocal) GetConfigurationVersion( /* OUT */ ConfigVersion RemoteAccessConfigurationVersion) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetConfigurationVersion")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="IsS2SVpnEnabled" type="bool "></param>
// <param name="IsVpnEnabled" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessLocal) SetVpnPorts( /* IN */ IsVpnEnabled bool,
	/* IN */ IsS2SVpnEnabled bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetVpnPorts", IsVpnEnabled, IsS2SVpnEnabled)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="CertificateRawData" type="uint8 []"></param>
// <param name="CertificateRawDataType" type="uint32 "></param>
// <param name="Purpose" type="uint32 "></param>

// <param name="EncodedCertificate" type="DACertificateContext "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessLocal) GetDACertificateFromRawData( /* IN */ CertificateRawData []uint8,
	/* IN */ CertificateRawDataType uint32,
	/* IN */ Purpose uint32,
	/* OUT */ EncodedCertificate DACertificateContext) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetDACertificateFromRawData", CertificateRawData, CertificateRawDataType, Purpose)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Flags" type="uint32 "></param>
// <param name="IncludeSelfSigned" type="bool "></param>
// <param name="SubjectName" type="string "></param>

// <param name="EncodedCertificate" type="DACertificateContext "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessLocal) GetDACertificateFromSubjectName( /* IN */ SubjectName string,
	/* IN */ IncludeSelfSigned bool,
	/* IN */ Flags uint32,
	/* OUT */ EncodedCertificate DACertificateContext) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetDACertificateFromSubjectName", SubjectName, IncludeSelfSigned, Flags)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ConnectToAddress" type="string []"></param>
// <param name="IncludeSelfSigned" type="bool "></param>
// <param name="MachineNamesForAllSites" type="string []"></param>
// <param name="SiteAddresses" type="DASiteAddresses []"></param>

// <param name="EncodedCertificate" type="DACertificateContext "></param>
// <param name="ResolvedAddress" type="string []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessLocal) FindNLSCertificate( /* IN */ SiteAddresses []DASiteAddresses,
	/* IN */ ConnectToAddress []string,
	/* IN */ MachineNamesForAllSites []string,
	/* IN */ IncludeSelfSigned bool,
	/* OUT */ ResolvedAddress []string,
	/* OUT */ EncodedCertificate DACertificateContext) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("FindNLSCertificate", SiteAddresses, ConnectToAddress, MachineNamesForAllSites, IncludeSelfSigned)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Purpose" type="uint32 "></param>

// <param name="EncodedCertificate" type="DACertificateContext []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessLocal) GetValidDACertificates( /* IN */ Purpose uint32,
	/* OUT */ EncodedCertificate []DACertificateContext) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetValidDACertificates", Purpose)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="CertificateHash" type="uint8 []"></param>
// <param name="UseHttps" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessLocal) SetSSTPCertificate( /* IN */ CertificateHash []uint8,
	/* IN */ UseHttps bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetSSTPCertificate", CertificateHash, UseHttps)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
