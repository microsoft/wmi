// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

// __ProviderRegistration struct
type __ProviderRegistration struct {
	__SystemClass

	//
	provider __Provider
}

// Setprovider sets the value of provider for the instance
func (instance *__ProviderRegistration) SetPropertyprovider(value __Provider) (err error) {
	return instance.SetProperty("provider", value)
}

// Getprovider gets the value of provider for the instance
func (instance *__ProviderRegistration) GetPropertyprovider() (value __Provider, err error) {
	retValue, err := instance.GetProperty("provider")
	if err != nil {
		return
	}
	value, ok := retValue.(__Provider)
	if !ok {
		// TODO: Set an error
	}
	return
}
