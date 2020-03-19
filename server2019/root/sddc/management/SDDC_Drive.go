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

// SDDC_Drive struct
type SDDC_Drive struct {
	*cim.WmiInstance

	//
	Alerts []SDDC_Alert

	//
	AverageLatency float64

	//
	FirmwareVersion string

	//
	Id string

	//
	IsIndicationEnabled bool

	//
	Location string

	//
	Manufacturer string

	//
	Model string

	//
	PowerOnHours uint32

	//
	SerialNumber string

	//
	Server string

	//
	Size uint64

	//
	SizeUsed uint64

	//
	Status []uint16

	//
	StatusCategory uint16

	//
	StoragePool string

	//
	TemperatureInCelsius uint8

	//
	TotalIops float64

	//
	TotalThroughput float64

	//
	Type uint16

	//
	UsedFor uint16

	//
	WearPercentage uint8
}

func NewSDDC_DriveEx1(instance *cim.WmiInstance) (newInstance *SDDC_Drive, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &SDDC_Drive{
		WmiInstance: tmp,
	}
	return
}

func NewSDDC_DriveEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SDDC_Drive, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SDDC_Drive{
		WmiInstance: tmp,
	}
	return
}

// SetAlerts sets the value of Alerts for the instance
func (instance *SDDC_Drive) SetPropertyAlerts(value []SDDC_Alert) (err error) {
	return instance.SetProperty("Alerts", value)
}

// GetAlerts gets the value of Alerts for the instance
func (instance *SDDC_Drive) GetPropertyAlerts() (value []SDDC_Alert, err error) {
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

// SetAverageLatency sets the value of AverageLatency for the instance
func (instance *SDDC_Drive) SetPropertyAverageLatency(value float64) (err error) {
	return instance.SetProperty("AverageLatency", value)
}

// GetAverageLatency gets the value of AverageLatency for the instance
func (instance *SDDC_Drive) GetPropertyAverageLatency() (value float64, err error) {
	retValue, err := instance.GetProperty("AverageLatency")
	if err != nil {
		return
	}
	value, ok := retValue.(float64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFirmwareVersion sets the value of FirmwareVersion for the instance
func (instance *SDDC_Drive) SetPropertyFirmwareVersion(value string) (err error) {
	return instance.SetProperty("FirmwareVersion", value)
}

// GetFirmwareVersion gets the value of FirmwareVersion for the instance
func (instance *SDDC_Drive) GetPropertyFirmwareVersion() (value string, err error) {
	retValue, err := instance.GetProperty("FirmwareVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetId sets the value of Id for the instance
func (instance *SDDC_Drive) SetPropertyId(value string) (err error) {
	return instance.SetProperty("Id", value)
}

// GetId gets the value of Id for the instance
func (instance *SDDC_Drive) GetPropertyId() (value string, err error) {
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

// SetIsIndicationEnabled sets the value of IsIndicationEnabled for the instance
func (instance *SDDC_Drive) SetPropertyIsIndicationEnabled(value bool) (err error) {
	return instance.SetProperty("IsIndicationEnabled", value)
}

// GetIsIndicationEnabled gets the value of IsIndicationEnabled for the instance
func (instance *SDDC_Drive) GetPropertyIsIndicationEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("IsIndicationEnabled")
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
func (instance *SDDC_Drive) SetPropertyLocation(value string) (err error) {
	return instance.SetProperty("Location", value)
}

// GetLocation gets the value of Location for the instance
func (instance *SDDC_Drive) GetPropertyLocation() (value string, err error) {
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
func (instance *SDDC_Drive) SetPropertyManufacturer(value string) (err error) {
	return instance.SetProperty("Manufacturer", value)
}

// GetManufacturer gets the value of Manufacturer for the instance
func (instance *SDDC_Drive) GetPropertyManufacturer() (value string, err error) {
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

// SetModel sets the value of Model for the instance
func (instance *SDDC_Drive) SetPropertyModel(value string) (err error) {
	return instance.SetProperty("Model", value)
}

// GetModel gets the value of Model for the instance
func (instance *SDDC_Drive) GetPropertyModel() (value string, err error) {
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

// SetPowerOnHours sets the value of PowerOnHours for the instance
func (instance *SDDC_Drive) SetPropertyPowerOnHours(value uint32) (err error) {
	return instance.SetProperty("PowerOnHours", value)
}

// GetPowerOnHours gets the value of PowerOnHours for the instance
func (instance *SDDC_Drive) GetPropertyPowerOnHours() (value uint32, err error) {
	retValue, err := instance.GetProperty("PowerOnHours")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSerialNumber sets the value of SerialNumber for the instance
func (instance *SDDC_Drive) SetPropertySerialNumber(value string) (err error) {
	return instance.SetProperty("SerialNumber", value)
}

// GetSerialNumber gets the value of SerialNumber for the instance
func (instance *SDDC_Drive) GetPropertySerialNumber() (value string, err error) {
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

// SetServer sets the value of Server for the instance
func (instance *SDDC_Drive) SetPropertyServer(value string) (err error) {
	return instance.SetProperty("Server", value)
}

// GetServer gets the value of Server for the instance
func (instance *SDDC_Drive) GetPropertyServer() (value string, err error) {
	retValue, err := instance.GetProperty("Server")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSize sets the value of Size for the instance
func (instance *SDDC_Drive) SetPropertySize(value uint64) (err error) {
	return instance.SetProperty("Size", value)
}

// GetSize gets the value of Size for the instance
func (instance *SDDC_Drive) GetPropertySize() (value uint64, err error) {
	retValue, err := instance.GetProperty("Size")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSizeUsed sets the value of SizeUsed for the instance
func (instance *SDDC_Drive) SetPropertySizeUsed(value uint64) (err error) {
	return instance.SetProperty("SizeUsed", value)
}

// GetSizeUsed gets the value of SizeUsed for the instance
func (instance *SDDC_Drive) GetPropertySizeUsed() (value uint64, err error) {
	retValue, err := instance.GetProperty("SizeUsed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStatus sets the value of Status for the instance
func (instance *SDDC_Drive) SetPropertyStatus(value []uint16) (err error) {
	return instance.SetProperty("Status", value)
}

// GetStatus gets the value of Status for the instance
func (instance *SDDC_Drive) GetPropertyStatus() (value []uint16, err error) {
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
func (instance *SDDC_Drive) SetPropertyStatusCategory(value uint16) (err error) {
	return instance.SetProperty("StatusCategory", value)
}

// GetStatusCategory gets the value of StatusCategory for the instance
func (instance *SDDC_Drive) GetPropertyStatusCategory() (value uint16, err error) {
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

// SetStoragePool sets the value of StoragePool for the instance
func (instance *SDDC_Drive) SetPropertyStoragePool(value string) (err error) {
	return instance.SetProperty("StoragePool", value)
}

// GetStoragePool gets the value of StoragePool for the instance
func (instance *SDDC_Drive) GetPropertyStoragePool() (value string, err error) {
	retValue, err := instance.GetProperty("StoragePool")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTemperatureInCelsius sets the value of TemperatureInCelsius for the instance
func (instance *SDDC_Drive) SetPropertyTemperatureInCelsius(value uint8) (err error) {
	return instance.SetProperty("TemperatureInCelsius", value)
}

// GetTemperatureInCelsius gets the value of TemperatureInCelsius for the instance
func (instance *SDDC_Drive) GetPropertyTemperatureInCelsius() (value uint8, err error) {
	retValue, err := instance.GetProperty("TemperatureInCelsius")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalIops sets the value of TotalIops for the instance
func (instance *SDDC_Drive) SetPropertyTotalIops(value float64) (err error) {
	return instance.SetProperty("TotalIops", value)
}

// GetTotalIops gets the value of TotalIops for the instance
func (instance *SDDC_Drive) GetPropertyTotalIops() (value float64, err error) {
	retValue, err := instance.GetProperty("TotalIops")
	if err != nil {
		return
	}
	value, ok := retValue.(float64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalThroughput sets the value of TotalThroughput for the instance
func (instance *SDDC_Drive) SetPropertyTotalThroughput(value float64) (err error) {
	return instance.SetProperty("TotalThroughput", value)
}

// GetTotalThroughput gets the value of TotalThroughput for the instance
func (instance *SDDC_Drive) GetPropertyTotalThroughput() (value float64, err error) {
	retValue, err := instance.GetProperty("TotalThroughput")
	if err != nil {
		return
	}
	value, ok := retValue.(float64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetType sets the value of Type for the instance
func (instance *SDDC_Drive) SetPropertyType(value uint16) (err error) {
	return instance.SetProperty("Type", value)
}

// GetType gets the value of Type for the instance
func (instance *SDDC_Drive) GetPropertyType() (value uint16, err error) {
	retValue, err := instance.GetProperty("Type")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUsedFor sets the value of UsedFor for the instance
func (instance *SDDC_Drive) SetPropertyUsedFor(value uint16) (err error) {
	return instance.SetProperty("UsedFor", value)
}

// GetUsedFor gets the value of UsedFor for the instance
func (instance *SDDC_Drive) GetPropertyUsedFor() (value uint16, err error) {
	retValue, err := instance.GetProperty("UsedFor")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWearPercentage sets the value of WearPercentage for the instance
func (instance *SDDC_Drive) SetPropertyWearPercentage(value uint8) (err error) {
	return instance.SetProperty("WearPercentage", value)
}

// GetWearPercentage gets the value of WearPercentage for the instance
func (instance *SDDC_Drive) GetPropertyWearPercentage() (value uint8, err error) {
	retValue, err := instance.GetProperty("WearPercentage")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *SDDC_Drive) Retire() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Retire")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *SDDC_Drive) Unretire() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Unretire")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ImagePath" type="string "></param>
// <param name="SlotNumber" type="uint16 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *SDDC_Drive) UpdateFirmware( /* IN */ ImagePath string,
	/* IN */ SlotNumber uint16) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("UpdateFirmware", ImagePath, SlotNumber)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *SDDC_Drive) TurnOnLight() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("TurnOnLight")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *SDDC_Drive) TurnOffLight() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("TurnOffLight")
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
func (instance *SDDC_Drive) GetMetrics( /* IN */ SeriesName string,
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
