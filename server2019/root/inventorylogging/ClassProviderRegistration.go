// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.InventoryLogging
//////////////////////////////////////////////
package inventorylogging

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __ClassProviderRegistration struct
type __ClassProviderRegistration struct {
	*__ObjectProviderRegistration

	//
	CacheRefreshInterval string

	//
	PerUserSchema bool

	//
	ReferencedSetQueries []string

	//
	ResultSetQueries []string

	//
	ReSynchroniseOnNamespaceOpen bool

	//
	UnsupportedQueries []string

	//
	Version uint32
}

func New__ClassProviderRegistrationEx1(instance *cim.WmiInstance) (newInstance *__ClassProviderRegistration, err error) {
	tmp, err := New__ObjectProviderRegistrationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__ClassProviderRegistration{
		__ObjectProviderRegistration: tmp,
	}
	return
}

func New__ClassProviderRegistrationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__ClassProviderRegistration, err error) {
	tmp, err := New__ObjectProviderRegistrationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__ClassProviderRegistration{
		__ObjectProviderRegistration: tmp,
	}
	return
}

// SetCacheRefreshInterval sets the value of CacheRefreshInterval for the instance
func (instance *__ClassProviderRegistration) SetPropertyCacheRefreshInterval(value string) (err error) {
	return instance.SetProperty("CacheRefreshInterval", value)
}

// GetCacheRefreshInterval gets the value of CacheRefreshInterval for the instance
func (instance *__ClassProviderRegistration) GetPropertyCacheRefreshInterval() (value string, err error) {
	retValue, err := instance.GetProperty("CacheRefreshInterval")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPerUserSchema sets the value of PerUserSchema for the instance
func (instance *__ClassProviderRegistration) SetPropertyPerUserSchema(value bool) (err error) {
	return instance.SetProperty("PerUserSchema", value)
}

// GetPerUserSchema gets the value of PerUserSchema for the instance
func (instance *__ClassProviderRegistration) GetPropertyPerUserSchema() (value bool, err error) {
	retValue, err := instance.GetProperty("PerUserSchema")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReferencedSetQueries sets the value of ReferencedSetQueries for the instance
func (instance *__ClassProviderRegistration) SetPropertyReferencedSetQueries(value []string) (err error) {
	return instance.SetProperty("ReferencedSetQueries", value)
}

// GetReferencedSetQueries gets the value of ReferencedSetQueries for the instance
func (instance *__ClassProviderRegistration) GetPropertyReferencedSetQueries() (value []string, err error) {
	retValue, err := instance.GetProperty("ReferencedSetQueries")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResultSetQueries sets the value of ResultSetQueries for the instance
func (instance *__ClassProviderRegistration) SetPropertyResultSetQueries(value []string) (err error) {
	return instance.SetProperty("ResultSetQueries", value)
}

// GetResultSetQueries gets the value of ResultSetQueries for the instance
func (instance *__ClassProviderRegistration) GetPropertyResultSetQueries() (value []string, err error) {
	retValue, err := instance.GetProperty("ResultSetQueries")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReSynchroniseOnNamespaceOpen sets the value of ReSynchroniseOnNamespaceOpen for the instance
func (instance *__ClassProviderRegistration) SetPropertyReSynchroniseOnNamespaceOpen(value bool) (err error) {
	return instance.SetProperty("ReSynchroniseOnNamespaceOpen", value)
}

// GetReSynchroniseOnNamespaceOpen gets the value of ReSynchroniseOnNamespaceOpen for the instance
func (instance *__ClassProviderRegistration) GetPropertyReSynchroniseOnNamespaceOpen() (value bool, err error) {
	retValue, err := instance.GetProperty("ReSynchroniseOnNamespaceOpen")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUnsupportedQueries sets the value of UnsupportedQueries for the instance
func (instance *__ClassProviderRegistration) SetPropertyUnsupportedQueries(value []string) (err error) {
	return instance.SetProperty("UnsupportedQueries", value)
}

// GetUnsupportedQueries gets the value of UnsupportedQueries for the instance
func (instance *__ClassProviderRegistration) GetPropertyUnsupportedQueries() (value []string, err error) {
	retValue, err := instance.GetProperty("UnsupportedQueries")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVersion sets the value of Version for the instance
func (instance *__ClassProviderRegistration) SetPropertyVersion(value uint32) (err error) {
	return instance.SetProperty("Version", value)
}

// GetVersion gets the value of Version for the instance
func (instance *__ClassProviderRegistration) GetPropertyVersion() (value uint32, err error) {
	retValue, err := instance.GetProperty("Version")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
