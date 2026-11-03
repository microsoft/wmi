// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.SecurityCenter
//////////////////////////////////////////////
package securitycenter

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
	companyName string

	//
	displayName string

	//
	enabled bool

	//
	instanceGuid string

	//
	pathToSignedProductExe string

	//
	productHasNotifiedUser bool

	//
	productState uint8

	//
	productWantsWscNotifications bool

	//
	versionNumber string
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

// SetcompanyName sets the value of companyName for the instance
func (instance *FirewallProduct) SetPropertycompanyName(value string) (err error) {
	return instance.SetProperty("companyName", (value))
}

// GetcompanyName gets the value of companyName for the instance
func (instance *FirewallProduct) GetPropertycompanyName() (value string, err error) {
	retValue, err := instance.GetProperty("companyName")
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

// Setenabled sets the value of enabled for the instance
func (instance *FirewallProduct) SetPropertyenabled(value bool) (err error) {
	return instance.SetProperty("enabled", (value))
}

// Getenabled gets the value of enabled for the instance
func (instance *FirewallProduct) GetPropertyenabled() (value bool, err error) {
	retValue, err := instance.GetProperty("enabled")
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

// SetproductHasNotifiedUser sets the value of productHasNotifiedUser for the instance
func (instance *FirewallProduct) SetPropertyproductHasNotifiedUser(value bool) (err error) {
	return instance.SetProperty("productHasNotifiedUser", (value))
}

// GetproductHasNotifiedUser gets the value of productHasNotifiedUser for the instance
func (instance *FirewallProduct) GetPropertyproductHasNotifiedUser() (value bool, err error) {
	retValue, err := instance.GetProperty("productHasNotifiedUser")
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

// SetproductState sets the value of productState for the instance
func (instance *FirewallProduct) SetPropertyproductState(value uint8) (err error) {
	return instance.SetProperty("productState", (value))
}

// GetproductState gets the value of productState for the instance
func (instance *FirewallProduct) GetPropertyproductState() (value uint8, err error) {
	retValue, err := instance.GetProperty("productState")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetproductWantsWscNotifications sets the value of productWantsWscNotifications for the instance
func (instance *FirewallProduct) SetPropertyproductWantsWscNotifications(value bool) (err error) {
	return instance.SetProperty("productWantsWscNotifications", (value))
}

// GetproductWantsWscNotifications gets the value of productWantsWscNotifications for the instance
func (instance *FirewallProduct) GetPropertyproductWantsWscNotifications() (value bool, err error) {
	retValue, err := instance.GetProperty("productWantsWscNotifications")
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

// SetversionNumber sets the value of versionNumber for the instance
func (instance *FirewallProduct) SetPropertyversionNumber(value string) (err error) {
	return instance.SetProperty("versionNumber", (value))
}

// GetversionNumber gets the value of versionNumber for the instance
func (instance *FirewallProduct) GetPropertyversionNumber() (value string, err error) {
	retValue, err := instance.GetProperty("versionNumber")
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
