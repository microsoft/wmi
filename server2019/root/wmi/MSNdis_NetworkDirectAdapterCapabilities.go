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

// MSNdis_NetworkDirectAdapterCapabilities struct
type MSNdis_NetworkDirectAdapterCapabilities struct { 
	*MSNdis

	// 
	MaxCqCount uint32

	// 
	MaxInboundReadLimit uint32

	// 
	MaxMrCount uint32

	// 
	MaxMwCount uint32

	// 
	MaxOutboundReadLimit uint32

	// 
	MaxPdCount uint32

	// 
	MaxQpCount uint32

	// 
	MaxSrqCount uint32

	// 
	MissingCounterMask uint64

	// 
	NdAdapterInfo MSNdis_NetworkDirectAdapterInfo
}

	func NewMSNdis_NetworkDirectAdapterCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *MSNdis_NetworkDirectAdapterCapabilities, err error) {tmp, err := NewMSNdisEx1(instance)
		
	if err != nil { return }
	newInstance = &MSNdis_NetworkDirectAdapterCapabilities {
	MSNdis: tmp,
	}
	return
	}
	

	func NewMSNdis_NetworkDirectAdapterCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSNdis_NetworkDirectAdapterCapabilities, err error) {tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSNdis_NetworkDirectAdapterCapabilities {
	MSNdis: tmp,
	}
	return
	}
	

// SetMaxCqCount sets the value of MaxCqCount for the instance
func (instance *MSNdis_NetworkDirectAdapterCapabilities) SetPropertyMaxCqCount(value uint32) (err error) { 
    return instance.SetProperty("MaxCqCount", (value))
}

// GetMaxCqCount gets the value of MaxCqCount for the instance
func (instance *MSNdis_NetworkDirectAdapterCapabilities) GetPropertyMaxCqCount()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MaxCqCount")
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

// SetMaxInboundReadLimit sets the value of MaxInboundReadLimit for the instance
func (instance *MSNdis_NetworkDirectAdapterCapabilities) SetPropertyMaxInboundReadLimit(value uint32) (err error) { 
    return instance.SetProperty("MaxInboundReadLimit", (value))
}

// GetMaxInboundReadLimit gets the value of MaxInboundReadLimit for the instance
func (instance *MSNdis_NetworkDirectAdapterCapabilities) GetPropertyMaxInboundReadLimit()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MaxInboundReadLimit")
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

// SetMaxMrCount sets the value of MaxMrCount for the instance
func (instance *MSNdis_NetworkDirectAdapterCapabilities) SetPropertyMaxMrCount(value uint32) (err error) { 
    return instance.SetProperty("MaxMrCount", (value))
}

// GetMaxMrCount gets the value of MaxMrCount for the instance
func (instance *MSNdis_NetworkDirectAdapterCapabilities) GetPropertyMaxMrCount()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MaxMrCount")
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

// SetMaxMwCount sets the value of MaxMwCount for the instance
func (instance *MSNdis_NetworkDirectAdapterCapabilities) SetPropertyMaxMwCount(value uint32) (err error) { 
    return instance.SetProperty("MaxMwCount", (value))
}

// GetMaxMwCount gets the value of MaxMwCount for the instance
func (instance *MSNdis_NetworkDirectAdapterCapabilities) GetPropertyMaxMwCount()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MaxMwCount")
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

// SetMaxOutboundReadLimit sets the value of MaxOutboundReadLimit for the instance
func (instance *MSNdis_NetworkDirectAdapterCapabilities) SetPropertyMaxOutboundReadLimit(value uint32) (err error) { 
    return instance.SetProperty("MaxOutboundReadLimit", (value))
}

// GetMaxOutboundReadLimit gets the value of MaxOutboundReadLimit for the instance
func (instance *MSNdis_NetworkDirectAdapterCapabilities) GetPropertyMaxOutboundReadLimit()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MaxOutboundReadLimit")
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

// SetMaxPdCount sets the value of MaxPdCount for the instance
func (instance *MSNdis_NetworkDirectAdapterCapabilities) SetPropertyMaxPdCount(value uint32) (err error) { 
    return instance.SetProperty("MaxPdCount", (value))
}

// GetMaxPdCount gets the value of MaxPdCount for the instance
func (instance *MSNdis_NetworkDirectAdapterCapabilities) GetPropertyMaxPdCount()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MaxPdCount")
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

// SetMaxQpCount sets the value of MaxQpCount for the instance
func (instance *MSNdis_NetworkDirectAdapterCapabilities) SetPropertyMaxQpCount(value uint32) (err error) { 
    return instance.SetProperty("MaxQpCount", (value))
}

// GetMaxQpCount gets the value of MaxQpCount for the instance
func (instance *MSNdis_NetworkDirectAdapterCapabilities) GetPropertyMaxQpCount()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MaxQpCount")
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

// SetMaxSrqCount sets the value of MaxSrqCount for the instance
func (instance *MSNdis_NetworkDirectAdapterCapabilities) SetPropertyMaxSrqCount(value uint32) (err error) { 
    return instance.SetProperty("MaxSrqCount", (value))
}

// GetMaxSrqCount gets the value of MaxSrqCount for the instance
func (instance *MSNdis_NetworkDirectAdapterCapabilities) GetPropertyMaxSrqCount()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MaxSrqCount")
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

// SetMissingCounterMask sets the value of MissingCounterMask for the instance
func (instance *MSNdis_NetworkDirectAdapterCapabilities) SetPropertyMissingCounterMask(value uint64) (err error) { 
    return instance.SetProperty("MissingCounterMask", (value))
}

// GetMissingCounterMask gets the value of MissingCounterMask for the instance
func (instance *MSNdis_NetworkDirectAdapterCapabilities) GetPropertyMissingCounterMask()(value uint64, err error) { 
    retValue, err := instance.GetProperty("MissingCounterMask")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetNdAdapterInfo sets the value of NdAdapterInfo for the instance
func (instance *MSNdis_NetworkDirectAdapterCapabilities) SetPropertyNdAdapterInfo(value MSNdis_NetworkDirectAdapterInfo) (err error) { 
    return instance.SetProperty("NdAdapterInfo", (value))
}

// GetNdAdapterInfo gets the value of NdAdapterInfo for the instance
func (instance *MSNdis_NetworkDirectAdapterCapabilities) GetPropertyNdAdapterInfo()(value MSNdis_NetworkDirectAdapterInfo, err error) { 
    retValue, err := instance.GetProperty("NdAdapterInfo")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(MSNdis_NetworkDirectAdapterInfo); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " MSNdis_NetworkDirectAdapterInfo is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = MSNdis_NetworkDirectAdapterInfo(valuetmp)

    return
}

