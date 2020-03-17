// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.protectionManagement
//////////////////////////////////////////////
package protectionmanagement

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_MpScan struct
type MSFT_MpScan struct {
	cim.WmiInstance
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
