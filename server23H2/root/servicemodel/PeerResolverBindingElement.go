// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// PeerResolverBindingElement struct
type PeerResolverBindingElement struct {
	*BindingElement

	// Determines how referrals are shared among peers.
	ReferralPolicy string
}

func NewPeerResolverBindingElementEx1(instance *cim.WmiInstance) (newInstance *PeerResolverBindingElement, err error) {
	tmp, err := NewBindingElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &PeerResolverBindingElement{
		BindingElement: tmp,
	}
	return
}

func NewPeerResolverBindingElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PeerResolverBindingElement, err error) {
	tmp, err := NewBindingElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PeerResolverBindingElement{
		BindingElement: tmp,
	}
	return
}

// SetReferralPolicy sets the value of ReferralPolicy for the instance
func (instance *PeerResolverBindingElement) SetPropertyReferralPolicy(value string) (err error) {
	return instance.SetProperty("ReferralPolicy", (value))
}

// GetReferralPolicy gets the value of ReferralPolicy for the instance
func (instance *PeerResolverBindingElement) GetPropertyReferralPolicy() (value string, err error) {
	retValue, err := instance.GetProperty("ReferralPolicy")
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
