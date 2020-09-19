// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.Windows.ServerManager
//////////////////////////////////////////////
package servermanager

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_ServerManagerTasks struct
type MSFT_ServerManagerTasks struct {
	*cim.WmiInstance
}

func NewMSFT_ServerManagerTasksEx1(instance *cim.WmiInstance) (newInstance *MSFT_ServerManagerTasks, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_ServerManagerTasks{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_ServerManagerTasksEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ServerManagerTasks, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ServerManagerTasks{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="BatchSize" type="uint32 "></param>
// <param name="CollectorName" type="string "></param>
// <param name="CounterPaths" type="string []"></param>
// <param name="Timestamps" type="string []"></param>

// <param name="cmdletOutput" type="MSFT_ServerPerformanceCounterSamples []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ServerManagerTasks) GetCounterSamplesAtTime( /* IN */ CollectorName string,
	/* IN */ CounterPaths []string,
	/* IN */ Timestamps []string,
	/* IN */ BatchSize uint32,
	/* OUT */ cmdletOutput []MSFT_ServerPerformanceCounterSamples) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetCounterSamplesAtTime", CollectorName, CounterPaths, Timestamps, BatchSize)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="BatchSize" type="uint32 "></param>
// <param name="CollectorName" type="string "></param>
// <param name="CounterPaths" type="string []"></param>
// <param name="EndTime" type="string "></param>
// <param name="StartTime" type="string "></param>

// <param name="cmdletOutput" type="MSFT_ServerPerformanceCounterSamples []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ServerManagerTasks) GetCounterSamplesInTimeRange( /* IN */ CollectorName string,
	/* IN */ CounterPaths []string,
	/* IN */ StartTime string,
	/* IN */ EndTime string,
	/* IN */ BatchSize uint32,
	/* OUT */ cmdletOutput []MSFT_ServerPerformanceCounterSamples) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetCounterSamplesInTimeRange", CollectorName, CounterPaths, StartTime, EndTime, BatchSize)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="CollectorName" type="string "></param>

// <param name="cmdletOutput" type="uint8 "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ServerManagerTasks) GetPerformanceCollectorState( /* IN */ CollectorName string,
	/* OUT */ cmdletOutput uint8) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetPerformanceCollectorState", CollectorName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="CollectorName" type="string "></param>
// <param name="State" type="uint8 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ServerManagerTasks) SetPerformanceCollectorState( /* IN */ CollectorName string,
	/* IN */ State uint8) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetPerformanceCollectorState", CollectorName, State)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="BatchSize" type="uint32 "></param>
// <param name="BpaXPaths" type="string []"></param>

// <param name="cmdletOutput" type="MSFT_ServerBpaResult []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ServerManagerTasks) GetServerBpaResult( /* IN */ BpaXPaths []string,
	/* IN */ BatchSize uint32,
	/* OUT */ cmdletOutput []MSFT_ServerBpaResult) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetServerBpaResult", BpaXPaths, BatchSize)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="cmdletOutput" type="string []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ServerManagerTasks) GetServerClusterName( /* OUT */ cmdletOutput []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetServerClusterName")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="BatchSize" type="uint32 "></param>
// <param name="EndTimes" type="uint64 []"></param>
// <param name="Levels" type="uint8 []"></param>
// <param name="Logs" type="string []"></param>
// <param name="QueryFileIds" type="int32 []"></param>
// <param name="QueryFiles" type="string []"></param>
// <param name="StartTimes" type="uint64 []"></param>

// <param name="cmdletOutput" type="MSFT_ServerEventDetail []"></param>
// <param name="LatestEventTimestamp" type="uint64 "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ServerManagerTasks) GetServerEventDetail( /* IN */ Logs []string,
	/* IN */ Levels []uint8,
	/* IN */ StartTimes []uint64,
	/* IN */ EndTimes []uint64,
	/* IN */ BatchSize uint32,
	/* IN */ QueryFiles []string,
	/* IN */ QueryFileIds []int32,
	/* OUT */ LatestEventTimestamp uint64,
	/* OUT */ cmdletOutput []MSFT_ServerEventDetail) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetServerEventDetail", Logs, Levels, StartTimes, EndTimes, BatchSize, QueryFiles, QueryFileIds)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="BatchSize" type="uint32 "></param>
// <param name="FilterFlags" type="uint32 "></param>

// <param name="cmdletOutput" type="MSFT_ServerFeature []"></param>
// <param name="RequiresReboot" type="bool "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ServerManagerTasks) GetServerFeature( /* IN */ FilterFlags uint32,
	/* IN */ BatchSize uint32,
	/* OUT */ RequiresReboot bool,
	/* OUT */ cmdletOutput []MSFT_ServerFeature) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetServerFeature", FilterFlags, BatchSize)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="SMServerId" type="string "></param>

// <param name="ClusterInformation" type="MSFT_ServerClusterInformation "></param>
// <param name="EventLogs" type="MSFT_ServerEventLog []"></param>
// <param name="NetworkAdapters" type="MSFT_ServerNetworkAdapter []"></param>
// <param name="OperatingSystem" type="MSFT_ServerOperatingSystem "></param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="ServerInventory" type="MSFT_ServerInventory "></param>
func (instance *MSFT_ServerManagerTasks) GetServerInventory( /* IN */ SMServerId string,
	/* OUT */ ServerInventory MSFT_ServerInventory,
	/* OUT */ OperatingSystem MSFT_ServerOperatingSystem,
	/* OUT */ ClusterInformation MSFT_ServerClusterInformation,
	/* OUT */ NetworkAdapters []MSFT_ServerNetworkAdapter,
	/* OUT */ EventLogs []MSFT_ServerEventLog) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetServerInventory", SMServerId)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="BatchSize" type="uint32 "></param>
// <param name="Services" type="string []"></param>

// <param name="cmdletOutput" type="MSFT_ServerServiceDetail []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ServerManagerTasks) GetServerServiceDetail( /* IN */ Services []string,
	/* IN */ BatchSize uint32,
	/* OUT */ cmdletOutput []MSFT_ServerServiceDetail) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetServerServiceDetail", Services, BatchSize)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="CollectorName" type="string "></param>
// <param name="MillisecondThreshold" type="uint64 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ServerManagerTasks) RemoveServerPerformanceLog( /* IN */ CollectorName string,
	/* IN */ MillisecondThreshold uint64) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RemoveServerPerformanceLog", CollectorName, MillisecondThreshold)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="BatchSize" type="uint32 "></param>
// <param name="FilterXml" type="string "></param>
// <param name="ReverseDirection" type="bool "></param>
// <param name="Skip" type="uint64 "></param>
// <param name="Top" type="uint64 "></param>

// <param name="cmdletOutput" type="MSFT_ServerEventDetail []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_ServerManagerTasks) GetServerEventDetailEx( /* IN */ FilterXml string,
	/* IN */ Skip uint64,
	/* IN */ Top uint64,
	/* IN */ ReverseDirection bool,
	/* IN */ BatchSize uint32,
	/* OUT */ cmdletOutput []MSFT_ServerEventDetail) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetServerEventDetailEx", FilterXml, Skip, Top, ReverseDirection, BatchSize)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
