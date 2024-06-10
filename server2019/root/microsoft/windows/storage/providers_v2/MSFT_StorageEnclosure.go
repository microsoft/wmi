// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_StorageEnclosure struct
type MSFT_StorageEnclosure struct {
	*MSFT_StorageFaultDomain

	//
	BusType StorageEnclosure_BusType

	// An array containing the operational status of each current sensor of the enclosure.
	///0 - 'Unknown'
	///2 - 'OK': The element is present and working with no issues detected.
	///3 - 'Degraded': The element detected one or more non-critical issues.
	///6 - 'Error': The element detected one or more critical issues.
	///7 - 'Non-Recoverable Error': The element detected one or more non-recoverable issues.
	///0xD009 - 'Not Installed': The element is not present.
	///0xD00A - 'Not Available': The element is present but has problems.
	///0xD00B - 'No Access Allowed': No access is allowed to the element.
	///0xD00C - 'Not Reported'
	CurrentSensorOperationalStatus []StorageEnclosure_CurrentSensorOperationalStatus

	// DeviceId is an address or other identifier that uniquely names the enclosure. For example, DeviceId is the enclosure GUID in Storage Spaces provider.
	DeviceId string

	// An array containing the operational status of each fan of the enclosure.
	///0 - 'Unknown'
	///2 - 'OK': The element is present and working with no issues detected.
	///3 - 'Degraded': The element detected one or more non-critical issues.
	///6 - 'Error': The element detected one or more critical issues.
	///7 - 'Non-Recoverable Error': The element detected one or more non-recoverable issues.
	///0xD009 - 'Not Installed': The element is not present.
	///0xD00A - 'Not Available': The element is present but has problems.
	///0xD00B - 'No Access Allowed': No access is allowed to the element.
	///0xD00C - 'Not Reported'
	FanOperationalStatus []StorageEnclosure_FanOperationalStatus

	// This field is a string representation of the enclosure's firmware version.
	FirmwareVersion string

	// An array containing the operational status of each controller of the enclosure.
	///0 - 'Unknown'
	///2 - 'OK': The element is present and working with no issues detected.
	///3 - 'Degraded': The element detected one or more non-critical issues.
	///6 - 'Error': The element detected one or more critical issues.
	///7 - 'Non-Recoverable Error': The element detected one or more non-recoverable issues.
	///0xD009 - 'Not Installed': The element is not present.
	///0xD00A - 'Not Available': The element is present but has problems.
	///0xD00B - 'No Access Allowed': No access is allowed to the element.
	///0xD00C - 'Not Reported'
	IOControllerOperationalStatus []StorageEnclosure_IOControllerOperationalStatus

	// Number of slots hosted within the enclosure
	NumberOfSlots uint32

	// An array containing the operational status of each power supply of the enclosure.
	///0 - 'Unknown'
	///2 - 'OK': The element is present and working with no issues detected.
	///3 - 'Degraded': The element detected one or more non-critical issues.
	///6 - 'Error': The element detected one or more critical issues.
	///7 - 'Non-Recoverable Error': The element detected one or more non-recoverable issues.
	///0xD009 - 'Not Installed': The element is not present.
	///0xD00A - 'Not Available': The element is present but has problems.
	///0xD00B - 'No Access Allowed': No access is allowed to the element.
	///0xD00C - 'Not Reported'
	PowerSupplyOperationalStatus []StorageEnclosure_PowerSupplyOperationalStatus

	//
	SlotOperationalStatus []uint16

	// An array containing the operational status of each temperature sensor of the enclosure.
	///0 - 'Unknown'
	///2 - 'OK': The element is present and working with no issues detected.
	///3 - 'Degraded': The element detected one or more non-critical issues.
	///6 - 'Error': The element detected one or more critical issues.
	///7 - 'Non-Recoverable Error': The element detected one or more non-recoverable issues.
	///0xD009 - 'Not Installed': The element is not present.
	///0xD00A - 'Not Available': The element is present but has problems.
	///0xD00B - 'No Access Allowed': No access is allowed to the element.
	///0xD00C - 'Not Reported'
	TemperatureSensorOperationalStatus []StorageEnclosure_TemperatureSensorOperationalStatus

	// An array containing the operational status of each voltage sensor of the enclosure.
	///0 - 'Unknown'
	///2 - 'OK': The element is present and working with no issues detected.
	///3 - 'Degraded': The element detected one or more non-critical issues.
	///6 - 'Error': The element detected one or more critical issues.
	///7 - 'Non-Recoverable Error': The element detected one or more non-recoverable issues.
	///0xD009 - 'Not Installed': The element is not present.
	///0xD00A - 'Not Available': The element is present but has problems.
	///0xD00B - 'No Access Allowed': No access is allowed to the element.
	///0xD00C - 'Not Reported'
	VoltageSensorOperationalStatus []StorageEnclosure_VoltageSensorOperationalStatus
}

func NewMSFT_StorageEnclosureEx1(instance *cim.WmiInstance) (newInstance *MSFT_StorageEnclosure, err error) {
	tmp, err := NewMSFT_StorageFaultDomainEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_StorageEnclosure{
		MSFT_StorageFaultDomain: tmp,
	}
	return
}

func NewMSFT_StorageEnclosureEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_StorageEnclosure, err error) {
	tmp, err := NewMSFT_StorageFaultDomainEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_StorageEnclosure{
		MSFT_StorageFaultDomain: tmp,
	}
	return
}

// SetBusType sets the value of BusType for the instance
func (instance *MSFT_StorageEnclosure) SetPropertyBusType(value StorageEnclosure_BusType) (err error) {
	return instance.SetProperty("BusType", (value))
}

// GetBusType gets the value of BusType for the instance
func (instance *MSFT_StorageEnclosure) GetPropertyBusType() (value StorageEnclosure_BusType, err error) {
	retValue, err := instance.GetProperty("BusType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = StorageEnclosure_BusType(valuetmp)

	return
}

// SetCurrentSensorOperationalStatus sets the value of CurrentSensorOperationalStatus for the instance
func (instance *MSFT_StorageEnclosure) SetPropertyCurrentSensorOperationalStatus(value []StorageEnclosure_CurrentSensorOperationalStatus) (err error) {
	return instance.SetProperty("CurrentSensorOperationalStatus", (value))
}

// GetCurrentSensorOperationalStatus gets the value of CurrentSensorOperationalStatus for the instance
func (instance *MSFT_StorageEnclosure) GetPropertyCurrentSensorOperationalStatus() (value []StorageEnclosure_CurrentSensorOperationalStatus, err error) {
	retValue, err := instance.GetProperty("CurrentSensorOperationalStatus")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, StorageEnclosure_CurrentSensorOperationalStatus(valuetmp))
	}

	return
}

// SetDeviceId sets the value of DeviceId for the instance
func (instance *MSFT_StorageEnclosure) SetPropertyDeviceId(value string) (err error) {
	return instance.SetProperty("DeviceId", (value))
}

// GetDeviceId gets the value of DeviceId for the instance
func (instance *MSFT_StorageEnclosure) GetPropertyDeviceId() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetFanOperationalStatus sets the value of FanOperationalStatus for the instance
func (instance *MSFT_StorageEnclosure) SetPropertyFanOperationalStatus(value []StorageEnclosure_FanOperationalStatus) (err error) {
	return instance.SetProperty("FanOperationalStatus", (value))
}

// GetFanOperationalStatus gets the value of FanOperationalStatus for the instance
func (instance *MSFT_StorageEnclosure) GetPropertyFanOperationalStatus() (value []StorageEnclosure_FanOperationalStatus, err error) {
	retValue, err := instance.GetProperty("FanOperationalStatus")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, StorageEnclosure_FanOperationalStatus(valuetmp))
	}

	return
}

// SetFirmwareVersion sets the value of FirmwareVersion for the instance
func (instance *MSFT_StorageEnclosure) SetPropertyFirmwareVersion(value string) (err error) {
	return instance.SetProperty("FirmwareVersion", (value))
}

// GetFirmwareVersion gets the value of FirmwareVersion for the instance
func (instance *MSFT_StorageEnclosure) GetPropertyFirmwareVersion() (value string, err error) {
	retValue, err := instance.GetProperty("FirmwareVersion")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetIOControllerOperationalStatus sets the value of IOControllerOperationalStatus for the instance
func (instance *MSFT_StorageEnclosure) SetPropertyIOControllerOperationalStatus(value []StorageEnclosure_IOControllerOperationalStatus) (err error) {
	return instance.SetProperty("IOControllerOperationalStatus", (value))
}

// GetIOControllerOperationalStatus gets the value of IOControllerOperationalStatus for the instance
func (instance *MSFT_StorageEnclosure) GetPropertyIOControllerOperationalStatus() (value []StorageEnclosure_IOControllerOperationalStatus, err error) {
	retValue, err := instance.GetProperty("IOControllerOperationalStatus")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, StorageEnclosure_IOControllerOperationalStatus(valuetmp))
	}

	return
}

// SetNumberOfSlots sets the value of NumberOfSlots for the instance
func (instance *MSFT_StorageEnclosure) SetPropertyNumberOfSlots(value uint32) (err error) {
	return instance.SetProperty("NumberOfSlots", (value))
}

// GetNumberOfSlots gets the value of NumberOfSlots for the instance
func (instance *MSFT_StorageEnclosure) GetPropertyNumberOfSlots() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfSlots")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetPowerSupplyOperationalStatus sets the value of PowerSupplyOperationalStatus for the instance
func (instance *MSFT_StorageEnclosure) SetPropertyPowerSupplyOperationalStatus(value []StorageEnclosure_PowerSupplyOperationalStatus) (err error) {
	return instance.SetProperty("PowerSupplyOperationalStatus", (value))
}

// GetPowerSupplyOperationalStatus gets the value of PowerSupplyOperationalStatus for the instance
func (instance *MSFT_StorageEnclosure) GetPropertyPowerSupplyOperationalStatus() (value []StorageEnclosure_PowerSupplyOperationalStatus, err error) {
	retValue, err := instance.GetProperty("PowerSupplyOperationalStatus")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, StorageEnclosure_PowerSupplyOperationalStatus(valuetmp))
	}

	return
}

// SetSlotOperationalStatus sets the value of SlotOperationalStatus for the instance
func (instance *MSFT_StorageEnclosure) SetPropertySlotOperationalStatus(value []uint16) (err error) {
	return instance.SetProperty("SlotOperationalStatus", (value))
}

// GetSlotOperationalStatus gets the value of SlotOperationalStatus for the instance
func (instance *MSFT_StorageEnclosure) GetPropertySlotOperationalStatus() (value []uint16, err error) {
	retValue, err := instance.GetProperty("SlotOperationalStatus")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint16)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint16(valuetmp))
	}

	return
}

// SetTemperatureSensorOperationalStatus sets the value of TemperatureSensorOperationalStatus for the instance
func (instance *MSFT_StorageEnclosure) SetPropertyTemperatureSensorOperationalStatus(value []StorageEnclosure_TemperatureSensorOperationalStatus) (err error) {
	return instance.SetProperty("TemperatureSensorOperationalStatus", (value))
}

// GetTemperatureSensorOperationalStatus gets the value of TemperatureSensorOperationalStatus for the instance
func (instance *MSFT_StorageEnclosure) GetPropertyTemperatureSensorOperationalStatus() (value []StorageEnclosure_TemperatureSensorOperationalStatus, err error) {
	retValue, err := instance.GetProperty("TemperatureSensorOperationalStatus")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, StorageEnclosure_TemperatureSensorOperationalStatus(valuetmp))
	}

	return
}

// SetVoltageSensorOperationalStatus sets the value of VoltageSensorOperationalStatus for the instance
func (instance *MSFT_StorageEnclosure) SetPropertyVoltageSensorOperationalStatus(value []StorageEnclosure_VoltageSensorOperationalStatus) (err error) {
	return instance.SetProperty("VoltageSensorOperationalStatus", (value))
}

// GetVoltageSensorOperationalStatus gets the value of VoltageSensorOperationalStatus for the instance
func (instance *MSFT_StorageEnclosure) GetPropertyVoltageSensorOperationalStatus() (value []StorageEnclosure_VoltageSensorOperationalStatus, err error) {
	retValue, err := instance.GetProperty("VoltageSensorOperationalStatus")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, StorageEnclosure_VoltageSensorOperationalStatus(valuetmp))
	}

	return
}

// This method allows a user to perform certain identification tasks on the enclosure and its elements.

// <param name="Enable" type="bool ">If set to TRUE, this instructs the enclosure to enable its identification LED on the specified element. The identification LED should remain enabled until a second call to IdentifyElement on the same element is made with this parameter specified as FALSE.</param>
// <param name="SlotNumbers" type="uint32 []">The numbers of the slots on which to enable or disable identification.</param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageEnclosure) IdentifyElement( /* IN */ Enable bool,
	/* IN */ SlotNumbers []uint32,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("IdentifyElement", Enable, SlotNumbers)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method returns the vendor specific data from an enclosure.

// <param name="PageNumber" type="uint16 ">Denotes the page number for which vendor data is requested.</param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="VendorData" type="string ">The vendor specific data (page 04h for example) from an enclosure.</param>
func (instance *MSFT_StorageEnclosure) GetVendorData( /* IN */ PageNumber uint16,
	/* OUT */ VendorData string,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetVendorData", PageNumber)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="EnableMaintenanceMode" type="bool "></param>
// <param name="IgnoreDetachedVirtualDisks" type="bool "></param>
// <param name="Manufacturer" type="string "></param>
// <param name="Model" type="string "></param>
// <param name="Timeout" type="uint32 "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageEnclosure) Maintenance( /* IN */ EnableMaintenanceMode bool,
	/* IN */ Timeout uint32,
	/* IN */ Model string,
	/* IN */ Manufacturer string,
	/* IN */ IgnoreDetachedVirtualDisks bool,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Maintenance", EnableMaintenanceMode, Timeout, Model, Manufacturer, IgnoreDetachedVirtualDisks)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Off" type="bool "></param>
// <param name="SlotNumbers" type="uint32 []"></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageEnclosure) PowerElement( /* IN */ Off bool,
	/* IN */ SlotNumbers []uint32,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("PowerElement", Off, SlotNumbers)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ActiveSlotNumber" type="uint16 "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="FirmwareVersionInSlot" type="string []"></param>
// <param name="IsSlotWritable" type="bool []"></param>
// <param name="NumberOfSlots" type="uint16 "></param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="SlotNumber" type="uint16 []"></param>
// <param name="SupportsUpdate" type="bool "></param>
func (instance *MSFT_StorageEnclosure) GetFirmwareInformation( /* OUT */ SupportsUpdate bool,
	/* OUT */ NumberOfSlots uint16,
	/* OUT */ ActiveSlotNumber uint16,
	/* OUT */ SlotNumber []uint16,
	/* OUT */ IsSlotWritable []bool,
	/* OUT */ FirmwareVersionInSlot []string,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetFirmwareInformation")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ImagePath" type="string "></param>
// <param name="SlotNumber" type="uint16 "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageEnclosure) UpdateFirmware( /* IN */ ImagePath string,
	/* IN */ SlotNumber uint16,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("UpdateFirmware", ImagePath, SlotNumber)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
