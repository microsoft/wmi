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

// PS_DAEntryPoint struct
type PS_DAEntryPoint struct {
	*cim.WmiInstance
}

func NewPS_DAEntryPointEx1(instance *cim.WmiInstance) (newInstance *PS_DAEntryPoint, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_DAEntryPoint{
		WmiInstance: tmp,
	}
	return
}

func NewPS_DAEntryPointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_DAEntryPoint, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_DAEntryPoint{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="ClientIPv6Prefix" type="string "></param>
// <param name="ComputerName" type="string "></param>
// <param name="ConnectToAddress" type="string "></param>
// <param name="DeployNat" type="bool "></param>
// <param name="Force" type="bool "></param>
// <param name="GslbIP" type="string "></param>
// <param name="InternalInterface" type="string "></param>
// <param name="InternetInterface" type="string "></param>
// <param name="Name" type="string "></param>
// <param name="NoPrerequisite" type="bool "></param>
// <param name="PassThru" type="bool "></param>
// <param name="RemoteAccessServer" type="string "></param>
// <param name="ServerGpoName" type="string "></param>

// <param name="cmdletOutput" type="DAEntryPoint "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAEntryPoint) Add( /* IN */ ComputerName string,
	/* IN */ ServerGpoName string,
	/* IN */ Name string,
	/* IN */ GslbIP string,
	/* IN */ DeployNat bool,
	/* IN */ NoPrerequisite bool,
	/* IN */ PassThru bool,
	/* IN */ RemoteAccessServer string,
	/* IN */ InternalInterface string,
	/* IN */ InternetInterface string,
	/* IN */ ClientIPv6Prefix string,
	/* IN */ ConnectToAddress string,
	/* IN */ Force bool,
	/* OUT */ cmdletOutput DAEntryPoint) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Add", ComputerName, ServerGpoName, Name, GslbIP, DeployNat, NoPrerequisite, PassThru, RemoteAccessServer, InternalInterface, InternetInterface, ClientIPv6Prefix, ConnectToAddress, Force)
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
// <param name="Name" type="string "></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="DAEntryPoint "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAEntryPoint) Remove( /* IN */ ComputerName string,
	/* IN */ Name string,
	/* IN */ PassThru bool,
	/* IN */ Force bool,
	/* OUT */ cmdletOutput DAEntryPoint) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Remove", ComputerName, Name, PassThru, Force)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerName" type="string "></param>
// <param name="GslbIP" type="string "></param>
// <param name="Name" type="string "></param>
// <param name="NewName" type="string "></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="DAEntryPoint "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAEntryPoint) Set( /* IN */ ComputerName string,
	/* IN */ Name string,
	/* IN */ GslbIP string,
	/* IN */ PassThru bool,
	/* IN */ NewName string,
	/* OUT */ cmdletOutput DAEntryPoint) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Set", ComputerName, Name, GslbIP, PassThru, NewName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerName" type="string "></param>
// <param name="Name" type="string "></param>

// <param name="cmdletOutput" type="DAEntryPoint "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAEntryPoint) Get( /* IN */ ComputerName string,
	/* IN */ Name string,
	/* OUT */ cmdletOutput DAEntryPoint) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", ComputerName, Name)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
