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

// GroupDependency struct
type GroupDependency struct { 
	*CollectionElement

	// 
	GroupId string
}

	func NewGroupDependencyEx1(instance *cim.WmiInstance) (newInstance *GroupDependency, err error) {tmp, err := NewCollectionElementEx1(instance)
		
	if err != nil { return }
	newInstance = &GroupDependency {
	CollectionElement: tmp,
	}
	return
	}
	

	func NewGroupDependencyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *GroupDependency, err error) {tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &GroupDependency {
	CollectionElement: tmp,
	}
	return
	}
	

// SetGroupId sets the value of GroupId for the instance
func (instance *GroupDependency) SetPropertyGroupId(value string) (err error) { 
    return instance.SetProperty("GroupId", (value))
}

// GetGroupId gets the value of GroupId for the instance
func (instance *GroupDependency) GetPropertyGroupId()(value string, err error) { 
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

