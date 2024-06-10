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

// PS_IPFilter struct
type PS_IPFilter struct {
	*cim.WmiInstance
}

func NewPS_IPFilterEx1(instance *cim.WmiInstance) (newInstance *PS_IPFilter, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_IPFilter{
		WmiInstance: tmp,
	}
	return
}

func NewPS_IPFilterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_IPFilter, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_IPFilter{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="Action" type="uint32 "></param>
// <param name="AddressFamily" type="uint32 "></param>
// <param name="Direction" type="uint32 "></param>
// <param name="InterfaceAlias" type="string "></param>
// <param name="List" type="string []"></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="InterfaceIpFilter []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_IPFilter) Add( /* IN */ InterfaceAlias string,
	/* IN */ Action uint32,
	/* IN */ List []string,
	/* IN */ Direction uint32,
	/* IN */ PassThru bool,
	/* IN */ AddressFamily uint32,
	/* OUT */ cmdletOutput []InterfaceIpFilter) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Add", InterfaceAlias, Action, List, Direction, PassThru, AddressFamily)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AddressFamily" type="uint32 "></param>
// <param name="Direction" type="uint32 "></param>
// <param name="Force" type="bool "></param>
// <param name="InterfaceAlias" type="string "></param>
// <param name="List" type="string []"></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="InterfaceIpFilter []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_IPFilter) Remove( /* IN */ InterfaceAlias string,
	/* IN */ Direction uint32,
	/* IN */ List []string,
	/* IN */ PassThru bool,
	/* IN */ Force bool,
	/* IN */ AddressFamily uint32,
	/* OUT */ cmdletOutput []InterfaceIpFilter) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Remove", InterfaceAlias, Direction, List, PassThru, Force, AddressFamily)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Direction" type="uint32 "></param>
// <param name="InterfaceAlias" type="string "></param>

// <param name="cmdletOutput" type="InterfaceIpFilter []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_IPFilter) Get( /* IN */ InterfaceAlias string,
	/* IN */ Direction uint32,
	/* OUT */ cmdletOutput []InterfaceIpFilter) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", InterfaceAlias, Direction)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Action" type="uint32 "></param>
// <param name="AddressFamily" type="uint32 "></param>
// <param name="Direction" type="uint32 "></param>
// <param name="InterfaceAlias" type="string "></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="InterfaceIpFilter []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_IPFilter) Set( /* IN */ Action uint32,
	/* IN */ Direction uint32,
	/* IN */ InterfaceAlias string,
	/* IN */ AddressFamily uint32,
	/* IN */ PassThru bool,
	/* OUT */ cmdletOutput []InterfaceIpFilter) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Set", Action, Direction, InterfaceAlias, AddressFamily, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
