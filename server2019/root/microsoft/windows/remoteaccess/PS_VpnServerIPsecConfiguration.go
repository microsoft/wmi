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

// PS_VpnServerIPsecConfiguration struct
type PS_VpnServerIPsecConfiguration struct { 
	*cim.WmiInstance
}

	func NewPS_VpnServerIPsecConfigurationEx1(instance *cim.WmiInstance) (newInstance *PS_VpnServerIPsecConfiguration, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &PS_VpnServerIPsecConfiguration {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewPS_VpnServerIPsecConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *PS_VpnServerIPsecConfiguration, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &PS_VpnServerIPsecConfiguration {
	WmiInstance: tmp,
	}
	return
	}
	

// 

// <param name="AuthenticationTransformConstants" type="uint32 "></param>
// <param name="CipherTransformConstants" type="uint32 "></param>
// <param name="CustomPolicy" type="bool "></param>
// <param name="DHGroup" type="uint32 "></param>
// <param name="EncryptionMethod" type="uint32 "></param>
// <param name="GrePorts" type="uint32 "></param>
// <param name="IdleDisconnectSeconds" type="uint32 "></param>
// <param name="Ikev2Ports" type="uint32 "></param>
// <param name="IntegrityCheckMethod" type="uint32 "></param>
// <param name="L2tpPorts" type="uint32 "></param>
// <param name="MMSALifeTimeSeconds" type="uint32 "></param>
// <param name="PassThru" type="bool "></param>
// <param name="PfsGroup" type="uint32 "></param>
// <param name="SADataSizeForRenegotiationKilobytes" type="uint32 "></param>
// <param name="SALifeTimeSeconds" type="uint32 "></param>
// <param name="SstpPorts" type="uint32 "></param>
// <param name="TunnelType" type="uint32 "></param>

// <param name="cmdletOutput" type="VpnServerIPsecConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnServerIPsecConfiguration) SetByCustomPolicy( /* IN */ PassThru bool,
 /* IN */ TunnelType uint32,
 /* IN */ IdleDisconnectSeconds uint32,
 /* IN */ SALifeTimeSeconds uint32,
 /* IN */ MMSALifeTimeSeconds uint32,
 /* IN */ SADataSizeForRenegotiationKilobytes uint32,
 /* IN */ CustomPolicy bool,
 /* IN */ EncryptionMethod uint32,
 /* IN */ IntegrityCheckMethod uint32,
 /* IN */ SstpPorts uint32,
 /* IN */ CipherTransformConstants uint32,
 /* IN */ PfsGroup uint32,
 /* IN */ AuthenticationTransformConstants uint32,
 /* IN */ DHGroup uint32,
 /* IN */ Ikev2Ports uint32,
 /* IN */ L2tpPorts uint32,
 /* IN */ GrePorts uint32,
 /* OUT */ cmdletOutput VpnServerIPsecConfiguration) (result uint32, err error) {retVal, err := instance.InvokeMethod("SetByCustomPolicy" , PassThru, TunnelType, IdleDisconnectSeconds, SALifeTimeSeconds, MMSALifeTimeSeconds, SADataSizeForRenegotiationKilobytes, CustomPolicy, EncryptionMethod, IntegrityCheckMethod, SstpPorts, CipherTransformConstants, PfsGroup, AuthenticationTransformConstants, DHGroup, Ikev2Ports, L2tpPorts, GrePorts)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="EncryptionType" type="string "></param>
// <param name="GrePorts" type="uint32 "></param>
// <param name="IdleDisconnectSeconds" type="uint32 "></param>
// <param name="Ikev2Ports" type="uint32 "></param>
// <param name="L2tpPorts" type="uint32 "></param>
// <param name="MMSALifeTimeSeconds" type="uint32 "></param>
// <param name="PassThru" type="bool "></param>
// <param name="SADataSizeForRenegotiationKilobytes" type="uint32 "></param>
// <param name="SALifeTimeSeconds" type="uint32 "></param>
// <param name="SstpPorts" type="uint32 "></param>
// <param name="TunnelType" type="uint32 "></param>

// <param name="cmdletOutput" type="VpnServerIPsecConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnServerIPsecConfiguration) SetByEncryptionType( /* IN */ PassThru bool,
 /* IN */ TunnelType uint32,
 /* IN */ EncryptionType string,
 /* IN */ IdleDisconnectSeconds uint32,
 /* IN */ Ikev2Ports uint32,
 /* IN */ SADataSizeForRenegotiationKilobytes uint32,
 /* IN */ SALifeTimeSeconds uint32,
 /* IN */ MMSALifeTimeSeconds uint32,
 /* IN */ L2tpPorts uint32,
 /* IN */ SstpPorts uint32,
 /* IN */ GrePorts uint32,
 /* OUT */ cmdletOutput VpnServerIPsecConfiguration) (result uint32, err error) {retVal, err := instance.InvokeMethod("SetByEncryptionType" , PassThru, TunnelType, EncryptionType, IdleDisconnectSeconds, Ikev2Ports, SADataSizeForRenegotiationKilobytes, SALifeTimeSeconds, MMSALifeTimeSeconds, L2tpPorts, SstpPorts, GrePorts)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="PassThru" type="bool "></param>
// <param name="RevertToDefault" type="bool "></param>
// <param name="TunnelType" type="uint32 "></param>

// <param name="cmdletOutput" type="VpnServerIPsecConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnServerIPsecConfiguration) SetByRevertToDefault( /* IN */ PassThru bool,
 /* IN */ TunnelType uint32,
 /* IN */ RevertToDefault bool,
 /* OUT */ cmdletOutput VpnServerIPsecConfiguration) (result uint32, err error) {retVal, err := instance.InvokeMethod("SetByRevertToDefault" , PassThru, TunnelType, RevertToDefault)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="TunnelType" type="uint32 "></param>

// <param name="cmdletOutput" type="VpnServerIPsecConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnServerIPsecConfiguration) Get( /* IN */ TunnelType uint32,
 /* OUT */ cmdletOutput VpnServerIPsecConfiguration) (result uint32, err error) {retVal, err := instance.InvokeMethod("Get" , TunnelType)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


