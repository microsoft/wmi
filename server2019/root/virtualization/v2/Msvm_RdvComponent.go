// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// Msvm_RdvComponent struct
type Msvm_RdvComponent struct {
	CIM_LogicalDevice
}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_RdvComponent) EnableEndPoints() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("EnableEndPoints")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
