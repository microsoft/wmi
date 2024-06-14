// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// HttpClientCache struct
type HttpClientCache struct {
	*EmbeddedObject

	//
	CacheControlCustom string

	//
	CacheControlMaxAge string

	//
	CacheControlMode int32

	//
	HttpExpires string

	//
	SetEtag bool
}

func NewHttpClientCacheEx1(instance *cim.WmiInstance) (newInstance *HttpClientCache, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &HttpClientCache{
		EmbeddedObject: tmp,
	}
	return
}

func NewHttpClientCacheEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HttpClientCache, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HttpClientCache{
		EmbeddedObject: tmp,
	}
	return
}

// SetCacheControlCustom sets the value of CacheControlCustom for the instance
func (instance *HttpClientCache) SetPropertyCacheControlCustom(value string) (err error) {
	return instance.SetProperty("CacheControlCustom", (value))
}

// GetCacheControlCustom gets the value of CacheControlCustom for the instance
func (instance *HttpClientCache) GetPropertyCacheControlCustom() (value string, err error) {
	retValue, err := instance.GetProperty("CacheControlCustom")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetCacheControlMaxAge sets the value of CacheControlMaxAge for the instance
func (instance *HttpClientCache) SetPropertyCacheControlMaxAge(value string) (err error) {
	return instance.SetProperty("CacheControlMaxAge", (value))
}

// GetCacheControlMaxAge gets the value of CacheControlMaxAge for the instance
func (instance *HttpClientCache) GetPropertyCacheControlMaxAge() (value string, err error) {
	retValue, err := instance.GetProperty("CacheControlMaxAge")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetCacheControlMode sets the value of CacheControlMode for the instance
func (instance *HttpClientCache) SetPropertyCacheControlMode(value int32) (err error) {
	return instance.SetProperty("CacheControlMode", (value))
}

// GetCacheControlMode gets the value of CacheControlMode for the instance
func (instance *HttpClientCache) GetPropertyCacheControlMode() (value int32, err error) {
	retValue, err := instance.GetProperty("CacheControlMode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetHttpExpires sets the value of HttpExpires for the instance
func (instance *HttpClientCache) SetPropertyHttpExpires(value string) (err error) {
	return instance.SetProperty("HttpExpires", (value))
}

// GetHttpExpires gets the value of HttpExpires for the instance
func (instance *HttpClientCache) GetPropertyHttpExpires() (value string, err error) {
	retValue, err := instance.GetProperty("HttpExpires")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetSetEtag sets the value of SetEtag for the instance
func (instance *HttpClientCache) SetPropertySetEtag(value bool) (err error) {
	return instance.SetProperty("SetEtag", (value))
}

// GetSetEtag gets the value of SetEtag for the instance
func (instance *HttpClientCache) GetPropertySetEtag() (value bool, err error) {
	retValue, err := instance.GetProperty("SetEtag")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}
