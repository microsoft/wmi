// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// TransactionsDefaultSettingsSection struct
type TransactionsDefaultSettingsSection struct { 
	*ConfigurationSection

	// 
	DistributedTransactionManagerName string

	// 
	Timeout string
}

	func NewTransactionsDefaultSettingsSectionEx1(instance *cim.WmiInstance) (newInstance *TransactionsDefaultSettingsSection, err error) {tmp, err := NewConfigurationSectionEx1(instance)
		
	if err != nil { return }
	newInstance = &TransactionsDefaultSettingsSection {
	ConfigurationSection: tmp,
	}
	return
	}
	

	func NewTransactionsDefaultSettingsSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *TransactionsDefaultSettingsSection, err error) {tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &TransactionsDefaultSettingsSection {
	ConfigurationSection: tmp,
	}
	return
	}
	

// SetDistributedTransactionManagerName sets the value of DistributedTransactionManagerName for the instance
func (instance *TransactionsDefaultSettingsSection) SetPropertyDistributedTransactionManagerName(value string) (err error) { 
    return instance.SetProperty("DistributedTransactionManagerName", (value))
}

// GetDistributedTransactionManagerName gets the value of DistributedTransactionManagerName for the instance
func (instance *TransactionsDefaultSettingsSection) GetPropertyDistributedTransactionManagerName()(value string, err error) { 
    retValue, err := instance.GetProperty("DistributedTransactionManagerName")
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

// SetTimeout sets the value of Timeout for the instance
func (instance *TransactionsDefaultSettingsSection) SetPropertyTimeout(value string) (err error) { 
    return instance.SetProperty("Timeout", (value))
}

// GetTimeout gets the value of Timeout for the instance
func (instance *TransactionsDefaultSettingsSection) GetPropertyTimeout()(value string, err error) { 
    retValue, err := instance.GetProperty("Timeout")
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

