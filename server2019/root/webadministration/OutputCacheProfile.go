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

// OutputCacheProfile struct
type OutputCacheProfile struct { 
	*CollectionElement

	// 
	Duration int32

	// 
	Enabled bool

	// 
	Location int32

	// 
	Name string

	// 
	NoStore bool

	// 
	SqlDependency string

	// 
	VaryByControl string

	// 
	VaryByCustom string

	// 
	VaryByHeader string

	// 
	VaryByParam string
}

	func NewOutputCacheProfileEx1(instance *cim.WmiInstance) (newInstance *OutputCacheProfile, err error) {tmp, err := NewCollectionElementEx1(instance)
		
	if err != nil { return }
	newInstance = &OutputCacheProfile {
	CollectionElement: tmp,
	}
	return
	}
	

	func NewOutputCacheProfileEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *OutputCacheProfile, err error) {tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &OutputCacheProfile {
	CollectionElement: tmp,
	}
	return
	}
	

// SetDuration sets the value of Duration for the instance
func (instance *OutputCacheProfile) SetPropertyDuration(value int32) (err error) { 
    return instance.SetProperty("Duration", (value))
}

// GetDuration gets the value of Duration for the instance
func (instance *OutputCacheProfile) GetPropertyDuration()(value int32, err error) { 
    retValue, err := instance.GetProperty("Duration")
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
func (instance *OutputCacheProfile) SetPropertyEnabled(value bool) (err error) { 
    return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *OutputCacheProfile) GetPropertyEnabled()(value bool, err error) { 
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

// SetLocation sets the value of Location for the instance
func (instance *OutputCacheProfile) SetPropertyLocation(value int32) (err error) { 
    return instance.SetProperty("Location", (value))
}

// GetLocation gets the value of Location for the instance
func (instance *OutputCacheProfile) GetPropertyLocation()(value int32, err error) { 
    retValue, err := instance.GetProperty("Location")
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

// SetName sets the value of Name for the instance
func (instance *OutputCacheProfile) SetPropertyName(value string) (err error) { 
    return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *OutputCacheProfile) GetPropertyName()(value string, err error) { 
    retValue, err := instance.GetProperty("Name")
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

// SetNoStore sets the value of NoStore for the instance
func (instance *OutputCacheProfile) SetPropertyNoStore(value bool) (err error) { 
    return instance.SetProperty("NoStore", (value))
}

// GetNoStore gets the value of NoStore for the instance
func (instance *OutputCacheProfile) GetPropertyNoStore()(value bool, err error) { 
    retValue, err := instance.GetProperty("NoStore")
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

// SetSqlDependency sets the value of SqlDependency for the instance
func (instance *OutputCacheProfile) SetPropertySqlDependency(value string) (err error) { 
    return instance.SetProperty("SqlDependency", (value))
}

// GetSqlDependency gets the value of SqlDependency for the instance
func (instance *OutputCacheProfile) GetPropertySqlDependency()(value string, err error) { 
    retValue, err := instance.GetProperty("SqlDependency")
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

// SetVaryByControl sets the value of VaryByControl for the instance
func (instance *OutputCacheProfile) SetPropertyVaryByControl(value string) (err error) { 
    return instance.SetProperty("VaryByControl", (value))
}

// GetVaryByControl gets the value of VaryByControl for the instance
func (instance *OutputCacheProfile) GetPropertyVaryByControl()(value string, err error) { 
    retValue, err := instance.GetProperty("VaryByControl")
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

// SetVaryByCustom sets the value of VaryByCustom for the instance
func (instance *OutputCacheProfile) SetPropertyVaryByCustom(value string) (err error) { 
    return instance.SetProperty("VaryByCustom", (value))
}

// GetVaryByCustom gets the value of VaryByCustom for the instance
func (instance *OutputCacheProfile) GetPropertyVaryByCustom()(value string, err error) { 
    retValue, err := instance.GetProperty("VaryByCustom")
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

// SetVaryByHeader sets the value of VaryByHeader for the instance
func (instance *OutputCacheProfile) SetPropertyVaryByHeader(value string) (err error) { 
    return instance.SetProperty("VaryByHeader", (value))
}

// GetVaryByHeader gets the value of VaryByHeader for the instance
func (instance *OutputCacheProfile) GetPropertyVaryByHeader()(value string, err error) { 
    retValue, err := instance.GetProperty("VaryByHeader")
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

// SetVaryByParam sets the value of VaryByParam for the instance
func (instance *OutputCacheProfile) SetPropertyVaryByParam(value string) (err error) { 
    return instance.SetProperty("VaryByParam", (value))
}

// GetVaryByParam gets the value of VaryByParam for the instance
func (instance *OutputCacheProfile) GetPropertyVaryByParam()(value string, err error) { 
    retValue, err := instance.GetProperty("VaryByParam")
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

