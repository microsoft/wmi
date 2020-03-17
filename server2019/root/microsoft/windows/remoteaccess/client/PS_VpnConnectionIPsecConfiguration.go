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

// PS_VpnConnectionIPsecConfiguration struct
type PS_VpnConnectionIPsecConfiguration struct {
	cim.WmiInstance
}

//

// <param name="AllUserConnection" type="bool "></param>
// <param name="AuthenticationTransformConstants" type="uint32 "></param>
// <param name="CipherTransformConstants" type="uint32 "></param>
// <param name="ConnectionName" type="string "></param>
// <param name="DHGroup" type="uint32 "></param>
// <param name="EncryptionMethod" type="uint32 "></param>
// <param name="Force" type="bool "></param>
// <param name="IntegrityCheckMethod" type="uint32 "></param>
// <param name="PassThru" type="bool "></param>
// <param name="PfsGroup" type="uint32 "></param>

// <param name="cmdletOutput" type="VpnConnectionIPsecConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnConnectionIPsecConfiguration) SetByCustomPolicy( /* IN */ ConnectionName string,
	/* IN */ AuthenticationTransformConstants uint32,
	/* IN */ CipherTransformConstants uint32,
	/* IN */ EncryptionMethod uint32,
	/* IN */ IntegrityCheckMethod uint32,
	/* IN */ PfsGroup uint32,
	/* IN */ DHGroup uint32,
	/* IN */ PassThru bool,
	/* IN */ Force bool,
	/* IN */ AllUserConnection bool,
	/* OUT */ cmdletOutput VpnConnectionIPsecConfiguration) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetByCustomPolicy", ConnectionName, AuthenticationTransformConstants, CipherTransformConstants, EncryptionMethod, IntegrityCheckMethod, PfsGroup, DHGroup, PassThru, Force, AllUserConnection)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AllUserConnection" type="bool "></param>
// <param name="ConnectionName" type="string "></param>
// <param name="Force" type="bool "></param>
// <param name="PassThru" type="bool "></param>
// <param name="RevertToDefault" type="bool "></param>

// <param name="cmdletOutput" type="VpnConnectionIPsecConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnConnectionIPsecConfiguration) SetIpSecByDefault( /* IN */ ConnectionName string,
	/* IN */ RevertToDefault bool,
	/* IN */ PassThru bool,
	/* IN */ Force bool,
	/* IN */ AllUserConnection bool,
	/* OUT */ cmdletOutput VpnConnectionIPsecConfiguration) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetIpSecByDefault", ConnectionName, RevertToDefault, PassThru, Force, AllUserConnection)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
