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

// MSiSCSIInitiator_Portal struct
type MSiSCSIInitiator_Portal struct { 
	*cim.WmiInstance

	// 
	Address string

	// 
	Index uint32

	// 
	Port uint16

	// 
	SymbolicName string
}

	func NewMSiSCSIInitiator_PortalEx1(instance *cim.WmiInstance) (newInstance *MSiSCSIInitiator_Portal, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MSiSCSIInitiator_Portal {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMSiSCSIInitiator_PortalEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSiSCSIInitiator_Portal, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSiSCSIInitiator_Portal {
	WmiInstance: tmp,
	}
	return
	}
	

// SetAddress sets the value of Address for the instance
func (instance *MSiSCSIInitiator_Portal) SetPropertyAddress(value string) (err error) { 
    return instance.SetProperty("Address", (value))
}

// GetAddress gets the value of Address for the instance
func (instance *MSiSCSIInitiator_Portal) GetPropertyAddress()(value string, err error) { 
    retValue, err := instance.GetProperty("Address")
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

// SetIndex sets the value of Index for the instance
func (instance *MSiSCSIInitiator_Portal) SetPropertyIndex(value uint32) (err error) { 
    return instance.SetProperty("Index", (value))
}

// GetIndex gets the value of Index for the instance
func (instance *MSiSCSIInitiator_Portal) GetPropertyIndex()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Index")
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

// SetPort sets the value of Port for the instance
func (instance *MSiSCSIInitiator_Portal) SetPropertyPort(value uint16) (err error) { 
    return instance.SetProperty("Port", (value))
}

// GetPort gets the value of Port for the instance
func (instance *MSiSCSIInitiator_Portal) GetPropertyPort()(value uint16, err error) { 
    retValue, err := instance.GetProperty("Port")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint16); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint16(valuetmp)

    return
}

// SetSymbolicName sets the value of SymbolicName for the instance
func (instance *MSiSCSIInitiator_Portal) SetPropertySymbolicName(value string) (err error) { 
    return instance.SetProperty("SymbolicName", (value))
}

// GetSymbolicName gets the value of SymbolicName for the instance
func (instance *MSiSCSIInitiator_Portal) GetPropertySymbolicName()(value string, err error) { 
    retValue, err := instance.GetProperty("SymbolicName")
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

