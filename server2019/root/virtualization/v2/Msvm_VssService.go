// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// Msvm_VssService struct
type Msvm_VssService struct {
	Msvm_GuestService
}

//

// <param name="GuestClusterInformation" type="Msvm_GuestClusterInformation ">Guest cluster information returned</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_VssService) QueryGuestClusterInformation( /* OUT */ GuestClusterInformation Msvm_GuestClusterInformation) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("QueryGuestClusterInformation")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
