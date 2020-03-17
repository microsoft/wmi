// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_PlatformIdentifier struct
type MSFT_PlatformIdentifier struct {
	cim.WmiInstance
}

//

// <param name="Name" type="string "></param>

// <param name="Identifier" type="string "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PlatformIdentifier) GetPlatformIdentifier( /* IN */ Name string,
	/* OUT */ Identifier string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetPlatformIdentifier", Name)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
