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

// MSiSCSIInitiator_TargetClass struct
type MSiSCSIInitiator_TargetClass struct { 
	*cim.WmiInstance

	// 
	DiscoveryMechanism string

	// 
	InitiatorName string

	// 
	LoginOptions MSiSCSIInitiator_TargetLoginOptions

	// 
	Mappings MSiSCSIInitiator_TargetMappings

	// 
	PortalGroups []MSiSCSIInitiator_PortalGroup

	// 
	ProtocolType TargetClass_ProtocolType

	// 
	TargetAlias string

	// 
	TargetFlags uint32

	// 
	TargetName string
}

	func NewMSiSCSIInitiator_TargetClassEx1(instance *cim.WmiInstance) (newInstance *MSiSCSIInitiator_TargetClass, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MSiSCSIInitiator_TargetClass {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMSiSCSIInitiator_TargetClassEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSiSCSIInitiator_TargetClass, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSiSCSIInitiator_TargetClass {
	WmiInstance: tmp,
	}
	return
	}
	

// SetDiscoveryMechanism sets the value of DiscoveryMechanism for the instance
func (instance *MSiSCSIInitiator_TargetClass) SetPropertyDiscoveryMechanism(value string) (err error) { 
    return instance.SetProperty("DiscoveryMechanism", (value))
}

// GetDiscoveryMechanism gets the value of DiscoveryMechanism for the instance
func (instance *MSiSCSIInitiator_TargetClass) GetPropertyDiscoveryMechanism()(value string, err error) { 
    retValue, err := instance.GetProperty("DiscoveryMechanism")
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

// SetInitiatorName sets the value of InitiatorName for the instance
func (instance *MSiSCSIInitiator_TargetClass) SetPropertyInitiatorName(value string) (err error) { 
    return instance.SetProperty("InitiatorName", (value))
}

// GetInitiatorName gets the value of InitiatorName for the instance
func (instance *MSiSCSIInitiator_TargetClass) GetPropertyInitiatorName()(value string, err error) { 
    retValue, err := instance.GetProperty("InitiatorName")
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

// SetLoginOptions sets the value of LoginOptions for the instance
func (instance *MSiSCSIInitiator_TargetClass) SetPropertyLoginOptions(value MSiSCSIInitiator_TargetLoginOptions) (err error) { 
    return instance.SetProperty("LoginOptions", (value))
}

// GetLoginOptions gets the value of LoginOptions for the instance
func (instance *MSiSCSIInitiator_TargetClass) GetPropertyLoginOptions()(value MSiSCSIInitiator_TargetLoginOptions, err error) { 
    retValue, err := instance.GetProperty("LoginOptions")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(MSiSCSIInitiator_TargetLoginOptions); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " MSiSCSIInitiator_TargetLoginOptions is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = MSiSCSIInitiator_TargetLoginOptions(valuetmp)

    return
}

// SetMappings sets the value of Mappings for the instance
func (instance *MSiSCSIInitiator_TargetClass) SetPropertyMappings(value MSiSCSIInitiator_TargetMappings) (err error) { 
    return instance.SetProperty("Mappings", (value))
}

// GetMappings gets the value of Mappings for the instance
func (instance *MSiSCSIInitiator_TargetClass) GetPropertyMappings()(value MSiSCSIInitiator_TargetMappings, err error) { 
    retValue, err := instance.GetProperty("Mappings")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(MSiSCSIInitiator_TargetMappings); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " MSiSCSIInitiator_TargetMappings is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = MSiSCSIInitiator_TargetMappings(valuetmp)

    return
}

// SetPortalGroups sets the value of PortalGroups for the instance
func (instance *MSiSCSIInitiator_TargetClass) SetPropertyPortalGroups(value []MSiSCSIInitiator_PortalGroup) (err error) { 
    return instance.SetProperty("PortalGroups", (value))
}

// GetPortalGroups gets the value of PortalGroups for the instance
func (instance *MSiSCSIInitiator_TargetClass) GetPropertyPortalGroups()(value []MSiSCSIInitiator_PortalGroup, err error) { 
    retValue, err := instance.GetProperty("PortalGroups")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(MSiSCSIInitiator_PortalGroup); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " MSiSCSIInitiator_PortalGroup is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, MSiSCSIInitiator_PortalGroup(valuetmp))
    }

    return
}

// SetProtocolType sets the value of ProtocolType for the instance
func (instance *MSiSCSIInitiator_TargetClass) SetPropertyProtocolType(value TargetClass_ProtocolType) (err error) { 
    return instance.SetProperty("ProtocolType", (value))
}

// GetProtocolType gets the value of ProtocolType for the instance
func (instance *MSiSCSIInitiator_TargetClass) GetPropertyProtocolType()(value TargetClass_ProtocolType, err error) { 
    retValue, err := instance.GetProperty("ProtocolType")
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

    value = TargetClass_ProtocolType(valuetmp)

    return
}

// SetTargetAlias sets the value of TargetAlias for the instance
func (instance *MSiSCSIInitiator_TargetClass) SetPropertyTargetAlias(value string) (err error) { 
    return instance.SetProperty("TargetAlias", (value))
}

// GetTargetAlias gets the value of TargetAlias for the instance
func (instance *MSiSCSIInitiator_TargetClass) GetPropertyTargetAlias()(value string, err error) { 
    retValue, err := instance.GetProperty("TargetAlias")
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

// SetTargetFlags sets the value of TargetFlags for the instance
func (instance *MSiSCSIInitiator_TargetClass) SetPropertyTargetFlags(value uint32) (err error) { 
    return instance.SetProperty("TargetFlags", (value))
}

// GetTargetFlags gets the value of TargetFlags for the instance
func (instance *MSiSCSIInitiator_TargetClass) GetPropertyTargetFlags()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TargetFlags")
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

// SetTargetName sets the value of TargetName for the instance
func (instance *MSiSCSIInitiator_TargetClass) SetPropertyTargetName(value string) (err error) { 
    return instance.SetProperty("TargetName", (value))
}

// GetTargetName gets the value of TargetName for the instance
func (instance *MSiSCSIInitiator_TargetClass) GetPropertyTargetName()(value string, err error) { 
    retValue, err := instance.GetProperty("TargetName")
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

// 

// <param name="InitiatorPortNumber" type="uint32 "></param>
// <param name="IsInformationalSession" type="bool "></param>
// <param name="IsPersistent" type="bool "></param>
// <param name="key" type="uint8 []"></param>
// <param name="LoginOptions" type="MSiSCSIInitiator_TargetLoginOptions "></param>
// <param name="Mappings" type="MSiSCSIInitiator_TargetMappings []"></param>
// <param name="SecurityFlags" type="uint64 "></param>
// <param name="TargetPortal" type="MSiSCSIInitiator_Portal "></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="UniqueConnectionId" type="string "></param>
// <param name="UniqueSessionId" type="string "></param>
func (instance *MSiSCSIInitiator_TargetClass) Login( /* IN */ IsInformationalSession bool,
 /* IN */ InitiatorPortNumber uint32,
 /* IN */ TargetPortal MSiSCSIInitiator_Portal,
 /* IN */ SecurityFlags uint64,
 /* IN */ Mappings []MSiSCSIInitiator_TargetMappings,
 /* IN */ LoginOptions MSiSCSIInitiator_TargetLoginOptions,
 /* IN */ key []uint8,
 /* IN */ IsPersistent bool,
 /* OUT */ UniqueSessionId string,
 /* OUT */ UniqueConnectionId string) (result uint32, err error) {retVal, err := instance.InvokeMethod("Login" , IsInformationalSession, InitiatorPortNumber, TargetPortal, SecurityFlags, Mappings, LoginOptions, key, IsPersistent)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


