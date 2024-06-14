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

// SampledProfileInterval_V2 struct
type SampledProfileInterval_V2 struct {
	*PerfInfo_V2

	//
	NewInterval uint32

	//
	OldInterval uint32

	//
	Source uint32
}

func NewSampledProfileInterval_V2Ex1(instance *cim.WmiInstance) (newInstance *SampledProfileInterval_V2, err error) {
	tmp, err := NewPerfInfo_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &SampledProfileInterval_V2{
		PerfInfo_V2: tmp,
	}
	return
}

func NewSampledProfileInterval_V2Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SampledProfileInterval_V2, err error) {
	tmp, err := NewPerfInfo_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SampledProfileInterval_V2{
		PerfInfo_V2: tmp,
	}
	return
}

// SetNewInterval sets the value of NewInterval for the instance
func (instance *SampledProfileInterval_V2) SetPropertyNewInterval(value uint32) (err error) {
	return instance.SetProperty("NewInterval", (value))
}

// GetNewInterval gets the value of NewInterval for the instance
func (instance *SampledProfileInterval_V2) GetPropertyNewInterval() (value uint32, err error) {
	retValue, err := instance.GetProperty("NewInterval")
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

// SetOldInterval sets the value of OldInterval for the instance
func (instance *SampledProfileInterval_V2) SetPropertyOldInterval(value uint32) (err error) {
	return instance.SetProperty("OldInterval", (value))
}

// GetOldInterval gets the value of OldInterval for the instance
func (instance *SampledProfileInterval_V2) GetPropertyOldInterval() (value uint32, err error) {
	retValue, err := instance.GetProperty("OldInterval")
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

// SetSource sets the value of Source for the instance
func (instance *SampledProfileInterval_V2) SetPropertySource(value uint32) (err error) {
	return instance.SetProperty("Source", (value))
}

// GetSource gets the value of Source for the instance
func (instance *SampledProfileInterval_V2) GetPropertySource() (value uint32, err error) {
	retValue, err := instance.GetProperty("Source")
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
