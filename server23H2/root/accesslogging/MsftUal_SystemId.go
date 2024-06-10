// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.AccessLogging
//
// ////////////////////////////////////////////
package accesslogging

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MsftUal_SystemId struct
type MsftUal_SystemId struct {
	*cim.WmiInstance

	// The number of cores for an instance of the physical processor in the system. For example, for a dual-core processor system, this property has a value of 2.
	CoresPerPhysicalProcessor uint32

	// The date and time that the current operating system first became operational with this set of system identity properties. If the properties of a system change, then a new System Identity record is created.
	CreationTime string

	// The number of logical processors for an instance of a Hyper-Thread capable physical processor in the system. For example, in a Hyper-Thread quad-core processor system, this property has a value of 8.
	LogicalProcessorsPerPhysicalProcessor uint32

	// The maximum system memory size (in bytes). For a virtual machine, this value represents what the hypervisor configured for the virtual machine’s memory size.
	MaximumMemory uint64

	// The build number of the operating system.
	OSBuildNumber uint32

	// The code for the country or region that an operating system uses. Values are based on international phone dialing prefixes.
	OSCountryCode string

	// The number, in minutes, an operating system is offset from Greenwich mean time (GMT). The number is positive, negative, or zero.
	OSCurrentTimeZone int16

	// If True, the daylight savings mode is ON.
	OSDaylightInEffect bool

	// The date and time the operating system was last restarted.
	OSLastBootUpTime string

	// The major portion of the version number of the operating system.
	OSMajor uint32

	// The minor portion of the version number of the operating system.
	OSMinor uint32

	// An integer that represents the operating system platform. The possible values of the data property are "1" to indicate an unsupported Windows system and "2" to indicate a supported Windows system.
	OSPlatformId uint32

	// An enumeration type that identifies the operating system that you are running.
	OSProductType uint32

	// The operating system product serial identification number.
	OSSerialNumber string

	// The SuiteMask of the local system.
	OSSuiteMask uint32

	// The number of physical processors currently available on a system.
	PhysicalProcessorCount uint32

	// The major version number of the service pack.
	ServicePackMajor uint32

	// The minor version number of the service pack.
	ServicePackMinor uint32

	// The server name according to the domain name server (DNS).
	SystemDNSHostName string

	// The name of the domain, or workgroup, to which the server belongs.
	SystemDomainName string

	// The name of the BIOS manufacturer.
	SystemManufacturer string

	// The product name specified in the system BIOS.
	SystemProductName string

	// The unit identification for the local server.
	SystemSerialNumber string

	// The SMBIOS reported universally unique identifier for this server unit.
	SystemSMBIOSUUID string
}

func NewMsftUal_SystemIdEx1(instance *cim.WmiInstance) (newInstance *MsftUal_SystemId, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MsftUal_SystemId{
		WmiInstance: tmp,
	}
	return
}

func NewMsftUal_SystemIdEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MsftUal_SystemId, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MsftUal_SystemId{
		WmiInstance: tmp,
	}
	return
}

// SetCoresPerPhysicalProcessor sets the value of CoresPerPhysicalProcessor for the instance
func (instance *MsftUal_SystemId) SetPropertyCoresPerPhysicalProcessor(value uint32) (err error) {
	return instance.SetProperty("CoresPerPhysicalProcessor", (value))
}

// GetCoresPerPhysicalProcessor gets the value of CoresPerPhysicalProcessor for the instance
func (instance *MsftUal_SystemId) GetPropertyCoresPerPhysicalProcessor() (value uint32, err error) {
	retValue, err := instance.GetProperty("CoresPerPhysicalProcessor")
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

// SetCreationTime sets the value of CreationTime for the instance
func (instance *MsftUal_SystemId) SetPropertyCreationTime(value string) (err error) {
	return instance.SetProperty("CreationTime", (value))
}

// GetCreationTime gets the value of CreationTime for the instance
func (instance *MsftUal_SystemId) GetPropertyCreationTime() (value string, err error) {
	retValue, err := instance.GetProperty("CreationTime")
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

// SetLogicalProcessorsPerPhysicalProcessor sets the value of LogicalProcessorsPerPhysicalProcessor for the instance
func (instance *MsftUal_SystemId) SetPropertyLogicalProcessorsPerPhysicalProcessor(value uint32) (err error) {
	return instance.SetProperty("LogicalProcessorsPerPhysicalProcessor", (value))
}

// GetLogicalProcessorsPerPhysicalProcessor gets the value of LogicalProcessorsPerPhysicalProcessor for the instance
func (instance *MsftUal_SystemId) GetPropertyLogicalProcessorsPerPhysicalProcessor() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogicalProcessorsPerPhysicalProcessor")
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

// SetMaximumMemory sets the value of MaximumMemory for the instance
func (instance *MsftUal_SystemId) SetPropertyMaximumMemory(value uint64) (err error) {
	return instance.SetProperty("MaximumMemory", (value))
}

// GetMaximumMemory gets the value of MaximumMemory for the instance
func (instance *MsftUal_SystemId) GetPropertyMaximumMemory() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaximumMemory")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetOSBuildNumber sets the value of OSBuildNumber for the instance
func (instance *MsftUal_SystemId) SetPropertyOSBuildNumber(value uint32) (err error) {
	return instance.SetProperty("OSBuildNumber", (value))
}

// GetOSBuildNumber gets the value of OSBuildNumber for the instance
func (instance *MsftUal_SystemId) GetPropertyOSBuildNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("OSBuildNumber")
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

// SetOSCountryCode sets the value of OSCountryCode for the instance
func (instance *MsftUal_SystemId) SetPropertyOSCountryCode(value string) (err error) {
	return instance.SetProperty("OSCountryCode", (value))
}

// GetOSCountryCode gets the value of OSCountryCode for the instance
func (instance *MsftUal_SystemId) GetPropertyOSCountryCode() (value string, err error) {
	retValue, err := instance.GetProperty("OSCountryCode")
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

// SetOSCurrentTimeZone sets the value of OSCurrentTimeZone for the instance
func (instance *MsftUal_SystemId) SetPropertyOSCurrentTimeZone(value int16) (err error) {
	return instance.SetProperty("OSCurrentTimeZone", (value))
}

// GetOSCurrentTimeZone gets the value of OSCurrentTimeZone for the instance
func (instance *MsftUal_SystemId) GetPropertyOSCurrentTimeZone() (value int16, err error) {
	retValue, err := instance.GetProperty("OSCurrentTimeZone")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int16(valuetmp)

	return
}

// SetOSDaylightInEffect sets the value of OSDaylightInEffect for the instance
func (instance *MsftUal_SystemId) SetPropertyOSDaylightInEffect(value bool) (err error) {
	return instance.SetProperty("OSDaylightInEffect", (value))
}

// GetOSDaylightInEffect gets the value of OSDaylightInEffect for the instance
func (instance *MsftUal_SystemId) GetPropertyOSDaylightInEffect() (value bool, err error) {
	retValue, err := instance.GetProperty("OSDaylightInEffect")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetOSLastBootUpTime sets the value of OSLastBootUpTime for the instance
func (instance *MsftUal_SystemId) SetPropertyOSLastBootUpTime(value string) (err error) {
	return instance.SetProperty("OSLastBootUpTime", (value))
}

// GetOSLastBootUpTime gets the value of OSLastBootUpTime for the instance
func (instance *MsftUal_SystemId) GetPropertyOSLastBootUpTime() (value string, err error) {
	retValue, err := instance.GetProperty("OSLastBootUpTime")
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

// SetOSMajor sets the value of OSMajor for the instance
func (instance *MsftUal_SystemId) SetPropertyOSMajor(value uint32) (err error) {
	return instance.SetProperty("OSMajor", (value))
}

// GetOSMajor gets the value of OSMajor for the instance
func (instance *MsftUal_SystemId) GetPropertyOSMajor() (value uint32, err error) {
	retValue, err := instance.GetProperty("OSMajor")
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

// SetOSMinor sets the value of OSMinor for the instance
func (instance *MsftUal_SystemId) SetPropertyOSMinor(value uint32) (err error) {
	return instance.SetProperty("OSMinor", (value))
}

// GetOSMinor gets the value of OSMinor for the instance
func (instance *MsftUal_SystemId) GetPropertyOSMinor() (value uint32, err error) {
	retValue, err := instance.GetProperty("OSMinor")
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

// SetOSPlatformId sets the value of OSPlatformId for the instance
func (instance *MsftUal_SystemId) SetPropertyOSPlatformId(value uint32) (err error) {
	return instance.SetProperty("OSPlatformId", (value))
}

// GetOSPlatformId gets the value of OSPlatformId for the instance
func (instance *MsftUal_SystemId) GetPropertyOSPlatformId() (value uint32, err error) {
	retValue, err := instance.GetProperty("OSPlatformId")
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

// SetOSProductType sets the value of OSProductType for the instance
func (instance *MsftUal_SystemId) SetPropertyOSProductType(value uint32) (err error) {
	return instance.SetProperty("OSProductType", (value))
}

// GetOSProductType gets the value of OSProductType for the instance
func (instance *MsftUal_SystemId) GetPropertyOSProductType() (value uint32, err error) {
	retValue, err := instance.GetProperty("OSProductType")
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

// SetOSSerialNumber sets the value of OSSerialNumber for the instance
func (instance *MsftUal_SystemId) SetPropertyOSSerialNumber(value string) (err error) {
	return instance.SetProperty("OSSerialNumber", (value))
}

// GetOSSerialNumber gets the value of OSSerialNumber for the instance
func (instance *MsftUal_SystemId) GetPropertyOSSerialNumber() (value string, err error) {
	retValue, err := instance.GetProperty("OSSerialNumber")
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

// SetOSSuiteMask sets the value of OSSuiteMask for the instance
func (instance *MsftUal_SystemId) SetPropertyOSSuiteMask(value uint32) (err error) {
	return instance.SetProperty("OSSuiteMask", (value))
}

// GetOSSuiteMask gets the value of OSSuiteMask for the instance
func (instance *MsftUal_SystemId) GetPropertyOSSuiteMask() (value uint32, err error) {
	retValue, err := instance.GetProperty("OSSuiteMask")
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

// SetPhysicalProcessorCount sets the value of PhysicalProcessorCount for the instance
func (instance *MsftUal_SystemId) SetPropertyPhysicalProcessorCount(value uint32) (err error) {
	return instance.SetProperty("PhysicalProcessorCount", (value))
}

// GetPhysicalProcessorCount gets the value of PhysicalProcessorCount for the instance
func (instance *MsftUal_SystemId) GetPropertyPhysicalProcessorCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("PhysicalProcessorCount")
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

// SetServicePackMajor sets the value of ServicePackMajor for the instance
func (instance *MsftUal_SystemId) SetPropertyServicePackMajor(value uint32) (err error) {
	return instance.SetProperty("ServicePackMajor", (value))
}

// GetServicePackMajor gets the value of ServicePackMajor for the instance
func (instance *MsftUal_SystemId) GetPropertyServicePackMajor() (value uint32, err error) {
	retValue, err := instance.GetProperty("ServicePackMajor")
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

// SetServicePackMinor sets the value of ServicePackMinor for the instance
func (instance *MsftUal_SystemId) SetPropertyServicePackMinor(value uint32) (err error) {
	return instance.SetProperty("ServicePackMinor", (value))
}

// GetServicePackMinor gets the value of ServicePackMinor for the instance
func (instance *MsftUal_SystemId) GetPropertyServicePackMinor() (value uint32, err error) {
	retValue, err := instance.GetProperty("ServicePackMinor")
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

// SetSystemDNSHostName sets the value of SystemDNSHostName for the instance
func (instance *MsftUal_SystemId) SetPropertySystemDNSHostName(value string) (err error) {
	return instance.SetProperty("SystemDNSHostName", (value))
}

// GetSystemDNSHostName gets the value of SystemDNSHostName for the instance
func (instance *MsftUal_SystemId) GetPropertySystemDNSHostName() (value string, err error) {
	retValue, err := instance.GetProperty("SystemDNSHostName")
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

// SetSystemDomainName sets the value of SystemDomainName for the instance
func (instance *MsftUal_SystemId) SetPropertySystemDomainName(value string) (err error) {
	return instance.SetProperty("SystemDomainName", (value))
}

// GetSystemDomainName gets the value of SystemDomainName for the instance
func (instance *MsftUal_SystemId) GetPropertySystemDomainName() (value string, err error) {
	retValue, err := instance.GetProperty("SystemDomainName")
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

// SetSystemManufacturer sets the value of SystemManufacturer for the instance
func (instance *MsftUal_SystemId) SetPropertySystemManufacturer(value string) (err error) {
	return instance.SetProperty("SystemManufacturer", (value))
}

// GetSystemManufacturer gets the value of SystemManufacturer for the instance
func (instance *MsftUal_SystemId) GetPropertySystemManufacturer() (value string, err error) {
	retValue, err := instance.GetProperty("SystemManufacturer")
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

// SetSystemProductName sets the value of SystemProductName for the instance
func (instance *MsftUal_SystemId) SetPropertySystemProductName(value string) (err error) {
	return instance.SetProperty("SystemProductName", (value))
}

// GetSystemProductName gets the value of SystemProductName for the instance
func (instance *MsftUal_SystemId) GetPropertySystemProductName() (value string, err error) {
	retValue, err := instance.GetProperty("SystemProductName")
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

// SetSystemSerialNumber sets the value of SystemSerialNumber for the instance
func (instance *MsftUal_SystemId) SetPropertySystemSerialNumber(value string) (err error) {
	return instance.SetProperty("SystemSerialNumber", (value))
}

// GetSystemSerialNumber gets the value of SystemSerialNumber for the instance
func (instance *MsftUal_SystemId) GetPropertySystemSerialNumber() (value string, err error) {
	retValue, err := instance.GetProperty("SystemSerialNumber")
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

// SetSystemSMBIOSUUID sets the value of SystemSMBIOSUUID for the instance
func (instance *MsftUal_SystemId) SetPropertySystemSMBIOSUUID(value string) (err error) {
	return instance.SetProperty("SystemSMBIOSUUID", (value))
}

// GetSystemSMBIOSUUID gets the value of SystemSMBIOSUUID for the instance
func (instance *MsftUal_SystemId) GetPropertySystemSMBIOSUUID() (value string, err error) {
	retValue, err := instance.GetProperty("SystemSMBIOSUUID")
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
