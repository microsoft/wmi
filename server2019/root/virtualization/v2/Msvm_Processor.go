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

// Msvm_Processor struct
type Msvm_Processor struct { 
	*CIM_Processor

	// 
	LoadPercentageHistory []uint16
}

	func NewMsvm_ProcessorEx1(instance *cim.WmiInstance) (newInstance *Msvm_Processor, err error) {tmp, err := NewCIM_ProcessorEx1(instance)
		
	if err != nil { return }
	newInstance = &Msvm_Processor {
	CIM_Processor: tmp,
	}
	return
	}
	

	func NewMsvm_ProcessorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Msvm_Processor, err error) {tmp, err := NewCIM_ProcessorEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Msvm_Processor {
	CIM_Processor: tmp,
	}
	return
	}
	

// SetLoadPercentageHistory sets the value of LoadPercentageHistory for the instance
func (instance *Msvm_Processor) SetPropertyLoadPercentageHistory(value []uint16) (err error) { 
    return instance.SetProperty("LoadPercentageHistory", (value))
}

// GetLoadPercentageHistory gets the value of LoadPercentageHistory for the instance
func (instance *Msvm_Processor) GetPropertyLoadPercentageHistory()(value []uint16, err error) { 
    retValue, err := instance.GetProperty("LoadPercentageHistory")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(uint16); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, uint16(valuetmp))
    }

    return
}
func  (instance* Msvm_Processor) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
		 return instance.GetRelated("Msvm_ComputerSystem"); 
	}
	
func  (instance* Msvm_Processor) GetRelatedNumaNode() (value *cim.WmiInstance, err error) {
		 return instance.GetRelated("Msvm_NumaNode"); 
	}
	
func  (instance* Msvm_Processor) GetRelatedProcessorPool() (value *cim.WmiInstance, err error) {
		 return instance.GetRelated("Msvm_ProcessorPool"); 
	}
	

