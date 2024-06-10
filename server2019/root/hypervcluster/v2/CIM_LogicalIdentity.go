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
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_LogicalIdentity struct
type CIM_LogicalIdentity struct {
	*cim.WmiInstance

	// SameElement represents an alternate aspect of the ManagedElement.
	SameElement CIM_ManagedElement

	// SystemElement represents one aspect of the Managed Element. The use of 'System' in the role name does not limit the scope of the association. The role name was defined in the original association, where the referenced elements were limited to LogicalElements. Since that time, it has been found valuable to instantiate these types of relationships for ManagedElements, such as Collections. So, the referenced elements of the association were redefined to be ManagedElements. Unfortunately, the role name could not be changed without deprecating the entire association. This was not deemed necessary just to correct the role name.
	SystemElement CIM_ManagedElement
}

func NewCIM_LogicalIdentityEx1(instance *cim.WmiInstance) (newInstance *CIM_LogicalIdentity, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &CIM_LogicalIdentity{
		WmiInstance: tmp,
	}
	return
}

func NewCIM_LogicalIdentityEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_LogicalIdentity, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_LogicalIdentity{
		WmiInstance: tmp,
	}
	return
}

// SetSameElement sets the value of SameElement for the instance
func (instance *CIM_LogicalIdentity) SetPropertySameElement(value CIM_ManagedElement) (err error) {
	return instance.SetProperty("SameElement", (value))
}

// GetSameElement gets the value of SameElement for the instance
func (instance *CIM_LogicalIdentity) GetPropertySameElement() (value CIM_ManagedElement, err error) {
	retValue, err := instance.GetProperty("SameElement")
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

// SetSystemElement sets the value of SystemElement for the instance
func (instance *CIM_LogicalIdentity) SetPropertySystemElement(value CIM_ManagedElement) (err error) {
	return instance.SetProperty("SystemElement", (value))
}

// GetSystemElement gets the value of SystemElement for the instance
func (instance *CIM_LogicalIdentity) GetPropertySystemElement() (value CIM_ManagedElement, err error) {
	retValue, err := instance.GetProperty("SystemElement")
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
