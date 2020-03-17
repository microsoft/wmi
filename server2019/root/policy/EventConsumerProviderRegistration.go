// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Policy
//////////////////////////////////////////////
package policy

// __EventConsumerProviderRegistration struct
type __EventConsumerProviderRegistration struct {
	__ProviderRegistration

	//
	ConsumerClassNames []string
}

// SetConsumerClassNames sets the value of ConsumerClassNames for the instance
func (instance *__EventConsumerProviderRegistration) SetPropertyConsumerClassNames(value []string) (err error) {
	return instance.SetProperty("ConsumerClassNames", value)
}

// GetConsumerClassNames gets the value of ConsumerClassNames for the instance
func (instance *__EventConsumerProviderRegistration) GetPropertyConsumerClassNames() (value []string, err error) {
	retValue, err := instance.GetProperty("ConsumerClassNames")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
