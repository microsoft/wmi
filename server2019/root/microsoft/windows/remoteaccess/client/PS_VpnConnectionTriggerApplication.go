// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess.Client
//////////////////////////////////////////////
package client

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// PS_VpnConnectionTriggerApplication struct
type PS_VpnConnectionTriggerApplication struct {
	*cim.WmiInstance
}

func NewPS_VpnConnectionTriggerApplicationEx1(instance *cim.WmiInstance) (newInstance *PS_VpnConnectionTriggerApplication, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_VpnConnectionTriggerApplication{
		WmiInstance: tmp,
	}
	return
}

func NewPS_VpnConnectionTriggerApplicationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_VpnConnectionTriggerApplication, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_VpnConnectionTriggerApplication{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="ApplicationID" type="string []"></param>
// <param name="ConnectionName" type="string "></param>
// <param name="Force" type="bool "></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="VpnConnectionTriggerApplication "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnConnectionTriggerApplication) Add( /* IN */ ConnectionName string,
	/* IN */ ApplicationID []string,
	/* IN */ PassThru bool,
	/* IN */ Force bool,
	/* OUT */ cmdletOutput VpnConnectionTriggerApplication) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Add", ConnectionName, ApplicationID, PassThru, Force)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ApplicationID" type="string []"></param>
// <param name="ConnectionName" type="string "></param>
// <param name="Force" type="bool "></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="VpnConnectionTriggerApplication "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnConnectionTriggerApplication) Remove( /* IN */ ConnectionName string,
	/* IN */ ApplicationID []string,
	/* IN */ PassThru bool,
	/* IN */ Force bool,
	/* OUT */ cmdletOutput VpnConnectionTriggerApplication) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Remove", ConnectionName, ApplicationID, PassThru, Force)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
