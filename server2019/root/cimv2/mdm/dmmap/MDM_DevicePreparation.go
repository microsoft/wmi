// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MDM_DevicePreparation struct
type MDM_DevicePreparation struct { 
	*cim.WmiInstance

	// 
	InstanceID string

	// 
	PageEnabled bool

	// 
	PageSettings string

	// 
	PageStatus int32

	// 
	ParentID string
}

	func NewMDM_DevicePreparationEx1(instance *cim.WmiInstance) (newInstance *MDM_DevicePreparation, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MDM_DevicePreparation {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMDM_DevicePreparationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MDM_DevicePreparation, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MDM_DevicePreparation {
	WmiInstance: tmp,
	}
	return
	}
	

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_DevicePreparation) SetPropertyInstanceID(value string) (err error) { 
    return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_DevicePreparation) GetPropertyInstanceID()(value string, err error) { 
    retValue, err := instance.GetProperty("InstanceID")
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

// SetPageEnabled sets the value of PageEnabled for the instance
func (instance *MDM_DevicePreparation) SetPropertyPageEnabled(value bool) (err error) { 
    return instance.SetProperty("PageEnabled", (value))
}

// GetPageEnabled gets the value of PageEnabled for the instance
func (instance *MDM_DevicePreparation) GetPropertyPageEnabled()(value bool, err error) { 
    retValue, err := instance.GetProperty("PageEnabled")
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

// SetPageSettings sets the value of PageSettings for the instance
func (instance *MDM_DevicePreparation) SetPropertyPageSettings(value string) (err error) { 
    return instance.SetProperty("PageSettings", (value))
}

// GetPageSettings gets the value of PageSettings for the instance
func (instance *MDM_DevicePreparation) GetPropertyPageSettings()(value string, err error) { 
    retValue, err := instance.GetProperty("PageSettings")
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

// SetPageStatus sets the value of PageStatus for the instance
func (instance *MDM_DevicePreparation) SetPropertyPageStatus(value int32) (err error) { 
    return instance.SetProperty("PageStatus", (value))
}

// GetPageStatus gets the value of PageStatus for the instance
func (instance *MDM_DevicePreparation) GetPropertyPageStatus()(value int32, err error) { 
    retValue, err := instance.GetProperty("PageStatus")
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_DevicePreparation) SetPropertyParentID(value string) (err error) { 
    return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_DevicePreparation) GetPropertyParentID()(value string, err error) { 
    retValue, err := instance.GetProperty("ParentID")
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

