// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.RSOP
//////////////////////////////////////////////
package rsop
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
)

// RsopLoggingModeProvider struct
type RsopLoggingModeProvider struct { 
	*cim.WmiInstance
}

	func NewRsopLoggingModeProviderEx1(instance *cim.WmiInstance) (newInstance *RsopLoggingModeProvider, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &RsopLoggingModeProvider {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewRsopLoggingModeProviderEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *RsopLoggingModeProvider, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &RsopLoggingModeProvider {
	WmiInstance: tmp,
	}
	return
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
 /* OUT */ ExtendedInfo uint32) (err error) {_, err = instance.InvokeMethod("RsopCreateSession" , flags, userSid)
	if err != nil { return }
	return
	
}


// 

// <param name="nameSpace" type="string "></param>

// <param name="hResult" type="uint32 "></param>
func (instance *RsopLoggingModeProvider) RsopDeleteSession( /* IN */ nameSpace string,
 /* OUT */ hResult uint32) (err error) {_, err = instance.InvokeMethod("RsopDeleteSession" , nameSpace)
	if err != nil { return }
	return
	
}


// 

// <param name="hResult" type="uint32 "></param>
// <param name="userSids" type="string []"></param>
func (instance *RsopLoggingModeProvider) RsopEnumerateUsers( /* OUT */ userSids []string,
 /* OUT */ hResult uint32) (err error) {_, err = instance.InvokeMethod("RsopEnumerateUsers" )
	if err != nil { return }
	return
	
}


