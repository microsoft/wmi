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

// ads_ms_sql_sqlrepository struct
type ads_ms_sql_sqlrepository struct {
	*ds_top

	//
	DS_mS_SQL_Build int32

	//
	DS_mS_SQL_Contact string

	//
	DS_mS_SQL_Description string

	//
	DS_mS_SQL_InformationDirectory bool

	//
	DS_mS_SQL_Name string

	//
	DS_mS_SQL_Status int64

	//
	DS_mS_SQL_Version string
}

func Newads_ms_sql_sqlrepositoryEx1(instance *cim.WmiInstance) (newInstance *ads_ms_sql_sqlrepository, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_ms_sql_sqlrepository{
		ds_top: tmp,
	}
	return
}

func Newads_ms_sql_sqlrepositoryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_ms_sql_sqlrepository, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_ms_sql_sqlrepository{
		ds_top: tmp,
	}
	return
}

// SetDS_mS_SQL_Build sets the value of DS_mS_SQL_Build for the instance
func (instance *ads_ms_sql_sqlrepository) SetPropertyDS_mS_SQL_Build(value int32) (err error) {
	return instance.SetProperty("DS_mS_SQL_Build", (value))
}

// GetDS_mS_SQL_Build gets the value of DS_mS_SQL_Build for the instance
func (instance *ads_ms_sql_sqlrepository) GetPropertyDS_mS_SQL_Build() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_Build")
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

// SetDS_mS_SQL_Contact sets the value of DS_mS_SQL_Contact for the instance
func (instance *ads_ms_sql_sqlrepository) SetPropertyDS_mS_SQL_Contact(value string) (err error) {
	return instance.SetProperty("DS_mS_SQL_Contact", (value))
}

// GetDS_mS_SQL_Contact gets the value of DS_mS_SQL_Contact for the instance
func (instance *ads_ms_sql_sqlrepository) GetPropertyDS_mS_SQL_Contact() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_Contact")
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

// SetDS_mS_SQL_Description sets the value of DS_mS_SQL_Description for the instance
func (instance *ads_ms_sql_sqlrepository) SetPropertyDS_mS_SQL_Description(value string) (err error) {
	return instance.SetProperty("DS_mS_SQL_Description", (value))
}

// GetDS_mS_SQL_Description gets the value of DS_mS_SQL_Description for the instance
func (instance *ads_ms_sql_sqlrepository) GetPropertyDS_mS_SQL_Description() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_Description")
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

// SetDS_mS_SQL_InformationDirectory sets the value of DS_mS_SQL_InformationDirectory for the instance
func (instance *ads_ms_sql_sqlrepository) SetPropertyDS_mS_SQL_InformationDirectory(value bool) (err error) {
	return instance.SetProperty("DS_mS_SQL_InformationDirectory", (value))
}

// GetDS_mS_SQL_InformationDirectory gets the value of DS_mS_SQL_InformationDirectory for the instance
func (instance *ads_ms_sql_sqlrepository) GetPropertyDS_mS_SQL_InformationDirectory() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_InformationDirectory")
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

// SetDS_mS_SQL_Name sets the value of DS_mS_SQL_Name for the instance
func (instance *ads_ms_sql_sqlrepository) SetPropertyDS_mS_SQL_Name(value string) (err error) {
	return instance.SetProperty("DS_mS_SQL_Name", (value))
}

// GetDS_mS_SQL_Name gets the value of DS_mS_SQL_Name for the instance
func (instance *ads_ms_sql_sqlrepository) GetPropertyDS_mS_SQL_Name() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_Name")
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

// SetDS_mS_SQL_Status sets the value of DS_mS_SQL_Status for the instance
func (instance *ads_ms_sql_sqlrepository) SetPropertyDS_mS_SQL_Status(value int64) (err error) {
	return instance.SetProperty("DS_mS_SQL_Status", (value))
}

// GetDS_mS_SQL_Status gets the value of DS_mS_SQL_Status for the instance
func (instance *ads_ms_sql_sqlrepository) GetPropertyDS_mS_SQL_Status() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_Status")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetDS_mS_SQL_Version sets the value of DS_mS_SQL_Version for the instance
func (instance *ads_ms_sql_sqlrepository) SetPropertyDS_mS_SQL_Version(value string) (err error) {
	return instance.SetProperty("DS_mS_SQL_Version", (value))
}

// GetDS_mS_SQL_Version gets the value of DS_mS_SQL_Version for the instance
func (instance *ads_ms_sql_sqlrepository) GetPropertyDS_mS_SQL_Version() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_Version")
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
