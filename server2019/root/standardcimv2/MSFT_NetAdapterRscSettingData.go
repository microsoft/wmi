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

// MSFT_NetAdapterRscSettingData struct
type MSFT_NetAdapterRscSettingData struct { 
	*MSFT_NetAdapterSettingData

	// 
	IPv4Enabled bool

	// 
	IPv4FailureReason uint32

	// 
	IPv4OperationalState bool

	// 
	IPv6Enabled bool

	// 
	IPv6FailureReason uint32

	// 
	IPv6OperationalState bool

	// 
	RscHardwareCapabilities MSFT_NetAdapterRscCapabilities
}

	func NewMSFT_NetAdapterRscSettingDataEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterRscSettingData, err error) {tmp, err := NewMSFT_NetAdapterSettingDataEx1(instance)
		
	if err != nil { return }
	newInstance = &MSFT_NetAdapterRscSettingData {
	MSFT_NetAdapterSettingData: tmp,
	}
	return
	}
	

	func NewMSFT_NetAdapterRscSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_NetAdapterRscSettingData, err error) {tmp, err := NewMSFT_NetAdapterSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_NetAdapterRscSettingData {
	MSFT_NetAdapterSettingData: tmp,
	}
	return
	}
	

// SetIPv4Enabled sets the value of IPv4Enabled for the instance
func (instance *MSFT_NetAdapterRscSettingData) SetPropertyIPv4Enabled(value bool) (err error) { 
    return instance.SetProperty("IPv4Enabled", (value))
}

// GetIPv4Enabled gets the value of IPv4Enabled for the instance
func (instance *MSFT_NetAdapterRscSettingData) GetPropertyIPv4Enabled()(value bool, err error) { 
    retValue, err := instance.GetProperty("IPv4Enabled")
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

// SetIPv4FailureReason sets the value of IPv4FailureReason for the instance
func (instance *MSFT_NetAdapterRscSettingData) SetPropertyIPv4FailureReason(value uint32) (err error) { 
    return instance.SetProperty("IPv4FailureReason", (value))
}

// GetIPv4FailureReason gets the value of IPv4FailureReason for the instance
func (instance *MSFT_NetAdapterRscSettingData) GetPropertyIPv4FailureReason()(value uint32, err error) { 
    retValue, err := instance.GetProperty("IPv4FailureReason")
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

// SetIPv4OperationalState sets the value of IPv4OperationalState for the instance
func (instance *MSFT_NetAdapterRscSettingData) SetPropertyIPv4OperationalState(value bool) (err error) { 
    return instance.SetProperty("IPv4OperationalState", (value))
}

// GetIPv4OperationalState gets the value of IPv4OperationalState for the instance
func (instance *MSFT_NetAdapterRscSettingData) GetPropertyIPv4OperationalState()(value bool, err error) { 
    retValue, err := instance.GetProperty("IPv4OperationalState")
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

// SetIPv6Enabled sets the value of IPv6Enabled for the instance
func (instance *MSFT_NetAdapterRscSettingData) SetPropertyIPv6Enabled(value bool) (err error) { 
    return instance.SetProperty("IPv6Enabled", (value))
}

// GetIPv6Enabled gets the value of IPv6Enabled for the instance
func (instance *MSFT_NetAdapterRscSettingData) GetPropertyIPv6Enabled()(value bool, err error) { 
    retValue, err := instance.GetProperty("IPv6Enabled")
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

// SetIPv6FailureReason sets the value of IPv6FailureReason for the instance
func (instance *MSFT_NetAdapterRscSettingData) SetPropertyIPv6FailureReason(value uint32) (err error) { 
    return instance.SetProperty("IPv6FailureReason", (value))
}

// GetIPv6FailureReason gets the value of IPv6FailureReason for the instance
func (instance *MSFT_NetAdapterRscSettingData) GetPropertyIPv6FailureReason()(value uint32, err error) { 
    retValue, err := instance.GetProperty("IPv6FailureReason")
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

// SetIPv6OperationalState sets the value of IPv6OperationalState for the instance
func (instance *MSFT_NetAdapterRscSettingData) SetPropertyIPv6OperationalState(value bool) (err error) { 
    return instance.SetProperty("IPv6OperationalState", (value))
}

// GetIPv6OperationalState gets the value of IPv6OperationalState for the instance
func (instance *MSFT_NetAdapterRscSettingData) GetPropertyIPv6OperationalState()(value bool, err error) { 
    retValue, err := instance.GetProperty("IPv6OperationalState")
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

// SetRscHardwareCapabilities sets the value of RscHardwareCapabilities for the instance
func (instance *MSFT_NetAdapterRscSettingData) SetPropertyRscHardwareCapabilities(value MSFT_NetAdapterRscCapabilities) (err error) { 
    return instance.SetProperty("RscHardwareCapabilities", (value))
}

// GetRscHardwareCapabilities gets the value of RscHardwareCapabilities for the instance
func (instance *MSFT_NetAdapterRscSettingData) GetPropertyRscHardwareCapabilities()(value MSFT_NetAdapterRscCapabilities, err error) { 
    retValue, err := instance.GetProperty("RscHardwareCapabilities")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(MSFT_NetAdapterRscCapabilities); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " MSFT_NetAdapterRscCapabilities is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = MSFT_NetAdapterRscCapabilities(valuetmp)

    return
}

// 

// <param name="IPv4" type="bool "></param>
// <param name="IPv6" type="bool "></param>

// <param name="cmdletOutput" type="MSFT_NetAdapterRscSettingData "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapterRscSettingData) Enable( /* IN */ IPv4 bool,
 /* IN */ IPv6 bool,
 /* OUT */ cmdletOutput MSFT_NetAdapterRscSettingData) (result uint32, err error) {retVal, err := instance.InvokeMethod("Enable" , IPv4, IPv6)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="IPv4" type="bool "></param>
// <param name="IPv6" type="bool "></param>

// <param name="cmdletOutput" type="MSFT_NetAdapterRscSettingData "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapterRscSettingData) Disable( /* IN */ IPv4 bool,
 /* IN */ IPv6 bool,
 /* OUT */ cmdletOutput MSFT_NetAdapterRscSettingData) (result uint32, err error) {retVal, err := instance.InvokeMethod("Disable" , IPv4, IPv6)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


