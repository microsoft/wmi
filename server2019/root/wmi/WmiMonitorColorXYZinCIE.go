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

// WmiMonitorColorXYZinCIE struct
type WmiMonitorColorXYZinCIE struct { 
	*cim.WmiInstance

	// 
	X uint16

	// 
	Y uint16
}

	func NewWmiMonitorColorXYZinCIEEx1(instance *cim.WmiInstance) (newInstance *WmiMonitorColorXYZinCIE, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &WmiMonitorColorXYZinCIE {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewWmiMonitorColorXYZinCIEEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *WmiMonitorColorXYZinCIE, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &WmiMonitorColorXYZinCIE {
	WmiInstance: tmp,
	}
	return
	}
	

// SetX sets the value of X for the instance
func (instance *WmiMonitorColorXYZinCIE) SetPropertyX(value uint16) (err error) { 
    return instance.SetProperty("X", (value))
}

// GetX gets the value of X for the instance
func (instance *WmiMonitorColorXYZinCIE) GetPropertyX()(value uint16, err error) { 
    retValue, err := instance.GetProperty("X")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint16); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint16(valuetmp)

    return
}

// SetY sets the value of Y for the instance
func (instance *WmiMonitorColorXYZinCIE) SetPropertyY(value uint16) (err error) { 
    return instance.SetProperty("Y", (value))
}

// GetY gets the value of Y for the instance
func (instance *WmiMonitorColorXYZinCIE) GetPropertyY()(value uint16, err error) { 
    retValue, err := instance.GetProperty("Y")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint16); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint16(valuetmp)

    return
}

