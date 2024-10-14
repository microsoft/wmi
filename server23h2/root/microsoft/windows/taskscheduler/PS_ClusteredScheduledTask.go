// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Microsoft.Windows.TaskScheduler
//////////////////////////////////////////////
package taskscheduler

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// PS_ClusteredScheduledTask struct
type PS_ClusteredScheduledTask struct {
	*cim.WmiInstance
}

func NewPS_ClusteredScheduledTaskEx1(instance *cim.WmiInstance) (newInstance *PS_ClusteredScheduledTask, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_ClusteredScheduledTask{
		WmiInstance: tmp,
	}
	return
}

func NewPS_ClusteredScheduledTaskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_ClusteredScheduledTask, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_ClusteredScheduledTask{
		WmiInstance: tmp,
	}
	return
}

// 99

// <param name="Cluster" type="string "></param>
// <param name="InputObject" type="MSFT_ScheduledTask ">24</param>
// <param name="Resource" type="string "></param>
// <param name="TaskName" type="string ">100</param>
// <param name="TaskType" type="int32 "></param>

// <param name="cmdletOutput" type="MSFT_ClusteredScheduledTask ">101</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_ClusteredScheduledTask) RegisterByObject( /* IN */ InputObject MSFT_ScheduledTask,
	/* IN */ Cluster string,
	/* IN */ TaskName string,
	/* IN */ TaskType int32,
	/* IN */ Resource string,
	/* OUT */ cmdletOutput MSFT_ClusteredScheduledTask) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("RegisterByObject", InputObject, Cluster, TaskName, TaskType, Resource)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// 99

// <param name="Action" type="MSFT_TaskAction []">31</param>
// <param name="Cluster" type="string "></param>
// <param name="Description" type="string ">32</param>
// <param name="Resource" type="string "></param>
// <param name="Settings" type="MSFT_TaskSettings ">33</param>
// <param name="TaskName" type="string ">100</param>
// <param name="TaskType" type="int32 "></param>
// <param name="Trigger" type="MSFT_TaskTrigger []">86</param>

// <param name="cmdletOutput" type="MSFT_ClusteredScheduledTask ">101</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_ClusteredScheduledTask) RegisterByParams( /* IN */ TaskName string,
	/* IN */ Settings MSFT_TaskSettings,
	/* IN */ Description string,
	/* IN */ Trigger []MSFT_TaskTrigger,
	/* IN */ Cluster string,
	/* IN */ Action []MSFT_TaskAction,
	/* IN */ TaskType int32,
	/* IN */ Resource string,
	/* OUT */ cmdletOutput MSFT_ClusteredScheduledTask) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("RegisterByParams", TaskName, Settings, Description, Trigger, Cluster, Action, TaskType, Resource)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// 99

// <param name="Cluster" type="string "></param>
// <param name="Resource" type="string "></param>
// <param name="TaskName" type="string ">100</param>
// <param name="TaskType" type="int32 "></param>
// <param name="Xml" type="string ">35</param>

// <param name="cmdletOutput" type="MSFT_ClusteredScheduledTask ">101</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_ClusteredScheduledTask) RegisterByXml( /* IN */ Xml string,
	/* IN */ TaskName string,
	/* IN */ Cluster string,
	/* IN */ TaskType int32,
	/* IN */ Resource string,
	/* OUT */ cmdletOutput MSFT_ClusteredScheduledTask) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("RegisterByXml", Xml, TaskName, Cluster, TaskType, Resource)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// 102

// <param name="Cluster" type="string "></param>
// <param name="TaskName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_ClusteredScheduledTask) UnregisterByName( /* IN */ Cluster string,
	/* IN */ TaskName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("UnregisterByName", Cluster, TaskName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

// 102

// <param name="InputObject" type="MSFT_ClusteredScheduledTask "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_ClusteredScheduledTask) UnregisterByObject( /* IN */ InputObject MSFT_ClusteredScheduledTask) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("UnregisterByObject", InputObject)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

// 103

// <param name="Cluster" type="string ">105</param>
// <param name="TaskName" type="string ">104</param>
// <param name="TaskType" type="int32 "></param>

// <param name="cmdletOutput" type="MSFT_ClusteredScheduledTask []">106</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_ClusteredScheduledTask) Get( /* IN */ TaskName string,
	/* IN */ Cluster string,
	/* IN */ TaskType int32,
	/* OUT */ cmdletOutput []MSFT_ClusteredScheduledTask) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", TaskName, Cluster, TaskType)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// 107

// <param name="Cluster" type="string "></param>
// <param name="InputObject" type="MSFT_ScheduledTask ">24</param>
// <param name="TaskName" type="string "></param>

// <param name="cmdletOutput" type="MSFT_ClusteredScheduledTask "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_ClusteredScheduledTask) SetByObject( /* IN */ TaskName string,
	/* IN */ Cluster string,
	/* IN */ InputObject MSFT_ScheduledTask,
	/* OUT */ cmdletOutput MSFT_ClusteredScheduledTask) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetByObject", TaskName, Cluster, InputObject)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// 107

// <param name="Action" type="MSFT_TaskAction []">31</param>
// <param name="Cluster" type="string "></param>
// <param name="Description" type="string ">32</param>
// <param name="Settings" type="MSFT_TaskSettings ">33</param>
// <param name="TaskName" type="string "></param>
// <param name="Trigger" type="MSFT_TaskTrigger []">86</param>

// <param name="cmdletOutput" type="MSFT_ClusteredScheduledTask "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_ClusteredScheduledTask) SetByParams( /* IN */ Action []MSFT_TaskAction,
	/* IN */ Settings MSFT_TaskSettings,
	/* IN */ Cluster string,
	/* IN */ TaskName string,
	/* IN */ Trigger []MSFT_TaskTrigger,
	/* IN */ Description string,
	/* OUT */ cmdletOutput MSFT_ClusteredScheduledTask) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetByParams", Action, Settings, Cluster, TaskName, Trigger, Description)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// 107

// <param name="Cluster" type="string "></param>
// <param name="TaskName" type="string "></param>
// <param name="Xml" type="string ">35</param>

// <param name="cmdletOutput" type="MSFT_ClusteredScheduledTask "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_ClusteredScheduledTask) SetByXml( /* IN */ Xml string,
	/* IN */ Cluster string,
	/* IN */ TaskName string,
	/* OUT */ cmdletOutput MSFT_ClusteredScheduledTask) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetByXml", Xml, Cluster, TaskName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
