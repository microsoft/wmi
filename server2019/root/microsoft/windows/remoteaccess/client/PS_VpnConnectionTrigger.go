// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess.Client
//////////////////////////////////////////////
package client

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// PS_VpnConnectionTrigger struct
type PS_VpnConnectionTrigger struct {
	cim.WmiInstance
}

//

// <param name="ConnectionName" type="string "></param>

// <param name="cmdletOutput" type="VpnConnectionTrigger "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnConnectionTrigger) Get( /* IN */ ConnectionName string,
	/* OUT */ cmdletOutput VpnConnectionTrigger) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", ConnectionName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
