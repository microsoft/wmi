// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.RemoteAccess
//
// ////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// BgpRouteAggregateConfig struct
type BgpRouteAggregateConfig struct {
	*cim.WmiInstance

	//
	AttributePolicy []string

	//
	Prefix string

	//
	PreserveASPath uint32

	//
	RoutingDomain string

	//
	SummaryOnly uint32
}

func NewBgpRouteAggregateConfigEx1(instance *cim.WmiInstance) (newInstance *BgpRouteAggregateConfig, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &BgpRouteAggregateConfig{
		WmiInstance: tmp,
	}
	return
}

func NewBgpRouteAggregateConfigEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *BgpRouteAggregateConfig, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &BgpRouteAggregateConfig{
		WmiInstance: tmp,
	}
	return
}

// SetAttributePolicy sets the value of AttributePolicy for the instance
func (instance *BgpRouteAggregateConfig) SetPropertyAttributePolicy(value []string) (err error) {
	return instance.SetProperty("AttributePolicy", (value))
}

// GetAttributePolicy gets the value of AttributePolicy for the instance
func (instance *BgpRouteAggregateConfig) GetPropertyAttributePolicy() (value []string, err error) {
	retValue, err := instance.GetProperty("AttributePolicy")
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

// SetPrefix sets the value of Prefix for the instance
func (instance *BgpRouteAggregateConfig) SetPropertyPrefix(value string) (err error) {
	return instance.SetProperty("Prefix", (value))
}

// GetPrefix gets the value of Prefix for the instance
func (instance *BgpRouteAggregateConfig) GetPropertyPrefix() (value string, err error) {
	retValue, err := instance.GetProperty("Prefix")
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

// SetPreserveASPath sets the value of PreserveASPath for the instance
func (instance *BgpRouteAggregateConfig) SetPropertyPreserveASPath(value uint32) (err error) {
	return instance.SetProperty("PreserveASPath", (value))
}

// GetPreserveASPath gets the value of PreserveASPath for the instance
func (instance *BgpRouteAggregateConfig) GetPropertyPreserveASPath() (value uint32, err error) {
	retValue, err := instance.GetProperty("PreserveASPath")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetRoutingDomain sets the value of RoutingDomain for the instance
func (instance *BgpRouteAggregateConfig) SetPropertyRoutingDomain(value string) (err error) {
	return instance.SetProperty("RoutingDomain", (value))
}

// GetRoutingDomain gets the value of RoutingDomain for the instance
func (instance *BgpRouteAggregateConfig) GetPropertyRoutingDomain() (value string, err error) {
	retValue, err := instance.GetProperty("RoutingDomain")
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

// SetSummaryOnly sets the value of SummaryOnly for the instance
func (instance *BgpRouteAggregateConfig) SetPropertySummaryOnly(value uint32) (err error) {
	return instance.SetProperty("SummaryOnly", (value))
}

// GetSummaryOnly gets the value of SummaryOnly for the instance
func (instance *BgpRouteAggregateConfig) GetPropertySummaryOnly() (value uint32, err error) {
	retValue, err := instance.GetProperty("SummaryOnly")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}
