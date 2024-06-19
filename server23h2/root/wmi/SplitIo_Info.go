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

// SplitIo_Info struct
type SplitIo_Info struct {
	*SplitIo

	//
	ChildIrp uint32

	//
	ParentIrp uint32
}

func NewSplitIo_InfoEx1(instance *cim.WmiInstance) (newInstance *SplitIo_Info, err error) {
	tmp, err := NewSplitIoEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SplitIo_Info{
		SplitIo: tmp,
	}
	return
}

func NewSplitIo_InfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SplitIo_Info, err error) {
	tmp, err := NewSplitIoEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SplitIo_Info{
		SplitIo: tmp,
	}
	return
}

// SetChildIrp sets the value of ChildIrp for the instance
func (instance *SplitIo_Info) SetPropertyChildIrp(value uint32) (err error) {
	return instance.SetProperty("ChildIrp", (value))
}

// GetChildIrp gets the value of ChildIrp for the instance
func (instance *SplitIo_Info) GetPropertyChildIrp() (value uint32, err error) {
	retValue, err := instance.GetProperty("ChildIrp")
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

// SetParentIrp sets the value of ParentIrp for the instance
func (instance *SplitIo_Info) SetPropertyParentIrp(value uint32) (err error) {
	return instance.SetProperty("ParentIrp", (value))
}

// GetParentIrp gets the value of ParentIrp for the instance
func (instance *SplitIo_Info) GetPropertyParentIrp() (value uint32, err error) {
	retValue, err := instance.GetProperty("ParentIrp")
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
