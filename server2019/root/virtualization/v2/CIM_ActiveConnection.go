// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_ActiveConnection struct
type CIM_ActiveConnection struct {
	*CIM_SAPSAPDependency

	// TRUE means that this connection is unidirectional; FALSE means that this connection is bidirectional. When the connection is unidirectional, the "speaker" should be defined as the Antecedent reference. In a bidirectional connection, the selection of which AccessPoint is the Antecedent or Dependent is immaterial.
	IsUnidirectional bool

	// Note: The use of this element is deprecated because it is incorrectly placed on the association. Unicast, broadcast, or other traffic types are not a function of the connection between the referenced endpoints, but rather are a function of the addressing, protocol and basic functionality of the endpoints.
	///Deprecated description: A string that describes the type of traffic that is being carried over this instance when its Type property is set, for example, to 1 (Other).
	OtherTrafficDescription string

	// Note: The use of this element is deprecated because it is incorrectly placed on the association. Unicast, broadcast, or other traffic types are not a function of the connection between the referenced endpoints, but rather are a function of the addressing, protocol and basic functionality of the endpoints.
	///Deprecated description: The type of traffic that is carried over this connection.
	TrafficType ActiveConnection_TrafficType
}

func NewCIM_ActiveConnectionEx1(instance *cim.WmiInstance) (newInstance *CIM_ActiveConnection, err error) {
	tmp, err := NewCIM_SAPSAPDependencyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_ActiveConnection{
		CIM_SAPSAPDependency: tmp,
	}
	return
}

func NewCIM_ActiveConnectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ActiveConnection, err error) {
	tmp, err := NewCIM_SAPSAPDependencyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ActiveConnection{
		CIM_SAPSAPDependency: tmp,
	}
	return
}

// SetIsUnidirectional sets the value of IsUnidirectional for the instance
func (instance *CIM_ActiveConnection) SetPropertyIsUnidirectional(value bool) (err error) {
	return instance.SetProperty("IsUnidirectional", (value))
}

// GetIsUnidirectional gets the value of IsUnidirectional for the instance
func (instance *CIM_ActiveConnection) GetPropertyIsUnidirectional() (value bool, err error) {
	retValue, err := instance.GetProperty("IsUnidirectional")
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

// SetOtherTrafficDescription sets the value of OtherTrafficDescription for the instance
func (instance *CIM_ActiveConnection) SetPropertyOtherTrafficDescription(value string) (err error) {
	return instance.SetProperty("OtherTrafficDescription", (value))
}

// GetOtherTrafficDescription gets the value of OtherTrafficDescription for the instance
func (instance *CIM_ActiveConnection) GetPropertyOtherTrafficDescription() (value string, err error) {
	retValue, err := instance.GetProperty("OtherTrafficDescription")
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

// SetTrafficType sets the value of TrafficType for the instance
func (instance *CIM_ActiveConnection) SetPropertyTrafficType(value ActiveConnection_TrafficType) (err error) {
	return instance.SetProperty("TrafficType", (value))
}

// GetTrafficType gets the value of TrafficType for the instance
func (instance *CIM_ActiveConnection) GetPropertyTrafficType() (value ActiveConnection_TrafficType, err error) {
	retValue, err := instance.GetProperty("TrafficType")
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

	value = ActiveConnection_TrafficType(valuetmp)

	return
}
