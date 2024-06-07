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

// Registry_HiveDestroy struct
type Registry_HiveDestroy struct { 
	*Registry

	// 
	FileName string

	// 
	Hive uint32

	// 
	Path string
}

	func NewRegistry_HiveDestroyEx1(instance *cim.WmiInstance) (newInstance *Registry_HiveDestroy, err error) {tmp, err := NewRegistryEx1(instance)
		
	if err != nil { return }
	newInstance = &Registry_HiveDestroy {
	Registry: tmp,
	}
	return
	}
	

	func NewRegistry_HiveDestroyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Registry_HiveDestroy, err error) {tmp, err := NewRegistryEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Registry_HiveDestroy {
	Registry: tmp,
	}
	return
	}
	

// SetFileName sets the value of FileName for the instance
func (instance *Registry_HiveDestroy) SetPropertyFileName(value string) (err error) { 
    return instance.SetProperty("FileName", (value))
}

// GetFileName gets the value of FileName for the instance
func (instance *Registry_HiveDestroy) GetPropertyFileName()(value string, err error) { 
    retValue, err := instance.GetProperty("FileName")
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

// SetHive sets the value of Hive for the instance
func (instance *Registry_HiveDestroy) SetPropertyHive(value uint32) (err error) { 
    return instance.SetProperty("Hive", (value))
}

// GetHive gets the value of Hive for the instance
func (instance *Registry_HiveDestroy) GetPropertyHive()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Hive")
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

// SetPath sets the value of Path for the instance
func (instance *Registry_HiveDestroy) SetPropertyPath(value string) (err error) { 
    return instance.SetProperty("Path", (value))
}

// GetPath gets the value of Path for the instance
func (instance *Registry_HiveDestroy) GetPropertyPath()(value string, err error) { 
    retValue, err := instance.GetProperty("Path")
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

