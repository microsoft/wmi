// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_Group struct
type Win32_Group struct {
	Win32_Account
}

//

// <param name="Name" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Group) Rename( /* IN */ Name string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Rename", Name)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
