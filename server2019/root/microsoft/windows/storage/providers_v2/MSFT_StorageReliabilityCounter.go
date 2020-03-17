// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

// MSFT_StorageReliabilityCounter struct
type MSFT_StorageReliabilityCounter struct {
	MSFT_StorageObject

	// DeviceId identifies the associated storage device. When associated with an MSFT_PhysicalDisk, it will be the same as its DeviceId field. When associated with an MSFT_Disk, it will be the same as its Number field.
	DeviceId string

	//
	FlushLatencyMax uint64

	// Number of load-unload cycles performed by the storage device.
	LoadUnloadCycleCount uint32

	// Maximum number of load-unload cycles within which the storage device is capable of normal operation.
	LoadUnloadCycleCountMax uint32

	// Year and week of storage device manufacture.
	ManufactureDate string

	// Length of time, in hours, the storage device has been powered on since manufacture.
	PowerOnHours uint32

	// Read errors corrected by the storage device.
	ReadErrorsCorrected uint64

	// Total read errors encountered by the storage device.
	ReadErrorsTotal uint64

	// Read errors not corrected by the storage device.
	ReadErrorsUncorrected uint64

	//
	ReadLatencyMax uint64

	// Number of start-stop cycles performed by the storage device.
	StartStopCycleCount uint32

	// Maximum number of start-stop cycles within which the storage device is capable of normal operation.
	StartStopCycleCountMax uint32

	// The current temperature of the storage device in Celsius.
	Temperature uint8

	// The maximum temperature in Celsius at which the storage device is capable of normal operation.
	TemperatureMax uint8

	// Storage device wear indicator, in percentage. At 100 percent, the estimated wear limit will have been reached.
	Wear uint8

	// Write errors corrected by the storage device.
	WriteErrorsCorrected uint64

	// Total write errors encountered by the storage device.
	WriteErrorsTotal uint64

	// Write errors not corrected by the storage device.
	WriteErrorsUncorrected uint64

	//
	WriteLatencyMax uint64
}

// SetDeviceId sets the value of DeviceId for the instance
func (instance *MSFT_StorageReliabilityCounter) SetPropertyDeviceId(value string) (err error) {
	return instance.SetProperty("DeviceId", value)
}

// GetDeviceId gets the value of DeviceId for the instance
func (instance *MSFT_StorageReliabilityCounter) GetPropertyDeviceId() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFlushLatencyMax sets the value of FlushLatencyMax for the instance
func (instance *MSFT_StorageReliabilityCounter) SetPropertyFlushLatencyMax(value uint64) (err error) {
	return instance.SetProperty("FlushLatencyMax", value)
}

// GetFlushLatencyMax gets the value of FlushLatencyMax for the instance
func (instance *MSFT_StorageReliabilityCounter) GetPropertyFlushLatencyMax() (value uint64, err error) {
	retValue, err := instance.GetProperty("FlushLatencyMax")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLoadUnloadCycleCount sets the value of LoadUnloadCycleCount for the instance
func (instance *MSFT_StorageReliabilityCounter) SetPropertyLoadUnloadCycleCount(value uint32) (err error) {
	return instance.SetProperty("LoadUnloadCycleCount", value)
}

// GetLoadUnloadCycleCount gets the value of LoadUnloadCycleCount for the instance
func (instance *MSFT_StorageReliabilityCounter) GetPropertyLoadUnloadCycleCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("LoadUnloadCycleCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLoadUnloadCycleCountMax sets the value of LoadUnloadCycleCountMax for the instance
func (instance *MSFT_StorageReliabilityCounter) SetPropertyLoadUnloadCycleCountMax(value uint32) (err error) {
	return instance.SetProperty("LoadUnloadCycleCountMax", value)
}

// GetLoadUnloadCycleCountMax gets the value of LoadUnloadCycleCountMax for the instance
func (instance *MSFT_StorageReliabilityCounter) GetPropertyLoadUnloadCycleCountMax() (value uint32, err error) {
	retValue, err := instance.GetProperty("LoadUnloadCycleCountMax")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetManufactureDate sets the value of ManufactureDate for the instance
func (instance *MSFT_StorageReliabilityCounter) SetPropertyManufactureDate(value string) (err error) {
	return instance.SetProperty("ManufactureDate", value)
}

// GetManufactureDate gets the value of ManufactureDate for the instance
func (instance *MSFT_StorageReliabilityCounter) GetPropertyManufactureDate() (value string, err error) {
	retValue, err := instance.GetProperty("ManufactureDate")
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
func (instance *MSFT_StorageReliabilityCounter) SetPropertyPowerOnHours(value uint32) (err error) {
	return instance.SetProperty("PowerOnHours", value)
}

// GetPowerOnHours gets the value of PowerOnHours for the instance
func (instance *MSFT_StorageReliabilityCounter) GetPropertyPowerOnHours() (value uint32, err error) {
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

// SetReadErrorsCorrected sets the value of ReadErrorsCorrected for the instance
func (instance *MSFT_StorageReliabilityCounter) SetPropertyReadErrorsCorrected(value uint64) (err error) {
	return instance.SetProperty("ReadErrorsCorrected", value)
}

// GetReadErrorsCorrected gets the value of ReadErrorsCorrected for the instance
func (instance *MSFT_StorageReliabilityCounter) GetPropertyReadErrorsCorrected() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadErrorsCorrected")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReadErrorsTotal sets the value of ReadErrorsTotal for the instance
func (instance *MSFT_StorageReliabilityCounter) SetPropertyReadErrorsTotal(value uint64) (err error) {
	return instance.SetProperty("ReadErrorsTotal", value)
}

// GetReadErrorsTotal gets the value of ReadErrorsTotal for the instance
func (instance *MSFT_StorageReliabilityCounter) GetPropertyReadErrorsTotal() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadErrorsTotal")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReadErrorsUncorrected sets the value of ReadErrorsUncorrected for the instance
func (instance *MSFT_StorageReliabilityCounter) SetPropertyReadErrorsUncorrected(value uint64) (err error) {
	return instance.SetProperty("ReadErrorsUncorrected", value)
}

// GetReadErrorsUncorrected gets the value of ReadErrorsUncorrected for the instance
func (instance *MSFT_StorageReliabilityCounter) GetPropertyReadErrorsUncorrected() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadErrorsUncorrected")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReadLatencyMax sets the value of ReadLatencyMax for the instance
func (instance *MSFT_StorageReliabilityCounter) SetPropertyReadLatencyMax(value uint64) (err error) {
	return instance.SetProperty("ReadLatencyMax", value)
}

// GetReadLatencyMax gets the value of ReadLatencyMax for the instance
func (instance *MSFT_StorageReliabilityCounter) GetPropertyReadLatencyMax() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadLatencyMax")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStartStopCycleCount sets the value of StartStopCycleCount for the instance
func (instance *MSFT_StorageReliabilityCounter) SetPropertyStartStopCycleCount(value uint32) (err error) {
	return instance.SetProperty("StartStopCycleCount", value)
}

// GetStartStopCycleCount gets the value of StartStopCycleCount for the instance
func (instance *MSFT_StorageReliabilityCounter) GetPropertyStartStopCycleCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("StartStopCycleCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStartStopCycleCountMax sets the value of StartStopCycleCountMax for the instance
func (instance *MSFT_StorageReliabilityCounter) SetPropertyStartStopCycleCountMax(value uint32) (err error) {
	return instance.SetProperty("StartStopCycleCountMax", value)
}

// GetStartStopCycleCountMax gets the value of StartStopCycleCountMax for the instance
func (instance *MSFT_StorageReliabilityCounter) GetPropertyStartStopCycleCountMax() (value uint32, err error) {
	retValue, err := instance.GetProperty("StartStopCycleCountMax")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTemperature sets the value of Temperature for the instance
func (instance *MSFT_StorageReliabilityCounter) SetPropertyTemperature(value uint8) (err error) {
	return instance.SetProperty("Temperature", value)
}

// GetTemperature gets the value of Temperature for the instance
func (instance *MSFT_StorageReliabilityCounter) GetPropertyTemperature() (value uint8, err error) {
	retValue, err := instance.GetProperty("Temperature")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTemperatureMax sets the value of TemperatureMax for the instance
func (instance *MSFT_StorageReliabilityCounter) SetPropertyTemperatureMax(value uint8) (err error) {
	return instance.SetProperty("TemperatureMax", value)
}

// GetTemperatureMax gets the value of TemperatureMax for the instance
func (instance *MSFT_StorageReliabilityCounter) GetPropertyTemperatureMax() (value uint8, err error) {
	retValue, err := instance.GetProperty("TemperatureMax")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWear sets the value of Wear for the instance
func (instance *MSFT_StorageReliabilityCounter) SetPropertyWear(value uint8) (err error) {
	return instance.SetProperty("Wear", value)
}

// GetWear gets the value of Wear for the instance
func (instance *MSFT_StorageReliabilityCounter) GetPropertyWear() (value uint8, err error) {
	retValue, err := instance.GetProperty("Wear")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWriteErrorsCorrected sets the value of WriteErrorsCorrected for the instance
func (instance *MSFT_StorageReliabilityCounter) SetPropertyWriteErrorsCorrected(value uint64) (err error) {
	return instance.SetProperty("WriteErrorsCorrected", value)
}

// GetWriteErrorsCorrected gets the value of WriteErrorsCorrected for the instance
func (instance *MSFT_StorageReliabilityCounter) GetPropertyWriteErrorsCorrected() (value uint64, err error) {
	retValue, err := instance.GetProperty("WriteErrorsCorrected")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWriteErrorsTotal sets the value of WriteErrorsTotal for the instance
func (instance *MSFT_StorageReliabilityCounter) SetPropertyWriteErrorsTotal(value uint64) (err error) {
	return instance.SetProperty("WriteErrorsTotal", value)
}

// GetWriteErrorsTotal gets the value of WriteErrorsTotal for the instance
func (instance *MSFT_StorageReliabilityCounter) GetPropertyWriteErrorsTotal() (value uint64, err error) {
	retValue, err := instance.GetProperty("WriteErrorsTotal")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWriteErrorsUncorrected sets the value of WriteErrorsUncorrected for the instance
func (instance *MSFT_StorageReliabilityCounter) SetPropertyWriteErrorsUncorrected(value uint64) (err error) {
	return instance.SetProperty("WriteErrorsUncorrected", value)
}

// GetWriteErrorsUncorrected gets the value of WriteErrorsUncorrected for the instance
func (instance *MSFT_StorageReliabilityCounter) GetPropertyWriteErrorsUncorrected() (value uint64, err error) {
	retValue, err := instance.GetProperty("WriteErrorsUncorrected")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWriteLatencyMax sets the value of WriteLatencyMax for the instance
func (instance *MSFT_StorageReliabilityCounter) SetPropertyWriteLatencyMax(value uint64) (err error) {
	return instance.SetProperty("WriteLatencyMax", value)
}

// GetWriteLatencyMax gets the value of WriteLatencyMax for the instance
func (instance *MSFT_StorageReliabilityCounter) GetPropertyWriteLatencyMax() (value uint64, err error) {
	retValue, err := instance.GetProperty("WriteLatencyMax")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
