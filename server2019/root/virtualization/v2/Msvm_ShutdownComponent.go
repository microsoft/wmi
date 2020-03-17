// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// Msvm_ShutdownComponent struct
type Msvm_ShutdownComponent struct {
	CIM_LogicalDevice
}

//

// <param name="Force" type="bool "></param>
// <param name="Reason" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ShutdownComponent) InitiateShutdown( /* IN */ Force bool,
	/* IN */ Reason string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("InitiateShutdown", Force, Reason)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Force" type="bool "></param>
// <param name="Reason" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ShutdownComponent) InitiateReboot( /* IN */ Force bool,
	/* IN */ Reason string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("InitiateReboot", Force, Reason)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ShutdownComponent) InitiateHibernate() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("InitiateHibernate")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
