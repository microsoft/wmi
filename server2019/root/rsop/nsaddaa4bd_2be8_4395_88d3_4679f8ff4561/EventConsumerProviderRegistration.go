// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.RSOP.NSADDAA4BD_2BE8_4395_88D3_4679F8FF4561
//////////////////////////////////////////////
package nsaddaa4bd_2be8_4395_88d3_4679f8ff4561
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// __EventConsumerProviderRegistration struct
type __EventConsumerProviderRegistration struct { 
	*__ProviderRegistration

	// 
	ConsumerClassNames []string
}

	func New__EventConsumerProviderRegistrationEx1(instance *cim.WmiInstance) (newInstance *__EventConsumerProviderRegistration, err error) {tmp, err := New__ProviderRegistrationEx1(instance)
		
	if err != nil { return }
	newInstance = &__EventConsumerProviderRegistration {
	__ProviderRegistration: tmp,
	}
	return
	}
	

	func New__EventConsumerProviderRegistrationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *__EventConsumerProviderRegistration, err error) {tmp, err := New__ProviderRegistrationEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &__EventConsumerProviderRegistration {
	__ProviderRegistration: tmp,
	}
	return
	}
	

// SetConsumerClassNames sets the value of ConsumerClassNames for the instance
func (instance *__EventConsumerProviderRegistration) SetPropertyConsumerClassNames(value []string) (err error) { 
    return instance.SetProperty("ConsumerClassNames", (value))
}

// GetConsumerClassNames gets the value of ConsumerClassNames for the instance
func (instance *__EventConsumerProviderRegistration) GetPropertyConsumerClassNames()(value []string, err error) { 
    retValue, err := instance.GetProperty("ConsumerClassNames")
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

