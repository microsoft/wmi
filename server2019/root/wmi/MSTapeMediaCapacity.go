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

// MSTapeMediaCapacity struct
type MSTapeMediaCapacity struct { 
	*MSTapeDriver

	// 
	Active bool

	// 
	AvailableCapacity uint64

	// 
	BlockSize uint32

	// 
	InstanceName string

	// 
	MaximumCapacity uint64

	// 
	MediaWriteProtected bool

	// 
	PartitionCount uint32
}

	func NewMSTapeMediaCapacityEx1(instance *cim.WmiInstance) (newInstance *MSTapeMediaCapacity, err error) {tmp, err := NewMSTapeDriverEx1(instance)
		
	if err != nil { return }
	newInstance = &MSTapeMediaCapacity {
	MSTapeDriver: tmp,
	}
	return
	}
	

	func NewMSTapeMediaCapacityEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSTapeMediaCapacity, err error) {tmp, err := NewMSTapeDriverEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSTapeMediaCapacity {
	MSTapeDriver: tmp,
	}
	return
	}
	

// SetActive sets the value of Active for the instance
func (instance *MSTapeMediaCapacity) SetPropertyActive(value bool) (err error) { 
    return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSTapeMediaCapacity) GetPropertyActive()(value bool, err error) { 
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

// SetAvailableCapacity sets the value of AvailableCapacity for the instance
func (instance *MSTapeMediaCapacity) SetPropertyAvailableCapacity(value uint64) (err error) { 
    return instance.SetProperty("AvailableCapacity", (value))
}

// GetAvailableCapacity gets the value of AvailableCapacity for the instance
func (instance *MSTapeMediaCapacity) GetPropertyAvailableCapacity()(value uint64, err error) { 
    retValue, err := instance.GetProperty("AvailableCapacity")
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

// SetBlockSize sets the value of BlockSize for the instance
func (instance *MSTapeMediaCapacity) SetPropertyBlockSize(value uint32) (err error) { 
    return instance.SetProperty("BlockSize", (value))
}

// GetBlockSize gets the value of BlockSize for the instance
func (instance *MSTapeMediaCapacity) GetPropertyBlockSize()(value uint32, err error) { 
    retValue, err := instance.GetProperty("BlockSize")
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
func (instance *MSTapeMediaCapacity) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSTapeMediaCapacity) GetPropertyInstanceName()(value string, err error) { 
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

// SetMaximumCapacity sets the value of MaximumCapacity for the instance
func (instance *MSTapeMediaCapacity) SetPropertyMaximumCapacity(value uint64) (err error) { 
    return instance.SetProperty("MaximumCapacity", (value))
}

// GetMaximumCapacity gets the value of MaximumCapacity for the instance
func (instance *MSTapeMediaCapacity) GetPropertyMaximumCapacity()(value uint64, err error) { 
    retValue, err := instance.GetProperty("MaximumCapacity")
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

// SetMediaWriteProtected sets the value of MediaWriteProtected for the instance
func (instance *MSTapeMediaCapacity) SetPropertyMediaWriteProtected(value bool) (err error) { 
    return instance.SetProperty("MediaWriteProtected", (value))
}

// GetMediaWriteProtected gets the value of MediaWriteProtected for the instance
func (instance *MSTapeMediaCapacity) GetPropertyMediaWriteProtected()(value bool, err error) { 
    retValue, err := instance.GetProperty("MediaWriteProtected")
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

// SetPartitionCount sets the value of PartitionCount for the instance
func (instance *MSTapeMediaCapacity) SetPropertyPartitionCount(value uint32) (err error) { 
    return instance.SetProperty("PartitionCount", (value))
}

// GetPartitionCount gets the value of PartitionCount for the instance
func (instance *MSTapeMediaCapacity) GetPropertyPartitionCount()(value uint32, err error) { 
    retValue, err := instance.GetProperty("PartitionCount")
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

