// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ServiceSecurityAuditBehavior struct
type ServiceSecurityAuditBehavior struct {
	*Behavior

	// The location of the audit log.
	AuditLogLocation string

	// The type of message authentication level that is used to log audit events.
	MessageAuthenticationAuditLevel string

	// The types of authorization events that are recorded in the audit log.
	ServiceAuthorizationAuditLevel string

	// A boolean value that specifies the behavior for suppressing failures of writing to the audit log.
	SuppressAuditFailure bool
}

func NewServiceSecurityAuditBehaviorEx1(instance *cim.WmiInstance) (newInstance *ServiceSecurityAuditBehavior, err error) {
	tmp, err := NewBehaviorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ServiceSecurityAuditBehavior{
		Behavior: tmp,
	}
	return
}

func NewServiceSecurityAuditBehaviorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ServiceSecurityAuditBehavior, err error) {
	tmp, err := NewBehaviorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ServiceSecurityAuditBehavior{
		Behavior: tmp,
	}
	return
}

// SetAuditLogLocation sets the value of AuditLogLocation for the instance
func (instance *ServiceSecurityAuditBehavior) SetPropertyAuditLogLocation(value string) (err error) {
	return instance.SetProperty("AuditLogLocation", (value))
}

// GetAuditLogLocation gets the value of AuditLogLocation for the instance
func (instance *ServiceSecurityAuditBehavior) GetPropertyAuditLogLocation() (value string, err error) {
	retValue, err := instance.GetProperty("AuditLogLocation")
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

// SetMessageAuthenticationAuditLevel sets the value of MessageAuthenticationAuditLevel for the instance
func (instance *ServiceSecurityAuditBehavior) SetPropertyMessageAuthenticationAuditLevel(value string) (err error) {
	return instance.SetProperty("MessageAuthenticationAuditLevel", (value))
}

// GetMessageAuthenticationAuditLevel gets the value of MessageAuthenticationAuditLevel for the instance
func (instance *ServiceSecurityAuditBehavior) GetPropertyMessageAuthenticationAuditLevel() (value string, err error) {
	retValue, err := instance.GetProperty("MessageAuthenticationAuditLevel")
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

// SetServiceAuthorizationAuditLevel sets the value of ServiceAuthorizationAuditLevel for the instance
func (instance *ServiceSecurityAuditBehavior) SetPropertyServiceAuthorizationAuditLevel(value string) (err error) {
	return instance.SetProperty("ServiceAuthorizationAuditLevel", (value))
}

// GetServiceAuthorizationAuditLevel gets the value of ServiceAuthorizationAuditLevel for the instance
func (instance *ServiceSecurityAuditBehavior) GetPropertyServiceAuthorizationAuditLevel() (value string, err error) {
	retValue, err := instance.GetProperty("ServiceAuthorizationAuditLevel")
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

// SetSuppressAuditFailure sets the value of SuppressAuditFailure for the instance
func (instance *ServiceSecurityAuditBehavior) SetPropertySuppressAuditFailure(value bool) (err error) {
	return instance.SetProperty("SuppressAuditFailure", (value))
}

// GetSuppressAuditFailure gets the value of SuppressAuditFailure for the instance
func (instance *ServiceSecurityAuditBehavior) GetPropertySuppressAuditFailure() (value bool, err error) {
	retValue, err := instance.GetProperty("SuppressAuditFailure")
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
