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

// ServerRuntimeSection struct
type ServerRuntimeSection struct { 
	*ConfigurationSection

	// 
	AlternateHostName string

	// 
	AppConcurrentRequestLimit uint32

	// 
	AuthenticatedUserOverride int32

	// 
	Enabled bool

	// 
	EnableNagling bool

	// 
	FrequentHitThreshold uint32

	// 
	FrequentHitTimePeriod string

	// 
	MaxRequestEntityAllowed uint32

	// 
	UploadReadAheadSize uint32
}

	func NewServerRuntimeSectionEx1(instance *cim.WmiInstance) (newInstance *ServerRuntimeSection, err error) {tmp, err := NewConfigurationSectionEx1(instance)
		
	if err != nil { return }
	newInstance = &ServerRuntimeSection {
	ConfigurationSection: tmp,
	}
	return
	}
	

	func NewServerRuntimeSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ServerRuntimeSection, err error) {tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ServerRuntimeSection {
	ConfigurationSection: tmp,
	}
	return
	}
	

// SetAlternateHostName sets the value of AlternateHostName for the instance
func (instance *ServerRuntimeSection) SetPropertyAlternateHostName(value string) (err error) { 
    return instance.SetProperty("AlternateHostName", (value))
}

// GetAlternateHostName gets the value of AlternateHostName for the instance
func (instance *ServerRuntimeSection) GetPropertyAlternateHostName()(value string, err error) { 
    retValue, err := instance.GetProperty("AlternateHostName")
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

// SetAppConcurrentRequestLimit sets the value of AppConcurrentRequestLimit for the instance
func (instance *ServerRuntimeSection) SetPropertyAppConcurrentRequestLimit(value uint32) (err error) { 
    return instance.SetProperty("AppConcurrentRequestLimit", (value))
}

// GetAppConcurrentRequestLimit gets the value of AppConcurrentRequestLimit for the instance
func (instance *ServerRuntimeSection) GetPropertyAppConcurrentRequestLimit()(value uint32, err error) { 
    retValue, err := instance.GetProperty("AppConcurrentRequestLimit")
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

// SetAuthenticatedUserOverride sets the value of AuthenticatedUserOverride for the instance
func (instance *ServerRuntimeSection) SetPropertyAuthenticatedUserOverride(value int32) (err error) { 
    return instance.SetProperty("AuthenticatedUserOverride", (value))
}

// GetAuthenticatedUserOverride gets the value of AuthenticatedUserOverride for the instance
func (instance *ServerRuntimeSection) GetPropertyAuthenticatedUserOverride()(value int32, err error) { 
    retValue, err := instance.GetProperty("AuthenticatedUserOverride")
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

// SetEnabled sets the value of Enabled for the instance
func (instance *ServerRuntimeSection) SetPropertyEnabled(value bool) (err error) { 
    return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *ServerRuntimeSection) GetPropertyEnabled()(value bool, err error) { 
    retValue, err := instance.GetProperty("Enabled")
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

// SetEnableNagling sets the value of EnableNagling for the instance
func (instance *ServerRuntimeSection) SetPropertyEnableNagling(value bool) (err error) { 
    return instance.SetProperty("EnableNagling", (value))
}

// GetEnableNagling gets the value of EnableNagling for the instance
func (instance *ServerRuntimeSection) GetPropertyEnableNagling()(value bool, err error) { 
    retValue, err := instance.GetProperty("EnableNagling")
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

// SetFrequentHitThreshold sets the value of FrequentHitThreshold for the instance
func (instance *ServerRuntimeSection) SetPropertyFrequentHitThreshold(value uint32) (err error) { 
    return instance.SetProperty("FrequentHitThreshold", (value))
}

// GetFrequentHitThreshold gets the value of FrequentHitThreshold for the instance
func (instance *ServerRuntimeSection) GetPropertyFrequentHitThreshold()(value uint32, err error) { 
    retValue, err := instance.GetProperty("FrequentHitThreshold")
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

// SetFrequentHitTimePeriod sets the value of FrequentHitTimePeriod for the instance
func (instance *ServerRuntimeSection) SetPropertyFrequentHitTimePeriod(value string) (err error) { 
    return instance.SetProperty("FrequentHitTimePeriod", (value))
}

// GetFrequentHitTimePeriod gets the value of FrequentHitTimePeriod for the instance
func (instance *ServerRuntimeSection) GetPropertyFrequentHitTimePeriod()(value string, err error) { 
    retValue, err := instance.GetProperty("FrequentHitTimePeriod")
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

// SetMaxRequestEntityAllowed sets the value of MaxRequestEntityAllowed for the instance
func (instance *ServerRuntimeSection) SetPropertyMaxRequestEntityAllowed(value uint32) (err error) { 
    return instance.SetProperty("MaxRequestEntityAllowed", (value))
}

// GetMaxRequestEntityAllowed gets the value of MaxRequestEntityAllowed for the instance
func (instance *ServerRuntimeSection) GetPropertyMaxRequestEntityAllowed()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MaxRequestEntityAllowed")
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

// SetUploadReadAheadSize sets the value of UploadReadAheadSize for the instance
func (instance *ServerRuntimeSection) SetPropertyUploadReadAheadSize(value uint32) (err error) { 
    return instance.SetProperty("UploadReadAheadSize", (value))
}

// GetUploadReadAheadSize gets the value of UploadReadAheadSize for the instance
func (instance *ServerRuntimeSection) GetPropertyUploadReadAheadSize()(value uint32, err error) { 
    retValue, err := instance.GetProperty("UploadReadAheadSize")
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

