// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_Service struct
type Win32_Service struct {
	*Win32_BaseService

	//
	CheckPoint uint32

	//
	DelayedAutoStart bool

	//
	ProcessId uint32

	//
	WaitHint uint32
}

func NewWin32_ServiceEx1(instance *cim.WmiInstance) (newInstance *Win32_Service, err error) {
	tmp, err := NewWin32_BaseServiceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_Service{
		Win32_BaseService: tmp,
	}
	return
}

func NewWin32_ServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_Service, err error) {
	tmp, err := NewWin32_BaseServiceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_Service{
		Win32_BaseService: tmp,
	}
	return
}

// SetCheckPoint sets the value of CheckPoint for the instance
func (instance *Win32_Service) SetPropertyCheckPoint(value uint32) (err error) {
	return instance.SetProperty("CheckPoint", value)
}

// GetCheckPoint gets the value of CheckPoint for the instance
func (instance *Win32_Service) GetPropertyCheckPoint() (value uint32, err error) {
	retValue, err := instance.GetProperty("CheckPoint")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDelayedAutoStart sets the value of DelayedAutoStart for the instance
func (instance *Win32_Service) SetPropertyDelayedAutoStart(value bool) (err error) {
	return instance.SetProperty("DelayedAutoStart", value)
}

// GetDelayedAutoStart gets the value of DelayedAutoStart for the instance
func (instance *Win32_Service) GetPropertyDelayedAutoStart() (value bool, err error) {
	retValue, err := instance.GetProperty("DelayedAutoStart")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProcessId sets the value of ProcessId for the instance
func (instance *Win32_Service) SetPropertyProcessId(value uint32) (err error) {
	return instance.SetProperty("ProcessId", value)
}

// GetProcessId gets the value of ProcessId for the instance
func (instance *Win32_Service) GetPropertyProcessId() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProcessId")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWaitHint sets the value of WaitHint for the instance
func (instance *Win32_Service) SetPropertyWaitHint(value uint32) (err error) {
	return instance.SetProperty("WaitHint", value)
}

// GetWaitHint gets the value of WaitHint for the instance
func (instance *Win32_Service) GetPropertyWaitHint() (value uint32, err error) {
	retValue, err := instance.GetProperty("WaitHint")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="Descriptor" type="Win32_SecurityDescriptor "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Service) GetSecurityDescriptor( /* OUT */ Descriptor Win32_SecurityDescriptor) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetSecurityDescriptor")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Descriptor" type="Win32_SecurityDescriptor "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Service) SetSecurityDescriptor( /* IN */ Descriptor Win32_SecurityDescriptor) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetSecurityDescriptor", Descriptor)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
