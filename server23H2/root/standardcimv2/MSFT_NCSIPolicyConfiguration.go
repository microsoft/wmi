// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NCSIPolicyConfiguration struct
type MSFT_NCSIPolicyConfiguration struct {
	*MSFT_NetSettingData

	//
	CorporateDNSProbeHostAddress string

	//
	CorporateDNSProbeHostName string

	//
	CorporateSitePrefixList []string

	//
	CorporateWebsiteProbeURL string

	//
	DomainLocationDeterminationURL string

	//
	PolicyStore string
}

func NewMSFT_NCSIPolicyConfigurationEx1(instance *cim.WmiInstance) (newInstance *MSFT_NCSIPolicyConfiguration, err error) {
	tmp, err := NewMSFT_NetSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NCSIPolicyConfiguration{
		MSFT_NetSettingData: tmp,
	}
	return
}

func NewMSFT_NCSIPolicyConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NCSIPolicyConfiguration, err error) {
	tmp, err := NewMSFT_NetSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NCSIPolicyConfiguration{
		MSFT_NetSettingData: tmp,
	}
	return
}

// SetCorporateDNSProbeHostAddress sets the value of CorporateDNSProbeHostAddress for the instance
func (instance *MSFT_NCSIPolicyConfiguration) SetPropertyCorporateDNSProbeHostAddress(value string) (err error) {
	return instance.SetProperty("CorporateDNSProbeHostAddress", (value))
}

// GetCorporateDNSProbeHostAddress gets the value of CorporateDNSProbeHostAddress for the instance
func (instance *MSFT_NCSIPolicyConfiguration) GetPropertyCorporateDNSProbeHostAddress() (value string, err error) {
	retValue, err := instance.GetProperty("CorporateDNSProbeHostAddress")
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

// SetCorporateDNSProbeHostName sets the value of CorporateDNSProbeHostName for the instance
func (instance *MSFT_NCSIPolicyConfiguration) SetPropertyCorporateDNSProbeHostName(value string) (err error) {
	return instance.SetProperty("CorporateDNSProbeHostName", (value))
}

// GetCorporateDNSProbeHostName gets the value of CorporateDNSProbeHostName for the instance
func (instance *MSFT_NCSIPolicyConfiguration) GetPropertyCorporateDNSProbeHostName() (value string, err error) {
	retValue, err := instance.GetProperty("CorporateDNSProbeHostName")
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

// SetCorporateSitePrefixList sets the value of CorporateSitePrefixList for the instance
func (instance *MSFT_NCSIPolicyConfiguration) SetPropertyCorporateSitePrefixList(value []string) (err error) {
	return instance.SetProperty("CorporateSitePrefixList", (value))
}

// GetCorporateSitePrefixList gets the value of CorporateSitePrefixList for the instance
func (instance *MSFT_NCSIPolicyConfiguration) GetPropertyCorporateSitePrefixList() (value []string, err error) {
	retValue, err := instance.GetProperty("CorporateSitePrefixList")
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

// SetCorporateWebsiteProbeURL sets the value of CorporateWebsiteProbeURL for the instance
func (instance *MSFT_NCSIPolicyConfiguration) SetPropertyCorporateWebsiteProbeURL(value string) (err error) {
	return instance.SetProperty("CorporateWebsiteProbeURL", (value))
}

// GetCorporateWebsiteProbeURL gets the value of CorporateWebsiteProbeURL for the instance
func (instance *MSFT_NCSIPolicyConfiguration) GetPropertyCorporateWebsiteProbeURL() (value string, err error) {
	retValue, err := instance.GetProperty("CorporateWebsiteProbeURL")
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

// SetDomainLocationDeterminationURL sets the value of DomainLocationDeterminationURL for the instance
func (instance *MSFT_NCSIPolicyConfiguration) SetPropertyDomainLocationDeterminationURL(value string) (err error) {
	return instance.SetProperty("DomainLocationDeterminationURL", (value))
}

// GetDomainLocationDeterminationURL gets the value of DomainLocationDeterminationURL for the instance
func (instance *MSFT_NCSIPolicyConfiguration) GetPropertyDomainLocationDeterminationURL() (value string, err error) {
	retValue, err := instance.GetProperty("DomainLocationDeterminationURL")
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

// SetPolicyStore sets the value of PolicyStore for the instance
func (instance *MSFT_NCSIPolicyConfiguration) SetPropertyPolicyStore(value string) (err error) {
	return instance.SetProperty("PolicyStore", (value))
}

// GetPolicyStore gets the value of PolicyStore for the instance
func (instance *MSFT_NCSIPolicyConfiguration) GetPropertyPolicyStore() (value string, err error) {
	retValue, err := instance.GetProperty("PolicyStore")
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

//

// <param name="CorporateDNSProbeHostAddress" type="bool "></param>
// <param name="CorporateDNSProbeHostName" type="bool "></param>
// <param name="CorporateSitePrefixList" type="bool "></param>
// <param name="CorporateWebsiteProbeURL" type="bool "></param>
// <param name="DomainLocationDeterminationURL" type="bool "></param>
// <param name="PassThru" type="bool "></param>

// <param name="OutputObject" type="MSFT_NCSIPolicyConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NCSIPolicyConfiguration) Reset( /* IN */ CorporateDNSProbeHostAddress bool,
	/* IN */ CorporateDNSProbeHostName bool,
	/* IN */ CorporateSitePrefixList bool,
	/* IN */ CorporateWebsiteProbeURL bool,
	/* IN */ DomainLocationDeterminationURL bool,
	/* IN */ PassThru bool,
	/* OUT */ OutputObject MSFT_NCSIPolicyConfiguration) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Reset", CorporateDNSProbeHostAddress, CorporateDNSProbeHostName, CorporateSitePrefixList, CorporateWebsiteProbeURL, DomainLocationDeterminationURL, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
