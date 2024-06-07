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

// ClusPortDeviceInformation struct
type ClusPortDeviceInformation struct { 
	*cim.WmiInstance

	// 
	Active bool

	// Connected Node .
	ConnectedNode string

	// Connected Node Device Number .
	ConnectedNodeDeviceNumber uint32

	// Connected Node Id.
	ConnectedNodeId uint32

	// Device Path attribute.
	DeviceAttribute uint32

	// Device Guid .
	DeviceGuid string

	// Device Number.
	DeviceNumber uint32

	// ClusPort Device State.
	DeviceState uint32

	// ClusPort Device Type.
	DeviceType uint32

	// Flags
	Flags uint32

	// 
	InstanceName string

	// IsReadOnly.
	IsReadOnly bool

	// Number of Multi-Paths.
	NumberOfPaths uint32

	// Path Info.
	Paths []ClusPortPathInformation

	// ProductId
	ProductId string

	// SerialNumber
	SerialNumber string

	// VendorId
	VendorId string
}

	func NewClusPortDeviceInformationEx1(instance *cim.WmiInstance) (newInstance *ClusPortDeviceInformation, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &ClusPortDeviceInformation {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewClusPortDeviceInformationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ClusPortDeviceInformation, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ClusPortDeviceInformation {
	WmiInstance: tmp,
	}
	return
	}
	

// SetActive sets the value of Active for the instance
func (instance *ClusPortDeviceInformation) SetPropertyActive(value bool) (err error) { 
    return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *ClusPortDeviceInformation) GetPropertyActive()(value bool, err error) { 
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

// SetConnectedNode sets the value of ConnectedNode for the instance
func (instance *ClusPortDeviceInformation) SetPropertyConnectedNode(value string) (err error) { 
    return instance.SetProperty("ConnectedNode", (value))
}

// GetConnectedNode gets the value of ConnectedNode for the instance
func (instance *ClusPortDeviceInformation) GetPropertyConnectedNode()(value string, err error) { 
    retValue, err := instance.GetProperty("ConnectedNode")
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

// SetConnectedNodeDeviceNumber sets the value of ConnectedNodeDeviceNumber for the instance
func (instance *ClusPortDeviceInformation) SetPropertyConnectedNodeDeviceNumber(value uint32) (err error) { 
    return instance.SetProperty("ConnectedNodeDeviceNumber", (value))
}

// GetConnectedNodeDeviceNumber gets the value of ConnectedNodeDeviceNumber for the instance
func (instance *ClusPortDeviceInformation) GetPropertyConnectedNodeDeviceNumber()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ConnectedNodeDeviceNumber")
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

// SetConnectedNodeId sets the value of ConnectedNodeId for the instance
func (instance *ClusPortDeviceInformation) SetPropertyConnectedNodeId(value uint32) (err error) { 
    return instance.SetProperty("ConnectedNodeId", (value))
}

// GetConnectedNodeId gets the value of ConnectedNodeId for the instance
func (instance *ClusPortDeviceInformation) GetPropertyConnectedNodeId()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ConnectedNodeId")
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

// SetDeviceAttribute sets the value of DeviceAttribute for the instance
func (instance *ClusPortDeviceInformation) SetPropertyDeviceAttribute(value uint32) (err error) { 
    return instance.SetProperty("DeviceAttribute", (value))
}

// GetDeviceAttribute gets the value of DeviceAttribute for the instance
func (instance *ClusPortDeviceInformation) GetPropertyDeviceAttribute()(value uint32, err error) { 
    retValue, err := instance.GetProperty("DeviceAttribute")
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

// SetDeviceGuid sets the value of DeviceGuid for the instance
func (instance *ClusPortDeviceInformation) SetPropertyDeviceGuid(value string) (err error) { 
    return instance.SetProperty("DeviceGuid", (value))
}

// GetDeviceGuid gets the value of DeviceGuid for the instance
func (instance *ClusPortDeviceInformation) GetPropertyDeviceGuid()(value string, err error) { 
    retValue, err := instance.GetProperty("DeviceGuid")
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

// SetDeviceNumber sets the value of DeviceNumber for the instance
func (instance *ClusPortDeviceInformation) SetPropertyDeviceNumber(value uint32) (err error) { 
    return instance.SetProperty("DeviceNumber", (value))
}

// GetDeviceNumber gets the value of DeviceNumber for the instance
func (instance *ClusPortDeviceInformation) GetPropertyDeviceNumber()(value uint32, err error) { 
    retValue, err := instance.GetProperty("DeviceNumber")
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

// SetDeviceState sets the value of DeviceState for the instance
func (instance *ClusPortDeviceInformation) SetPropertyDeviceState(value uint32) (err error) { 
    return instance.SetProperty("DeviceState", (value))
}

// GetDeviceState gets the value of DeviceState for the instance
func (instance *ClusPortDeviceInformation) GetPropertyDeviceState()(value uint32, err error) { 
    retValue, err := instance.GetProperty("DeviceState")
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

// SetDeviceType sets the value of DeviceType for the instance
func (instance *ClusPortDeviceInformation) SetPropertyDeviceType(value uint32) (err error) { 
    return instance.SetProperty("DeviceType", (value))
}

// GetDeviceType gets the value of DeviceType for the instance
func (instance *ClusPortDeviceInformation) GetPropertyDeviceType()(value uint32, err error) { 
    retValue, err := instance.GetProperty("DeviceType")
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

// SetFlags sets the value of Flags for the instance
func (instance *ClusPortDeviceInformation) SetPropertyFlags(value uint32) (err error) { 
    return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *ClusPortDeviceInformation) GetPropertyFlags()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Flags")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *ClusPortDeviceInformation) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *ClusPortDeviceInformation) GetPropertyInstanceName()(value string, err error) { 
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

// SetIsReadOnly sets the value of IsReadOnly for the instance
func (instance *ClusPortDeviceInformation) SetPropertyIsReadOnly(value bool) (err error) { 
    return instance.SetProperty("IsReadOnly", (value))
}

// GetIsReadOnly gets the value of IsReadOnly for the instance
func (instance *ClusPortDeviceInformation) GetPropertyIsReadOnly()(value bool, err error) { 
    retValue, err := instance.GetProperty("IsReadOnly")
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

// SetNumberOfPaths sets the value of NumberOfPaths for the instance
func (instance *ClusPortDeviceInformation) SetPropertyNumberOfPaths(value uint32) (err error) { 
    return instance.SetProperty("NumberOfPaths", (value))
}

// GetNumberOfPaths gets the value of NumberOfPaths for the instance
func (instance *ClusPortDeviceInformation) GetPropertyNumberOfPaths()(value uint32, err error) { 
    retValue, err := instance.GetProperty("NumberOfPaths")
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

// SetPaths sets the value of Paths for the instance
func (instance *ClusPortDeviceInformation) SetPropertyPaths(value []ClusPortPathInformation) (err error) { 
    return instance.SetProperty("Paths", (value))
}

// GetPaths gets the value of Paths for the instance
func (instance *ClusPortDeviceInformation) GetPropertyPaths()(value []ClusPortPathInformation, err error) { 
    retValue, err := instance.GetProperty("Paths")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(ClusPortPathInformation); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " ClusPortPathInformation is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, ClusPortPathInformation(valuetmp))
    }

    return
}

// SetProductId sets the value of ProductId for the instance
func (instance *ClusPortDeviceInformation) SetPropertyProductId(value string) (err error) { 
    return instance.SetProperty("ProductId", (value))
}

// GetProductId gets the value of ProductId for the instance
func (instance *ClusPortDeviceInformation) GetPropertyProductId()(value string, err error) { 
    retValue, err := instance.GetProperty("ProductId")
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
func (instance *ClusPortDeviceInformation) SetPropertySerialNumber(value string) (err error) { 
    return instance.SetProperty("SerialNumber", (value))
}

// GetSerialNumber gets the value of SerialNumber for the instance
func (instance *ClusPortDeviceInformation) GetPropertySerialNumber()(value string, err error) { 
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

// SetVendorId sets the value of VendorId for the instance
func (instance *ClusPortDeviceInformation) SetPropertyVendorId(value string) (err error) { 
    return instance.SetProperty("VendorId", (value))
}

// GetVendorId gets the value of VendorId for the instance
func (instance *ClusPortDeviceInformation) GetPropertyVendorId()(value string, err error) { 
    retValue, err := instance.GetProperty("VendorId")
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

