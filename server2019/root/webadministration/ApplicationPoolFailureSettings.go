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

// ApplicationPoolFailureSettings struct
type ApplicationPoolFailureSettings struct { 
	*EmbeddedObject

	// 
	AutoShutdownExe string

	// 
	AutoShutdownParams string

	// 
	LoadBalancerCapabilities int32

	// 
	OrphanActionExe string

	// 
	OrphanActionParams string

	// 
	OrphanWorkerProcess bool

	// 
	RapidFailProtection bool

	// 
	RapidFailProtectionInterval string

	// 
	RapidFailProtectionMaxCrashes uint32
}

	func NewApplicationPoolFailureSettingsEx1(instance *cim.WmiInstance) (newInstance *ApplicationPoolFailureSettings, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &ApplicationPoolFailureSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewApplicationPoolFailureSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ApplicationPoolFailureSettings, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ApplicationPoolFailureSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetAutoShutdownExe sets the value of AutoShutdownExe for the instance
func (instance *ApplicationPoolFailureSettings) SetPropertyAutoShutdownExe(value string) (err error) { 
    return instance.SetProperty("AutoShutdownExe", (value))
}

// GetAutoShutdownExe gets the value of AutoShutdownExe for the instance
func (instance *ApplicationPoolFailureSettings) GetPropertyAutoShutdownExe()(value string, err error) { 
    retValue, err := instance.GetProperty("AutoShutdownExe")
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

// SetAutoShutdownParams sets the value of AutoShutdownParams for the instance
func (instance *ApplicationPoolFailureSettings) SetPropertyAutoShutdownParams(value string) (err error) { 
    return instance.SetProperty("AutoShutdownParams", (value))
}

// GetAutoShutdownParams gets the value of AutoShutdownParams for the instance
func (instance *ApplicationPoolFailureSettings) GetPropertyAutoShutdownParams()(value string, err error) { 
    retValue, err := instance.GetProperty("AutoShutdownParams")
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

// SetLoadBalancerCapabilities sets the value of LoadBalancerCapabilities for the instance
func (instance *ApplicationPoolFailureSettings) SetPropertyLoadBalancerCapabilities(value int32) (err error) { 
    return instance.SetProperty("LoadBalancerCapabilities", (value))
}

// GetLoadBalancerCapabilities gets the value of LoadBalancerCapabilities for the instance
func (instance *ApplicationPoolFailureSettings) GetPropertyLoadBalancerCapabilities()(value int32, err error) { 
    retValue, err := instance.GetProperty("LoadBalancerCapabilities")
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

// SetOrphanActionExe sets the value of OrphanActionExe for the instance
func (instance *ApplicationPoolFailureSettings) SetPropertyOrphanActionExe(value string) (err error) { 
    return instance.SetProperty("OrphanActionExe", (value))
}

// GetOrphanActionExe gets the value of OrphanActionExe for the instance
func (instance *ApplicationPoolFailureSettings) GetPropertyOrphanActionExe()(value string, err error) { 
    retValue, err := instance.GetProperty("OrphanActionExe")
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

// SetOrphanActionParams sets the value of OrphanActionParams for the instance
func (instance *ApplicationPoolFailureSettings) SetPropertyOrphanActionParams(value string) (err error) { 
    return instance.SetProperty("OrphanActionParams", (value))
}

// GetOrphanActionParams gets the value of OrphanActionParams for the instance
func (instance *ApplicationPoolFailureSettings) GetPropertyOrphanActionParams()(value string, err error) { 
    retValue, err := instance.GetProperty("OrphanActionParams")
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

// SetOrphanWorkerProcess sets the value of OrphanWorkerProcess for the instance
func (instance *ApplicationPoolFailureSettings) SetPropertyOrphanWorkerProcess(value bool) (err error) { 
    return instance.SetProperty("OrphanWorkerProcess", (value))
}

// GetOrphanWorkerProcess gets the value of OrphanWorkerProcess for the instance
func (instance *ApplicationPoolFailureSettings) GetPropertyOrphanWorkerProcess()(value bool, err error) { 
    retValue, err := instance.GetProperty("OrphanWorkerProcess")
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

// SetRapidFailProtection sets the value of RapidFailProtection for the instance
func (instance *ApplicationPoolFailureSettings) SetPropertyRapidFailProtection(value bool) (err error) { 
    return instance.SetProperty("RapidFailProtection", (value))
}

// GetRapidFailProtection gets the value of RapidFailProtection for the instance
func (instance *ApplicationPoolFailureSettings) GetPropertyRapidFailProtection()(value bool, err error) { 
    retValue, err := instance.GetProperty("RapidFailProtection")
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

// SetRapidFailProtectionInterval sets the value of RapidFailProtectionInterval for the instance
func (instance *ApplicationPoolFailureSettings) SetPropertyRapidFailProtectionInterval(value string) (err error) { 
    return instance.SetProperty("RapidFailProtectionInterval", (value))
}

// GetRapidFailProtectionInterval gets the value of RapidFailProtectionInterval for the instance
func (instance *ApplicationPoolFailureSettings) GetPropertyRapidFailProtectionInterval()(value string, err error) { 
    retValue, err := instance.GetProperty("RapidFailProtectionInterval")
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

// SetRapidFailProtectionMaxCrashes sets the value of RapidFailProtectionMaxCrashes for the instance
func (instance *ApplicationPoolFailureSettings) SetPropertyRapidFailProtectionMaxCrashes(value uint32) (err error) { 
    return instance.SetProperty("RapidFailProtectionMaxCrashes", (value))
}

// GetRapidFailProtectionMaxCrashes gets the value of RapidFailProtectionMaxCrashes for the instance
func (instance *ApplicationPoolFailureSettings) GetPropertyRapidFailProtectionMaxCrashes()(value uint32, err error) { 
    retValue, err := instance.GetProperty("RapidFailProtectionMaxCrashes")
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

