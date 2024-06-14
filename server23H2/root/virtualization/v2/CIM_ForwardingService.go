// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_ForwardingService struct
type CIM_ForwardingService struct {
	*CIM_NetworkService

	// This defines the type of protocol that is being forwarded when the value of the ProtocolType attribute is 1 (i.e., "Other"). This provides for future extensibility.
	OtherProtocolType string

	// This defines the type of protocol that is being forwarded.
	ProtocolType ForwardingService_ProtocolType
}

func NewCIM_ForwardingServiceEx1(instance *cim.WmiInstance) (newInstance *CIM_ForwardingService, err error) {
	tmp, err := NewCIM_NetworkServiceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_ForwardingService{
		CIM_NetworkService: tmp,
	}
	return
}

func NewCIM_ForwardingServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ForwardingService, err error) {
	tmp, err := NewCIM_NetworkServiceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ForwardingService{
		CIM_NetworkService: tmp,
	}
	return
}

// SetOtherProtocolType sets the value of OtherProtocolType for the instance
func (instance *CIM_ForwardingService) SetPropertyOtherProtocolType(value string) (err error) {
	return instance.SetProperty("OtherProtocolType", (value))
}

// GetOtherProtocolType gets the value of OtherProtocolType for the instance
func (instance *CIM_ForwardingService) GetPropertyOtherProtocolType() (value string, err error) {
	retValue, err := instance.GetProperty("OtherProtocolType")
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

// SetProtocolType sets the value of ProtocolType for the instance
func (instance *CIM_ForwardingService) SetPropertyProtocolType(value ForwardingService_ProtocolType) (err error) {
	return instance.SetProperty("ProtocolType", (value))
}

// GetProtocolType gets the value of ProtocolType for the instance
func (instance *CIM_ForwardingService) GetPropertyProtocolType() (value ForwardingService_ProtocolType, err error) {
	retValue, err := instance.GetProperty("ProtocolType")
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

	value = ForwardingService_ProtocolType(valuetmp)

	return
}
