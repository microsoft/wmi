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

// ads_dnszone struct
type ads_dnszone struct {
	*ds_top

	//
	DS_dc string

	//
	DS_dnsAllowDynamic bool

	//
	DS_dnsAllowXFR bool

	//
	DS_dnsNotifySecondaries []int32

	//
	DS_dNSProperty []Uint8Array

	//
	DS_dnsSecureSecondaries []int32

	//
	DS_managedBy string

	//
	DS_msDNS_DNSKEYRecords []Uint8Array

	//
	DS_msDNS_DNSKEYRecordSetTTL int32

	//
	DS_msDNS_DSRecordAlgorithms int32

	//
	DS_msDNS_DSRecordSetTTL int32

	//
	DS_msDNS_IsSigned bool

	//
	DS_msDNS_MaintainTrustAnchor int32

	//
	DS_msDNS_NSEC3CurrentSalt string

	//
	DS_msDNS_NSEC3HashAlgorithm int32

	//
	DS_msDNS_NSEC3Iterations int32

	//
	DS_msDNS_NSEC3OptOut bool

	//
	DS_msDNS_NSEC3RandomSaltLength int32

	//
	DS_msDNS_NSEC3UserSalt string

	//
	DS_msDNS_ParentHasSecureDelegation bool

	//
	DS_msDNS_PropagationTime int32

	//
	DS_msDNS_RFC5011KeyRollovers bool

	//
	DS_msDNS_SecureDelegationPollingPeriod int32

	//
	DS_msDNS_SignatureInceptionOffset int32

	//
	DS_msDNS_SigningKeyDescriptors []Uint8Array

	//
	DS_msDNS_SigningKeys []Uint8Array

	//
	DS_msDNS_SignWithNSEC3 bool
}

func Newads_dnszoneEx1(instance *cim.WmiInstance) (newInstance *ads_dnszone, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_dnszone{
		ds_top: tmp,
	}
	return
}

func Newads_dnszoneEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_dnszone, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_dnszone{
		ds_top: tmp,
	}
	return
}

// SetDS_dc sets the value of DS_dc for the instance
func (instance *ads_dnszone) SetPropertyDS_dc(value string) (err error) {
	return instance.SetProperty("DS_dc", (value))
}

// GetDS_dc gets the value of DS_dc for the instance
func (instance *ads_dnszone) GetPropertyDS_dc() (value string, err error) {
	retValue, err := instance.GetProperty("DS_dc")
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

// SetDS_dnsAllowDynamic sets the value of DS_dnsAllowDynamic for the instance
func (instance *ads_dnszone) SetPropertyDS_dnsAllowDynamic(value bool) (err error) {
	return instance.SetProperty("DS_dnsAllowDynamic", (value))
}

// GetDS_dnsAllowDynamic gets the value of DS_dnsAllowDynamic for the instance
func (instance *ads_dnszone) GetPropertyDS_dnsAllowDynamic() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_dnsAllowDynamic")
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

// SetDS_dnsAllowXFR sets the value of DS_dnsAllowXFR for the instance
func (instance *ads_dnszone) SetPropertyDS_dnsAllowXFR(value bool) (err error) {
	return instance.SetProperty("DS_dnsAllowXFR", (value))
}

// GetDS_dnsAllowXFR gets the value of DS_dnsAllowXFR for the instance
func (instance *ads_dnszone) GetPropertyDS_dnsAllowXFR() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_dnsAllowXFR")
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

// SetDS_dnsNotifySecondaries sets the value of DS_dnsNotifySecondaries for the instance
func (instance *ads_dnszone) SetPropertyDS_dnsNotifySecondaries(value []int32) (err error) {
	return instance.SetProperty("DS_dnsNotifySecondaries", (value))
}

// GetDS_dnsNotifySecondaries gets the value of DS_dnsNotifySecondaries for the instance
func (instance *ads_dnszone) GetPropertyDS_dnsNotifySecondaries() (value []int32, err error) {
	retValue, err := instance.GetProperty("DS_dnsNotifySecondaries")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, int32(valuetmp))
	}

	return
}

// SetDS_dNSProperty sets the value of DS_dNSProperty for the instance
func (instance *ads_dnszone) SetPropertyDS_dNSProperty(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_dNSProperty", (value))
}

// GetDS_dNSProperty gets the value of DS_dNSProperty for the instance
func (instance *ads_dnszone) GetPropertyDS_dNSProperty() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_dNSProperty")
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

// SetDS_dnsSecureSecondaries sets the value of DS_dnsSecureSecondaries for the instance
func (instance *ads_dnszone) SetPropertyDS_dnsSecureSecondaries(value []int32) (err error) {
	return instance.SetProperty("DS_dnsSecureSecondaries", (value))
}

// GetDS_dnsSecureSecondaries gets the value of DS_dnsSecureSecondaries for the instance
func (instance *ads_dnszone) GetPropertyDS_dnsSecureSecondaries() (value []int32, err error) {
	retValue, err := instance.GetProperty("DS_dnsSecureSecondaries")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, int32(valuetmp))
	}

	return
}

// SetDS_managedBy sets the value of DS_managedBy for the instance
func (instance *ads_dnszone) SetPropertyDS_managedBy(value string) (err error) {
	return instance.SetProperty("DS_managedBy", (value))
}

// GetDS_managedBy gets the value of DS_managedBy for the instance
func (instance *ads_dnszone) GetPropertyDS_managedBy() (value string, err error) {
	retValue, err := instance.GetProperty("DS_managedBy")
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

// SetDS_msDNS_DNSKEYRecords sets the value of DS_msDNS_DNSKEYRecords for the instance
func (instance *ads_dnszone) SetPropertyDS_msDNS_DNSKEYRecords(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_msDNS_DNSKEYRecords", (value))
}

// GetDS_msDNS_DNSKEYRecords gets the value of DS_msDNS_DNSKEYRecords for the instance
func (instance *ads_dnszone) GetPropertyDS_msDNS_DNSKEYRecords() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDNS_DNSKEYRecords")
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

// SetDS_msDNS_DNSKEYRecordSetTTL sets the value of DS_msDNS_DNSKEYRecordSetTTL for the instance
func (instance *ads_dnszone) SetPropertyDS_msDNS_DNSKEYRecordSetTTL(value int32) (err error) {
	return instance.SetProperty("DS_msDNS_DNSKEYRecordSetTTL", (value))
}

// GetDS_msDNS_DNSKEYRecordSetTTL gets the value of DS_msDNS_DNSKEYRecordSetTTL for the instance
func (instance *ads_dnszone) GetPropertyDS_msDNS_DNSKEYRecordSetTTL() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDNS_DNSKEYRecordSetTTL")
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

// SetDS_msDNS_DSRecordAlgorithms sets the value of DS_msDNS_DSRecordAlgorithms for the instance
func (instance *ads_dnszone) SetPropertyDS_msDNS_DSRecordAlgorithms(value int32) (err error) {
	return instance.SetProperty("DS_msDNS_DSRecordAlgorithms", (value))
}

// GetDS_msDNS_DSRecordAlgorithms gets the value of DS_msDNS_DSRecordAlgorithms for the instance
func (instance *ads_dnszone) GetPropertyDS_msDNS_DSRecordAlgorithms() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDNS_DSRecordAlgorithms")
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

// SetDS_msDNS_DSRecordSetTTL sets the value of DS_msDNS_DSRecordSetTTL for the instance
func (instance *ads_dnszone) SetPropertyDS_msDNS_DSRecordSetTTL(value int32) (err error) {
	return instance.SetProperty("DS_msDNS_DSRecordSetTTL", (value))
}

// GetDS_msDNS_DSRecordSetTTL gets the value of DS_msDNS_DSRecordSetTTL for the instance
func (instance *ads_dnszone) GetPropertyDS_msDNS_DSRecordSetTTL() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDNS_DSRecordSetTTL")
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

// SetDS_msDNS_IsSigned sets the value of DS_msDNS_IsSigned for the instance
func (instance *ads_dnszone) SetPropertyDS_msDNS_IsSigned(value bool) (err error) {
	return instance.SetProperty("DS_msDNS_IsSigned", (value))
}

// GetDS_msDNS_IsSigned gets the value of DS_msDNS_IsSigned for the instance
func (instance *ads_dnszone) GetPropertyDS_msDNS_IsSigned() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_msDNS_IsSigned")
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

// SetDS_msDNS_MaintainTrustAnchor sets the value of DS_msDNS_MaintainTrustAnchor for the instance
func (instance *ads_dnszone) SetPropertyDS_msDNS_MaintainTrustAnchor(value int32) (err error) {
	return instance.SetProperty("DS_msDNS_MaintainTrustAnchor", (value))
}

// GetDS_msDNS_MaintainTrustAnchor gets the value of DS_msDNS_MaintainTrustAnchor for the instance
func (instance *ads_dnszone) GetPropertyDS_msDNS_MaintainTrustAnchor() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDNS_MaintainTrustAnchor")
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

// SetDS_msDNS_NSEC3CurrentSalt sets the value of DS_msDNS_NSEC3CurrentSalt for the instance
func (instance *ads_dnszone) SetPropertyDS_msDNS_NSEC3CurrentSalt(value string) (err error) {
	return instance.SetProperty("DS_msDNS_NSEC3CurrentSalt", (value))
}

// GetDS_msDNS_NSEC3CurrentSalt gets the value of DS_msDNS_NSEC3CurrentSalt for the instance
func (instance *ads_dnszone) GetPropertyDS_msDNS_NSEC3CurrentSalt() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDNS_NSEC3CurrentSalt")
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

// SetDS_msDNS_NSEC3HashAlgorithm sets the value of DS_msDNS_NSEC3HashAlgorithm for the instance
func (instance *ads_dnszone) SetPropertyDS_msDNS_NSEC3HashAlgorithm(value int32) (err error) {
	return instance.SetProperty("DS_msDNS_NSEC3HashAlgorithm", (value))
}

// GetDS_msDNS_NSEC3HashAlgorithm gets the value of DS_msDNS_NSEC3HashAlgorithm for the instance
func (instance *ads_dnszone) GetPropertyDS_msDNS_NSEC3HashAlgorithm() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDNS_NSEC3HashAlgorithm")
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

// SetDS_msDNS_NSEC3Iterations sets the value of DS_msDNS_NSEC3Iterations for the instance
func (instance *ads_dnszone) SetPropertyDS_msDNS_NSEC3Iterations(value int32) (err error) {
	return instance.SetProperty("DS_msDNS_NSEC3Iterations", (value))
}

// GetDS_msDNS_NSEC3Iterations gets the value of DS_msDNS_NSEC3Iterations for the instance
func (instance *ads_dnszone) GetPropertyDS_msDNS_NSEC3Iterations() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDNS_NSEC3Iterations")
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

// SetDS_msDNS_NSEC3OptOut sets the value of DS_msDNS_NSEC3OptOut for the instance
func (instance *ads_dnszone) SetPropertyDS_msDNS_NSEC3OptOut(value bool) (err error) {
	return instance.SetProperty("DS_msDNS_NSEC3OptOut", (value))
}

// GetDS_msDNS_NSEC3OptOut gets the value of DS_msDNS_NSEC3OptOut for the instance
func (instance *ads_dnszone) GetPropertyDS_msDNS_NSEC3OptOut() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_msDNS_NSEC3OptOut")
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

// SetDS_msDNS_NSEC3RandomSaltLength sets the value of DS_msDNS_NSEC3RandomSaltLength for the instance
func (instance *ads_dnszone) SetPropertyDS_msDNS_NSEC3RandomSaltLength(value int32) (err error) {
	return instance.SetProperty("DS_msDNS_NSEC3RandomSaltLength", (value))
}

// GetDS_msDNS_NSEC3RandomSaltLength gets the value of DS_msDNS_NSEC3RandomSaltLength for the instance
func (instance *ads_dnszone) GetPropertyDS_msDNS_NSEC3RandomSaltLength() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDNS_NSEC3RandomSaltLength")
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

// SetDS_msDNS_NSEC3UserSalt sets the value of DS_msDNS_NSEC3UserSalt for the instance
func (instance *ads_dnszone) SetPropertyDS_msDNS_NSEC3UserSalt(value string) (err error) {
	return instance.SetProperty("DS_msDNS_NSEC3UserSalt", (value))
}

// GetDS_msDNS_NSEC3UserSalt gets the value of DS_msDNS_NSEC3UserSalt for the instance
func (instance *ads_dnszone) GetPropertyDS_msDNS_NSEC3UserSalt() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDNS_NSEC3UserSalt")
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

// SetDS_msDNS_ParentHasSecureDelegation sets the value of DS_msDNS_ParentHasSecureDelegation for the instance
func (instance *ads_dnszone) SetPropertyDS_msDNS_ParentHasSecureDelegation(value bool) (err error) {
	return instance.SetProperty("DS_msDNS_ParentHasSecureDelegation", (value))
}

// GetDS_msDNS_ParentHasSecureDelegation gets the value of DS_msDNS_ParentHasSecureDelegation for the instance
func (instance *ads_dnszone) GetPropertyDS_msDNS_ParentHasSecureDelegation() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_msDNS_ParentHasSecureDelegation")
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

// SetDS_msDNS_PropagationTime sets the value of DS_msDNS_PropagationTime for the instance
func (instance *ads_dnszone) SetPropertyDS_msDNS_PropagationTime(value int32) (err error) {
	return instance.SetProperty("DS_msDNS_PropagationTime", (value))
}

// GetDS_msDNS_PropagationTime gets the value of DS_msDNS_PropagationTime for the instance
func (instance *ads_dnszone) GetPropertyDS_msDNS_PropagationTime() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDNS_PropagationTime")
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

// SetDS_msDNS_RFC5011KeyRollovers sets the value of DS_msDNS_RFC5011KeyRollovers for the instance
func (instance *ads_dnszone) SetPropertyDS_msDNS_RFC5011KeyRollovers(value bool) (err error) {
	return instance.SetProperty("DS_msDNS_RFC5011KeyRollovers", (value))
}

// GetDS_msDNS_RFC5011KeyRollovers gets the value of DS_msDNS_RFC5011KeyRollovers for the instance
func (instance *ads_dnszone) GetPropertyDS_msDNS_RFC5011KeyRollovers() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_msDNS_RFC5011KeyRollovers")
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

// SetDS_msDNS_SecureDelegationPollingPeriod sets the value of DS_msDNS_SecureDelegationPollingPeriod for the instance
func (instance *ads_dnszone) SetPropertyDS_msDNS_SecureDelegationPollingPeriod(value int32) (err error) {
	return instance.SetProperty("DS_msDNS_SecureDelegationPollingPeriod", (value))
}

// GetDS_msDNS_SecureDelegationPollingPeriod gets the value of DS_msDNS_SecureDelegationPollingPeriod for the instance
func (instance *ads_dnszone) GetPropertyDS_msDNS_SecureDelegationPollingPeriod() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDNS_SecureDelegationPollingPeriod")
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

// SetDS_msDNS_SignatureInceptionOffset sets the value of DS_msDNS_SignatureInceptionOffset for the instance
func (instance *ads_dnszone) SetPropertyDS_msDNS_SignatureInceptionOffset(value int32) (err error) {
	return instance.SetProperty("DS_msDNS_SignatureInceptionOffset", (value))
}

// GetDS_msDNS_SignatureInceptionOffset gets the value of DS_msDNS_SignatureInceptionOffset for the instance
func (instance *ads_dnszone) GetPropertyDS_msDNS_SignatureInceptionOffset() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDNS_SignatureInceptionOffset")
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

// SetDS_msDNS_SigningKeyDescriptors sets the value of DS_msDNS_SigningKeyDescriptors for the instance
func (instance *ads_dnszone) SetPropertyDS_msDNS_SigningKeyDescriptors(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_msDNS_SigningKeyDescriptors", (value))
}

// GetDS_msDNS_SigningKeyDescriptors gets the value of DS_msDNS_SigningKeyDescriptors for the instance
func (instance *ads_dnszone) GetPropertyDS_msDNS_SigningKeyDescriptors() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDNS_SigningKeyDescriptors")
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

// SetDS_msDNS_SigningKeys sets the value of DS_msDNS_SigningKeys for the instance
func (instance *ads_dnszone) SetPropertyDS_msDNS_SigningKeys(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_msDNS_SigningKeys", (value))
}

// GetDS_msDNS_SigningKeys gets the value of DS_msDNS_SigningKeys for the instance
func (instance *ads_dnszone) GetPropertyDS_msDNS_SigningKeys() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDNS_SigningKeys")
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

// SetDS_msDNS_SignWithNSEC3 sets the value of DS_msDNS_SignWithNSEC3 for the instance
func (instance *ads_dnszone) SetPropertyDS_msDNS_SignWithNSEC3(value bool) (err error) {
	return instance.SetProperty("DS_msDNS_SignWithNSEC3", (value))
}

// GetDS_msDNS_SignWithNSEC3 gets the value of DS_msDNS_SignWithNSEC3 for the instance
func (instance *ads_dnszone) GetPropertyDS_msDNS_SignWithNSEC3() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_msDNS_SignWithNSEC3")
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
