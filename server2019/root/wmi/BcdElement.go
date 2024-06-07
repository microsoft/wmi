// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// BcdElement struct
type BcdElement struct { 
	*cim.WmiInstance

	// This is the id of the object that this element belongs to.
	ObjectId string

	// This is the file path of the store that this element is a part of.
	StoreFilePath string

	// The upper 4 bits (28-31) determine the class of the element. The next 4 bits (24-27) determine the format of the element data. The lower 24 bits (0-23) determine the sub-type of the element.
	Type uint32
}

	func NewBcdElementEx1(instance *cim.WmiInstance) (newInstance *BcdElement, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &BcdElement {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewBcdElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *BcdElement, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &BcdElement {
	WmiInstance: tmp,
	}
	return
	}
	

// SetObjectId sets the value of ObjectId for the instance
func (instance *BcdElement) SetPropertyObjectId(value string) (err error) { 
    return instance.SetProperty("ObjectId", (value))
}

// GetObjectId gets the value of ObjectId for the instance
func (instance *BcdElement) GetPropertyObjectId()(value string, err error) { 
    retValue, err := instance.GetProperty("ObjectId")
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

// SetStoreFilePath sets the value of StoreFilePath for the instance
func (instance *BcdElement) SetPropertyStoreFilePath(value string) (err error) { 
    return instance.SetProperty("StoreFilePath", (value))
}

// GetStoreFilePath gets the value of StoreFilePath for the instance
func (instance *BcdElement) GetPropertyStoreFilePath()(value string, err error) { 
    retValue, err := instance.GetProperty("StoreFilePath")
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

// SetType sets the value of Type for the instance
func (instance *BcdElement) SetPropertyType(value uint32) (err error) { 
    return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *BcdElement) GetPropertyType()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Type")
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

