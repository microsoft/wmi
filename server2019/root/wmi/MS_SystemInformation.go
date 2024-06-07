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

// MS_SystemInformation struct
type MS_SystemInformation struct { 
	*cim.WmiInstance

	// 
	Active bool

	// 
	BaseBoardManufacturer string

	// 
	BaseBoardProduct string

	// 
	BaseBoardVersion string

	// 
	BiosMajorRelease uint8

	// 
	BiosMinorRelease uint8

	// 
	BIOSReleaseDate string

	// 
	BIOSVendor string

	// 
	BIOSVersion string

	// 
	ECFirmwareMajorRelease uint8

	// 
	ECFirmwareMinorRelease uint8

	// 
	InstanceName string

	// 
	SystemFamily string

	// 
	SystemManufacturer string

	// 
	SystemProductName string

	// 
	SystemSKU string

	// 
	SystemVersion string
}

	func NewMS_SystemInformationEx1(instance *cim.WmiInstance) (newInstance *MS_SystemInformation, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MS_SystemInformation {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMS_SystemInformationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MS_SystemInformation, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MS_SystemInformation {
	WmiInstance: tmp,
	}
	return
	}
	

// SetActive sets the value of Active for the instance
func (instance *MS_SystemInformation) SetPropertyActive(value bool) (err error) { 
    return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MS_SystemInformation) GetPropertyActive()(value bool, err error) { 
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

// SetBaseBoardManufacturer sets the value of BaseBoardManufacturer for the instance
func (instance *MS_SystemInformation) SetPropertyBaseBoardManufacturer(value string) (err error) { 
    return instance.SetProperty("BaseBoardManufacturer", (value))
}

// GetBaseBoardManufacturer gets the value of BaseBoardManufacturer for the instance
func (instance *MS_SystemInformation) GetPropertyBaseBoardManufacturer()(value string, err error) { 
    retValue, err := instance.GetProperty("BaseBoardManufacturer")
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

// SetBaseBoardProduct sets the value of BaseBoardProduct for the instance
func (instance *MS_SystemInformation) SetPropertyBaseBoardProduct(value string) (err error) { 
    return instance.SetProperty("BaseBoardProduct", (value))
}

// GetBaseBoardProduct gets the value of BaseBoardProduct for the instance
func (instance *MS_SystemInformation) GetPropertyBaseBoardProduct()(value string, err error) { 
    retValue, err := instance.GetProperty("BaseBoardProduct")
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

// SetBaseBoardVersion sets the value of BaseBoardVersion for the instance
func (instance *MS_SystemInformation) SetPropertyBaseBoardVersion(value string) (err error) { 
    return instance.SetProperty("BaseBoardVersion", (value))
}

// GetBaseBoardVersion gets the value of BaseBoardVersion for the instance
func (instance *MS_SystemInformation) GetPropertyBaseBoardVersion()(value string, err error) { 
    retValue, err := instance.GetProperty("BaseBoardVersion")
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

// SetBiosMajorRelease sets the value of BiosMajorRelease for the instance
func (instance *MS_SystemInformation) SetPropertyBiosMajorRelease(value uint8) (err error) { 
    return instance.SetProperty("BiosMajorRelease", (value))
}

// GetBiosMajorRelease gets the value of BiosMajorRelease for the instance
func (instance *MS_SystemInformation) GetPropertyBiosMajorRelease()(value uint8, err error) { 
    retValue, err := instance.GetProperty("BiosMajorRelease")
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

// SetBiosMinorRelease sets the value of BiosMinorRelease for the instance
func (instance *MS_SystemInformation) SetPropertyBiosMinorRelease(value uint8) (err error) { 
    return instance.SetProperty("BiosMinorRelease", (value))
}

// GetBiosMinorRelease gets the value of BiosMinorRelease for the instance
func (instance *MS_SystemInformation) GetPropertyBiosMinorRelease()(value uint8, err error) { 
    retValue, err := instance.GetProperty("BiosMinorRelease")
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

// SetBIOSReleaseDate sets the value of BIOSReleaseDate for the instance
func (instance *MS_SystemInformation) SetPropertyBIOSReleaseDate(value string) (err error) { 
    return instance.SetProperty("BIOSReleaseDate", (value))
}

// GetBIOSReleaseDate gets the value of BIOSReleaseDate for the instance
func (instance *MS_SystemInformation) GetPropertyBIOSReleaseDate()(value string, err error) { 
    retValue, err := instance.GetProperty("BIOSReleaseDate")
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

// SetBIOSVendor sets the value of BIOSVendor for the instance
func (instance *MS_SystemInformation) SetPropertyBIOSVendor(value string) (err error) { 
    return instance.SetProperty("BIOSVendor", (value))
}

// GetBIOSVendor gets the value of BIOSVendor for the instance
func (instance *MS_SystemInformation) GetPropertyBIOSVendor()(value string, err error) { 
    retValue, err := instance.GetProperty("BIOSVendor")
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

// SetBIOSVersion sets the value of BIOSVersion for the instance
func (instance *MS_SystemInformation) SetPropertyBIOSVersion(value string) (err error) { 
    return instance.SetProperty("BIOSVersion", (value))
}

// GetBIOSVersion gets the value of BIOSVersion for the instance
func (instance *MS_SystemInformation) GetPropertyBIOSVersion()(value string, err error) { 
    retValue, err := instance.GetProperty("BIOSVersion")
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

// SetECFirmwareMajorRelease sets the value of ECFirmwareMajorRelease for the instance
func (instance *MS_SystemInformation) SetPropertyECFirmwareMajorRelease(value uint8) (err error) { 
    return instance.SetProperty("ECFirmwareMajorRelease", (value))
}

// GetECFirmwareMajorRelease gets the value of ECFirmwareMajorRelease for the instance
func (instance *MS_SystemInformation) GetPropertyECFirmwareMajorRelease()(value uint8, err error) { 
    retValue, err := instance.GetProperty("ECFirmwareMajorRelease")
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

// SetECFirmwareMinorRelease sets the value of ECFirmwareMinorRelease for the instance
func (instance *MS_SystemInformation) SetPropertyECFirmwareMinorRelease(value uint8) (err error) { 
    return instance.SetProperty("ECFirmwareMinorRelease", (value))
}

// GetECFirmwareMinorRelease gets the value of ECFirmwareMinorRelease for the instance
func (instance *MS_SystemInformation) GetPropertyECFirmwareMinorRelease()(value uint8, err error) { 
    retValue, err := instance.GetProperty("ECFirmwareMinorRelease")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MS_SystemInformation) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MS_SystemInformation) GetPropertyInstanceName()(value string, err error) { 
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

// SetSystemFamily sets the value of SystemFamily for the instance
func (instance *MS_SystemInformation) SetPropertySystemFamily(value string) (err error) { 
    return instance.SetProperty("SystemFamily", (value))
}

// GetSystemFamily gets the value of SystemFamily for the instance
func (instance *MS_SystemInformation) GetPropertySystemFamily()(value string, err error) { 
    retValue, err := instance.GetProperty("SystemFamily")
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

// SetSystemManufacturer sets the value of SystemManufacturer for the instance
func (instance *MS_SystemInformation) SetPropertySystemManufacturer(value string) (err error) { 
    return instance.SetProperty("SystemManufacturer", (value))
}

// GetSystemManufacturer gets the value of SystemManufacturer for the instance
func (instance *MS_SystemInformation) GetPropertySystemManufacturer()(value string, err error) { 
    retValue, err := instance.GetProperty("SystemManufacturer")
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

// SetSystemProductName sets the value of SystemProductName for the instance
func (instance *MS_SystemInformation) SetPropertySystemProductName(value string) (err error) { 
    return instance.SetProperty("SystemProductName", (value))
}

// GetSystemProductName gets the value of SystemProductName for the instance
func (instance *MS_SystemInformation) GetPropertySystemProductName()(value string, err error) { 
    retValue, err := instance.GetProperty("SystemProductName")
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

// SetSystemSKU sets the value of SystemSKU for the instance
func (instance *MS_SystemInformation) SetPropertySystemSKU(value string) (err error) { 
    return instance.SetProperty("SystemSKU", (value))
}

// GetSystemSKU gets the value of SystemSKU for the instance
func (instance *MS_SystemInformation) GetPropertySystemSKU()(value string, err error) { 
    retValue, err := instance.GetProperty("SystemSKU")
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

// SetSystemVersion sets the value of SystemVersion for the instance
func (instance *MS_SystemInformation) SetPropertySystemVersion(value string) (err error) { 
    return instance.SetProperty("SystemVersion", (value))
}

// GetSystemVersion gets the value of SystemVersion for the instance
func (instance *MS_SystemInformation) GetPropertySystemVersion()(value string, err error) { 
    retValue, err := instance.GetProperty("SystemVersion")
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

