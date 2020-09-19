// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_ServiceAvailableToElement struct
type CIM_ServiceAvailableToElement struct {
	*cim.WmiInstance

	//
	ServiceProvided CIM_Service

	//
	UserOfService CIM_ManagedElement
}

func NewCIM_ServiceAvailableToElementEx1(instance *cim.WmiInstance) (newInstance *CIM_ServiceAvailableToElement, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &CIM_ServiceAvailableToElement{
		WmiInstance: tmp,
	}
	return
}

func NewCIM_ServiceAvailableToElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ServiceAvailableToElement, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ServiceAvailableToElement{
		WmiInstance: tmp,
	}
	return
}

// SetServiceProvided sets the value of ServiceProvided for the instance
func (instance *CIM_ServiceAvailableToElement) SetPropertyServiceProvided(value CIM_Service) (err error) {
	return instance.SetProperty("ServiceProvided", (value))
}

// GetServiceProvided gets the value of ServiceProvided for the instance
func (instance *CIM_ServiceAvailableToElement) GetPropertyServiceProvided() (value CIM_Service, err error) {
	retValue, err := instance.GetProperty("ServiceProvided")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(CIM_Service)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " CIM_Service is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = CIM_Service(valuetmp)

	return
}

// SetUserOfService sets the value of UserOfService for the instance
func (instance *CIM_ServiceAvailableToElement) SetPropertyUserOfService(value CIM_ManagedElement) (err error) {
	return instance.SetProperty("UserOfService", (value))
}

// GetUserOfService gets the value of UserOfService for the instance
func (instance *CIM_ServiceAvailableToElement) GetPropertyUserOfService() (value CIM_ManagedElement, err error) {
	retValue, err := instance.GetProperty("UserOfService")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(CIM_ManagedElement)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " CIM_ManagedElement is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = CIM_ManagedElement(valuetmp)

	return
}
