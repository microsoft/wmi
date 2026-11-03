// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.SecurityCenter2
//////////////////////////////////////////////
package securitycenter2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// AntiSpywareProduct struct
type AntiSpywareProduct struct {
	*cim.WmiInstance

	//
	displayName string

	//
	instanceGuid string

	//
	pathToSignedProductExe string

	//
	pathToSignedReportingExe string

	//
	productState uint32

	//
	timestamp string
}

func NewAntiSpywareProductEx1(instance *cim.WmiInstance) (newInstance *AntiSpywareProduct, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &AntiSpywareProduct{
		WmiInstance: tmp,
	}
	return
}

func NewAntiSpywareProductEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *AntiSpywareProduct, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &AntiSpywareProduct{
		WmiInstance: tmp,
	}
	return
}

// SetdisplayName sets the value of displayName for the instance
func (instance *AntiSpywareProduct) SetPropertydisplayName(value string) (err error) {
	return instance.SetProperty("displayName", (value))
}

// GetdisplayName gets the value of displayName for the instance
func (instance *AntiSpywareProduct) GetPropertydisplayName() (value string, err error) {
	retValue, err := instance.GetProperty("displayName")
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

// SetinstanceGuid sets the value of instanceGuid for the instance
func (instance *AntiSpywareProduct) SetPropertyinstanceGuid(value string) (err error) {
	return instance.SetProperty("instanceGuid", (value))
}

// GetinstanceGuid gets the value of instanceGuid for the instance
func (instance *AntiSpywareProduct) GetPropertyinstanceGuid() (value string, err error) {
	retValue, err := instance.GetProperty("instanceGuid")
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

// SetpathToSignedProductExe sets the value of pathToSignedProductExe for the instance
func (instance *AntiSpywareProduct) SetPropertypathToSignedProductExe(value string) (err error) {
	return instance.SetProperty("pathToSignedProductExe", (value))
}

// GetpathToSignedProductExe gets the value of pathToSignedProductExe for the instance
func (instance *AntiSpywareProduct) GetPropertypathToSignedProductExe() (value string, err error) {
	retValue, err := instance.GetProperty("pathToSignedProductExe")
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

// SetpathToSignedReportingExe sets the value of pathToSignedReportingExe for the instance
func (instance *AntiSpywareProduct) SetPropertypathToSignedReportingExe(value string) (err error) {
	return instance.SetProperty("pathToSignedReportingExe", (value))
}

// GetpathToSignedReportingExe gets the value of pathToSignedReportingExe for the instance
func (instance *AntiSpywareProduct) GetPropertypathToSignedReportingExe() (value string, err error) {
	retValue, err := instance.GetProperty("pathToSignedReportingExe")
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

// SetproductState sets the value of productState for the instance
func (instance *AntiSpywareProduct) SetPropertyproductState(value uint32) (err error) {
	return instance.SetProperty("productState", (value))
}

// GetproductState gets the value of productState for the instance
func (instance *AntiSpywareProduct) GetPropertyproductState() (value uint32, err error) {
	retValue, err := instance.GetProperty("productState")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// Settimestamp sets the value of timestamp for the instance
func (instance *AntiSpywareProduct) SetPropertytimestamp(value string) (err error) {
	return instance.SetProperty("timestamp", (value))
}

// Gettimestamp gets the value of timestamp for the instance
func (instance *AntiSpywareProduct) GetPropertytimestamp() (value string, err error) {
	retValue, err := instance.GetProperty("timestamp")
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
