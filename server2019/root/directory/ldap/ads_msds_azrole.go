// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// ads_msds_azrole struct
type ads_msds_azrole struct { 
	*ds_top

	// 
	DS_msDS_AzApplicationData string

	// 
	DS_msDS_AzGenericData string

	// 
	DS_msDS_AzObjectGuid Uint8Array

	// 
	DS_msDS_MembersForAzRole []string

	// 
	DS_msDS_OperationsForAzRole []string

	// 
	DS_msDS_TasksForAzRole []string
}

	func Newads_msds_azroleEx1(instance *cim.WmiInstance) (newInstance *ads_msds_azrole, err error) {tmp, err := Newds_topEx1(instance)
		
	if err != nil { return }
	newInstance = &ads_msds_azrole {
	ds_top: tmp,
	}
	return
	}
	

	func Newads_msds_azroleEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ads_msds_azrole, err error) {tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ads_msds_azrole {
	ds_top: tmp,
	}
	return
	}
	

// SetDS_msDS_AzApplicationData sets the value of DS_msDS_AzApplicationData for the instance
func (instance *ads_msds_azrole) SetPropertyDS_msDS_AzApplicationData(value string) (err error) { 
    return instance.SetProperty("DS_msDS_AzApplicationData", (value))
}

// GetDS_msDS_AzApplicationData gets the value of DS_msDS_AzApplicationData for the instance
func (instance *ads_msds_azrole) GetPropertyDS_msDS_AzApplicationData()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_AzApplicationData")
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

// SetDS_msDS_AzGenericData sets the value of DS_msDS_AzGenericData for the instance
func (instance *ads_msds_azrole) SetPropertyDS_msDS_AzGenericData(value string) (err error) { 
    return instance.SetProperty("DS_msDS_AzGenericData", (value))
}

// GetDS_msDS_AzGenericData gets the value of DS_msDS_AzGenericData for the instance
func (instance *ads_msds_azrole) GetPropertyDS_msDS_AzGenericData()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_AzGenericData")
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

// SetDS_msDS_AzObjectGuid sets the value of DS_msDS_AzObjectGuid for the instance
func (instance *ads_msds_azrole) SetPropertyDS_msDS_AzObjectGuid(value Uint8Array) (err error) { 
    return instance.SetProperty("DS_msDS_AzObjectGuid", (value))
}

// GetDS_msDS_AzObjectGuid gets the value of DS_msDS_AzObjectGuid for the instance
func (instance *ads_msds_azrole) GetPropertyDS_msDS_AzObjectGuid()(value Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_AzObjectGuid")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(Uint8Array); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = Uint8Array(valuetmp)

    return
}

// SetDS_msDS_MembersForAzRole sets the value of DS_msDS_MembersForAzRole for the instance
func (instance *ads_msds_azrole) SetPropertyDS_msDS_MembersForAzRole(value []string) (err error) { 
    return instance.SetProperty("DS_msDS_MembersForAzRole", (value))
}

// GetDS_msDS_MembersForAzRole gets the value of DS_msDS_MembersForAzRole for the instance
func (instance *ads_msds_azrole) GetPropertyDS_msDS_MembersForAzRole()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_MembersForAzRole")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(string); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, string(valuetmp))
    }

    return
}

// SetDS_msDS_OperationsForAzRole sets the value of DS_msDS_OperationsForAzRole for the instance
func (instance *ads_msds_azrole) SetPropertyDS_msDS_OperationsForAzRole(value []string) (err error) { 
    return instance.SetProperty("DS_msDS_OperationsForAzRole", (value))
}

// GetDS_msDS_OperationsForAzRole gets the value of DS_msDS_OperationsForAzRole for the instance
func (instance *ads_msds_azrole) GetPropertyDS_msDS_OperationsForAzRole()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_OperationsForAzRole")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(string); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, string(valuetmp))
    }

    return
}

// SetDS_msDS_TasksForAzRole sets the value of DS_msDS_TasksForAzRole for the instance
func (instance *ads_msds_azrole) SetPropertyDS_msDS_TasksForAzRole(value []string) (err error) { 
    return instance.SetProperty("DS_msDS_TasksForAzRole", (value))
}

// GetDS_msDS_TasksForAzRole gets the value of DS_msDS_TasksForAzRole for the instance
func (instance *ads_msds_azrole) GetPropertyDS_msDS_TasksForAzRole()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_TasksForAzRole")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(string); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, string(valuetmp))
    }

    return
}

