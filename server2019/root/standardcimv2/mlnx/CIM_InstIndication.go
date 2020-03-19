// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_InstIndication struct
type CIM_InstIndication struct {
	*CIM_Indication

	//
	SourceInstance interface{}

	//
	SourceInstanceHost string

	//
	SourceInstanceModelPath string
}

func NewCIM_InstIndicationEx1(instance *cim.WmiInstance) (newInstance *CIM_InstIndication, err error) {
	tmp, err := NewCIM_IndicationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_InstIndication{
		CIM_Indication: tmp,
	}
	return
}

func NewCIM_InstIndicationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_InstIndication, err error) {
	tmp, err := NewCIM_IndicationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_InstIndication{
		CIM_Indication: tmp,
	}
	return
}

// SetSourceInstance sets the value of SourceInstance for the instance
func (instance *CIM_InstIndication) SetPropertySourceInstance(value interface{}) (err error) {
	return instance.SetProperty("SourceInstance", value)
}

// GetSourceInstance gets the value of SourceInstance for the instance
func (instance *CIM_InstIndication) GetPropertySourceInstance() (value interface{}, err error) {
	retValue, err := instance.GetProperty("SourceInstance")
	if err != nil {
		return
	}
	value, ok := retValue.(interface{})
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSourceInstanceHost sets the value of SourceInstanceHost for the instance
func (instance *CIM_InstIndication) SetPropertySourceInstanceHost(value string) (err error) {
	return instance.SetProperty("SourceInstanceHost", value)
}

// GetSourceInstanceHost gets the value of SourceInstanceHost for the instance
func (instance *CIM_InstIndication) GetPropertySourceInstanceHost() (value string, err error) {
	retValue, err := instance.GetProperty("SourceInstanceHost")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSourceInstanceModelPath sets the value of SourceInstanceModelPath for the instance
func (instance *CIM_InstIndication) SetPropertySourceInstanceModelPath(value string) (err error) {
	return instance.SetProperty("SourceInstanceModelPath", value)
}

// GetSourceInstanceModelPath gets the value of SourceInstanceModelPath for the instance
func (instance *CIM_InstIndication) GetPropertySourceInstanceModelPath() (value string, err error) {
	retValue, err := instance.GetProperty("SourceInstanceModelPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
