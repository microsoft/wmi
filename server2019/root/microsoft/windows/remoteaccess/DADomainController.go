// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess
//////////////////////////////////////////////
package remoteaccess
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// DADomainController struct
type DADomainController struct { 
	*cim.WmiInstance

	// 
	DomainControllerName string

	// 
	EntryPointName string
}

	func NewDADomainControllerEx1(instance *cim.WmiInstance) (newInstance *DADomainController, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &DADomainController {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewDADomainControllerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *DADomainController, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &DADomainController {
	WmiInstance: tmp,
	}
	return
	}
	

// SetDomainControllerName sets the value of DomainControllerName for the instance
func (instance *DADomainController) SetPropertyDomainControllerName(value string) (err error) { 
    return instance.SetProperty("DomainControllerName", (value))
}

// GetDomainControllerName gets the value of DomainControllerName for the instance
func (instance *DADomainController) GetPropertyDomainControllerName()(value string, err error) { 
    retValue, err := instance.GetProperty("DomainControllerName")
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

// SetEntryPointName sets the value of EntryPointName for the instance
func (instance *DADomainController) SetPropertyEntryPointName(value string) (err error) { 
    return instance.SetProperty("EntryPointName", (value))
}

// GetEntryPointName gets the value of EntryPointName for the instance
func (instance *DADomainController) GetPropertyEntryPointName()(value string, err error) { 
    retValue, err := instance.GetProperty("EntryPointName")
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

