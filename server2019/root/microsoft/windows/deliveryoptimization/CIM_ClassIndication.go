// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.DeliveryOptimization
//////////////////////////////////////////////
package deliveryoptimization

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_ClassIndication struct
type CIM_ClassIndication struct {
	*CIM_Indication

	// The current definition of the class that is created, changed or deleted in the schema. In the case of a CIM_ClassDeletion Indication, the definition for the class just prior to deletion should be placed in this property.
	ClassDefinition interface{}
}

func NewCIM_ClassIndicationEx1(instance *cim.WmiInstance) (newInstance *CIM_ClassIndication, err error) {
	tmp, err := NewCIM_IndicationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_ClassIndication{
		CIM_Indication: tmp,
	}
	return
}

func NewCIM_ClassIndicationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ClassIndication, err error) {
	tmp, err := NewCIM_IndicationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ClassIndication{
		CIM_Indication: tmp,
	}
	return
}

// SetClassDefinition sets the value of ClassDefinition for the instance
func (instance *CIM_ClassIndication) SetPropertyClassDefinition(value interface{}) (err error) {
	return instance.SetProperty("ClassDefinition", value)
}

// GetClassDefinition gets the value of ClassDefinition for the instance
func (instance *CIM_ClassIndication) GetPropertyClassDefinition() (value interface{}, err error) {
	retValue, err := instance.GetProperty("ClassDefinition")
	if err != nil {
		return
	}
	value, ok := retValue.(interface{})
	if !ok {
		// TODO: Set an error
	}
	return
}
