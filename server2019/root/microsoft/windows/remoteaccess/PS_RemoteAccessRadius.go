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

// PS_RemoteAccessRadius struct
type PS_RemoteAccessRadius struct {
	*cim.WmiInstance
}

func NewPS_RemoteAccessRadiusEx1(instance *cim.WmiInstance) (newInstance *PS_RemoteAccessRadius, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_RemoteAccessRadius{
		WmiInstance: tmp,
	}
	return
}

func NewPS_RemoteAccessRadiusEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_RemoteAccessRadius, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_RemoteAccessRadius{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="AccountingOnOffMsg" type="string "></param>
// <param name="ComputerName" type="string "></param>
// <param name="EntrypointName" type="string "></param>
// <param name="MsgAuthenticator" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="Port" type="uint16 "></param>
// <param name="Purpose" type="string "></param>
// <param name="Score" type="uint8 "></param>
// <param name="ServerName" type="string "></param>
// <param name="SharedSecret" type="string "></param>
// <param name="Timeout" type="uint32 "></param>

// <param name="cmdletOutput" type="RemoteAccessRadiusServer []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessRadius) Add( /* IN */ ServerName string,
	/* IN */ SharedSecret string,
	/* IN */ ComputerName string,
	/* IN */ Port uint16,
	/* IN */ Score uint8,
	/* IN */ Timeout uint32,
	/* IN */ Purpose string,
	/* IN */ AccountingOnOffMsg string,
	/* IN */ MsgAuthenticator string,
	/* IN */ EntrypointName string,
	/* IN */ PassThru bool,
	/* OUT */ cmdletOutput []RemoteAccessRadiusServer) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Add", ServerName, SharedSecret, ComputerName, Port, Score, Timeout, Purpose, AccountingOnOffMsg, MsgAuthenticator, EntrypointName, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerName" type="string "></param>
// <param name="EntrypointName" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="Purpose" type="string "></param>
// <param name="ServerName" type="string "></param>

// <param name="cmdletOutput" type="RemoteAccessRadiusServerPurpose []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessRadius) Remove( /* IN */ ServerName string,
	/* IN */ Purpose string,
	/* IN */ ComputerName string,
	/* IN */ PassThru bool,
	/* IN */ EntrypointName string,
	/* OUT */ cmdletOutput []RemoteAccessRadiusServerPurpose) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Remove", ServerName, Purpose, ComputerName, PassThru, EntrypointName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AccountingOnOffMsg" type="string "></param>
// <param name="ComputerName" type="string "></param>
// <param name="EntrypointName" type="string "></param>
// <param name="MsgAuthenticator" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="Port" type="uint16 "></param>
// <param name="Purpose" type="string "></param>
// <param name="Score" type="uint8 "></param>
// <param name="ServerName" type="string "></param>
// <param name="SharedSecret" type="string "></param>
// <param name="Timeout" type="uint32 "></param>

// <param name="cmdletOutput" type="RemoteAccessRadiusServer []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessRadius) Set( /* IN */ ComputerName string,
	/* IN */ Purpose string,
	/* IN */ Port uint16,
	/* IN */ Score uint8,
	/* IN */ ServerName string,
	/* IN */ Timeout uint32,
	/* IN */ SharedSecret string,
	/* IN */ AccountingOnOffMsg string,
	/* IN */ MsgAuthenticator string,
	/* IN */ EntrypointName string,
	/* IN */ PassThru bool,
	/* OUT */ cmdletOutput []RemoteAccessRadiusServer) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Set", ComputerName, Purpose, Port, Score, ServerName, Timeout, SharedSecret, AccountingOnOffMsg, MsgAuthenticator, EntrypointName, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerName" type="string "></param>
// <param name="EntrypointName" type="string "></param>
// <param name="Purpose" type="string "></param>

// <param name="cmdletOutput" type="RemoteAccessRadiusServer []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessRadius) Get( /* IN */ Purpose string,
	/* IN */ ComputerName string,
	/* IN */ EntrypointName string,
	/* OUT */ cmdletOutput []RemoteAccessRadiusServer) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", Purpose, ComputerName, EntrypointName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
