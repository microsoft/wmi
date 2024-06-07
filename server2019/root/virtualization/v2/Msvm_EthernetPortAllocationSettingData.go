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

// Msvm_EthernetPortAllocationSettingData struct
type Msvm_EthernetPortAllocationSettingData struct { 
	*CIM_EthernetPortAllocationSettingData

	// 
	CompartmentGuid string

	// EnabledState is an integer enumeration that indicates whether the allocation request is enabled or disabled. When an allocation request is marked as Disabled (3), then the allocation is not processed. The EnabledState for an active configuration is always marked as Enabled (2).
	EnabledState EthernetPortAllocationSettingData_EnabledState

	// The last known friendly name of the switch this port had a hard affinity to, if any.
	LastKnownSwitchName string

	// 
	PortName string

	// 
	RequiredFeatureHints []string

	// 
	RequiredFeatures []string

	// 
	TestReplicaPoolID string

	// 
	TestReplicaSwitchName string
}

	func NewMsvm_EthernetPortAllocationSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_EthernetPortAllocationSettingData, err error) {tmp, err := NewCIM_EthernetPortAllocationSettingDataEx1(instance)
		
	if err != nil { return }
	newInstance = &Msvm_EthernetPortAllocationSettingData {
	CIM_EthernetPortAllocationSettingData: tmp,
	}
	return
	}
	

	func NewMsvm_EthernetPortAllocationSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Msvm_EthernetPortAllocationSettingData, err error) {tmp, err := NewCIM_EthernetPortAllocationSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Msvm_EthernetPortAllocationSettingData {
	CIM_EthernetPortAllocationSettingData: tmp,
	}
	return
	}
	

// SetCompartmentGuid sets the value of CompartmentGuid for the instance
func (instance *Msvm_EthernetPortAllocationSettingData) SetPropertyCompartmentGuid(value string) (err error) { 
    return instance.SetProperty("CompartmentGuid", (value))
}

// GetCompartmentGuid gets the value of CompartmentGuid for the instance
func (instance *Msvm_EthernetPortAllocationSettingData) GetPropertyCompartmentGuid()(value string, err error) { 
    retValue, err := instance.GetProperty("CompartmentGuid")
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

// SetEnabledState sets the value of EnabledState for the instance
func (instance *Msvm_EthernetPortAllocationSettingData) SetPropertyEnabledState(value EthernetPortAllocationSettingData_EnabledState) (err error) { 
    return instance.SetProperty("EnabledState", (value))
}

// GetEnabledState gets the value of EnabledState for the instance
func (instance *Msvm_EthernetPortAllocationSettingData) GetPropertyEnabledState()(value EthernetPortAllocationSettingData_EnabledState, err error) { 
    retValue, err := instance.GetProperty("EnabledState")
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

    value = EthernetPortAllocationSettingData_EnabledState(valuetmp)

    return
}

// SetLastKnownSwitchName sets the value of LastKnownSwitchName for the instance
func (instance *Msvm_EthernetPortAllocationSettingData) SetPropertyLastKnownSwitchName(value string) (err error) { 
    return instance.SetProperty("LastKnownSwitchName", (value))
}

// GetLastKnownSwitchName gets the value of LastKnownSwitchName for the instance
func (instance *Msvm_EthernetPortAllocationSettingData) GetPropertyLastKnownSwitchName()(value string, err error) { 
    retValue, err := instance.GetProperty("LastKnownSwitchName")
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

// SetPortName sets the value of PortName for the instance
func (instance *Msvm_EthernetPortAllocationSettingData) SetPropertyPortName(value string) (err error) { 
    return instance.SetProperty("PortName", (value))
}

// GetPortName gets the value of PortName for the instance
func (instance *Msvm_EthernetPortAllocationSettingData) GetPropertyPortName()(value string, err error) { 
    retValue, err := instance.GetProperty("PortName")
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

// SetRequiredFeatureHints sets the value of RequiredFeatureHints for the instance
func (instance *Msvm_EthernetPortAllocationSettingData) SetPropertyRequiredFeatureHints(value []string) (err error) { 
    return instance.SetProperty("RequiredFeatureHints", (value))
}

// GetRequiredFeatureHints gets the value of RequiredFeatureHints for the instance
func (instance *Msvm_EthernetPortAllocationSettingData) GetPropertyRequiredFeatureHints()(value []string, err error) { 
    retValue, err := instance.GetProperty("RequiredFeatureHints")
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

// SetRequiredFeatures sets the value of RequiredFeatures for the instance
func (instance *Msvm_EthernetPortAllocationSettingData) SetPropertyRequiredFeatures(value []string) (err error) { 
    return instance.SetProperty("RequiredFeatures", (value))
}

// GetRequiredFeatures gets the value of RequiredFeatures for the instance
func (instance *Msvm_EthernetPortAllocationSettingData) GetPropertyRequiredFeatures()(value []string, err error) { 
    retValue, err := instance.GetProperty("RequiredFeatures")
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

// SetTestReplicaPoolID sets the value of TestReplicaPoolID for the instance
func (instance *Msvm_EthernetPortAllocationSettingData) SetPropertyTestReplicaPoolID(value string) (err error) { 
    return instance.SetProperty("TestReplicaPoolID", (value))
}

// GetTestReplicaPoolID gets the value of TestReplicaPoolID for the instance
func (instance *Msvm_EthernetPortAllocationSettingData) GetPropertyTestReplicaPoolID()(value string, err error) { 
    retValue, err := instance.GetProperty("TestReplicaPoolID")
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

// SetTestReplicaSwitchName sets the value of TestReplicaSwitchName for the instance
func (instance *Msvm_EthernetPortAllocationSettingData) SetPropertyTestReplicaSwitchName(value string) (err error) { 
    return instance.SetProperty("TestReplicaSwitchName", (value))
}

// GetTestReplicaSwitchName gets the value of TestReplicaSwitchName for the instance
func (instance *Msvm_EthernetPortAllocationSettingData) GetPropertyTestReplicaSwitchName()(value string, err error) { 
    retValue, err := instance.GetProperty("TestReplicaSwitchName")
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
func  (instance* Msvm_EthernetPortAllocationSettingData) GetRelatedVirtualEthernetSwitchSettingData() (value *cim.WmiInstance, err error) {
		 return instance.GetRelated("Msvm_VirtualEthernetSwitchSettingData"); 
	}
	
func  (instance* Msvm_EthernetPortAllocationSettingData) GetRelatedEthernetSwitchPort() (value *cim.WmiInstance, err error) {
		 return instance.GetRelated("Msvm_EthernetSwitchPort"); 
	}
	

