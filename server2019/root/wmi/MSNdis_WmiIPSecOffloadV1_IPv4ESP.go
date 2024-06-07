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

// MSNdis_WmiIPSecOffloadV1_IPv4ESP struct
type MSNdis_WmiIPSecOffloadV1_IPv4ESP struct { 
	*MSNdis

	// 
	Des uint32

	// 
	NullEsp uint32

	// 
	Receive uint32

	// 
	Reserved uint32

	// 
	Send uint32

	// 
	Transport uint32

	// 
	TripleDes uint32

	// 
	Tunnel uint32
}

	func NewMSNdis_WmiIPSecOffloadV1_IPv4ESPEx1(instance *cim.WmiInstance) (newInstance *MSNdis_WmiIPSecOffloadV1_IPv4ESP, err error) {tmp, err := NewMSNdisEx1(instance)
		
	if err != nil { return }
	newInstance = &MSNdis_WmiIPSecOffloadV1_IPv4ESP {
	MSNdis: tmp,
	}
	return
	}
	

	func NewMSNdis_WmiIPSecOffloadV1_IPv4ESPEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSNdis_WmiIPSecOffloadV1_IPv4ESP, err error) {tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSNdis_WmiIPSecOffloadV1_IPv4ESP {
	MSNdis: tmp,
	}
	return
	}
	

// SetDes sets the value of Des for the instance
func (instance *MSNdis_WmiIPSecOffloadV1_IPv4ESP) SetPropertyDes(value uint32) (err error) { 
    return instance.SetProperty("Des", (value))
}

// GetDes gets the value of Des for the instance
func (instance *MSNdis_WmiIPSecOffloadV1_IPv4ESP) GetPropertyDes()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Des")
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

// SetNullEsp sets the value of NullEsp for the instance
func (instance *MSNdis_WmiIPSecOffloadV1_IPv4ESP) SetPropertyNullEsp(value uint32) (err error) { 
    return instance.SetProperty("NullEsp", (value))
}

// GetNullEsp gets the value of NullEsp for the instance
func (instance *MSNdis_WmiIPSecOffloadV1_IPv4ESP) GetPropertyNullEsp()(value uint32, err error) { 
    retValue, err := instance.GetProperty("NullEsp")
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

// SetReceive sets the value of Receive for the instance
func (instance *MSNdis_WmiIPSecOffloadV1_IPv4ESP) SetPropertyReceive(value uint32) (err error) { 
    return instance.SetProperty("Receive", (value))
}

// GetReceive gets the value of Receive for the instance
func (instance *MSNdis_WmiIPSecOffloadV1_IPv4ESP) GetPropertyReceive()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Receive")
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

// SetReserved sets the value of Reserved for the instance
func (instance *MSNdis_WmiIPSecOffloadV1_IPv4ESP) SetPropertyReserved(value uint32) (err error) { 
    return instance.SetProperty("Reserved", (value))
}

// GetReserved gets the value of Reserved for the instance
func (instance *MSNdis_WmiIPSecOffloadV1_IPv4ESP) GetPropertyReserved()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Reserved")
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

// SetSend sets the value of Send for the instance
func (instance *MSNdis_WmiIPSecOffloadV1_IPv4ESP) SetPropertySend(value uint32) (err error) { 
    return instance.SetProperty("Send", (value))
}

// GetSend gets the value of Send for the instance
func (instance *MSNdis_WmiIPSecOffloadV1_IPv4ESP) GetPropertySend()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Send")
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

// SetTransport sets the value of Transport for the instance
func (instance *MSNdis_WmiIPSecOffloadV1_IPv4ESP) SetPropertyTransport(value uint32) (err error) { 
    return instance.SetProperty("Transport", (value))
}

// GetTransport gets the value of Transport for the instance
func (instance *MSNdis_WmiIPSecOffloadV1_IPv4ESP) GetPropertyTransport()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Transport")
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

// SetTripleDes sets the value of TripleDes for the instance
func (instance *MSNdis_WmiIPSecOffloadV1_IPv4ESP) SetPropertyTripleDes(value uint32) (err error) { 
    return instance.SetProperty("TripleDes", (value))
}

// GetTripleDes gets the value of TripleDes for the instance
func (instance *MSNdis_WmiIPSecOffloadV1_IPv4ESP) GetPropertyTripleDes()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TripleDes")
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

// SetTunnel sets the value of Tunnel for the instance
func (instance *MSNdis_WmiIPSecOffloadV1_IPv4ESP) SetPropertyTunnel(value uint32) (err error) { 
    return instance.SetProperty("Tunnel", (value))
}

// GetTunnel gets the value of Tunnel for the instance
func (instance *MSNdis_WmiIPSecOffloadV1_IPv4ESP) GetPropertyTunnel()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Tunnel")
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

