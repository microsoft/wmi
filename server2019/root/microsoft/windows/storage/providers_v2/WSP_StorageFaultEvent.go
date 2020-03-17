// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

// WSP_StorageFaultEvent struct
type WSP_StorageFaultEvent struct {
	MSFT_StorageFaultEvent
}

// This method manually fires alerts

// <param name="FaultAlert" type="WSP_StorageFaultEvent ">Copy of the alert payload to be fired</param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *WSP_StorageFaultEvent) FireAlert( /* IN */ FaultAlert WSP_StorageFaultEvent) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("FireAlert", FaultAlert)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
