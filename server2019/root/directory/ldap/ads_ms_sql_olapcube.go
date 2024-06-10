// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.directory.LDAP
//
// ////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ads_ms_sql_olapcube struct
type ads_ms_sql_olapcube struct {
	*ds_top

	//
	DS_mS_SQL_Contact string

	//
	DS_mS_SQL_Description string

	//
	DS_mS_SQL_InformationURL string

	//
	DS_mS_SQL_Keywords []string

	//
	DS_mS_SQL_LastUpdatedDate string

	//
	DS_mS_SQL_Name string

	//
	DS_mS_SQL_PublicationURL string

	//
	DS_mS_SQL_Size int64

	//
	DS_mS_SQL_Status int64
}

func Newads_ms_sql_olapcubeEx1(instance *cim.WmiInstance) (newInstance *ads_ms_sql_olapcube, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_ms_sql_olapcube{
		ds_top: tmp,
	}
	return
}

func Newads_ms_sql_olapcubeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_ms_sql_olapcube, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_ms_sql_olapcube{
		ds_top: tmp,
	}
	return
}

// SetDS_mS_SQL_Contact sets the value of DS_mS_SQL_Contact for the instance
func (instance *ads_ms_sql_olapcube) SetPropertyDS_mS_SQL_Contact(value string) (err error) {
	return instance.SetProperty("DS_mS_SQL_Contact", (value))
}

// GetDS_mS_SQL_Contact gets the value of DS_mS_SQL_Contact for the instance
func (instance *ads_ms_sql_olapcube) GetPropertyDS_mS_SQL_Contact() (value string, err error) {
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
func (instance *ads_ms_sql_olapcube) SetPropertyDS_mS_SQL_Description(value string) (err error) {
	return instance.SetProperty("DS_mS_SQL_Description", (value))
}

// GetDS_mS_SQL_Description gets the value of DS_mS_SQL_Description for the instance
func (instance *ads_ms_sql_olapcube) GetPropertyDS_mS_SQL_Description() (value string, err error) {
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

// SetDS_mS_SQL_InformationURL sets the value of DS_mS_SQL_InformationURL for the instance
func (instance *ads_ms_sql_olapcube) SetPropertyDS_mS_SQL_InformationURL(value string) (err error) {
	return instance.SetProperty("DS_mS_SQL_InformationURL", (value))
}

// GetDS_mS_SQL_InformationURL gets the value of DS_mS_SQL_InformationURL for the instance
func (instance *ads_ms_sql_olapcube) GetPropertyDS_mS_SQL_InformationURL() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_InformationURL")
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

// SetDS_mS_SQL_Keywords sets the value of DS_mS_SQL_Keywords for the instance
func (instance *ads_ms_sql_olapcube) SetPropertyDS_mS_SQL_Keywords(value []string) (err error) {
	return instance.SetProperty("DS_mS_SQL_Keywords", (value))
}

// GetDS_mS_SQL_Keywords gets the value of DS_mS_SQL_Keywords for the instance
func (instance *ads_ms_sql_olapcube) GetPropertyDS_mS_SQL_Keywords() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_Keywords")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetDS_mS_SQL_LastUpdatedDate sets the value of DS_mS_SQL_LastUpdatedDate for the instance
func (instance *ads_ms_sql_olapcube) SetPropertyDS_mS_SQL_LastUpdatedDate(value string) (err error) {
	return instance.SetProperty("DS_mS_SQL_LastUpdatedDate", (value))
}

// GetDS_mS_SQL_LastUpdatedDate gets the value of DS_mS_SQL_LastUpdatedDate for the instance
func (instance *ads_ms_sql_olapcube) GetPropertyDS_mS_SQL_LastUpdatedDate() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_LastUpdatedDate")
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

// SetDS_mS_SQL_Name sets the value of DS_mS_SQL_Name for the instance
func (instance *ads_ms_sql_olapcube) SetPropertyDS_mS_SQL_Name(value string) (err error) {
	return instance.SetProperty("DS_mS_SQL_Name", (value))
}

// GetDS_mS_SQL_Name gets the value of DS_mS_SQL_Name for the instance
func (instance *ads_ms_sql_olapcube) GetPropertyDS_mS_SQL_Name() (value string, err error) {
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

// SetDS_mS_SQL_PublicationURL sets the value of DS_mS_SQL_PublicationURL for the instance
func (instance *ads_ms_sql_olapcube) SetPropertyDS_mS_SQL_PublicationURL(value string) (err error) {
	return instance.SetProperty("DS_mS_SQL_PublicationURL", (value))
}

// GetDS_mS_SQL_PublicationURL gets the value of DS_mS_SQL_PublicationURL for the instance
func (instance *ads_ms_sql_olapcube) GetPropertyDS_mS_SQL_PublicationURL() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_PublicationURL")
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

// SetDS_mS_SQL_Size sets the value of DS_mS_SQL_Size for the instance
func (instance *ads_ms_sql_olapcube) SetPropertyDS_mS_SQL_Size(value int64) (err error) {
	return instance.SetProperty("DS_mS_SQL_Size", (value))
}

// GetDS_mS_SQL_Size gets the value of DS_mS_SQL_Size for the instance
func (instance *ads_ms_sql_olapcube) GetPropertyDS_mS_SQL_Size() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_Size")
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

// SetDS_mS_SQL_Status sets the value of DS_mS_SQL_Status for the instance
func (instance *ads_ms_sql_olapcube) SetPropertyDS_mS_SQL_Status(value int64) (err error) {
	return instance.SetProperty("DS_mS_SQL_Status", (value))
}

// GetDS_mS_SQL_Status gets the value of DS_mS_SQL_Status for the instance
func (instance *ads_ms_sql_olapcube) GetPropertyDS_mS_SQL_Status() (value int64, err error) {
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
