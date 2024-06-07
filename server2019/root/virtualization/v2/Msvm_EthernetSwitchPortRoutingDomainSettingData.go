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

// Msvm_EthernetSwitchPortRoutingDomainSettingData struct
type Msvm_EthernetSwitchPortRoutingDomainSettingData struct { 
	*Msvm_EthernetSwitchPortFeatureSettingData

	// 
	IsolationIdList []uint32

	// 
	IsolationIdNameList []string

	// 
	RoutingDomainGuid string

	// 
	RoutingDomainName string
}

	func NewMsvm_EthernetSwitchPortRoutingDomainSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_EthernetSwitchPortRoutingDomainSettingData, err error) {tmp, err := NewMsvm_EthernetSwitchPortFeatureSettingDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Msvm_EthernetSwitchPortRoutingDomainSettingData {
	Msvm_EthernetSwitchPortFeatureSettingData: tmp,
	}
	return
	}
	

	func NewMsvm_EthernetSwitchPortRoutingDomainSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Msvm_EthernetSwitchPortRoutingDomainSettingData, err error) {tmp, err := NewMsvm_EthernetSwitchPortFeatureSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Msvm_EthernetSwitchPortRoutingDomainSettingData {
	Msvm_EthernetSwitchPortFeatureSettingData: tmp,
	}
	return
	}
	

// SetIsolationIdList sets the value of IsolationIdList for the instance
func (instance *Msvm_EthernetSwitchPortRoutingDomainSettingData) SetPropertyIsolationIdList(value []uint32) (err error) { 
    return instance.SetProperty("IsolationIdList", (value))
}

// GetIsolationIdList gets the value of IsolationIdList for the instance
func (instance *Msvm_EthernetSwitchPortRoutingDomainSettingData) GetPropertyIsolationIdList()(value []uint32, err error) { 
    retValue, err := instance.GetProperty("IsolationIdList")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(uint32); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, uint32(valuetmp))
    }

    return
}

// SetIsolationIdNameList sets the value of IsolationIdNameList for the instance
func (instance *Msvm_EthernetSwitchPortRoutingDomainSettingData) SetPropertyIsolationIdNameList(value []string) (err error) { 
    return instance.SetProperty("IsolationIdNameList", (value))
}

// GetIsolationIdNameList gets the value of IsolationIdNameList for the instance
func (instance *Msvm_EthernetSwitchPortRoutingDomainSettingData) GetPropertyIsolationIdNameList()(value []string, err error) { 
    retValue, err := instance.GetProperty("IsolationIdNameList")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(string); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, string(valuetmp))
    }

    return
}

// SetRoutingDomainGuid sets the value of RoutingDomainGuid for the instance
func (instance *Msvm_EthernetSwitchPortRoutingDomainSettingData) SetPropertyRoutingDomainGuid(value string) (err error) { 
    return instance.SetProperty("RoutingDomainGuid", (value))
}

// GetRoutingDomainGuid gets the value of RoutingDomainGuid for the instance
func (instance *Msvm_EthernetSwitchPortRoutingDomainSettingData) GetPropertyRoutingDomainGuid()(value string, err error) { 
    retValue, err := instance.GetProperty("RoutingDomainGuid")
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

// SetRoutingDomainName sets the value of RoutingDomainName for the instance
func (instance *Msvm_EthernetSwitchPortRoutingDomainSettingData) SetPropertyRoutingDomainName(value string) (err error) { 
    return instance.SetProperty("RoutingDomainName", (value))
}

// GetRoutingDomainName gets the value of RoutingDomainName for the instance
func (instance *Msvm_EthernetSwitchPortRoutingDomainSettingData) GetPropertyRoutingDomainName()(value string, err error) { 
    retValue, err := instance.GetProperty("RoutingDomainName")
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
func  (instance* Msvm_EthernetSwitchPortRoutingDomainSettingData) GetRelatedEthernetSwitchFeatureCapabilities() (value *cim.WmiInstance, err error) {
		 return instance.GetRelated("Msvm_EthernetSwitchFeatureCapabilities"); 
	}
	

