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

// VpnS2SCustomInterface struct
type VpnS2SCustomInterface struct { 
	*VpnS2SInterface

	// 
	AuthenticationTransformConstants uint32

	// 
	CipherTransformConstants uint32

	// 
	CustomPolicy bool

	// 
	DHGroup uint32

	// 
	EncryptionMethod uint32

	// 
	IntegrityCheckMethod uint32

	// 
	PfsGroup uint32
}

	func NewVpnS2SCustomInterfaceEx1(instance *cim.WmiInstance) (newInstance *VpnS2SCustomInterface, err error) {tmp, err := NewVpnS2SInterfaceEx1(instance)
		
	if err != nil { return }
	newInstance = &VpnS2SCustomInterface {
	VpnS2SInterface: tmp,
	}
	return
	}
	

	func NewVpnS2SCustomInterfaceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *VpnS2SCustomInterface, err error) {tmp, err := NewVpnS2SInterfaceEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &VpnS2SCustomInterface {
	VpnS2SInterface: tmp,
	}
	return
	}
	

// SetAuthenticationTransformConstants sets the value of AuthenticationTransformConstants for the instance
func (instance *VpnS2SCustomInterface) SetPropertyAuthenticationTransformConstants(value uint32) (err error) { 
    return instance.SetProperty("AuthenticationTransformConstants", (value))
}

// GetAuthenticationTransformConstants gets the value of AuthenticationTransformConstants for the instance
func (instance *VpnS2SCustomInterface) GetPropertyAuthenticationTransformConstants()(value uint32, err error) { 
    retValue, err := instance.GetProperty("AuthenticationTransformConstants")
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

// SetCipherTransformConstants sets the value of CipherTransformConstants for the instance
func (instance *VpnS2SCustomInterface) SetPropertyCipherTransformConstants(value uint32) (err error) { 
    return instance.SetProperty("CipherTransformConstants", (value))
}

// GetCipherTransformConstants gets the value of CipherTransformConstants for the instance
func (instance *VpnS2SCustomInterface) GetPropertyCipherTransformConstants()(value uint32, err error) { 
    retValue, err := instance.GetProperty("CipherTransformConstants")
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
func (instance *VpnS2SCustomInterface) SetPropertyCustomPolicy(value bool) (err error) { 
    return instance.SetProperty("CustomPolicy", (value))
}

// GetCustomPolicy gets the value of CustomPolicy for the instance
func (instance *VpnS2SCustomInterface) GetPropertyCustomPolicy()(value bool, err error) { 
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
func (instance *VpnS2SCustomInterface) SetPropertyDHGroup(value uint32) (err error) { 
    return instance.SetProperty("DHGroup", (value))
}

// GetDHGroup gets the value of DHGroup for the instance
func (instance *VpnS2SCustomInterface) GetPropertyDHGroup()(value uint32, err error) { 
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
func (instance *VpnS2SCustomInterface) SetPropertyEncryptionMethod(value uint32) (err error) { 
    return instance.SetProperty("EncryptionMethod", (value))
}

// GetEncryptionMethod gets the value of EncryptionMethod for the instance
func (instance *VpnS2SCustomInterface) GetPropertyEncryptionMethod()(value uint32, err error) { 
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

// SetIntegrityCheckMethod sets the value of IntegrityCheckMethod for the instance
func (instance *VpnS2SCustomInterface) SetPropertyIntegrityCheckMethod(value uint32) (err error) { 
    return instance.SetProperty("IntegrityCheckMethod", (value))
}

// GetIntegrityCheckMethod gets the value of IntegrityCheckMethod for the instance
func (instance *VpnS2SCustomInterface) GetPropertyIntegrityCheckMethod()(value uint32, err error) { 
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

// SetPfsGroup sets the value of PfsGroup for the instance
func (instance *VpnS2SCustomInterface) SetPropertyPfsGroup(value uint32) (err error) { 
    return instance.SetProperty("PfsGroup", (value))
}

// GetPfsGroup gets the value of PfsGroup for the instance
func (instance *VpnS2SCustomInterface) GetPropertyPfsGroup()(value uint32, err error) { 
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

