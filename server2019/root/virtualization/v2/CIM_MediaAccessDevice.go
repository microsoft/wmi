// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_MediaAccessDevice struct
type CIM_MediaAccessDevice struct {
	*CIM_LogicalDevice

	// Capabilities of the MediaAccessDevice. For example, the Device may support "Random Access", removeable media and "Automatic Cleaning". In this case, the values 3, 7 and 9 would be written to the array.
	///Several of the enumerated values require some explanation: 1) Value 11, Supports Dual Sided Media, distinguishes a Device that can access both sides of dual sided Media, from a Device that reads only a single side and requires the Media to be flipped; and, 2) Value 12, Predismount Eject Not Required, indicates that Media does not have to be explicitly ejected from the Device before being accessed by a PickerElement.
	Capabilities []MediaAccessDevice_Capabilities

	// An array of free-form strings providing more detailed explanations for any of the AccessDevice features indicated in the Capabilities array. Note, each entry of this array is related to the entry in the Capabilities array that is located at the same index.
	CapabilityDescriptions []string

	// A free form string indicating the algorithm or tool used by the device to support compression. If it is not possible or not desired to describe the compression scheme (perhaps because it is not known), recommend using the following words: "Unknown" to represent that it is not known whether the device supports compression capabilities or not, "Compressed" to represent that the device supports compression capabilities but either its compression scheme is not known or not disclosed, and "Not Compressed" to represent that the devices does not support compression capabilities.
	CompressionMethod string

	// Default block size, in bytes, for this Device.
	DefaultBlockSize uint64

	// ErrorMethodology is a free-form string describing the type(s) of error detection and correction supported by this Device.
	ErrorMethodology string

	// The date and time on which the Device was last cleaned.
	LastCleaned string

	// Time in milliseconds from 'load' to being able to read or write a Media. For example, for DiskDrives, this is the interval between a disk not spinning to the disk reporting that it is ready for read/write (ie, the disk spinning at nominal speeds). For TapeDrives, this is the time from a Media being injected to reporting that it is ready for an application. This is usually at the tape's BOT area.
	LoadTime uint64

	// Time in milliseconds to move from the first location on the Media to the location that is furthest with respect to time. For a DiskDrive, this represents full seek + full rotational delay. For TapeDrives, this represents a search from the beginning of the tape to the most physically distant point. (The end of a tape may be at its most physically distant point, but this is not necessarily true.)
	MaxAccessTime uint64

	// Maximum block size, in bytes, for media accessed by this Device.
	MaxBlockSize uint64

	// Maximum size, in KBytes, of media supported by this Device. KBytes is interpreted as the number of bytes multiplied by 1000 (NOT the number of bytes multiplied by 1024).
	MaxMediaSize uint64

	// An unsigned integer indicating the maximum 'units' that can be used, with respect to the AccessDevice, before the Device should be cleaned. The property, UnitsDescription, defines how 'units' should be interpreted.
	MaxUnitsBeforeCleaning uint64

	// True indicates that the media is locked in the Device and can not be ejected. For non-removeable Devices, this value should be true.
	MediaIsLocked bool

	// Minimum block size, in bytes, for media accessed by this Device.
	MinBlockSize uint64

	// For a MediaAccessDevice that supports removable Media, the number of times that Media have been mounted for data transfer or to clean the Device. For Devices accessing nonremovable Media, such as hard disks, this property is not applicable and should be set to 0.
	MountCount uint64

	// Boolean indicating that the MediaAccessDevice needs cleaning. Whether manual or automatic cleaning is possible is indicated in the Capabilities array property.
	NeedsCleaning bool

	// When the MediaAccessDevice supports multiple individual Media, this property defines the maximum number which can be supported or inserted.
	NumberOfMediaSupported uint32

	// An enumeration indicating the operational security defined for the MediaAccessDevice. For example, information that the Device is "Read Only" (value=4) or "Boot Bypass" (value=6) can be described using this property.
	Security MediaAccessDevice_Security

	// For a MediaAccessDevice that supports removable Media, the most recent date and time that Media was mounted on the Device. For Devices accessing nonremovable Media, such as hard disks, this property has no meaning and is not applicable.
	TimeOfLastMount string

	// For a MediaAccessDevice that supports removable Media, the total time (in seconds) that Media have been mounted for data transfer or to clean the Device. For Devices accessing nonremovable Media, such as hard disks, this property is not applicable and should be set to 0.
	TotalMountTime uint64

	// The sustained data transfer rate in KB/sec that the Device can read from and write to a Media. This is a sustained, raw data rate. Maximum rates or rates assuming compression should not be reported in this property.
	UncompressedDataRate uint32

	// Defines 'Units' relative to its use in the property, MaxUnitsBeforeCleaning. This describes the criteria used to determine when the MediaAccessDevice should be cleaned.
	UnitsDescription string

	// An unsigned integer indicating the currently used 'units' of the AccessDevice, helpful to describe when the Device may require cleaning. The property, UnitsDescription, defines how 'units' should be interpreted.
	UnitsUsed uint64

	// Time in milliseconds from being able to read or write a Media to its 'unload'. For example, for DiskDrives, this is the interval between a disk spinning at nominal speeds and a disk not spinning. For TapeDrives, this is the time for a Media to go from its BOT to being fully ejected and accessible to a PickerElement or human operator.
	UnloadTime uint64
}

func NewCIM_MediaAccessDeviceEx1(instance *cim.WmiInstance) (newInstance *CIM_MediaAccessDevice, err error) {
	tmp, err := NewCIM_LogicalDeviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_MediaAccessDevice{
		CIM_LogicalDevice: tmp,
	}
	return
}

func NewCIM_MediaAccessDeviceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_MediaAccessDevice, err error) {
	tmp, err := NewCIM_LogicalDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_MediaAccessDevice{
		CIM_LogicalDevice: tmp,
	}
	return
}

// SetCapabilities sets the value of Capabilities for the instance
func (instance *CIM_MediaAccessDevice) SetPropertyCapabilities(value []MediaAccessDevice_Capabilities) (err error) {
	return instance.SetProperty("Capabilities", value)
}

// GetCapabilities gets the value of Capabilities for the instance
func (instance *CIM_MediaAccessDevice) GetPropertyCapabilities() (value []MediaAccessDevice_Capabilities, err error) {
	retValue, err := instance.GetProperty("Capabilities")
	if err != nil {
		return
	}
	value, ok := retValue.([]MediaAccessDevice_Capabilities)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCapabilityDescriptions sets the value of CapabilityDescriptions for the instance
func (instance *CIM_MediaAccessDevice) SetPropertyCapabilityDescriptions(value []string) (err error) {
	return instance.SetProperty("CapabilityDescriptions", value)
}

// GetCapabilityDescriptions gets the value of CapabilityDescriptions for the instance
func (instance *CIM_MediaAccessDevice) GetPropertyCapabilityDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("CapabilityDescriptions")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCompressionMethod sets the value of CompressionMethod for the instance
func (instance *CIM_MediaAccessDevice) SetPropertyCompressionMethod(value string) (err error) {
	return instance.SetProperty("CompressionMethod", value)
}

// GetCompressionMethod gets the value of CompressionMethod for the instance
func (instance *CIM_MediaAccessDevice) GetPropertyCompressionMethod() (value string, err error) {
	retValue, err := instance.GetProperty("CompressionMethod")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDefaultBlockSize sets the value of DefaultBlockSize for the instance
func (instance *CIM_MediaAccessDevice) SetPropertyDefaultBlockSize(value uint64) (err error) {
	return instance.SetProperty("DefaultBlockSize", value)
}

// GetDefaultBlockSize gets the value of DefaultBlockSize for the instance
func (instance *CIM_MediaAccessDevice) GetPropertyDefaultBlockSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("DefaultBlockSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetErrorMethodology sets the value of ErrorMethodology for the instance
func (instance *CIM_MediaAccessDevice) SetPropertyErrorMethodology(value string) (err error) {
	return instance.SetProperty("ErrorMethodology", value)
}

// GetErrorMethodology gets the value of ErrorMethodology for the instance
func (instance *CIM_MediaAccessDevice) GetPropertyErrorMethodology() (value string, err error) {
	retValue, err := instance.GetProperty("ErrorMethodology")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLastCleaned sets the value of LastCleaned for the instance
func (instance *CIM_MediaAccessDevice) SetPropertyLastCleaned(value string) (err error) {
	return instance.SetProperty("LastCleaned", value)
}

// GetLastCleaned gets the value of LastCleaned for the instance
func (instance *CIM_MediaAccessDevice) GetPropertyLastCleaned() (value string, err error) {
	retValue, err := instance.GetProperty("LastCleaned")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLoadTime sets the value of LoadTime for the instance
func (instance *CIM_MediaAccessDevice) SetPropertyLoadTime(value uint64) (err error) {
	return instance.SetProperty("LoadTime", value)
}

// GetLoadTime gets the value of LoadTime for the instance
func (instance *CIM_MediaAccessDevice) GetPropertyLoadTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("LoadTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxAccessTime sets the value of MaxAccessTime for the instance
func (instance *CIM_MediaAccessDevice) SetPropertyMaxAccessTime(value uint64) (err error) {
	return instance.SetProperty("MaxAccessTime", value)
}

// GetMaxAccessTime gets the value of MaxAccessTime for the instance
func (instance *CIM_MediaAccessDevice) GetPropertyMaxAccessTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxAccessTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxBlockSize sets the value of MaxBlockSize for the instance
func (instance *CIM_MediaAccessDevice) SetPropertyMaxBlockSize(value uint64) (err error) {
	return instance.SetProperty("MaxBlockSize", value)
}

// GetMaxBlockSize gets the value of MaxBlockSize for the instance
func (instance *CIM_MediaAccessDevice) GetPropertyMaxBlockSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxBlockSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxMediaSize sets the value of MaxMediaSize for the instance
func (instance *CIM_MediaAccessDevice) SetPropertyMaxMediaSize(value uint64) (err error) {
	return instance.SetProperty("MaxMediaSize", value)
}

// GetMaxMediaSize gets the value of MaxMediaSize for the instance
func (instance *CIM_MediaAccessDevice) GetPropertyMaxMediaSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxMediaSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxUnitsBeforeCleaning sets the value of MaxUnitsBeforeCleaning for the instance
func (instance *CIM_MediaAccessDevice) SetPropertyMaxUnitsBeforeCleaning(value uint64) (err error) {
	return instance.SetProperty("MaxUnitsBeforeCleaning", value)
}

// GetMaxUnitsBeforeCleaning gets the value of MaxUnitsBeforeCleaning for the instance
func (instance *CIM_MediaAccessDevice) GetPropertyMaxUnitsBeforeCleaning() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxUnitsBeforeCleaning")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMediaIsLocked sets the value of MediaIsLocked for the instance
func (instance *CIM_MediaAccessDevice) SetPropertyMediaIsLocked(value bool) (err error) {
	return instance.SetProperty("MediaIsLocked", value)
}

// GetMediaIsLocked gets the value of MediaIsLocked for the instance
func (instance *CIM_MediaAccessDevice) GetPropertyMediaIsLocked() (value bool, err error) {
	retValue, err := instance.GetProperty("MediaIsLocked")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMinBlockSize sets the value of MinBlockSize for the instance
func (instance *CIM_MediaAccessDevice) SetPropertyMinBlockSize(value uint64) (err error) {
	return instance.SetProperty("MinBlockSize", value)
}

// GetMinBlockSize gets the value of MinBlockSize for the instance
func (instance *CIM_MediaAccessDevice) GetPropertyMinBlockSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("MinBlockSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMountCount sets the value of MountCount for the instance
func (instance *CIM_MediaAccessDevice) SetPropertyMountCount(value uint64) (err error) {
	return instance.SetProperty("MountCount", value)
}

// GetMountCount gets the value of MountCount for the instance
func (instance *CIM_MediaAccessDevice) GetPropertyMountCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("MountCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNeedsCleaning sets the value of NeedsCleaning for the instance
func (instance *CIM_MediaAccessDevice) SetPropertyNeedsCleaning(value bool) (err error) {
	return instance.SetProperty("NeedsCleaning", value)
}

// GetNeedsCleaning gets the value of NeedsCleaning for the instance
func (instance *CIM_MediaAccessDevice) GetPropertyNeedsCleaning() (value bool, err error) {
	retValue, err := instance.GetProperty("NeedsCleaning")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfMediaSupported sets the value of NumberOfMediaSupported for the instance
func (instance *CIM_MediaAccessDevice) SetPropertyNumberOfMediaSupported(value uint32) (err error) {
	return instance.SetProperty("NumberOfMediaSupported", value)
}

// GetNumberOfMediaSupported gets the value of NumberOfMediaSupported for the instance
func (instance *CIM_MediaAccessDevice) GetPropertyNumberOfMediaSupported() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfMediaSupported")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSecurity sets the value of Security for the instance
func (instance *CIM_MediaAccessDevice) SetPropertySecurity(value MediaAccessDevice_Security) (err error) {
	return instance.SetProperty("Security", value)
}

// GetSecurity gets the value of Security for the instance
func (instance *CIM_MediaAccessDevice) GetPropertySecurity() (value MediaAccessDevice_Security, err error) {
	retValue, err := instance.GetProperty("Security")
	if err != nil {
		return
	}
	value, ok := retValue.(MediaAccessDevice_Security)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimeOfLastMount sets the value of TimeOfLastMount for the instance
func (instance *CIM_MediaAccessDevice) SetPropertyTimeOfLastMount(value string) (err error) {
	return instance.SetProperty("TimeOfLastMount", value)
}

// GetTimeOfLastMount gets the value of TimeOfLastMount for the instance
func (instance *CIM_MediaAccessDevice) GetPropertyTimeOfLastMount() (value string, err error) {
	retValue, err := instance.GetProperty("TimeOfLastMount")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalMountTime sets the value of TotalMountTime for the instance
func (instance *CIM_MediaAccessDevice) SetPropertyTotalMountTime(value uint64) (err error) {
	return instance.SetProperty("TotalMountTime", value)
}

// GetTotalMountTime gets the value of TotalMountTime for the instance
func (instance *CIM_MediaAccessDevice) GetPropertyTotalMountTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalMountTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUncompressedDataRate sets the value of UncompressedDataRate for the instance
func (instance *CIM_MediaAccessDevice) SetPropertyUncompressedDataRate(value uint32) (err error) {
	return instance.SetProperty("UncompressedDataRate", value)
}

// GetUncompressedDataRate gets the value of UncompressedDataRate for the instance
func (instance *CIM_MediaAccessDevice) GetPropertyUncompressedDataRate() (value uint32, err error) {
	retValue, err := instance.GetProperty("UncompressedDataRate")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUnitsDescription sets the value of UnitsDescription for the instance
func (instance *CIM_MediaAccessDevice) SetPropertyUnitsDescription(value string) (err error) {
	return instance.SetProperty("UnitsDescription", value)
}

// GetUnitsDescription gets the value of UnitsDescription for the instance
func (instance *CIM_MediaAccessDevice) GetPropertyUnitsDescription() (value string, err error) {
	retValue, err := instance.GetProperty("UnitsDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUnitsUsed sets the value of UnitsUsed for the instance
func (instance *CIM_MediaAccessDevice) SetPropertyUnitsUsed(value uint64) (err error) {
	return instance.SetProperty("UnitsUsed", value)
}

// GetUnitsUsed gets the value of UnitsUsed for the instance
func (instance *CIM_MediaAccessDevice) GetPropertyUnitsUsed() (value uint64, err error) {
	retValue, err := instance.GetProperty("UnitsUsed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUnloadTime sets the value of UnloadTime for the instance
func (instance *CIM_MediaAccessDevice) SetPropertyUnloadTime(value uint64) (err error) {
	return instance.SetProperty("UnloadTime", value)
}

// GetUnloadTime gets the value of UnloadTime for the instance
func (instance *CIM_MediaAccessDevice) GetPropertyUnloadTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("UnloadTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Method to lock and unlock the media in a removeable Access Device. The method takes one parameter as input - a boolean indicating whether to lock or unlock. TRUE indicates that the media should be locked in the Device, FALSE indicates that the media should be unlocked. The method returns 0 if successful, 1 if not supported, and any other value if an error occurred. The set of possible return codes should be specified in a ValueMap qualifier on the method. The strings to which the ValueMap contents are 'translated' should be specified as a Values array qualifier on the method.

// <param name="Lock" type="bool ">If TRUE, lock the media. If FALSE release the media.</param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_MediaAccessDevice) LockMedia( /* IN */ Lock bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("LockMedia", Lock)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
