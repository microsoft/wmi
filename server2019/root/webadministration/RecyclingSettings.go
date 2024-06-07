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

// RecyclingSettings struct
type RecyclingSettings struct { 
	*EmbeddedObject

	// 
	DisallowOverlappingRotation bool

	// 
	DisallowRotationOnConfigChange bool

	// 
	LogEventOnRecycle int32

	// 
	PeriodicRestart PeriodicRestartSettings
}

	func NewRecyclingSettingsEx1(instance *cim.WmiInstance) (newInstance *RecyclingSettings, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &RecyclingSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewRecyclingSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *RecyclingSettings, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &RecyclingSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetDisallowOverlappingRotation sets the value of DisallowOverlappingRotation for the instance
func (instance *RecyclingSettings) SetPropertyDisallowOverlappingRotation(value bool) (err error) { 
    return instance.SetProperty("DisallowOverlappingRotation", (value))
}

// GetDisallowOverlappingRotation gets the value of DisallowOverlappingRotation for the instance
func (instance *RecyclingSettings) GetPropertyDisallowOverlappingRotation()(value bool, err error) { 
    retValue, err := instance.GetProperty("DisallowOverlappingRotation")
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

// SetDisallowRotationOnConfigChange sets the value of DisallowRotationOnConfigChange for the instance
func (instance *RecyclingSettings) SetPropertyDisallowRotationOnConfigChange(value bool) (err error) { 
    return instance.SetProperty("DisallowRotationOnConfigChange", (value))
}

// GetDisallowRotationOnConfigChange gets the value of DisallowRotationOnConfigChange for the instance
func (instance *RecyclingSettings) GetPropertyDisallowRotationOnConfigChange()(value bool, err error) { 
    retValue, err := instance.GetProperty("DisallowRotationOnConfigChange")
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

// SetLogEventOnRecycle sets the value of LogEventOnRecycle for the instance
func (instance *RecyclingSettings) SetPropertyLogEventOnRecycle(value int32) (err error) { 
    return instance.SetProperty("LogEventOnRecycle", (value))
}

// GetLogEventOnRecycle gets the value of LogEventOnRecycle for the instance
func (instance *RecyclingSettings) GetPropertyLogEventOnRecycle()(value int32, err error) { 
    retValue, err := instance.GetProperty("LogEventOnRecycle")
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

// SetPeriodicRestart sets the value of PeriodicRestart for the instance
func (instance *RecyclingSettings) SetPropertyPeriodicRestart(value PeriodicRestartSettings) (err error) { 
    return instance.SetProperty("PeriodicRestart", (value))
}

// GetPeriodicRestart gets the value of PeriodicRestart for the instance
func (instance *RecyclingSettings) GetPropertyPeriodicRestart()(value PeriodicRestartSettings, err error) { 
    retValue, err := instance.GetProperty("PeriodicRestart")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(PeriodicRestartSettings); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " PeriodicRestartSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = PeriodicRestartSettings(valuetmp)

    return
}

