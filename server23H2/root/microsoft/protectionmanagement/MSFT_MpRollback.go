// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.ProtectionManagement
//////////////////////////////////////////////
package protectionmanagement

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_MpRollback struct
type MSFT_MpRollback struct {
	*cim.WmiInstance
}

func NewMSFT_MpRollbackEx1(instance *cim.WmiInstance) (newInstance *MSFT_MpRollback, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_MpRollback{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_MpRollbackEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_MpRollback, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_MpRollback{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="Engine" type="bool "></param>
// <param name="Platform" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MpRollback) Start( /* IN */ Engine bool,
	/* IN */ Platform bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Start", Engine, Platform)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
