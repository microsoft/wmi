// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_Indication struct
type CIM_Indication struct {
	*cim.WmiInstance

	//
	CorrelatedIndications []string

	//
	IndicationFilterName string

	//
	IndicationIdentifier string

	//
	IndicationTime string

	//
	OtherSeverity string

	//
	PerceivedSeverity Indication_PerceivedSeverity

	//
	SequenceContext string

	//
	SequenceNumber int64
}

func NewCIM_IndicationEx1(instance *cim.WmiInstance) (newInstance *CIM_Indication, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &CIM_Indication{
		WmiInstance: tmp,
	}
	return
}

func NewCIM_IndicationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_Indication, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_Indication{
		WmiInstance: tmp,
	}
	return
}

// SetCorrelatedIndications sets the value of CorrelatedIndications for the instance
func (instance *CIM_Indication) SetPropertyCorrelatedIndications(value []string) (err error) {
	return instance.SetProperty("CorrelatedIndications", value)
}

// GetCorrelatedIndications gets the value of CorrelatedIndications for the instance
func (instance *CIM_Indication) GetPropertyCorrelatedIndications() (value []string, err error) {
	retValue, err := instance.GetProperty("CorrelatedIndications")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIndicationFilterName sets the value of IndicationFilterName for the instance
func (instance *CIM_Indication) SetPropertyIndicationFilterName(value string) (err error) {
	return instance.SetProperty("IndicationFilterName", value)
}

// GetIndicationFilterName gets the value of IndicationFilterName for the instance
func (instance *CIM_Indication) GetPropertyIndicationFilterName() (value string, err error) {
	retValue, err := instance.GetProperty("IndicationFilterName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIndicationIdentifier sets the value of IndicationIdentifier for the instance
func (instance *CIM_Indication) SetPropertyIndicationIdentifier(value string) (err error) {
	return instance.SetProperty("IndicationIdentifier", value)
}

// GetIndicationIdentifier gets the value of IndicationIdentifier for the instance
func (instance *CIM_Indication) GetPropertyIndicationIdentifier() (value string, err error) {
	retValue, err := instance.GetProperty("IndicationIdentifier")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIndicationTime sets the value of IndicationTime for the instance
func (instance *CIM_Indication) SetPropertyIndicationTime(value string) (err error) {
	return instance.SetProperty("IndicationTime", value)
}

// GetIndicationTime gets the value of IndicationTime for the instance
func (instance *CIM_Indication) GetPropertyIndicationTime() (value string, err error) {
	retValue, err := instance.GetProperty("IndicationTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherSeverity sets the value of OtherSeverity for the instance
func (instance *CIM_Indication) SetPropertyOtherSeverity(value string) (err error) {
	return instance.SetProperty("OtherSeverity", value)
}

// GetOtherSeverity gets the value of OtherSeverity for the instance
func (instance *CIM_Indication) GetPropertyOtherSeverity() (value string, err error) {
	retValue, err := instance.GetProperty("OtherSeverity")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPerceivedSeverity sets the value of PerceivedSeverity for the instance
func (instance *CIM_Indication) SetPropertyPerceivedSeverity(value Indication_PerceivedSeverity) (err error) {
	return instance.SetProperty("PerceivedSeverity", value)
}

// GetPerceivedSeverity gets the value of PerceivedSeverity for the instance
func (instance *CIM_Indication) GetPropertyPerceivedSeverity() (value Indication_PerceivedSeverity, err error) {
	retValue, err := instance.GetProperty("PerceivedSeverity")
	if err != nil {
		return
	}
	value, ok := retValue.(Indication_PerceivedSeverity)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSequenceContext sets the value of SequenceContext for the instance
func (instance *CIM_Indication) SetPropertySequenceContext(value string) (err error) {
	return instance.SetProperty("SequenceContext", value)
}

// GetSequenceContext gets the value of SequenceContext for the instance
func (instance *CIM_Indication) GetPropertySequenceContext() (value string, err error) {
	retValue, err := instance.GetProperty("SequenceContext")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSequenceNumber sets the value of SequenceNumber for the instance
func (instance *CIM_Indication) SetPropertySequenceNumber(value int64) (err error) {
	return instance.SetProperty("SequenceNumber", value)
}

// GetSequenceNumber gets the value of SequenceNumber for the instance
func (instance *CIM_Indication) GetPropertySequenceNumber() (value int64, err error) {
	retValue, err := instance.GetProperty("SequenceNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(int64)
	if !ok {
		// TODO: Set an error
	}
	return
}
