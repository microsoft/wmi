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

// MSiSCSIInitiator_ConnectionInformation struct
type MSiSCSIInitiator_ConnectionInformation struct { 
	*cim.WmiInstance

	// 
	CID []uint8

	// 
	ConnectionID string

	// 
	InitiatorAddress string

	// 
	InitiatorPort uint16

	// 
	TargetAddress string

	// 
	TargetPort uint16
}

	func NewMSiSCSIInitiator_ConnectionInformationEx1(instance *cim.WmiInstance) (newInstance *MSiSCSIInitiator_ConnectionInformation, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MSiSCSIInitiator_ConnectionInformation {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMSiSCSIInitiator_ConnectionInformationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSiSCSIInitiator_ConnectionInformation, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSiSCSIInitiator_ConnectionInformation {
	WmiInstance: tmp,
	}
	return
	}
	

// SetCID sets the value of CID for the instance
func (instance *MSiSCSIInitiator_ConnectionInformation) SetPropertyCID(value []uint8) (err error) { 
    return instance.SetProperty("CID", (value))
}

// GetCID gets the value of CID for the instance
func (instance *MSiSCSIInitiator_ConnectionInformation) GetPropertyCID()(value []uint8, err error) { 
    retValue, err := instance.GetProperty("CID")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(uint8); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, uint8(valuetmp))
    }

    return
}

// SetConnectionID sets the value of ConnectionID for the instance
func (instance *MSiSCSIInitiator_ConnectionInformation) SetPropertyConnectionID(value string) (err error) { 
    return instance.SetProperty("ConnectionID", (value))
}

// GetConnectionID gets the value of ConnectionID for the instance
func (instance *MSiSCSIInitiator_ConnectionInformation) GetPropertyConnectionID()(value string, err error) { 
    retValue, err := instance.GetProperty("ConnectionID")
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

// SetInitiatorAddress sets the value of InitiatorAddress for the instance
func (instance *MSiSCSIInitiator_ConnectionInformation) SetPropertyInitiatorAddress(value string) (err error) { 
    return instance.SetProperty("InitiatorAddress", (value))
}

// GetInitiatorAddress gets the value of InitiatorAddress for the instance
func (instance *MSiSCSIInitiator_ConnectionInformation) GetPropertyInitiatorAddress()(value string, err error) { 
    retValue, err := instance.GetProperty("InitiatorAddress")
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

// SetInitiatorPort sets the value of InitiatorPort for the instance
func (instance *MSiSCSIInitiator_ConnectionInformation) SetPropertyInitiatorPort(value uint16) (err error) { 
    return instance.SetProperty("InitiatorPort", (value))
}

// GetInitiatorPort gets the value of InitiatorPort for the instance
func (instance *MSiSCSIInitiator_ConnectionInformation) GetPropertyInitiatorPort()(value uint16, err error) { 
    retValue, err := instance.GetProperty("InitiatorPort")
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

// SetTargetAddress sets the value of TargetAddress for the instance
func (instance *MSiSCSIInitiator_ConnectionInformation) SetPropertyTargetAddress(value string) (err error) { 
    return instance.SetProperty("TargetAddress", (value))
}

// GetTargetAddress gets the value of TargetAddress for the instance
func (instance *MSiSCSIInitiator_ConnectionInformation) GetPropertyTargetAddress()(value string, err error) { 
    retValue, err := instance.GetProperty("TargetAddress")
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

// SetTargetPort sets the value of TargetPort for the instance
func (instance *MSiSCSIInitiator_ConnectionInformation) SetPropertyTargetPort(value uint16) (err error) { 
    return instance.SetProperty("TargetPort", (value))
}

// GetTargetPort gets the value of TargetPort for the instance
func (instance *MSiSCSIInitiator_ConnectionInformation) GetPropertyTargetPort()(value uint16, err error) { 
    retValue, err := instance.GetProperty("TargetPort")
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

