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

// MSNdis_WmiTcpIpChecksumOffload_IPv6TransmitReceive struct
type MSNdis_WmiTcpIpChecksumOffload_IPv6TransmitReceive struct { 
	*MSNdis

	// 
	Encapsulation uint32

	// 
	IpExtensionHeadersSupported uint32

	// 
	TcpChecksum uint32

	// 
	TcpOptionsSupported uint32

	// 
	UdpChecksum uint32
}

	func NewMSNdis_WmiTcpIpChecksumOffload_IPv6TransmitReceiveEx1(instance *cim.WmiInstance) (newInstance *MSNdis_WmiTcpIpChecksumOffload_IPv6TransmitReceive, err error) {tmp, err := NewMSNdisEx1(instance)
		
	if err != nil { return }
	newInstance = &MSNdis_WmiTcpIpChecksumOffload_IPv6TransmitReceive {
	MSNdis: tmp,
	}
	return
	}
	

	func NewMSNdis_WmiTcpIpChecksumOffload_IPv6TransmitReceiveEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSNdis_WmiTcpIpChecksumOffload_IPv6TransmitReceive, err error) {tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSNdis_WmiTcpIpChecksumOffload_IPv6TransmitReceive {
	MSNdis: tmp,
	}
	return
	}
	

// SetEncapsulation sets the value of Encapsulation for the instance
func (instance *MSNdis_WmiTcpIpChecksumOffload_IPv6TransmitReceive) SetPropertyEncapsulation(value uint32) (err error) { 
    return instance.SetProperty("Encapsulation", (value))
}

// GetEncapsulation gets the value of Encapsulation for the instance
func (instance *MSNdis_WmiTcpIpChecksumOffload_IPv6TransmitReceive) GetPropertyEncapsulation()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Encapsulation")
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

// SetIpExtensionHeadersSupported sets the value of IpExtensionHeadersSupported for the instance
func (instance *MSNdis_WmiTcpIpChecksumOffload_IPv6TransmitReceive) SetPropertyIpExtensionHeadersSupported(value uint32) (err error) { 
    return instance.SetProperty("IpExtensionHeadersSupported", (value))
}

// GetIpExtensionHeadersSupported gets the value of IpExtensionHeadersSupported for the instance
func (instance *MSNdis_WmiTcpIpChecksumOffload_IPv6TransmitReceive) GetPropertyIpExtensionHeadersSupported()(value uint32, err error) { 
    retValue, err := instance.GetProperty("IpExtensionHeadersSupported")
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

// SetTcpChecksum sets the value of TcpChecksum for the instance
func (instance *MSNdis_WmiTcpIpChecksumOffload_IPv6TransmitReceive) SetPropertyTcpChecksum(value uint32) (err error) { 
    return instance.SetProperty("TcpChecksum", (value))
}

// GetTcpChecksum gets the value of TcpChecksum for the instance
func (instance *MSNdis_WmiTcpIpChecksumOffload_IPv6TransmitReceive) GetPropertyTcpChecksum()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TcpChecksum")
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

// SetTcpOptionsSupported sets the value of TcpOptionsSupported for the instance
func (instance *MSNdis_WmiTcpIpChecksumOffload_IPv6TransmitReceive) SetPropertyTcpOptionsSupported(value uint32) (err error) { 
    return instance.SetProperty("TcpOptionsSupported", (value))
}

// GetTcpOptionsSupported gets the value of TcpOptionsSupported for the instance
func (instance *MSNdis_WmiTcpIpChecksumOffload_IPv6TransmitReceive) GetPropertyTcpOptionsSupported()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TcpOptionsSupported")
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

// SetUdpChecksum sets the value of UdpChecksum for the instance
func (instance *MSNdis_WmiTcpIpChecksumOffload_IPv6TransmitReceive) SetPropertyUdpChecksum(value uint32) (err error) { 
    return instance.SetProperty("UdpChecksum", (value))
}

// GetUdpChecksum gets the value of UdpChecksum for the instance
func (instance *MSNdis_WmiTcpIpChecksumOffload_IPv6TransmitReceive) GetPropertyUdpChecksum()(value uint32, err error) { 
    retValue, err := instance.GetProperty("UdpChecksum")
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

