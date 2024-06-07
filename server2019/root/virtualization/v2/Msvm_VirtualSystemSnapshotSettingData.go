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

// Msvm_VirtualSystemSnapshotSettingData struct
type Msvm_VirtualSystemSnapshotSettingData struct { 
	*CIM_SettingData

	// 
	ConsistencyLevel VirtualSystemSnapshotSettingData_ConsistencyLevel

	// 
	GuestBackupType VirtualSystemSnapshotSettingData_GuestBackupType

	// 
	IgnoreNonSnapshottableDisks bool
}

	func NewMsvm_VirtualSystemSnapshotSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_VirtualSystemSnapshotSettingData, err error) {tmp, err := NewCIM_SettingDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Msvm_VirtualSystemSnapshotSettingData {
	CIM_SettingData: tmp,
	}
	return
	}
	

	func NewMsvm_VirtualSystemSnapshotSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Msvm_VirtualSystemSnapshotSettingData, err error) {tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Msvm_VirtualSystemSnapshotSettingData {
	CIM_SettingData: tmp,
	}
	return
	}
	

// SetConsistencyLevel sets the value of ConsistencyLevel for the instance
func (instance *Msvm_VirtualSystemSnapshotSettingData) SetPropertyConsistencyLevel(value VirtualSystemSnapshotSettingData_ConsistencyLevel) (err error) { 
    return instance.SetProperty("ConsistencyLevel", (value))
}

// GetConsistencyLevel gets the value of ConsistencyLevel for the instance
func (instance *Msvm_VirtualSystemSnapshotSettingData) GetPropertyConsistencyLevel()(value VirtualSystemSnapshotSettingData_ConsistencyLevel, err error) { 
    retValue, err := instance.GetProperty("ConsistencyLevel")
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

    value = VirtualSystemSnapshotSettingData_ConsistencyLevel(valuetmp)

    return
}

// SetGuestBackupType sets the value of GuestBackupType for the instance
func (instance *Msvm_VirtualSystemSnapshotSettingData) SetPropertyGuestBackupType(value VirtualSystemSnapshotSettingData_GuestBackupType) (err error) { 
    return instance.SetProperty("GuestBackupType", (value))
}

// GetGuestBackupType gets the value of GuestBackupType for the instance
func (instance *Msvm_VirtualSystemSnapshotSettingData) GetPropertyGuestBackupType()(value VirtualSystemSnapshotSettingData_GuestBackupType, err error) { 
    retValue, err := instance.GetProperty("GuestBackupType")
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

    value = VirtualSystemSnapshotSettingData_GuestBackupType(valuetmp)

    return
}

// SetIgnoreNonSnapshottableDisks sets the value of IgnoreNonSnapshottableDisks for the instance
func (instance *Msvm_VirtualSystemSnapshotSettingData) SetPropertyIgnoreNonSnapshottableDisks(value bool) (err error) { 
    return instance.SetProperty("IgnoreNonSnapshottableDisks", (value))
}

// GetIgnoreNonSnapshottableDisks gets the value of IgnoreNonSnapshottableDisks for the instance
func (instance *Msvm_VirtualSystemSnapshotSettingData) GetPropertyIgnoreNonSnapshottableDisks()(value bool, err error) { 
    retValue, err := instance.GetProperty("IgnoreNonSnapshottableDisks")
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

