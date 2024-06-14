// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess
//////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// BgpRoutingPolicyConfig struct
type BgpRoutingPolicyConfig struct {
	*cim.WmiInstance

	//
	AddCommunity []string

	//
	ClearMED bool

	//
	IgnorePrefix []string

	//
	MatchASNRange []uint32

	//
	MatchCommunity []string

	//
	MatchNextHop []string

	//
	MatchPrefix []string

	//
	NewLocalPref uint32

	//
	NewMED uint32

	//
	NewNextHop string

	//
	PolicyName string

	//
	PolicyType uint32

	//
	RemoveAllCommunities bool

	//
	RemoveCommunity []string

	//
	RoutingDomain string
}

func NewBgpRoutingPolicyConfigEx1(instance *cim.WmiInstance) (newInstance *BgpRoutingPolicyConfig, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &BgpRoutingPolicyConfig{
		WmiInstance: tmp,
	}
	return
}

func NewBgpRoutingPolicyConfigEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *BgpRoutingPolicyConfig, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &BgpRoutingPolicyConfig{
		WmiInstance: tmp,
	}
	return
}

// SetAddCommunity sets the value of AddCommunity for the instance
func (instance *BgpRoutingPolicyConfig) SetPropertyAddCommunity(value []string) (err error) {
	return instance.SetProperty("AddCommunity", (value))
}

// GetAddCommunity gets the value of AddCommunity for the instance
func (instance *BgpRoutingPolicyConfig) GetPropertyAddCommunity() (value []string, err error) {
	retValue, err := instance.GetProperty("AddCommunity")
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

// SetClearMED sets the value of ClearMED for the instance
func (instance *BgpRoutingPolicyConfig) SetPropertyClearMED(value bool) (err error) {
	return instance.SetProperty("ClearMED", (value))
}

// GetClearMED gets the value of ClearMED for the instance
func (instance *BgpRoutingPolicyConfig) GetPropertyClearMED() (value bool, err error) {
	retValue, err := instance.GetProperty("ClearMED")
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

// SetIgnorePrefix sets the value of IgnorePrefix for the instance
func (instance *BgpRoutingPolicyConfig) SetPropertyIgnorePrefix(value []string) (err error) {
	return instance.SetProperty("IgnorePrefix", (value))
}

// GetIgnorePrefix gets the value of IgnorePrefix for the instance
func (instance *BgpRoutingPolicyConfig) GetPropertyIgnorePrefix() (value []string, err error) {
	retValue, err := instance.GetProperty("IgnorePrefix")
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

// SetMatchASNRange sets the value of MatchASNRange for the instance
func (instance *BgpRoutingPolicyConfig) SetPropertyMatchASNRange(value []uint32) (err error) {
	return instance.SetProperty("MatchASNRange", (value))
}

// GetMatchASNRange gets the value of MatchASNRange for the instance
func (instance *BgpRoutingPolicyConfig) GetPropertyMatchASNRange() (value []uint32, err error) {
	retValue, err := instance.GetProperty("MatchASNRange")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint32(valuetmp))
	}

	return
}

// SetMatchCommunity sets the value of MatchCommunity for the instance
func (instance *BgpRoutingPolicyConfig) SetPropertyMatchCommunity(value []string) (err error) {
	return instance.SetProperty("MatchCommunity", (value))
}

// GetMatchCommunity gets the value of MatchCommunity for the instance
func (instance *BgpRoutingPolicyConfig) GetPropertyMatchCommunity() (value []string, err error) {
	retValue, err := instance.GetProperty("MatchCommunity")
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

// SetMatchNextHop sets the value of MatchNextHop for the instance
func (instance *BgpRoutingPolicyConfig) SetPropertyMatchNextHop(value []string) (err error) {
	return instance.SetProperty("MatchNextHop", (value))
}

// GetMatchNextHop gets the value of MatchNextHop for the instance
func (instance *BgpRoutingPolicyConfig) GetPropertyMatchNextHop() (value []string, err error) {
	retValue, err := instance.GetProperty("MatchNextHop")
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

// SetMatchPrefix sets the value of MatchPrefix for the instance
func (instance *BgpRoutingPolicyConfig) SetPropertyMatchPrefix(value []string) (err error) {
	return instance.SetProperty("MatchPrefix", (value))
}

// GetMatchPrefix gets the value of MatchPrefix for the instance
func (instance *BgpRoutingPolicyConfig) GetPropertyMatchPrefix() (value []string, err error) {
	retValue, err := instance.GetProperty("MatchPrefix")
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

// SetNewLocalPref sets the value of NewLocalPref for the instance
func (instance *BgpRoutingPolicyConfig) SetPropertyNewLocalPref(value uint32) (err error) {
	return instance.SetProperty("NewLocalPref", (value))
}

// GetNewLocalPref gets the value of NewLocalPref for the instance
func (instance *BgpRoutingPolicyConfig) GetPropertyNewLocalPref() (value uint32, err error) {
	retValue, err := instance.GetProperty("NewLocalPref")
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

// SetNewMED sets the value of NewMED for the instance
func (instance *BgpRoutingPolicyConfig) SetPropertyNewMED(value uint32) (err error) {
	return instance.SetProperty("NewMED", (value))
}

// GetNewMED gets the value of NewMED for the instance
func (instance *BgpRoutingPolicyConfig) GetPropertyNewMED() (value uint32, err error) {
	retValue, err := instance.GetProperty("NewMED")
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

// SetNewNextHop sets the value of NewNextHop for the instance
func (instance *BgpRoutingPolicyConfig) SetPropertyNewNextHop(value string) (err error) {
	return instance.SetProperty("NewNextHop", (value))
}

// GetNewNextHop gets the value of NewNextHop for the instance
func (instance *BgpRoutingPolicyConfig) GetPropertyNewNextHop() (value string, err error) {
	retValue, err := instance.GetProperty("NewNextHop")
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

// SetPolicyName sets the value of PolicyName for the instance
func (instance *BgpRoutingPolicyConfig) SetPropertyPolicyName(value string) (err error) {
	return instance.SetProperty("PolicyName", (value))
}

// GetPolicyName gets the value of PolicyName for the instance
func (instance *BgpRoutingPolicyConfig) GetPropertyPolicyName() (value string, err error) {
	retValue, err := instance.GetProperty("PolicyName")
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

// SetPolicyType sets the value of PolicyType for the instance
func (instance *BgpRoutingPolicyConfig) SetPropertyPolicyType(value uint32) (err error) {
	return instance.SetProperty("PolicyType", (value))
}

// GetPolicyType gets the value of PolicyType for the instance
func (instance *BgpRoutingPolicyConfig) GetPropertyPolicyType() (value uint32, err error) {
	retValue, err := instance.GetProperty("PolicyType")
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

// SetRemoveAllCommunities sets the value of RemoveAllCommunities for the instance
func (instance *BgpRoutingPolicyConfig) SetPropertyRemoveAllCommunities(value bool) (err error) {
	return instance.SetProperty("RemoveAllCommunities", (value))
}

// GetRemoveAllCommunities gets the value of RemoveAllCommunities for the instance
func (instance *BgpRoutingPolicyConfig) GetPropertyRemoveAllCommunities() (value bool, err error) {
	retValue, err := instance.GetProperty("RemoveAllCommunities")
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

// SetRemoveCommunity sets the value of RemoveCommunity for the instance
func (instance *BgpRoutingPolicyConfig) SetPropertyRemoveCommunity(value []string) (err error) {
	return instance.SetProperty("RemoveCommunity", (value))
}

// GetRemoveCommunity gets the value of RemoveCommunity for the instance
func (instance *BgpRoutingPolicyConfig) GetPropertyRemoveCommunity() (value []string, err error) {
	retValue, err := instance.GetProperty("RemoveCommunity")
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

// SetRoutingDomain sets the value of RoutingDomain for the instance
func (instance *BgpRoutingPolicyConfig) SetPropertyRoutingDomain(value string) (err error) {
	return instance.SetProperty("RoutingDomain", (value))
}

// GetRoutingDomain gets the value of RoutingDomain for the instance
func (instance *BgpRoutingPolicyConfig) GetPropertyRoutingDomain() (value string, err error) {
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
