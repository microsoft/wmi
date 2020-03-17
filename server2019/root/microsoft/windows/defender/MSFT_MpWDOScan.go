// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Defender
//////////////////////////////////////////////
package defender

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_MpWDOScan struct
type MSFT_MpWDOScan struct {
	cim.WmiInstance
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
