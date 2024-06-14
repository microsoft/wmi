// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// SegmentElement struct
type SegmentElement struct {
	*CollectionElement

	//
	Segment string
}

func NewSegmentElementEx1(instance *cim.WmiInstance) (newInstance *SegmentElement, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SegmentElement{
		CollectionElement: tmp,
	}
	return
}

func NewSegmentElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SegmentElement, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SegmentElement{
		CollectionElement: tmp,
	}
	return
}

// SetSegment sets the value of Segment for the instance
func (instance *SegmentElement) SetPropertySegment(value string) (err error) {
	return instance.SetProperty("Segment", (value))
}

// GetSegment gets the value of Segment for the instance
func (instance *SegmentElement) GetPropertySegment() (value string, err error) {
	retValue, err := instance.GetProperty("Segment")
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
