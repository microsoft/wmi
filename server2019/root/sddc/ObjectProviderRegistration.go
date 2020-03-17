// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.SDDC
//////////////////////////////////////////////
package sddc

// __ObjectProviderRegistration struct
type __ObjectProviderRegistration struct {
	__ProviderRegistration

	//
	InteractionType ObjectProviderRegistration_InteractionType

	//
	QuerySupportLevels []string

	//
	SupportsBatching bool

	//
	SupportsDelete bool

	//
	SupportsEnumeration bool

	//
	SupportsGet bool

	//
	SupportsPut bool

	//
	SupportsTransactions bool
}

// SetInteractionType sets the value of InteractionType for the instance
func (instance *__ObjectProviderRegistration) SetPropertyInteractionType(value ObjectProviderRegistration_InteractionType) (err error) {
	return instance.SetProperty("InteractionType", value)
}

// GetInteractionType gets the value of InteractionType for the instance
func (instance *__ObjectProviderRegistration) GetPropertyInteractionType() (value ObjectProviderRegistration_InteractionType, err error) {
	retValue, err := instance.GetProperty("InteractionType")
	if err != nil {
		return
	}
	value, ok := retValue.(ObjectProviderRegistration_InteractionType)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetQuerySupportLevels sets the value of QuerySupportLevels for the instance
func (instance *__ObjectProviderRegistration) SetPropertyQuerySupportLevels(value []string) (err error) {
	return instance.SetProperty("QuerySupportLevels", value)
}

// GetQuerySupportLevels gets the value of QuerySupportLevels for the instance
func (instance *__ObjectProviderRegistration) GetPropertyQuerySupportLevels() (value []string, err error) {
	retValue, err := instance.GetProperty("QuerySupportLevels")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportsBatching sets the value of SupportsBatching for the instance
func (instance *__ObjectProviderRegistration) SetPropertySupportsBatching(value bool) (err error) {
	return instance.SetProperty("SupportsBatching", value)
}

// GetSupportsBatching gets the value of SupportsBatching for the instance
func (instance *__ObjectProviderRegistration) GetPropertySupportsBatching() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsBatching")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportsDelete sets the value of SupportsDelete for the instance
func (instance *__ObjectProviderRegistration) SetPropertySupportsDelete(value bool) (err error) {
	return instance.SetProperty("SupportsDelete", value)
}

// GetSupportsDelete gets the value of SupportsDelete for the instance
func (instance *__ObjectProviderRegistration) GetPropertySupportsDelete() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsDelete")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportsEnumeration sets the value of SupportsEnumeration for the instance
func (instance *__ObjectProviderRegistration) SetPropertySupportsEnumeration(value bool) (err error) {
	return instance.SetProperty("SupportsEnumeration", value)
}

// GetSupportsEnumeration gets the value of SupportsEnumeration for the instance
func (instance *__ObjectProviderRegistration) GetPropertySupportsEnumeration() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsEnumeration")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportsGet sets the value of SupportsGet for the instance
func (instance *__ObjectProviderRegistration) SetPropertySupportsGet(value bool) (err error) {
	return instance.SetProperty("SupportsGet", value)
}

// GetSupportsGet gets the value of SupportsGet for the instance
func (instance *__ObjectProviderRegistration) GetPropertySupportsGet() (value bool, err error) {
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
func (instance *__ObjectProviderRegistration) SetPropertySupportsPut(value bool) (err error) {
	return instance.SetProperty("SupportsPut", value)
}

// GetSupportsPut gets the value of SupportsPut for the instance
func (instance *__ObjectProviderRegistration) GetPropertySupportsPut() (value bool, err error) {
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

// SetSupportsTransactions sets the value of SupportsTransactions for the instance
func (instance *__ObjectProviderRegistration) SetPropertySupportsTransactions(value bool) (err error) {
	return instance.SetProperty("SupportsTransactions", value)
}

// GetSupportsTransactions gets the value of SupportsTransactions for the instance
func (instance *__ObjectProviderRegistration) GetPropertySupportsTransactions() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsTransactions")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
