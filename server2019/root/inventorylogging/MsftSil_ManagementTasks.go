// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.InventoryLogging
//////////////////////////////////////////////
package inventorylogging

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MsftSil_ManagementTasks struct
type MsftSil_ManagementTasks struct {
	*cim.WmiInstance
}

func NewMsftSil_ManagementTasksEx1(instance *cim.WmiInstance) (newInstance *MsftSil_ManagementTasks, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MsftSil_ManagementTasks{
		WmiInstance: tmp,
	}
	return
}

func NewMsftSil_ManagementTasksEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MsftSil_ManagementTasks, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MsftSil_ManagementTasks{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="certificateThumbprint" type="string "></param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="uri" type="string "></param>
func (instance *MsftSil_ManagementTasks) GetTargetUri( /* OUT */ uri string,
	/* OUT */ certificateThumbprint string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetTargetUri")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="certificateThumbprint" type="string "></param>
// <param name="uri" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MsftSil_ManagementTasks) SetTargetUri( /* IN */ uri string,
	/* IN */ certificateThumbprint string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetTargetUri", uri, certificateThumbprint)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
// <param name="state" type="uint8 "></param>
func (instance *MsftSil_ManagementTasks) GetLoggingState( /* OUT */ state uint8) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetLoggingState")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="state" type="uint8 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MsftSil_ManagementTasks) SetLoggingState( /* IN */ state uint8) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetLoggingState", state)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
// <param name="time" type="string "></param>
func (instance *MsftSil_ManagementTasks) GetLoggingTime( /* OUT */ time string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetLoggingTime")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="time" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MsftSil_ManagementTasks) SetLoggingTime( /* IN */ time string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetLoggingTime", time)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
