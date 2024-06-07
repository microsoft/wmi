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

// ApplicationDependency struct
type ApplicationDependency struct { 
	*CollectionElement

	// 
	Application []GroupDependency

	// 
	GroupId string

	// 
	Name string
}

	func NewApplicationDependencyEx1(instance *cim.WmiInstance) (newInstance *ApplicationDependency, err error) {tmp, err := NewCollectionElementEx1(instance)
		
	if err != nil { return }
	newInstance = &ApplicationDependency {
	CollectionElement: tmp,
	}
	return
	}
	

	func NewApplicationDependencyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ApplicationDependency, err error) {tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ApplicationDependency {
	CollectionElement: tmp,
	}
	return
	}
	

// SetApplication sets the value of Application for the instance
func (instance *ApplicationDependency) SetPropertyApplication(value []GroupDependency) (err error) { 
    return instance.SetProperty("Application", (value))
}

// GetApplication gets the value of Application for the instance
func (instance *ApplicationDependency) GetPropertyApplication()(value []GroupDependency, err error) { 
    retValue, err := instance.GetProperty("Application")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(GroupDependency); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " GroupDependency is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, GroupDependency(valuetmp))
    }

    return
}

// SetGroupId sets the value of GroupId for the instance
func (instance *ApplicationDependency) SetPropertyGroupId(value string) (err error) { 
    return instance.SetProperty("GroupId", (value))
}

// GetGroupId gets the value of GroupId for the instance
func (instance *ApplicationDependency) GetPropertyGroupId()(value string, err error) { 
    retValue, err := instance.GetProperty("GroupId")
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
func (instance *ApplicationDependency) SetPropertyName(value string) (err error) { 
    return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *ApplicationDependency) GetPropertyName()(value string, err error) { 
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

