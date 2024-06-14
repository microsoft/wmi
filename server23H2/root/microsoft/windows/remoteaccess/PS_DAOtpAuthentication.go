// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess
//////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// PS_DAOtpAuthentication struct
type PS_DAOtpAuthentication struct {
	*cim.WmiInstance
}

func NewPS_DAOtpAuthenticationEx1(instance *cim.WmiInstance) (newInstance *PS_DAOtpAuthentication, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_DAOtpAuthentication{
		WmiInstance: tmp,
	}
	return
}

func NewPS_DAOtpAuthenticationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_DAOtpAuthentication, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_DAOtpAuthentication{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="CAServer" type="string []"></param>
// <param name="CertificateTemplateName" type="string "></param>
// <param name="ComputerName" type="string "></param>
// <param name="Force" type="bool "></param>
// <param name="PassThru" type="bool "></param>
// <param name="RadiusPort" type="uint16 "></param>
// <param name="RadiusServer" type="string "></param>
// <param name="SharedSecret" type="string "></param>
// <param name="SigningCertificateTemplateName" type="string "></param>
// <param name="UserSecurityGroupName" type="string "></param>

// <param name="cmdletOutput" type="DAOtpAuthentication "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAOtpAuthentication) Enable( /* IN */ RadiusServer string,
	/* IN */ ComputerName string,
	/* IN */ RadiusPort uint16,
	/* IN */ CAServer []string,
	/* IN */ CertificateTemplateName string,
	/* IN */ SharedSecret string,
	/* IN */ UserSecurityGroupName string,
	/* IN */ Force bool,
	/* IN */ PassThru bool,
	/* IN */ SigningCertificateTemplateName string,
	/* OUT */ cmdletOutput DAOtpAuthentication) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Enable", RadiusServer, ComputerName, RadiusPort, CAServer, CertificateTemplateName, SharedSecret, UserSecurityGroupName, Force, PassThru, SigningCertificateTemplateName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerName" type="string "></param>
// <param name="Force" type="bool "></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="DAOtpAuthentication "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAOtpAuthentication) Disable( /* IN */ ComputerName string,
	/* IN */ Force bool,
	/* IN */ PassThru bool,
	/* OUT */ cmdletOutput DAOtpAuthentication) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Disable", ComputerName, Force, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerName" type="string "></param>

// <param name="cmdletOutput" type="DAOtpAuthentication "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAOtpAuthentication) Get( /* IN */ ComputerName string,
	/* OUT */ cmdletOutput DAOtpAuthentication) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", ComputerName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="CAServer" type="string []"></param>
// <param name="CertificateTemplateName" type="string "></param>
// <param name="ComputerName" type="string "></param>
// <param name="DisableTwoFactorBypass" type="bool "></param>
// <param name="PassThru" type="bool "></param>
// <param name="SigningCertificateTemplateName" type="string "></param>
// <param name="UserSecurityGroupName" type="string "></param>

// <param name="cmdletOutput" type="DAOtpAuthentication "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAOtpAuthentication) Set( /* IN */ ComputerName string,
	/* IN */ CertificateTemplateName string,
	/* IN */ CAServer []string,
	/* IN */ UserSecurityGroupName string,
	/* IN */ DisableTwoFactorBypass bool,
	/* IN */ PassThru bool,
	/* IN */ SigningCertificateTemplateName string,
	/* OUT */ cmdletOutput DAOtpAuthentication) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Set", ComputerName, CertificateTemplateName, CAServer, UserSecurityGroupName, DisableTwoFactorBypass, PassThru, SigningCertificateTemplateName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
