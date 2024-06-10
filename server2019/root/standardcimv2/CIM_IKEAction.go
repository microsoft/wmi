// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.StandardCimv2
//
// ////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_IKEAction struct
type CIM_IKEAction struct {
	*CIM_SANegotiationAction

	//
	AggressiveModeGroupID uint16

	//
	ExchangeMode uint16

	//
	UseIKEIdentityType uint16

	//
	VendorID string
}

func NewCIM_IKEActionEx1(instance *cim.WmiInstance) (newInstance *CIM_IKEAction, err error) {
	tmp, err := NewCIM_SANegotiationActionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_IKEAction{
		CIM_SANegotiationAction: tmp,
	}
	return
}

func NewCIM_IKEActionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_IKEAction, err error) {
	tmp, err := NewCIM_SANegotiationActionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_IKEAction{
		CIM_SANegotiationAction: tmp,
	}
	return
}

// SetAggressiveModeGroupID sets the value of AggressiveModeGroupID for the instance
func (instance *CIM_IKEAction) SetPropertyAggressiveModeGroupID(value uint16) (err error) {
	return instance.SetProperty("AggressiveModeGroupID", (value))
}

// GetAggressiveModeGroupID gets the value of AggressiveModeGroupID for the instance
func (instance *CIM_IKEAction) GetPropertyAggressiveModeGroupID() (value uint16, err error) {
	retValue, err := instance.GetProperty("AggressiveModeGroupID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetExchangeMode sets the value of ExchangeMode for the instance
func (instance *CIM_IKEAction) SetPropertyExchangeMode(value uint16) (err error) {
	return instance.SetProperty("ExchangeMode", (value))
}

// GetExchangeMode gets the value of ExchangeMode for the instance
func (instance *CIM_IKEAction) GetPropertyExchangeMode() (value uint16, err error) {
	retValue, err := instance.GetProperty("ExchangeMode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetUseIKEIdentityType sets the value of UseIKEIdentityType for the instance
func (instance *CIM_IKEAction) SetPropertyUseIKEIdentityType(value uint16) (err error) {
	return instance.SetProperty("UseIKEIdentityType", (value))
}

// GetUseIKEIdentityType gets the value of UseIKEIdentityType for the instance
func (instance *CIM_IKEAction) GetPropertyUseIKEIdentityType() (value uint16, err error) {
	retValue, err := instance.GetProperty("UseIKEIdentityType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetVendorID sets the value of VendorID for the instance
func (instance *CIM_IKEAction) SetPropertyVendorID(value string) (err error) {
	return instance.SetProperty("VendorID", (value))
}

// GetVendorID gets the value of VendorID for the instance
func (instance *CIM_IKEAction) GetPropertyVendorID() (value string, err error) {
	retValue, err := instance.GetProperty("VendorID")
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
