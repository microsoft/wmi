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

// Msvm_VssService struct
type Msvm_VssService struct {
	*Msvm_GuestService
}

func NewMsvm_VssServiceEx1(instance *cim.WmiInstance) (newInstance *Msvm_VssService, err error) {
	tmp, err := NewMsvm_GuestServiceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_VssService{
		Msvm_GuestService: tmp,
	}
	return
}

func NewMsvm_VssServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VssService, err error) {
	tmp, err := NewMsvm_GuestServiceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VssService{
		Msvm_GuestService: tmp,
	}
	return
}

//

// <param name="GuestClusterInformation" type="Msvm_GuestClusterInformation ">Guest cluster information returned</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VssService) QueryGuestClusterInformation( /* OUT */ GuestClusterInformation Msvm_GuestClusterInformation) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("QueryGuestClusterInformation")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
