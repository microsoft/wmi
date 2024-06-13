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

// ads_acsresourcelimits struct
type ads_acsresourcelimits struct {
	*ds_top

	//
	DS_aCSAllocableRSVPBandwidth int64

	//
	DS_aCSMaxPeakBandwidth int64

	//
	DS_aCSMaxPeakBandwidthPerFlow int64

	//
	DS_aCSMaxTokenRatePerFlow int64

	//
	DS_aCSServiceType int32
}

func Newads_acsresourcelimitsEx1(instance *cim.WmiInstance) (newInstance *ads_acsresourcelimits, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_acsresourcelimits{
		ds_top: tmp,
	}
	return
}

func Newads_acsresourcelimitsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_acsresourcelimits, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_acsresourcelimits{
		ds_top: tmp,
	}
	return
}

// SetDS_aCSAllocableRSVPBandwidth sets the value of DS_aCSAllocableRSVPBandwidth for the instance
func (instance *ads_acsresourcelimits) SetPropertyDS_aCSAllocableRSVPBandwidth(value int64) (err error) {
	return instance.SetProperty("DS_aCSAllocableRSVPBandwidth", (value))
}

// GetDS_aCSAllocableRSVPBandwidth gets the value of DS_aCSAllocableRSVPBandwidth for the instance
func (instance *ads_acsresourcelimits) GetPropertyDS_aCSAllocableRSVPBandwidth() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_aCSAllocableRSVPBandwidth")
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

// SetDS_aCSMaxPeakBandwidth sets the value of DS_aCSMaxPeakBandwidth for the instance
func (instance *ads_acsresourcelimits) SetPropertyDS_aCSMaxPeakBandwidth(value int64) (err error) {
	return instance.SetProperty("DS_aCSMaxPeakBandwidth", (value))
}

// GetDS_aCSMaxPeakBandwidth gets the value of DS_aCSMaxPeakBandwidth for the instance
func (instance *ads_acsresourcelimits) GetPropertyDS_aCSMaxPeakBandwidth() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_aCSMaxPeakBandwidth")
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
func (instance *ads_acsresourcelimits) SetPropertyDS_aCSMaxPeakBandwidthPerFlow(value int64) (err error) {
	return instance.SetProperty("DS_aCSMaxPeakBandwidthPerFlow", (value))
}

// GetDS_aCSMaxPeakBandwidthPerFlow gets the value of DS_aCSMaxPeakBandwidthPerFlow for the instance
func (instance *ads_acsresourcelimits) GetPropertyDS_aCSMaxPeakBandwidthPerFlow() (value int64, err error) {
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

// SetDS_aCSMaxTokenRatePerFlow sets the value of DS_aCSMaxTokenRatePerFlow for the instance
func (instance *ads_acsresourcelimits) SetPropertyDS_aCSMaxTokenRatePerFlow(value int64) (err error) {
	return instance.SetProperty("DS_aCSMaxTokenRatePerFlow", (value))
}

// GetDS_aCSMaxTokenRatePerFlow gets the value of DS_aCSMaxTokenRatePerFlow for the instance
func (instance *ads_acsresourcelimits) GetPropertyDS_aCSMaxTokenRatePerFlow() (value int64, err error) {
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

// SetDS_aCSServiceType sets the value of DS_aCSServiceType for the instance
func (instance *ads_acsresourcelimits) SetPropertyDS_aCSServiceType(value int32) (err error) {
	return instance.SetProperty("DS_aCSServiceType", (value))
}

// GetDS_aCSServiceType gets the value of DS_aCSServiceType for the instance
func (instance *ads_acsresourcelimits) GetPropertyDS_aCSServiceType() (value int32, err error) {
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
