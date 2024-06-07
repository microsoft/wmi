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

// MSNdis_PMAdminConfigParam struct
type MSNdis_PMAdminConfigParam struct { 
	*MSNdis

	// 
	DeviceSleepOnDisconnect MSNdis_PMAdminConfigState

	// 
	Header MSNdis_ObjectHeader

	// 
	PMARPOffload MSNdis_PMAdminConfigState

	// 
	PMNDOffload MSNdis_PMAdminConfigState

	// 
	PMWiFiRekeyOffload MSNdis_PMAdminConfigState

	// 
	WakeOnMagicPacket MSNdis_PMAdminConfigState

	// 
	WakeOnPattern MSNdis_PMAdminConfigState
}

	func NewMSNdis_PMAdminConfigParamEx1(instance *cim.WmiInstance) (newInstance *MSNdis_PMAdminConfigParam, err error) {tmp, err := NewMSNdisEx1(instance)
		
	if err != nil { return }
	newInstance = &MSNdis_PMAdminConfigParam {
	MSNdis: tmp,
	}
	return
	}
	

	func NewMSNdis_PMAdminConfigParamEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSNdis_PMAdminConfigParam, err error) {tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSNdis_PMAdminConfigParam {
	MSNdis: tmp,
	}
	return
	}
	

// SetDeviceSleepOnDisconnect sets the value of DeviceSleepOnDisconnect for the instance
func (instance *MSNdis_PMAdminConfigParam) SetPropertyDeviceSleepOnDisconnect(value MSNdis_PMAdminConfigState) (err error) { 
    return instance.SetProperty("DeviceSleepOnDisconnect", (value))
}

// GetDeviceSleepOnDisconnect gets the value of DeviceSleepOnDisconnect for the instance
func (instance *MSNdis_PMAdminConfigParam) GetPropertyDeviceSleepOnDisconnect()(value MSNdis_PMAdminConfigState, err error) { 
    retValue, err := instance.GetProperty("DeviceSleepOnDisconnect")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(MSNdis_PMAdminConfigState); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " MSNdis_PMAdminConfigState is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = MSNdis_PMAdminConfigState(valuetmp)

    return
}

// SetHeader sets the value of Header for the instance
func (instance *MSNdis_PMAdminConfigParam) SetPropertyHeader(value MSNdis_ObjectHeader) (err error) { 
    return instance.SetProperty("Header", (value))
}

// GetHeader gets the value of Header for the instance
func (instance *MSNdis_PMAdminConfigParam) GetPropertyHeader()(value MSNdis_ObjectHeader, err error) { 
    retValue, err := instance.GetProperty("Header")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(MSNdis_ObjectHeader); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " MSNdis_ObjectHeader is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = MSNdis_ObjectHeader(valuetmp)

    return
}

// SetPMARPOffload sets the value of PMARPOffload for the instance
func (instance *MSNdis_PMAdminConfigParam) SetPropertyPMARPOffload(value MSNdis_PMAdminConfigState) (err error) { 
    return instance.SetProperty("PMARPOffload", (value))
}

// GetPMARPOffload gets the value of PMARPOffload for the instance
func (instance *MSNdis_PMAdminConfigParam) GetPropertyPMARPOffload()(value MSNdis_PMAdminConfigState, err error) { 
    retValue, err := instance.GetProperty("PMARPOffload")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(MSNdis_PMAdminConfigState); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " MSNdis_PMAdminConfigState is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = MSNdis_PMAdminConfigState(valuetmp)

    return
}

// SetPMNDOffload sets the value of PMNDOffload for the instance
func (instance *MSNdis_PMAdminConfigParam) SetPropertyPMNDOffload(value MSNdis_PMAdminConfigState) (err error) { 
    return instance.SetProperty("PMNDOffload", (value))
}

// GetPMNDOffload gets the value of PMNDOffload for the instance
func (instance *MSNdis_PMAdminConfigParam) GetPropertyPMNDOffload()(value MSNdis_PMAdminConfigState, err error) { 
    retValue, err := instance.GetProperty("PMNDOffload")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(MSNdis_PMAdminConfigState); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " MSNdis_PMAdminConfigState is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = MSNdis_PMAdminConfigState(valuetmp)

    return
}

// SetPMWiFiRekeyOffload sets the value of PMWiFiRekeyOffload for the instance
func (instance *MSNdis_PMAdminConfigParam) SetPropertyPMWiFiRekeyOffload(value MSNdis_PMAdminConfigState) (err error) { 
    return instance.SetProperty("PMWiFiRekeyOffload", (value))
}

// GetPMWiFiRekeyOffload gets the value of PMWiFiRekeyOffload for the instance
func (instance *MSNdis_PMAdminConfigParam) GetPropertyPMWiFiRekeyOffload()(value MSNdis_PMAdminConfigState, err error) { 
    retValue, err := instance.GetProperty("PMWiFiRekeyOffload")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(MSNdis_PMAdminConfigState); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " MSNdis_PMAdminConfigState is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = MSNdis_PMAdminConfigState(valuetmp)

    return
}

// SetWakeOnMagicPacket sets the value of WakeOnMagicPacket for the instance
func (instance *MSNdis_PMAdminConfigParam) SetPropertyWakeOnMagicPacket(value MSNdis_PMAdminConfigState) (err error) { 
    return instance.SetProperty("WakeOnMagicPacket", (value))
}

// GetWakeOnMagicPacket gets the value of WakeOnMagicPacket for the instance
func (instance *MSNdis_PMAdminConfigParam) GetPropertyWakeOnMagicPacket()(value MSNdis_PMAdminConfigState, err error) { 
    retValue, err := instance.GetProperty("WakeOnMagicPacket")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(MSNdis_PMAdminConfigState); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " MSNdis_PMAdminConfigState is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = MSNdis_PMAdminConfigState(valuetmp)

    return
}

// SetWakeOnPattern sets the value of WakeOnPattern for the instance
func (instance *MSNdis_PMAdminConfigParam) SetPropertyWakeOnPattern(value MSNdis_PMAdminConfigState) (err error) { 
    return instance.SetProperty("WakeOnPattern", (value))
}

// GetWakeOnPattern gets the value of WakeOnPattern for the instance
func (instance *MSNdis_PMAdminConfigParam) GetPropertyWakeOnPattern()(value MSNdis_PMAdminConfigState, err error) { 
    retValue, err := instance.GetProperty("WakeOnPattern")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(MSNdis_PMAdminConfigState); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " MSNdis_PMAdminConfigState is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = MSNdis_PMAdminConfigState(valuetmp)

    return
}

