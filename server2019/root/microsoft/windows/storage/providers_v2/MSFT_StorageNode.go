// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSFT_StorageNode struct
type MSFT_StorageNode struct { 
	*MSFT_StorageObject

	// This field is a string representation of the node's firmware version.
	FirmwareVersion string

	// 
	Manufacturer string

	// 
	Model string

	// Name is a human-readable string used to identify a storage node.
	Name string

	// NameFormat describes the format of the Name identifier.
	NameFormat StorageNode_NameFormat

	// Indicates the current status of the node.
	OperationalStatus StorageNode_OperationalStatus

	// This field is an array of custom identifier for the node. If this field is set, the OtherIdentifyingInfoDescription field must also be set.
	OtherIdentifyingInfo []string

	// An array of string description of the format used in the custom identifiers defined in the OtherIdentifyingInfo field. There must be a 1:1 mapping between this array and OtherIdentifyingInfo.
	OtherIdentifyingInfoDescription []string

	// 
	SerialNumber string
}

	func NewMSFT_StorageNodeEx1(instance *cim.WmiInstance) (newInstance *MSFT_StorageNode, err error) {tmp, err := NewMSFT_StorageObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &MSFT_StorageNode {
	MSFT_StorageObject: tmp,
	}
	return
	}
	

	func NewMSFT_StorageNodeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_StorageNode, err error) {tmp, err := NewMSFT_StorageObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_StorageNode {
	MSFT_StorageObject: tmp,
	}
	return
	}
	

// SetFirmwareVersion sets the value of FirmwareVersion for the instance
func (instance *MSFT_StorageNode) SetPropertyFirmwareVersion(value string) (err error) { 
    return instance.SetProperty("FirmwareVersion", (value))
}

// GetFirmwareVersion gets the value of FirmwareVersion for the instance
func (instance *MSFT_StorageNode) GetPropertyFirmwareVersion()(value string, err error) { 
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

// SetManufacturer sets the value of Manufacturer for the instance
func (instance *MSFT_StorageNode) SetPropertyManufacturer(value string) (err error) { 
    return instance.SetProperty("Manufacturer", (value))
}

// GetManufacturer gets the value of Manufacturer for the instance
func (instance *MSFT_StorageNode) GetPropertyManufacturer()(value string, err error) { 
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

// SetModel sets the value of Model for the instance
func (instance *MSFT_StorageNode) SetPropertyModel(value string) (err error) { 
    return instance.SetProperty("Model", (value))
}

// GetModel gets the value of Model for the instance
func (instance *MSFT_StorageNode) GetPropertyModel()(value string, err error) { 
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

// SetName sets the value of Name for the instance
func (instance *MSFT_StorageNode) SetPropertyName(value string) (err error) { 
    return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSFT_StorageNode) GetPropertyName()(value string, err error) { 
    retValue, err := instance.GetProperty("Name")
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

// SetNameFormat sets the value of NameFormat for the instance
func (instance *MSFT_StorageNode) SetPropertyNameFormat(value StorageNode_NameFormat) (err error) { 
    return instance.SetProperty("NameFormat", (value))
}

// GetNameFormat gets the value of NameFormat for the instance
func (instance *MSFT_StorageNode) GetPropertyNameFormat()(value StorageNode_NameFormat, err error) { 
    retValue, err := instance.GetProperty("NameFormat")
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

    value = StorageNode_NameFormat(valuetmp)

    return
}

// SetOperationalStatus sets the value of OperationalStatus for the instance
func (instance *MSFT_StorageNode) SetPropertyOperationalStatus(value StorageNode_OperationalStatus) (err error) { 
    return instance.SetProperty("OperationalStatus", (value))
}

// GetOperationalStatus gets the value of OperationalStatus for the instance
func (instance *MSFT_StorageNode) GetPropertyOperationalStatus()(value StorageNode_OperationalStatus, err error) { 
    retValue, err := instance.GetProperty("OperationalStatus")
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

    value = StorageNode_OperationalStatus(valuetmp)

    return
}

// SetOtherIdentifyingInfo sets the value of OtherIdentifyingInfo for the instance
func (instance *MSFT_StorageNode) SetPropertyOtherIdentifyingInfo(value []string) (err error) { 
    return instance.SetProperty("OtherIdentifyingInfo", (value))
}

// GetOtherIdentifyingInfo gets the value of OtherIdentifyingInfo for the instance
func (instance *MSFT_StorageNode) GetPropertyOtherIdentifyingInfo()(value []string, err error) { 
    retValue, err := instance.GetProperty("OtherIdentifyingInfo")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(string); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, string(valuetmp))
    }

    return
}

// SetOtherIdentifyingInfoDescription sets the value of OtherIdentifyingInfoDescription for the instance
func (instance *MSFT_StorageNode) SetPropertyOtherIdentifyingInfoDescription(value []string) (err error) { 
    return instance.SetProperty("OtherIdentifyingInfoDescription", (value))
}

// GetOtherIdentifyingInfoDescription gets the value of OtherIdentifyingInfoDescription for the instance
func (instance *MSFT_StorageNode) GetPropertyOtherIdentifyingInfoDescription()(value []string, err error) { 
    retValue, err := instance.GetProperty("OtherIdentifyingInfoDescription")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(string); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, string(valuetmp))
    }

    return
}

// SetSerialNumber sets the value of SerialNumber for the instance
func (instance *MSFT_StorageNode) SetPropertySerialNumber(value string) (err error) { 
    return instance.SetProperty("SerialNumber", (value))
}

// GetSerialNumber gets the value of SerialNumber for the instance
func (instance *MSFT_StorageNode) GetPropertySerialNumber()(value string, err error) { 
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

