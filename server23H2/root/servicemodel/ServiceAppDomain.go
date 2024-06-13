// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ServiceAppDomain struct
type ServiceAppDomain struct {
	*cim.WmiInstance

	// Contains properties of the appdomain.
	AppDomainInfo AppDomainInfo

	// The service of this appdomain.
	Service Service
}

func NewServiceAppDomainEx1(instance *cim.WmiInstance) (newInstance *ServiceAppDomain, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &ServiceAppDomain{
		WmiInstance: tmp,
	}
	return
}

func NewServiceAppDomainEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ServiceAppDomain, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ServiceAppDomain{
		WmiInstance: tmp,
	}
	return
}

// SetAppDomainInfo sets the value of AppDomainInfo for the instance
func (instance *ServiceAppDomain) SetPropertyAppDomainInfo(value AppDomainInfo) (err error) {
	return instance.SetProperty("AppDomainInfo", (value))
}

// GetAppDomainInfo gets the value of AppDomainInfo for the instance
func (instance *ServiceAppDomain) GetPropertyAppDomainInfo() (value AppDomainInfo, err error) {
	retValue, err := instance.GetProperty("AppDomainInfo")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(AppDomainInfo)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " AppDomainInfo is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = AppDomainInfo(valuetmp)

	return
}

// SetService sets the value of Service for the instance
func (instance *ServiceAppDomain) SetPropertyService(value Service) (err error) {
	return instance.SetProperty("Service", (value))
}

// GetService gets the value of Service for the instance
func (instance *ServiceAppDomain) GetPropertyService() (value Service, err error) {
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
