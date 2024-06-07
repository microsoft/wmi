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

// Msvm_EthernetSwitchPortAclSettingData struct
type Msvm_EthernetSwitchPortAclSettingData struct { 
	*Msvm_EthernetSwitchPortFeatureSettingData

	// 
	AclType uint8

	// 
	Action uint8

	// 
	Applicability uint8

	// 
	Direction uint8

	// 
	LocalAddress string

	// 
	LocalAddressPrefixLength uint8

	// 
	Name string

	// 
	RemoteAddress string

	// 
	RemoteAddressPrefixLength uint8
}

	func NewMsvm_EthernetSwitchPortAclSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_EthernetSwitchPortAclSettingData, err error) {tmp, err := NewMsvm_EthernetSwitchPortFeatureSettingDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Msvm_EthernetSwitchPortAclSettingData {
	Msvm_EthernetSwitchPortFeatureSettingData: tmp,
	}
	return
	}
	

	func NewMsvm_EthernetSwitchPortAclSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Msvm_EthernetSwitchPortAclSettingData, err error) {tmp, err := NewMsvm_EthernetSwitchPortFeatureSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Msvm_EthernetSwitchPortAclSettingData {
	Msvm_EthernetSwitchPortFeatureSettingData: tmp,
	}
	return
	}
	

// SetAclType sets the value of AclType for the instance
func (instance *Msvm_EthernetSwitchPortAclSettingData) SetPropertyAclType(value uint8) (err error) { 
    return instance.SetProperty("AclType", (value))
}

// GetAclType gets the value of AclType for the instance
func (instance *Msvm_EthernetSwitchPortAclSettingData) GetPropertyAclType()(value uint8, err error) { 
    retValue, err := instance.GetProperty("AclType")
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

// SetAction sets the value of Action for the instance
func (instance *Msvm_EthernetSwitchPortAclSettingData) SetPropertyAction(value uint8) (err error) { 
    return instance.SetProperty("Action", (value))
}

// GetAction gets the value of Action for the instance
func (instance *Msvm_EthernetSwitchPortAclSettingData) GetPropertyAction()(value uint8, err error) { 
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

// SetApplicability sets the value of Applicability for the instance
func (instance *Msvm_EthernetSwitchPortAclSettingData) SetPropertyApplicability(value uint8) (err error) { 
    return instance.SetProperty("Applicability", (value))
}

// GetApplicability gets the value of Applicability for the instance
func (instance *Msvm_EthernetSwitchPortAclSettingData) GetPropertyApplicability()(value uint8, err error) { 
    retValue, err := instance.GetProperty("Applicability")
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
func (instance *Msvm_EthernetSwitchPortAclSettingData) SetPropertyDirection(value uint8) (err error) { 
    return instance.SetProperty("Direction", (value))
}

// GetDirection gets the value of Direction for the instance
func (instance *Msvm_EthernetSwitchPortAclSettingData) GetPropertyDirection()(value uint8, err error) { 
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

// SetLocalAddress sets the value of LocalAddress for the instance
func (instance *Msvm_EthernetSwitchPortAclSettingData) SetPropertyLocalAddress(value string) (err error) { 
    return instance.SetProperty("LocalAddress", (value))
}

// GetLocalAddress gets the value of LocalAddress for the instance
func (instance *Msvm_EthernetSwitchPortAclSettingData) GetPropertyLocalAddress()(value string, err error) { 
    retValue, err := instance.GetProperty("LocalAddress")
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

// SetLocalAddressPrefixLength sets the value of LocalAddressPrefixLength for the instance
func (instance *Msvm_EthernetSwitchPortAclSettingData) SetPropertyLocalAddressPrefixLength(value uint8) (err error) { 
    return instance.SetProperty("LocalAddressPrefixLength", (value))
}

// GetLocalAddressPrefixLength gets the value of LocalAddressPrefixLength for the instance
func (instance *Msvm_EthernetSwitchPortAclSettingData) GetPropertyLocalAddressPrefixLength()(value uint8, err error) { 
    retValue, err := instance.GetProperty("LocalAddressPrefixLength")
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

// SetName sets the value of Name for the instance
func (instance *Msvm_EthernetSwitchPortAclSettingData) SetPropertyName(value string) (err error) { 
    return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *Msvm_EthernetSwitchPortAclSettingData) GetPropertyName()(value string, err error) { 
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

// SetRemoteAddress sets the value of RemoteAddress for the instance
func (instance *Msvm_EthernetSwitchPortAclSettingData) SetPropertyRemoteAddress(value string) (err error) { 
    return instance.SetProperty("RemoteAddress", (value))
}

// GetRemoteAddress gets the value of RemoteAddress for the instance
func (instance *Msvm_EthernetSwitchPortAclSettingData) GetPropertyRemoteAddress()(value string, err error) { 
    retValue, err := instance.GetProperty("RemoteAddress")
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

// SetRemoteAddressPrefixLength sets the value of RemoteAddressPrefixLength for the instance
func (instance *Msvm_EthernetSwitchPortAclSettingData) SetPropertyRemoteAddressPrefixLength(value uint8) (err error) { 
    return instance.SetProperty("RemoteAddressPrefixLength", (value))
}

// GetRemoteAddressPrefixLength gets the value of RemoteAddressPrefixLength for the instance
func (instance *Msvm_EthernetSwitchPortAclSettingData) GetPropertyRemoteAddressPrefixLength()(value uint8, err error) { 
    retValue, err := instance.GetProperty("RemoteAddressPrefixLength")
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
func  (instance* Msvm_EthernetSwitchPortAclSettingData) GetRelatedEthernetSwitchFeatureCapabilities() (value *cim.WmiInstance, err error) {
		 return instance.GetRelated("Msvm_EthernetSwitchFeatureCapabilities"); 
	}
	

