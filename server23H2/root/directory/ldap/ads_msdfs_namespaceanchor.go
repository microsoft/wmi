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

// ads_msdfs_namespaceanchor struct
type ads_msdfs_namespaceanchor struct {
	*ds_top

	//
	DS_msDFS_SchemaMajorVersion int32
}

func Newads_msdfs_namespaceanchorEx1(instance *cim.WmiInstance) (newInstance *ads_msdfs_namespaceanchor, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msdfs_namespaceanchor{
		ds_top: tmp,
	}
	return
}

func Newads_msdfs_namespaceanchorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msdfs_namespaceanchor, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msdfs_namespaceanchor{
		ds_top: tmp,
	}
	return
}

// SetDS_msDFS_SchemaMajorVersion sets the value of DS_msDFS_SchemaMajorVersion for the instance
func (instance *ads_msdfs_namespaceanchor) SetPropertyDS_msDFS_SchemaMajorVersion(value int32) (err error) {
	return instance.SetProperty("DS_msDFS_SchemaMajorVersion", (value))
}

// GetDS_msDFS_SchemaMajorVersion gets the value of DS_msDFS_SchemaMajorVersion for the instance
func (instance *ads_msdfs_namespaceanchor) GetPropertyDS_msDFS_SchemaMajorVersion() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDFS_SchemaMajorVersion")
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
