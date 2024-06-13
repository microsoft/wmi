// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_ShutdownComponent struct
type Msvm_ShutdownComponent struct {
	*CIM_LogicalDevice
}

func NewMsvm_ShutdownComponentEx1(instance *cim.WmiInstance) (newInstance *Msvm_ShutdownComponent, err error) {
	tmp, err := NewCIM_LogicalDeviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_ShutdownComponent{
		CIM_LogicalDevice: tmp,
	}
	return
}

func NewMsvm_ShutdownComponentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_ShutdownComponent, err error) {
	tmp, err := NewCIM_LogicalDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_ShutdownComponent{
		CIM_LogicalDevice: tmp,
	}
	return
}

//

// <param name="Force" type="bool "></param>
// <param name="Reason" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ShutdownComponent) InitiateShutdown( /* IN */ Force bool,
	/* IN */ Reason string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("InitiateShutdown", Force, Reason)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Force" type="bool "></param>
// <param name="Reason" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ShutdownComponent) InitiateReboot( /* IN */ Force bool,
	/* IN */ Reason string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("InitiateReboot", Force, Reason)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ShutdownComponent) InitiateHibernate() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("InitiateHibernate")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
