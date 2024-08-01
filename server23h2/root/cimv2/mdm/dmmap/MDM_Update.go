// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_Update struct
type MDM_Update struct {
	*cim.WmiInstance

	//
	DeferUpgrade int32

	//
	InstanceID string

	//
	LastSuccessfulScanTime string

	//
	ParentID string
}

func NewMDM_UpdateEx1(instance *cim.WmiInstance) (newInstance *MDM_Update, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Update{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_UpdateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Update, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Update{
		WmiInstance: tmp,
	}
	return
}

// SetDeferUpgrade sets the value of DeferUpgrade for the instance
func (instance *MDM_Update) SetPropertyDeferUpgrade(value int32) (err error) {
	return instance.SetProperty("DeferUpgrade", (value))
}

// GetDeferUpgrade gets the value of DeferUpgrade for the instance
func (instance *MDM_Update) GetPropertyDeferUpgrade() (value int32, err error) {
	retValue, err := instance.GetProperty("DeferUpgrade")
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

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Update) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Update) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
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

// SetLastSuccessfulScanTime sets the value of LastSuccessfulScanTime for the instance
func (instance *MDM_Update) SetPropertyLastSuccessfulScanTime(value string) (err error) {
	return instance.SetProperty("LastSuccessfulScanTime", (value))
}

// GetLastSuccessfulScanTime gets the value of LastSuccessfulScanTime for the instance
func (instance *MDM_Update) GetPropertyLastSuccessfulScanTime() (value string, err error) {
	retValue, err := instance.GetProperty("LastSuccessfulScanTime")
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Update) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Update) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
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
