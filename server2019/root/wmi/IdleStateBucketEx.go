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

// IdleStateBucketEx struct
type IdleStateBucketEx struct { 
	*cim.WmiInstance

	// 
	Count uint32

	// 
	MaxTimeUs uint32

	// 
	MinTimeUs uint32

	// 
	TotalTimeUs uint64
}

	func NewIdleStateBucketExEx1(instance *cim.WmiInstance) (newInstance *IdleStateBucketEx, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &IdleStateBucketEx {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewIdleStateBucketExEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *IdleStateBucketEx, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &IdleStateBucketEx {
	WmiInstance: tmp,
	}
	return
	}
	

// SetCount sets the value of Count for the instance
func (instance *IdleStateBucketEx) SetPropertyCount(value uint32) (err error) { 
    return instance.SetProperty("Count", (value))
}

// GetCount gets the value of Count for the instance
func (instance *IdleStateBucketEx) GetPropertyCount()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Count")
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

// SetMaxTimeUs sets the value of MaxTimeUs for the instance
func (instance *IdleStateBucketEx) SetPropertyMaxTimeUs(value uint32) (err error) { 
    return instance.SetProperty("MaxTimeUs", (value))
}

// GetMaxTimeUs gets the value of MaxTimeUs for the instance
func (instance *IdleStateBucketEx) GetPropertyMaxTimeUs()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MaxTimeUs")
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

// SetMinTimeUs sets the value of MinTimeUs for the instance
func (instance *IdleStateBucketEx) SetPropertyMinTimeUs(value uint32) (err error) { 
    return instance.SetProperty("MinTimeUs", (value))
}

// GetMinTimeUs gets the value of MinTimeUs for the instance
func (instance *IdleStateBucketEx) GetPropertyMinTimeUs()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MinTimeUs")
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

// SetTotalTimeUs sets the value of TotalTimeUs for the instance
func (instance *IdleStateBucketEx) SetPropertyTotalTimeUs(value uint64) (err error) { 
    return instance.SetProperty("TotalTimeUs", (value))
}

// GetTotalTimeUs gets the value of TotalTimeUs for the instance
func (instance *IdleStateBucketEx) GetPropertyTotalTimeUs()(value uint64, err error) { 
    retValue, err := instance.GetProperty("TotalTimeUs")
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

