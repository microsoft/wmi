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

// PS_DAServer struct
type PS_DAServer struct { 
	*cim.WmiInstance
}

	func NewPS_DAServerEx1(instance *cim.WmiInstance) (newInstance *PS_DAServer, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &PS_DAServer {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewPS_DAServerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *PS_DAServer, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &PS_DAServer {
	WmiInstance: tmp,
	}
	return
	}
	

// 

// <param name="ComputerName" type="string "></param>
// <param name="DAInstallType" type="string "></param>
// <param name="Force" type="bool "></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="DAServer "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAServer) SetByChangeDAInstallationType( /* IN */ ComputerName string,
 /* IN */ PassThru bool,
 /* IN */ Force bool,
 /* IN */ DAInstallType string,
 /* OUT */ cmdletOutput DAServer) (result uint32, err error) {retVal, err := instance.InvokeMethod("SetByChangeDAInstallationType" , ComputerName, PassThru, Force, DAInstallType)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="ClientIPv6Prefix" type="string "></param>
// <param name="ComputerName" type="string "></param>
// <param name="ConnectToAddress" type="string "></param>
// <param name="DisableComputerCertAuthentication" type="bool "></param>
// <param name="Force" type="bool "></param>
// <param name="InternalIPv6Prefix" type="string []"></param>
// <param name="PassThru" type="bool "></param>
// <param name="TeredoState" type="string "></param>

// <param name="cmdletOutput" type="DAServer "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAServer) SetByDisableComputerCertAuth( /* IN */ ComputerName string,
 /* IN */ PassThru bool,
 /* IN */ Force bool,
 /* IN */ InternalIPv6Prefix []string,
 /* IN */ ClientIPv6Prefix string,
 /* IN */ DisableComputerCertAuthentication bool,
 /* IN */ TeredoState string,
 /* IN */ ConnectToAddress string,
 /* OUT */ cmdletOutput DAServer) (result uint32, err error) {retVal, err := instance.InvokeMethod("SetByDisableComputerCertAuth" , ComputerName, PassThru, Force, InternalIPv6Prefix, ClientIPv6Prefix, DisableComputerCertAuthentication, TeredoState, ConnectToAddress)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="ClientIPv6Prefix" type="string "></param>
// <param name="ComputerName" type="string "></param>
// <param name="ConnectToAddress" type="string "></param>
// <param name="EntrypointName" type="string "></param>
// <param name="Force" type="bool "></param>
// <param name="IntermediateRootCertificate" type="bool "></param>
// <param name="InternalIPv6Prefix" type="string []"></param>
// <param name="IPsecRootCertificate" type="uint8 []"></param>
// <param name="PassThru" type="bool "></param>
// <param name="TeredoState" type="string "></param>
// <param name="UserAuthentication" type="string "></param>

// <param name="cmdletOutput" type="DAServer "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAServer) SetByEnableComputerCertAuth( /* IN */ ComputerName string,
 /* IN */ PassThru bool,
 /* IN */ Force bool,
 /* IN */ UserAuthentication string,
 /* IN */ IPsecRootCertificate []uint8,
 /* IN */ IntermediateRootCertificate bool,
 /* IN */ EntrypointName string,
 /* IN */ ClientIPv6Prefix string,
 /* IN */ ConnectToAddress string,
 /* IN */ InternalIPv6Prefix []string,
 /* IN */ TeredoState string,
 /* OUT */ cmdletOutput DAServer) (result uint32, err error) {retVal, err := instance.InvokeMethod("SetByEnableComputerCertAuth" , ComputerName, PassThru, Force, UserAuthentication, IPsecRootCertificate, IntermediateRootCertificate, EntrypointName, ClientIPv6Prefix, ConnectToAddress, InternalIPv6Prefix, TeredoState)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="ComputerName" type="string "></param>
// <param name="EntrypointName" type="string "></param>

// <param name="cmdletOutput" type="DAServer "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAServer) Get( /* IN */ ComputerName string,
 /* IN */ EntrypointName string,
 /* OUT */ cmdletOutput DAServer) (result uint32, err error) {retVal, err := instance.InvokeMethod("Get" , ComputerName, EntrypointName)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


