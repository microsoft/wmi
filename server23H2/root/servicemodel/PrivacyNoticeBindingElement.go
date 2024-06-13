// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// PrivacyNoticeBindingElement struct
type PrivacyNoticeBindingElement struct {
	*BindingElement

	// The privacy notice version.
	PrivacyNoticeVersion int32

	// The URI at which the privacy notice is located.
	Url string
}

func NewPrivacyNoticeBindingElementEx1(instance *cim.WmiInstance) (newInstance *PrivacyNoticeBindingElement, err error) {
	tmp, err := NewBindingElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &PrivacyNoticeBindingElement{
		BindingElement: tmp,
	}
	return
}

func NewPrivacyNoticeBindingElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PrivacyNoticeBindingElement, err error) {
	tmp, err := NewBindingElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PrivacyNoticeBindingElement{
		BindingElement: tmp,
	}
	return
}

// SetPrivacyNoticeVersion sets the value of PrivacyNoticeVersion for the instance
func (instance *PrivacyNoticeBindingElement) SetPropertyPrivacyNoticeVersion(value int32) (err error) {
	return instance.SetProperty("PrivacyNoticeVersion", (value))
}

// GetPrivacyNoticeVersion gets the value of PrivacyNoticeVersion for the instance
func (instance *PrivacyNoticeBindingElement) GetPropertyPrivacyNoticeVersion() (value int32, err error) {
	retValue, err := instance.GetProperty("PrivacyNoticeVersion")
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

	value = int32(valuetmp)

	return
}

// SetUrl sets the value of Url for the instance
func (instance *PrivacyNoticeBindingElement) SetPropertyUrl(value string) (err error) {
	return instance.SetProperty("Url", (value))
}

// GetUrl gets the value of Url for the instance
func (instance *PrivacyNoticeBindingElement) GetPropertyUrl() (value string, err error) {
	retValue, err := instance.GetProperty("Url")
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
