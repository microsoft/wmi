// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ActivityTransfer struct
type ActivityTransfer struct {
	*WSAT_TraceEvent

	// Activity ID
	ActivityID interface{}

	// Related Activity ID
	RelatedActivityID interface{}
}

func NewActivityTransferEx1(instance *cim.WmiInstance) (newInstance *ActivityTransfer, err error) {
	tmp, err := NewWSAT_TraceEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ActivityTransfer{
		WSAT_TraceEvent: tmp,
	}
	return
}

func NewActivityTransferEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ActivityTransfer, err error) {
	tmp, err := NewWSAT_TraceEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ActivityTransfer{
		WSAT_TraceEvent: tmp,
	}
	return
}

// SetActivityID sets the value of ActivityID for the instance
func (instance *ActivityTransfer) SetPropertyActivityID(value interface{}) (err error) {
	return instance.SetProperty("ActivityID", (value))
}

// GetActivityID gets the value of ActivityID for the instance
func (instance *ActivityTransfer) GetPropertyActivityID() (value interface{}, err error) {
	retValue, err := instance.GetProperty("ActivityID")
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

// SetRelatedActivityID sets the value of RelatedActivityID for the instance
func (instance *ActivityTransfer) SetPropertyRelatedActivityID(value interface{}) (err error) {
	return instance.SetProperty("RelatedActivityID", (value))
}

// GetRelatedActivityID gets the value of RelatedActivityID for the instance
func (instance *ActivityTransfer) GetPropertyRelatedActivityID() (value interface{}, err error) {
	retValue, err := instance.GetProperty("RelatedActivityID")
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
