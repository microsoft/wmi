// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.InventoryLogging
//////////////////////////////////////////////
package inventorylogging

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msft_MiStreamTasks struct
type Msft_MiStreamTasks struct {
	cim.WmiInstance
}

//

// <param name="Filename" type="string "></param>

// <param name="Results" type="interface{} []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msft_MiStreamTasks) Collect( /* IN */ Filename string,
	/* OUT */ Results []interface{}) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Collect", Filename)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="CheckCollectionHistory" type="bool "></param>
// <param name="Filename" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msft_MiStreamTasks) Push( /* IN */ Filename string,
	/* IN */ CheckCollectionHistory bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Push", Filename, CheckCollectionHistory)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msft_MiStreamTasks) Flush() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Flush")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
