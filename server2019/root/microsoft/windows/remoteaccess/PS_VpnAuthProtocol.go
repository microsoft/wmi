// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess
//////////////////////////////////////////////
package remoteaccess
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
)

// PS_VpnAuthProtocol struct
type PS_VpnAuthProtocol struct { 
	*cim.WmiInstance
}

	func NewPS_VpnAuthProtocolEx1(instance *cim.WmiInstance) (newInstance *PS_VpnAuthProtocol, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &PS_VpnAuthProtocol {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewPS_VpnAuthProtocolEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *PS_VpnAuthProtocol, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &PS_VpnAuthProtocol {
	WmiInstance: tmp,
	}
	return
	}
	

// 

// <param name="CertificateAdvertised" type="uint8 []"></param>
// <param name="CertificateEKUsToAccept" type="string []"></param>
// <param name="PassThru" type="bool "></param>
// <param name="RootCertificateNameToAccept" type="uint8 []"></param>
// <param name="SharedSecret" type="string "></param>
// <param name="TunnelAuthProtocolsAdvertised" type="string "></param>
// <param name="UserAuthProtocolAccepted" type="string []"></param>

// <param name="cmdletOutput" type="VpnAuthProtocol "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnAuthProtocol) Set( /* IN */ UserAuthProtocolAccepted []string,
 /* IN */ TunnelAuthProtocolsAdvertised string,
 /* IN */ RootCertificateNameToAccept []uint8,
 /* IN */ CertificateAdvertised []uint8,
 /* IN */ SharedSecret string,
 /* IN */ PassThru bool,
 /* IN */ CertificateEKUsToAccept []string,
 /* OUT */ cmdletOutput VpnAuthProtocol) (result uint32, err error) {retVal, err := instance.InvokeMethod("Set" , UserAuthProtocolAccepted, TunnelAuthProtocolsAdvertised, RootCertificateNameToAccept, CertificateAdvertised, SharedSecret, PassThru, CertificateEKUsToAccept)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="cmdletOutput" type="VpnAuthProtocol "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnAuthProtocol) Get( /* OUT */ cmdletOutput VpnAuthProtocol) (result uint32, err error) {retVal, err := instance.InvokeMethod("Get" )
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


