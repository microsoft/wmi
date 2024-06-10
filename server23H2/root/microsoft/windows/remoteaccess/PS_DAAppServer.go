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

// PS_DAAppServer struct
type PS_DAAppServer struct {
	*cim.WmiInstance
}

func NewPS_DAAppServerEx1(instance *cim.WmiInstance) (newInstance *PS_DAAppServer, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_DAAppServer{
		WmiInstance: tmp,
	}
	return
}

func NewPS_DAAppServerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_DAAppServer, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_DAAppServer{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="ComputerName" type="string "></param>
// <param name="GpoName" type="string []"></param>
// <param name="PassThru" type="bool "></param>
// <param name="SecurityGroupNameList" type="string []"></param>

// <param name="cmdletOutput" type="DAAppServer "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAAppServer) AddByAppServerSGGpo( /* IN */ GpoName []string,
	/* IN */ ComputerName string,
	/* IN */ PassThru bool,
	/* IN */ SecurityGroupNameList []string,
	/* OUT */ cmdletOutput DAAppServer) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("AddByAppServerSGGpo", GpoName, ComputerName, PassThru, SecurityGroupNameList)
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
// <param name="Name" type="string []"></param>
// <param name="PassThru" type="bool "></param>
// <param name="SecurityGroupName" type="string "></param>

// <param name="cmdletOutput" type="DAAppServer "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAAppServer) AddByAppServerToSGGpo( /* IN */ GpoName []string,
	/* IN */ ComputerName string,
	/* IN */ PassThru bool,
	/* IN */ SecurityGroupName string,
	/* IN */ Name []string,
	/* OUT */ cmdletOutput DAAppServer) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("AddByAppServerToSGGpo", GpoName, ComputerName, PassThru, SecurityGroupName, Name)
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
// <param name="Name" type="string []"></param>
// <param name="PassThru" type="bool "></param>
// <param name="SecurityGroupName" type="string "></param>

// <param name="cmdletOutput" type="DAAppServer "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAAppServer) RemoveByAppServerFromSGGpo( /* IN */ ComputerName string,
	/* IN */ PassThru bool,
	/* IN */ SecurityGroupName string,
	/* IN */ Name []string,
	/* IN */ DomainName []string,
	/* OUT */ cmdletOutput DAAppServer) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("RemoveByAppServerFromSGGpo", ComputerName, PassThru, SecurityGroupName, Name, DomainName)
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

// <param name="cmdletOutput" type="DAAppServer "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAAppServer) RemoveByAppServerSGGpo( /* IN */ ComputerName string,
	/* IN */ PassThru bool,
	/* IN */ SecurityGroupNameList []string,
	/* IN */ DomainName []string,
	/* OUT */ cmdletOutput DAAppServer) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("RemoveByAppServerSGGpo", ComputerName, PassThru, SecurityGroupNameList, DomainName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerName" type="string "></param>

// <param name="cmdletOutput" type="DAAppServer "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAAppServer) Get( /* IN */ ComputerName string,
	/* OUT */ cmdletOutput DAAppServer) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", ComputerName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
