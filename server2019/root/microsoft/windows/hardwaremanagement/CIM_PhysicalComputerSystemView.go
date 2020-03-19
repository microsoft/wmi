// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.HardwareManagement
//////////////////////////////////////////////
package hardwaremanagement

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_PhysicalComputerSystemView struct
type CIM_PhysicalComputerSystemView struct {
	*CIM_View

	//
	CurrentBIOSBuildNumber uint16

	//
	CurrentBIOSMajorVersion uint16

	//
	CurrentBIOSMinorVersion uint16

	//
	CurrentBIOSRevisionNumber uint16

	//
	CurrentBIOSVersionString string

	//
	CurrentManagementFirmwareBuildNumber uint16

	//
	CurrentManagementFirmwareElementName string

	//
	CurrentManagementFirmwareMajorVersion uint16

	//
	CurrentManagementFirmwareMinorVersion uint16

	//
	CurrentManagementFirmwareRevisionNumber uint16

	//
	CurrentManagementFirmwareVersionString string

	//
	Dedicated []uint16

	//
	EnabledState uint16

	//
	FRUInfoSupported bool

	//
	HealthState uint16

	//
	IdentifyingDescriptions []string

	//
	LogCurrentNumberOfRecords []uint64

	//
	LogInstanceID []string

	//
	LogMaxNumberOfRecords []uint64

	//
	LogOverwritePolicy []uint16

	//
	LogState []uint16

	//
	Manufacturer string

	//
	MemoryBlockSize uint64

	//
	MemoryConsumableBlocks uint64

	//
	MemoryNumberOfBlocks uint64

	//
	Model string

	//
	NumberOfProcessorCores uint16

	//
	NumberOfProcessors uint16

	//
	NumberOfProcessorThreads uint16

	//
	NumericSensorBaseUnits []uint16

	//
	NumericSensorContext []string

	//
	NumericSensorCurrentReading []int32

	//
	NumericSensorCurrentState []string

	//
	NumericSensorElementName []string

	//
	NumericSensorEnabledState []uint16

	//
	NumericSensorHealthState []uint16

	//
	NumericSensorLowerThresholdCritical []int32

	//
	NumericSensorLowerThresholdFatal []int32

	//
	NumericSensorLowerThresholdNonCritical []int32

	//
	NumericSensorOtherSensorTypeDescription []string

	//
	NumericSensorPrimaryStatus []uint16

	//
	NumericSensorRateUnits []uint16

	//
	NumericSensorSensorType []uint16

	//
	NumericSensorUnitModifier []int32

	//
	NumericSensorUpperThresholdCritical []int32

	//
	NumericSensorUpperThresholdFatal []int32

	//
	NumericSensorUpperThresholdNonCritical []int32

	//
	OneTimeBootSource uint8

	//
	OperationalStatus []uint16

	//
	OSEnabledState uint16

	//
	OSType uint16

	//
	OSVersion string

	//
	OtherDedicatedDescriptions []string

	//
	OtherIdentifyingInfo []string

	//
	PartNumber string

	//
	PersistentBootConfigOrder []uint8

	//
	PowerAllocationLimit uint64

	//
	PowerUtilizationMode uint16

	//
	PowerUtilizationModesSupported []uint16

	//
	ProcessorCurrentClockSpeed uint32

	//
	ProcessorFamily uint16

	//
	ProcessorMaxClockSpeed uint32

	//
	RequestedState uint16

	//
	SerialNumber string

	//
	SKU string

	//
	StructuredBootString []string

	//
	Tag string

	//
	Version string
}

func NewCIM_PhysicalComputerSystemViewEx1(instance *cim.WmiInstance) (newInstance *CIM_PhysicalComputerSystemView, err error) {
	tmp, err := NewCIM_ViewEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_PhysicalComputerSystemView{
		CIM_View: tmp,
	}
	return
}

func NewCIM_PhysicalComputerSystemViewEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_PhysicalComputerSystemView, err error) {
	tmp, err := NewCIM_ViewEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_PhysicalComputerSystemView{
		CIM_View: tmp,
	}
	return
}

// SetCurrentBIOSBuildNumber sets the value of CurrentBIOSBuildNumber for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyCurrentBIOSBuildNumber(value uint16) (err error) {
	return instance.SetProperty("CurrentBIOSBuildNumber", value)
}

// GetCurrentBIOSBuildNumber gets the value of CurrentBIOSBuildNumber for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyCurrentBIOSBuildNumber() (value uint16, err error) {
	retValue, err := instance.GetProperty("CurrentBIOSBuildNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentBIOSMajorVersion sets the value of CurrentBIOSMajorVersion for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyCurrentBIOSMajorVersion(value uint16) (err error) {
	return instance.SetProperty("CurrentBIOSMajorVersion", value)
}

// GetCurrentBIOSMajorVersion gets the value of CurrentBIOSMajorVersion for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyCurrentBIOSMajorVersion() (value uint16, err error) {
	retValue, err := instance.GetProperty("CurrentBIOSMajorVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentBIOSMinorVersion sets the value of CurrentBIOSMinorVersion for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyCurrentBIOSMinorVersion(value uint16) (err error) {
	return instance.SetProperty("CurrentBIOSMinorVersion", value)
}

// GetCurrentBIOSMinorVersion gets the value of CurrentBIOSMinorVersion for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyCurrentBIOSMinorVersion() (value uint16, err error) {
	retValue, err := instance.GetProperty("CurrentBIOSMinorVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentBIOSRevisionNumber sets the value of CurrentBIOSRevisionNumber for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyCurrentBIOSRevisionNumber(value uint16) (err error) {
	return instance.SetProperty("CurrentBIOSRevisionNumber", value)
}

// GetCurrentBIOSRevisionNumber gets the value of CurrentBIOSRevisionNumber for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyCurrentBIOSRevisionNumber() (value uint16, err error) {
	retValue, err := instance.GetProperty("CurrentBIOSRevisionNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentBIOSVersionString sets the value of CurrentBIOSVersionString for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyCurrentBIOSVersionString(value string) (err error) {
	return instance.SetProperty("CurrentBIOSVersionString", value)
}

// GetCurrentBIOSVersionString gets the value of CurrentBIOSVersionString for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyCurrentBIOSVersionString() (value string, err error) {
	retValue, err := instance.GetProperty("CurrentBIOSVersionString")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentManagementFirmwareBuildNumber sets the value of CurrentManagementFirmwareBuildNumber for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyCurrentManagementFirmwareBuildNumber(value uint16) (err error) {
	return instance.SetProperty("CurrentManagementFirmwareBuildNumber", value)
}

// GetCurrentManagementFirmwareBuildNumber gets the value of CurrentManagementFirmwareBuildNumber for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyCurrentManagementFirmwareBuildNumber() (value uint16, err error) {
	retValue, err := instance.GetProperty("CurrentManagementFirmwareBuildNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentManagementFirmwareElementName sets the value of CurrentManagementFirmwareElementName for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyCurrentManagementFirmwareElementName(value string) (err error) {
	return instance.SetProperty("CurrentManagementFirmwareElementName", value)
}

// GetCurrentManagementFirmwareElementName gets the value of CurrentManagementFirmwareElementName for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyCurrentManagementFirmwareElementName() (value string, err error) {
	retValue, err := instance.GetProperty("CurrentManagementFirmwareElementName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentManagementFirmwareMajorVersion sets the value of CurrentManagementFirmwareMajorVersion for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyCurrentManagementFirmwareMajorVersion(value uint16) (err error) {
	return instance.SetProperty("CurrentManagementFirmwareMajorVersion", value)
}

// GetCurrentManagementFirmwareMajorVersion gets the value of CurrentManagementFirmwareMajorVersion for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyCurrentManagementFirmwareMajorVersion() (value uint16, err error) {
	retValue, err := instance.GetProperty("CurrentManagementFirmwareMajorVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentManagementFirmwareMinorVersion sets the value of CurrentManagementFirmwareMinorVersion for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyCurrentManagementFirmwareMinorVersion(value uint16) (err error) {
	return instance.SetProperty("CurrentManagementFirmwareMinorVersion", value)
}

// GetCurrentManagementFirmwareMinorVersion gets the value of CurrentManagementFirmwareMinorVersion for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyCurrentManagementFirmwareMinorVersion() (value uint16, err error) {
	retValue, err := instance.GetProperty("CurrentManagementFirmwareMinorVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentManagementFirmwareRevisionNumber sets the value of CurrentManagementFirmwareRevisionNumber for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyCurrentManagementFirmwareRevisionNumber(value uint16) (err error) {
	return instance.SetProperty("CurrentManagementFirmwareRevisionNumber", value)
}

// GetCurrentManagementFirmwareRevisionNumber gets the value of CurrentManagementFirmwareRevisionNumber for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyCurrentManagementFirmwareRevisionNumber() (value uint16, err error) {
	retValue, err := instance.GetProperty("CurrentManagementFirmwareRevisionNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentManagementFirmwareVersionString sets the value of CurrentManagementFirmwareVersionString for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyCurrentManagementFirmwareVersionString(value string) (err error) {
	return instance.SetProperty("CurrentManagementFirmwareVersionString", value)
}

// GetCurrentManagementFirmwareVersionString gets the value of CurrentManagementFirmwareVersionString for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyCurrentManagementFirmwareVersionString() (value string, err error) {
	retValue, err := instance.GetProperty("CurrentManagementFirmwareVersionString")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDedicated sets the value of Dedicated for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyDedicated(value []uint16) (err error) {
	return instance.SetProperty("Dedicated", value)
}

// GetDedicated gets the value of Dedicated for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyDedicated() (value []uint16, err error) {
	retValue, err := instance.GetProperty("Dedicated")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnabledState sets the value of EnabledState for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyEnabledState(value uint16) (err error) {
	return instance.SetProperty("EnabledState", value)
}

// GetEnabledState gets the value of EnabledState for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyEnabledState() (value uint16, err error) {
	retValue, err := instance.GetProperty("EnabledState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFRUInfoSupported sets the value of FRUInfoSupported for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyFRUInfoSupported(value bool) (err error) {
	return instance.SetProperty("FRUInfoSupported", value)
}

// GetFRUInfoSupported gets the value of FRUInfoSupported for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyFRUInfoSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("FRUInfoSupported")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHealthState sets the value of HealthState for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyHealthState(value uint16) (err error) {
	return instance.SetProperty("HealthState", value)
}

// GetHealthState gets the value of HealthState for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyHealthState() (value uint16, err error) {
	retValue, err := instance.GetProperty("HealthState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIdentifyingDescriptions sets the value of IdentifyingDescriptions for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyIdentifyingDescriptions(value []string) (err error) {
	return instance.SetProperty("IdentifyingDescriptions", value)
}

// GetIdentifyingDescriptions gets the value of IdentifyingDescriptions for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyIdentifyingDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("IdentifyingDescriptions")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogCurrentNumberOfRecords sets the value of LogCurrentNumberOfRecords for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyLogCurrentNumberOfRecords(value []uint64) (err error) {
	return instance.SetProperty("LogCurrentNumberOfRecords", value)
}

// GetLogCurrentNumberOfRecords gets the value of LogCurrentNumberOfRecords for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyLogCurrentNumberOfRecords() (value []uint64, err error) {
	retValue, err := instance.GetProperty("LogCurrentNumberOfRecords")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogInstanceID sets the value of LogInstanceID for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyLogInstanceID(value []string) (err error) {
	return instance.SetProperty("LogInstanceID", value)
}

// GetLogInstanceID gets the value of LogInstanceID for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyLogInstanceID() (value []string, err error) {
	retValue, err := instance.GetProperty("LogInstanceID")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogMaxNumberOfRecords sets the value of LogMaxNumberOfRecords for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyLogMaxNumberOfRecords(value []uint64) (err error) {
	return instance.SetProperty("LogMaxNumberOfRecords", value)
}

// GetLogMaxNumberOfRecords gets the value of LogMaxNumberOfRecords for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyLogMaxNumberOfRecords() (value []uint64, err error) {
	retValue, err := instance.GetProperty("LogMaxNumberOfRecords")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogOverwritePolicy sets the value of LogOverwritePolicy for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyLogOverwritePolicy(value []uint16) (err error) {
	return instance.SetProperty("LogOverwritePolicy", value)
}

// GetLogOverwritePolicy gets the value of LogOverwritePolicy for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyLogOverwritePolicy() (value []uint16, err error) {
	retValue, err := instance.GetProperty("LogOverwritePolicy")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogState sets the value of LogState for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyLogState(value []uint16) (err error) {
	return instance.SetProperty("LogState", value)
}

// GetLogState gets the value of LogState for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyLogState() (value []uint16, err error) {
	retValue, err := instance.GetProperty("LogState")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetManufacturer sets the value of Manufacturer for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyManufacturer(value string) (err error) {
	return instance.SetProperty("Manufacturer", value)
}

// GetManufacturer gets the value of Manufacturer for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyManufacturer() (value string, err error) {
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

// SetMemoryBlockSize sets the value of MemoryBlockSize for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyMemoryBlockSize(value uint64) (err error) {
	return instance.SetProperty("MemoryBlockSize", value)
}

// GetMemoryBlockSize gets the value of MemoryBlockSize for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyMemoryBlockSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("MemoryBlockSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMemoryConsumableBlocks sets the value of MemoryConsumableBlocks for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyMemoryConsumableBlocks(value uint64) (err error) {
	return instance.SetProperty("MemoryConsumableBlocks", value)
}

// GetMemoryConsumableBlocks gets the value of MemoryConsumableBlocks for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyMemoryConsumableBlocks() (value uint64, err error) {
	retValue, err := instance.GetProperty("MemoryConsumableBlocks")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMemoryNumberOfBlocks sets the value of MemoryNumberOfBlocks for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyMemoryNumberOfBlocks(value uint64) (err error) {
	return instance.SetProperty("MemoryNumberOfBlocks", value)
}

// GetMemoryNumberOfBlocks gets the value of MemoryNumberOfBlocks for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyMemoryNumberOfBlocks() (value uint64, err error) {
	retValue, err := instance.GetProperty("MemoryNumberOfBlocks")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetModel sets the value of Model for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyModel(value string) (err error) {
	return instance.SetProperty("Model", value)
}

// GetModel gets the value of Model for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyModel() (value string, err error) {
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

// SetNumberOfProcessorCores sets the value of NumberOfProcessorCores for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyNumberOfProcessorCores(value uint16) (err error) {
	return instance.SetProperty("NumberOfProcessorCores", value)
}

// GetNumberOfProcessorCores gets the value of NumberOfProcessorCores for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyNumberOfProcessorCores() (value uint16, err error) {
	retValue, err := instance.GetProperty("NumberOfProcessorCores")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfProcessors sets the value of NumberOfProcessors for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyNumberOfProcessors(value uint16) (err error) {
	return instance.SetProperty("NumberOfProcessors", value)
}

// GetNumberOfProcessors gets the value of NumberOfProcessors for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyNumberOfProcessors() (value uint16, err error) {
	retValue, err := instance.GetProperty("NumberOfProcessors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfProcessorThreads sets the value of NumberOfProcessorThreads for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyNumberOfProcessorThreads(value uint16) (err error) {
	return instance.SetProperty("NumberOfProcessorThreads", value)
}

// GetNumberOfProcessorThreads gets the value of NumberOfProcessorThreads for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyNumberOfProcessorThreads() (value uint16, err error) {
	retValue, err := instance.GetProperty("NumberOfProcessorThreads")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumericSensorBaseUnits sets the value of NumericSensorBaseUnits for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyNumericSensorBaseUnits(value []uint16) (err error) {
	return instance.SetProperty("NumericSensorBaseUnits", value)
}

// GetNumericSensorBaseUnits gets the value of NumericSensorBaseUnits for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyNumericSensorBaseUnits() (value []uint16, err error) {
	retValue, err := instance.GetProperty("NumericSensorBaseUnits")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumericSensorContext sets the value of NumericSensorContext for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyNumericSensorContext(value []string) (err error) {
	return instance.SetProperty("NumericSensorContext", value)
}

// GetNumericSensorContext gets the value of NumericSensorContext for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyNumericSensorContext() (value []string, err error) {
	retValue, err := instance.GetProperty("NumericSensorContext")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumericSensorCurrentReading sets the value of NumericSensorCurrentReading for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyNumericSensorCurrentReading(value []int32) (err error) {
	return instance.SetProperty("NumericSensorCurrentReading", value)
}

// GetNumericSensorCurrentReading gets the value of NumericSensorCurrentReading for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyNumericSensorCurrentReading() (value []int32, err error) {
	retValue, err := instance.GetProperty("NumericSensorCurrentReading")
	if err != nil {
		return
	}
	value, ok := retValue.([]int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumericSensorCurrentState sets the value of NumericSensorCurrentState for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyNumericSensorCurrentState(value []string) (err error) {
	return instance.SetProperty("NumericSensorCurrentState", value)
}

// GetNumericSensorCurrentState gets the value of NumericSensorCurrentState for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyNumericSensorCurrentState() (value []string, err error) {
	retValue, err := instance.GetProperty("NumericSensorCurrentState")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumericSensorElementName sets the value of NumericSensorElementName for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyNumericSensorElementName(value []string) (err error) {
	return instance.SetProperty("NumericSensorElementName", value)
}

// GetNumericSensorElementName gets the value of NumericSensorElementName for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyNumericSensorElementName() (value []string, err error) {
	retValue, err := instance.GetProperty("NumericSensorElementName")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumericSensorEnabledState sets the value of NumericSensorEnabledState for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyNumericSensorEnabledState(value []uint16) (err error) {
	return instance.SetProperty("NumericSensorEnabledState", value)
}

// GetNumericSensorEnabledState gets the value of NumericSensorEnabledState for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyNumericSensorEnabledState() (value []uint16, err error) {
	retValue, err := instance.GetProperty("NumericSensorEnabledState")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumericSensorHealthState sets the value of NumericSensorHealthState for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyNumericSensorHealthState(value []uint16) (err error) {
	return instance.SetProperty("NumericSensorHealthState", value)
}

// GetNumericSensorHealthState gets the value of NumericSensorHealthState for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyNumericSensorHealthState() (value []uint16, err error) {
	retValue, err := instance.GetProperty("NumericSensorHealthState")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumericSensorLowerThresholdCritical sets the value of NumericSensorLowerThresholdCritical for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyNumericSensorLowerThresholdCritical(value []int32) (err error) {
	return instance.SetProperty("NumericSensorLowerThresholdCritical", value)
}

// GetNumericSensorLowerThresholdCritical gets the value of NumericSensorLowerThresholdCritical for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyNumericSensorLowerThresholdCritical() (value []int32, err error) {
	retValue, err := instance.GetProperty("NumericSensorLowerThresholdCritical")
	if err != nil {
		return
	}
	value, ok := retValue.([]int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumericSensorLowerThresholdFatal sets the value of NumericSensorLowerThresholdFatal for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyNumericSensorLowerThresholdFatal(value []int32) (err error) {
	return instance.SetProperty("NumericSensorLowerThresholdFatal", value)
}

// GetNumericSensorLowerThresholdFatal gets the value of NumericSensorLowerThresholdFatal for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyNumericSensorLowerThresholdFatal() (value []int32, err error) {
	retValue, err := instance.GetProperty("NumericSensorLowerThresholdFatal")
	if err != nil {
		return
	}
	value, ok := retValue.([]int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumericSensorLowerThresholdNonCritical sets the value of NumericSensorLowerThresholdNonCritical for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyNumericSensorLowerThresholdNonCritical(value []int32) (err error) {
	return instance.SetProperty("NumericSensorLowerThresholdNonCritical", value)
}

// GetNumericSensorLowerThresholdNonCritical gets the value of NumericSensorLowerThresholdNonCritical for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyNumericSensorLowerThresholdNonCritical() (value []int32, err error) {
	retValue, err := instance.GetProperty("NumericSensorLowerThresholdNonCritical")
	if err != nil {
		return
	}
	value, ok := retValue.([]int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumericSensorOtherSensorTypeDescription sets the value of NumericSensorOtherSensorTypeDescription for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyNumericSensorOtherSensorTypeDescription(value []string) (err error) {
	return instance.SetProperty("NumericSensorOtherSensorTypeDescription", value)
}

// GetNumericSensorOtherSensorTypeDescription gets the value of NumericSensorOtherSensorTypeDescription for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyNumericSensorOtherSensorTypeDescription() (value []string, err error) {
	retValue, err := instance.GetProperty("NumericSensorOtherSensorTypeDescription")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumericSensorPrimaryStatus sets the value of NumericSensorPrimaryStatus for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyNumericSensorPrimaryStatus(value []uint16) (err error) {
	return instance.SetProperty("NumericSensorPrimaryStatus", value)
}

// GetNumericSensorPrimaryStatus gets the value of NumericSensorPrimaryStatus for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyNumericSensorPrimaryStatus() (value []uint16, err error) {
	retValue, err := instance.GetProperty("NumericSensorPrimaryStatus")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumericSensorRateUnits sets the value of NumericSensorRateUnits for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyNumericSensorRateUnits(value []uint16) (err error) {
	return instance.SetProperty("NumericSensorRateUnits", value)
}

// GetNumericSensorRateUnits gets the value of NumericSensorRateUnits for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyNumericSensorRateUnits() (value []uint16, err error) {
	retValue, err := instance.GetProperty("NumericSensorRateUnits")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumericSensorSensorType sets the value of NumericSensorSensorType for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyNumericSensorSensorType(value []uint16) (err error) {
	return instance.SetProperty("NumericSensorSensorType", value)
}

// GetNumericSensorSensorType gets the value of NumericSensorSensorType for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyNumericSensorSensorType() (value []uint16, err error) {
	retValue, err := instance.GetProperty("NumericSensorSensorType")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumericSensorUnitModifier sets the value of NumericSensorUnitModifier for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyNumericSensorUnitModifier(value []int32) (err error) {
	return instance.SetProperty("NumericSensorUnitModifier", value)
}

// GetNumericSensorUnitModifier gets the value of NumericSensorUnitModifier for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyNumericSensorUnitModifier() (value []int32, err error) {
	retValue, err := instance.GetProperty("NumericSensorUnitModifier")
	if err != nil {
		return
	}
	value, ok := retValue.([]int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumericSensorUpperThresholdCritical sets the value of NumericSensorUpperThresholdCritical for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyNumericSensorUpperThresholdCritical(value []int32) (err error) {
	return instance.SetProperty("NumericSensorUpperThresholdCritical", value)
}

// GetNumericSensorUpperThresholdCritical gets the value of NumericSensorUpperThresholdCritical for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyNumericSensorUpperThresholdCritical() (value []int32, err error) {
	retValue, err := instance.GetProperty("NumericSensorUpperThresholdCritical")
	if err != nil {
		return
	}
	value, ok := retValue.([]int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumericSensorUpperThresholdFatal sets the value of NumericSensorUpperThresholdFatal for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyNumericSensorUpperThresholdFatal(value []int32) (err error) {
	return instance.SetProperty("NumericSensorUpperThresholdFatal", value)
}

// GetNumericSensorUpperThresholdFatal gets the value of NumericSensorUpperThresholdFatal for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyNumericSensorUpperThresholdFatal() (value []int32, err error) {
	retValue, err := instance.GetProperty("NumericSensorUpperThresholdFatal")
	if err != nil {
		return
	}
	value, ok := retValue.([]int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumericSensorUpperThresholdNonCritical sets the value of NumericSensorUpperThresholdNonCritical for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyNumericSensorUpperThresholdNonCritical(value []int32) (err error) {
	return instance.SetProperty("NumericSensorUpperThresholdNonCritical", value)
}

// GetNumericSensorUpperThresholdNonCritical gets the value of NumericSensorUpperThresholdNonCritical for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyNumericSensorUpperThresholdNonCritical() (value []int32, err error) {
	retValue, err := instance.GetProperty("NumericSensorUpperThresholdNonCritical")
	if err != nil {
		return
	}
	value, ok := retValue.([]int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOneTimeBootSource sets the value of OneTimeBootSource for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyOneTimeBootSource(value uint8) (err error) {
	return instance.SetProperty("OneTimeBootSource", value)
}

// GetOneTimeBootSource gets the value of OneTimeBootSource for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyOneTimeBootSource() (value uint8, err error) {
	retValue, err := instance.GetProperty("OneTimeBootSource")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOperationalStatus sets the value of OperationalStatus for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyOperationalStatus(value []uint16) (err error) {
	return instance.SetProperty("OperationalStatus", value)
}

// GetOperationalStatus gets the value of OperationalStatus for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyOperationalStatus() (value []uint16, err error) {
	retValue, err := instance.GetProperty("OperationalStatus")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOSEnabledState sets the value of OSEnabledState for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyOSEnabledState(value uint16) (err error) {
	return instance.SetProperty("OSEnabledState", value)
}

// GetOSEnabledState gets the value of OSEnabledState for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyOSEnabledState() (value uint16, err error) {
	retValue, err := instance.GetProperty("OSEnabledState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOSType sets the value of OSType for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyOSType(value uint16) (err error) {
	return instance.SetProperty("OSType", value)
}

// GetOSType gets the value of OSType for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyOSType() (value uint16, err error) {
	retValue, err := instance.GetProperty("OSType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOSVersion sets the value of OSVersion for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyOSVersion(value string) (err error) {
	return instance.SetProperty("OSVersion", value)
}

// GetOSVersion gets the value of OSVersion for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyOSVersion() (value string, err error) {
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

// SetOtherDedicatedDescriptions sets the value of OtherDedicatedDescriptions for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyOtherDedicatedDescriptions(value []string) (err error) {
	return instance.SetProperty("OtherDedicatedDescriptions", value)
}

// GetOtherDedicatedDescriptions gets the value of OtherDedicatedDescriptions for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyOtherDedicatedDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("OtherDedicatedDescriptions")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherIdentifyingInfo sets the value of OtherIdentifyingInfo for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyOtherIdentifyingInfo(value []string) (err error) {
	return instance.SetProperty("OtherIdentifyingInfo", value)
}

// GetOtherIdentifyingInfo gets the value of OtherIdentifyingInfo for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyOtherIdentifyingInfo() (value []string, err error) {
	retValue, err := instance.GetProperty("OtherIdentifyingInfo")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPartNumber sets the value of PartNumber for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyPartNumber(value string) (err error) {
	return instance.SetProperty("PartNumber", value)
}

// GetPartNumber gets the value of PartNumber for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyPartNumber() (value string, err error) {
	retValue, err := instance.GetProperty("PartNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPersistentBootConfigOrder sets the value of PersistentBootConfigOrder for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyPersistentBootConfigOrder(value []uint8) (err error) {
	return instance.SetProperty("PersistentBootConfigOrder", value)
}

// GetPersistentBootConfigOrder gets the value of PersistentBootConfigOrder for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyPersistentBootConfigOrder() (value []uint8, err error) {
	retValue, err := instance.GetProperty("PersistentBootConfigOrder")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPowerAllocationLimit sets the value of PowerAllocationLimit for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyPowerAllocationLimit(value uint64) (err error) {
	return instance.SetProperty("PowerAllocationLimit", value)
}

// GetPowerAllocationLimit gets the value of PowerAllocationLimit for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyPowerAllocationLimit() (value uint64, err error) {
	retValue, err := instance.GetProperty("PowerAllocationLimit")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPowerUtilizationMode sets the value of PowerUtilizationMode for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyPowerUtilizationMode(value uint16) (err error) {
	return instance.SetProperty("PowerUtilizationMode", value)
}

// GetPowerUtilizationMode gets the value of PowerUtilizationMode for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyPowerUtilizationMode() (value uint16, err error) {
	retValue, err := instance.GetProperty("PowerUtilizationMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPowerUtilizationModesSupported sets the value of PowerUtilizationModesSupported for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyPowerUtilizationModesSupported(value []uint16) (err error) {
	return instance.SetProperty("PowerUtilizationModesSupported", value)
}

// GetPowerUtilizationModesSupported gets the value of PowerUtilizationModesSupported for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyPowerUtilizationModesSupported() (value []uint16, err error) {
	retValue, err := instance.GetProperty("PowerUtilizationModesSupported")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProcessorCurrentClockSpeed sets the value of ProcessorCurrentClockSpeed for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyProcessorCurrentClockSpeed(value uint32) (err error) {
	return instance.SetProperty("ProcessorCurrentClockSpeed", value)
}

// GetProcessorCurrentClockSpeed gets the value of ProcessorCurrentClockSpeed for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyProcessorCurrentClockSpeed() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProcessorCurrentClockSpeed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProcessorFamily sets the value of ProcessorFamily for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyProcessorFamily(value uint16) (err error) {
	return instance.SetProperty("ProcessorFamily", value)
}

// GetProcessorFamily gets the value of ProcessorFamily for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyProcessorFamily() (value uint16, err error) {
	retValue, err := instance.GetProperty("ProcessorFamily")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProcessorMaxClockSpeed sets the value of ProcessorMaxClockSpeed for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyProcessorMaxClockSpeed(value uint32) (err error) {
	return instance.SetProperty("ProcessorMaxClockSpeed", value)
}

// GetProcessorMaxClockSpeed gets the value of ProcessorMaxClockSpeed for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyProcessorMaxClockSpeed() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProcessorMaxClockSpeed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequestedState sets the value of RequestedState for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyRequestedState(value uint16) (err error) {
	return instance.SetProperty("RequestedState", value)
}

// GetRequestedState gets the value of RequestedState for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyRequestedState() (value uint16, err error) {
	retValue, err := instance.GetProperty("RequestedState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSerialNumber sets the value of SerialNumber for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertySerialNumber(value string) (err error) {
	return instance.SetProperty("SerialNumber", value)
}

// GetSerialNumber gets the value of SerialNumber for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertySerialNumber() (value string, err error) {
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

// SetSKU sets the value of SKU for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertySKU(value string) (err error) {
	return instance.SetProperty("SKU", value)
}

// GetSKU gets the value of SKU for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertySKU() (value string, err error) {
	retValue, err := instance.GetProperty("SKU")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStructuredBootString sets the value of StructuredBootString for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyStructuredBootString(value []string) (err error) {
	return instance.SetProperty("StructuredBootString", value)
}

// GetStructuredBootString gets the value of StructuredBootString for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyStructuredBootString() (value []string, err error) {
	retValue, err := instance.GetProperty("StructuredBootString")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTag sets the value of Tag for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyTag(value string) (err error) {
	return instance.SetProperty("Tag", value)
}

// GetTag gets the value of Tag for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyTag() (value string, err error) {
	retValue, err := instance.GetProperty("Tag")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVersion sets the value of Version for the instance
func (instance *CIM_PhysicalComputerSystemView) SetPropertyVersion(value string) (err error) {
	return instance.SetProperty("Version", value)
}

// GetVersion gets the value of Version for the instance
func (instance *CIM_PhysicalComputerSystemView) GetPropertyVersion() (value string, err error) {
	retValue, err := instance.GetProperty("Version")
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

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="RequestedState" type="uint16 "></param>
// <param name="TimeoutPeriod" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_PhysicalComputerSystemView) RequestStateChange( /* IN */ RequestedState uint16,
	/* IN/OUT */ Job CIM_ConcreteJob,
	/* IN */ TimeoutPeriod string,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("RequestStateChange", Action, PercentComplete, Timeout, RequestedState, TimeoutPeriod)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="LogInstanceID" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_PhysicalComputerSystemView) ClearLog( /* IN */ LogInstanceID string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ClearLog", LogInstanceID)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Classifications" type="uint16 []"></param>
// <param name="InstallOptions" type="uint16 []"></param>
// <param name="InstallOptionsValues" type="string []"></param>
// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="URI" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_PhysicalComputerSystemView) InstallSoftwareFromURI( /* IN/OUT */ Job CIM_ConcreteJob,
	/* IN */ Classifications []uint16,
	/* IN */ URI string,
	/* IN */ InstallOptions []uint16,
	/* IN */ InstallOptionsValues []string,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("InstallSoftwareFromURI", Action, PercentComplete, Timeout, Classifications, URI, InstallOptions, InstallOptionsValues)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="StructuredBootString" type="string []"></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_PhysicalComputerSystemView) ModifyPersistentBootConfigOrder( /* IN */ StructuredBootString []string,
	/* IN/OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("ModifyPersistentBootConfigOrder", Action, PercentComplete, Timeout, StructuredBootString)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="StructuredBootString" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_PhysicalComputerSystemView) SetOneTimeBootSource( /* IN */ StructuredBootString string,
	/* IN/OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("SetOneTimeBootSource", Action, PercentComplete, Timeout, StructuredBootString)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
