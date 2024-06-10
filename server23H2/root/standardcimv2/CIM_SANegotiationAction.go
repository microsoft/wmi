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

// CIM_SANegotiationAction struct
type CIM_SANegotiationAction struct {
	*CIM_SAAction

	//
	IdleDurationSeconds uint64

	//
	MinLifetimeKilobytes uint64

	//
	MinLifetimeSeconds uint64
}

func NewCIM_SANegotiationActionEx1(instance *cim.WmiInstance) (newInstance *CIM_SANegotiationAction, err error) {
	tmp, err := NewCIM_SAActionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_SANegotiationAction{
		CIM_SAAction: tmp,
	}
	return
}

func NewCIM_SANegotiationActionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_SANegotiationAction, err error) {
	tmp, err := NewCIM_SAActionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_SANegotiationAction{
		CIM_SAAction: tmp,
	}
	return
}

// SetIdleDurationSeconds sets the value of IdleDurationSeconds for the instance
func (instance *CIM_SANegotiationAction) SetPropertyIdleDurationSeconds(value uint64) (err error) {
	return instance.SetProperty("IdleDurationSeconds", (value))
}

// GetIdleDurationSeconds gets the value of IdleDurationSeconds for the instance
func (instance *CIM_SANegotiationAction) GetPropertyIdleDurationSeconds() (value uint64, err error) {
	retValue, err := instance.GetProperty("IdleDurationSeconds")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetMinLifetimeKilobytes sets the value of MinLifetimeKilobytes for the instance
func (instance *CIM_SANegotiationAction) SetPropertyMinLifetimeKilobytes(value uint64) (err error) {
	return instance.SetProperty("MinLifetimeKilobytes", (value))
}

// GetMinLifetimeKilobytes gets the value of MinLifetimeKilobytes for the instance
func (instance *CIM_SANegotiationAction) GetPropertyMinLifetimeKilobytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("MinLifetimeKilobytes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetMinLifetimeSeconds sets the value of MinLifetimeSeconds for the instance
func (instance *CIM_SANegotiationAction) SetPropertyMinLifetimeSeconds(value uint64) (err error) {
	return instance.SetProperty("MinLifetimeSeconds", (value))
}

// GetMinLifetimeSeconds gets the value of MinLifetimeSeconds for the instance
func (instance *CIM_SANegotiationAction) GetPropertyMinLifetimeSeconds() (value uint64, err error) {
	retValue, err := instance.GetProperty("MinLifetimeSeconds")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}
