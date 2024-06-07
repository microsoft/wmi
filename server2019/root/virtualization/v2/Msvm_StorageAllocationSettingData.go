// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// Msvm_StorageAllocationSettingData struct
type Msvm_StorageAllocationSettingData struct { 
	*CIM_StorageAllocationSettingData

	// 
	CachingMode uint16

	// 
	IgnoreFlushes bool

	// 
	IOPSAllocationUnits string

	// 
	IOPSLimit uint64

	// 
	IOPSReservation uint64

	// 
	PersistentReservationsSupported bool

	// 
	SnapshotId string

	// 
	StorageQoSPolicyID string

	// 
	WriteHardeningMethod uint16
}

	func NewMsvm_StorageAllocationSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_StorageAllocationSettingData, err error) {tmp, err := NewCIM_StorageAllocationSettingDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Msvm_StorageAllocationSettingData {
	CIM_StorageAllocationSettingData: tmp,
	}
	return
	}
	

	func NewMsvm_StorageAllocationSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Msvm_StorageAllocationSettingData, err error) {tmp, err := NewCIM_StorageAllocationSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Msvm_StorageAllocationSettingData {
	CIM_StorageAllocationSettingData: tmp,
	}
	return
	}
	

// SetCachingMode sets the value of CachingMode for the instance
func (instance *Msvm_StorageAllocationSettingData) SetPropertyCachingMode(value uint16) (err error) { 
    return instance.SetProperty("CachingMode", (value))
}

// GetCachingMode gets the value of CachingMode for the instance
func (instance *Msvm_StorageAllocationSettingData) GetPropertyCachingMode()(value uint16, err error) { 
    retValue, err := instance.GetProperty("CachingMode")
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

// SetIgnoreFlushes sets the value of IgnoreFlushes for the instance
func (instance *Msvm_StorageAllocationSettingData) SetPropertyIgnoreFlushes(value bool) (err error) { 
    return instance.SetProperty("IgnoreFlushes", (value))
}

// GetIgnoreFlushes gets the value of IgnoreFlushes for the instance
func (instance *Msvm_StorageAllocationSettingData) GetPropertyIgnoreFlushes()(value bool, err error) { 
    retValue, err := instance.GetProperty("IgnoreFlushes")
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

// SetIOPSAllocationUnits sets the value of IOPSAllocationUnits for the instance
func (instance *Msvm_StorageAllocationSettingData) SetPropertyIOPSAllocationUnits(value string) (err error) { 
    return instance.SetProperty("IOPSAllocationUnits", (value))
}

// GetIOPSAllocationUnits gets the value of IOPSAllocationUnits for the instance
func (instance *Msvm_StorageAllocationSettingData) GetPropertyIOPSAllocationUnits()(value string, err error) { 
    retValue, err := instance.GetProperty("IOPSAllocationUnits")
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

// SetIOPSLimit sets the value of IOPSLimit for the instance
func (instance *Msvm_StorageAllocationSettingData) SetPropertyIOPSLimit(value uint64) (err error) { 
    return instance.SetProperty("IOPSLimit", (value))
}

// GetIOPSLimit gets the value of IOPSLimit for the instance
func (instance *Msvm_StorageAllocationSettingData) GetPropertyIOPSLimit()(value uint64, err error) { 
    retValue, err := instance.GetProperty("IOPSLimit")
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

// SetIOPSReservation sets the value of IOPSReservation for the instance
func (instance *Msvm_StorageAllocationSettingData) SetPropertyIOPSReservation(value uint64) (err error) { 
    return instance.SetProperty("IOPSReservation", (value))
}

// GetIOPSReservation gets the value of IOPSReservation for the instance
func (instance *Msvm_StorageAllocationSettingData) GetPropertyIOPSReservation()(value uint64, err error) { 
    retValue, err := instance.GetProperty("IOPSReservation")
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

// SetPersistentReservationsSupported sets the value of PersistentReservationsSupported for the instance
func (instance *Msvm_StorageAllocationSettingData) SetPropertyPersistentReservationsSupported(value bool) (err error) { 
    return instance.SetProperty("PersistentReservationsSupported", (value))
}

// GetPersistentReservationsSupported gets the value of PersistentReservationsSupported for the instance
func (instance *Msvm_StorageAllocationSettingData) GetPropertyPersistentReservationsSupported()(value bool, err error) { 
    retValue, err := instance.GetProperty("PersistentReservationsSupported")
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

// SetSnapshotId sets the value of SnapshotId for the instance
func (instance *Msvm_StorageAllocationSettingData) SetPropertySnapshotId(value string) (err error) { 
    return instance.SetProperty("SnapshotId", (value))
}

// GetSnapshotId gets the value of SnapshotId for the instance
func (instance *Msvm_StorageAllocationSettingData) GetPropertySnapshotId()(value string, err error) { 
    retValue, err := instance.GetProperty("SnapshotId")
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

// SetStorageQoSPolicyID sets the value of StorageQoSPolicyID for the instance
func (instance *Msvm_StorageAllocationSettingData) SetPropertyStorageQoSPolicyID(value string) (err error) { 
    return instance.SetProperty("StorageQoSPolicyID", (value))
}

// GetStorageQoSPolicyID gets the value of StorageQoSPolicyID for the instance
func (instance *Msvm_StorageAllocationSettingData) GetPropertyStorageQoSPolicyID()(value string, err error) { 
    retValue, err := instance.GetProperty("StorageQoSPolicyID")
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

// SetWriteHardeningMethod sets the value of WriteHardeningMethod for the instance
func (instance *Msvm_StorageAllocationSettingData) SetPropertyWriteHardeningMethod(value uint16) (err error) { 
    return instance.SetProperty("WriteHardeningMethod", (value))
}

// GetWriteHardeningMethod gets the value of WriteHardeningMethod for the instance
func (instance *Msvm_StorageAllocationSettingData) GetPropertyWriteHardeningMethod()(value uint16, err error) { 
    retValue, err := instance.GetProperty("WriteHardeningMethod")
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
func  (instance* Msvm_StorageAllocationSettingData) GetRelatedAllocationCapabilities() (value *cim.WmiInstance, err error) {
		 return instance.GetRelated("Msvm_AllocationCapabilities"); 
	}
	

