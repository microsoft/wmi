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

// MSDiskDriver_Performance struct
type MSDiskDriver_Performance struct { 
	*MSDiskDriver

	// 
	Active bool

	// 
	DeviceName string

	// 
	InstanceName string

	// 
	PerfData MSDiskDriver_PerformanceData
}

	func NewMSDiskDriver_PerformanceEx1(instance *cim.WmiInstance) (newInstance *MSDiskDriver_Performance, err error) {tmp, err := NewMSDiskDriverEx1(instance)
		
	if err != nil { return }
	newInstance = &MSDiskDriver_Performance {
	MSDiskDriver: tmp,
	}
	return
	}
	

	func NewMSDiskDriver_PerformanceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSDiskDriver_Performance, err error) {tmp, err := NewMSDiskDriverEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSDiskDriver_Performance {
	MSDiskDriver: tmp,
	}
	return
	}
	

// SetActive sets the value of Active for the instance
func (instance *MSDiskDriver_Performance) SetPropertyActive(value bool) (err error) { 
    return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSDiskDriver_Performance) GetPropertyActive()(value bool, err error) { 
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

// SetDeviceName sets the value of DeviceName for the instance
func (instance *MSDiskDriver_Performance) SetPropertyDeviceName(value string) (err error) { 
    return instance.SetProperty("DeviceName", (value))
}

// GetDeviceName gets the value of DeviceName for the instance
func (instance *MSDiskDriver_Performance) GetPropertyDeviceName()(value string, err error) { 
    retValue, err := instance.GetProperty("DeviceName")
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
func (instance *MSDiskDriver_Performance) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSDiskDriver_Performance) GetPropertyInstanceName()(value string, err error) { 
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

// SetPerfData sets the value of PerfData for the instance
func (instance *MSDiskDriver_Performance) SetPropertyPerfData(value MSDiskDriver_PerformanceData) (err error) { 
    return instance.SetProperty("PerfData", (value))
}

// GetPerfData gets the value of PerfData for the instance
func (instance *MSDiskDriver_Performance) GetPropertyPerfData()(value MSDiskDriver_PerformanceData, err error) { 
    retValue, err := instance.GetProperty("PerfData")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(MSDiskDriver_PerformanceData); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " MSDiskDriver_PerformanceData is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = MSDiskDriver_PerformanceData(valuetmp)

    return
}

