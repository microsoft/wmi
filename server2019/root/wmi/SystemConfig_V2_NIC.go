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
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// SystemConfig_V2_NIC struct
type SystemConfig_V2_NIC struct { 
	*SystemConfig_V2

	// 
	DnsServerAddresses string

	// 
	IpAddresses string

	// 
	Ipv4Index uint32

	// 
	Ipv6Index uint32

	// 
	NICDescription string

	// 
	PhysicalAddr uint64

	// 
	PhysicalAddrLen uint32
}

	func NewSystemConfig_V2_NICEx1(instance *cim.WmiInstance) (newInstance *SystemConfig_V2_NIC, err error) {tmp, err := NewSystemConfig_V2Ex1(instance)
		
	if err != nil { return }
	newInstance = &SystemConfig_V2_NIC {
	SystemConfig_V2: tmp,
	}
	return
	}
	

	func NewSystemConfig_V2_NICEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SystemConfig_V2_NIC, err error) {tmp, err := NewSystemConfig_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SystemConfig_V2_NIC {
	SystemConfig_V2: tmp,
	}
	return
	}
	

// SetDnsServerAddresses sets the value of DnsServerAddresses for the instance
func (instance *SystemConfig_V2_NIC) SetPropertyDnsServerAddresses(value string) (err error) { 
    return instance.SetProperty("DnsServerAddresses", (value))
}

// GetDnsServerAddresses gets the value of DnsServerAddresses for the instance
func (instance *SystemConfig_V2_NIC) GetPropertyDnsServerAddresses()(value string, err error) { 
    retValue, err := instance.GetProperty("DnsServerAddresses")
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

// SetIpAddresses sets the value of IpAddresses for the instance
func (instance *SystemConfig_V2_NIC) SetPropertyIpAddresses(value string) (err error) { 
    return instance.SetProperty("IpAddresses", (value))
}

// GetIpAddresses gets the value of IpAddresses for the instance
func (instance *SystemConfig_V2_NIC) GetPropertyIpAddresses()(value string, err error) { 
    retValue, err := instance.GetProperty("IpAddresses")
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

// SetIpv4Index sets the value of Ipv4Index for the instance
func (instance *SystemConfig_V2_NIC) SetPropertyIpv4Index(value uint32) (err error) { 
    return instance.SetProperty("Ipv4Index", (value))
}

// GetIpv4Index gets the value of Ipv4Index for the instance
func (instance *SystemConfig_V2_NIC) GetPropertyIpv4Index()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Ipv4Index")
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

// SetIpv6Index sets the value of Ipv6Index for the instance
func (instance *SystemConfig_V2_NIC) SetPropertyIpv6Index(value uint32) (err error) { 
    return instance.SetProperty("Ipv6Index", (value))
}

// GetIpv6Index gets the value of Ipv6Index for the instance
func (instance *SystemConfig_V2_NIC) GetPropertyIpv6Index()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Ipv6Index")
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

// SetNICDescription sets the value of NICDescription for the instance
func (instance *SystemConfig_V2_NIC) SetPropertyNICDescription(value string) (err error) { 
    return instance.SetProperty("NICDescription", (value))
}

// GetNICDescription gets the value of NICDescription for the instance
func (instance *SystemConfig_V2_NIC) GetPropertyNICDescription()(value string, err error) { 
    retValue, err := instance.GetProperty("NICDescription")
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

// SetPhysicalAddr sets the value of PhysicalAddr for the instance
func (instance *SystemConfig_V2_NIC) SetPropertyPhysicalAddr(value uint64) (err error) { 
    return instance.SetProperty("PhysicalAddr", (value))
}

// GetPhysicalAddr gets the value of PhysicalAddr for the instance
func (instance *SystemConfig_V2_NIC) GetPropertyPhysicalAddr()(value uint64, err error) { 
    retValue, err := instance.GetProperty("PhysicalAddr")
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

// SetPhysicalAddrLen sets the value of PhysicalAddrLen for the instance
func (instance *SystemConfig_V2_NIC) SetPropertyPhysicalAddrLen(value uint32) (err error) { 
    return instance.SetProperty("PhysicalAddrLen", (value))
}

// GetPhysicalAddrLen gets the value of PhysicalAddrLen for the instance
func (instance *SystemConfig_V2_NIC) GetPropertyPhysicalAddrLen()(value uint32, err error) { 
    retValue, err := instance.GetProperty("PhysicalAddrLen")
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

