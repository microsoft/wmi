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

// SqlCacheDependencySection struct
type SqlCacheDependencySection struct { 
	*ConfigurationSectionWithCollection

	// 
	Databases SqlCacheDependencyDatabaseSettings

	// 
	Enabled bool

	// 
	PollTime int32
}

	func NewSqlCacheDependencySectionEx1(instance *cim.WmiInstance) (newInstance *SqlCacheDependencySection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx1(instance)
		
	if err != nil { return }
	newInstance = &SqlCacheDependencySection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

	func NewSqlCacheDependencySectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SqlCacheDependencySection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SqlCacheDependencySection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

// SetDatabases sets the value of Databases for the instance
func (instance *SqlCacheDependencySection) SetPropertyDatabases(value SqlCacheDependencyDatabaseSettings) (err error) { 
    return instance.SetProperty("Databases", (value))
}

// GetDatabases gets the value of Databases for the instance
func (instance *SqlCacheDependencySection) GetPropertyDatabases()(value SqlCacheDependencyDatabaseSettings, err error) { 
    retValue, err := instance.GetProperty("Databases")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(SqlCacheDependencyDatabaseSettings); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " SqlCacheDependencyDatabaseSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = SqlCacheDependencyDatabaseSettings(valuetmp)

    return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *SqlCacheDependencySection) SetPropertyEnabled(value bool) (err error) { 
    return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *SqlCacheDependencySection) GetPropertyEnabled()(value bool, err error) { 
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

// SetPollTime sets the value of PollTime for the instance
func (instance *SqlCacheDependencySection) SetPropertyPollTime(value int32) (err error) { 
    return instance.SetProperty("PollTime", (value))
}

// GetPollTime gets the value of PollTime for the instance
func (instance *SqlCacheDependencySection) GetPropertyPollTime()(value int32, err error) { 
    retValue, err := instance.GetProperty("PollTime")
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

