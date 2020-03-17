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

// PS_EapConfiguration struct
type PS_EapConfiguration struct {
	cim.WmiInstance
}

//

// <param name="UseWinlogonCredential" type="bool "></param>

// <param name="cmdletOutput" type="EapConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_EapConfiguration) NewByEapMsChapv2Auth( /* IN */ UseWinlogonCredential bool,
	/* OUT */ cmdletOutput EapConfiguration) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("NewByEapMsChapv2Auth", UseWinlogonCredential)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Tls" type="bool "></param>
// <param name="UserCertificate" type="bool "></param>
// <param name="VerifyServerIdentity" type="bool "></param>

// <param name="cmdletOutput" type="EapConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_EapConfiguration) NewByEapTlsAuth( /* IN */ Tls bool,
	/* IN */ UserCertificate bool,
	/* IN */ VerifyServerIdentity bool,
	/* OUT */ cmdletOutput EapConfiguration) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("NewByEapTlsAuth", Tls, UserCertificate, VerifyServerIdentity)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Ttls" type="bool "></param>
// <param name="TunnledEapAuthMethod" type="string "></param>
// <param name="TunnledNonEapAuthMethod" type="string "></param>
// <param name="UseWinlogonCredential" type="bool "></param>

// <param name="cmdletOutput" type="EapConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_EapConfiguration) NewByEapTtlsAuth( /* IN */ Ttls bool,
	/* IN */ UseWinlogonCredential bool,
	/* IN */ TunnledNonEapAuthMethod string,
	/* IN */ TunnledEapAuthMethod string,
	/* OUT */ cmdletOutput EapConfiguration) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("NewByEapTtlsAuth", Ttls, UseWinlogonCredential, TunnledNonEapAuthMethod, TunnledEapAuthMethod)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="EnableNap" type="bool "></param>
// <param name="FastReconnect" type="bool "></param>
// <param name="Peap" type="bool "></param>
// <param name="TunnledEapAuthMethod" type="string "></param>
// <param name="VerifyServerIdentity" type="bool "></param>

// <param name="cmdletOutput" type="EapConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_EapConfiguration) NewByPeapAuth( /* IN */ Peap bool,
	/* IN */ VerifyServerIdentity bool,
	/* IN */ EnableNap bool,
	/* IN */ FastReconnect bool,
	/* IN */ TunnledEapAuthMethod string,
	/* OUT */ cmdletOutput EapConfiguration) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("NewByPeapAuth", Peap, VerifyServerIdentity, EnableNap, FastReconnect, TunnledEapAuthMethod)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
