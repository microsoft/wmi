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

// ConfigurationHistorySection struct
type ConfigurationHistorySection struct { 
	*ConfigurationSection

	// 
	Enabled bool

	// 
	HistoryPath string

	// 
	MaxHistories uint32

	// 
	Period string
}

	func NewConfigurationHistorySectionEx1(instance *cim.WmiInstance) (newInstance *ConfigurationHistorySection, err error) {tmp, err := NewConfigurationSectionEx1(instance)
		
	if err != nil { return }
	newInstance = &ConfigurationHistorySection {
	ConfigurationSection: tmp,
	}
	return
	}
	

	func NewConfigurationHistorySectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ConfigurationHistorySection, err error) {tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ConfigurationHistorySection {
	ConfigurationSection: tmp,
	}
	return
	}
	

// SetEnabled sets the value of Enabled for the instance
func (instance *ConfigurationHistorySection) SetPropertyEnabled(value bool) (err error) { 
    return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *ConfigurationHistorySection) GetPropertyEnabled()(value bool, err error) { 
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

// SetHistoryPath sets the value of HistoryPath for the instance
func (instance *ConfigurationHistorySection) SetPropertyHistoryPath(value string) (err error) { 
    return instance.SetProperty("HistoryPath", (value))
}

// GetHistoryPath gets the value of HistoryPath for the instance
func (instance *ConfigurationHistorySection) GetPropertyHistoryPath()(value string, err error) { 
    retValue, err := instance.GetProperty("HistoryPath")
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

// SetMaxHistories sets the value of MaxHistories for the instance
func (instance *ConfigurationHistorySection) SetPropertyMaxHistories(value uint32) (err error) { 
    return instance.SetProperty("MaxHistories", (value))
}

// GetMaxHistories gets the value of MaxHistories for the instance
func (instance *ConfigurationHistorySection) GetPropertyMaxHistories()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MaxHistories")
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

// SetPeriod sets the value of Period for the instance
func (instance *ConfigurationHistorySection) SetPropertyPeriod(value string) (err error) { 
    return instance.SetProperty("Period", (value))
}

// GetPeriod gets the value of Period for the instance
func (instance *ConfigurationHistorySection) GetPropertyPeriod()(value string, err error) { 
    retValue, err := instance.GetProperty("Period")
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

