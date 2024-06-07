// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSiSCSI_HBAInformation struct
type MSiSCSI_HBAInformation struct { 
	*cim.WmiInstance

	// 
	Active bool

	// A text string set by the manufacturer describing the Asic version
	AsicVersion string

	// TRUE if Bi-directionsal SCSI comamnd supported
	BiDiScsiCommands bool

	// TRUE if the adapter caches are valid
	CacheValid bool

	// A text string specifying the name of the driver for the adapter
	DriverName string

	// A text string set by the manufacturer describing the firmware version of adapter
	FirmwareVersion string

	// **typedef** Bit flags that indicate various functionality supported
	FunctionalitySupported uint32

	// This is the GUID value last set by the SetGenerationalGuid method in the MSiSCSI_Operations class.
	GenerationalGuid []uint8

	// 
	InstanceName string

	// TRUE if TCP/IP traffic is integrated with the Windows networking TCP/IP stack via a software only initiator. An adapter with its own TCP/IP stack would set this to FALSE.
	IntegratedTCPIP bool

	// Maxumum CDB length supported by the adapter
	MaxCDBLength uint32

	// TRUE if this adapter is a multifunction device, that is it also exposes a netcard interface
	MultifunctionDevice bool

	// Number of ports (or TCP/IP addresses) on the adapter
	NumberOfPorts uint32

	// A text string set by the manufacturer describing the option rom version of adapter
	OptionRomVersion string

	// If TRUE the iSCSI Initiator service will perform any DNS lookup and pass binary IP addresses to the adapter; the adapter must be on the same network as the Windows TCP/IP stack. If FALSE then DNS must be available on adapter.
	RequiresBinaryIpAddresses bool

	// A text string set by the manufacturer describing the serial number of adapter
	SerialNumber string

	// **typedef** Current status of adapter
	Status HBAInformation_Status

	// Id that is globally unique for all instances of iSCSI initiators. Use the address of the Adapter Extension or another address owned by the device driver.
	UniqueAdapterId uint64

	// A text string describing the manufacturer of adapter
	VendorID string

	// A text string set by the manufacturer describing the model of adapter
	VendorModel string

	// A text string set by the manufacturer describing the version of adapter
	VendorVersion string

	// Maximum version number of the iSCSI spec supported by adapter
	VersionMax uint8

	// Minimum version number of the iScsi spec supported by adapter
	VersionMin uint8
}

	func NewMSiSCSI_HBAInformationEx1(instance *cim.WmiInstance) (newInstance *MSiSCSI_HBAInformation, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MSiSCSI_HBAInformation {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMSiSCSI_HBAInformationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSiSCSI_HBAInformation, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSiSCSI_HBAInformation {
	WmiInstance: tmp,
	}
	return
	}
	

// SetActive sets the value of Active for the instance
func (instance *MSiSCSI_HBAInformation) SetPropertyActive(value bool) (err error) { 
    return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSiSCSI_HBAInformation) GetPropertyActive()(value bool, err error) { 
    retValue, err := instance.GetProperty("Active")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetAsicVersion sets the value of AsicVersion for the instance
func (instance *MSiSCSI_HBAInformation) SetPropertyAsicVersion(value string) (err error) { 
    return instance.SetProperty("AsicVersion", (value))
}

// GetAsicVersion gets the value of AsicVersion for the instance
func (instance *MSiSCSI_HBAInformation) GetPropertyAsicVersion()(value string, err error) { 
    retValue, err := instance.GetProperty("AsicVersion")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetBiDiScsiCommands sets the value of BiDiScsiCommands for the instance
func (instance *MSiSCSI_HBAInformation) SetPropertyBiDiScsiCommands(value bool) (err error) { 
    return instance.SetProperty("BiDiScsiCommands", (value))
}

// GetBiDiScsiCommands gets the value of BiDiScsiCommands for the instance
func (instance *MSiSCSI_HBAInformation) GetPropertyBiDiScsiCommands()(value bool, err error) { 
    retValue, err := instance.GetProperty("BiDiScsiCommands")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetCacheValid sets the value of CacheValid for the instance
func (instance *MSiSCSI_HBAInformation) SetPropertyCacheValid(value bool) (err error) { 
    return instance.SetProperty("CacheValid", (value))
}

// GetCacheValid gets the value of CacheValid for the instance
func (instance *MSiSCSI_HBAInformation) GetPropertyCacheValid()(value bool, err error) { 
    retValue, err := instance.GetProperty("CacheValid")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetDriverName sets the value of DriverName for the instance
func (instance *MSiSCSI_HBAInformation) SetPropertyDriverName(value string) (err error) { 
    return instance.SetProperty("DriverName", (value))
}

// GetDriverName gets the value of DriverName for the instance
func (instance *MSiSCSI_HBAInformation) GetPropertyDriverName()(value string, err error) { 
    retValue, err := instance.GetProperty("DriverName")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetFirmwareVersion sets the value of FirmwareVersion for the instance
func (instance *MSiSCSI_HBAInformation) SetPropertyFirmwareVersion(value string) (err error) { 
    return instance.SetProperty("FirmwareVersion", (value))
}

// GetFirmwareVersion gets the value of FirmwareVersion for the instance
func (instance *MSiSCSI_HBAInformation) GetPropertyFirmwareVersion()(value string, err error) { 
    retValue, err := instance.GetProperty("FirmwareVersion")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetFunctionalitySupported sets the value of FunctionalitySupported for the instance
func (instance *MSiSCSI_HBAInformation) SetPropertyFunctionalitySupported(value uint32) (err error) { 
    return instance.SetProperty("FunctionalitySupported", (value))
}

// GetFunctionalitySupported gets the value of FunctionalitySupported for the instance
func (instance *MSiSCSI_HBAInformation) GetPropertyFunctionalitySupported()(value uint32, err error) { 
    retValue, err := instance.GetProperty("FunctionalitySupported")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetGenerationalGuid sets the value of GenerationalGuid for the instance
func (instance *MSiSCSI_HBAInformation) SetPropertyGenerationalGuid(value []uint8) (err error) { 
    return instance.SetProperty("GenerationalGuid", (value))
}

// GetGenerationalGuid gets the value of GenerationalGuid for the instance
func (instance *MSiSCSI_HBAInformation) GetPropertyGenerationalGuid()(value []uint8, err error) { 
    retValue, err := instance.GetProperty("GenerationalGuid")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(uint8); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, uint8(valuetmp))
    }

    return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSiSCSI_HBAInformation) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSiSCSI_HBAInformation) GetPropertyInstanceName()(value string, err error) { 
    retValue, err := instance.GetProperty("InstanceName")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetIntegratedTCPIP sets the value of IntegratedTCPIP for the instance
func (instance *MSiSCSI_HBAInformation) SetPropertyIntegratedTCPIP(value bool) (err error) { 
    return instance.SetProperty("IntegratedTCPIP", (value))
}

// GetIntegratedTCPIP gets the value of IntegratedTCPIP for the instance
func (instance *MSiSCSI_HBAInformation) GetPropertyIntegratedTCPIP()(value bool, err error) { 
    retValue, err := instance.GetProperty("IntegratedTCPIP")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetMaxCDBLength sets the value of MaxCDBLength for the instance
func (instance *MSiSCSI_HBAInformation) SetPropertyMaxCDBLength(value uint32) (err error) { 
    return instance.SetProperty("MaxCDBLength", (value))
}

// GetMaxCDBLength gets the value of MaxCDBLength for the instance
func (instance *MSiSCSI_HBAInformation) GetPropertyMaxCDBLength()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MaxCDBLength")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetMultifunctionDevice sets the value of MultifunctionDevice for the instance
func (instance *MSiSCSI_HBAInformation) SetPropertyMultifunctionDevice(value bool) (err error) { 
    return instance.SetProperty("MultifunctionDevice", (value))
}

// GetMultifunctionDevice gets the value of MultifunctionDevice for the instance
func (instance *MSiSCSI_HBAInformation) GetPropertyMultifunctionDevice()(value bool, err error) { 
    retValue, err := instance.GetProperty("MultifunctionDevice")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetNumberOfPorts sets the value of NumberOfPorts for the instance
func (instance *MSiSCSI_HBAInformation) SetPropertyNumberOfPorts(value uint32) (err error) { 
    return instance.SetProperty("NumberOfPorts", (value))
}

// GetNumberOfPorts gets the value of NumberOfPorts for the instance
func (instance *MSiSCSI_HBAInformation) GetPropertyNumberOfPorts()(value uint32, err error) { 
    retValue, err := instance.GetProperty("NumberOfPorts")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetOptionRomVersion sets the value of OptionRomVersion for the instance
func (instance *MSiSCSI_HBAInformation) SetPropertyOptionRomVersion(value string) (err error) { 
    return instance.SetProperty("OptionRomVersion", (value))
}

// GetOptionRomVersion gets the value of OptionRomVersion for the instance
func (instance *MSiSCSI_HBAInformation) GetPropertyOptionRomVersion()(value string, err error) { 
    retValue, err := instance.GetProperty("OptionRomVersion")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetRequiresBinaryIpAddresses sets the value of RequiresBinaryIpAddresses for the instance
func (instance *MSiSCSI_HBAInformation) SetPropertyRequiresBinaryIpAddresses(value bool) (err error) { 
    return instance.SetProperty("RequiresBinaryIpAddresses", (value))
}

// GetRequiresBinaryIpAddresses gets the value of RequiresBinaryIpAddresses for the instance
func (instance *MSiSCSI_HBAInformation) GetPropertyRequiresBinaryIpAddresses()(value bool, err error) { 
    retValue, err := instance.GetProperty("RequiresBinaryIpAddresses")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetSerialNumber sets the value of SerialNumber for the instance
func (instance *MSiSCSI_HBAInformation) SetPropertySerialNumber(value string) (err error) { 
    return instance.SetProperty("SerialNumber", (value))
}

// GetSerialNumber gets the value of SerialNumber for the instance
func (instance *MSiSCSI_HBAInformation) GetPropertySerialNumber()(value string, err error) { 
    retValue, err := instance.GetProperty("SerialNumber")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetStatus sets the value of Status for the instance
func (instance *MSiSCSI_HBAInformation) SetPropertyStatus(value HBAInformation_Status) (err error) { 
    return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *MSiSCSI_HBAInformation) GetPropertyStatus()(value HBAInformation_Status, err error) { 
    retValue, err := instance.GetProperty("Status")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = HBAInformation_Status(valuetmp)

    return
}

// SetUniqueAdapterId sets the value of UniqueAdapterId for the instance
func (instance *MSiSCSI_HBAInformation) SetPropertyUniqueAdapterId(value uint64) (err error) { 
    return instance.SetProperty("UniqueAdapterId", (value))
}

// GetUniqueAdapterId gets the value of UniqueAdapterId for the instance
func (instance *MSiSCSI_HBAInformation) GetPropertyUniqueAdapterId()(value uint64, err error) { 
    retValue, err := instance.GetProperty("UniqueAdapterId")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetVendorID sets the value of VendorID for the instance
func (instance *MSiSCSI_HBAInformation) SetPropertyVendorID(value string) (err error) { 
    return instance.SetProperty("VendorID", (value))
}

// GetVendorID gets the value of VendorID for the instance
func (instance *MSiSCSI_HBAInformation) GetPropertyVendorID()(value string, err error) { 
    retValue, err := instance.GetProperty("VendorID")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetVendorModel sets the value of VendorModel for the instance
func (instance *MSiSCSI_HBAInformation) SetPropertyVendorModel(value string) (err error) { 
    return instance.SetProperty("VendorModel", (value))
}

// GetVendorModel gets the value of VendorModel for the instance
func (instance *MSiSCSI_HBAInformation) GetPropertyVendorModel()(value string, err error) { 
    retValue, err := instance.GetProperty("VendorModel")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetVendorVersion sets the value of VendorVersion for the instance
func (instance *MSiSCSI_HBAInformation) SetPropertyVendorVersion(value string) (err error) { 
    return instance.SetProperty("VendorVersion", (value))
}

// GetVendorVersion gets the value of VendorVersion for the instance
func (instance *MSiSCSI_HBAInformation) GetPropertyVendorVersion()(value string, err error) { 
    retValue, err := instance.GetProperty("VendorVersion")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetVersionMax sets the value of VersionMax for the instance
func (instance *MSiSCSI_HBAInformation) SetPropertyVersionMax(value uint8) (err error) { 
    return instance.SetProperty("VersionMax", (value))
}

// GetVersionMax gets the value of VersionMax for the instance
func (instance *MSiSCSI_HBAInformation) GetPropertyVersionMax()(value uint8, err error) { 
    retValue, err := instance.GetProperty("VersionMax")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// SetVersionMin sets the value of VersionMin for the instance
func (instance *MSiSCSI_HBAInformation) SetPropertyVersionMin(value uint8) (err error) { 
    return instance.SetProperty("VersionMin", (value))
}

// GetVersionMin gets the value of VersionMin for the instance
func (instance *MSiSCSI_HBAInformation) GetPropertyVersionMin()(value uint8, err error) { 
    retValue, err := instance.GetProperty("VersionMin")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

