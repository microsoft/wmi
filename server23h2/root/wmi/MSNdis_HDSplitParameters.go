// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSNdis_HDSplitParameters struct
type MSNdis_HDSplitParameters struct {
	*MSNdis

	//
	HDSplitCombineFlags uint32

	//
	Header MSNdis_ObjectHeader
}

func NewMSNdis_HDSplitParametersEx1(instance *cim.WmiInstance) (newInstance *MSNdis_HDSplitParameters, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_HDSplitParameters{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_HDSplitParametersEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_HDSplitParameters, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_HDSplitParameters{
		MSNdis: tmp,
	}
	return
}

// SetHDSplitCombineFlags sets the value of HDSplitCombineFlags for the instance
func (instance *MSNdis_HDSplitParameters) SetPropertyHDSplitCombineFlags(value uint32) (err error) {
	return instance.SetProperty("HDSplitCombineFlags", (value))
}

// GetHDSplitCombineFlags gets the value of HDSplitCombineFlags for the instance
func (instance *MSNdis_HDSplitParameters) GetPropertyHDSplitCombineFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("HDSplitCombineFlags")
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

// SetHeader sets the value of Header for the instance
func (instance *MSNdis_HDSplitParameters) SetPropertyHeader(value MSNdis_ObjectHeader) (err error) {
	return instance.SetProperty("Header", (value))
}

// GetHeader gets the value of Header for the instance
func (instance *MSNdis_HDSplitParameters) GetPropertyHeader() (value MSNdis_ObjectHeader, err error) {
	retValue, err := instance.GetProperty("Header")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_ObjectHeader)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_ObjectHeader is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_ObjectHeader(valuetmp)

	return
}
