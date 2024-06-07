// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// ExtendedProtectionPolicy struct
type ExtendedProtectionPolicy struct { 
	*cim.WmiInstance

	// A channel binding token used when authenticating clients.
	CustomChannelBinding string

	// A list of service principal names accepted by the service.
	CustomServiceNames []string

	// A value that specifies when ExtendedProtection should be enforced.
	PolicyEnforcement string

	// A value that specifies the protection scenario being enforced by the policy.
	ProtectionScenario string
}

	func NewExtendedProtectionPolicyEx1(instance *cim.WmiInstance) (newInstance *ExtendedProtectionPolicy, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &ExtendedProtectionPolicy {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewExtendedProtectionPolicyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ExtendedProtectionPolicy, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ExtendedProtectionPolicy {
	WmiInstance: tmp,
	}
	return
	}
	

// SetCustomChannelBinding sets the value of CustomChannelBinding for the instance
func (instance *ExtendedProtectionPolicy) SetPropertyCustomChannelBinding(value string) (err error) { 
    return instance.SetProperty("CustomChannelBinding", (value))
}

// GetCustomChannelBinding gets the value of CustomChannelBinding for the instance
func (instance *ExtendedProtectionPolicy) GetPropertyCustomChannelBinding()(value string, err error) { 
    retValue, err := instance.GetProperty("CustomChannelBinding")
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

// SetCustomServiceNames sets the value of CustomServiceNames for the instance
func (instance *ExtendedProtectionPolicy) SetPropertyCustomServiceNames(value []string) (err error) { 
    return instance.SetProperty("CustomServiceNames", (value))
}

// GetCustomServiceNames gets the value of CustomServiceNames for the instance
func (instance *ExtendedProtectionPolicy) GetPropertyCustomServiceNames()(value []string, err error) { 
    retValue, err := instance.GetProperty("CustomServiceNames")
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

// SetPolicyEnforcement sets the value of PolicyEnforcement for the instance
func (instance *ExtendedProtectionPolicy) SetPropertyPolicyEnforcement(value string) (err error) { 
    return instance.SetProperty("PolicyEnforcement", (value))
}

// GetPolicyEnforcement gets the value of PolicyEnforcement for the instance
func (instance *ExtendedProtectionPolicy) GetPropertyPolicyEnforcement()(value string, err error) { 
    retValue, err := instance.GetProperty("PolicyEnforcement")
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

// SetProtectionScenario sets the value of ProtectionScenario for the instance
func (instance *ExtendedProtectionPolicy) SetPropertyProtectionScenario(value string) (err error) { 
    return instance.SetProperty("ProtectionScenario", (value))
}

// GetProtectionScenario gets the value of ProtectionScenario for the instance
func (instance *ExtendedProtectionPolicy) GetPropertyProtectionScenario()(value string, err error) { 
    retValue, err := instance.GetProperty("ProtectionScenario")
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

