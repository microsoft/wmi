// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.RSOP
//////////////////////////////////////////////
package rsop

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// RsopLoggingModeProvider struct
type RsopLoggingModeProvider struct {
	cim.WmiInstance
}

//

// <param name="flags" type="uint32 "></param>
// <param name="userSid" type="string "></param>

// <param name="ExtendedInfo" type="uint32 "></param>
// <param name="hResult" type="uint32 "></param>
// <param name="nameSpace" type="string "></param>
func (instance *RsopLoggingModeProvider) RsopCreateSession( /* IN */ flags uint32,
	/* IN */ userSid string,
	/* OUT */ nameSpace string,
	/* OUT */ hResult uint32,
	/* OUT */ ExtendedInfo uint32) (err error) {
	_, err = instance.InvokeMethod("RsopCreateSession", flags, userSid)
	if err != nil {
		return
	}
	return

}

//

// <param name="nameSpace" type="string "></param>

// <param name="hResult" type="uint32 "></param>
func (instance *RsopLoggingModeProvider) RsopDeleteSession( /* IN */ nameSpace string,
	/* OUT */ hResult uint32) (err error) {
	_, err = instance.InvokeMethod("RsopDeleteSession", nameSpace)
	if err != nil {
		return
	}
	return

}

//

// <param name="hResult" type="uint32 "></param>
// <param name="userSids" type="string []"></param>
func (instance *RsopLoggingModeProvider) RsopEnumerateUsers( /* OUT */ userSids []string,
	/* OUT */ hResult uint32) (err error) {
	_, err = instance.InvokeMethod("RsopEnumerateUsers")
	if err != nil {
		return
	}
	return

}
