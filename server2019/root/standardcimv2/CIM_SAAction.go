// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// CIM_SAAction struct
type CIM_SAAction struct { 
	*CIM_PolicyAction

	// 
	DoPacketLogging bool
}

	func NewCIM_SAActionEx1(instance *cim.WmiInstance) (newInstance *CIM_SAAction, err error) {tmp, err := NewCIM_PolicyActionEx1(instance)
		
	if err != nil { return }
	newInstance = &CIM_SAAction {
	CIM_PolicyAction: tmp,
	}
	return
	}
	

	func NewCIM_SAActionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *CIM_SAAction, err error) {tmp, err := NewCIM_PolicyActionEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &CIM_SAAction {
	CIM_PolicyAction: tmp,
	}
	return
	}
	

// SetDoPacketLogging sets the value of DoPacketLogging for the instance
func (instance *CIM_SAAction) SetPropertyDoPacketLogging(value bool) (err error) { 
    return instance.SetProperty("DoPacketLogging", (value))
}

// GetDoPacketLogging gets the value of DoPacketLogging for the instance
func (instance *CIM_SAAction) GetPropertyDoPacketLogging()(value bool, err error) { 
    retValue, err := instance.GetProperty("DoPacketLogging")
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

