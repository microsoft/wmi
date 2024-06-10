// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.ServiceModel
//
// ////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ServiceToEndpointAssociation struct
type ServiceToEndpointAssociation struct {
	*cim.WmiInstance

	// The endpoint associated with the service.
	Endpoint Endpoint

	// The service associated with the endpoint.
	Service Service
}

func NewServiceToEndpointAssociationEx1(instance *cim.WmiInstance) (newInstance *ServiceToEndpointAssociation, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &ServiceToEndpointAssociation{
		WmiInstance: tmp,
	}
	return
}

func NewServiceToEndpointAssociationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ServiceToEndpointAssociation, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ServiceToEndpointAssociation{
		WmiInstance: tmp,
	}
	return
}

// SetEndpoint sets the value of Endpoint for the instance
func (instance *ServiceToEndpointAssociation) SetPropertyEndpoint(value Endpoint) (err error) {
	return instance.SetProperty("Endpoint", (value))
}

// GetEndpoint gets the value of Endpoint for the instance
func (instance *ServiceToEndpointAssociation) GetPropertyEndpoint() (value Endpoint, err error) {
	retValue, err := instance.GetProperty("Endpoint")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(Endpoint)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " Endpoint is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = Endpoint(valuetmp)

	return
}

// SetService sets the value of Service for the instance
func (instance *ServiceToEndpointAssociation) SetPropertyService(value Service) (err error) {
	return instance.SetProperty("Service", (value))
}

// GetService gets the value of Service for the instance
func (instance *ServiceToEndpointAssociation) GetPropertyService() (value Service, err error) {
	retValue, err := instance.GetProperty("Service")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(Service)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " Service is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = Service(valuetmp)

	return
}
