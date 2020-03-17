// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.ManagementTools
//////////////////////////////////////////////
package managementtools

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_MTRegistryTasks struct
type MSFT_MTRegistryTasks struct {
	cim.WmiInstance
}

//

// <param name="KeyName" type="string "></param>
// <param name="Options" type="uint8 "></param>
// <param name="Value" type="string "></param>

// <param name="Results" type="MSFT_MTRegistryObject []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MTRegistryTasks) Search( /* IN */ Value string,
	/* IN */ KeyName string,
	/* IN */ Options uint8,
	/* OUT */ Results []MSFT_MTRegistryObject) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Search", Value, KeyName, Options)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
