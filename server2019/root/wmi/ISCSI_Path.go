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

// ISCSI_Path struct
type ISCSI_Path struct { 
	*cim.WmiInstance

	// Status of the path - connected, disconnected, reconnecting
	ConnectionStatus Path_ConnectionStatus

	// Estimated speed of the connection in MegaBits Per Second
	EstimatedLinkSpeed uint64

	// Weight assigned to the path
	PathWeight uint32

	// Flag set to 1 if the path is a primary path, 0 otherwise.
	PrimaryPath uint32

	// Flag set to 1 if TCP offload is supported for this connection, 0 otherwise.
	TCPOffLoadAvailable uint32

	// iSCSI Unique connection id
	UniqueConnectionId uint64
}

	func NewISCSI_PathEx1(instance *cim.WmiInstance) (newInstance *ISCSI_Path, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &ISCSI_Path {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewISCSI_PathEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ISCSI_Path, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ISCSI_Path {
	WmiInstance: tmp,
	}
	return
	}
	

// SetConnectionStatus sets the value of ConnectionStatus for the instance
func (instance *ISCSI_Path) SetPropertyConnectionStatus(value Path_ConnectionStatus) (err error) { 
    return instance.SetProperty("ConnectionStatus", (value))
}

// GetConnectionStatus gets the value of ConnectionStatus for the instance
func (instance *ISCSI_Path) GetPropertyConnectionStatus()(value Path_ConnectionStatus, err error) { 
    retValue, err := instance.GetProperty("ConnectionStatus")
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

    value = Path_ConnectionStatus(valuetmp)

    return
}

// SetEstimatedLinkSpeed sets the value of EstimatedLinkSpeed for the instance
func (instance *ISCSI_Path) SetPropertyEstimatedLinkSpeed(value uint64) (err error) { 
    return instance.SetProperty("EstimatedLinkSpeed", (value))
}

// GetEstimatedLinkSpeed gets the value of EstimatedLinkSpeed for the instance
func (instance *ISCSI_Path) GetPropertyEstimatedLinkSpeed()(value uint64, err error) { 
    retValue, err := instance.GetProperty("EstimatedLinkSpeed")
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

// SetPathWeight sets the value of PathWeight for the instance
func (instance *ISCSI_Path) SetPropertyPathWeight(value uint32) (err error) { 
    return instance.SetProperty("PathWeight", (value))
}

// GetPathWeight gets the value of PathWeight for the instance
func (instance *ISCSI_Path) GetPropertyPathWeight()(value uint32, err error) { 
    retValue, err := instance.GetProperty("PathWeight")
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

// SetPrimaryPath sets the value of PrimaryPath for the instance
func (instance *ISCSI_Path) SetPropertyPrimaryPath(value uint32) (err error) { 
    return instance.SetProperty("PrimaryPath", (value))
}

// GetPrimaryPath gets the value of PrimaryPath for the instance
func (instance *ISCSI_Path) GetPropertyPrimaryPath()(value uint32, err error) { 
    retValue, err := instance.GetProperty("PrimaryPath")
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

// SetTCPOffLoadAvailable sets the value of TCPOffLoadAvailable for the instance
func (instance *ISCSI_Path) SetPropertyTCPOffLoadAvailable(value uint32) (err error) { 
    return instance.SetProperty("TCPOffLoadAvailable", (value))
}

// GetTCPOffLoadAvailable gets the value of TCPOffLoadAvailable for the instance
func (instance *ISCSI_Path) GetPropertyTCPOffLoadAvailable()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TCPOffLoadAvailable")
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

// SetUniqueConnectionId sets the value of UniqueConnectionId for the instance
func (instance *ISCSI_Path) SetPropertyUniqueConnectionId(value uint64) (err error) { 
    return instance.SetProperty("UniqueConnectionId", (value))
}

// GetUniqueConnectionId gets the value of UniqueConnectionId for the instance
func (instance *ISCSI_Path) GetPropertyUniqueConnectionId()(value uint64, err error) { 
    retValue, err := instance.GetProperty("UniqueConnectionId")
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

