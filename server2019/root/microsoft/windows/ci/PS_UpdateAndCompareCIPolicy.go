// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.CI
//////////////////////////////////////////////
package ci

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// PS_UpdateAndCompareCIPolicy struct
type PS_UpdateAndCompareCIPolicy struct {
	cim.WmiInstance
}

//

// <param name="FilePath" type="string "></param>

// <param name="cmdletOutput" type="uint64 "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_UpdateAndCompareCIPolicy) Compare( /* IN */ FilePath string,
	/* OUT */ cmdletOutput uint64) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Compare", FilePath)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="FilePath" type="string "></param>

// <param name="cmdletOutput" type="uint64 "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_UpdateAndCompareCIPolicy) Update( /* IN */ FilePath string,
	/* OUT */ cmdletOutput uint64) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Update", FilePath)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="cmdletOutput" type="uint64 "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_UpdateAndCompareCIPolicy) Delete( /* OUT */ cmdletOutput uint64) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Delete")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
