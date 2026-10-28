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

// FirewallProduct struct
type FirewallProduct struct {
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

func NewFirewallProductEx1(instance *cim.WmiInstance) (newInstance *FirewallProduct, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &FirewallProduct{
		WmiInstance: tmp,
	}
	return
}

func NewFirewallProductEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FirewallProduct, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FirewallProduct{
		WmiInstance: tmp,
	}
	return
}

// SetdisplayName sets the value of displayName for the instance
func (instance *FirewallProduct) SetPropertydisplayName(value string) (err error) {
	return instance.SetProperty("displayName", (value))
}

// GetdisplayName gets the value of displayName for the instance
func (instance *FirewallProduct) GetPropertydisplayName() (value string, err error) {
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
func (instance *FirewallProduct) SetPropertyinstanceGuid(value string) (err error) {
	return instance.SetProperty("instanceGuid", (value))
}

// GetinstanceGuid gets the value of instanceGuid for the instance
func (instance *FirewallProduct) GetPropertyinstanceGuid() (value string, err error) {
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
func (instance *FirewallProduct) SetPropertypathToSignedProductExe(value string) (err error) {
	return instance.SetProperty("pathToSignedProductExe", (value))
}

// GetpathToSignedProductExe gets the value of pathToSignedProductExe for the instance
func (instance *FirewallProduct) GetPropertypathToSignedProductExe() (value string, err error) {
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
func (instance *FirewallProduct) SetPropertypathToSignedReportingExe(value string) (err error) {
	return instance.SetProperty("pathToSignedReportingExe", (value))
}

// GetpathToSignedReportingExe gets the value of pathToSignedReportingExe for the instance
func (instance *FirewallProduct) GetPropertypathToSignedReportingExe() (value string, err error) {
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
func (instance *FirewallProduct) SetPropertyproductState(value uint32) (err error) {
	return instance.SetProperty("productState", (value))
}

// GetproductState gets the value of productState for the instance
func (instance *FirewallProduct) GetPropertyproductState() (value uint32, err error) {
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
func (instance *FirewallProduct) SetPropertytimestamp(value string) (err error) {
	return instance.SetProperty("timestamp", (value))
}

// Gettimestamp gets the value of timestamp for the instance
func (instance *FirewallProduct) GetPropertytimestamp() (value string, err error) {
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
