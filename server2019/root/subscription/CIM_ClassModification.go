// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.subscription
//////////////////////////////////////////////
package subscription

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
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
	return instance.SetProperty("PreviousClassDefinition", value)
}

// GetPreviousClassDefinition gets the value of PreviousClassDefinition for the instance
func (instance *CIM_ClassModification) GetPropertyPreviousClassDefinition() (value interface{}, err error) {
	retValue, err := instance.GetProperty("PreviousClassDefinition")
	if err != nil {
		return
	}
	value, ok := retValue.(interface{})
	if !ok {
		// TODO: Set an error
	}
	return
}
