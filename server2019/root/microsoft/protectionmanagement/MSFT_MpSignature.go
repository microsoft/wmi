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

// MSFT_MpSignature struct
type MSFT_MpSignature struct {
	cim.WmiInstance
}

//

// <param name="UpdateSource" type="uint8 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MpSignature) Update( /* IN */ UpdateSource uint8) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Update", UpdateSource)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
