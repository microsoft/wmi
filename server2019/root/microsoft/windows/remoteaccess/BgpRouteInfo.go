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

// BgpRouteInfo struct
type BgpRouteInfo struct {
	*cim.WmiInstance

	//
	Aggregate bool

	//
	Aggregator string

	//
	ASPath []ASPathSegment

	//
	ClusterList []uint32

	//
	Community []string

	//
	LearnedFromPeer string

	//
	LocalPref uint32

	//
	MED uint32

	//
	Network string

	//
	NextHop string

	//
	Origin string

	//
	OriginatorId string

	//
	RoutingDomain string

	//
	State []uint32
}

func NewBgpRouteInfoEx1(instance *cim.WmiInstance) (newInstance *BgpRouteInfo, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &BgpRouteInfo{
		WmiInstance: tmp,
	}
	return
}

func NewBgpRouteInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *BgpRouteInfo, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &BgpRouteInfo{
		WmiInstance: tmp,
	}
	return
}

// SetAggregate sets the value of Aggregate for the instance
func (instance *BgpRouteInfo) SetPropertyAggregate(value bool) (err error) {
	return instance.SetProperty("Aggregate", (value))
}

// GetAggregate gets the value of Aggregate for the instance
func (instance *BgpRouteInfo) GetPropertyAggregate() (value bool, err error) {
	retValue, err := instance.GetProperty("Aggregate")
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

// SetAggregator sets the value of Aggregator for the instance
func (instance *BgpRouteInfo) SetPropertyAggregator(value string) (err error) {
	return instance.SetProperty("Aggregator", (value))
}

// GetAggregator gets the value of Aggregator for the instance
func (instance *BgpRouteInfo) GetPropertyAggregator() (value string, err error) {
	retValue, err := instance.GetProperty("Aggregator")
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

// SetASPath sets the value of ASPath for the instance
func (instance *BgpRouteInfo) SetPropertyASPath(value []ASPathSegment) (err error) {
	return instance.SetProperty("ASPath", (value))
}

// GetASPath gets the value of ASPath for the instance
func (instance *BgpRouteInfo) GetPropertyASPath() (value []ASPathSegment, err error) {
	retValue, err := instance.GetProperty("ASPath")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(ASPathSegment)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " ASPathSegment is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, ASPathSegment(valuetmp))
	}

	return
}

// SetClusterList sets the value of ClusterList for the instance
func (instance *BgpRouteInfo) SetPropertyClusterList(value []uint32) (err error) {
	return instance.SetProperty("ClusterList", (value))
}

// GetClusterList gets the value of ClusterList for the instance
func (instance *BgpRouteInfo) GetPropertyClusterList() (value []uint32, err error) {
	retValue, err := instance.GetProperty("ClusterList")
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

// SetCommunity sets the value of Community for the instance
func (instance *BgpRouteInfo) SetPropertyCommunity(value []string) (err error) {
	return instance.SetProperty("Community", (value))
}

// GetCommunity gets the value of Community for the instance
func (instance *BgpRouteInfo) GetPropertyCommunity() (value []string, err error) {
	retValue, err := instance.GetProperty("Community")
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

// SetLearnedFromPeer sets the value of LearnedFromPeer for the instance
func (instance *BgpRouteInfo) SetPropertyLearnedFromPeer(value string) (err error) {
	return instance.SetProperty("LearnedFromPeer", (value))
}

// GetLearnedFromPeer gets the value of LearnedFromPeer for the instance
func (instance *BgpRouteInfo) GetPropertyLearnedFromPeer() (value string, err error) {
	retValue, err := instance.GetProperty("LearnedFromPeer")
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

// SetLocalPref sets the value of LocalPref for the instance
func (instance *BgpRouteInfo) SetPropertyLocalPref(value uint32) (err error) {
	return instance.SetProperty("LocalPref", (value))
}

// GetLocalPref gets the value of LocalPref for the instance
func (instance *BgpRouteInfo) GetPropertyLocalPref() (value uint32, err error) {
	retValue, err := instance.GetProperty("LocalPref")
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

// SetMED sets the value of MED for the instance
func (instance *BgpRouteInfo) SetPropertyMED(value uint32) (err error) {
	return instance.SetProperty("MED", (value))
}

// GetMED gets the value of MED for the instance
func (instance *BgpRouteInfo) GetPropertyMED() (value uint32, err error) {
	retValue, err := instance.GetProperty("MED")
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

// SetNetwork sets the value of Network for the instance
func (instance *BgpRouteInfo) SetPropertyNetwork(value string) (err error) {
	return instance.SetProperty("Network", (value))
}

// GetNetwork gets the value of Network for the instance
func (instance *BgpRouteInfo) GetPropertyNetwork() (value string, err error) {
	retValue, err := instance.GetProperty("Network")
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

// SetNextHop sets the value of NextHop for the instance
func (instance *BgpRouteInfo) SetPropertyNextHop(value string) (err error) {
	return instance.SetProperty("NextHop", (value))
}

// GetNextHop gets the value of NextHop for the instance
func (instance *BgpRouteInfo) GetPropertyNextHop() (value string, err error) {
	retValue, err := instance.GetProperty("NextHop")
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

// SetOrigin sets the value of Origin for the instance
func (instance *BgpRouteInfo) SetPropertyOrigin(value string) (err error) {
	return instance.SetProperty("Origin", (value))
}

// GetOrigin gets the value of Origin for the instance
func (instance *BgpRouteInfo) GetPropertyOrigin() (value string, err error) {
	retValue, err := instance.GetProperty("Origin")
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

// SetOriginatorId sets the value of OriginatorId for the instance
func (instance *BgpRouteInfo) SetPropertyOriginatorId(value string) (err error) {
	return instance.SetProperty("OriginatorId", (value))
}

// GetOriginatorId gets the value of OriginatorId for the instance
func (instance *BgpRouteInfo) GetPropertyOriginatorId() (value string, err error) {
	retValue, err := instance.GetProperty("OriginatorId")
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

// SetRoutingDomain sets the value of RoutingDomain for the instance
func (instance *BgpRouteInfo) SetPropertyRoutingDomain(value string) (err error) {
	return instance.SetProperty("RoutingDomain", (value))
}

// GetRoutingDomain gets the value of RoutingDomain for the instance
func (instance *BgpRouteInfo) GetPropertyRoutingDomain() (value string, err error) {
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

// SetState sets the value of State for the instance
func (instance *BgpRouteInfo) SetPropertyState(value []uint32) (err error) {
	return instance.SetProperty("State", (value))
}

// GetState gets the value of State for the instance
func (instance *BgpRouteInfo) GetPropertyState() (value []uint32, err error) {
	retValue, err := instance.GetProperty("State")
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
