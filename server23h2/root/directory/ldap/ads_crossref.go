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

// ads_crossref struct
type ads_crossref struct {
	*ds_top

	//
	DS_dnsRoot []string

	//
	DS_Enabled bool

	//
	DS_msDS_Behavior_Version int32

	//
	DS_msDS_DnsRootAlias string

	//
	DS_msDS_NC_Replica_Locations []string

	//
	DS_msDS_NC_RO_Replica_Locations []string

	//
	DS_msDS_Replication_Notify_First_DSA_Delay int32

	//
	DS_msDS_Replication_Notify_Subsequent_DSA_Delay int32

	//
	DS_msDS_SDReferenceDomain string

	//
	DS_nCName string

	//
	DS_nETBIOSName string

	//
	DS_nTMixedDomain int32

	//
	DS_rootTrust []string

	//
	DS_superiorDNSRoot string

	//
	DS_trustParent string
}

func Newads_crossrefEx1(instance *cim.WmiInstance) (newInstance *ads_crossref, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_crossref{
		ds_top: tmp,
	}
	return
}

func Newads_crossrefEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_crossref, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_crossref{
		ds_top: tmp,
	}
	return
}

// SetDS_dnsRoot sets the value of DS_dnsRoot for the instance
func (instance *ads_crossref) SetPropertyDS_dnsRoot(value []string) (err error) {
	return instance.SetProperty("DS_dnsRoot", (value))
}

// GetDS_dnsRoot gets the value of DS_dnsRoot for the instance
func (instance *ads_crossref) GetPropertyDS_dnsRoot() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_dnsRoot")
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

// SetDS_Enabled sets the value of DS_Enabled for the instance
func (instance *ads_crossref) SetPropertyDS_Enabled(value bool) (err error) {
	return instance.SetProperty("DS_Enabled", (value))
}

// GetDS_Enabled gets the value of DS_Enabled for the instance
func (instance *ads_crossref) GetPropertyDS_Enabled() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_Enabled")
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

// SetDS_msDS_Behavior_Version sets the value of DS_msDS_Behavior_Version for the instance
func (instance *ads_crossref) SetPropertyDS_msDS_Behavior_Version(value int32) (err error) {
	return instance.SetProperty("DS_msDS_Behavior_Version", (value))
}

// GetDS_msDS_Behavior_Version gets the value of DS_msDS_Behavior_Version for the instance
func (instance *ads_crossref) GetPropertyDS_msDS_Behavior_Version() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_Behavior_Version")
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

// SetDS_msDS_DnsRootAlias sets the value of DS_msDS_DnsRootAlias for the instance
func (instance *ads_crossref) SetPropertyDS_msDS_DnsRootAlias(value string) (err error) {
	return instance.SetProperty("DS_msDS_DnsRootAlias", (value))
}

// GetDS_msDS_DnsRootAlias gets the value of DS_msDS_DnsRootAlias for the instance
func (instance *ads_crossref) GetPropertyDS_msDS_DnsRootAlias() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_DnsRootAlias")
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

// SetDS_msDS_NC_Replica_Locations sets the value of DS_msDS_NC_Replica_Locations for the instance
func (instance *ads_crossref) SetPropertyDS_msDS_NC_Replica_Locations(value []string) (err error) {
	return instance.SetProperty("DS_msDS_NC_Replica_Locations", (value))
}

// GetDS_msDS_NC_Replica_Locations gets the value of DS_msDS_NC_Replica_Locations for the instance
func (instance *ads_crossref) GetPropertyDS_msDS_NC_Replica_Locations() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_NC_Replica_Locations")
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

// SetDS_msDS_NC_RO_Replica_Locations sets the value of DS_msDS_NC_RO_Replica_Locations for the instance
func (instance *ads_crossref) SetPropertyDS_msDS_NC_RO_Replica_Locations(value []string) (err error) {
	return instance.SetProperty("DS_msDS_NC_RO_Replica_Locations", (value))
}

// GetDS_msDS_NC_RO_Replica_Locations gets the value of DS_msDS_NC_RO_Replica_Locations for the instance
func (instance *ads_crossref) GetPropertyDS_msDS_NC_RO_Replica_Locations() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_NC_RO_Replica_Locations")
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

// SetDS_msDS_Replication_Notify_First_DSA_Delay sets the value of DS_msDS_Replication_Notify_First_DSA_Delay for the instance
func (instance *ads_crossref) SetPropertyDS_msDS_Replication_Notify_First_DSA_Delay(value int32) (err error) {
	return instance.SetProperty("DS_msDS_Replication_Notify_First_DSA_Delay", (value))
}

// GetDS_msDS_Replication_Notify_First_DSA_Delay gets the value of DS_msDS_Replication_Notify_First_DSA_Delay for the instance
func (instance *ads_crossref) GetPropertyDS_msDS_Replication_Notify_First_DSA_Delay() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_Replication_Notify_First_DSA_Delay")
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

// SetDS_msDS_Replication_Notify_Subsequent_DSA_Delay sets the value of DS_msDS_Replication_Notify_Subsequent_DSA_Delay for the instance
func (instance *ads_crossref) SetPropertyDS_msDS_Replication_Notify_Subsequent_DSA_Delay(value int32) (err error) {
	return instance.SetProperty("DS_msDS_Replication_Notify_Subsequent_DSA_Delay", (value))
}

// GetDS_msDS_Replication_Notify_Subsequent_DSA_Delay gets the value of DS_msDS_Replication_Notify_Subsequent_DSA_Delay for the instance
func (instance *ads_crossref) GetPropertyDS_msDS_Replication_Notify_Subsequent_DSA_Delay() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_Replication_Notify_Subsequent_DSA_Delay")
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

// SetDS_msDS_SDReferenceDomain sets the value of DS_msDS_SDReferenceDomain for the instance
func (instance *ads_crossref) SetPropertyDS_msDS_SDReferenceDomain(value string) (err error) {
	return instance.SetProperty("DS_msDS_SDReferenceDomain", (value))
}

// GetDS_msDS_SDReferenceDomain gets the value of DS_msDS_SDReferenceDomain for the instance
func (instance *ads_crossref) GetPropertyDS_msDS_SDReferenceDomain() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_SDReferenceDomain")
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

// SetDS_nCName sets the value of DS_nCName for the instance
func (instance *ads_crossref) SetPropertyDS_nCName(value string) (err error) {
	return instance.SetProperty("DS_nCName", (value))
}

// GetDS_nCName gets the value of DS_nCName for the instance
func (instance *ads_crossref) GetPropertyDS_nCName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_nCName")
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

// SetDS_nETBIOSName sets the value of DS_nETBIOSName for the instance
func (instance *ads_crossref) SetPropertyDS_nETBIOSName(value string) (err error) {
	return instance.SetProperty("DS_nETBIOSName", (value))
}

// GetDS_nETBIOSName gets the value of DS_nETBIOSName for the instance
func (instance *ads_crossref) GetPropertyDS_nETBIOSName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_nETBIOSName")
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

// SetDS_nTMixedDomain sets the value of DS_nTMixedDomain for the instance
func (instance *ads_crossref) SetPropertyDS_nTMixedDomain(value int32) (err error) {
	return instance.SetProperty("DS_nTMixedDomain", (value))
}

// GetDS_nTMixedDomain gets the value of DS_nTMixedDomain for the instance
func (instance *ads_crossref) GetPropertyDS_nTMixedDomain() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_nTMixedDomain")
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

// SetDS_rootTrust sets the value of DS_rootTrust for the instance
func (instance *ads_crossref) SetPropertyDS_rootTrust(value []string) (err error) {
	return instance.SetProperty("DS_rootTrust", (value))
}

// GetDS_rootTrust gets the value of DS_rootTrust for the instance
func (instance *ads_crossref) GetPropertyDS_rootTrust() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_rootTrust")
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

// SetDS_superiorDNSRoot sets the value of DS_superiorDNSRoot for the instance
func (instance *ads_crossref) SetPropertyDS_superiorDNSRoot(value string) (err error) {
	return instance.SetProperty("DS_superiorDNSRoot", (value))
}

// GetDS_superiorDNSRoot gets the value of DS_superiorDNSRoot for the instance
func (instance *ads_crossref) GetPropertyDS_superiorDNSRoot() (value string, err error) {
	retValue, err := instance.GetProperty("DS_superiorDNSRoot")
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

// SetDS_trustParent sets the value of DS_trustParent for the instance
func (instance *ads_crossref) SetPropertyDS_trustParent(value string) (err error) {
	return instance.SetProperty("DS_trustParent", (value))
}

// GetDS_trustParent gets the value of DS_trustParent for the instance
func (instance *ads_crossref) GetPropertyDS_trustParent() (value string, err error) {
	retValue, err := instance.GetProperty("DS_trustParent")
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
