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

// MS_SM_AdapterInformationQuery struct
type MS_SM_AdapterInformationQuery struct { 
	*cim.WmiInstance

	// 
	Active bool

	// 
	DriverName string

	// 
	DriverVersion string

	// 
	FirmwareVersion string

	// 
	HardwareVersion string

	// 
	HBAStatus uint32

	// 
	HBASymbolicName string

	// 
	InstanceName string

	// 
	Manufacturer string

	// 
	MfgDomain string

	// 
	Model string

	// 
	ModelDescription string

	// 
	NumberOfPorts uint32

	// 
	OptionROMVersion string

	// 
	RedundantFirmwareVersion string

	// 
	RedundantOptionROMVersion string

	// 
	SerialNumber string

	// 
	UniqueAdapterId uint64

	// 
	VendorSpecificID uint32
}

	func NewMS_SM_AdapterInformationQueryEx1(instance *cim.WmiInstance) (newInstance *MS_SM_AdapterInformationQuery, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MS_SM_AdapterInformationQuery {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMS_SM_AdapterInformationQueryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MS_SM_AdapterInformationQuery, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MS_SM_AdapterInformationQuery {
	WmiInstance: tmp,
	}
	return
	}
	

// SetActive sets the value of Active for the instance
func (instance *MS_SM_AdapterInformationQuery) SetPropertyActive(value bool) (err error) { 
    return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MS_SM_AdapterInformationQuery) GetPropertyActive()(value bool, err error) { 
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

// SetDriverName sets the value of DriverName for the instance
func (instance *MS_SM_AdapterInformationQuery) SetPropertyDriverName(value string) (err error) { 
    return instance.SetProperty("DriverName", (value))
}

// GetDriverName gets the value of DriverName for the instance
func (instance *MS_SM_AdapterInformationQuery) GetPropertyDriverName()(value string, err error) { 
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

// SetDriverVersion sets the value of DriverVersion for the instance
func (instance *MS_SM_AdapterInformationQuery) SetPropertyDriverVersion(value string) (err error) { 
    return instance.SetProperty("DriverVersion", (value))
}

// GetDriverVersion gets the value of DriverVersion for the instance
func (instance *MS_SM_AdapterInformationQuery) GetPropertyDriverVersion()(value string, err error) { 
    retValue, err := instance.GetProperty("DriverVersion")
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
func (instance *MS_SM_AdapterInformationQuery) SetPropertyFirmwareVersion(value string) (err error) { 
    return instance.SetProperty("FirmwareVersion", (value))
}

// GetFirmwareVersion gets the value of FirmwareVersion for the instance
func (instance *MS_SM_AdapterInformationQuery) GetPropertyFirmwareVersion()(value string, err error) { 
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

// SetHardwareVersion sets the value of HardwareVersion for the instance
func (instance *MS_SM_AdapterInformationQuery) SetPropertyHardwareVersion(value string) (err error) { 
    return instance.SetProperty("HardwareVersion", (value))
}

// GetHardwareVersion gets the value of HardwareVersion for the instance
func (instance *MS_SM_AdapterInformationQuery) GetPropertyHardwareVersion()(value string, err error) { 
    retValue, err := instance.GetProperty("HardwareVersion")
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

// SetHBAStatus sets the value of HBAStatus for the instance
func (instance *MS_SM_AdapterInformationQuery) SetPropertyHBAStatus(value uint32) (err error) { 
    return instance.SetProperty("HBAStatus", (value))
}

// GetHBAStatus gets the value of HBAStatus for the instance
func (instance *MS_SM_AdapterInformationQuery) GetPropertyHBAStatus()(value uint32, err error) { 
    retValue, err := instance.GetProperty("HBAStatus")
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

// SetHBASymbolicName sets the value of HBASymbolicName for the instance
func (instance *MS_SM_AdapterInformationQuery) SetPropertyHBASymbolicName(value string) (err error) { 
    return instance.SetProperty("HBASymbolicName", (value))
}

// GetHBASymbolicName gets the value of HBASymbolicName for the instance
func (instance *MS_SM_AdapterInformationQuery) GetPropertyHBASymbolicName()(value string, err error) { 
    retValue, err := instance.GetProperty("HBASymbolicName")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MS_SM_AdapterInformationQuery) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MS_SM_AdapterInformationQuery) GetPropertyInstanceName()(value string, err error) { 
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

// SetManufacturer sets the value of Manufacturer for the instance
func (instance *MS_SM_AdapterInformationQuery) SetPropertyManufacturer(value string) (err error) { 
    return instance.SetProperty("Manufacturer", (value))
}

// GetManufacturer gets the value of Manufacturer for the instance
func (instance *MS_SM_AdapterInformationQuery) GetPropertyManufacturer()(value string, err error) { 
    retValue, err := instance.GetProperty("Manufacturer")
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

// SetMfgDomain sets the value of MfgDomain for the instance
func (instance *MS_SM_AdapterInformationQuery) SetPropertyMfgDomain(value string) (err error) { 
    return instance.SetProperty("MfgDomain", (value))
}

// GetMfgDomain gets the value of MfgDomain for the instance
func (instance *MS_SM_AdapterInformationQuery) GetPropertyMfgDomain()(value string, err error) { 
    retValue, err := instance.GetProperty("MfgDomain")
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

// SetModel sets the value of Model for the instance
func (instance *MS_SM_AdapterInformationQuery) SetPropertyModel(value string) (err error) { 
    return instance.SetProperty("Model", (value))
}

// GetModel gets the value of Model for the instance
func (instance *MS_SM_AdapterInformationQuery) GetPropertyModel()(value string, err error) { 
    retValue, err := instance.GetProperty("Model")
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

// SetModelDescription sets the value of ModelDescription for the instance
func (instance *MS_SM_AdapterInformationQuery) SetPropertyModelDescription(value string) (err error) { 
    return instance.SetProperty("ModelDescription", (value))
}

// GetModelDescription gets the value of ModelDescription for the instance
func (instance *MS_SM_AdapterInformationQuery) GetPropertyModelDescription()(value string, err error) { 
    retValue, err := instance.GetProperty("ModelDescription")
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

// SetNumberOfPorts sets the value of NumberOfPorts for the instance
func (instance *MS_SM_AdapterInformationQuery) SetPropertyNumberOfPorts(value uint32) (err error) { 
    return instance.SetProperty("NumberOfPorts", (value))
}

// GetNumberOfPorts gets the value of NumberOfPorts for the instance
func (instance *MS_SM_AdapterInformationQuery) GetPropertyNumberOfPorts()(value uint32, err error) { 
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

// SetOptionROMVersion sets the value of OptionROMVersion for the instance
func (instance *MS_SM_AdapterInformationQuery) SetPropertyOptionROMVersion(value string) (err error) { 
    return instance.SetProperty("OptionROMVersion", (value))
}

// GetOptionROMVersion gets the value of OptionROMVersion for the instance
func (instance *MS_SM_AdapterInformationQuery) GetPropertyOptionROMVersion()(value string, err error) { 
    retValue, err := instance.GetProperty("OptionROMVersion")
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

// SetRedundantFirmwareVersion sets the value of RedundantFirmwareVersion for the instance
func (instance *MS_SM_AdapterInformationQuery) SetPropertyRedundantFirmwareVersion(value string) (err error) { 
    return instance.SetProperty("RedundantFirmwareVersion", (value))
}

// GetRedundantFirmwareVersion gets the value of RedundantFirmwareVersion for the instance
func (instance *MS_SM_AdapterInformationQuery) GetPropertyRedundantFirmwareVersion()(value string, err error) { 
    retValue, err := instance.GetProperty("RedundantFirmwareVersion")
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

// SetRedundantOptionROMVersion sets the value of RedundantOptionROMVersion for the instance
func (instance *MS_SM_AdapterInformationQuery) SetPropertyRedundantOptionROMVersion(value string) (err error) { 
    return instance.SetProperty("RedundantOptionROMVersion", (value))
}

// GetRedundantOptionROMVersion gets the value of RedundantOptionROMVersion for the instance
func (instance *MS_SM_AdapterInformationQuery) GetPropertyRedundantOptionROMVersion()(value string, err error) { 
    retValue, err := instance.GetProperty("RedundantOptionROMVersion")
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

// SetSerialNumber sets the value of SerialNumber for the instance
func (instance *MS_SM_AdapterInformationQuery) SetPropertySerialNumber(value string) (err error) { 
    return instance.SetProperty("SerialNumber", (value))
}

// GetSerialNumber gets the value of SerialNumber for the instance
func (instance *MS_SM_AdapterInformationQuery) GetPropertySerialNumber()(value string, err error) { 
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

// SetUniqueAdapterId sets the value of UniqueAdapterId for the instance
func (instance *MS_SM_AdapterInformationQuery) SetPropertyUniqueAdapterId(value uint64) (err error) { 
    return instance.SetProperty("UniqueAdapterId", (value))
}

// GetUniqueAdapterId gets the value of UniqueAdapterId for the instance
func (instance *MS_SM_AdapterInformationQuery) GetPropertyUniqueAdapterId()(value uint64, err error) { 
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

// SetVendorSpecificID sets the value of VendorSpecificID for the instance
func (instance *MS_SM_AdapterInformationQuery) SetPropertyVendorSpecificID(value uint32) (err error) { 
    return instance.SetProperty("VendorSpecificID", (value))
}

// GetVendorSpecificID gets the value of VendorSpecificID for the instance
func (instance *MS_SM_AdapterInformationQuery) GetPropertyVendorSpecificID()(value uint32, err error) { 
    retValue, err := instance.GetProperty("VendorSpecificID")
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

