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

// MDM_VPNv2_PluginProfile02 struct
type MDM_VPNv2_PluginProfile02 struct { 
	*cim.WmiInstance

	// 
	CustomConfiguration string

	// 
	InstanceID string

	// 
	ParentID string

	// 
	PluginPackageFamilyName string

	// 
	ServerUrlList string
}

	func NewMDM_VPNv2_PluginProfile02Ex1(instance *cim.WmiInstance) (newInstance *MDM_VPNv2_PluginProfile02, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MDM_VPNv2_PluginProfile02 {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMDM_VPNv2_PluginProfile02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MDM_VPNv2_PluginProfile02, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MDM_VPNv2_PluginProfile02 {
	WmiInstance: tmp,
	}
	return
	}
	

// SetCustomConfiguration sets the value of CustomConfiguration for the instance
func (instance *MDM_VPNv2_PluginProfile02) SetPropertyCustomConfiguration(value string) (err error) { 
    return instance.SetProperty("CustomConfiguration", (value))
}

// GetCustomConfiguration gets the value of CustomConfiguration for the instance
func (instance *MDM_VPNv2_PluginProfile02) GetPropertyCustomConfiguration()(value string, err error) { 
    retValue, err := instance.GetProperty("CustomConfiguration")
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
func (instance *MDM_VPNv2_PluginProfile02) SetPropertyInstanceID(value string) (err error) { 
    return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_VPNv2_PluginProfile02) GetPropertyInstanceID()(value string, err error) { 
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
func (instance *MDM_VPNv2_PluginProfile02) SetPropertyParentID(value string) (err error) { 
    return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_VPNv2_PluginProfile02) GetPropertyParentID()(value string, err error) { 
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

// SetPluginPackageFamilyName sets the value of PluginPackageFamilyName for the instance
func (instance *MDM_VPNv2_PluginProfile02) SetPropertyPluginPackageFamilyName(value string) (err error) { 
    return instance.SetProperty("PluginPackageFamilyName", (value))
}

// GetPluginPackageFamilyName gets the value of PluginPackageFamilyName for the instance
func (instance *MDM_VPNv2_PluginProfile02) GetPropertyPluginPackageFamilyName()(value string, err error) { 
    retValue, err := instance.GetProperty("PluginPackageFamilyName")
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

// SetServerUrlList sets the value of ServerUrlList for the instance
func (instance *MDM_VPNv2_PluginProfile02) SetPropertyServerUrlList(value string) (err error) { 
    return instance.SetProperty("ServerUrlList", (value))
}

// GetServerUrlList gets the value of ServerUrlList for the instance
func (instance *MDM_VPNv2_PluginProfile02) GetPropertyServerUrlList()(value string, err error) { 
    retValue, err := instance.GetProperty("ServerUrlList")
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

