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

// MSFT_MpHeartBeat struct
type MSFT_MpHeartBeat struct {
	cim.WmiInstance
}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MpHeartBeat) Send() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Send")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
