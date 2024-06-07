// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// Msvm_EthernetSwitchPortExtendedAclSettingData struct
type Msvm_EthernetSwitchPortExtendedAclSettingData struct { 
	*Msvm_EthernetSwitchPortFeatureSettingData

	// 
	Action uint8

	// 
	Direction uint8

	// 
	IdleSessionTimeout uint16

	// 
	IsolationID uint32

	// 
	LocalIPAddress string

	// 
	LocalPort string

	// 
	Name string

	// 
	Protocol string

	// 
	RemoteIPAddress string

	// 
	RemotePort string

	// 
	Stateful bool

	// 
	Weight uint16
}

	func NewMsvm_EthernetSwitchPortExtendedAclSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_EthernetSwitchPortExtendedAclSettingData, err error) {tmp, err := NewMsvm_EthernetSwitchPortFeatureSettingDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Msvm_EthernetSwitchPortExtendedAclSettingData {
	Msvm_EthernetSwitchPortFeatureSettingData: tmp,
	}
	return
	}
	

	func NewMsvm_EthernetSwitchPortExtendedAclSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Msvm_EthernetSwitchPortExtendedAclSettingData, err error) {tmp, err := NewMsvm_EthernetSwitchPortFeatureSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Msvm_EthernetSwitchPortExtendedAclSettingData {
	Msvm_EthernetSwitchPortFeatureSettingData: tmp,
	}
	return
	}
	

// SetAction sets the value of Action for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) SetPropertyAction(value uint8) (err error) { 
    return instance.SetProperty("Action", (value))
}

// GetAction gets the value of Action for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) GetPropertyAction()(value uint8, err error) { 
    retValue, err := instance.GetProperty("Action")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// SetDirection sets the value of Direction for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) SetPropertyDirection(value uint8) (err error) { 
    return instance.SetProperty("Direction", (value))
}

// GetDirection gets the value of Direction for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) GetPropertyDirection()(value uint8, err error) { 
    retValue, err := instance.GetProperty("Direction")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// SetIdleSessionTimeout sets the value of IdleSessionTimeout for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) SetPropertyIdleSessionTimeout(value uint16) (err error) { 
    return instance.SetProperty("IdleSessionTimeout", (value))
}

// GetIdleSessionTimeout gets the value of IdleSessionTimeout for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) GetPropertyIdleSessionTimeout()(value uint16, err error) { 
    retValue, err := instance.GetProperty("IdleSessionTimeout")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint16); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint16(valuetmp)

    return
}

// SetIsolationID sets the value of IsolationID for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) SetPropertyIsolationID(value uint32) (err error) { 
    return instance.SetProperty("IsolationID", (value))
}

// GetIsolationID gets the value of IsolationID for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) GetPropertyIsolationID()(value uint32, err error) { 
    retValue, err := instance.GetProperty("IsolationID")
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

// SetLocalIPAddress sets the value of LocalIPAddress for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) SetPropertyLocalIPAddress(value string) (err error) { 
    return instance.SetProperty("LocalIPAddress", (value))
}

// GetLocalIPAddress gets the value of LocalIPAddress for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) GetPropertyLocalIPAddress()(value string, err error) { 
    retValue, err := instance.GetProperty("LocalIPAddress")
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

// SetLocalPort sets the value of LocalPort for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) SetPropertyLocalPort(value string) (err error) { 
    return instance.SetProperty("LocalPort", (value))
}

// GetLocalPort gets the value of LocalPort for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) GetPropertyLocalPort()(value string, err error) { 
    retValue, err := instance.GetProperty("LocalPort")
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

// SetName sets the value of Name for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) SetPropertyName(value string) (err error) { 
    return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) GetPropertyName()(value string, err error) { 
    retValue, err := instance.GetProperty("Name")
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

// SetProtocol sets the value of Protocol for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) SetPropertyProtocol(value string) (err error) { 
    return instance.SetProperty("Protocol", (value))
}

// GetProtocol gets the value of Protocol for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) GetPropertyProtocol()(value string, err error) { 
    retValue, err := instance.GetProperty("Protocol")
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

// SetRemoteIPAddress sets the value of RemoteIPAddress for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) SetPropertyRemoteIPAddress(value string) (err error) { 
    return instance.SetProperty("RemoteIPAddress", (value))
}

// GetRemoteIPAddress gets the value of RemoteIPAddress for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) GetPropertyRemoteIPAddress()(value string, err error) { 
    retValue, err := instance.GetProperty("RemoteIPAddress")
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

// SetRemotePort sets the value of RemotePort for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) SetPropertyRemotePort(value string) (err error) { 
    return instance.SetProperty("RemotePort", (value))
}

// GetRemotePort gets the value of RemotePort for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) GetPropertyRemotePort()(value string, err error) { 
    retValue, err := instance.GetProperty("RemotePort")
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

// SetStateful sets the value of Stateful for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) SetPropertyStateful(value bool) (err error) { 
    return instance.SetProperty("Stateful", (value))
}

// GetStateful gets the value of Stateful for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) GetPropertyStateful()(value bool, err error) { 
    retValue, err := instance.GetProperty("Stateful")
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

// SetWeight sets the value of Weight for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) SetPropertyWeight(value uint16) (err error) { 
    return instance.SetProperty("Weight", (value))
}

// GetWeight gets the value of Weight for the instance
func (instance *Msvm_EthernetSwitchPortExtendedAclSettingData) GetPropertyWeight()(value uint16, err error) { 
    retValue, err := instance.GetProperty("Weight")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint16); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint16(valuetmp)

    return
}
func  (instance* Msvm_EthernetSwitchPortExtendedAclSettingData) GetRelatedEthernetSwitchFeatureCapabilities() (value *cim.WmiInstance, err error) {
		 return instance.GetRelated("Msvm_EthernetSwitchFeatureCapabilities"); 
	}
	

