// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WebAdministration
//
// ////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CachingProfileElement struct
type CachingProfileElement struct {
	*CollectionElement

	//
	Duration string

	//
	Extension string

	//
	KernelCachePolicy int32

	//
	Location int32

	//
	Policy int32

	//
	VaryByHeaders string

	//
	VaryByQueryString string
}

func NewCachingProfileElementEx1(instance *cim.WmiInstance) (newInstance *CachingProfileElement, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CachingProfileElement{
		CollectionElement: tmp,
	}
	return
}

func NewCachingProfileElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CachingProfileElement, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CachingProfileElement{
		CollectionElement: tmp,
	}
	return
}

// SetDuration sets the value of Duration for the instance
func (instance *CachingProfileElement) SetPropertyDuration(value string) (err error) {
	return instance.SetProperty("Duration", (value))
}

// GetDuration gets the value of Duration for the instance
func (instance *CachingProfileElement) GetPropertyDuration() (value string, err error) {
	retValue, err := instance.GetProperty("Duration")
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

// SetExtension sets the value of Extension for the instance
func (instance *CachingProfileElement) SetPropertyExtension(value string) (err error) {
	return instance.SetProperty("Extension", (value))
}

// GetExtension gets the value of Extension for the instance
func (instance *CachingProfileElement) GetPropertyExtension() (value string, err error) {
	retValue, err := instance.GetProperty("Extension")
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

// SetKernelCachePolicy sets the value of KernelCachePolicy for the instance
func (instance *CachingProfileElement) SetPropertyKernelCachePolicy(value int32) (err error) {
	return instance.SetProperty("KernelCachePolicy", (value))
}

// GetKernelCachePolicy gets the value of KernelCachePolicy for the instance
func (instance *CachingProfileElement) GetPropertyKernelCachePolicy() (value int32, err error) {
	retValue, err := instance.GetProperty("KernelCachePolicy")
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

// SetLocation sets the value of Location for the instance
func (instance *CachingProfileElement) SetPropertyLocation(value int32) (err error) {
	return instance.SetProperty("Location", (value))
}

// GetLocation gets the value of Location for the instance
func (instance *CachingProfileElement) GetPropertyLocation() (value int32, err error) {
	retValue, err := instance.GetProperty("Location")
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

// SetPolicy sets the value of Policy for the instance
func (instance *CachingProfileElement) SetPropertyPolicy(value int32) (err error) {
	return instance.SetProperty("Policy", (value))
}

// GetPolicy gets the value of Policy for the instance
func (instance *CachingProfileElement) GetPropertyPolicy() (value int32, err error) {
	retValue, err := instance.GetProperty("Policy")
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

// SetVaryByHeaders sets the value of VaryByHeaders for the instance
func (instance *CachingProfileElement) SetPropertyVaryByHeaders(value string) (err error) {
	return instance.SetProperty("VaryByHeaders", (value))
}

// GetVaryByHeaders gets the value of VaryByHeaders for the instance
func (instance *CachingProfileElement) GetPropertyVaryByHeaders() (value string, err error) {
	retValue, err := instance.GetProperty("VaryByHeaders")
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

// SetVaryByQueryString sets the value of VaryByQueryString for the instance
func (instance *CachingProfileElement) SetPropertyVaryByQueryString(value string) (err error) {
	return instance.SetProperty("VaryByQueryString", (value))
}

// GetVaryByQueryString gets the value of VaryByQueryString for the instance
func (instance *CachingProfileElement) GetPropertyVaryByQueryString() (value string, err error) {
	retValue, err := instance.GetProperty("VaryByQueryString")
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
