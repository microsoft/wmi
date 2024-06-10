// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_AvailableDiagnosticService struct
type CIM_AvailableDiagnosticService struct {
	*CIM_ServiceAvailableToElement

	//
	EstimatedDurationOfService AvailableDiagnosticService_EstimatedDurationOfService

	//
	EstimatedDurationQualifier uint32
}

func NewCIM_AvailableDiagnosticServiceEx1(instance *cim.WmiInstance) (newInstance *CIM_AvailableDiagnosticService, err error) {
	tmp, err := NewCIM_ServiceAvailableToElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_AvailableDiagnosticService{
		CIM_ServiceAvailableToElement: tmp,
	}
	return
}

func NewCIM_AvailableDiagnosticServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_AvailableDiagnosticService, err error) {
	tmp, err := NewCIM_ServiceAvailableToElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_AvailableDiagnosticService{
		CIM_ServiceAvailableToElement: tmp,
	}
	return
}

// SetEstimatedDurationOfService sets the value of EstimatedDurationOfService for the instance
func (instance *CIM_AvailableDiagnosticService) SetPropertyEstimatedDurationOfService(value AvailableDiagnosticService_EstimatedDurationOfService) (err error) {
	return instance.SetProperty("EstimatedDurationOfService", (value))
}

// GetEstimatedDurationOfService gets the value of EstimatedDurationOfService for the instance
func (instance *CIM_AvailableDiagnosticService) GetPropertyEstimatedDurationOfService() (value AvailableDiagnosticService_EstimatedDurationOfService, err error) {
	retValue, err := instance.GetProperty("EstimatedDurationOfService")
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

	value = AvailableDiagnosticService_EstimatedDurationOfService(valuetmp)

	return
}

// SetEstimatedDurationQualifier sets the value of EstimatedDurationQualifier for the instance
func (instance *CIM_AvailableDiagnosticService) SetPropertyEstimatedDurationQualifier(value uint32) (err error) {
	return instance.SetProperty("EstimatedDurationQualifier", (value))
}

// GetEstimatedDurationQualifier gets the value of EstimatedDurationQualifier for the instance
func (instance *CIM_AvailableDiagnosticService) GetPropertyEstimatedDurationQualifier() (value uint32, err error) {
	retValue, err := instance.GetProperty("EstimatedDurationQualifier")
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
