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

// ads_msmq_custom_recipient struct
type ads_msmq_custom_recipient struct {
	*ds_top

	//
	DS_msMQ_Recipient_FormatName string
}

func Newads_msmq_custom_recipientEx1(instance *cim.WmiInstance) (newInstance *ads_msmq_custom_recipient, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msmq_custom_recipient{
		ds_top: tmp,
	}
	return
}

func Newads_msmq_custom_recipientEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msmq_custom_recipient, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msmq_custom_recipient{
		ds_top: tmp,
	}
	return
}

// SetDS_msMQ_Recipient_FormatName sets the value of DS_msMQ_Recipient_FormatName for the instance
func (instance *ads_msmq_custom_recipient) SetPropertyDS_msMQ_Recipient_FormatName(value string) (err error) {
	return instance.SetProperty("DS_msMQ_Recipient_FormatName", (value))
}

// GetDS_msMQ_Recipient_FormatName gets the value of DS_msMQ_Recipient_FormatName for the instance
func (instance *ads_msmq_custom_recipient) GetPropertyDS_msMQ_Recipient_FormatName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msMQ_Recipient_FormatName")
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
