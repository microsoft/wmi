// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

// WSP_HealthActionEvent struct
type WSP_HealthActionEvent struct {
	MSFT_HealthActionEvent
}

// This method manually fires alerts

// <param name="HealthActionAlert" type="WSP_HealthActionEvent ">Copy of the alert payload to be fired</param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *WSP_HealthActionEvent) FireAlert( /* IN */ HealthActionAlert WSP_HealthActionEvent) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("FireAlert", HealthActionAlert)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
