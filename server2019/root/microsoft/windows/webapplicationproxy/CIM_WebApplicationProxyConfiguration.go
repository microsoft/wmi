// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.WebApplicationProxy
//////////////////////////////////////////////
package webapplicationproxy
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
)

// CIM_WebApplicationProxyConfiguration struct
type CIM_WebApplicationProxyConfiguration struct { 
	*cim.WmiInstance
}

	func NewCIM_WebApplicationProxyConfigurationEx1(instance *cim.WmiInstance) (newInstance *CIM_WebApplicationProxyConfiguration, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &CIM_WebApplicationProxyConfiguration {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewCIM_WebApplicationProxyConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *CIM_WebApplicationProxyConfiguration, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &CIM_WebApplicationProxyConfiguration {
	WmiInstance: tmp,
	}
	return
	}
	

// 

// <param name="cmdletOutput" type="AppProxyGlobalConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_WebApplicationProxyConfiguration) Get( /* OUT */ cmdletOutput AppProxyGlobalConfiguration) (result uint32, err error) {retVal, err := instance.InvokeMethod("Get" )
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="ADFSSignOutURL" type="string "></param>
// <param name="ADFSTokenAcceptanceDurationSec" type="uint32 "></param>
// <param name="ADFSTokenSigningCertificatePublicKey" type="string "></param>
// <param name="ADFSUrl" type="string "></param>
// <param name="ADFSWebApplicationProxyRelyingPartyUri" type="string "></param>
// <param name="ConfigurationChangesPollingIntervalSec" type="uint32 "></param>
// <param name="ConnectedServersName" type="string []"></param>
// <param name="OAuthAuthenticationURL" type="string "></param>
// <param name="RegenerateAccessCookiesEncryptionKey" type="bool "></param>
// <param name="UpgradeConfigurationVersion" type="bool "></param>
// <param name="UserIdleTimeoutAction" type="string "></param>
// <param name="UserIdleTimeoutSec" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_WebApplicationProxyConfiguration) Set( /* IN */ ADFSUrl string,
 /* IN */ ADFSTokenSigningCertificatePublicKey string,
 /* IN */ ADFSWebApplicationProxyRelyingPartyUri string,
 /* IN */ RegenerateAccessCookiesEncryptionKey bool,
 /* IN */ ConnectedServersName []string,
 /* IN */ OAuthAuthenticationURL string,
 /* IN */ ConfigurationChangesPollingIntervalSec uint32,
 /* IN */ UpgradeConfigurationVersion bool,
 /* IN */ ADFSTokenAcceptanceDurationSec uint32,
 /* IN */ ADFSSignOutURL string,
 /* IN */ UserIdleTimeoutSec uint32,
 /* IN */ UserIdleTimeoutAction string) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("Set" , ADFSUrl, ADFSTokenSigningCertificatePublicKey, ADFSWebApplicationProxyRelyingPartyUri, RegenerateAccessCookiesEncryptionKey, ConnectedServersName, OAuthAuthenticationURL, ConfigurationChangesPollingIntervalSec, UpgradeConfigurationVersion, ADFSTokenAcceptanceDurationSec, ADFSSignOutURL, UserIdleTimeoutSec, UserIdleTimeoutAction);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


