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

// MSFT_MpWDOScan struct
type MSFT_MpWDOScan struct {
	*cim.WmiInstance
}

func NewMSFT_MpWDOScanEx1(instance *cim.WmiInstance) (newInstance *MSFT_MpWDOScan, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_MpWDOScan{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_MpWDOScanEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_MpWDOScan, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_MpWDOScan{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MpWDOScan) Start() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Start")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
