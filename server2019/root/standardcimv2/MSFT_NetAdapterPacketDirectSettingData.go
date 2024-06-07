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

// MSFT_NetAdapterPacketDirectSettingData struct
type MSFT_NetAdapterPacketDirectSettingData struct { 
	*MSFT_NetAdapterSettingData

	// 
	Capabilities MSFT_NetAdapter_PacketDirectCapabilities

	// 
	DiagnosticCode uint32

	// 
	DmaAddressWidth uint8

	// 
	DomainId uint32

	// 
	Enabled bool

	// 
	Operational bool
}

	func NewMSFT_NetAdapterPacketDirectSettingDataEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterPacketDirectSettingData, err error) {tmp, err := NewMSFT_NetAdapterSettingDataEx1(instance)
		
	if err != nil { return }
	newInstance = &MSFT_NetAdapterPacketDirectSettingData {
	MSFT_NetAdapterSettingData: tmp,
	}
	return
	}
	

	func NewMSFT_NetAdapterPacketDirectSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_NetAdapterPacketDirectSettingData, err error) {tmp, err := NewMSFT_NetAdapterSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_NetAdapterPacketDirectSettingData {
	MSFT_NetAdapterSettingData: tmp,
	}
	return
	}
	

// SetCapabilities sets the value of Capabilities for the instance
func (instance *MSFT_NetAdapterPacketDirectSettingData) SetPropertyCapabilities(value MSFT_NetAdapter_PacketDirectCapabilities) (err error) { 
    return instance.SetProperty("Capabilities", (value))
}

// GetCapabilities gets the value of Capabilities for the instance
func (instance *MSFT_NetAdapterPacketDirectSettingData) GetPropertyCapabilities()(value MSFT_NetAdapter_PacketDirectCapabilities, err error) { 
    retValue, err := instance.GetProperty("Capabilities")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(MSFT_NetAdapter_PacketDirectCapabilities); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " MSFT_NetAdapter_PacketDirectCapabilities is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = MSFT_NetAdapter_PacketDirectCapabilities(valuetmp)

    return
}

// SetDiagnosticCode sets the value of DiagnosticCode for the instance
func (instance *MSFT_NetAdapterPacketDirectSettingData) SetPropertyDiagnosticCode(value uint32) (err error) { 
    return instance.SetProperty("DiagnosticCode", (value))
}

// GetDiagnosticCode gets the value of DiagnosticCode for the instance
func (instance *MSFT_NetAdapterPacketDirectSettingData) GetPropertyDiagnosticCode()(value uint32, err error) { 
    retValue, err := instance.GetProperty("DiagnosticCode")
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

// SetDmaAddressWidth sets the value of DmaAddressWidth for the instance
func (instance *MSFT_NetAdapterPacketDirectSettingData) SetPropertyDmaAddressWidth(value uint8) (err error) { 
    return instance.SetProperty("DmaAddressWidth", (value))
}

// GetDmaAddressWidth gets the value of DmaAddressWidth for the instance
func (instance *MSFT_NetAdapterPacketDirectSettingData) GetPropertyDmaAddressWidth()(value uint8, err error) { 
    retValue, err := instance.GetProperty("DmaAddressWidth")
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

// SetDomainId sets the value of DomainId for the instance
func (instance *MSFT_NetAdapterPacketDirectSettingData) SetPropertyDomainId(value uint32) (err error) { 
    return instance.SetProperty("DomainId", (value))
}

// GetDomainId gets the value of DomainId for the instance
func (instance *MSFT_NetAdapterPacketDirectSettingData) GetPropertyDomainId()(value uint32, err error) { 
    retValue, err := instance.GetProperty("DomainId")
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

// SetEnabled sets the value of Enabled for the instance
func (instance *MSFT_NetAdapterPacketDirectSettingData) SetPropertyEnabled(value bool) (err error) { 
    return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *MSFT_NetAdapterPacketDirectSettingData) GetPropertyEnabled()(value bool, err error) { 
    retValue, err := instance.GetProperty("Enabled")
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

// SetOperational sets the value of Operational for the instance
func (instance *MSFT_NetAdapterPacketDirectSettingData) SetPropertyOperational(value bool) (err error) { 
    return instance.SetProperty("Operational", (value))
}

// GetOperational gets the value of Operational for the instance
func (instance *MSFT_NetAdapterPacketDirectSettingData) GetPropertyOperational()(value bool, err error) { 
    retValue, err := instance.GetProperty("Operational")
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

// 

// <param name="CmdletOutput" type="MSFT_NetAdapterPacketDirectSettingData "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapterPacketDirectSettingData) Enable( /* OUT */ CmdletOutput MSFT_NetAdapterPacketDirectSettingData) (result uint32, err error) {retVal, err := instance.InvokeMethod("Enable" )
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="CmdletOutput" type="MSFT_NetAdapterPacketDirectSettingData "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapterPacketDirectSettingData) Disable( /* OUT */ CmdletOutput MSFT_NetAdapterPacketDirectSettingData) (result uint32, err error) {retVal, err := instance.InvokeMethod("Disable" )
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


