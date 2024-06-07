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
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSFC_TM struct
type MSFC_TM struct { 
	*cim.WmiInstance

	// 
	tm_hour uint32

	// 
	tm_isdst uint32

	// 
	tm_mday uint32

	// 
	tm_min uint32

	// 
	tm_mon uint32

	// 
	tm_sec uint32

	// 
	tm_wday uint32

	// 
	tm_yday uint32

	// 
	tm_year uint32
}

	func NewMSFC_TMEx1(instance *cim.WmiInstance) (newInstance *MSFC_TM, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MSFC_TM {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMSFC_TMEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFC_TM, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFC_TM {
	WmiInstance: tmp,
	}
	return
	}
	

// Settm_hour sets the value of tm_hour for the instance
func (instance *MSFC_TM) SetPropertytm_hour(value uint32) (err error) { 
    return instance.SetProperty("tm_hour", (value))
}

// Gettm_hour gets the value of tm_hour for the instance
func (instance *MSFC_TM) GetPropertytm_hour()(value uint32, err error) { 
    retValue, err := instance.GetProperty("tm_hour")
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

// Settm_isdst sets the value of tm_isdst for the instance
func (instance *MSFC_TM) SetPropertytm_isdst(value uint32) (err error) { 
    return instance.SetProperty("tm_isdst", (value))
}

// Gettm_isdst gets the value of tm_isdst for the instance
func (instance *MSFC_TM) GetPropertytm_isdst()(value uint32, err error) { 
    retValue, err := instance.GetProperty("tm_isdst")
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

// Settm_mday sets the value of tm_mday for the instance
func (instance *MSFC_TM) SetPropertytm_mday(value uint32) (err error) { 
    return instance.SetProperty("tm_mday", (value))
}

// Gettm_mday gets the value of tm_mday for the instance
func (instance *MSFC_TM) GetPropertytm_mday()(value uint32, err error) { 
    retValue, err := instance.GetProperty("tm_mday")
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

// Settm_min sets the value of tm_min for the instance
func (instance *MSFC_TM) SetPropertytm_min(value uint32) (err error) { 
    return instance.SetProperty("tm_min", (value))
}

// Gettm_min gets the value of tm_min for the instance
func (instance *MSFC_TM) GetPropertytm_min()(value uint32, err error) { 
    retValue, err := instance.GetProperty("tm_min")
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

// Settm_mon sets the value of tm_mon for the instance
func (instance *MSFC_TM) SetPropertytm_mon(value uint32) (err error) { 
    return instance.SetProperty("tm_mon", (value))
}

// Gettm_mon gets the value of tm_mon for the instance
func (instance *MSFC_TM) GetPropertytm_mon()(value uint32, err error) { 
    retValue, err := instance.GetProperty("tm_mon")
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

// Settm_sec sets the value of tm_sec for the instance
func (instance *MSFC_TM) SetPropertytm_sec(value uint32) (err error) { 
    return instance.SetProperty("tm_sec", (value))
}

// Gettm_sec gets the value of tm_sec for the instance
func (instance *MSFC_TM) GetPropertytm_sec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("tm_sec")
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

// Settm_wday sets the value of tm_wday for the instance
func (instance *MSFC_TM) SetPropertytm_wday(value uint32) (err error) { 
    return instance.SetProperty("tm_wday", (value))
}

// Gettm_wday gets the value of tm_wday for the instance
func (instance *MSFC_TM) GetPropertytm_wday()(value uint32, err error) { 
    retValue, err := instance.GetProperty("tm_wday")
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

// Settm_yday sets the value of tm_yday for the instance
func (instance *MSFC_TM) SetPropertytm_yday(value uint32) (err error) { 
    return instance.SetProperty("tm_yday", (value))
}

// Gettm_yday gets the value of tm_yday for the instance
func (instance *MSFC_TM) GetPropertytm_yday()(value uint32, err error) { 
    retValue, err := instance.GetProperty("tm_yday")
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

// Settm_year sets the value of tm_year for the instance
func (instance *MSFC_TM) SetPropertytm_year(value uint32) (err error) { 
    return instance.SetProperty("tm_year", (value))
}

// Gettm_year gets the value of tm_year for the instance
func (instance *MSFC_TM) GetPropertytm_year()(value uint32, err error) { 
    retValue, err := instance.GetProperty("tm_year")
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

