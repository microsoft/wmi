// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ads_mstpm_informationobject struct
type ads_mstpm_informationobject struct {
	*ds_top

	//
	DS_msTPM_OwnerInformation string

	//
	DS_msTPM_OwnerInformationTemp string

	//
	DS_msTPM_SrkPubThumbprint Uint8Array
}

func Newads_mstpm_informationobjectEx1(instance *cim.WmiInstance) (newInstance *ads_mstpm_informationobject, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_mstpm_informationobject{
		ds_top: tmp,
	}
	return
}

func Newads_mstpm_informationobjectEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_mstpm_informationobject, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_mstpm_informationobject{
		ds_top: tmp,
	}
	return
}

// SetDS_msTPM_OwnerInformation sets the value of DS_msTPM_OwnerInformation for the instance
func (instance *ads_mstpm_informationobject) SetPropertyDS_msTPM_OwnerInformation(value string) (err error) {
	return instance.SetProperty("DS_msTPM_OwnerInformation", (value))
}

// GetDS_msTPM_OwnerInformation gets the value of DS_msTPM_OwnerInformation for the instance
func (instance *ads_mstpm_informationobject) GetPropertyDS_msTPM_OwnerInformation() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msTPM_OwnerInformation")
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

// SetDS_msTPM_OwnerInformationTemp sets the value of DS_msTPM_OwnerInformationTemp for the instance
func (instance *ads_mstpm_informationobject) SetPropertyDS_msTPM_OwnerInformationTemp(value string) (err error) {
	return instance.SetProperty("DS_msTPM_OwnerInformationTemp", (value))
}

// GetDS_msTPM_OwnerInformationTemp gets the value of DS_msTPM_OwnerInformationTemp for the instance
func (instance *ads_mstpm_informationobject) GetPropertyDS_msTPM_OwnerInformationTemp() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msTPM_OwnerInformationTemp")
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

// SetDS_msTPM_SrkPubThumbprint sets the value of DS_msTPM_SrkPubThumbprint for the instance
func (instance *ads_mstpm_informationobject) SetPropertyDS_msTPM_SrkPubThumbprint(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msTPM_SrkPubThumbprint", (value))
}

// GetDS_msTPM_SrkPubThumbprint gets the value of DS_msTPM_SrkPubThumbprint for the instance
func (instance *ads_mstpm_informationobject) GetPropertyDS_msTPM_SrkPubThumbprint() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msTPM_SrkPubThumbprint")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(Uint8Array)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = Uint8Array(valuetmp)

	return
}
