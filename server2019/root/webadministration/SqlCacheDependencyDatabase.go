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

// SqlCacheDependencyDatabase struct
type SqlCacheDependencyDatabase struct { 
	*CollectionElement

	// 
	ConnectionStringName string

	// 
	Name string

	// 
	PollTime int32
}

	func NewSqlCacheDependencyDatabaseEx1(instance *cim.WmiInstance) (newInstance *SqlCacheDependencyDatabase, err error) {tmp, err := NewCollectionElementEx1(instance)
		
	if err != nil { return }
	newInstance = &SqlCacheDependencyDatabase {
	CollectionElement: tmp,
	}
	return
	}
	

	func NewSqlCacheDependencyDatabaseEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SqlCacheDependencyDatabase, err error) {tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SqlCacheDependencyDatabase {
	CollectionElement: tmp,
	}
	return
	}
	

// SetConnectionStringName sets the value of ConnectionStringName for the instance
func (instance *SqlCacheDependencyDatabase) SetPropertyConnectionStringName(value string) (err error) { 
    return instance.SetProperty("ConnectionStringName", (value))
}

// GetConnectionStringName gets the value of ConnectionStringName for the instance
func (instance *SqlCacheDependencyDatabase) GetPropertyConnectionStringName()(value string, err error) { 
    retValue, err := instance.GetProperty("ConnectionStringName")
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

// SetName sets the value of Name for the instance
func (instance *SqlCacheDependencyDatabase) SetPropertyName(value string) (err error) { 
    return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *SqlCacheDependencyDatabase) GetPropertyName()(value string, err error) { 
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

// SetPollTime sets the value of PollTime for the instance
func (instance *SqlCacheDependencyDatabase) SetPropertyPollTime(value int32) (err error) { 
    return instance.SetProperty("PollTime", (value))
}

// GetPollTime gets the value of PollTime for the instance
func (instance *SqlCacheDependencyDatabase) GetPropertyPollTime()(value int32, err error) { 
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

