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
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSStorageDriver_ATAPISmartData struct
type MSStorageDriver_ATAPISmartData struct { 
	*MSStorageDriver

	// 
	Active bool

	// 
	Checksum uint8

	// 
	ErrorLogCapability uint8

	// 
	ExtendedPollTimeInMinutes uint8

	// 
	InstanceName string

	// 
	Length uint32

	// 
	OfflineCollectCapability uint8

	// 
	OfflineCollectionStatus ATAPISmartData_OfflineCollectionStatus

	// Reserved
	Reserved []uint8

	// 
	SelfTestStatus ATAPISmartData_SelfTestStatus

	// 
	ShortPollTimeInMinutes uint8

	// 
	SmartCapability uint16

	// 
	TotalTime uint16

	// 
	VendorSpecific []uint8

	// 
	VendorSpecific2 uint8

	// 
	VendorSpecific3 uint8

	// 
	VendorSpecific4 []uint8
}

	func NewMSStorageDriver_ATAPISmartDataEx1(instance *cim.WmiInstance) (newInstance *MSStorageDriver_ATAPISmartData, err error) {tmp, err := NewMSStorageDriverEx1(instance)
		
	if err != nil { return }
	newInstance = &MSStorageDriver_ATAPISmartData {
	MSStorageDriver: tmp,
	}
	return
	}
	

	func NewMSStorageDriver_ATAPISmartDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSStorageDriver_ATAPISmartData, err error) {tmp, err := NewMSStorageDriverEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSStorageDriver_ATAPISmartData {
	MSStorageDriver: tmp,
	}
	return
	}
	

// SetActive sets the value of Active for the instance
func (instance *MSStorageDriver_ATAPISmartData) SetPropertyActive(value bool) (err error) { 
    return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSStorageDriver_ATAPISmartData) GetPropertyActive()(value bool, err error) { 
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

// SetChecksum sets the value of Checksum for the instance
func (instance *MSStorageDriver_ATAPISmartData) SetPropertyChecksum(value uint8) (err error) { 
    return instance.SetProperty("Checksum", (value))
}

// GetChecksum gets the value of Checksum for the instance
func (instance *MSStorageDriver_ATAPISmartData) GetPropertyChecksum()(value uint8, err error) { 
    retValue, err := instance.GetProperty("Checksum")
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

// SetErrorLogCapability sets the value of ErrorLogCapability for the instance
func (instance *MSStorageDriver_ATAPISmartData) SetPropertyErrorLogCapability(value uint8) (err error) { 
    return instance.SetProperty("ErrorLogCapability", (value))
}

// GetErrorLogCapability gets the value of ErrorLogCapability for the instance
func (instance *MSStorageDriver_ATAPISmartData) GetPropertyErrorLogCapability()(value uint8, err error) { 
    retValue, err := instance.GetProperty("ErrorLogCapability")
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

// SetExtendedPollTimeInMinutes sets the value of ExtendedPollTimeInMinutes for the instance
func (instance *MSStorageDriver_ATAPISmartData) SetPropertyExtendedPollTimeInMinutes(value uint8) (err error) { 
    return instance.SetProperty("ExtendedPollTimeInMinutes", (value))
}

// GetExtendedPollTimeInMinutes gets the value of ExtendedPollTimeInMinutes for the instance
func (instance *MSStorageDriver_ATAPISmartData) GetPropertyExtendedPollTimeInMinutes()(value uint8, err error) { 
    retValue, err := instance.GetProperty("ExtendedPollTimeInMinutes")
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
func (instance *MSStorageDriver_ATAPISmartData) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSStorageDriver_ATAPISmartData) GetPropertyInstanceName()(value string, err error) { 
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

// SetLength sets the value of Length for the instance
func (instance *MSStorageDriver_ATAPISmartData) SetPropertyLength(value uint32) (err error) { 
    return instance.SetProperty("Length", (value))
}

// GetLength gets the value of Length for the instance
func (instance *MSStorageDriver_ATAPISmartData) GetPropertyLength()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Length")
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

// SetOfflineCollectCapability sets the value of OfflineCollectCapability for the instance
func (instance *MSStorageDriver_ATAPISmartData) SetPropertyOfflineCollectCapability(value uint8) (err error) { 
    return instance.SetProperty("OfflineCollectCapability", (value))
}

// GetOfflineCollectCapability gets the value of OfflineCollectCapability for the instance
func (instance *MSStorageDriver_ATAPISmartData) GetPropertyOfflineCollectCapability()(value uint8, err error) { 
    retValue, err := instance.GetProperty("OfflineCollectCapability")
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

// SetOfflineCollectionStatus sets the value of OfflineCollectionStatus for the instance
func (instance *MSStorageDriver_ATAPISmartData) SetPropertyOfflineCollectionStatus(value ATAPISmartData_OfflineCollectionStatus) (err error) { 
    return instance.SetProperty("OfflineCollectionStatus", (value))
}

// GetOfflineCollectionStatus gets the value of OfflineCollectionStatus for the instance
func (instance *MSStorageDriver_ATAPISmartData) GetPropertyOfflineCollectionStatus()(value ATAPISmartData_OfflineCollectionStatus, err error) { 
    retValue, err := instance.GetProperty("OfflineCollectionStatus")
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

    value = ATAPISmartData_OfflineCollectionStatus(valuetmp)

    return
}

// SetReserved sets the value of Reserved for the instance
func (instance *MSStorageDriver_ATAPISmartData) SetPropertyReserved(value []uint8) (err error) { 
    return instance.SetProperty("Reserved", (value))
}

// GetReserved gets the value of Reserved for the instance
func (instance *MSStorageDriver_ATAPISmartData) GetPropertyReserved()(value []uint8, err error) { 
    retValue, err := instance.GetProperty("Reserved")
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

// SetSelfTestStatus sets the value of SelfTestStatus for the instance
func (instance *MSStorageDriver_ATAPISmartData) SetPropertySelfTestStatus(value ATAPISmartData_SelfTestStatus) (err error) { 
    return instance.SetProperty("SelfTestStatus", (value))
}

// GetSelfTestStatus gets the value of SelfTestStatus for the instance
func (instance *MSStorageDriver_ATAPISmartData) GetPropertySelfTestStatus()(value ATAPISmartData_SelfTestStatus, err error) { 
    retValue, err := instance.GetProperty("SelfTestStatus")
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

    value = ATAPISmartData_SelfTestStatus(valuetmp)

    return
}

// SetShortPollTimeInMinutes sets the value of ShortPollTimeInMinutes for the instance
func (instance *MSStorageDriver_ATAPISmartData) SetPropertyShortPollTimeInMinutes(value uint8) (err error) { 
    return instance.SetProperty("ShortPollTimeInMinutes", (value))
}

// GetShortPollTimeInMinutes gets the value of ShortPollTimeInMinutes for the instance
func (instance *MSStorageDriver_ATAPISmartData) GetPropertyShortPollTimeInMinutes()(value uint8, err error) { 
    retValue, err := instance.GetProperty("ShortPollTimeInMinutes")
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

// SetSmartCapability sets the value of SmartCapability for the instance
func (instance *MSStorageDriver_ATAPISmartData) SetPropertySmartCapability(value uint16) (err error) { 
    return instance.SetProperty("SmartCapability", (value))
}

// GetSmartCapability gets the value of SmartCapability for the instance
func (instance *MSStorageDriver_ATAPISmartData) GetPropertySmartCapability()(value uint16, err error) { 
    retValue, err := instance.GetProperty("SmartCapability")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint16); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint16(valuetmp)

    return
}

// SetTotalTime sets the value of TotalTime for the instance
func (instance *MSStorageDriver_ATAPISmartData) SetPropertyTotalTime(value uint16) (err error) { 
    return instance.SetProperty("TotalTime", (value))
}

// GetTotalTime gets the value of TotalTime for the instance
func (instance *MSStorageDriver_ATAPISmartData) GetPropertyTotalTime()(value uint16, err error) { 
    retValue, err := instance.GetProperty("TotalTime")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint16); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint16(valuetmp)

    return
}

// SetVendorSpecific sets the value of VendorSpecific for the instance
func (instance *MSStorageDriver_ATAPISmartData) SetPropertyVendorSpecific(value []uint8) (err error) { 
    return instance.SetProperty("VendorSpecific", (value))
}

// GetVendorSpecific gets the value of VendorSpecific for the instance
func (instance *MSStorageDriver_ATAPISmartData) GetPropertyVendorSpecific()(value []uint8, err error) { 
    retValue, err := instance.GetProperty("VendorSpecific")
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

// SetVendorSpecific2 sets the value of VendorSpecific2 for the instance
func (instance *MSStorageDriver_ATAPISmartData) SetPropertyVendorSpecific2(value uint8) (err error) { 
    return instance.SetProperty("VendorSpecific2", (value))
}

// GetVendorSpecific2 gets the value of VendorSpecific2 for the instance
func (instance *MSStorageDriver_ATAPISmartData) GetPropertyVendorSpecific2()(value uint8, err error) { 
    retValue, err := instance.GetProperty("VendorSpecific2")
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

// SetVendorSpecific3 sets the value of VendorSpecific3 for the instance
func (instance *MSStorageDriver_ATAPISmartData) SetPropertyVendorSpecific3(value uint8) (err error) { 
    return instance.SetProperty("VendorSpecific3", (value))
}

// GetVendorSpecific3 gets the value of VendorSpecific3 for the instance
func (instance *MSStorageDriver_ATAPISmartData) GetPropertyVendorSpecific3()(value uint8, err error) { 
    retValue, err := instance.GetProperty("VendorSpecific3")
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

// SetVendorSpecific4 sets the value of VendorSpecific4 for the instance
func (instance *MSStorageDriver_ATAPISmartData) SetPropertyVendorSpecific4(value []uint8) (err error) { 
    return instance.SetProperty("VendorSpecific4", (value))
}

// GetVendorSpecific4 gets the value of VendorSpecific4 for the instance
func (instance *MSStorageDriver_ATAPISmartData) GetPropertyVendorSpecific4()(value []uint8, err error) { 
    retValue, err := instance.GetProperty("VendorSpecific4")
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

