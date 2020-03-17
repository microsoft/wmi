// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Attestation
//////////////////////////////////////////////
package attestation

// __PropertyProviderRegistration struct
type __PropertyProviderRegistration struct {
	__ProviderRegistration

	//
	SupportsGet bool

	//
	SupportsPut bool
}

// SetSupportsGet sets the value of SupportsGet for the instance
func (instance *__PropertyProviderRegistration) SetPropertySupportsGet(value bool) (err error) {
	return instance.SetProperty("SupportsGet", value)
}

// GetSupportsGet gets the value of SupportsGet for the instance
func (instance *__PropertyProviderRegistration) GetPropertySupportsGet() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsGet")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportsPut sets the value of SupportsPut for the instance
func (instance *__PropertyProviderRegistration) SetPropertySupportsPut(value bool) (err error) {
	return instance.SetProperty("SupportsPut", value)
}

// GetSupportsPut gets the value of SupportsPut for the instance
func (instance *__PropertyProviderRegistration) GetPropertySupportsPut() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsPut")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
