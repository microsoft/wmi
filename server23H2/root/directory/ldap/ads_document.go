// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ads_document struct
type ads_document struct {
	*ds_top

	//
	DS_documentAuthor []string

	//
	DS_documentIdentifier []string

	//
	DS_documentLocation []string

	//
	DS_documentPublisher []string

	//
	DS_documentTitle []string

	//
	DS_documentVersion []string

	//
	DS_l string

	//
	DS_o []string

	//
	DS_ou []string

	//
	DS_seeAlso []string
}

func Newads_documentEx1(instance *cim.WmiInstance) (newInstance *ads_document, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_document{
		ds_top: tmp,
	}
	return
}

func Newads_documentEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_document, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_document{
		ds_top: tmp,
	}
	return
}

// SetDS_documentAuthor sets the value of DS_documentAuthor for the instance
func (instance *ads_document) SetPropertyDS_documentAuthor(value []string) (err error) {
	return instance.SetProperty("DS_documentAuthor", (value))
}

// GetDS_documentAuthor gets the value of DS_documentAuthor for the instance
func (instance *ads_document) GetPropertyDS_documentAuthor() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_documentAuthor")
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

// SetDS_documentIdentifier sets the value of DS_documentIdentifier for the instance
func (instance *ads_document) SetPropertyDS_documentIdentifier(value []string) (err error) {
	return instance.SetProperty("DS_documentIdentifier", (value))
}

// GetDS_documentIdentifier gets the value of DS_documentIdentifier for the instance
func (instance *ads_document) GetPropertyDS_documentIdentifier() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_documentIdentifier")
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

// SetDS_documentLocation sets the value of DS_documentLocation for the instance
func (instance *ads_document) SetPropertyDS_documentLocation(value []string) (err error) {
	return instance.SetProperty("DS_documentLocation", (value))
}

// GetDS_documentLocation gets the value of DS_documentLocation for the instance
func (instance *ads_document) GetPropertyDS_documentLocation() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_documentLocation")
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

// SetDS_documentPublisher sets the value of DS_documentPublisher for the instance
func (instance *ads_document) SetPropertyDS_documentPublisher(value []string) (err error) {
	return instance.SetProperty("DS_documentPublisher", (value))
}

// GetDS_documentPublisher gets the value of DS_documentPublisher for the instance
func (instance *ads_document) GetPropertyDS_documentPublisher() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_documentPublisher")
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

// SetDS_documentTitle sets the value of DS_documentTitle for the instance
func (instance *ads_document) SetPropertyDS_documentTitle(value []string) (err error) {
	return instance.SetProperty("DS_documentTitle", (value))
}

// GetDS_documentTitle gets the value of DS_documentTitle for the instance
func (instance *ads_document) GetPropertyDS_documentTitle() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_documentTitle")
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

// SetDS_documentVersion sets the value of DS_documentVersion for the instance
func (instance *ads_document) SetPropertyDS_documentVersion(value []string) (err error) {
	return instance.SetProperty("DS_documentVersion", (value))
}

// GetDS_documentVersion gets the value of DS_documentVersion for the instance
func (instance *ads_document) GetPropertyDS_documentVersion() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_documentVersion")
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

// SetDS_l sets the value of DS_l for the instance
func (instance *ads_document) SetPropertyDS_l(value string) (err error) {
	return instance.SetProperty("DS_l", (value))
}

// GetDS_l gets the value of DS_l for the instance
func (instance *ads_document) GetPropertyDS_l() (value string, err error) {
	retValue, err := instance.GetProperty("DS_l")
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

// SetDS_o sets the value of DS_o for the instance
func (instance *ads_document) SetPropertyDS_o(value []string) (err error) {
	return instance.SetProperty("DS_o", (value))
}

// GetDS_o gets the value of DS_o for the instance
func (instance *ads_document) GetPropertyDS_o() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_o")
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

// SetDS_ou sets the value of DS_ou for the instance
func (instance *ads_document) SetPropertyDS_ou(value []string) (err error) {
	return instance.SetProperty("DS_ou", (value))
}

// GetDS_ou gets the value of DS_ou for the instance
func (instance *ads_document) GetPropertyDS_ou() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_ou")
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

// SetDS_seeAlso sets the value of DS_seeAlso for the instance
func (instance *ads_document) SetPropertyDS_seeAlso(value []string) (err error) {
	return instance.SetProperty("DS_seeAlso", (value))
}

// GetDS_seeAlso gets the value of DS_seeAlso for the instance
func (instance *ads_document) GetPropertyDS_seeAlso() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_seeAlso")
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
