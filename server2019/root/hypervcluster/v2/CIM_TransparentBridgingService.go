// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.HyperVCluster.v2
//
// ////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_TransparentBridgingService struct
type CIM_TransparentBridgingService struct {
	*CIM_ForwardingService

	// The timeout period in seconds for aging out dynamically learned forwarding information. 802.1D-1990 recommends a default of 300 seconds.
	AgingTime uint32

	// Filtering Database Identifier used by VLAN-aware switches that have more than one filtering database.
	FID uint32
}

func NewCIM_TransparentBridgingServiceEx1(instance *cim.WmiInstance) (newInstance *CIM_TransparentBridgingService, err error) {
	tmp, err := NewCIM_ForwardingServiceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_TransparentBridgingService{
		CIM_ForwardingService: tmp,
	}
	return
}

func NewCIM_TransparentBridgingServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_TransparentBridgingService, err error) {
	tmp, err := NewCIM_ForwardingServiceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_TransparentBridgingService{
		CIM_ForwardingService: tmp,
	}
	return
}

// SetAgingTime sets the value of AgingTime for the instance
func (instance *CIM_TransparentBridgingService) SetPropertyAgingTime(value uint32) (err error) {
	return instance.SetProperty("AgingTime", (value))
}

// GetAgingTime gets the value of AgingTime for the instance
func (instance *CIM_TransparentBridgingService) GetPropertyAgingTime() (value uint32, err error) {
	retValue, err := instance.GetProperty("AgingTime")
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

// SetFID sets the value of FID for the instance
func (instance *CIM_TransparentBridgingService) SetPropertyFID(value uint32) (err error) {
	return instance.SetProperty("FID", (value))
}

// GetFID gets the value of FID for the instance
func (instance *CIM_TransparentBridgingService) GetPropertyFID() (value uint32, err error) {
	retValue, err := instance.GetProperty("FID")
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
