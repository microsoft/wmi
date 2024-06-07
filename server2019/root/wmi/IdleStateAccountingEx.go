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

// IdleStateAccountingEx struct
type IdleStateAccountingEx struct { 
	*cim.WmiInstance

	// 
	FailedTransitions uint32

	// 
	IdleTimeBuckets []IdleStateBucketEx

	// 
	IdleTransitions uint32

	// 
	InvalidBucketIndex uint32

	// 
	MaxTimeUs uint32

	// 
	MinTimeUs uint32

	// 
	TotalTime uint64
}

	func NewIdleStateAccountingExEx1(instance *cim.WmiInstance) (newInstance *IdleStateAccountingEx, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &IdleStateAccountingEx {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewIdleStateAccountingExEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *IdleStateAccountingEx, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &IdleStateAccountingEx {
	WmiInstance: tmp,
	}
	return
	}
	

// SetFailedTransitions sets the value of FailedTransitions for the instance
func (instance *IdleStateAccountingEx) SetPropertyFailedTransitions(value uint32) (err error) { 
    return instance.SetProperty("FailedTransitions", (value))
}

// GetFailedTransitions gets the value of FailedTransitions for the instance
func (instance *IdleStateAccountingEx) GetPropertyFailedTransitions()(value uint32, err error) { 
    retValue, err := instance.GetProperty("FailedTransitions")
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

// SetIdleTimeBuckets sets the value of IdleTimeBuckets for the instance
func (instance *IdleStateAccountingEx) SetPropertyIdleTimeBuckets(value []IdleStateBucketEx) (err error) { 
    return instance.SetProperty("IdleTimeBuckets", (value))
}

// GetIdleTimeBuckets gets the value of IdleTimeBuckets for the instance
func (instance *IdleStateAccountingEx) GetPropertyIdleTimeBuckets()(value []IdleStateBucketEx, err error) { 
    retValue, err := instance.GetProperty("IdleTimeBuckets")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(IdleStateBucketEx); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " IdleStateBucketEx is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, IdleStateBucketEx(valuetmp))
    }

    return
}

// SetIdleTransitions sets the value of IdleTransitions for the instance
func (instance *IdleStateAccountingEx) SetPropertyIdleTransitions(value uint32) (err error) { 
    return instance.SetProperty("IdleTransitions", (value))
}

// GetIdleTransitions gets the value of IdleTransitions for the instance
func (instance *IdleStateAccountingEx) GetPropertyIdleTransitions()(value uint32, err error) { 
    retValue, err := instance.GetProperty("IdleTransitions")
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

// SetInvalidBucketIndex sets the value of InvalidBucketIndex for the instance
func (instance *IdleStateAccountingEx) SetPropertyInvalidBucketIndex(value uint32) (err error) { 
    return instance.SetProperty("InvalidBucketIndex", (value))
}

// GetInvalidBucketIndex gets the value of InvalidBucketIndex for the instance
func (instance *IdleStateAccountingEx) GetPropertyInvalidBucketIndex()(value uint32, err error) { 
    retValue, err := instance.GetProperty("InvalidBucketIndex")
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
func (instance *IdleStateAccountingEx) SetPropertyMaxTimeUs(value uint32) (err error) { 
    return instance.SetProperty("MaxTimeUs", (value))
}

// GetMaxTimeUs gets the value of MaxTimeUs for the instance
func (instance *IdleStateAccountingEx) GetPropertyMaxTimeUs()(value uint32, err error) { 
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
func (instance *IdleStateAccountingEx) SetPropertyMinTimeUs(value uint32) (err error) { 
    return instance.SetProperty("MinTimeUs", (value))
}

// GetMinTimeUs gets the value of MinTimeUs for the instance
func (instance *IdleStateAccountingEx) GetPropertyMinTimeUs()(value uint32, err error) { 
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

// SetTotalTime sets the value of TotalTime for the instance
func (instance *IdleStateAccountingEx) SetPropertyTotalTime(value uint64) (err error) { 
    return instance.SetProperty("TotalTime", (value))
}

// GetTotalTime gets the value of TotalTime for the instance
func (instance *IdleStateAccountingEx) GetPropertyTotalTime()(value uint64, err error) { 
    retValue, err := instance.GetProperty("TotalTime")
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

