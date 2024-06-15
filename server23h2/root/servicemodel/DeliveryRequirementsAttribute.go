// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// DeliveryRequirementsAttribute struct
type DeliveryRequirementsAttribute struct {
	*Behavior

	// Specifies whether the binding for a service supports contracts.
	QueuedDeliveryRequirements string

	// Specifies whether the binding supports ordered messages.
	RequireOrderedDelivery bool

	// The contract to which it applies.
	TargetContract string
}

func NewDeliveryRequirementsAttributeEx1(instance *cim.WmiInstance) (newInstance *DeliveryRequirementsAttribute, err error) {
	tmp, err := NewBehaviorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &DeliveryRequirementsAttribute{
		Behavior: tmp,
	}
	return
}

func NewDeliveryRequirementsAttributeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DeliveryRequirementsAttribute, err error) {
	tmp, err := NewBehaviorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DeliveryRequirementsAttribute{
		Behavior: tmp,
	}
	return
}

// SetQueuedDeliveryRequirements sets the value of QueuedDeliveryRequirements for the instance
func (instance *DeliveryRequirementsAttribute) SetPropertyQueuedDeliveryRequirements(value string) (err error) {
	return instance.SetProperty("QueuedDeliveryRequirements", (value))
}

// GetQueuedDeliveryRequirements gets the value of QueuedDeliveryRequirements for the instance
func (instance *DeliveryRequirementsAttribute) GetPropertyQueuedDeliveryRequirements() (value string, err error) {
	retValue, err := instance.GetProperty("QueuedDeliveryRequirements")
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

// SetRequireOrderedDelivery sets the value of RequireOrderedDelivery for the instance
func (instance *DeliveryRequirementsAttribute) SetPropertyRequireOrderedDelivery(value bool) (err error) {
	return instance.SetProperty("RequireOrderedDelivery", (value))
}

// GetRequireOrderedDelivery gets the value of RequireOrderedDelivery for the instance
func (instance *DeliveryRequirementsAttribute) GetPropertyRequireOrderedDelivery() (value bool, err error) {
	retValue, err := instance.GetProperty("RequireOrderedDelivery")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetTargetContract sets the value of TargetContract for the instance
func (instance *DeliveryRequirementsAttribute) SetPropertyTargetContract(value string) (err error) {
	return instance.SetProperty("TargetContract", (value))
}

// GetTargetContract gets the value of TargetContract for the instance
func (instance *DeliveryRequirementsAttribute) GetPropertyTargetContract() (value string, err error) {
	retValue, err := instance.GetProperty("TargetContract")
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
