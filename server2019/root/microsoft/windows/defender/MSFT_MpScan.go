// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.Defender
//////////////////////////////////////////////
package defender

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_MpScan struct
type MSFT_MpScan struct {
	*cim.WmiInstance
}

func NewMSFT_MpScanEx1(instance *cim.WmiInstance) (newInstance *MSFT_MpScan, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_MpScan{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_MpScanEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_MpScan, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_MpScan{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="ScanPath" type="string "></param>
// <param name="ScanType" type="uint8 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MpScan) Start( /* IN */ ScanType uint8,
	/* IN */ ScanPath string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Start", ScanType, ScanPath)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
