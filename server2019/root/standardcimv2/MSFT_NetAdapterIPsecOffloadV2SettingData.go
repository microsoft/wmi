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

// MSFT_NetAdapterIPsecOffloadV2SettingData struct
type MSFT_NetAdapterIPsecOffloadV2SettingData struct { 
	*MSFT_NetAdapterSettingData

	// 
	AhEnabled bool

	// 
	AhEspCombinedEnabled bool

	// 
	AhEspCombinedSupported bool

	// 
	AhSupported bool

	// 
	AuthenticationAlgorithmsEnabled uint32

	// 
	AuthenticationAlgorithmsSupported uint32

	// 
	Enabled bool

	// 
	EncryptionAlgorithmsEnabled uint32

	// 
	EncryptionAlgorithmsSupported uint32

	// 
	EspEnabled bool

	// 
	EspSupported bool

	// 
	IPv4OptionsEnabled bool

	// 
	IPv4OptionsSupported bool

	// 
	IPv6Enabled bool

	// 
	IPv6NonIPsecExtensionHeadersEnabled bool

	// 
	IPv6NonIPsecExtensionHeadersSupported bool

	// 
	IPv6Supported bool

	// 
	LsoEnabled bool

	// 
	LsoSupported bool

	// 
	SaOffloadCapacityEnabled uint32

	// 
	SaOffloadCapacitySupported uint32

	// 
	TransportEnabled bool

	// 
	TransportSupported bool

	// 
	TunnelEnabled bool

	// 
	TunnelSupported bool

	// 
	UdpEspEnabled uint32

	// 
	UdpEspSupported uint32
}

	func NewMSFT_NetAdapterIPsecOffloadV2SettingDataEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterIPsecOffloadV2SettingData, err error) {tmp, err := NewMSFT_NetAdapterSettingDataEx1(instance)
		
	if err != nil { return }
	newInstance = &MSFT_NetAdapterIPsecOffloadV2SettingData {
	MSFT_NetAdapterSettingData: tmp,
	}
	return
	}
	

	func NewMSFT_NetAdapterIPsecOffloadV2SettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_NetAdapterIPsecOffloadV2SettingData, err error) {tmp, err := NewMSFT_NetAdapterSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_NetAdapterIPsecOffloadV2SettingData {
	MSFT_NetAdapterSettingData: tmp,
	}
	return
	}
	

// SetAhEnabled sets the value of AhEnabled for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) SetPropertyAhEnabled(value bool) (err error) { 
    return instance.SetProperty("AhEnabled", (value))
}

// GetAhEnabled gets the value of AhEnabled for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) GetPropertyAhEnabled()(value bool, err error) { 
    retValue, err := instance.GetProperty("AhEnabled")
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

// SetAhEspCombinedEnabled sets the value of AhEspCombinedEnabled for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) SetPropertyAhEspCombinedEnabled(value bool) (err error) { 
    return instance.SetProperty("AhEspCombinedEnabled", (value))
}

// GetAhEspCombinedEnabled gets the value of AhEspCombinedEnabled for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) GetPropertyAhEspCombinedEnabled()(value bool, err error) { 
    retValue, err := instance.GetProperty("AhEspCombinedEnabled")
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

// SetAhEspCombinedSupported sets the value of AhEspCombinedSupported for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) SetPropertyAhEspCombinedSupported(value bool) (err error) { 
    return instance.SetProperty("AhEspCombinedSupported", (value))
}

// GetAhEspCombinedSupported gets the value of AhEspCombinedSupported for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) GetPropertyAhEspCombinedSupported()(value bool, err error) { 
    retValue, err := instance.GetProperty("AhEspCombinedSupported")
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

// SetAhSupported sets the value of AhSupported for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) SetPropertyAhSupported(value bool) (err error) { 
    return instance.SetProperty("AhSupported", (value))
}

// GetAhSupported gets the value of AhSupported for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) GetPropertyAhSupported()(value bool, err error) { 
    retValue, err := instance.GetProperty("AhSupported")
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

// SetAuthenticationAlgorithmsEnabled sets the value of AuthenticationAlgorithmsEnabled for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) SetPropertyAuthenticationAlgorithmsEnabled(value uint32) (err error) { 
    return instance.SetProperty("AuthenticationAlgorithmsEnabled", (value))
}

// GetAuthenticationAlgorithmsEnabled gets the value of AuthenticationAlgorithmsEnabled for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) GetPropertyAuthenticationAlgorithmsEnabled()(value uint32, err error) { 
    retValue, err := instance.GetProperty("AuthenticationAlgorithmsEnabled")
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

// SetAuthenticationAlgorithmsSupported sets the value of AuthenticationAlgorithmsSupported for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) SetPropertyAuthenticationAlgorithmsSupported(value uint32) (err error) { 
    return instance.SetProperty("AuthenticationAlgorithmsSupported", (value))
}

// GetAuthenticationAlgorithmsSupported gets the value of AuthenticationAlgorithmsSupported for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) GetPropertyAuthenticationAlgorithmsSupported()(value uint32, err error) { 
    retValue, err := instance.GetProperty("AuthenticationAlgorithmsSupported")
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
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) SetPropertyEnabled(value bool) (err error) { 
    return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) GetPropertyEnabled()(value bool, err error) { 
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

// SetEncryptionAlgorithmsEnabled sets the value of EncryptionAlgorithmsEnabled for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) SetPropertyEncryptionAlgorithmsEnabled(value uint32) (err error) { 
    return instance.SetProperty("EncryptionAlgorithmsEnabled", (value))
}

// GetEncryptionAlgorithmsEnabled gets the value of EncryptionAlgorithmsEnabled for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) GetPropertyEncryptionAlgorithmsEnabled()(value uint32, err error) { 
    retValue, err := instance.GetProperty("EncryptionAlgorithmsEnabled")
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

// SetEncryptionAlgorithmsSupported sets the value of EncryptionAlgorithmsSupported for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) SetPropertyEncryptionAlgorithmsSupported(value uint32) (err error) { 
    return instance.SetProperty("EncryptionAlgorithmsSupported", (value))
}

// GetEncryptionAlgorithmsSupported gets the value of EncryptionAlgorithmsSupported for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) GetPropertyEncryptionAlgorithmsSupported()(value uint32, err error) { 
    retValue, err := instance.GetProperty("EncryptionAlgorithmsSupported")
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

// SetEspEnabled sets the value of EspEnabled for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) SetPropertyEspEnabled(value bool) (err error) { 
    return instance.SetProperty("EspEnabled", (value))
}

// GetEspEnabled gets the value of EspEnabled for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) GetPropertyEspEnabled()(value bool, err error) { 
    retValue, err := instance.GetProperty("EspEnabled")
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

// SetEspSupported sets the value of EspSupported for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) SetPropertyEspSupported(value bool) (err error) { 
    return instance.SetProperty("EspSupported", (value))
}

// GetEspSupported gets the value of EspSupported for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) GetPropertyEspSupported()(value bool, err error) { 
    retValue, err := instance.GetProperty("EspSupported")
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

// SetIPv4OptionsEnabled sets the value of IPv4OptionsEnabled for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) SetPropertyIPv4OptionsEnabled(value bool) (err error) { 
    return instance.SetProperty("IPv4OptionsEnabled", (value))
}

// GetIPv4OptionsEnabled gets the value of IPv4OptionsEnabled for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) GetPropertyIPv4OptionsEnabled()(value bool, err error) { 
    retValue, err := instance.GetProperty("IPv4OptionsEnabled")
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

// SetIPv4OptionsSupported sets the value of IPv4OptionsSupported for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) SetPropertyIPv4OptionsSupported(value bool) (err error) { 
    return instance.SetProperty("IPv4OptionsSupported", (value))
}

// GetIPv4OptionsSupported gets the value of IPv4OptionsSupported for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) GetPropertyIPv4OptionsSupported()(value bool, err error) { 
    retValue, err := instance.GetProperty("IPv4OptionsSupported")
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
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) SetPropertyIPv6Enabled(value bool) (err error) { 
    return instance.SetProperty("IPv6Enabled", (value))
}

// GetIPv6Enabled gets the value of IPv6Enabled for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) GetPropertyIPv6Enabled()(value bool, err error) { 
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

// SetIPv6NonIPsecExtensionHeadersEnabled sets the value of IPv6NonIPsecExtensionHeadersEnabled for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) SetPropertyIPv6NonIPsecExtensionHeadersEnabled(value bool) (err error) { 
    return instance.SetProperty("IPv6NonIPsecExtensionHeadersEnabled", (value))
}

// GetIPv6NonIPsecExtensionHeadersEnabled gets the value of IPv6NonIPsecExtensionHeadersEnabled for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) GetPropertyIPv6NonIPsecExtensionHeadersEnabled()(value bool, err error) { 
    retValue, err := instance.GetProperty("IPv6NonIPsecExtensionHeadersEnabled")
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

// SetIPv6NonIPsecExtensionHeadersSupported sets the value of IPv6NonIPsecExtensionHeadersSupported for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) SetPropertyIPv6NonIPsecExtensionHeadersSupported(value bool) (err error) { 
    return instance.SetProperty("IPv6NonIPsecExtensionHeadersSupported", (value))
}

// GetIPv6NonIPsecExtensionHeadersSupported gets the value of IPv6NonIPsecExtensionHeadersSupported for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) GetPropertyIPv6NonIPsecExtensionHeadersSupported()(value bool, err error) { 
    retValue, err := instance.GetProperty("IPv6NonIPsecExtensionHeadersSupported")
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

// SetIPv6Supported sets the value of IPv6Supported for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) SetPropertyIPv6Supported(value bool) (err error) { 
    return instance.SetProperty("IPv6Supported", (value))
}

// GetIPv6Supported gets the value of IPv6Supported for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) GetPropertyIPv6Supported()(value bool, err error) { 
    retValue, err := instance.GetProperty("IPv6Supported")
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

// SetLsoEnabled sets the value of LsoEnabled for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) SetPropertyLsoEnabled(value bool) (err error) { 
    return instance.SetProperty("LsoEnabled", (value))
}

// GetLsoEnabled gets the value of LsoEnabled for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) GetPropertyLsoEnabled()(value bool, err error) { 
    retValue, err := instance.GetProperty("LsoEnabled")
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

// SetLsoSupported sets the value of LsoSupported for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) SetPropertyLsoSupported(value bool) (err error) { 
    return instance.SetProperty("LsoSupported", (value))
}

// GetLsoSupported gets the value of LsoSupported for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) GetPropertyLsoSupported()(value bool, err error) { 
    retValue, err := instance.GetProperty("LsoSupported")
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

// SetSaOffloadCapacityEnabled sets the value of SaOffloadCapacityEnabled for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) SetPropertySaOffloadCapacityEnabled(value uint32) (err error) { 
    return instance.SetProperty("SaOffloadCapacityEnabled", (value))
}

// GetSaOffloadCapacityEnabled gets the value of SaOffloadCapacityEnabled for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) GetPropertySaOffloadCapacityEnabled()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SaOffloadCapacityEnabled")
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

// SetSaOffloadCapacitySupported sets the value of SaOffloadCapacitySupported for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) SetPropertySaOffloadCapacitySupported(value uint32) (err error) { 
    return instance.SetProperty("SaOffloadCapacitySupported", (value))
}

// GetSaOffloadCapacitySupported gets the value of SaOffloadCapacitySupported for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) GetPropertySaOffloadCapacitySupported()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SaOffloadCapacitySupported")
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

// SetTransportEnabled sets the value of TransportEnabled for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) SetPropertyTransportEnabled(value bool) (err error) { 
    return instance.SetProperty("TransportEnabled", (value))
}

// GetTransportEnabled gets the value of TransportEnabled for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) GetPropertyTransportEnabled()(value bool, err error) { 
    retValue, err := instance.GetProperty("TransportEnabled")
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

// SetTransportSupported sets the value of TransportSupported for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) SetPropertyTransportSupported(value bool) (err error) { 
    return instance.SetProperty("TransportSupported", (value))
}

// GetTransportSupported gets the value of TransportSupported for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) GetPropertyTransportSupported()(value bool, err error) { 
    retValue, err := instance.GetProperty("TransportSupported")
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

// SetTunnelEnabled sets the value of TunnelEnabled for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) SetPropertyTunnelEnabled(value bool) (err error) { 
    return instance.SetProperty("TunnelEnabled", (value))
}

// GetTunnelEnabled gets the value of TunnelEnabled for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) GetPropertyTunnelEnabled()(value bool, err error) { 
    retValue, err := instance.GetProperty("TunnelEnabled")
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

// SetTunnelSupported sets the value of TunnelSupported for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) SetPropertyTunnelSupported(value bool) (err error) { 
    return instance.SetProperty("TunnelSupported", (value))
}

// GetTunnelSupported gets the value of TunnelSupported for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) GetPropertyTunnelSupported()(value bool, err error) { 
    retValue, err := instance.GetProperty("TunnelSupported")
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

// SetUdpEspEnabled sets the value of UdpEspEnabled for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) SetPropertyUdpEspEnabled(value uint32) (err error) { 
    return instance.SetProperty("UdpEspEnabled", (value))
}

// GetUdpEspEnabled gets the value of UdpEspEnabled for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) GetPropertyUdpEspEnabled()(value uint32, err error) { 
    retValue, err := instance.GetProperty("UdpEspEnabled")
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

// SetUdpEspSupported sets the value of UdpEspSupported for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) SetPropertyUdpEspSupported(value uint32) (err error) { 
    return instance.SetProperty("UdpEspSupported", (value))
}

// GetUdpEspSupported gets the value of UdpEspSupported for the instance
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) GetPropertyUdpEspSupported()(value uint32, err error) { 
    retValue, err := instance.GetProperty("UdpEspSupported")
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

// 

// <param name="Enabled" type="bool "></param>
// <param name="NoRestart" type="bool "></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="MSFT_NetAdapterIPsecOffloadV2SettingData "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) Set( /* IN */ Enabled bool,
 /* IN */ NoRestart bool,
 /* IN */ PassThru bool,
 /* OUT */ cmdletOutput MSFT_NetAdapterIPsecOffloadV2SettingData) (result uint32, err error) {retVal, err := instance.InvokeMethod("Set" , Enabled, NoRestart, PassThru)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="NoRestart" type="bool "></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="MSFT_NetAdapterIPsecOffloadV2SettingData "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) Enable( /* IN */ NoRestart bool,
 /* IN */ PassThru bool,
 /* OUT */ cmdletOutput MSFT_NetAdapterIPsecOffloadV2SettingData) (result uint32, err error) {retVal, err := instance.InvokeMethod("Enable" , NoRestart, PassThru)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="NoRestart" type="bool "></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="MSFT_NetAdapterIPsecOffloadV2SettingData "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapterIPsecOffloadV2SettingData) Disable( /* IN */ NoRestart bool,
 /* IN */ PassThru bool,
 /* OUT */ cmdletOutput MSFT_NetAdapterIPsecOffloadV2SettingData) (result uint32, err error) {retVal, err := instance.InvokeMethod("Disable" , NoRestart, PassThru)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


