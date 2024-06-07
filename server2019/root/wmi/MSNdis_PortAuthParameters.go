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

// MSNdis_PortAuthParameters struct
type MSNdis_PortAuthParameters struct { 
	*MSNdis

	// 
	Header MSNdis_ObjectHeader

	// 
	RcvAuthorizationState uint32

	// 
	RcvControlState uint32

	// 
	SendAuthorizationState uint32

	// 
	SendControlState uint32
}

	func NewMSNdis_PortAuthParametersEx1(instance *cim.WmiInstance) (newInstance *MSNdis_PortAuthParameters, err error) {tmp, err := NewMSNdisEx1(instance)
		
	if err != nil { return }
	newInstance = &MSNdis_PortAuthParameters {
	MSNdis: tmp,
	}
	return
	}
	

	func NewMSNdis_PortAuthParametersEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSNdis_PortAuthParameters, err error) {tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSNdis_PortAuthParameters {
	MSNdis: tmp,
	}
	return
	}
	

// SetHeader sets the value of Header for the instance
func (instance *MSNdis_PortAuthParameters) SetPropertyHeader(value MSNdis_ObjectHeader) (err error) { 
    return instance.SetProperty("Header", (value))
}

// GetHeader gets the value of Header for the instance
func (instance *MSNdis_PortAuthParameters) GetPropertyHeader()(value MSNdis_ObjectHeader, err error) { 
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

// SetRcvAuthorizationState sets the value of RcvAuthorizationState for the instance
func (instance *MSNdis_PortAuthParameters) SetPropertyRcvAuthorizationState(value uint32) (err error) { 
    return instance.SetProperty("RcvAuthorizationState", (value))
}

// GetRcvAuthorizationState gets the value of RcvAuthorizationState for the instance
func (instance *MSNdis_PortAuthParameters) GetPropertyRcvAuthorizationState()(value uint32, err error) { 
    retValue, err := instance.GetProperty("RcvAuthorizationState")
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

// SetRcvControlState sets the value of RcvControlState for the instance
func (instance *MSNdis_PortAuthParameters) SetPropertyRcvControlState(value uint32) (err error) { 
    return instance.SetProperty("RcvControlState", (value))
}

// GetRcvControlState gets the value of RcvControlState for the instance
func (instance *MSNdis_PortAuthParameters) GetPropertyRcvControlState()(value uint32, err error) { 
    retValue, err := instance.GetProperty("RcvControlState")
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

// SetSendAuthorizationState sets the value of SendAuthorizationState for the instance
func (instance *MSNdis_PortAuthParameters) SetPropertySendAuthorizationState(value uint32) (err error) { 
    return instance.SetProperty("SendAuthorizationState", (value))
}

// GetSendAuthorizationState gets the value of SendAuthorizationState for the instance
func (instance *MSNdis_PortAuthParameters) GetPropertySendAuthorizationState()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SendAuthorizationState")
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

// SetSendControlState sets the value of SendControlState for the instance
func (instance *MSNdis_PortAuthParameters) SetPropertySendControlState(value uint32) (err error) { 
    return instance.SetProperty("SendControlState", (value))
}

// GetSendControlState gets the value of SendControlState for the instance
func (instance *MSNdis_PortAuthParameters) GetPropertySendControlState()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SendControlState")
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

