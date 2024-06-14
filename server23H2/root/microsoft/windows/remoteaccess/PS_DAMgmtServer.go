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

// PS_DAMgmtServer struct
type PS_DAMgmtServer struct {
	*cim.WmiInstance
}

func NewPS_DAMgmtServerEx1(instance *cim.WmiInstance) (newInstance *PS_DAMgmtServer, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_DAMgmtServer{
		WmiInstance: tmp,
	}
	return
}

func NewPS_DAMgmtServerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_DAMgmtServer, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_DAMgmtServer{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="ComputerName" type="string "></param>
// <param name="MgmtServer" type="string []"></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="string []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAMgmtServer) Add( /* IN */ MgmtServer []string,
	/* IN */ ComputerName string,
	/* IN */ PassThru bool,
	/* OUT */ cmdletOutput []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Add", MgmtServer, ComputerName, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerName" type="string "></param>
// <param name="MgmtServer" type="string []"></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="string []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAMgmtServer) Remove( /* IN */ MgmtServer []string,
	/* IN */ ComputerName string,
	/* IN */ PassThru bool,
	/* OUT */ cmdletOutput []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Remove", MgmtServer, ComputerName, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerName" type="string "></param>
// <param name="Type" type="string "></param>

// <param name="cmdletOutput" type="string []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAMgmtServer) Get( /* IN */ ComputerName string,
	/* IN */ Type string,
	/* OUT */ cmdletOutput []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", ComputerName, Type)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerName" type="string "></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="DAMgmtServer "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAMgmtServer) Update( /* IN */ PassThru bool,
	/* IN */ ComputerName string,
	/* OUT */ cmdletOutput DAMgmtServer) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Update", PassThru, ComputerName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
