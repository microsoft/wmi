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
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSFT_NetLldpTlv struct
type MSFT_NetLldpTlv struct { 
	*cim.WmiInstance

	// 
	Data []uint8

	// 
	Oui []uint8

	// 
	OuiSubtype uint32

	// 
	TlvName string

	// 
	TlvType uint32
}

	func NewMSFT_NetLldpTlvEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetLldpTlv, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MSFT_NetLldpTlv {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMSFT_NetLldpTlvEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_NetLldpTlv, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_NetLldpTlv {
	WmiInstance: tmp,
	}
	return
	}
	

// SetData sets the value of Data for the instance
func (instance *MSFT_NetLldpTlv) SetPropertyData(value []uint8) (err error) { 
    return instance.SetProperty("Data", (value))
}

// GetData gets the value of Data for the instance
func (instance *MSFT_NetLldpTlv) GetPropertyData()(value []uint8, err error) { 
    retValue, err := instance.GetProperty("Data")
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

// SetOui sets the value of Oui for the instance
func (instance *MSFT_NetLldpTlv) SetPropertyOui(value []uint8) (err error) { 
    return instance.SetProperty("Oui", (value))
}

// GetOui gets the value of Oui for the instance
func (instance *MSFT_NetLldpTlv) GetPropertyOui()(value []uint8, err error) { 
    retValue, err := instance.GetProperty("Oui")
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

// SetOuiSubtype sets the value of OuiSubtype for the instance
func (instance *MSFT_NetLldpTlv) SetPropertyOuiSubtype(value uint32) (err error) { 
    return instance.SetProperty("OuiSubtype", (value))
}

// GetOuiSubtype gets the value of OuiSubtype for the instance
func (instance *MSFT_NetLldpTlv) GetPropertyOuiSubtype()(value uint32, err error) { 
    retValue, err := instance.GetProperty("OuiSubtype")
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

// SetTlvName sets the value of TlvName for the instance
func (instance *MSFT_NetLldpTlv) SetPropertyTlvName(value string) (err error) { 
    return instance.SetProperty("TlvName", (value))
}

// GetTlvName gets the value of TlvName for the instance
func (instance *MSFT_NetLldpTlv) GetPropertyTlvName()(value string, err error) { 
    retValue, err := instance.GetProperty("TlvName")
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

// SetTlvType sets the value of TlvType for the instance
func (instance *MSFT_NetLldpTlv) SetPropertyTlvType(value uint32) (err error) { 
    return instance.SetProperty("TlvType", (value))
}

// GetTlvType gets the value of TlvType for the instance
func (instance *MSFT_NetLldpTlv) GetPropertyTlvType()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TlvType")
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

