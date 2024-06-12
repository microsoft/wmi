// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ads_msdfs_deletedlinkv2 struct
type ads_msdfs_deletedlinkv2 struct {
	*ds_top

	//
	DS_msDFS_Commentv2 string

	//
	DS_msDFS_LastModifiedv2 string

	//
	DS_msDFS_LinkIdentityGUIDv2 Uint8Array

	//
	DS_msDFS_LinkPathv2 string

	//
	DS_msDFS_NamespaceIdentityGUIDv2 Uint8Array

	//
	DS_msDFS_ShortNameLinkPathv2 string
}

func Newads_msdfs_deletedlinkv2Ex1(instance *cim.WmiInstance) (newInstance *ads_msdfs_deletedlinkv2, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msdfs_deletedlinkv2{
		ds_top: tmp,
	}
	return
}

func Newads_msdfs_deletedlinkv2Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msdfs_deletedlinkv2, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msdfs_deletedlinkv2{
		ds_top: tmp,
	}
	return
}

// SetDS_msDFS_Commentv2 sets the value of DS_msDFS_Commentv2 for the instance
func (instance *ads_msdfs_deletedlinkv2) SetPropertyDS_msDFS_Commentv2(value string) (err error) {
	return instance.SetProperty("DS_msDFS_Commentv2", (value))
}

// GetDS_msDFS_Commentv2 gets the value of DS_msDFS_Commentv2 for the instance
func (instance *ads_msdfs_deletedlinkv2) GetPropertyDS_msDFS_Commentv2() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDFS_Commentv2")
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

// SetDS_msDFS_LastModifiedv2 sets the value of DS_msDFS_LastModifiedv2 for the instance
func (instance *ads_msdfs_deletedlinkv2) SetPropertyDS_msDFS_LastModifiedv2(value string) (err error) {
	return instance.SetProperty("DS_msDFS_LastModifiedv2", (value))
}

// GetDS_msDFS_LastModifiedv2 gets the value of DS_msDFS_LastModifiedv2 for the instance
func (instance *ads_msdfs_deletedlinkv2) GetPropertyDS_msDFS_LastModifiedv2() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDFS_LastModifiedv2")
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

// SetDS_msDFS_LinkIdentityGUIDv2 sets the value of DS_msDFS_LinkIdentityGUIDv2 for the instance
func (instance *ads_msdfs_deletedlinkv2) SetPropertyDS_msDFS_LinkIdentityGUIDv2(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDFS_LinkIdentityGUIDv2", (value))
}

// GetDS_msDFS_LinkIdentityGUIDv2 gets the value of DS_msDFS_LinkIdentityGUIDv2 for the instance
func (instance *ads_msdfs_deletedlinkv2) GetPropertyDS_msDFS_LinkIdentityGUIDv2() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDFS_LinkIdentityGUIDv2")
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

// SetDS_msDFS_LinkPathv2 sets the value of DS_msDFS_LinkPathv2 for the instance
func (instance *ads_msdfs_deletedlinkv2) SetPropertyDS_msDFS_LinkPathv2(value string) (err error) {
	return instance.SetProperty("DS_msDFS_LinkPathv2", (value))
}

// GetDS_msDFS_LinkPathv2 gets the value of DS_msDFS_LinkPathv2 for the instance
func (instance *ads_msdfs_deletedlinkv2) GetPropertyDS_msDFS_LinkPathv2() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDFS_LinkPathv2")
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

// SetDS_msDFS_NamespaceIdentityGUIDv2 sets the value of DS_msDFS_NamespaceIdentityGUIDv2 for the instance
func (instance *ads_msdfs_deletedlinkv2) SetPropertyDS_msDFS_NamespaceIdentityGUIDv2(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDFS_NamespaceIdentityGUIDv2", (value))
}

// GetDS_msDFS_NamespaceIdentityGUIDv2 gets the value of DS_msDFS_NamespaceIdentityGUIDv2 for the instance
func (instance *ads_msdfs_deletedlinkv2) GetPropertyDS_msDFS_NamespaceIdentityGUIDv2() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDFS_NamespaceIdentityGUIDv2")
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

// SetDS_msDFS_ShortNameLinkPathv2 sets the value of DS_msDFS_ShortNameLinkPathv2 for the instance
func (instance *ads_msdfs_deletedlinkv2) SetPropertyDS_msDFS_ShortNameLinkPathv2(value string) (err error) {
	return instance.SetProperty("DS_msDFS_ShortNameLinkPathv2", (value))
}

// GetDS_msDFS_ShortNameLinkPathv2 gets the value of DS_msDFS_ShortNameLinkPathv2 for the instance
func (instance *ads_msdfs_deletedlinkv2) GetPropertyDS_msDFS_ShortNameLinkPathv2() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDFS_ShortNameLinkPathv2")
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
