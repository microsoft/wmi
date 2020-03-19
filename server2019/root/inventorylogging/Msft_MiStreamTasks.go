// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.InventoryLogging
//////////////////////////////////////////////
package inventorylogging

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msft_MiStreamTasks struct
type Msft_MiStreamTasks struct {
	*cim.WmiInstance
}

func NewMsft_MiStreamTasksEx1(instance *cim.WmiInstance) (newInstance *Msft_MiStreamTasks, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &Msft_MiStreamTasks{
		WmiInstance: tmp,
	}
	return
}

func NewMsft_MiStreamTasksEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msft_MiStreamTasks, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msft_MiStreamTasks{
		WmiInstance: tmp,
	}
	return
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
