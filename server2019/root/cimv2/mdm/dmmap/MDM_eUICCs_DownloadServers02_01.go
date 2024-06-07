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

// MDM_eUICCs_DownloadServers02_01 struct
type MDM_eUICCs_DownloadServers02_01 struct { 
	*cim.WmiInstance

	// 
	AutoEnable bool

	// 
	DiscoveryState int32

	// 
	InstanceID string

	// 
	ParentID string
}

	func NewMDM_eUICCs_DownloadServers02_01Ex1(instance *cim.WmiInstance) (newInstance *MDM_eUICCs_DownloadServers02_01, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MDM_eUICCs_DownloadServers02_01 {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMDM_eUICCs_DownloadServers02_01Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MDM_eUICCs_DownloadServers02_01, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MDM_eUICCs_DownloadServers02_01 {
	WmiInstance: tmp,
	}
	return
	}
	

// SetAutoEnable sets the value of AutoEnable for the instance
func (instance *MDM_eUICCs_DownloadServers02_01) SetPropertyAutoEnable(value bool) (err error) { 
    return instance.SetProperty("AutoEnable", (value))
}

// GetAutoEnable gets the value of AutoEnable for the instance
func (instance *MDM_eUICCs_DownloadServers02_01) GetPropertyAutoEnable()(value bool, err error) { 
    retValue, err := instance.GetProperty("AutoEnable")
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

// SetDiscoveryState sets the value of DiscoveryState for the instance
func (instance *MDM_eUICCs_DownloadServers02_01) SetPropertyDiscoveryState(value int32) (err error) { 
    return instance.SetProperty("DiscoveryState", (value))
}

// GetDiscoveryState gets the value of DiscoveryState for the instance
func (instance *MDM_eUICCs_DownloadServers02_01) GetPropertyDiscoveryState()(value int32, err error) { 
    retValue, err := instance.GetProperty("DiscoveryState")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_eUICCs_DownloadServers02_01) SetPropertyInstanceID(value string) (err error) { 
    return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_eUICCs_DownloadServers02_01) GetPropertyInstanceID()(value string, err error) { 
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
func (instance *MDM_eUICCs_DownloadServers02_01) SetPropertyParentID(value string) (err error) { 
    return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_eUICCs_DownloadServers02_01) GetPropertyParentID()(value string, err error) { 
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

