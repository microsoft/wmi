// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSFT_NetEventPacketCaptureProvider struct
type MSFT_NetEventPacketCaptureProvider struct { 
	*MSFT_NetEventProviderBase

	// 
	CaptureType uint8

	// 
	EtherType []uint16

	// 
	IPAddresses []string

	// 
	IPProtocols []uint8

	// 
	LinkLayerAddress []string

	// 
	MultiLayer bool

	// 
	TruncationLength uint16

	// 
	VmCaptureDirection uint8
}

	func NewMSFT_NetEventPacketCaptureProviderEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetEventPacketCaptureProvider, err error) {tmp, err := NewMSFT_NetEventProviderBaseEx1(instance)
		
	if err != nil { return }
	newInstance = &MSFT_NetEventPacketCaptureProvider {
	MSFT_NetEventProviderBase: tmp,
	}
	return
	}
	

	func NewMSFT_NetEventPacketCaptureProviderEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_NetEventPacketCaptureProvider, err error) {tmp, err := NewMSFT_NetEventProviderBaseEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_NetEventPacketCaptureProvider {
	MSFT_NetEventProviderBase: tmp,
	}
	return
	}
	

// SetCaptureType sets the value of CaptureType for the instance
func (instance *MSFT_NetEventPacketCaptureProvider) SetPropertyCaptureType(value uint8) (err error) { 
    return instance.SetProperty("CaptureType", (value))
}

// GetCaptureType gets the value of CaptureType for the instance
func (instance *MSFT_NetEventPacketCaptureProvider) GetPropertyCaptureType()(value uint8, err error) { 
    retValue, err := instance.GetProperty("CaptureType")
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

// SetEtherType sets the value of EtherType for the instance
func (instance *MSFT_NetEventPacketCaptureProvider) SetPropertyEtherType(value []uint16) (err error) { 
    return instance.SetProperty("EtherType", (value))
}

// GetEtherType gets the value of EtherType for the instance
func (instance *MSFT_NetEventPacketCaptureProvider) GetPropertyEtherType()(value []uint16, err error) { 
    retValue, err := instance.GetProperty("EtherType")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(uint16); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, uint16(valuetmp))
    }

    return
}

// SetIPAddresses sets the value of IPAddresses for the instance
func (instance *MSFT_NetEventPacketCaptureProvider) SetPropertyIPAddresses(value []string) (err error) { 
    return instance.SetProperty("IPAddresses", (value))
}

// GetIPAddresses gets the value of IPAddresses for the instance
func (instance *MSFT_NetEventPacketCaptureProvider) GetPropertyIPAddresses()(value []string, err error) { 
    retValue, err := instance.GetProperty("IPAddresses")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(string); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, string(valuetmp))
    }

    return
}

// SetIPProtocols sets the value of IPProtocols for the instance
func (instance *MSFT_NetEventPacketCaptureProvider) SetPropertyIPProtocols(value []uint8) (err error) { 
    return instance.SetProperty("IPProtocols", (value))
}

// GetIPProtocols gets the value of IPProtocols for the instance
func (instance *MSFT_NetEventPacketCaptureProvider) GetPropertyIPProtocols()(value []uint8, err error) { 
    retValue, err := instance.GetProperty("IPProtocols")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(uint8); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, uint8(valuetmp))
    }

    return
}

// SetLinkLayerAddress sets the value of LinkLayerAddress for the instance
func (instance *MSFT_NetEventPacketCaptureProvider) SetPropertyLinkLayerAddress(value []string) (err error) { 
    return instance.SetProperty("LinkLayerAddress", (value))
}

// GetLinkLayerAddress gets the value of LinkLayerAddress for the instance
func (instance *MSFT_NetEventPacketCaptureProvider) GetPropertyLinkLayerAddress()(value []string, err error) { 
    retValue, err := instance.GetProperty("LinkLayerAddress")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(string); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, string(valuetmp))
    }

    return
}

// SetMultiLayer sets the value of MultiLayer for the instance
func (instance *MSFT_NetEventPacketCaptureProvider) SetPropertyMultiLayer(value bool) (err error) { 
    return instance.SetProperty("MultiLayer", (value))
}

// GetMultiLayer gets the value of MultiLayer for the instance
func (instance *MSFT_NetEventPacketCaptureProvider) GetPropertyMultiLayer()(value bool, err error) { 
    retValue, err := instance.GetProperty("MultiLayer")
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

// SetTruncationLength sets the value of TruncationLength for the instance
func (instance *MSFT_NetEventPacketCaptureProvider) SetPropertyTruncationLength(value uint16) (err error) { 
    return instance.SetProperty("TruncationLength", (value))
}

// GetTruncationLength gets the value of TruncationLength for the instance
func (instance *MSFT_NetEventPacketCaptureProvider) GetPropertyTruncationLength()(value uint16, err error) { 
    retValue, err := instance.GetProperty("TruncationLength")
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

// SetVmCaptureDirection sets the value of VmCaptureDirection for the instance
func (instance *MSFT_NetEventPacketCaptureProvider) SetPropertyVmCaptureDirection(value uint8) (err error) { 
    return instance.SetProperty("VmCaptureDirection", (value))
}

// GetVmCaptureDirection gets the value of VmCaptureDirection for the instance
func (instance *MSFT_NetEventPacketCaptureProvider) GetPropertyVmCaptureDirection()(value uint8, err error) { 
    retValue, err := instance.GetProperty("VmCaptureDirection")
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

