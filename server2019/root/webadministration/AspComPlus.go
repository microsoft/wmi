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

// AspComPlus struct
type AspComPlus struct { 
	*EmbeddedObject

	// 
	AppServiceFlags int32

	// 
	ExecuteInMta bool

	// 
	PartitionId string

	// 
	SxsName string

	// 
	TrackThreadingModel bool
}

	func NewAspComPlusEx1(instance *cim.WmiInstance) (newInstance *AspComPlus, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &AspComPlus {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewAspComPlusEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *AspComPlus, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &AspComPlus {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetAppServiceFlags sets the value of AppServiceFlags for the instance
func (instance *AspComPlus) SetPropertyAppServiceFlags(value int32) (err error) { 
    return instance.SetProperty("AppServiceFlags", (value))
}

// GetAppServiceFlags gets the value of AppServiceFlags for the instance
func (instance *AspComPlus) GetPropertyAppServiceFlags()(value int32, err error) { 
    retValue, err := instance.GetProperty("AppServiceFlags")
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

// SetExecuteInMta sets the value of ExecuteInMta for the instance
func (instance *AspComPlus) SetPropertyExecuteInMta(value bool) (err error) { 
    return instance.SetProperty("ExecuteInMta", (value))
}

// GetExecuteInMta gets the value of ExecuteInMta for the instance
func (instance *AspComPlus) GetPropertyExecuteInMta()(value bool, err error) { 
    retValue, err := instance.GetProperty("ExecuteInMta")
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

// SetPartitionId sets the value of PartitionId for the instance
func (instance *AspComPlus) SetPropertyPartitionId(value string) (err error) { 
    return instance.SetProperty("PartitionId", (value))
}

// GetPartitionId gets the value of PartitionId for the instance
func (instance *AspComPlus) GetPropertyPartitionId()(value string, err error) { 
    retValue, err := instance.GetProperty("PartitionId")
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

// SetSxsName sets the value of SxsName for the instance
func (instance *AspComPlus) SetPropertySxsName(value string) (err error) { 
    return instance.SetProperty("SxsName", (value))
}

// GetSxsName gets the value of SxsName for the instance
func (instance *AspComPlus) GetPropertySxsName()(value string, err error) { 
    retValue, err := instance.GetProperty("SxsName")
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

// SetTrackThreadingModel sets the value of TrackThreadingModel for the instance
func (instance *AspComPlus) SetPropertyTrackThreadingModel(value bool) (err error) { 
    return instance.SetProperty("TrackThreadingModel", (value))
}

// GetTrackThreadingModel gets the value of TrackThreadingModel for the instance
func (instance *AspComPlus) GetPropertyTrackThreadingModel()(value bool, err error) { 
    retValue, err := instance.GetProperty("TrackThreadingModel")
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

