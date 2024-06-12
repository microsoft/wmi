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

// ads_ms_sql_sqlpublication struct
type ads_ms_sql_sqlpublication struct {
	*ds_top

	//
	DS_mS_SQL_AllowAnonymousSubscription bool

	//
	DS_mS_SQL_AllowImmediateUpdatingSubscription bool

	//
	DS_mS_SQL_AllowKnownPullSubscription bool

	//
	DS_mS_SQL_AllowQueuedUpdatingSubscription bool

	//
	DS_mS_SQL_AllowSnapshotFilesFTPDownloading bool

	//
	DS_mS_SQL_Database string

	//
	DS_mS_SQL_Description string

	//
	DS_mS_SQL_Name string

	//
	DS_mS_SQL_Publisher string

	//
	DS_mS_SQL_Status int64

	//
	DS_mS_SQL_ThirdParty bool

	//
	DS_mS_SQL_Type string
}

func Newads_ms_sql_sqlpublicationEx1(instance *cim.WmiInstance) (newInstance *ads_ms_sql_sqlpublication, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_ms_sql_sqlpublication{
		ds_top: tmp,
	}
	return
}

func Newads_ms_sql_sqlpublicationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_ms_sql_sqlpublication, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_ms_sql_sqlpublication{
		ds_top: tmp,
	}
	return
}

// SetDS_mS_SQL_AllowAnonymousSubscription sets the value of DS_mS_SQL_AllowAnonymousSubscription for the instance
func (instance *ads_ms_sql_sqlpublication) SetPropertyDS_mS_SQL_AllowAnonymousSubscription(value bool) (err error) {
	return instance.SetProperty("DS_mS_SQL_AllowAnonymousSubscription", (value))
}

// GetDS_mS_SQL_AllowAnonymousSubscription gets the value of DS_mS_SQL_AllowAnonymousSubscription for the instance
func (instance *ads_ms_sql_sqlpublication) GetPropertyDS_mS_SQL_AllowAnonymousSubscription() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_AllowAnonymousSubscription")
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

// SetDS_mS_SQL_AllowImmediateUpdatingSubscription sets the value of DS_mS_SQL_AllowImmediateUpdatingSubscription for the instance
func (instance *ads_ms_sql_sqlpublication) SetPropertyDS_mS_SQL_AllowImmediateUpdatingSubscription(value bool) (err error) {
	return instance.SetProperty("DS_mS_SQL_AllowImmediateUpdatingSubscription", (value))
}

// GetDS_mS_SQL_AllowImmediateUpdatingSubscription gets the value of DS_mS_SQL_AllowImmediateUpdatingSubscription for the instance
func (instance *ads_ms_sql_sqlpublication) GetPropertyDS_mS_SQL_AllowImmediateUpdatingSubscription() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_AllowImmediateUpdatingSubscription")
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

// SetDS_mS_SQL_AllowKnownPullSubscription sets the value of DS_mS_SQL_AllowKnownPullSubscription for the instance
func (instance *ads_ms_sql_sqlpublication) SetPropertyDS_mS_SQL_AllowKnownPullSubscription(value bool) (err error) {
	return instance.SetProperty("DS_mS_SQL_AllowKnownPullSubscription", (value))
}

// GetDS_mS_SQL_AllowKnownPullSubscription gets the value of DS_mS_SQL_AllowKnownPullSubscription for the instance
func (instance *ads_ms_sql_sqlpublication) GetPropertyDS_mS_SQL_AllowKnownPullSubscription() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_AllowKnownPullSubscription")
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

// SetDS_mS_SQL_AllowQueuedUpdatingSubscription sets the value of DS_mS_SQL_AllowQueuedUpdatingSubscription for the instance
func (instance *ads_ms_sql_sqlpublication) SetPropertyDS_mS_SQL_AllowQueuedUpdatingSubscription(value bool) (err error) {
	return instance.SetProperty("DS_mS_SQL_AllowQueuedUpdatingSubscription", (value))
}

// GetDS_mS_SQL_AllowQueuedUpdatingSubscription gets the value of DS_mS_SQL_AllowQueuedUpdatingSubscription for the instance
func (instance *ads_ms_sql_sqlpublication) GetPropertyDS_mS_SQL_AllowQueuedUpdatingSubscription() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_AllowQueuedUpdatingSubscription")
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

// SetDS_mS_SQL_AllowSnapshotFilesFTPDownloading sets the value of DS_mS_SQL_AllowSnapshotFilesFTPDownloading for the instance
func (instance *ads_ms_sql_sqlpublication) SetPropertyDS_mS_SQL_AllowSnapshotFilesFTPDownloading(value bool) (err error) {
	return instance.SetProperty("DS_mS_SQL_AllowSnapshotFilesFTPDownloading", (value))
}

// GetDS_mS_SQL_AllowSnapshotFilesFTPDownloading gets the value of DS_mS_SQL_AllowSnapshotFilesFTPDownloading for the instance
func (instance *ads_ms_sql_sqlpublication) GetPropertyDS_mS_SQL_AllowSnapshotFilesFTPDownloading() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_AllowSnapshotFilesFTPDownloading")
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

// SetDS_mS_SQL_Database sets the value of DS_mS_SQL_Database for the instance
func (instance *ads_ms_sql_sqlpublication) SetPropertyDS_mS_SQL_Database(value string) (err error) {
	return instance.SetProperty("DS_mS_SQL_Database", (value))
}

// GetDS_mS_SQL_Database gets the value of DS_mS_SQL_Database for the instance
func (instance *ads_ms_sql_sqlpublication) GetPropertyDS_mS_SQL_Database() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_Database")
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
func (instance *ads_ms_sql_sqlpublication) SetPropertyDS_mS_SQL_Description(value string) (err error) {
	return instance.SetProperty("DS_mS_SQL_Description", (value))
}

// GetDS_mS_SQL_Description gets the value of DS_mS_SQL_Description for the instance
func (instance *ads_ms_sql_sqlpublication) GetPropertyDS_mS_SQL_Description() (value string, err error) {
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

// SetDS_mS_SQL_Name sets the value of DS_mS_SQL_Name for the instance
func (instance *ads_ms_sql_sqlpublication) SetPropertyDS_mS_SQL_Name(value string) (err error) {
	return instance.SetProperty("DS_mS_SQL_Name", (value))
}

// GetDS_mS_SQL_Name gets the value of DS_mS_SQL_Name for the instance
func (instance *ads_ms_sql_sqlpublication) GetPropertyDS_mS_SQL_Name() (value string, err error) {
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

// SetDS_mS_SQL_Publisher sets the value of DS_mS_SQL_Publisher for the instance
func (instance *ads_ms_sql_sqlpublication) SetPropertyDS_mS_SQL_Publisher(value string) (err error) {
	return instance.SetProperty("DS_mS_SQL_Publisher", (value))
}

// GetDS_mS_SQL_Publisher gets the value of DS_mS_SQL_Publisher for the instance
func (instance *ads_ms_sql_sqlpublication) GetPropertyDS_mS_SQL_Publisher() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_Publisher")
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
func (instance *ads_ms_sql_sqlpublication) SetPropertyDS_mS_SQL_Status(value int64) (err error) {
	return instance.SetProperty("DS_mS_SQL_Status", (value))
}

// GetDS_mS_SQL_Status gets the value of DS_mS_SQL_Status for the instance
func (instance *ads_ms_sql_sqlpublication) GetPropertyDS_mS_SQL_Status() (value int64, err error) {
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

// SetDS_mS_SQL_ThirdParty sets the value of DS_mS_SQL_ThirdParty for the instance
func (instance *ads_ms_sql_sqlpublication) SetPropertyDS_mS_SQL_ThirdParty(value bool) (err error) {
	return instance.SetProperty("DS_mS_SQL_ThirdParty", (value))
}

// GetDS_mS_SQL_ThirdParty gets the value of DS_mS_SQL_ThirdParty for the instance
func (instance *ads_ms_sql_sqlpublication) GetPropertyDS_mS_SQL_ThirdParty() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_ThirdParty")
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

// SetDS_mS_SQL_Type sets the value of DS_mS_SQL_Type for the instance
func (instance *ads_ms_sql_sqlpublication) SetPropertyDS_mS_SQL_Type(value string) (err error) {
	return instance.SetProperty("DS_mS_SQL_Type", (value))
}

// GetDS_mS_SQL_Type gets the value of DS_mS_SQL_Type for the instance
func (instance *ads_ms_sql_sqlpublication) GetPropertyDS_mS_SQL_Type() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mS_SQL_Type")
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
