// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfigurationProxy
//////////////////////////////////////////////
package desiredstateconfigurationproxy

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_DscProxy struct
type MSFT_DscProxy struct {
	cim.WmiInstance
}

// Get resource state.

// <param name="ConfigurationData" type="uint8 []"></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="state" type="bool "></param>
func (instance *MSFT_DscProxy) GetResourceState( /* IN */ ConfigurationData []uint8,
	/* OUT */ state bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetResourceState", ConfigurationData)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
