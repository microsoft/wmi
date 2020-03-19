// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.HomeNet
//////////////////////////////////////////////
package homenet

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __EventConsumer struct
type __EventConsumer struct {
	*__IndicationRelated

	//
	CreatorSID []uint8

	//
	MachineName string

	//
	MaximumQueueSize uint32
}

func New__EventConsumerEx1(instance *cim.WmiInstance) (newInstance *__EventConsumer, err error) {
	tmp, err := New__IndicationRelatedEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__EventConsumer{
		__IndicationRelated: tmp,
	}
	return
}

func New__EventConsumerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__EventConsumer, err error) {
	tmp, err := New__IndicationRelatedEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__EventConsumer{
		__IndicationRelated: tmp,
	}
	return
}

// SetCreatorSID sets the value of CreatorSID for the instance
func (instance *__EventConsumer) SetPropertyCreatorSID(value []uint8) (err error) {
	return instance.SetProperty("CreatorSID", value)
}

// GetCreatorSID gets the value of CreatorSID for the instance
func (instance *__EventConsumer) GetPropertyCreatorSID() (value []uint8, err error) {
	retValue, err := instance.GetProperty("CreatorSID")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMachineName sets the value of MachineName for the instance
func (instance *__EventConsumer) SetPropertyMachineName(value string) (err error) {
	return instance.SetProperty("MachineName", value)
}

// GetMachineName gets the value of MachineName for the instance
func (instance *__EventConsumer) GetPropertyMachineName() (value string, err error) {
	retValue, err := instance.GetProperty("MachineName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaximumQueueSize sets the value of MaximumQueueSize for the instance
func (instance *__EventConsumer) SetPropertyMaximumQueueSize(value uint32) (err error) {
	return instance.SetProperty("MaximumQueueSize", value)
}

// GetMaximumQueueSize gets the value of MaximumQueueSize for the instance
func (instance *__EventConsumer) GetPropertyMaximumQueueSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumQueueSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
