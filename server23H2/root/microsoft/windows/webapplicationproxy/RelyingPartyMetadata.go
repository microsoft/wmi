// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.WebApplicationProxy
//
// ////////////////////////////////////////////
package webapplicationproxy

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RelyingPartyMetadata struct
type RelyingPartyMetadata struct {
	*cim.WmiInstance

	//
	Enabled bool

	//
	ID string

	//
	Name string

	//
	NonClaimsAware bool

	//
	Published bool
}

func NewRelyingPartyMetadataEx1(instance *cim.WmiInstance) (newInstance *RelyingPartyMetadata, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RelyingPartyMetadata{
		WmiInstance: tmp,
	}
	return
}

func NewRelyingPartyMetadataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RelyingPartyMetadata, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RelyingPartyMetadata{
		WmiInstance: tmp,
	}
	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *RelyingPartyMetadata) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *RelyingPartyMetadata) GetPropertyEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("Enabled")
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

// SetID sets the value of ID for the instance
func (instance *RelyingPartyMetadata) SetPropertyID(value string) (err error) {
	return instance.SetProperty("ID", (value))
}

// GetID gets the value of ID for the instance
func (instance *RelyingPartyMetadata) GetPropertyID() (value string, err error) {
	retValue, err := instance.GetProperty("ID")
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

// SetName sets the value of Name for the instance
func (instance *RelyingPartyMetadata) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *RelyingPartyMetadata) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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

// SetNonClaimsAware sets the value of NonClaimsAware for the instance
func (instance *RelyingPartyMetadata) SetPropertyNonClaimsAware(value bool) (err error) {
	return instance.SetProperty("NonClaimsAware", (value))
}

// GetNonClaimsAware gets the value of NonClaimsAware for the instance
func (instance *RelyingPartyMetadata) GetPropertyNonClaimsAware() (value bool, err error) {
	retValue, err := instance.GetProperty("NonClaimsAware")
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

// SetPublished sets the value of Published for the instance
func (instance *RelyingPartyMetadata) SetPropertyPublished(value bool) (err error) {
	return instance.SetProperty("Published", (value))
}

// GetPublished gets the value of Published for the instance
func (instance *RelyingPartyMetadata) GetPropertyPublished() (value bool, err error) {
	retValue, err := instance.GetProperty("Published")
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
