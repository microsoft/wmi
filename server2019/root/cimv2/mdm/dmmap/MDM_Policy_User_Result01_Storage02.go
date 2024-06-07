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

// MDM_Policy_User_Result01_Storage02 struct
type MDM_Policy_User_Result01_Storage02 struct { 
	*cim.WmiInstance

	// 
	InstanceID string

	// 
	ParentID string

	// 
	WPDDevicesDenyReadAccessPerUser string

	// 
	WPDDevicesDenyWriteAccessPerUser string
}

	func NewMDM_Policy_User_Result01_Storage02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_User_Result01_Storage02, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MDM_Policy_User_Result01_Storage02 {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMDM_Policy_User_Result01_Storage02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MDM_Policy_User_Result01_Storage02, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MDM_Policy_User_Result01_Storage02 {
	WmiInstance: tmp,
	}
	return
	}
	

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_User_Result01_Storage02) SetPropertyInstanceID(value string) (err error) { 
    return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_User_Result01_Storage02) GetPropertyInstanceID()(value string, err error) { 
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
func (instance *MDM_Policy_User_Result01_Storage02) SetPropertyParentID(value string) (err error) { 
    return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_User_Result01_Storage02) GetPropertyParentID()(value string, err error) { 
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

// SetWPDDevicesDenyReadAccessPerUser sets the value of WPDDevicesDenyReadAccessPerUser for the instance
func (instance *MDM_Policy_User_Result01_Storage02) SetPropertyWPDDevicesDenyReadAccessPerUser(value string) (err error) { 
    return instance.SetProperty("WPDDevicesDenyReadAccessPerUser", (value))
}

// GetWPDDevicesDenyReadAccessPerUser gets the value of WPDDevicesDenyReadAccessPerUser for the instance
func (instance *MDM_Policy_User_Result01_Storage02) GetPropertyWPDDevicesDenyReadAccessPerUser()(value string, err error) { 
    retValue, err := instance.GetProperty("WPDDevicesDenyReadAccessPerUser")
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

// SetWPDDevicesDenyWriteAccessPerUser sets the value of WPDDevicesDenyWriteAccessPerUser for the instance
func (instance *MDM_Policy_User_Result01_Storage02) SetPropertyWPDDevicesDenyWriteAccessPerUser(value string) (err error) { 
    return instance.SetProperty("WPDDevicesDenyWriteAccessPerUser", (value))
}

// GetWPDDevicesDenyWriteAccessPerUser gets the value of WPDDevicesDenyWriteAccessPerUser for the instance
func (instance *MDM_Policy_User_Result01_Storage02) GetPropertyWPDDevicesDenyWriteAccessPerUser()(value string, err error) { 
    retValue, err := instance.GetProperty("WPDDevicesDenyWriteAccessPerUser")
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

