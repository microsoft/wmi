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

// MDM_Policy_Result01_WirelessDisplay02 struct
type MDM_Policy_Result01_WirelessDisplay02 struct { 
	*cim.WmiInstance

	// 
	AllowMdnsAdvertisement int32

	// 
	AllowMdnsDiscovery int32

	// 
	AllowMovementDetectionOnInfrastructure int32

	// 
	AllowPCReceiverToBeTCPServer int32

	// 
	AllowPCSenderToBeTCPClient int32

	// 
	AllowProjectionFromPC int32

	// 
	AllowProjectionFromPCOverInfrastructure int32

	// 
	AllowProjectionToPC int32

	// 
	AllowProjectionToPCOverInfrastructure int32

	// 
	AllowUserInputFromWirelessDisplayReceiver int32

	// 
	InstanceID string

	// 
	ParentID string

	// 
	RequirePinForPairing int32
}

	func NewMDM_Policy_Result01_WirelessDisplay02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Result01_WirelessDisplay02, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MDM_Policy_Result01_WirelessDisplay02 {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMDM_Policy_Result01_WirelessDisplay02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MDM_Policy_Result01_WirelessDisplay02, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MDM_Policy_Result01_WirelessDisplay02 {
	WmiInstance: tmp,
	}
	return
	}
	

// SetAllowMdnsAdvertisement sets the value of AllowMdnsAdvertisement for the instance
func (instance *MDM_Policy_Result01_WirelessDisplay02) SetPropertyAllowMdnsAdvertisement(value int32) (err error) { 
    return instance.SetProperty("AllowMdnsAdvertisement", (value))
}

// GetAllowMdnsAdvertisement gets the value of AllowMdnsAdvertisement for the instance
func (instance *MDM_Policy_Result01_WirelessDisplay02) GetPropertyAllowMdnsAdvertisement()(value int32, err error) { 
    retValue, err := instance.GetProperty("AllowMdnsAdvertisement")
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

// SetAllowMdnsDiscovery sets the value of AllowMdnsDiscovery for the instance
func (instance *MDM_Policy_Result01_WirelessDisplay02) SetPropertyAllowMdnsDiscovery(value int32) (err error) { 
    return instance.SetProperty("AllowMdnsDiscovery", (value))
}

// GetAllowMdnsDiscovery gets the value of AllowMdnsDiscovery for the instance
func (instance *MDM_Policy_Result01_WirelessDisplay02) GetPropertyAllowMdnsDiscovery()(value int32, err error) { 
    retValue, err := instance.GetProperty("AllowMdnsDiscovery")
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

// SetAllowMovementDetectionOnInfrastructure sets the value of AllowMovementDetectionOnInfrastructure for the instance
func (instance *MDM_Policy_Result01_WirelessDisplay02) SetPropertyAllowMovementDetectionOnInfrastructure(value int32) (err error) { 
    return instance.SetProperty("AllowMovementDetectionOnInfrastructure", (value))
}

// GetAllowMovementDetectionOnInfrastructure gets the value of AllowMovementDetectionOnInfrastructure for the instance
func (instance *MDM_Policy_Result01_WirelessDisplay02) GetPropertyAllowMovementDetectionOnInfrastructure()(value int32, err error) { 
    retValue, err := instance.GetProperty("AllowMovementDetectionOnInfrastructure")
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

// SetAllowPCReceiverToBeTCPServer sets the value of AllowPCReceiverToBeTCPServer for the instance
func (instance *MDM_Policy_Result01_WirelessDisplay02) SetPropertyAllowPCReceiverToBeTCPServer(value int32) (err error) { 
    return instance.SetProperty("AllowPCReceiverToBeTCPServer", (value))
}

// GetAllowPCReceiverToBeTCPServer gets the value of AllowPCReceiverToBeTCPServer for the instance
func (instance *MDM_Policy_Result01_WirelessDisplay02) GetPropertyAllowPCReceiverToBeTCPServer()(value int32, err error) { 
    retValue, err := instance.GetProperty("AllowPCReceiverToBeTCPServer")
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

// SetAllowPCSenderToBeTCPClient sets the value of AllowPCSenderToBeTCPClient for the instance
func (instance *MDM_Policy_Result01_WirelessDisplay02) SetPropertyAllowPCSenderToBeTCPClient(value int32) (err error) { 
    return instance.SetProperty("AllowPCSenderToBeTCPClient", (value))
}

// GetAllowPCSenderToBeTCPClient gets the value of AllowPCSenderToBeTCPClient for the instance
func (instance *MDM_Policy_Result01_WirelessDisplay02) GetPropertyAllowPCSenderToBeTCPClient()(value int32, err error) { 
    retValue, err := instance.GetProperty("AllowPCSenderToBeTCPClient")
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

// SetAllowProjectionFromPC sets the value of AllowProjectionFromPC for the instance
func (instance *MDM_Policy_Result01_WirelessDisplay02) SetPropertyAllowProjectionFromPC(value int32) (err error) { 
    return instance.SetProperty("AllowProjectionFromPC", (value))
}

// GetAllowProjectionFromPC gets the value of AllowProjectionFromPC for the instance
func (instance *MDM_Policy_Result01_WirelessDisplay02) GetPropertyAllowProjectionFromPC()(value int32, err error) { 
    retValue, err := instance.GetProperty("AllowProjectionFromPC")
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

// SetAllowProjectionFromPCOverInfrastructure sets the value of AllowProjectionFromPCOverInfrastructure for the instance
func (instance *MDM_Policy_Result01_WirelessDisplay02) SetPropertyAllowProjectionFromPCOverInfrastructure(value int32) (err error) { 
    return instance.SetProperty("AllowProjectionFromPCOverInfrastructure", (value))
}

// GetAllowProjectionFromPCOverInfrastructure gets the value of AllowProjectionFromPCOverInfrastructure for the instance
func (instance *MDM_Policy_Result01_WirelessDisplay02) GetPropertyAllowProjectionFromPCOverInfrastructure()(value int32, err error) { 
    retValue, err := instance.GetProperty("AllowProjectionFromPCOverInfrastructure")
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

// SetAllowProjectionToPC sets the value of AllowProjectionToPC for the instance
func (instance *MDM_Policy_Result01_WirelessDisplay02) SetPropertyAllowProjectionToPC(value int32) (err error) { 
    return instance.SetProperty("AllowProjectionToPC", (value))
}

// GetAllowProjectionToPC gets the value of AllowProjectionToPC for the instance
func (instance *MDM_Policy_Result01_WirelessDisplay02) GetPropertyAllowProjectionToPC()(value int32, err error) { 
    retValue, err := instance.GetProperty("AllowProjectionToPC")
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

// SetAllowProjectionToPCOverInfrastructure sets the value of AllowProjectionToPCOverInfrastructure for the instance
func (instance *MDM_Policy_Result01_WirelessDisplay02) SetPropertyAllowProjectionToPCOverInfrastructure(value int32) (err error) { 
    return instance.SetProperty("AllowProjectionToPCOverInfrastructure", (value))
}

// GetAllowProjectionToPCOverInfrastructure gets the value of AllowProjectionToPCOverInfrastructure for the instance
func (instance *MDM_Policy_Result01_WirelessDisplay02) GetPropertyAllowProjectionToPCOverInfrastructure()(value int32, err error) { 
    retValue, err := instance.GetProperty("AllowProjectionToPCOverInfrastructure")
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

// SetAllowUserInputFromWirelessDisplayReceiver sets the value of AllowUserInputFromWirelessDisplayReceiver for the instance
func (instance *MDM_Policy_Result01_WirelessDisplay02) SetPropertyAllowUserInputFromWirelessDisplayReceiver(value int32) (err error) { 
    return instance.SetProperty("AllowUserInputFromWirelessDisplayReceiver", (value))
}

// GetAllowUserInputFromWirelessDisplayReceiver gets the value of AllowUserInputFromWirelessDisplayReceiver for the instance
func (instance *MDM_Policy_Result01_WirelessDisplay02) GetPropertyAllowUserInputFromWirelessDisplayReceiver()(value int32, err error) { 
    retValue, err := instance.GetProperty("AllowUserInputFromWirelessDisplayReceiver")
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
func (instance *MDM_Policy_Result01_WirelessDisplay02) SetPropertyInstanceID(value string) (err error) { 
    return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_WirelessDisplay02) GetPropertyInstanceID()(value string, err error) { 
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
func (instance *MDM_Policy_Result01_WirelessDisplay02) SetPropertyParentID(value string) (err error) { 
    return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_WirelessDisplay02) GetPropertyParentID()(value string, err error) { 
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

// SetRequirePinForPairing sets the value of RequirePinForPairing for the instance
func (instance *MDM_Policy_Result01_WirelessDisplay02) SetPropertyRequirePinForPairing(value int32) (err error) { 
    return instance.SetProperty("RequirePinForPairing", (value))
}

// GetRequirePinForPairing gets the value of RequirePinForPairing for the instance
func (instance *MDM_Policy_Result01_WirelessDisplay02) GetPropertyRequirePinForPairing()(value int32, err error) { 
    retValue, err := instance.GetProperty("RequirePinForPairing")
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

