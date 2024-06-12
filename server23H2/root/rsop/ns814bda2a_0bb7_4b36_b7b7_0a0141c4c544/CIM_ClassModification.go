// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS814BDA2A_0BB7_4B36_B7B7_0A0141C4C544
//////////////////////////////////////////////
package ns814bda2a_0bb7_4b36_b7b7_0a0141c4c544

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_ClassModification struct
type CIM_ClassModification struct {
	*CIM_ClassIndication

	// A copy of the 'previous' class definition whose change generated the Indication. PreviousClassDefinition contains an 'older' copy of the class' information, as compared to what is found in the ClassDefinition property (inherited from ClassIndication).
	PreviousClassDefinition interface{}
}

func NewCIM_ClassModificationEx1(instance *cim.WmiInstance) (newInstance *CIM_ClassModification, err error) {
	tmp, err := NewCIM_ClassIndicationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_ClassModification{
		CIM_ClassIndication: tmp,
	}
	return
}

func NewCIM_ClassModificationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ClassModification, err error) {
	tmp, err := NewCIM_ClassIndicationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ClassModification{
		CIM_ClassIndication: tmp,
	}
	return
}

// SetPreviousClassDefinition sets the value of PreviousClassDefinition for the instance
func (instance *CIM_ClassModification) SetPropertyPreviousClassDefinition(value interface{}) (err error) {
	return instance.SetProperty("PreviousClassDefinition", (value))
}

// GetPreviousClassDefinition gets the value of PreviousClassDefinition for the instance
func (instance *CIM_ClassModification) GetPropertyPreviousClassDefinition() (value interface{}, err error) {
	retValue, err := instance.GetProperty("PreviousClassDefinition")
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
