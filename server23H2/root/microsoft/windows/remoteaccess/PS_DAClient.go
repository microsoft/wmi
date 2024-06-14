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

// PS_DAClient struct
type PS_DAClient struct {
	*cim.WmiInstance
}

func NewPS_DAClientEx1(instance *cim.WmiInstance) (newInstance *PS_DAClient, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_DAClient{
		WmiInstance: tmp,
	}
	return
}

func NewPS_DAClientEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_DAClient, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_DAClient{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="ComputerName" type="string "></param>
// <param name="DownlevelGpoName" type="string []"></param>
// <param name="DownlevelSecurityGroupNameList" type="string []"></param>
// <param name="EntrypointName" type="string "></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="DAClient "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAClient) AddByClientDownlevelSGGpo( /* IN */ ComputerName string,
	/* IN */ PassThru bool,
	/* IN */ DownlevelGpoName []string,
	/* IN */ DownlevelSecurityGroupNameList []string,
	/* IN */ EntrypointName string,
	/* OUT */ cmdletOutput DAClient) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("AddByClientDownlevelSGGpo", ComputerName, PassThru, DownlevelGpoName, DownlevelSecurityGroupNameList, EntrypointName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerName" type="string "></param>
// <param name="GpoName" type="string []"></param>
// <param name="PassThru" type="bool "></param>
// <param name="SecurityGroupNameList" type="string []"></param>

// <param name="cmdletOutput" type="DAClient "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAClient) AddByClientSGGpo( /* IN */ ComputerName string,
	/* IN */ PassThru bool,
	/* IN */ SecurityGroupNameList []string,
	/* IN */ GpoName []string,
	/* OUT */ cmdletOutput DAClient) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("AddByClientSGGpo", ComputerName, PassThru, SecurityGroupNameList, GpoName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerName" type="string "></param>
// <param name="DownlevelDomainName" type="string []"></param>
// <param name="DownlevelSecurityGroupNameList" type="string []"></param>
// <param name="EntrypointName" type="string "></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="DAClient "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAClient) RemoveByClientDownlevelSGGpo( /* IN */ ComputerName string,
	/* IN */ PassThru bool,
	/* IN */ DownlevelSecurityGroupNameList []string,
	/* IN */ EntrypointName string,
	/* IN */ DownlevelDomainName []string,
	/* OUT */ cmdletOutput DAClient) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("RemoveByClientDownlevelSGGpo", ComputerName, PassThru, DownlevelSecurityGroupNameList, EntrypointName, DownlevelDomainName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerName" type="string "></param>
// <param name="DomainName" type="string []"></param>
// <param name="PassThru" type="bool "></param>
// <param name="SecurityGroupNameList" type="string []"></param>

// <param name="cmdletOutput" type="DAClient "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAClient) RemoveByClientSGGpo( /* IN */ ComputerName string,
	/* IN */ PassThru bool,
	/* IN */ SecurityGroupNameList []string,
	/* IN */ DomainName []string,
	/* OUT */ cmdletOutput DAClient) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("RemoveByClientSGGpo", ComputerName, PassThru, SecurityGroupNameList, DomainName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerName" type="string "></param>
// <param name="Downlevel" type="string "></param>
// <param name="ForceTunnel" type="string "></param>
// <param name="OnlyRemoteComputers" type="string "></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="DAClientSettings "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAClient) Set( /* IN */ ComputerName string,
	/* IN */ ForceTunnel string,
	/* IN */ OnlyRemoteComputers string,
	/* IN */ Downlevel string,
	/* IN */ PassThru bool,
	/* OUT */ cmdletOutput DAClientSettings) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Set", ComputerName, ForceTunnel, OnlyRemoteComputers, Downlevel, PassThru)
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

// <param name="cmdletOutput" type="DAClient "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAClient) Get( /* IN */ ComputerName string,
	/* IN */ EntrypointName string,
	/* OUT */ cmdletOutput DAClient) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", ComputerName, EntrypointName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
