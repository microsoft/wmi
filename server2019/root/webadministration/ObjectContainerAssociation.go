// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WebAdministration
//
// ////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ObjectContainerAssociation struct
type ObjectContainerAssociation struct {
	*cim.WmiInstance

	//
	Container interface{}

	//
	Element interface{}
}

func NewObjectContainerAssociationEx1(instance *cim.WmiInstance) (newInstance *ObjectContainerAssociation, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &ObjectContainerAssociation{
		WmiInstance: tmp,
	}
	return
}

func NewObjectContainerAssociationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ObjectContainerAssociation, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ObjectContainerAssociation{
		WmiInstance: tmp,
	}
	return
}

// SetContainer sets the value of Container for the instance
func (instance *ObjectContainerAssociation) SetPropertyContainer(value interface{}) (err error) {
	return instance.SetProperty("Container", (value))
}

// GetContainer gets the value of Container for the instance
func (instance *ObjectContainerAssociation) GetPropertyContainer() (value interface{}, err error) {
	retValue, err := instance.GetProperty("Container")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}

// SetElement sets the value of Element for the instance
func (instance *ObjectContainerAssociation) SetPropertyElement(value interface{}) (err error) {
	return instance.SetProperty("Element", (value))
}

// GetElement gets the value of Element for the instance
func (instance *ObjectContainerAssociation) GetPropertyElement() (value interface{}, err error) {
	retValue, err := instance.GetProperty("Element")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}
