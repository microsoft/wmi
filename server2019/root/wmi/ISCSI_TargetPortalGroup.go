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

// ISCSI_TargetPortalGroup struct
type ISCSI_TargetPortalGroup struct { 
	*cim.WmiInstance

	// 
	PortalCount uint32

	// 
	Portals []ISCSI_TargetPortal
}

	func NewISCSI_TargetPortalGroupEx1(instance *cim.WmiInstance) (newInstance *ISCSI_TargetPortalGroup, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &ISCSI_TargetPortalGroup {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewISCSI_TargetPortalGroupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ISCSI_TargetPortalGroup, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ISCSI_TargetPortalGroup {
	WmiInstance: tmp,
	}
	return
	}
	

// SetPortalCount sets the value of PortalCount for the instance
func (instance *ISCSI_TargetPortalGroup) SetPropertyPortalCount(value uint32) (err error) { 
    return instance.SetProperty("PortalCount", (value))
}

// GetPortalCount gets the value of PortalCount for the instance
func (instance *ISCSI_TargetPortalGroup) GetPropertyPortalCount()(value uint32, err error) { 
    retValue, err := instance.GetProperty("PortalCount")
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

// SetPortals sets the value of Portals for the instance
func (instance *ISCSI_TargetPortalGroup) SetPropertyPortals(value []ISCSI_TargetPortal) (err error) { 
    return instance.SetProperty("Portals", (value))
}

// GetPortals gets the value of Portals for the instance
func (instance *ISCSI_TargetPortalGroup) GetPropertyPortals()(value []ISCSI_TargetPortal, err error) { 
    retValue, err := instance.GetProperty("Portals")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(ISCSI_TargetPortal); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " ISCSI_TargetPortal is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, ISCSI_TargetPortal(valuetmp))
    }

    return
}

