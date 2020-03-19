// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root
//////////////////////////////////////////////
package root

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __ObjectProviderCacheControl struct
type __ObjectProviderCacheControl struct {
	*__CacheControl

	//
	ClearAfter string
}

func New__ObjectProviderCacheControlEx1(instance *cim.WmiInstance) (newInstance *__ObjectProviderCacheControl, err error) {
	tmp, err := New__CacheControlEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__ObjectProviderCacheControl{
		__CacheControl: tmp,
	}
	return
}

func New__ObjectProviderCacheControlEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__ObjectProviderCacheControl, err error) {
	tmp, err := New__CacheControlEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__ObjectProviderCacheControl{
		__CacheControl: tmp,
	}
	return
}

// SetClearAfter sets the value of ClearAfter for the instance
func (instance *__ObjectProviderCacheControl) SetPropertyClearAfter(value string) (err error) {
	return instance.SetProperty("ClearAfter", value)
}

// GetClearAfter gets the value of ClearAfter for the instance
func (instance *__ObjectProviderCacheControl) GetPropertyClearAfter() (value string, err error) {
	retValue, err := instance.GetProperty("ClearAfter")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
