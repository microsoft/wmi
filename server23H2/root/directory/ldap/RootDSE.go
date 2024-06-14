// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RootDSE struct
type RootDSE struct {
	*cim.WmiInstance

	//
	configurationNamingContext string

	//
	currentTime string

	//
	defaultNamingContext string

	//
	dnsHostName string

	//
	dsServiceName string

	//
	highestCommittedUSN string

	//
	LDAPServiceName string

	//
	namingContexts []string

	//
	rootDomainNamingContext string

	//
	schemaNamingContext string

	//
	serverName string

	//
	subschemaSubentry string

	//
	supportedCapabilities string

	//
	supportedControl []string

	//
	supportedLDAPPolicies []string

	//
	supportedLDAPVersion []string

	//
	supportedSASLMechanisms []string
}

func NewRootDSEEx1(instance *cim.WmiInstance) (newInstance *RootDSE, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RootDSE{
		WmiInstance: tmp,
	}
	return
}

func NewRootDSEEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RootDSE, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RootDSE{
		WmiInstance: tmp,
	}
	return
}

// SetconfigurationNamingContext sets the value of configurationNamingContext for the instance
func (instance *RootDSE) SetPropertyconfigurationNamingContext(value string) (err error) {
	return instance.SetProperty("configurationNamingContext", (value))
}

// GetconfigurationNamingContext gets the value of configurationNamingContext for the instance
func (instance *RootDSE) GetPropertyconfigurationNamingContext() (value string, err error) {
	retValue, err := instance.GetProperty("configurationNamingContext")
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

// SetcurrentTime sets the value of currentTime for the instance
func (instance *RootDSE) SetPropertycurrentTime(value string) (err error) {
	return instance.SetProperty("currentTime", (value))
}

// GetcurrentTime gets the value of currentTime for the instance
func (instance *RootDSE) GetPropertycurrentTime() (value string, err error) {
	retValue, err := instance.GetProperty("currentTime")
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

// SetdefaultNamingContext sets the value of defaultNamingContext for the instance
func (instance *RootDSE) SetPropertydefaultNamingContext(value string) (err error) {
	return instance.SetProperty("defaultNamingContext", (value))
}

// GetdefaultNamingContext gets the value of defaultNamingContext for the instance
func (instance *RootDSE) GetPropertydefaultNamingContext() (value string, err error) {
	retValue, err := instance.GetProperty("defaultNamingContext")
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

// SetdnsHostName sets the value of dnsHostName for the instance
func (instance *RootDSE) SetPropertydnsHostName(value string) (err error) {
	return instance.SetProperty("dnsHostName", (value))
}

// GetdnsHostName gets the value of dnsHostName for the instance
func (instance *RootDSE) GetPropertydnsHostName() (value string, err error) {
	retValue, err := instance.GetProperty("dnsHostName")
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

// SetdsServiceName sets the value of dsServiceName for the instance
func (instance *RootDSE) SetPropertydsServiceName(value string) (err error) {
	return instance.SetProperty("dsServiceName", (value))
}

// GetdsServiceName gets the value of dsServiceName for the instance
func (instance *RootDSE) GetPropertydsServiceName() (value string, err error) {
	retValue, err := instance.GetProperty("dsServiceName")
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

// SethighestCommittedUSN sets the value of highestCommittedUSN for the instance
func (instance *RootDSE) SetPropertyhighestCommittedUSN(value string) (err error) {
	return instance.SetProperty("highestCommittedUSN", (value))
}

// GethighestCommittedUSN gets the value of highestCommittedUSN for the instance
func (instance *RootDSE) GetPropertyhighestCommittedUSN() (value string, err error) {
	retValue, err := instance.GetProperty("highestCommittedUSN")
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

// SetLDAPServiceName sets the value of LDAPServiceName for the instance
func (instance *RootDSE) SetPropertyLDAPServiceName(value string) (err error) {
	return instance.SetProperty("LDAPServiceName", (value))
}

// GetLDAPServiceName gets the value of LDAPServiceName for the instance
func (instance *RootDSE) GetPropertyLDAPServiceName() (value string, err error) {
	retValue, err := instance.GetProperty("LDAPServiceName")
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

// SetnamingContexts sets the value of namingContexts for the instance
func (instance *RootDSE) SetPropertynamingContexts(value []string) (err error) {
	return instance.SetProperty("namingContexts", (value))
}

// GetnamingContexts gets the value of namingContexts for the instance
func (instance *RootDSE) GetPropertynamingContexts() (value []string, err error) {
	retValue, err := instance.GetProperty("namingContexts")
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

// SetrootDomainNamingContext sets the value of rootDomainNamingContext for the instance
func (instance *RootDSE) SetPropertyrootDomainNamingContext(value string) (err error) {
	return instance.SetProperty("rootDomainNamingContext", (value))
}

// GetrootDomainNamingContext gets the value of rootDomainNamingContext for the instance
func (instance *RootDSE) GetPropertyrootDomainNamingContext() (value string, err error) {
	retValue, err := instance.GetProperty("rootDomainNamingContext")
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

// SetschemaNamingContext sets the value of schemaNamingContext for the instance
func (instance *RootDSE) SetPropertyschemaNamingContext(value string) (err error) {
	return instance.SetProperty("schemaNamingContext", (value))
}

// GetschemaNamingContext gets the value of schemaNamingContext for the instance
func (instance *RootDSE) GetPropertyschemaNamingContext() (value string, err error) {
	retValue, err := instance.GetProperty("schemaNamingContext")
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

// SetserverName sets the value of serverName for the instance
func (instance *RootDSE) SetPropertyserverName(value string) (err error) {
	return instance.SetProperty("serverName", (value))
}

// GetserverName gets the value of serverName for the instance
func (instance *RootDSE) GetPropertyserverName() (value string, err error) {
	retValue, err := instance.GetProperty("serverName")
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

// SetsubschemaSubentry sets the value of subschemaSubentry for the instance
func (instance *RootDSE) SetPropertysubschemaSubentry(value string) (err error) {
	return instance.SetProperty("subschemaSubentry", (value))
}

// GetsubschemaSubentry gets the value of subschemaSubentry for the instance
func (instance *RootDSE) GetPropertysubschemaSubentry() (value string, err error) {
	retValue, err := instance.GetProperty("subschemaSubentry")
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

// SetsupportedCapabilities sets the value of supportedCapabilities for the instance
func (instance *RootDSE) SetPropertysupportedCapabilities(value string) (err error) {
	return instance.SetProperty("supportedCapabilities", (value))
}

// GetsupportedCapabilities gets the value of supportedCapabilities for the instance
func (instance *RootDSE) GetPropertysupportedCapabilities() (value string, err error) {
	retValue, err := instance.GetProperty("supportedCapabilities")
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

// SetsupportedControl sets the value of supportedControl for the instance
func (instance *RootDSE) SetPropertysupportedControl(value []string) (err error) {
	return instance.SetProperty("supportedControl", (value))
}

// GetsupportedControl gets the value of supportedControl for the instance
func (instance *RootDSE) GetPropertysupportedControl() (value []string, err error) {
	retValue, err := instance.GetProperty("supportedControl")
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

// SetsupportedLDAPPolicies sets the value of supportedLDAPPolicies for the instance
func (instance *RootDSE) SetPropertysupportedLDAPPolicies(value []string) (err error) {
	return instance.SetProperty("supportedLDAPPolicies", (value))
}

// GetsupportedLDAPPolicies gets the value of supportedLDAPPolicies for the instance
func (instance *RootDSE) GetPropertysupportedLDAPPolicies() (value []string, err error) {
	retValue, err := instance.GetProperty("supportedLDAPPolicies")
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

// SetsupportedLDAPVersion sets the value of supportedLDAPVersion for the instance
func (instance *RootDSE) SetPropertysupportedLDAPVersion(value []string) (err error) {
	return instance.SetProperty("supportedLDAPVersion", (value))
}

// GetsupportedLDAPVersion gets the value of supportedLDAPVersion for the instance
func (instance *RootDSE) GetPropertysupportedLDAPVersion() (value []string, err error) {
	retValue, err := instance.GetProperty("supportedLDAPVersion")
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

// SetsupportedSASLMechanisms sets the value of supportedSASLMechanisms for the instance
func (instance *RootDSE) SetPropertysupportedSASLMechanisms(value []string) (err error) {
	return instance.SetProperty("supportedSASLMechanisms", (value))
}

// GetsupportedSASLMechanisms gets the value of supportedSASLMechanisms for the instance
func (instance *RootDSE) GetPropertysupportedSASLMechanisms() (value []string, err error) {
	retValue, err := instance.GetProperty("supportedSASLMechanisms")
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
