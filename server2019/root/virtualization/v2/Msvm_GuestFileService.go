// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.virtualization.v2
//
// ////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_GuestFileService struct
type Msvm_GuestFileService struct {
	*Msvm_GuestService
}

func NewMsvm_GuestFileServiceEx1(instance *cim.WmiInstance) (newInstance *Msvm_GuestFileService, err error) {
	tmp, err := NewMsvm_GuestServiceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_GuestFileService{
		Msvm_GuestService: tmp,
	}
	return
}

func NewMsvm_GuestFileServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_GuestFileService, err error) {
	tmp, err := NewMsvm_GuestServiceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_GuestFileService{
		Msvm_GuestService: tmp,
	}
	return
}

//

// <param name="CopyFileToGuestSettings" type="string []"></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_GuestFileService) CopyFilesToGuest( /* IN */ CopyFileToGuestSettings []string,
	/* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("CopyFilesToGuest", Action, PercentComplete, Timeout, CopyFileToGuestSettings)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
