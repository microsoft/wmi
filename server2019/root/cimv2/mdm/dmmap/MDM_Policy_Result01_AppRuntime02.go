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

// MDM_Policy_Result01_AppRuntime02 struct
type MDM_Policy_Result01_AppRuntime02 struct { 
	*cim.WmiInstance

	// 
	AllowMicrosoftAccountsToBeOptional string

	// 
	InstanceID string

	// 
	ParentID string
}

	func NewMDM_Policy_Result01_AppRuntime02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Result01_AppRuntime02, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MDM_Policy_Result01_AppRuntime02 {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMDM_Policy_Result01_AppRuntime02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MDM_Policy_Result01_AppRuntime02, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MDM_Policy_Result01_AppRuntime02 {
	WmiInstance: tmp,
	}
	return
	}
	

// SetAllowMicrosoftAccountsToBeOptional sets the value of AllowMicrosoftAccountsToBeOptional for the instance
func (instance *MDM_Policy_Result01_AppRuntime02) SetPropertyAllowMicrosoftAccountsToBeOptional(value string) (err error) { 
    return instance.SetProperty("AllowMicrosoftAccountsToBeOptional", (value))
}

// GetAllowMicrosoftAccountsToBeOptional gets the value of AllowMicrosoftAccountsToBeOptional for the instance
func (instance *MDM_Policy_Result01_AppRuntime02) GetPropertyAllowMicrosoftAccountsToBeOptional()(value string, err error) { 
    retValue, err := instance.GetProperty("AllowMicrosoftAccountsToBeOptional")
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

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_AppRuntime02) SetPropertyInstanceID(value string) (err error) { 
    return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_AppRuntime02) GetPropertyInstanceID()(value string, err error) { 
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_AppRuntime02) SetPropertyParentID(value string) (err error) { 
    return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_AppRuntime02) GetPropertyParentID()(value string, err error) { 
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

