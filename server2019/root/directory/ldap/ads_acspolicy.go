// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.directory.LDAP
//
// ////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ads_acspolicy struct
type ads_acspolicy struct {
	*ds_top

	//
	DS_aCSAggregateTokenRatePerUser int64

	//
	DS_aCSDirection int32

	//
	DS_aCSIdentityName []string

	//
	DS_aCSMaxAggregatePeakRatePerUser int64

	//
	DS_aCSMaxDurationPerFlow int32

	//
	DS_aCSMaximumSDUSize int64

	//
	DS_aCSMaxPeakBandwidthPerFlow int64

	//
	DS_aCSMaxTokenBucketPerFlow int64

	//
	DS_aCSMaxTokenRatePerFlow int64

	//
	DS_aCSMinimumDelayVariation int64

	//
	DS_aCSMinimumLatency int64

	//
	DS_aCSMinimumPolicedSize int64

	//
	DS_aCSPermissionBits int64

	//
	DS_aCSPriority int32

	//
	DS_aCSServiceType int32

	//
	DS_aCSTimeOfDay string

	//
	DS_aCSTotalNoOfFlows int32
}

func Newads_acspolicyEx1(instance *cim.WmiInstance) (newInstance *ads_acspolicy, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_acspolicy{
		ds_top: tmp,
	}
	return
}

func Newads_acspolicyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_acspolicy, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_acspolicy{
		ds_top: tmp,
	}
	return
}

// SetDS_aCSAggregateTokenRatePerUser sets the value of DS_aCSAggregateTokenRatePerUser for the instance
func (instance *ads_acspolicy) SetPropertyDS_aCSAggregateTokenRatePerUser(value int64) (err error) {
	return instance.SetProperty("DS_aCSAggregateTokenRatePerUser", (value))
}

// GetDS_aCSAggregateTokenRatePerUser gets the value of DS_aCSAggregateTokenRatePerUser for the instance
func (instance *ads_acspolicy) GetPropertyDS_aCSAggregateTokenRatePerUser() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_aCSAggregateTokenRatePerUser")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetDS_aCSDirection sets the value of DS_aCSDirection for the instance
func (instance *ads_acspolicy) SetPropertyDS_aCSDirection(value int32) (err error) {
	return instance.SetProperty("DS_aCSDirection", (value))
}

// GetDS_aCSDirection gets the value of DS_aCSDirection for the instance
func (instance *ads_acspolicy) GetPropertyDS_aCSDirection() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_aCSDirection")
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

// SetDS_aCSIdentityName sets the value of DS_aCSIdentityName for the instance
func (instance *ads_acspolicy) SetPropertyDS_aCSIdentityName(value []string) (err error) {
	return instance.SetProperty("DS_aCSIdentityName", (value))
}

// GetDS_aCSIdentityName gets the value of DS_aCSIdentityName for the instance
func (instance *ads_acspolicy) GetPropertyDS_aCSIdentityName() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_aCSIdentityName")
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

// SetDS_aCSMaxAggregatePeakRatePerUser sets the value of DS_aCSMaxAggregatePeakRatePerUser for the instance
func (instance *ads_acspolicy) SetPropertyDS_aCSMaxAggregatePeakRatePerUser(value int64) (err error) {
	return instance.SetProperty("DS_aCSMaxAggregatePeakRatePerUser", (value))
}

// GetDS_aCSMaxAggregatePeakRatePerUser gets the value of DS_aCSMaxAggregatePeakRatePerUser for the instance
func (instance *ads_acspolicy) GetPropertyDS_aCSMaxAggregatePeakRatePerUser() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_aCSMaxAggregatePeakRatePerUser")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetDS_aCSMaxDurationPerFlow sets the value of DS_aCSMaxDurationPerFlow for the instance
func (instance *ads_acspolicy) SetPropertyDS_aCSMaxDurationPerFlow(value int32) (err error) {
	return instance.SetProperty("DS_aCSMaxDurationPerFlow", (value))
}

// GetDS_aCSMaxDurationPerFlow gets the value of DS_aCSMaxDurationPerFlow for the instance
func (instance *ads_acspolicy) GetPropertyDS_aCSMaxDurationPerFlow() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_aCSMaxDurationPerFlow")
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

// SetDS_aCSMaximumSDUSize sets the value of DS_aCSMaximumSDUSize for the instance
func (instance *ads_acspolicy) SetPropertyDS_aCSMaximumSDUSize(value int64) (err error) {
	return instance.SetProperty("DS_aCSMaximumSDUSize", (value))
}

// GetDS_aCSMaximumSDUSize gets the value of DS_aCSMaximumSDUSize for the instance
func (instance *ads_acspolicy) GetPropertyDS_aCSMaximumSDUSize() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_aCSMaximumSDUSize")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetDS_aCSMaxPeakBandwidthPerFlow sets the value of DS_aCSMaxPeakBandwidthPerFlow for the instance
func (instance *ads_acspolicy) SetPropertyDS_aCSMaxPeakBandwidthPerFlow(value int64) (err error) {
	return instance.SetProperty("DS_aCSMaxPeakBandwidthPerFlow", (value))
}

// GetDS_aCSMaxPeakBandwidthPerFlow gets the value of DS_aCSMaxPeakBandwidthPerFlow for the instance
func (instance *ads_acspolicy) GetPropertyDS_aCSMaxPeakBandwidthPerFlow() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_aCSMaxPeakBandwidthPerFlow")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetDS_aCSMaxTokenBucketPerFlow sets the value of DS_aCSMaxTokenBucketPerFlow for the instance
func (instance *ads_acspolicy) SetPropertyDS_aCSMaxTokenBucketPerFlow(value int64) (err error) {
	return instance.SetProperty("DS_aCSMaxTokenBucketPerFlow", (value))
}

// GetDS_aCSMaxTokenBucketPerFlow gets the value of DS_aCSMaxTokenBucketPerFlow for the instance
func (instance *ads_acspolicy) GetPropertyDS_aCSMaxTokenBucketPerFlow() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_aCSMaxTokenBucketPerFlow")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetDS_aCSMaxTokenRatePerFlow sets the value of DS_aCSMaxTokenRatePerFlow for the instance
func (instance *ads_acspolicy) SetPropertyDS_aCSMaxTokenRatePerFlow(value int64) (err error) {
	return instance.SetProperty("DS_aCSMaxTokenRatePerFlow", (value))
}

// GetDS_aCSMaxTokenRatePerFlow gets the value of DS_aCSMaxTokenRatePerFlow for the instance
func (instance *ads_acspolicy) GetPropertyDS_aCSMaxTokenRatePerFlow() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_aCSMaxTokenRatePerFlow")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetDS_aCSMinimumDelayVariation sets the value of DS_aCSMinimumDelayVariation for the instance
func (instance *ads_acspolicy) SetPropertyDS_aCSMinimumDelayVariation(value int64) (err error) {
	return instance.SetProperty("DS_aCSMinimumDelayVariation", (value))
}

// GetDS_aCSMinimumDelayVariation gets the value of DS_aCSMinimumDelayVariation for the instance
func (instance *ads_acspolicy) GetPropertyDS_aCSMinimumDelayVariation() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_aCSMinimumDelayVariation")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetDS_aCSMinimumLatency sets the value of DS_aCSMinimumLatency for the instance
func (instance *ads_acspolicy) SetPropertyDS_aCSMinimumLatency(value int64) (err error) {
	return instance.SetProperty("DS_aCSMinimumLatency", (value))
}

// GetDS_aCSMinimumLatency gets the value of DS_aCSMinimumLatency for the instance
func (instance *ads_acspolicy) GetPropertyDS_aCSMinimumLatency() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_aCSMinimumLatency")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetDS_aCSMinimumPolicedSize sets the value of DS_aCSMinimumPolicedSize for the instance
func (instance *ads_acspolicy) SetPropertyDS_aCSMinimumPolicedSize(value int64) (err error) {
	return instance.SetProperty("DS_aCSMinimumPolicedSize", (value))
}

// GetDS_aCSMinimumPolicedSize gets the value of DS_aCSMinimumPolicedSize for the instance
func (instance *ads_acspolicy) GetPropertyDS_aCSMinimumPolicedSize() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_aCSMinimumPolicedSize")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetDS_aCSPermissionBits sets the value of DS_aCSPermissionBits for the instance
func (instance *ads_acspolicy) SetPropertyDS_aCSPermissionBits(value int64) (err error) {
	return instance.SetProperty("DS_aCSPermissionBits", (value))
}

// GetDS_aCSPermissionBits gets the value of DS_aCSPermissionBits for the instance
func (instance *ads_acspolicy) GetPropertyDS_aCSPermissionBits() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_aCSPermissionBits")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetDS_aCSPriority sets the value of DS_aCSPriority for the instance
func (instance *ads_acspolicy) SetPropertyDS_aCSPriority(value int32) (err error) {
	return instance.SetProperty("DS_aCSPriority", (value))
}

// GetDS_aCSPriority gets the value of DS_aCSPriority for the instance
func (instance *ads_acspolicy) GetPropertyDS_aCSPriority() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_aCSPriority")
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

// SetDS_aCSServiceType sets the value of DS_aCSServiceType for the instance
func (instance *ads_acspolicy) SetPropertyDS_aCSServiceType(value int32) (err error) {
	return instance.SetProperty("DS_aCSServiceType", (value))
}

// GetDS_aCSServiceType gets the value of DS_aCSServiceType for the instance
func (instance *ads_acspolicy) GetPropertyDS_aCSServiceType() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_aCSServiceType")
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

// SetDS_aCSTimeOfDay sets the value of DS_aCSTimeOfDay for the instance
func (instance *ads_acspolicy) SetPropertyDS_aCSTimeOfDay(value string) (err error) {
	return instance.SetProperty("DS_aCSTimeOfDay", (value))
}

// GetDS_aCSTimeOfDay gets the value of DS_aCSTimeOfDay for the instance
func (instance *ads_acspolicy) GetPropertyDS_aCSTimeOfDay() (value string, err error) {
	retValue, err := instance.GetProperty("DS_aCSTimeOfDay")
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

// SetDS_aCSTotalNoOfFlows sets the value of DS_aCSTotalNoOfFlows for the instance
func (instance *ads_acspolicy) SetPropertyDS_aCSTotalNoOfFlows(value int32) (err error) {
	return instance.SetProperty("DS_aCSTotalNoOfFlows", (value))
}

// GetDS_aCSTotalNoOfFlows gets the value of DS_aCSTotalNoOfFlows for the instance
func (instance *ads_acspolicy) GetPropertyDS_aCSTotalNoOfFlows() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_aCSTotalNoOfFlows")
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
