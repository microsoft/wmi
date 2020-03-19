// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.SDDC.Management
//////////////////////////////////////////////
package management

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// SDDC_Server struct
type SDDC_Server struct {
	*cim.WmiInstance

	//
	Alerts []SDDC_Alert

	//
	BuildNumber string

	//
	Cluster string

	//
	Domain string

	//
	FreePhysicalMemoryInBytes uint64

	//
	Id string

	//
	IsBitlockerFeatureInstalled bool

	//
	IsDataDedupFeatureInstalled bool

	//
	IsStorageReplicaFeatureInstalled bool

	//
	Location string

	//
	Manufacturer string

	//
	MemoryDIMMs []SDDC_Memory

	//
	Model string

	//
	Name string

	//
	OSName string

	//
	OSVersion string

	//
	Processors []SDDC_Processor

	//
	SerialNumber string

	//
	Status []uint16

	//
	StatusCategory uint16

	//
	TotalNetworkUsageInBytesPerSecond float64

	//
	TotalProcessorsIdlePercentage uint64

	//
	TotalRdmaUsageInBytesPerSecond float64

	//
	Uptime string
}

func NewSDDC_ServerEx1(instance *cim.WmiInstance) (newInstance *SDDC_Server, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &SDDC_Server{
		WmiInstance: tmp,
	}
	return
}

func NewSDDC_ServerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SDDC_Server, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SDDC_Server{
		WmiInstance: tmp,
	}
	return
}

// SetAlerts sets the value of Alerts for the instance
func (instance *SDDC_Server) SetPropertyAlerts(value []SDDC_Alert) (err error) {
	return instance.SetProperty("Alerts", value)
}

// GetAlerts gets the value of Alerts for the instance
func (instance *SDDC_Server) GetPropertyAlerts() (value []SDDC_Alert, err error) {
	retValue, err := instance.GetProperty("Alerts")
	if err != nil {
		return
	}
	value, ok := retValue.([]SDDC_Alert)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBuildNumber sets the value of BuildNumber for the instance
func (instance *SDDC_Server) SetPropertyBuildNumber(value string) (err error) {
	return instance.SetProperty("BuildNumber", value)
}

// GetBuildNumber gets the value of BuildNumber for the instance
func (instance *SDDC_Server) GetPropertyBuildNumber() (value string, err error) {
	retValue, err := instance.GetProperty("BuildNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCluster sets the value of Cluster for the instance
func (instance *SDDC_Server) SetPropertyCluster(value string) (err error) {
	return instance.SetProperty("Cluster", value)
}

// GetCluster gets the value of Cluster for the instance
func (instance *SDDC_Server) GetPropertyCluster() (value string, err error) {
	retValue, err := instance.GetProperty("Cluster")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDomain sets the value of Domain for the instance
func (instance *SDDC_Server) SetPropertyDomain(value string) (err error) {
	return instance.SetProperty("Domain", value)
}

// GetDomain gets the value of Domain for the instance
func (instance *SDDC_Server) GetPropertyDomain() (value string, err error) {
	retValue, err := instance.GetProperty("Domain")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFreePhysicalMemoryInBytes sets the value of FreePhysicalMemoryInBytes for the instance
func (instance *SDDC_Server) SetPropertyFreePhysicalMemoryInBytes(value uint64) (err error) {
	return instance.SetProperty("FreePhysicalMemoryInBytes", value)
}

// GetFreePhysicalMemoryInBytes gets the value of FreePhysicalMemoryInBytes for the instance
func (instance *SDDC_Server) GetPropertyFreePhysicalMemoryInBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("FreePhysicalMemoryInBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetId sets the value of Id for the instance
func (instance *SDDC_Server) SetPropertyId(value string) (err error) {
	return instance.SetProperty("Id", value)
}

// GetId gets the value of Id for the instance
func (instance *SDDC_Server) GetPropertyId() (value string, err error) {
	retValue, err := instance.GetProperty("Id")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsBitlockerFeatureInstalled sets the value of IsBitlockerFeatureInstalled for the instance
func (instance *SDDC_Server) SetPropertyIsBitlockerFeatureInstalled(value bool) (err error) {
	return instance.SetProperty("IsBitlockerFeatureInstalled", value)
}

// GetIsBitlockerFeatureInstalled gets the value of IsBitlockerFeatureInstalled for the instance
func (instance *SDDC_Server) GetPropertyIsBitlockerFeatureInstalled() (value bool, err error) {
	retValue, err := instance.GetProperty("IsBitlockerFeatureInstalled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsDataDedupFeatureInstalled sets the value of IsDataDedupFeatureInstalled for the instance
func (instance *SDDC_Server) SetPropertyIsDataDedupFeatureInstalled(value bool) (err error) {
	return instance.SetProperty("IsDataDedupFeatureInstalled", value)
}

// GetIsDataDedupFeatureInstalled gets the value of IsDataDedupFeatureInstalled for the instance
func (instance *SDDC_Server) GetPropertyIsDataDedupFeatureInstalled() (value bool, err error) {
	retValue, err := instance.GetProperty("IsDataDedupFeatureInstalled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsStorageReplicaFeatureInstalled sets the value of IsStorageReplicaFeatureInstalled for the instance
func (instance *SDDC_Server) SetPropertyIsStorageReplicaFeatureInstalled(value bool) (err error) {
	return instance.SetProperty("IsStorageReplicaFeatureInstalled", value)
}

// GetIsStorageReplicaFeatureInstalled gets the value of IsStorageReplicaFeatureInstalled for the instance
func (instance *SDDC_Server) GetPropertyIsStorageReplicaFeatureInstalled() (value bool, err error) {
	retValue, err := instance.GetProperty("IsStorageReplicaFeatureInstalled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocation sets the value of Location for the instance
func (instance *SDDC_Server) SetPropertyLocation(value string) (err error) {
	return instance.SetProperty("Location", value)
}

// GetLocation gets the value of Location for the instance
func (instance *SDDC_Server) GetPropertyLocation() (value string, err error) {
	retValue, err := instance.GetProperty("Location")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetManufacturer sets the value of Manufacturer for the instance
func (instance *SDDC_Server) SetPropertyManufacturer(value string) (err error) {
	return instance.SetProperty("Manufacturer", value)
}

// GetManufacturer gets the value of Manufacturer for the instance
func (instance *SDDC_Server) GetPropertyManufacturer() (value string, err error) {
	retValue, err := instance.GetProperty("Manufacturer")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMemoryDIMMs sets the value of MemoryDIMMs for the instance
func (instance *SDDC_Server) SetPropertyMemoryDIMMs(value []SDDC_Memory) (err error) {
	return instance.SetProperty("MemoryDIMMs", value)
}

// GetMemoryDIMMs gets the value of MemoryDIMMs for the instance
func (instance *SDDC_Server) GetPropertyMemoryDIMMs() (value []SDDC_Memory, err error) {
	retValue, err := instance.GetProperty("MemoryDIMMs")
	if err != nil {
		return
	}
	value, ok := retValue.([]SDDC_Memory)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetModel sets the value of Model for the instance
func (instance *SDDC_Server) SetPropertyModel(value string) (err error) {
	return instance.SetProperty("Model", value)
}

// GetModel gets the value of Model for the instance
func (instance *SDDC_Server) GetPropertyModel() (value string, err error) {
	retValue, err := instance.GetProperty("Model")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *SDDC_Server) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *SDDC_Server) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOSName sets the value of OSName for the instance
func (instance *SDDC_Server) SetPropertyOSName(value string) (err error) {
	return instance.SetProperty("OSName", value)
}

// GetOSName gets the value of OSName for the instance
func (instance *SDDC_Server) GetPropertyOSName() (value string, err error) {
	retValue, err := instance.GetProperty("OSName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOSVersion sets the value of OSVersion for the instance
func (instance *SDDC_Server) SetPropertyOSVersion(value string) (err error) {
	return instance.SetProperty("OSVersion", value)
}

// GetOSVersion gets the value of OSVersion for the instance
func (instance *SDDC_Server) GetPropertyOSVersion() (value string, err error) {
	retValue, err := instance.GetProperty("OSVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProcessors sets the value of Processors for the instance
func (instance *SDDC_Server) SetPropertyProcessors(value []SDDC_Processor) (err error) {
	return instance.SetProperty("Processors", value)
}

// GetProcessors gets the value of Processors for the instance
func (instance *SDDC_Server) GetPropertyProcessors() (value []SDDC_Processor, err error) {
	retValue, err := instance.GetProperty("Processors")
	if err != nil {
		return
	}
	value, ok := retValue.([]SDDC_Processor)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSerialNumber sets the value of SerialNumber for the instance
func (instance *SDDC_Server) SetPropertySerialNumber(value string) (err error) {
	return instance.SetProperty("SerialNumber", value)
}

// GetSerialNumber gets the value of SerialNumber for the instance
func (instance *SDDC_Server) GetPropertySerialNumber() (value string, err error) {
	retValue, err := instance.GetProperty("SerialNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStatus sets the value of Status for the instance
func (instance *SDDC_Server) SetPropertyStatus(value []uint16) (err error) {
	return instance.SetProperty("Status", value)
}

// GetStatus gets the value of Status for the instance
func (instance *SDDC_Server) GetPropertyStatus() (value []uint16, err error) {
	retValue, err := instance.GetProperty("Status")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStatusCategory sets the value of StatusCategory for the instance
func (instance *SDDC_Server) SetPropertyStatusCategory(value uint16) (err error) {
	return instance.SetProperty("StatusCategory", value)
}

// GetStatusCategory gets the value of StatusCategory for the instance
func (instance *SDDC_Server) GetPropertyStatusCategory() (value uint16, err error) {
	retValue, err := instance.GetProperty("StatusCategory")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalNetworkUsageInBytesPerSecond sets the value of TotalNetworkUsageInBytesPerSecond for the instance
func (instance *SDDC_Server) SetPropertyTotalNetworkUsageInBytesPerSecond(value float64) (err error) {
	return instance.SetProperty("TotalNetworkUsageInBytesPerSecond", value)
}

// GetTotalNetworkUsageInBytesPerSecond gets the value of TotalNetworkUsageInBytesPerSecond for the instance
func (instance *SDDC_Server) GetPropertyTotalNetworkUsageInBytesPerSecond() (value float64, err error) {
	retValue, err := instance.GetProperty("TotalNetworkUsageInBytesPerSecond")
	if err != nil {
		return
	}
	value, ok := retValue.(float64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalProcessorsIdlePercentage sets the value of TotalProcessorsIdlePercentage for the instance
func (instance *SDDC_Server) SetPropertyTotalProcessorsIdlePercentage(value uint64) (err error) {
	return instance.SetProperty("TotalProcessorsIdlePercentage", value)
}

// GetTotalProcessorsIdlePercentage gets the value of TotalProcessorsIdlePercentage for the instance
func (instance *SDDC_Server) GetPropertyTotalProcessorsIdlePercentage() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalProcessorsIdlePercentage")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalRdmaUsageInBytesPerSecond sets the value of TotalRdmaUsageInBytesPerSecond for the instance
func (instance *SDDC_Server) SetPropertyTotalRdmaUsageInBytesPerSecond(value float64) (err error) {
	return instance.SetProperty("TotalRdmaUsageInBytesPerSecond", value)
}

// GetTotalRdmaUsageInBytesPerSecond gets the value of TotalRdmaUsageInBytesPerSecond for the instance
func (instance *SDDC_Server) GetPropertyTotalRdmaUsageInBytesPerSecond() (value float64, err error) {
	retValue, err := instance.GetProperty("TotalRdmaUsageInBytesPerSecond")
	if err != nil {
		return
	}
	value, ok := retValue.(float64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUptime sets the value of Uptime for the instance
func (instance *SDDC_Server) SetPropertyUptime(value string) (err error) {
	return instance.SetProperty("Uptime", value)
}

// GetUptime gets the value of Uptime for the instance
func (instance *SDDC_Server) GetPropertyUptime() (value string, err error) {
	retValue, err := instance.GetProperty("Uptime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="ServerName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *SDDC_Server) AddServer( /* IN */ ServerName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("AddServer", ServerName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *SDDC_Server) RemoveServer() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RemoveServer")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="DrainServer" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *SDDC_Server) PauseServer( /* IN */ DrainServer bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("PauseServer", DrainServer)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="FailbackType" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *SDDC_Server) ResumeServer( /* IN */ FailbackType uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ResumeServer", FailbackType)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="SeriesName" type="string "></param>
// <param name="TimeFrame" type="uint16 "></param>

// <param name="Metric" type="SDDC_Metric "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *SDDC_Server) GetMetrics( /* IN */ SeriesName string,
	/* IN */ TimeFrame uint16,
	/* OUT */ Metric SDDC_Metric) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetMetrics", SeriesName, TimeFrame)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
