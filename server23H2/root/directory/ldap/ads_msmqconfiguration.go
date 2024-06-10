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

// ads_msmqconfiguration struct
type ads_msmqconfiguration struct {
	*ds_top

	//
	DS_mSMQComputerType string

	//
	DS_mSMQComputerTypeEx string

	//
	DS_mSMQDependentClientServices bool

	//
	DS_mSMQDsServices bool

	//
	DS_mSMQEncryptKey Uint8Array

	//
	DS_mSMQForeign bool

	//
	DS_mSMQInRoutingServers []string

	//
	DS_mSMQJournalQuota int32

	//
	DS_mSMQOSType int32

	//
	DS_mSMQOutRoutingServers []string

	//
	DS_mSMQOwnerID Uint8Array

	//
	DS_mSMQQuota int32

	//
	DS_mSMQRoutingServices bool

	//
	DS_mSMQServiceType int32

	//
	DS_mSMQSignKey Uint8Array

	//
	DS_mSMQSites []Uint8Array
}

func Newads_msmqconfigurationEx1(instance *cim.WmiInstance) (newInstance *ads_msmqconfiguration, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msmqconfiguration{
		ds_top: tmp,
	}
	return
}

func Newads_msmqconfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msmqconfiguration, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msmqconfiguration{
		ds_top: tmp,
	}
	return
}

// SetDS_mSMQComputerType sets the value of DS_mSMQComputerType for the instance
func (instance *ads_msmqconfiguration) SetPropertyDS_mSMQComputerType(value string) (err error) {
	return instance.SetProperty("DS_mSMQComputerType", (value))
}

// GetDS_mSMQComputerType gets the value of DS_mSMQComputerType for the instance
func (instance *ads_msmqconfiguration) GetPropertyDS_mSMQComputerType() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mSMQComputerType")
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

// SetDS_mSMQComputerTypeEx sets the value of DS_mSMQComputerTypeEx for the instance
func (instance *ads_msmqconfiguration) SetPropertyDS_mSMQComputerTypeEx(value string) (err error) {
	return instance.SetProperty("DS_mSMQComputerTypeEx", (value))
}

// GetDS_mSMQComputerTypeEx gets the value of DS_mSMQComputerTypeEx for the instance
func (instance *ads_msmqconfiguration) GetPropertyDS_mSMQComputerTypeEx() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mSMQComputerTypeEx")
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

// SetDS_mSMQDependentClientServices sets the value of DS_mSMQDependentClientServices for the instance
func (instance *ads_msmqconfiguration) SetPropertyDS_mSMQDependentClientServices(value bool) (err error) {
	return instance.SetProperty("DS_mSMQDependentClientServices", (value))
}

// GetDS_mSMQDependentClientServices gets the value of DS_mSMQDependentClientServices for the instance
func (instance *ads_msmqconfiguration) GetPropertyDS_mSMQDependentClientServices() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_mSMQDependentClientServices")
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

// SetDS_mSMQDsServices sets the value of DS_mSMQDsServices for the instance
func (instance *ads_msmqconfiguration) SetPropertyDS_mSMQDsServices(value bool) (err error) {
	return instance.SetProperty("DS_mSMQDsServices", (value))
}

// GetDS_mSMQDsServices gets the value of DS_mSMQDsServices for the instance
func (instance *ads_msmqconfiguration) GetPropertyDS_mSMQDsServices() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_mSMQDsServices")
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

// SetDS_mSMQEncryptKey sets the value of DS_mSMQEncryptKey for the instance
func (instance *ads_msmqconfiguration) SetPropertyDS_mSMQEncryptKey(value Uint8Array) (err error) {
	return instance.SetProperty("DS_mSMQEncryptKey", (value))
}

// GetDS_mSMQEncryptKey gets the value of DS_mSMQEncryptKey for the instance
func (instance *ads_msmqconfiguration) GetPropertyDS_mSMQEncryptKey() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_mSMQEncryptKey")
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

// SetDS_mSMQForeign sets the value of DS_mSMQForeign for the instance
func (instance *ads_msmqconfiguration) SetPropertyDS_mSMQForeign(value bool) (err error) {
	return instance.SetProperty("DS_mSMQForeign", (value))
}

// GetDS_mSMQForeign gets the value of DS_mSMQForeign for the instance
func (instance *ads_msmqconfiguration) GetPropertyDS_mSMQForeign() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_mSMQForeign")
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

// SetDS_mSMQInRoutingServers sets the value of DS_mSMQInRoutingServers for the instance
func (instance *ads_msmqconfiguration) SetPropertyDS_mSMQInRoutingServers(value []string) (err error) {
	return instance.SetProperty("DS_mSMQInRoutingServers", (value))
}

// GetDS_mSMQInRoutingServers gets the value of DS_mSMQInRoutingServers for the instance
func (instance *ads_msmqconfiguration) GetPropertyDS_mSMQInRoutingServers() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_mSMQInRoutingServers")
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

// SetDS_mSMQJournalQuota sets the value of DS_mSMQJournalQuota for the instance
func (instance *ads_msmqconfiguration) SetPropertyDS_mSMQJournalQuota(value int32) (err error) {
	return instance.SetProperty("DS_mSMQJournalQuota", (value))
}

// GetDS_mSMQJournalQuota gets the value of DS_mSMQJournalQuota for the instance
func (instance *ads_msmqconfiguration) GetPropertyDS_mSMQJournalQuota() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_mSMQJournalQuota")
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

// SetDS_mSMQOSType sets the value of DS_mSMQOSType for the instance
func (instance *ads_msmqconfiguration) SetPropertyDS_mSMQOSType(value int32) (err error) {
	return instance.SetProperty("DS_mSMQOSType", (value))
}

// GetDS_mSMQOSType gets the value of DS_mSMQOSType for the instance
func (instance *ads_msmqconfiguration) GetPropertyDS_mSMQOSType() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_mSMQOSType")
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

// SetDS_mSMQOutRoutingServers sets the value of DS_mSMQOutRoutingServers for the instance
func (instance *ads_msmqconfiguration) SetPropertyDS_mSMQOutRoutingServers(value []string) (err error) {
	return instance.SetProperty("DS_mSMQOutRoutingServers", (value))
}

// GetDS_mSMQOutRoutingServers gets the value of DS_mSMQOutRoutingServers for the instance
func (instance *ads_msmqconfiguration) GetPropertyDS_mSMQOutRoutingServers() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_mSMQOutRoutingServers")
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

// SetDS_mSMQOwnerID sets the value of DS_mSMQOwnerID for the instance
func (instance *ads_msmqconfiguration) SetPropertyDS_mSMQOwnerID(value Uint8Array) (err error) {
	return instance.SetProperty("DS_mSMQOwnerID", (value))
}

// GetDS_mSMQOwnerID gets the value of DS_mSMQOwnerID for the instance
func (instance *ads_msmqconfiguration) GetPropertyDS_mSMQOwnerID() (value Uint8Array, err error) {
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

// SetDS_mSMQQuota sets the value of DS_mSMQQuota for the instance
func (instance *ads_msmqconfiguration) SetPropertyDS_mSMQQuota(value int32) (err error) {
	return instance.SetProperty("DS_mSMQQuota", (value))
}

// GetDS_mSMQQuota gets the value of DS_mSMQQuota for the instance
func (instance *ads_msmqconfiguration) GetPropertyDS_mSMQQuota() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_mSMQQuota")
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

// SetDS_mSMQRoutingServices sets the value of DS_mSMQRoutingServices for the instance
func (instance *ads_msmqconfiguration) SetPropertyDS_mSMQRoutingServices(value bool) (err error) {
	return instance.SetProperty("DS_mSMQRoutingServices", (value))
}

// GetDS_mSMQRoutingServices gets the value of DS_mSMQRoutingServices for the instance
func (instance *ads_msmqconfiguration) GetPropertyDS_mSMQRoutingServices() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_mSMQRoutingServices")
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

// SetDS_mSMQServiceType sets the value of DS_mSMQServiceType for the instance
func (instance *ads_msmqconfiguration) SetPropertyDS_mSMQServiceType(value int32) (err error) {
	return instance.SetProperty("DS_mSMQServiceType", (value))
}

// GetDS_mSMQServiceType gets the value of DS_mSMQServiceType for the instance
func (instance *ads_msmqconfiguration) GetPropertyDS_mSMQServiceType() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_mSMQServiceType")
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

// SetDS_mSMQSignKey sets the value of DS_mSMQSignKey for the instance
func (instance *ads_msmqconfiguration) SetPropertyDS_mSMQSignKey(value Uint8Array) (err error) {
	return instance.SetProperty("DS_mSMQSignKey", (value))
}

// GetDS_mSMQSignKey gets the value of DS_mSMQSignKey for the instance
func (instance *ads_msmqconfiguration) GetPropertyDS_mSMQSignKey() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_mSMQSignKey")
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

// SetDS_mSMQSites sets the value of DS_mSMQSites for the instance
func (instance *ads_msmqconfiguration) SetPropertyDS_mSMQSites(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_mSMQSites", (value))
}

// GetDS_mSMQSites gets the value of DS_mSMQSites for the instance
func (instance *ads_msmqconfiguration) GetPropertyDS_mSMQSites() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_mSMQSites")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(Uint8Array)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, Uint8Array(valuetmp))
	}

	return
}
