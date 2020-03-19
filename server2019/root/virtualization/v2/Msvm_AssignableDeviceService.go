// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_AssignableDeviceService struct
type Msvm_AssignableDeviceService struct {
	*CIM_Service
}

func NewMsvm_AssignableDeviceServiceEx1(instance *cim.WmiInstance) (newInstance *Msvm_AssignableDeviceService, err error) {
	tmp, err := NewCIM_ServiceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_AssignableDeviceService{
		CIM_Service: tmp,
	}
	return
}

func NewMsvm_AssignableDeviceServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_AssignableDeviceService, err error) {
	tmp, err := NewCIM_ServiceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_AssignableDeviceService{
		CIM_Service: tmp,
	}
	return
}

//

// <param name="DismountSettingData" type="string "></param>

// <param name="DismountedDeviceInstancePath" type="string "></param>
// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_AssignableDeviceService) DismountAssignableDevice( /* IN */ DismountSettingData string,
	/* OUT */ DismountedDeviceInstancePath string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("DismountAssignableDevice", Action, PercentComplete, Timeout, DismountSettingData)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="DeviceInstancePath" type="string "></param>
// <param name="DeviceLocationPath" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="MountedDeviceInstancePath" type="string "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_AssignableDeviceService) MountAssignableDevice( /* IN */ DeviceInstancePath string,
	/* IN */ DeviceLocationPath string,
	/* OUT */ MountedDeviceInstancePath string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("MountAssignableDevice", Action, PercentComplete, Timeout, DeviceInstancePath, DeviceLocationPath)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

func (instance *Msvm_AssignableDeviceService) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ComputerSystem")
}
