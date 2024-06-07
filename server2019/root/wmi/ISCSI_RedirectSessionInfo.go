// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// ISCSI_RedirectSessionInfo struct
type ISCSI_RedirectSessionInfo struct { 
	*cim.WmiInstance

	// Number of elements in RedirectPortalList array
	ConnectionCount uint32

	// Redirect portal info - one element for each connection in the session
	RedirectPortalList []ISCSI_RedirectPortalInfo

	// Target portal group tag for this Session 
	TargetPortalGroupTag uint32

	// A uniquely generated session ID, it is the same id returned by the LoginToTarget method.  Do not confuse this with ISID or SSID.
	UniqueSessionId uint64
}

	func NewISCSI_RedirectSessionInfoEx1(instance *cim.WmiInstance) (newInstance *ISCSI_RedirectSessionInfo, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &ISCSI_RedirectSessionInfo {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewISCSI_RedirectSessionInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ISCSI_RedirectSessionInfo, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ISCSI_RedirectSessionInfo {
	WmiInstance: tmp,
	}
	return
	}
	

// SetConnectionCount sets the value of ConnectionCount for the instance
func (instance *ISCSI_RedirectSessionInfo) SetPropertyConnectionCount(value uint32) (err error) { 
    return instance.SetProperty("ConnectionCount", (value))
}

// GetConnectionCount gets the value of ConnectionCount for the instance
func (instance *ISCSI_RedirectSessionInfo) GetPropertyConnectionCount()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ConnectionCount")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetRedirectPortalList sets the value of RedirectPortalList for the instance
func (instance *ISCSI_RedirectSessionInfo) SetPropertyRedirectPortalList(value []ISCSI_RedirectPortalInfo) (err error) { 
    return instance.SetProperty("RedirectPortalList", (value))
}

// GetRedirectPortalList gets the value of RedirectPortalList for the instance
func (instance *ISCSI_RedirectSessionInfo) GetPropertyRedirectPortalList()(value []ISCSI_RedirectPortalInfo, err error) { 
    retValue, err := instance.GetProperty("RedirectPortalList")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(ISCSI_RedirectPortalInfo); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " ISCSI_RedirectPortalInfo is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, ISCSI_RedirectPortalInfo(valuetmp))
    }

    return
}

// SetTargetPortalGroupTag sets the value of TargetPortalGroupTag for the instance
func (instance *ISCSI_RedirectSessionInfo) SetPropertyTargetPortalGroupTag(value uint32) (err error) { 
    return instance.SetProperty("TargetPortalGroupTag", (value))
}

// GetTargetPortalGroupTag gets the value of TargetPortalGroupTag for the instance
func (instance *ISCSI_RedirectSessionInfo) GetPropertyTargetPortalGroupTag()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TargetPortalGroupTag")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetUniqueSessionId sets the value of UniqueSessionId for the instance
func (instance *ISCSI_RedirectSessionInfo) SetPropertyUniqueSessionId(value uint64) (err error) { 
    return instance.SetProperty("UniqueSessionId", (value))
}

// GetUniqueSessionId gets the value of UniqueSessionId for the instance
func (instance *ISCSI_RedirectSessionInfo) GetPropertyUniqueSessionId()(value uint64, err error) { 
    retValue, err := instance.GetProperty("UniqueSessionId")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

