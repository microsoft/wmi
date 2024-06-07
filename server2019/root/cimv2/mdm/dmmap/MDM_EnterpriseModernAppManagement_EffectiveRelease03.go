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

// MDM_EnterpriseModernAppManagement_EffectiveRelease03 struct
type MDM_EnterpriseModernAppManagement_EffectiveRelease03 struct { 
	*cim.WmiInstance

	// 
	ChannelId string

	// 
	InstanceID string

	// 
	ParentID string

	// 
	ReleaseManagementId string
}

	func NewMDM_EnterpriseModernAppManagement_EffectiveRelease03Ex1(instance *cim.WmiInstance) (newInstance *MDM_EnterpriseModernAppManagement_EffectiveRelease03, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MDM_EnterpriseModernAppManagement_EffectiveRelease03 {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMDM_EnterpriseModernAppManagement_EffectiveRelease03Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MDM_EnterpriseModernAppManagement_EffectiveRelease03, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MDM_EnterpriseModernAppManagement_EffectiveRelease03 {
	WmiInstance: tmp,
	}
	return
	}
	

// SetChannelId sets the value of ChannelId for the instance
func (instance *MDM_EnterpriseModernAppManagement_EffectiveRelease03) SetPropertyChannelId(value string) (err error) { 
    return instance.SetProperty("ChannelId", (value))
}

// GetChannelId gets the value of ChannelId for the instance
func (instance *MDM_EnterpriseModernAppManagement_EffectiveRelease03) GetPropertyChannelId()(value string, err error) { 
    retValue, err := instance.GetProperty("ChannelId")
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
func (instance *MDM_EnterpriseModernAppManagement_EffectiveRelease03) SetPropertyInstanceID(value string) (err error) { 
    return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_EnterpriseModernAppManagement_EffectiveRelease03) GetPropertyInstanceID()(value string, err error) { 
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
func (instance *MDM_EnterpriseModernAppManagement_EffectiveRelease03) SetPropertyParentID(value string) (err error) { 
    return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_EnterpriseModernAppManagement_EffectiveRelease03) GetPropertyParentID()(value string, err error) { 
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

// SetReleaseManagementId sets the value of ReleaseManagementId for the instance
func (instance *MDM_EnterpriseModernAppManagement_EffectiveRelease03) SetPropertyReleaseManagementId(value string) (err error) { 
    return instance.SetProperty("ReleaseManagementId", (value))
}

// GetReleaseManagementId gets the value of ReleaseManagementId for the instance
func (instance *MDM_EnterpriseModernAppManagement_EffectiveRelease03) GetPropertyReleaseManagementId()(value string, err error) { 
    retValue, err := instance.GetProperty("ReleaseManagementId")
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

