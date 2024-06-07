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

// DAMultiSite struct
type DAMultiSite struct { 
	*cim.WmiInstance

	// 
	DAEntryPoints []DAEntryPointBase

	// 
	EnterpriseName string

	// 
	GslbFqdn string

	// 
	ManualEntryPointSelectionAllowed string
}

	func NewDAMultiSiteEx1(instance *cim.WmiInstance) (newInstance *DAMultiSite, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &DAMultiSite {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewDAMultiSiteEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *DAMultiSite, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &DAMultiSite {
	WmiInstance: tmp,
	}
	return
	}
	

// SetDAEntryPoints sets the value of DAEntryPoints for the instance
func (instance *DAMultiSite) SetPropertyDAEntryPoints(value []DAEntryPointBase) (err error) { 
    return instance.SetProperty("DAEntryPoints", (value))
}

// GetDAEntryPoints gets the value of DAEntryPoints for the instance
func (instance *DAMultiSite) GetPropertyDAEntryPoints()(value []DAEntryPointBase, err error) { 
    retValue, err := instance.GetProperty("DAEntryPoints")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(DAEntryPointBase); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " DAEntryPointBase is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, DAEntryPointBase(valuetmp))
    }

    return
}

// SetEnterpriseName sets the value of EnterpriseName for the instance
func (instance *DAMultiSite) SetPropertyEnterpriseName(value string) (err error) { 
    return instance.SetProperty("EnterpriseName", (value))
}

// GetEnterpriseName gets the value of EnterpriseName for the instance
func (instance *DAMultiSite) GetPropertyEnterpriseName()(value string, err error) { 
    retValue, err := instance.GetProperty("EnterpriseName")
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

// SetGslbFqdn sets the value of GslbFqdn for the instance
func (instance *DAMultiSite) SetPropertyGslbFqdn(value string) (err error) { 
    return instance.SetProperty("GslbFqdn", (value))
}

// GetGslbFqdn gets the value of GslbFqdn for the instance
func (instance *DAMultiSite) GetPropertyGslbFqdn()(value string, err error) { 
    retValue, err := instance.GetProperty("GslbFqdn")
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

// SetManualEntryPointSelectionAllowed sets the value of ManualEntryPointSelectionAllowed for the instance
func (instance *DAMultiSite) SetPropertyManualEntryPointSelectionAllowed(value string) (err error) { 
    return instance.SetProperty("ManualEntryPointSelectionAllowed", (value))
}

// GetManualEntryPointSelectionAllowed gets the value of ManualEntryPointSelectionAllowed for the instance
func (instance *DAMultiSite) GetPropertyManualEntryPointSelectionAllowed()(value string, err error) { 
    retValue, err := instance.GetProperty("ManualEntryPointSelectionAllowed")
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

