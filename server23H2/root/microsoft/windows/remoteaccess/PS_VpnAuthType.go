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

// PS_VpnAuthType struct
type PS_VpnAuthType struct {
	*cim.WmiInstance
}

func NewPS_VpnAuthTypeEx1(instance *cim.WmiInstance) (newInstance *PS_VpnAuthType, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_VpnAuthType{
		WmiInstance: tmp,
	}
	return
}

func NewPS_VpnAuthTypeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_VpnAuthType, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_VpnAuthType{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="ComputerName" type="string "></param>
// <param name="EntrypointName" type="string "></param>
// <param name="MsgAuthenticator" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="RadiusPort" type="uint16 "></param>
// <param name="RadiusScore" type="uint8 "></param>
// <param name="RadiusServer" type="string "></param>
// <param name="RadiusTimeout" type="uint32 "></param>
// <param name="SharedSecret" type="string "></param>
// <param name="Type" type="string "></param>

// <param name="cmdletOutput" type="VpnAuth "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnAuthType) Set( /* IN */ Type string,
	/* IN */ RadiusServer string,
	/* IN */ SharedSecret string,
	/* IN */ RadiusTimeout uint32,
	/* IN */ RadiusScore uint8,
	/* IN */ RadiusPort uint16,
	/* IN */ ComputerName string,
	/* IN */ MsgAuthenticator string,
	/* IN */ EntrypointName string,
	/* IN */ PassThru bool,
	/* OUT */ cmdletOutput VpnAuth) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Set", Type, RadiusServer, SharedSecret, RadiusTimeout, RadiusScore, RadiusPort, ComputerName, MsgAuthenticator, EntrypointName, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
