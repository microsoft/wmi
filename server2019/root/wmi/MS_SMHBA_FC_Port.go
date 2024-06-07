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

// MS_SMHBA_FC_Port struct
type MS_SMHBA_FC_Port struct { 
	*cim.WmiInstance

	// 
	FabricName []uint8

	// 
	FcId uint32

	// 
	NodeWWN []uint8

	// 
	NumberofDiscoveredPorts uint32

	// 
	NumberofPhys uint8

	// 
	PortActiveFc4Types []uint8

	// 
	PortSupportedClassofService uint32

	// 
	PortSupportedFc4Types []uint8

	// 
	PortSymbolicName string

	// 
	PortWWN []uint8
}

	func NewMS_SMHBA_FC_PortEx1(instance *cim.WmiInstance) (newInstance *MS_SMHBA_FC_Port, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MS_SMHBA_FC_Port {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMS_SMHBA_FC_PortEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MS_SMHBA_FC_Port, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MS_SMHBA_FC_Port {
	WmiInstance: tmp,
	}
	return
	}
	

// SetFabricName sets the value of FabricName for the instance
func (instance *MS_SMHBA_FC_Port) SetPropertyFabricName(value []uint8) (err error) { 
    return instance.SetProperty("FabricName", (value))
}

// GetFabricName gets the value of FabricName for the instance
func (instance *MS_SMHBA_FC_Port) GetPropertyFabricName()(value []uint8, err error) { 
    retValue, err := instance.GetProperty("FabricName")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(uint8); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, uint8(valuetmp))
    }

    return
}

// SetFcId sets the value of FcId for the instance
func (instance *MS_SMHBA_FC_Port) SetPropertyFcId(value uint32) (err error) { 
    return instance.SetProperty("FcId", (value))
}

// GetFcId gets the value of FcId for the instance
func (instance *MS_SMHBA_FC_Port) GetPropertyFcId()(value uint32, err error) { 
    retValue, err := instance.GetProperty("FcId")
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

// SetNodeWWN sets the value of NodeWWN for the instance
func (instance *MS_SMHBA_FC_Port) SetPropertyNodeWWN(value []uint8) (err error) { 
    return instance.SetProperty("NodeWWN", (value))
}

// GetNodeWWN gets the value of NodeWWN for the instance
func (instance *MS_SMHBA_FC_Port) GetPropertyNodeWWN()(value []uint8, err error) { 
    retValue, err := instance.GetProperty("NodeWWN")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(uint8); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, uint8(valuetmp))
    }

    return
}

// SetNumberofDiscoveredPorts sets the value of NumberofDiscoveredPorts for the instance
func (instance *MS_SMHBA_FC_Port) SetPropertyNumberofDiscoveredPorts(value uint32) (err error) { 
    return instance.SetProperty("NumberofDiscoveredPorts", (value))
}

// GetNumberofDiscoveredPorts gets the value of NumberofDiscoveredPorts for the instance
func (instance *MS_SMHBA_FC_Port) GetPropertyNumberofDiscoveredPorts()(value uint32, err error) { 
    retValue, err := instance.GetProperty("NumberofDiscoveredPorts")
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

// SetNumberofPhys sets the value of NumberofPhys for the instance
func (instance *MS_SMHBA_FC_Port) SetPropertyNumberofPhys(value uint8) (err error) { 
    return instance.SetProperty("NumberofPhys", (value))
}

// GetNumberofPhys gets the value of NumberofPhys for the instance
func (instance *MS_SMHBA_FC_Port) GetPropertyNumberofPhys()(value uint8, err error) { 
    retValue, err := instance.GetProperty("NumberofPhys")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// SetPortActiveFc4Types sets the value of PortActiveFc4Types for the instance
func (instance *MS_SMHBA_FC_Port) SetPropertyPortActiveFc4Types(value []uint8) (err error) { 
    return instance.SetProperty("PortActiveFc4Types", (value))
}

// GetPortActiveFc4Types gets the value of PortActiveFc4Types for the instance
func (instance *MS_SMHBA_FC_Port) GetPropertyPortActiveFc4Types()(value []uint8, err error) { 
    retValue, err := instance.GetProperty("PortActiveFc4Types")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(uint8); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, uint8(valuetmp))
    }

    return
}

// SetPortSupportedClassofService sets the value of PortSupportedClassofService for the instance
func (instance *MS_SMHBA_FC_Port) SetPropertyPortSupportedClassofService(value uint32) (err error) { 
    return instance.SetProperty("PortSupportedClassofService", (value))
}

// GetPortSupportedClassofService gets the value of PortSupportedClassofService for the instance
func (instance *MS_SMHBA_FC_Port) GetPropertyPortSupportedClassofService()(value uint32, err error) { 
    retValue, err := instance.GetProperty("PortSupportedClassofService")
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

// SetPortSupportedFc4Types sets the value of PortSupportedFc4Types for the instance
func (instance *MS_SMHBA_FC_Port) SetPropertyPortSupportedFc4Types(value []uint8) (err error) { 
    return instance.SetProperty("PortSupportedFc4Types", (value))
}

// GetPortSupportedFc4Types gets the value of PortSupportedFc4Types for the instance
func (instance *MS_SMHBA_FC_Port) GetPropertyPortSupportedFc4Types()(value []uint8, err error) { 
    retValue, err := instance.GetProperty("PortSupportedFc4Types")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(uint8); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, uint8(valuetmp))
    }

    return
}

// SetPortSymbolicName sets the value of PortSymbolicName for the instance
func (instance *MS_SMHBA_FC_Port) SetPropertyPortSymbolicName(value string) (err error) { 
    return instance.SetProperty("PortSymbolicName", (value))
}

// GetPortSymbolicName gets the value of PortSymbolicName for the instance
func (instance *MS_SMHBA_FC_Port) GetPropertyPortSymbolicName()(value string, err error) { 
    retValue, err := instance.GetProperty("PortSymbolicName")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetPortWWN sets the value of PortWWN for the instance
func (instance *MS_SMHBA_FC_Port) SetPropertyPortWWN(value []uint8) (err error) { 
    return instance.SetProperty("PortWWN", (value))
}

// GetPortWWN gets the value of PortWWN for the instance
func (instance *MS_SMHBA_FC_Port) GetPropertyPortWWN()(value []uint8, err error) { 
    retValue, err := instance.GetProperty("PortWWN")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(uint8); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, uint8(valuetmp))
    }

    return
}

