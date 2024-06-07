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
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// VpnServerIPsecConfiguration struct
type VpnServerIPsecConfiguration struct { 
	*cim.WmiInstance

	// 
	GrePorts uint32

	// 
	IdleDisconnect uint32

	// 
	Ikev2Ports uint32

	// 
	L2tpPorts uint32

	// 
	MMSALifeTime uint32

	// 
	SADataSizeForRenegotiation uint32

	// 
	SALifeTime uint32

	// 
	SstpPorts uint32

	// 
	TunnelType uint32
}

	func NewVpnServerIPsecConfigurationEx1(instance *cim.WmiInstance) (newInstance *VpnServerIPsecConfiguration, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &VpnServerIPsecConfiguration {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewVpnServerIPsecConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *VpnServerIPsecConfiguration, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &VpnServerIPsecConfiguration {
	WmiInstance: tmp,
	}
	return
	}
	

// SetGrePorts sets the value of GrePorts for the instance
func (instance *VpnServerIPsecConfiguration) SetPropertyGrePorts(value uint32) (err error) { 
    return instance.SetProperty("GrePorts", (value))
}

// GetGrePorts gets the value of GrePorts for the instance
func (instance *VpnServerIPsecConfiguration) GetPropertyGrePorts()(value uint32, err error) { 
    retValue, err := instance.GetProperty("GrePorts")
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

// SetIdleDisconnect sets the value of IdleDisconnect for the instance
func (instance *VpnServerIPsecConfiguration) SetPropertyIdleDisconnect(value uint32) (err error) { 
    return instance.SetProperty("IdleDisconnect", (value))
}

// GetIdleDisconnect gets the value of IdleDisconnect for the instance
func (instance *VpnServerIPsecConfiguration) GetPropertyIdleDisconnect()(value uint32, err error) { 
    retValue, err := instance.GetProperty("IdleDisconnect")
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

// SetIkev2Ports sets the value of Ikev2Ports for the instance
func (instance *VpnServerIPsecConfiguration) SetPropertyIkev2Ports(value uint32) (err error) { 
    return instance.SetProperty("Ikev2Ports", (value))
}

// GetIkev2Ports gets the value of Ikev2Ports for the instance
func (instance *VpnServerIPsecConfiguration) GetPropertyIkev2Ports()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Ikev2Ports")
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

// SetL2tpPorts sets the value of L2tpPorts for the instance
func (instance *VpnServerIPsecConfiguration) SetPropertyL2tpPorts(value uint32) (err error) { 
    return instance.SetProperty("L2tpPorts", (value))
}

// GetL2tpPorts gets the value of L2tpPorts for the instance
func (instance *VpnServerIPsecConfiguration) GetPropertyL2tpPorts()(value uint32, err error) { 
    retValue, err := instance.GetProperty("L2tpPorts")
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

// SetMMSALifeTime sets the value of MMSALifeTime for the instance
func (instance *VpnServerIPsecConfiguration) SetPropertyMMSALifeTime(value uint32) (err error) { 
    return instance.SetProperty("MMSALifeTime", (value))
}

// GetMMSALifeTime gets the value of MMSALifeTime for the instance
func (instance *VpnServerIPsecConfiguration) GetPropertyMMSALifeTime()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MMSALifeTime")
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

// SetSADataSizeForRenegotiation sets the value of SADataSizeForRenegotiation for the instance
func (instance *VpnServerIPsecConfiguration) SetPropertySADataSizeForRenegotiation(value uint32) (err error) { 
    return instance.SetProperty("SADataSizeForRenegotiation", (value))
}

// GetSADataSizeForRenegotiation gets the value of SADataSizeForRenegotiation for the instance
func (instance *VpnServerIPsecConfiguration) GetPropertySADataSizeForRenegotiation()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SADataSizeForRenegotiation")
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

// SetSALifeTime sets the value of SALifeTime for the instance
func (instance *VpnServerIPsecConfiguration) SetPropertySALifeTime(value uint32) (err error) { 
    return instance.SetProperty("SALifeTime", (value))
}

// GetSALifeTime gets the value of SALifeTime for the instance
func (instance *VpnServerIPsecConfiguration) GetPropertySALifeTime()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SALifeTime")
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

// SetSstpPorts sets the value of SstpPorts for the instance
func (instance *VpnServerIPsecConfiguration) SetPropertySstpPorts(value uint32) (err error) { 
    return instance.SetProperty("SstpPorts", (value))
}

// GetSstpPorts gets the value of SstpPorts for the instance
func (instance *VpnServerIPsecConfiguration) GetPropertySstpPorts()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SstpPorts")
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

// SetTunnelType sets the value of TunnelType for the instance
func (instance *VpnServerIPsecConfiguration) SetPropertyTunnelType(value uint32) (err error) { 
    return instance.SetProperty("TunnelType", (value))
}

// GetTunnelType gets the value of TunnelType for the instance
func (instance *VpnServerIPsecConfiguration) GetPropertyTunnelType()(value uint32, err error) { 
    retValue, err := instance.GetProperty("TunnelType")
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

