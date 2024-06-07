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

// TrustLevel struct
type TrustLevel struct { 
	*CollectionElement

	// 
	Name string

	// 
	PolicyFile string
}

	func NewTrustLevelEx1(instance *cim.WmiInstance) (newInstance *TrustLevel, err error) {tmp, err := NewCollectionElementEx1(instance)
		
	if err != nil { return }
	newInstance = &TrustLevel {
	CollectionElement: tmp,
	}
	return
	}
	

	func NewTrustLevelEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *TrustLevel, err error) {tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &TrustLevel {
	CollectionElement: tmp,
	}
	return
	}
	

// SetName sets the value of Name for the instance
func (instance *TrustLevel) SetPropertyName(value string) (err error) { 
    return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *TrustLevel) GetPropertyName()(value string, err error) { 
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

// SetPolicyFile sets the value of PolicyFile for the instance
func (instance *TrustLevel) SetPropertyPolicyFile(value string) (err error) { 
    return instance.SetProperty("PolicyFile", (value))
}

// GetPolicyFile gets the value of PolicyFile for the instance
func (instance *TrustLevel) GetPropertyPolicyFile()(value string, err error) { 
    retValue, err := instance.GetProperty("PolicyFile")
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

