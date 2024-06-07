// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.msdtc
//////////////////////////////////////////////
package msdtc
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// DtcClusterTMMapping struct
type DtcClusterTMMapping struct { 
	*cim.WmiInstance

	// 
	Application string

	// 
	ApplicationType string

	// 
	ClusterResourceName string

	// 
	Local bool

	// 
	Name string
}

	func NewDtcClusterTMMappingEx1(instance *cim.WmiInstance) (newInstance *DtcClusterTMMapping, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &DtcClusterTMMapping {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewDtcClusterTMMappingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *DtcClusterTMMapping, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &DtcClusterTMMapping {
	WmiInstance: tmp,
	}
	return
	}
	

// SetApplication sets the value of Application for the instance
func (instance *DtcClusterTMMapping) SetPropertyApplication(value string) (err error) { 
    return instance.SetProperty("Application", (value))
}

// GetApplication gets the value of Application for the instance
func (instance *DtcClusterTMMapping) GetPropertyApplication()(value string, err error) { 
    retValue, err := instance.GetProperty("Application")
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

// SetApplicationType sets the value of ApplicationType for the instance
func (instance *DtcClusterTMMapping) SetPropertyApplicationType(value string) (err error) { 
    return instance.SetProperty("ApplicationType", (value))
}

// GetApplicationType gets the value of ApplicationType for the instance
func (instance *DtcClusterTMMapping) GetPropertyApplicationType()(value string, err error) { 
    retValue, err := instance.GetProperty("ApplicationType")
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

// SetClusterResourceName sets the value of ClusterResourceName for the instance
func (instance *DtcClusterTMMapping) SetPropertyClusterResourceName(value string) (err error) { 
    return instance.SetProperty("ClusterResourceName", (value))
}

// GetClusterResourceName gets the value of ClusterResourceName for the instance
func (instance *DtcClusterTMMapping) GetPropertyClusterResourceName()(value string, err error) { 
    retValue, err := instance.GetProperty("ClusterResourceName")
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

// SetLocal sets the value of Local for the instance
func (instance *DtcClusterTMMapping) SetPropertyLocal(value bool) (err error) { 
    return instance.SetProperty("Local", (value))
}

// GetLocal gets the value of Local for the instance
func (instance *DtcClusterTMMapping) GetPropertyLocal()(value bool, err error) { 
    retValue, err := instance.GetProperty("Local")
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

// SetName sets the value of Name for the instance
func (instance *DtcClusterTMMapping) SetPropertyName(value string) (err error) { 
    return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *DtcClusterTMMapping) GetPropertyName()(value string, err error) { 
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

