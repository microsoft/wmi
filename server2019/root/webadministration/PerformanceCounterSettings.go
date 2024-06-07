// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// PerformanceCounterSettings struct
type PerformanceCounterSettings struct { 
	*EmbeddedObject

	// 
	FileMappingSize int32
}

	func NewPerformanceCounterSettingsEx1(instance *cim.WmiInstance) (newInstance *PerformanceCounterSettings, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &PerformanceCounterSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewPerformanceCounterSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *PerformanceCounterSettings, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &PerformanceCounterSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetFileMappingSize sets the value of FileMappingSize for the instance
func (instance *PerformanceCounterSettings) SetPropertyFileMappingSize(value int32) (err error) { 
    return instance.SetProperty("FileMappingSize", (value))
}

// GetFileMappingSize gets the value of FileMappingSize for the instance
func (instance *PerformanceCounterSettings) GetPropertyFileMappingSize()(value int32, err error) { 
    retValue, err := instance.GetProperty("FileMappingSize")
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

    value = int32(valuetmp)

    return
}

