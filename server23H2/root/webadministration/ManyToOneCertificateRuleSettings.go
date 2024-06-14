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

// ManyToOneCertificateRuleSettings struct
type ManyToOneCertificateRuleSettings struct {
	*EmbeddedObject

	//
	Rules []ManyToOneCertificateMappingRuleElement
}

func NewManyToOneCertificateRuleSettingsEx1(instance *cim.WmiInstance) (newInstance *ManyToOneCertificateRuleSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ManyToOneCertificateRuleSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewManyToOneCertificateRuleSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ManyToOneCertificateRuleSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ManyToOneCertificateRuleSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetRules sets the value of Rules for the instance
func (instance *ManyToOneCertificateRuleSettings) SetPropertyRules(value []ManyToOneCertificateMappingRuleElement) (err error) {
	return instance.SetProperty("Rules", (value))
}

// GetRules gets the value of Rules for the instance
func (instance *ManyToOneCertificateRuleSettings) GetPropertyRules() (value []ManyToOneCertificateMappingRuleElement, err error) {
	retValue, err := instance.GetProperty("Rules")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(ManyToOneCertificateMappingRuleElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " ManyToOneCertificateMappingRuleElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, ManyToOneCertificateMappingRuleElement(valuetmp))
	}

	return
}