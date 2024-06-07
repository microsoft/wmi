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

// OutputCacheSection struct
type OutputCacheSection struct { 
	*ConfigurationSection

	// 
	EnableFragmentCache bool

	// 
	EnableKernelCacheForVaryByStar bool

	// 
	EnableOutputCache bool

	// 
	OmitVaryStar bool

	// 
	SendCacheControlHeader bool
}

	func NewOutputCacheSectionEx1(instance *cim.WmiInstance) (newInstance *OutputCacheSection, err error) {tmp, err := NewConfigurationSectionEx1(instance)
		
	if err != nil { return }
	newInstance = &OutputCacheSection {
	ConfigurationSection: tmp,
	}
	return
	}
	

	func NewOutputCacheSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *OutputCacheSection, err error) {tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &OutputCacheSection {
	ConfigurationSection: tmp,
	}
	return
	}
	

// SetEnableFragmentCache sets the value of EnableFragmentCache for the instance
func (instance *OutputCacheSection) SetPropertyEnableFragmentCache(value bool) (err error) { 
    return instance.SetProperty("EnableFragmentCache", (value))
}

// GetEnableFragmentCache gets the value of EnableFragmentCache for the instance
func (instance *OutputCacheSection) GetPropertyEnableFragmentCache()(value bool, err error) { 
    retValue, err := instance.GetProperty("EnableFragmentCache")
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

// SetEnableKernelCacheForVaryByStar sets the value of EnableKernelCacheForVaryByStar for the instance
func (instance *OutputCacheSection) SetPropertyEnableKernelCacheForVaryByStar(value bool) (err error) { 
    return instance.SetProperty("EnableKernelCacheForVaryByStar", (value))
}

// GetEnableKernelCacheForVaryByStar gets the value of EnableKernelCacheForVaryByStar for the instance
func (instance *OutputCacheSection) GetPropertyEnableKernelCacheForVaryByStar()(value bool, err error) { 
    retValue, err := instance.GetProperty("EnableKernelCacheForVaryByStar")
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

// SetEnableOutputCache sets the value of EnableOutputCache for the instance
func (instance *OutputCacheSection) SetPropertyEnableOutputCache(value bool) (err error) { 
    return instance.SetProperty("EnableOutputCache", (value))
}

// GetEnableOutputCache gets the value of EnableOutputCache for the instance
func (instance *OutputCacheSection) GetPropertyEnableOutputCache()(value bool, err error) { 
    retValue, err := instance.GetProperty("EnableOutputCache")
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

// SetOmitVaryStar sets the value of OmitVaryStar for the instance
func (instance *OutputCacheSection) SetPropertyOmitVaryStar(value bool) (err error) { 
    return instance.SetProperty("OmitVaryStar", (value))
}

// GetOmitVaryStar gets the value of OmitVaryStar for the instance
func (instance *OutputCacheSection) GetPropertyOmitVaryStar()(value bool, err error) { 
    retValue, err := instance.GetProperty("OmitVaryStar")
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

// SetSendCacheControlHeader sets the value of SendCacheControlHeader for the instance
func (instance *OutputCacheSection) SetPropertySendCacheControlHeader(value bool) (err error) { 
    return instance.SetProperty("SendCacheControlHeader", (value))
}

// GetSendCacheControlHeader gets the value of SendCacheControlHeader for the instance
func (instance *OutputCacheSection) GetPropertySendCacheControlHeader()(value bool, err error) { 
    retValue, err := instance.GetProperty("SendCacheControlHeader")
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

