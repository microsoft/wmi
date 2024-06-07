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

// MSNdis_ReceiveQueueInfoArray struct
type MSNdis_ReceiveQueueInfoArray struct { 
	*MSNdis

	// 
	ElementSize uint32

	// 
	FirstElementOffset uint32

	// 
	Header MSNdis_ObjectHeader

	// 
	NumElements uint32

	// 
	Queue []MSNdis_ReceiveQueueInfo
}

	func NewMSNdis_ReceiveQueueInfoArrayEx1(instance *cim.WmiInstance) (newInstance *MSNdis_ReceiveQueueInfoArray, err error) {tmp, err := NewMSNdisEx1(instance)
		
	if err != nil { return }
	newInstance = &MSNdis_ReceiveQueueInfoArray {
	MSNdis: tmp,
	}
	return
	}
	

	func NewMSNdis_ReceiveQueueInfoArrayEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSNdis_ReceiveQueueInfoArray, err error) {tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSNdis_ReceiveQueueInfoArray {
	MSNdis: tmp,
	}
	return
	}
	

// SetElementSize sets the value of ElementSize for the instance
func (instance *MSNdis_ReceiveQueueInfoArray) SetPropertyElementSize(value uint32) (err error) { 
    return instance.SetProperty("ElementSize", (value))
}

// GetElementSize gets the value of ElementSize for the instance
func (instance *MSNdis_ReceiveQueueInfoArray) GetPropertyElementSize()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ElementSize")
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

// SetFirstElementOffset sets the value of FirstElementOffset for the instance
func (instance *MSNdis_ReceiveQueueInfoArray) SetPropertyFirstElementOffset(value uint32) (err error) { 
    return instance.SetProperty("FirstElementOffset", (value))
}

// GetFirstElementOffset gets the value of FirstElementOffset for the instance
func (instance *MSNdis_ReceiveQueueInfoArray) GetPropertyFirstElementOffset()(value uint32, err error) { 
    retValue, err := instance.GetProperty("FirstElementOffset")
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

// SetHeader sets the value of Header for the instance
func (instance *MSNdis_ReceiveQueueInfoArray) SetPropertyHeader(value MSNdis_ObjectHeader) (err error) { 
    return instance.SetProperty("Header", (value))
}

// GetHeader gets the value of Header for the instance
func (instance *MSNdis_ReceiveQueueInfoArray) GetPropertyHeader()(value MSNdis_ObjectHeader, err error) { 
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

// SetNumElements sets the value of NumElements for the instance
func (instance *MSNdis_ReceiveQueueInfoArray) SetPropertyNumElements(value uint32) (err error) { 
    return instance.SetProperty("NumElements", (value))
}

// GetNumElements gets the value of NumElements for the instance
func (instance *MSNdis_ReceiveQueueInfoArray) GetPropertyNumElements()(value uint32, err error) { 
    retValue, err := instance.GetProperty("NumElements")
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

// SetQueue sets the value of Queue for the instance
func (instance *MSNdis_ReceiveQueueInfoArray) SetPropertyQueue(value []MSNdis_ReceiveQueueInfo) (err error) { 
    return instance.SetProperty("Queue", (value))
}

// GetQueue gets the value of Queue for the instance
func (instance *MSNdis_ReceiveQueueInfoArray) GetPropertyQueue()(value []MSNdis_ReceiveQueueInfo, err error) { 
    retValue, err := instance.GetProperty("Queue")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(MSNdis_ReceiveQueueInfo); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " MSNdis_ReceiveQueueInfo is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, MSNdis_ReceiveQueueInfo(valuetmp))
    }

    return
}

