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

// CIM_ElementConformsToProfile struct
type CIM_ElementConformsToProfile struct {
	*cim.WmiInstance

	//
	ConformantStandard CIM_RegisteredProfile

	//
	ManagedElement CIM_ManagedElement
}

func NewCIM_ElementConformsToProfileEx1(instance *cim.WmiInstance) (newInstance *CIM_ElementConformsToProfile, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &CIM_ElementConformsToProfile{
		WmiInstance: tmp,
	}
	return
}

func NewCIM_ElementConformsToProfileEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ElementConformsToProfile, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ElementConformsToProfile{
		WmiInstance: tmp,
	}
	return
}

// SetConformantStandard sets the value of ConformantStandard for the instance
func (instance *CIM_ElementConformsToProfile) SetPropertyConformantStandard(value CIM_RegisteredProfile) (err error) {
	return instance.SetProperty("ConformantStandard", (value))
}

// GetConformantStandard gets the value of ConformantStandard for the instance
func (instance *CIM_ElementConformsToProfile) GetPropertyConformantStandard() (value CIM_RegisteredProfile, err error) {
	retValue, err := instance.GetProperty("ConformantStandard")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(CIM_RegisteredProfile)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " CIM_RegisteredProfile is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = CIM_RegisteredProfile(valuetmp)

	return
}

// SetManagedElement sets the value of ManagedElement for the instance
func (instance *CIM_ElementConformsToProfile) SetPropertyManagedElement(value CIM_ManagedElement) (err error) {
	return instance.SetProperty("ManagedElement", (value))
}

// GetManagedElement gets the value of ManagedElement for the instance
func (instance *CIM_ElementConformsToProfile) GetPropertyManagedElement() (value CIM_ManagedElement, err error) {
	retValue, err := instance.GetProperty("ManagedElement")
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
