// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess
//////////////////////////////////////////////
package remoteaccess
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// VpnRoutingDomainCustomConfig struct
type VpnRoutingDomainCustomConfig struct { 
	*VpnRoutingDomainConfig

	// 
	AuthenticationTransformConstant uint32

	// 
	CipherTransformConstant uint32

	// 
	CustomPolicy bool

	// 
	DHGroup uint32

	// 
	EncryptionMethod uint32

	// 
	IdleDisconnectSec uint32

	// 
	IntegrityCheckMethod uint32

	// 
	MMSaLifeTimeSec uint32

	// 
	PfsGroup uint32

	// 
	SaLifeTimeSec uint32

	// 
	SaRenegotiationDataSizeKB uint32
}

	func NewVpnRoutingDomainCustomConfigEx1(instance *cim.WmiInstance) (newInstance *VpnRoutingDomainCustomConfig, err error) {tmp, err := NewVpnRoutingDomainConfigEx1(instance)
		
	if err != nil { return }
	newInstance = &VpnRoutingDomainCustomConfig {
	VpnRoutingDomainConfig: tmp,
	}
	return
	}
	

	func NewVpnRoutingDomainCustomConfigEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *VpnRoutingDomainCustomConfig, err error) {tmp, err := NewVpnRoutingDomainConfigEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &VpnRoutingDomainCustomConfig {
	VpnRoutingDomainConfig: tmp,
	}
	return
	}
	

// SetAuthenticationTransformConstant sets the value of AuthenticationTransformConstant for the instance
func (instance *VpnRoutingDomainCustomConfig) SetPropertyAuthenticationTransformConstant(value uint32) (err error) { 
    return instance.SetProperty("AuthenticationTransformConstant", (value))
}

// GetAuthenticationTransformConstant gets the value of AuthenticationTransformConstant for the instance
func (instance *VpnRoutingDomainCustomConfig) GetPropertyAuthenticationTransformConstant()(value uint32, err error) { 
    retValue, err := instance.GetProperty("AuthenticationTransformConstant")
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

// SetCipherTransformConstant sets the value of CipherTransformConstant for the instance
func (instance *VpnRoutingDomainCustomConfig) SetPropertyCipherTransformConstant(value uint32) (err error) { 
    return instance.SetProperty("CipherTransformConstant", (value))
}

// GetCipherTransformConstant gets the value of CipherTransformConstant for the instance
func (instance *VpnRoutingDomainCustomConfig) GetPropertyCipherTransformConstant()(value uint32, err error) { 
    retValue, err := instance.GetProperty("CipherTransformConstant")
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

// SetCustomPolicy sets the value of CustomPolicy for the instance
func (instance *VpnRoutingDomainCustomConfig) SetPropertyCustomPolicy(value bool) (err error) { 
    return instance.SetProperty("CustomPolicy", (value))
}

// GetCustomPolicy gets the value of CustomPolicy for the instance
func (instance *VpnRoutingDomainCustomConfig) GetPropertyCustomPolicy()(value bool, err error) { 
    retValue, err := instance.GetProperty("CustomPolicy")
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

// SetDHGroup sets the value of DHGroup for the instance
func (instance *VpnRoutingDomainCustomConfig) SetPropertyDHGroup(value uint32) (err error) { 
    return instance.SetProperty("DHGroup", (value))
}

// GetDHGroup gets the value of DHGroup for the instance
func (instance *VpnRoutingDomainCustomConfig) GetPropertyDHGroup()(value uint32, err error) { 
    retValue, err := instance.GetProperty("DHGroup")
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

// SetEncryptionMethod sets the value of EncryptionMethod for the instance
func (instance *VpnRoutingDomainCustomConfig) SetPropertyEncryptionMethod(value uint32) (err error) { 
    return instance.SetProperty("EncryptionMethod", (value))
}

// GetEncryptionMethod gets the value of EncryptionMethod for the instance
func (instance *VpnRoutingDomainCustomConfig) GetPropertyEncryptionMethod()(value uint32, err error) { 
    retValue, err := instance.GetProperty("EncryptionMethod")
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

// SetIdleDisconnectSec sets the value of IdleDisconnectSec for the instance
func (instance *VpnRoutingDomainCustomConfig) SetPropertyIdleDisconnectSec(value uint32) (err error) { 
    return instance.SetProperty("IdleDisconnectSec", (value))
}

// GetIdleDisconnectSec gets the value of IdleDisconnectSec for the instance
func (instance *VpnRoutingDomainCustomConfig) GetPropertyIdleDisconnectSec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("IdleDisconnectSec")
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

// SetIntegrityCheckMethod sets the value of IntegrityCheckMethod for the instance
func (instance *VpnRoutingDomainCustomConfig) SetPropertyIntegrityCheckMethod(value uint32) (err error) { 
    return instance.SetProperty("IntegrityCheckMethod", (value))
}

// GetIntegrityCheckMethod gets the value of IntegrityCheckMethod for the instance
func (instance *VpnRoutingDomainCustomConfig) GetPropertyIntegrityCheckMethod()(value uint32, err error) { 
    retValue, err := instance.GetProperty("IntegrityCheckMethod")
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

// SetMMSaLifeTimeSec sets the value of MMSaLifeTimeSec for the instance
func (instance *VpnRoutingDomainCustomConfig) SetPropertyMMSaLifeTimeSec(value uint32) (err error) { 
    return instance.SetProperty("MMSaLifeTimeSec", (value))
}

// GetMMSaLifeTimeSec gets the value of MMSaLifeTimeSec for the instance
func (instance *VpnRoutingDomainCustomConfig) GetPropertyMMSaLifeTimeSec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MMSaLifeTimeSec")
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

// SetPfsGroup sets the value of PfsGroup for the instance
func (instance *VpnRoutingDomainCustomConfig) SetPropertyPfsGroup(value uint32) (err error) { 
    return instance.SetProperty("PfsGroup", (value))
}

// GetPfsGroup gets the value of PfsGroup for the instance
func (instance *VpnRoutingDomainCustomConfig) GetPropertyPfsGroup()(value uint32, err error) { 
    retValue, err := instance.GetProperty("PfsGroup")
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

// SetSaLifeTimeSec sets the value of SaLifeTimeSec for the instance
func (instance *VpnRoutingDomainCustomConfig) SetPropertySaLifeTimeSec(value uint32) (err error) { 
    return instance.SetProperty("SaLifeTimeSec", (value))
}

// GetSaLifeTimeSec gets the value of SaLifeTimeSec for the instance
func (instance *VpnRoutingDomainCustomConfig) GetPropertySaLifeTimeSec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SaLifeTimeSec")
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

// SetSaRenegotiationDataSizeKB sets the value of SaRenegotiationDataSizeKB for the instance
func (instance *VpnRoutingDomainCustomConfig) SetPropertySaRenegotiationDataSizeKB(value uint32) (err error) { 
    return instance.SetProperty("SaRenegotiationDataSizeKB", (value))
}

// GetSaRenegotiationDataSizeKB gets the value of SaRenegotiationDataSizeKB for the instance
func (instance *VpnRoutingDomainCustomConfig) GetPropertySaRenegotiationDataSizeKB()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SaRenegotiationDataSizeKB")
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

