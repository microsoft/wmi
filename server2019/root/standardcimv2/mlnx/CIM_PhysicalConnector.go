// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_PhysicalConnector struct
type CIM_PhysicalConnector struct {
	*CIM_PhysicalElement

	//
	ConnectorDescription string

	//
	ConnectorElectricalCharacteristics []PhysicalConnector_ConnectorElectricalCharacteristics

	//
	ConnectorGender PhysicalConnector_ConnectorGender

	//
	ConnectorLayout PhysicalConnector_ConnectorLayout

	//
	ConnectorPinout string

	//
	ConnectorType []PhysicalConnector_ConnectorType

	//
	NumPhysicalPins uint32

	//
	OtherElectricalCharacteristics []string

	//
	OtherTypeDescription string
}

func NewCIM_PhysicalConnectorEx1(instance *cim.WmiInstance) (newInstance *CIM_PhysicalConnector, err error) {
	tmp, err := NewCIM_PhysicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_PhysicalConnector{
		CIM_PhysicalElement: tmp,
	}
	return
}

func NewCIM_PhysicalConnectorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_PhysicalConnector, err error) {
	tmp, err := NewCIM_PhysicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_PhysicalConnector{
		CIM_PhysicalElement: tmp,
	}
	return
}

// SetConnectorDescription sets the value of ConnectorDescription for the instance
func (instance *CIM_PhysicalConnector) SetPropertyConnectorDescription(value string) (err error) {
	return instance.SetProperty("ConnectorDescription", (value))
}

// GetConnectorDescription gets the value of ConnectorDescription for the instance
func (instance *CIM_PhysicalConnector) GetPropertyConnectorDescription() (value string, err error) {
	retValue, err := instance.GetProperty("ConnectorDescription")
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

// SetConnectorElectricalCharacteristics sets the value of ConnectorElectricalCharacteristics for the instance
func (instance *CIM_PhysicalConnector) SetPropertyConnectorElectricalCharacteristics(value []PhysicalConnector_ConnectorElectricalCharacteristics) (err error) {
	return instance.SetProperty("ConnectorElectricalCharacteristics", (value))
}

// GetConnectorElectricalCharacteristics gets the value of ConnectorElectricalCharacteristics for the instance
func (instance *CIM_PhysicalConnector) GetPropertyConnectorElectricalCharacteristics() (value []PhysicalConnector_ConnectorElectricalCharacteristics, err error) {
	retValue, err := instance.GetProperty("ConnectorElectricalCharacteristics")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, PhysicalConnector_ConnectorElectricalCharacteristics(valuetmp))
	}

	return
}

// SetConnectorGender sets the value of ConnectorGender for the instance
func (instance *CIM_PhysicalConnector) SetPropertyConnectorGender(value PhysicalConnector_ConnectorGender) (err error) {
	return instance.SetProperty("ConnectorGender", (value))
}

// GetConnectorGender gets the value of ConnectorGender for the instance
func (instance *CIM_PhysicalConnector) GetPropertyConnectorGender() (value PhysicalConnector_ConnectorGender, err error) {
	retValue, err := instance.GetProperty("ConnectorGender")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = PhysicalConnector_ConnectorGender(valuetmp)

	return
}

// SetConnectorLayout sets the value of ConnectorLayout for the instance
func (instance *CIM_PhysicalConnector) SetPropertyConnectorLayout(value PhysicalConnector_ConnectorLayout) (err error) {
	return instance.SetProperty("ConnectorLayout", (value))
}

// GetConnectorLayout gets the value of ConnectorLayout for the instance
func (instance *CIM_PhysicalConnector) GetPropertyConnectorLayout() (value PhysicalConnector_ConnectorLayout, err error) {
	retValue, err := instance.GetProperty("ConnectorLayout")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = PhysicalConnector_ConnectorLayout(valuetmp)

	return
}

// SetConnectorPinout sets the value of ConnectorPinout for the instance
func (instance *CIM_PhysicalConnector) SetPropertyConnectorPinout(value string) (err error) {
	return instance.SetProperty("ConnectorPinout", (value))
}

// GetConnectorPinout gets the value of ConnectorPinout for the instance
func (instance *CIM_PhysicalConnector) GetPropertyConnectorPinout() (value string, err error) {
	retValue, err := instance.GetProperty("ConnectorPinout")
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

// SetConnectorType sets the value of ConnectorType for the instance
func (instance *CIM_PhysicalConnector) SetPropertyConnectorType(value []PhysicalConnector_ConnectorType) (err error) {
	return instance.SetProperty("ConnectorType", (value))
}

// GetConnectorType gets the value of ConnectorType for the instance
func (instance *CIM_PhysicalConnector) GetPropertyConnectorType() (value []PhysicalConnector_ConnectorType, err error) {
	retValue, err := instance.GetProperty("ConnectorType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, PhysicalConnector_ConnectorType(valuetmp))
	}

	return
}

// SetNumPhysicalPins sets the value of NumPhysicalPins for the instance
func (instance *CIM_PhysicalConnector) SetPropertyNumPhysicalPins(value uint32) (err error) {
	return instance.SetProperty("NumPhysicalPins", (value))
}

// GetNumPhysicalPins gets the value of NumPhysicalPins for the instance
func (instance *CIM_PhysicalConnector) GetPropertyNumPhysicalPins() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumPhysicalPins")
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

// SetOtherElectricalCharacteristics sets the value of OtherElectricalCharacteristics for the instance
func (instance *CIM_PhysicalConnector) SetPropertyOtherElectricalCharacteristics(value []string) (err error) {
	return instance.SetProperty("OtherElectricalCharacteristics", (value))
}

// GetOtherElectricalCharacteristics gets the value of OtherElectricalCharacteristics for the instance
func (instance *CIM_PhysicalConnector) GetPropertyOtherElectricalCharacteristics() (value []string, err error) {
	retValue, err := instance.GetProperty("OtherElectricalCharacteristics")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetOtherTypeDescription sets the value of OtherTypeDescription for the instance
func (instance *CIM_PhysicalConnector) SetPropertyOtherTypeDescription(value string) (err error) {
	return instance.SetProperty("OtherTypeDescription", (value))
}

// GetOtherTypeDescription gets the value of OtherTypeDescription for the instance
func (instance *CIM_PhysicalConnector) GetPropertyOtherTypeDescription() (value string, err error) {
	retValue, err := instance.GetProperty("OtherTypeDescription")
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
