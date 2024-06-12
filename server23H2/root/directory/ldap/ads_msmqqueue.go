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

// ads_msmqqueue struct
type ads_msmqqueue struct {
	*ds_top

	//
	DS_MSMQ_MulticastAddress string

	//
	DS_MSMQ_SecuredSource bool

	//
	DS_mSMQAuthenticate bool

	//
	DS_mSMQBasePriority int32

	//
	DS_mSMQJournal bool

	//
	DS_mSMQLabel string

	//
	DS_mSMQLabelEx string

	//
	DS_mSMQOwnerID Uint8Array

	//
	DS_mSMQPrivacyLevel int32

	//
	DS_mSMQQueueJournalQuota int32

	//
	DS_mSMQQueueNameExt string

	//
	DS_mSMQQueueQuota int32

	//
	DS_mSMQQueueType Uint8Array

	//
	DS_mSMQTransactional bool
}

func Newads_msmqqueueEx1(instance *cim.WmiInstance) (newInstance *ads_msmqqueue, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msmqqueue{
		ds_top: tmp,
	}
	return
}

func Newads_msmqqueueEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msmqqueue, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msmqqueue{
		ds_top: tmp,
	}
	return
}

// SetDS_MSMQ_MulticastAddress sets the value of DS_MSMQ_MulticastAddress for the instance
func (instance *ads_msmqqueue) SetPropertyDS_MSMQ_MulticastAddress(value string) (err error) {
	return instance.SetProperty("DS_MSMQ_MulticastAddress", (value))
}

// GetDS_MSMQ_MulticastAddress gets the value of DS_MSMQ_MulticastAddress for the instance
func (instance *ads_msmqqueue) GetPropertyDS_MSMQ_MulticastAddress() (value string, err error) {
	retValue, err := instance.GetProperty("DS_MSMQ_MulticastAddress")
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

// SetDS_MSMQ_SecuredSource sets the value of DS_MSMQ_SecuredSource for the instance
func (instance *ads_msmqqueue) SetPropertyDS_MSMQ_SecuredSource(value bool) (err error) {
	return instance.SetProperty("DS_MSMQ_SecuredSource", (value))
}

// GetDS_MSMQ_SecuredSource gets the value of DS_MSMQ_SecuredSource for the instance
func (instance *ads_msmqqueue) GetPropertyDS_MSMQ_SecuredSource() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_MSMQ_SecuredSource")
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

// SetDS_mSMQAuthenticate sets the value of DS_mSMQAuthenticate for the instance
func (instance *ads_msmqqueue) SetPropertyDS_mSMQAuthenticate(value bool) (err error) {
	return instance.SetProperty("DS_mSMQAuthenticate", (value))
}

// GetDS_mSMQAuthenticate gets the value of DS_mSMQAuthenticate for the instance
func (instance *ads_msmqqueue) GetPropertyDS_mSMQAuthenticate() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_mSMQAuthenticate")
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

// SetDS_mSMQBasePriority sets the value of DS_mSMQBasePriority for the instance
func (instance *ads_msmqqueue) SetPropertyDS_mSMQBasePriority(value int32) (err error) {
	return instance.SetProperty("DS_mSMQBasePriority", (value))
}

// GetDS_mSMQBasePriority gets the value of DS_mSMQBasePriority for the instance
func (instance *ads_msmqqueue) GetPropertyDS_mSMQBasePriority() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_mSMQBasePriority")
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

// SetDS_mSMQJournal sets the value of DS_mSMQJournal for the instance
func (instance *ads_msmqqueue) SetPropertyDS_mSMQJournal(value bool) (err error) {
	return instance.SetProperty("DS_mSMQJournal", (value))
}

// GetDS_mSMQJournal gets the value of DS_mSMQJournal for the instance
func (instance *ads_msmqqueue) GetPropertyDS_mSMQJournal() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_mSMQJournal")
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

// SetDS_mSMQLabel sets the value of DS_mSMQLabel for the instance
func (instance *ads_msmqqueue) SetPropertyDS_mSMQLabel(value string) (err error) {
	return instance.SetProperty("DS_mSMQLabel", (value))
}

// GetDS_mSMQLabel gets the value of DS_mSMQLabel for the instance
func (instance *ads_msmqqueue) GetPropertyDS_mSMQLabel() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mSMQLabel")
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

// SetDS_mSMQLabelEx sets the value of DS_mSMQLabelEx for the instance
func (instance *ads_msmqqueue) SetPropertyDS_mSMQLabelEx(value string) (err error) {
	return instance.SetProperty("DS_mSMQLabelEx", (value))
}

// GetDS_mSMQLabelEx gets the value of DS_mSMQLabelEx for the instance
func (instance *ads_msmqqueue) GetPropertyDS_mSMQLabelEx() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mSMQLabelEx")
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

// SetDS_mSMQOwnerID sets the value of DS_mSMQOwnerID for the instance
func (instance *ads_msmqqueue) SetPropertyDS_mSMQOwnerID(value Uint8Array) (err error) {
	return instance.SetProperty("DS_mSMQOwnerID", (value))
}

// GetDS_mSMQOwnerID gets the value of DS_mSMQOwnerID for the instance
func (instance *ads_msmqqueue) GetPropertyDS_mSMQOwnerID() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_mSMQOwnerID")
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

// SetDS_mSMQPrivacyLevel sets the value of DS_mSMQPrivacyLevel for the instance
func (instance *ads_msmqqueue) SetPropertyDS_mSMQPrivacyLevel(value int32) (err error) {
	return instance.SetProperty("DS_mSMQPrivacyLevel", (value))
}

// GetDS_mSMQPrivacyLevel gets the value of DS_mSMQPrivacyLevel for the instance
func (instance *ads_msmqqueue) GetPropertyDS_mSMQPrivacyLevel() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_mSMQPrivacyLevel")
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

// SetDS_mSMQQueueJournalQuota sets the value of DS_mSMQQueueJournalQuota for the instance
func (instance *ads_msmqqueue) SetPropertyDS_mSMQQueueJournalQuota(value int32) (err error) {
	return instance.SetProperty("DS_mSMQQueueJournalQuota", (value))
}

// GetDS_mSMQQueueJournalQuota gets the value of DS_mSMQQueueJournalQuota for the instance
func (instance *ads_msmqqueue) GetPropertyDS_mSMQQueueJournalQuota() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_mSMQQueueJournalQuota")
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

// SetDS_mSMQQueueNameExt sets the value of DS_mSMQQueueNameExt for the instance
func (instance *ads_msmqqueue) SetPropertyDS_mSMQQueueNameExt(value string) (err error) {
	return instance.SetProperty("DS_mSMQQueueNameExt", (value))
}

// GetDS_mSMQQueueNameExt gets the value of DS_mSMQQueueNameExt for the instance
func (instance *ads_msmqqueue) GetPropertyDS_mSMQQueueNameExt() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mSMQQueueNameExt")
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

// SetDS_mSMQQueueQuota sets the value of DS_mSMQQueueQuota for the instance
func (instance *ads_msmqqueue) SetPropertyDS_mSMQQueueQuota(value int32) (err error) {
	return instance.SetProperty("DS_mSMQQueueQuota", (value))
}

// GetDS_mSMQQueueQuota gets the value of DS_mSMQQueueQuota for the instance
func (instance *ads_msmqqueue) GetPropertyDS_mSMQQueueQuota() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_mSMQQueueQuota")
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

// SetDS_mSMQQueueType sets the value of DS_mSMQQueueType for the instance
func (instance *ads_msmqqueue) SetPropertyDS_mSMQQueueType(value Uint8Array) (err error) {
	return instance.SetProperty("DS_mSMQQueueType", (value))
}

// GetDS_mSMQQueueType gets the value of DS_mSMQQueueType for the instance
func (instance *ads_msmqqueue) GetPropertyDS_mSMQQueueType() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_mSMQQueueType")
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

// SetDS_mSMQTransactional sets the value of DS_mSMQTransactional for the instance
func (instance *ads_msmqqueue) SetPropertyDS_mSMQTransactional(value bool) (err error) {
	return instance.SetProperty("DS_mSMQTransactional", (value))
}

// GetDS_mSMQTransactional gets the value of DS_mSMQTransactional for the instance
func (instance *ads_msmqqueue) GetPropertyDS_mSMQTransactional() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_mSMQTransactional")
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
