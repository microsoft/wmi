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

// SystemConfig_V1_NIC struct
type SystemConfig_V1_NIC struct { 
	*SystemConfig_V1

	// 
	Data uint32

	// 
	DhcpServer int32

	// 
	DnsServer1 int32

	// 
	DnsServer2 int32

	// 
	DnsServer3 int32

	// 
	DnsServer4 int32

	// 
	Gateway int32

	// 
	Index uint32

	// 
	IpAddress int32

	// 
	NICName []byte

	// 
	PhysicalAddr []byte

	// 
	PhysicalAddrLen uint32

	// 
	PrimaryWinsServer int32

	// 
	SecondaryWinsServer int32

	// 
	Size uint32

	// 
	SubnetMask int32
}

	func NewSystemConfig_V1_NICEx1(instance *cim.WmiInstance) (newInstance *SystemConfig_V1_NIC, err error) {tmp, err := NewSystemConfig_V1Ex1(instance)
		
	if err != nil { return }
	newInstance = &SystemConfig_V1_NIC {
	SystemConfig_V1: tmp,
	}
	return
	}
	

	func NewSystemConfig_V1_NICEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SystemConfig_V1_NIC, err error) {tmp, err := NewSystemConfig_V1Ex6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SystemConfig_V1_NIC {
	SystemConfig_V1: tmp,
	}
	return
	}
	

// SetData sets the value of Data for the instance
func (instance *SystemConfig_V1_NIC) SetPropertyData(value uint32) (err error) { 
    return instance.SetProperty("Data", (value))
}

// GetData gets the value of Data for the instance
func (instance *SystemConfig_V1_NIC) GetPropertyData()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Data")
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

// SetDhcpServer sets the value of DhcpServer for the instance
func (instance *SystemConfig_V1_NIC) SetPropertyDhcpServer(value int32) (err error) { 
    return instance.SetProperty("DhcpServer", (value))
}

// GetDhcpServer gets the value of DhcpServer for the instance
func (instance *SystemConfig_V1_NIC) GetPropertyDhcpServer()(value int32, err error) { 
    retValue, err := instance.GetProperty("DhcpServer")
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

// SetDnsServer1 sets the value of DnsServer1 for the instance
func (instance *SystemConfig_V1_NIC) SetPropertyDnsServer1(value int32) (err error) { 
    return instance.SetProperty("DnsServer1", (value))
}

// GetDnsServer1 gets the value of DnsServer1 for the instance
func (instance *SystemConfig_V1_NIC) GetPropertyDnsServer1()(value int32, err error) { 
    retValue, err := instance.GetProperty("DnsServer1")
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

// SetDnsServer2 sets the value of DnsServer2 for the instance
func (instance *SystemConfig_V1_NIC) SetPropertyDnsServer2(value int32) (err error) { 
    return instance.SetProperty("DnsServer2", (value))
}

// GetDnsServer2 gets the value of DnsServer2 for the instance
func (instance *SystemConfig_V1_NIC) GetPropertyDnsServer2()(value int32, err error) { 
    retValue, err := instance.GetProperty("DnsServer2")
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

// SetDnsServer3 sets the value of DnsServer3 for the instance
func (instance *SystemConfig_V1_NIC) SetPropertyDnsServer3(value int32) (err error) { 
    return instance.SetProperty("DnsServer3", (value))
}

// GetDnsServer3 gets the value of DnsServer3 for the instance
func (instance *SystemConfig_V1_NIC) GetPropertyDnsServer3()(value int32, err error) { 
    retValue, err := instance.GetProperty("DnsServer3")
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

// SetDnsServer4 sets the value of DnsServer4 for the instance
func (instance *SystemConfig_V1_NIC) SetPropertyDnsServer4(value int32) (err error) { 
    return instance.SetProperty("DnsServer4", (value))
}

// GetDnsServer4 gets the value of DnsServer4 for the instance
func (instance *SystemConfig_V1_NIC) GetPropertyDnsServer4()(value int32, err error) { 
    retValue, err := instance.GetProperty("DnsServer4")
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

// SetGateway sets the value of Gateway for the instance
func (instance *SystemConfig_V1_NIC) SetPropertyGateway(value int32) (err error) { 
    return instance.SetProperty("Gateway", (value))
}

// GetGateway gets the value of Gateway for the instance
func (instance *SystemConfig_V1_NIC) GetPropertyGateway()(value int32, err error) { 
    retValue, err := instance.GetProperty("Gateway")
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

// SetIndex sets the value of Index for the instance
func (instance *SystemConfig_V1_NIC) SetPropertyIndex(value uint32) (err error) { 
    return instance.SetProperty("Index", (value))
}

// GetIndex gets the value of Index for the instance
func (instance *SystemConfig_V1_NIC) GetPropertyIndex()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Index")
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

// SetIpAddress sets the value of IpAddress for the instance
func (instance *SystemConfig_V1_NIC) SetPropertyIpAddress(value int32) (err error) { 
    return instance.SetProperty("IpAddress", (value))
}

// GetIpAddress gets the value of IpAddress for the instance
func (instance *SystemConfig_V1_NIC) GetPropertyIpAddress()(value int32, err error) { 
    retValue, err := instance.GetProperty("IpAddress")
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

// SetNICName sets the value of NICName for the instance
func (instance *SystemConfig_V1_NIC) SetPropertyNICName(value []byte) (err error) { 
    return instance.SetProperty("NICName", (value))
}

// GetNICName gets the value of NICName for the instance
func (instance *SystemConfig_V1_NIC) GetPropertyNICName()(value []byte, err error) { 
    retValue, err := instance.GetProperty("NICName")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(byte); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " byte is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, byte(valuetmp))
    }

    return
}

// SetPhysicalAddr sets the value of PhysicalAddr for the instance
func (instance *SystemConfig_V1_NIC) SetPropertyPhysicalAddr(value []byte) (err error) { 
    return instance.SetProperty("PhysicalAddr", (value))
}

// GetPhysicalAddr gets the value of PhysicalAddr for the instance
func (instance *SystemConfig_V1_NIC) GetPropertyPhysicalAddr()(value []byte, err error) { 
    retValue, err := instance.GetProperty("PhysicalAddr")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(byte); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " byte is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, byte(valuetmp))
    }

    return
}

// SetPhysicalAddrLen sets the value of PhysicalAddrLen for the instance
func (instance *SystemConfig_V1_NIC) SetPropertyPhysicalAddrLen(value uint32) (err error) { 
    return instance.SetProperty("PhysicalAddrLen", (value))
}

// GetPhysicalAddrLen gets the value of PhysicalAddrLen for the instance
func (instance *SystemConfig_V1_NIC) GetPropertyPhysicalAddrLen()(value uint32, err error) { 
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

// SetPrimaryWinsServer sets the value of PrimaryWinsServer for the instance
func (instance *SystemConfig_V1_NIC) SetPropertyPrimaryWinsServer(value int32) (err error) { 
    return instance.SetProperty("PrimaryWinsServer", (value))
}

// GetPrimaryWinsServer gets the value of PrimaryWinsServer for the instance
func (instance *SystemConfig_V1_NIC) GetPropertyPrimaryWinsServer()(value int32, err error) { 
    retValue, err := instance.GetProperty("PrimaryWinsServer")
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

// SetSecondaryWinsServer sets the value of SecondaryWinsServer for the instance
func (instance *SystemConfig_V1_NIC) SetPropertySecondaryWinsServer(value int32) (err error) { 
    return instance.SetProperty("SecondaryWinsServer", (value))
}

// GetSecondaryWinsServer gets the value of SecondaryWinsServer for the instance
func (instance *SystemConfig_V1_NIC) GetPropertySecondaryWinsServer()(value int32, err error) { 
    retValue, err := instance.GetProperty("SecondaryWinsServer")
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

// SetSize sets the value of Size for the instance
func (instance *SystemConfig_V1_NIC) SetPropertySize(value uint32) (err error) { 
    return instance.SetProperty("Size", (value))
}

// GetSize gets the value of Size for the instance
func (instance *SystemConfig_V1_NIC) GetPropertySize()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Size")
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

// SetSubnetMask sets the value of SubnetMask for the instance
func (instance *SystemConfig_V1_NIC) SetPropertySubnetMask(value int32) (err error) { 
    return instance.SetProperty("SubnetMask", (value))
}

// GetSubnetMask gets the value of SubnetMask for the instance
func (instance *SystemConfig_V1_NIC) GetPropertySubnetMask()(value int32, err error) { 
    retValue, err := instance.GetProperty("SubnetMask")
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

