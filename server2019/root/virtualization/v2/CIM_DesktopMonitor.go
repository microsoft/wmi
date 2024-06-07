// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// CIM_DesktopMonitor struct
type CIM_DesktopMonitor struct { 
	*CIM_Display

	// Monitor's bandwidth in MHertz. If unknown, enter 0.
	Bandwidth uint32

	// The type of DesktopMonitor or CRT. For example, multiscan color or monochrome monitors (values 2 or 3, respectively) can be indicated in this property.
	DisplayType DesktopMonitor_DisplayType

	// The logical height of the Display in screen coordinates.
	ScreenHeight uint32

	// The logical width of the Display in screen coordinates.
	ScreenWidth uint32
}

	func NewCIM_DesktopMonitorEx1(instance *cim.WmiInstance) (newInstance *CIM_DesktopMonitor, err error) {tmp, err := NewCIM_DisplayEx1(instance)
		
	if err != nil { return }
	newInstance = &CIM_DesktopMonitor {
	CIM_Display: tmp,
	}
	return
	}
	

	func NewCIM_DesktopMonitorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *CIM_DesktopMonitor, err error) {tmp, err := NewCIM_DisplayEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &CIM_DesktopMonitor {
	CIM_Display: tmp,
	}
	return
	}
	

// SetBandwidth sets the value of Bandwidth for the instance
func (instance *CIM_DesktopMonitor) SetPropertyBandwidth(value uint32) (err error) { 
    return instance.SetProperty("Bandwidth", (value))
}

// GetBandwidth gets the value of Bandwidth for the instance
func (instance *CIM_DesktopMonitor) GetPropertyBandwidth()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Bandwidth")
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

// SetDisplayType sets the value of DisplayType for the instance
func (instance *CIM_DesktopMonitor) SetPropertyDisplayType(value DesktopMonitor_DisplayType) (err error) { 
    return instance.SetProperty("DisplayType", (value))
}

// GetDisplayType gets the value of DisplayType for the instance
func (instance *CIM_DesktopMonitor) GetPropertyDisplayType()(value DesktopMonitor_DisplayType, err error) { 
    retValue, err := instance.GetProperty("DisplayType")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = DesktopMonitor_DisplayType(valuetmp)

    return
}

// SetScreenHeight sets the value of ScreenHeight for the instance
func (instance *CIM_DesktopMonitor) SetPropertyScreenHeight(value uint32) (err error) { 
    return instance.SetProperty("ScreenHeight", (value))
}

// GetScreenHeight gets the value of ScreenHeight for the instance
func (instance *CIM_DesktopMonitor) GetPropertyScreenHeight()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ScreenHeight")
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

// SetScreenWidth sets the value of ScreenWidth for the instance
func (instance *CIM_DesktopMonitor) SetPropertyScreenWidth(value uint32) (err error) { 
    return instance.SetProperty("ScreenWidth", (value))
}

// GetScreenWidth gets the value of ScreenWidth for the instance
func (instance *CIM_DesktopMonitor) GetPropertyScreenWidth()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ScreenWidth")
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

