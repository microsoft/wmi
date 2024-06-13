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

// ads_organizationalperson struct
type ads_organizationalperson struct {
	*ads_person

	//
	DS_assistant string

	//
	DS_c string

	//
	DS_co string

	//
	DS_comment string

	//
	DS_company string

	//
	DS_countryCode int32

	//
	DS_department string

	//
	DS_destinationIndicator []string

	//
	DS_division string

	//
	DS_employeeID string

	//
	DS_facsimileTelephoneNumber string

	//
	DS_generationQualifier string

	//
	DS_givenName string

	//
	DS_homePhone string

	//
	DS_homePostalAddress string

	//
	DS_houseIdentifier []string

	//
	DS_initials string

	//
	DS_internationalISDNNumber []string

	//
	DS_ipPhone string

	//
	DS_l string

	//
	DS_mail string

	//
	DS_manager string

	//
	DS_mhsORAddress []string

	//
	DS_middleName string

	//
	DS_mobile string

	//
	DS_msDS_AllowedToActOnBehalfOfOtherIdentity Uint8Array

	//
	DS_msDS_AllowedToDelegateTo []string

	//
	DS_msDS_HABSeniorityIndex int32

	//
	DS_msDS_PhoneticCompanyName string

	//
	DS_msDS_PhoneticDepartment string

	//
	DS_msDS_PhoneticDisplayName string

	//
	DS_msDS_PhoneticFirstName string

	//
	DS_msDS_PhoneticLastName string

	//
	DS_msExchHouseIdentifier string

	//
	DS_o []string

	//
	DS_otherFacsimileTelephoneNumber []string

	//
	DS_otherHomePhone []string

	//
	DS_otherIpPhone []string

	//
	DS_otherMailbox []string

	//
	DS_otherMobile []string

	//
	DS_otherPager []string

	//
	DS_otherTelephone []string

	//
	DS_ou []string

	//
	DS_pager string

	//
	DS_personalTitle string

	//
	DS_physicalDeliveryOfficeName string

	//
	DS_postalAddress []string

	//
	DS_postalCode string

	//
	DS_postOfficeBox []string

	//
	DS_preferredDeliveryMethod []int32

	//
	DS_primaryInternationalISDNNumber string

	//
	DS_primaryTelexNumber string

	//
	DS_registeredAddress []Uint8Array

	//
	DS_st string

	//
	DS_street string

	//
	DS_streetAddress string

	//
	DS_teletexTerminalIdentifier []Uint8Array

	//
	DS_telexNumber []Uint8Array

	//
	DS_thumbnailLogo Uint8Array

	//
	DS_thumbnailPhoto Uint8Array

	//
	DS_title string

	//
	DS_x121Address []string
}

func Newads_organizationalpersonEx1(instance *cim.WmiInstance) (newInstance *ads_organizationalperson, err error) {
	tmp, err := Newads_personEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_organizationalperson{
		ads_person: tmp,
	}
	return
}

func Newads_organizationalpersonEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_organizationalperson, err error) {
	tmp, err := Newads_personEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_organizationalperson{
		ads_person: tmp,
	}
	return
}

// SetDS_assistant sets the value of DS_assistant for the instance
func (instance *ads_organizationalperson) SetPropertyDS_assistant(value string) (err error) {
	return instance.SetProperty("DS_assistant", (value))
}

// GetDS_assistant gets the value of DS_assistant for the instance
func (instance *ads_organizationalperson) GetPropertyDS_assistant() (value string, err error) {
	retValue, err := instance.GetProperty("DS_assistant")
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

// SetDS_c sets the value of DS_c for the instance
func (instance *ads_organizationalperson) SetPropertyDS_c(value string) (err error) {
	return instance.SetProperty("DS_c", (value))
}

// GetDS_c gets the value of DS_c for the instance
func (instance *ads_organizationalperson) GetPropertyDS_c() (value string, err error) {
	retValue, err := instance.GetProperty("DS_c")
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

// SetDS_co sets the value of DS_co for the instance
func (instance *ads_organizationalperson) SetPropertyDS_co(value string) (err error) {
	return instance.SetProperty("DS_co", (value))
}

// GetDS_co gets the value of DS_co for the instance
func (instance *ads_organizationalperson) GetPropertyDS_co() (value string, err error) {
	retValue, err := instance.GetProperty("DS_co")
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

// SetDS_comment sets the value of DS_comment for the instance
func (instance *ads_organizationalperson) SetPropertyDS_comment(value string) (err error) {
	return instance.SetProperty("DS_comment", (value))
}

// GetDS_comment gets the value of DS_comment for the instance
func (instance *ads_organizationalperson) GetPropertyDS_comment() (value string, err error) {
	retValue, err := instance.GetProperty("DS_comment")
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

// SetDS_company sets the value of DS_company for the instance
func (instance *ads_organizationalperson) SetPropertyDS_company(value string) (err error) {
	return instance.SetProperty("DS_company", (value))
}

// GetDS_company gets the value of DS_company for the instance
func (instance *ads_organizationalperson) GetPropertyDS_company() (value string, err error) {
	retValue, err := instance.GetProperty("DS_company")
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

// SetDS_countryCode sets the value of DS_countryCode for the instance
func (instance *ads_organizationalperson) SetPropertyDS_countryCode(value int32) (err error) {
	return instance.SetProperty("DS_countryCode", (value))
}

// GetDS_countryCode gets the value of DS_countryCode for the instance
func (instance *ads_organizationalperson) GetPropertyDS_countryCode() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_countryCode")
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

// SetDS_department sets the value of DS_department for the instance
func (instance *ads_organizationalperson) SetPropertyDS_department(value string) (err error) {
	return instance.SetProperty("DS_department", (value))
}

// GetDS_department gets the value of DS_department for the instance
func (instance *ads_organizationalperson) GetPropertyDS_department() (value string, err error) {
	retValue, err := instance.GetProperty("DS_department")
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

// SetDS_destinationIndicator sets the value of DS_destinationIndicator for the instance
func (instance *ads_organizationalperson) SetPropertyDS_destinationIndicator(value []string) (err error) {
	return instance.SetProperty("DS_destinationIndicator", (value))
}

// GetDS_destinationIndicator gets the value of DS_destinationIndicator for the instance
func (instance *ads_organizationalperson) GetPropertyDS_destinationIndicator() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_destinationIndicator")
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

// SetDS_division sets the value of DS_division for the instance
func (instance *ads_organizationalperson) SetPropertyDS_division(value string) (err error) {
	return instance.SetProperty("DS_division", (value))
}

// GetDS_division gets the value of DS_division for the instance
func (instance *ads_organizationalperson) GetPropertyDS_division() (value string, err error) {
	retValue, err := instance.GetProperty("DS_division")
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

// SetDS_employeeID sets the value of DS_employeeID for the instance
func (instance *ads_organizationalperson) SetPropertyDS_employeeID(value string) (err error) {
	return instance.SetProperty("DS_employeeID", (value))
}

// GetDS_employeeID gets the value of DS_employeeID for the instance
func (instance *ads_organizationalperson) GetPropertyDS_employeeID() (value string, err error) {
	retValue, err := instance.GetProperty("DS_employeeID")
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

// SetDS_facsimileTelephoneNumber sets the value of DS_facsimileTelephoneNumber for the instance
func (instance *ads_organizationalperson) SetPropertyDS_facsimileTelephoneNumber(value string) (err error) {
	return instance.SetProperty("DS_facsimileTelephoneNumber", (value))
}

// GetDS_facsimileTelephoneNumber gets the value of DS_facsimileTelephoneNumber for the instance
func (instance *ads_organizationalperson) GetPropertyDS_facsimileTelephoneNumber() (value string, err error) {
	retValue, err := instance.GetProperty("DS_facsimileTelephoneNumber")
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

// SetDS_generationQualifier sets the value of DS_generationQualifier for the instance
func (instance *ads_organizationalperson) SetPropertyDS_generationQualifier(value string) (err error) {
	return instance.SetProperty("DS_generationQualifier", (value))
}

// GetDS_generationQualifier gets the value of DS_generationQualifier for the instance
func (instance *ads_organizationalperson) GetPropertyDS_generationQualifier() (value string, err error) {
	retValue, err := instance.GetProperty("DS_generationQualifier")
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

// SetDS_givenName sets the value of DS_givenName for the instance
func (instance *ads_organizationalperson) SetPropertyDS_givenName(value string) (err error) {
	return instance.SetProperty("DS_givenName", (value))
}

// GetDS_givenName gets the value of DS_givenName for the instance
func (instance *ads_organizationalperson) GetPropertyDS_givenName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_givenName")
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

// SetDS_homePhone sets the value of DS_homePhone for the instance
func (instance *ads_organizationalperson) SetPropertyDS_homePhone(value string) (err error) {
	return instance.SetProperty("DS_homePhone", (value))
}

// GetDS_homePhone gets the value of DS_homePhone for the instance
func (instance *ads_organizationalperson) GetPropertyDS_homePhone() (value string, err error) {
	retValue, err := instance.GetProperty("DS_homePhone")
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

// SetDS_homePostalAddress sets the value of DS_homePostalAddress for the instance
func (instance *ads_organizationalperson) SetPropertyDS_homePostalAddress(value string) (err error) {
	return instance.SetProperty("DS_homePostalAddress", (value))
}

// GetDS_homePostalAddress gets the value of DS_homePostalAddress for the instance
func (instance *ads_organizationalperson) GetPropertyDS_homePostalAddress() (value string, err error) {
	retValue, err := instance.GetProperty("DS_homePostalAddress")
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

// SetDS_houseIdentifier sets the value of DS_houseIdentifier for the instance
func (instance *ads_organizationalperson) SetPropertyDS_houseIdentifier(value []string) (err error) {
	return instance.SetProperty("DS_houseIdentifier", (value))
}

// GetDS_houseIdentifier gets the value of DS_houseIdentifier for the instance
func (instance *ads_organizationalperson) GetPropertyDS_houseIdentifier() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_houseIdentifier")
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

// SetDS_initials sets the value of DS_initials for the instance
func (instance *ads_organizationalperson) SetPropertyDS_initials(value string) (err error) {
	return instance.SetProperty("DS_initials", (value))
}

// GetDS_initials gets the value of DS_initials for the instance
func (instance *ads_organizationalperson) GetPropertyDS_initials() (value string, err error) {
	retValue, err := instance.GetProperty("DS_initials")
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

// SetDS_internationalISDNNumber sets the value of DS_internationalISDNNumber for the instance
func (instance *ads_organizationalperson) SetPropertyDS_internationalISDNNumber(value []string) (err error) {
	return instance.SetProperty("DS_internationalISDNNumber", (value))
}

// GetDS_internationalISDNNumber gets the value of DS_internationalISDNNumber for the instance
func (instance *ads_organizationalperson) GetPropertyDS_internationalISDNNumber() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_internationalISDNNumber")
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

// SetDS_ipPhone sets the value of DS_ipPhone for the instance
func (instance *ads_organizationalperson) SetPropertyDS_ipPhone(value string) (err error) {
	return instance.SetProperty("DS_ipPhone", (value))
}

// GetDS_ipPhone gets the value of DS_ipPhone for the instance
func (instance *ads_organizationalperson) GetPropertyDS_ipPhone() (value string, err error) {
	retValue, err := instance.GetProperty("DS_ipPhone")
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

// SetDS_l sets the value of DS_l for the instance
func (instance *ads_organizationalperson) SetPropertyDS_l(value string) (err error) {
	return instance.SetProperty("DS_l", (value))
}

// GetDS_l gets the value of DS_l for the instance
func (instance *ads_organizationalperson) GetPropertyDS_l() (value string, err error) {
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

// SetDS_mail sets the value of DS_mail for the instance
func (instance *ads_organizationalperson) SetPropertyDS_mail(value string) (err error) {
	return instance.SetProperty("DS_mail", (value))
}

// GetDS_mail gets the value of DS_mail for the instance
func (instance *ads_organizationalperson) GetPropertyDS_mail() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mail")
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

// SetDS_manager sets the value of DS_manager for the instance
func (instance *ads_organizationalperson) SetPropertyDS_manager(value string) (err error) {
	return instance.SetProperty("DS_manager", (value))
}

// GetDS_manager gets the value of DS_manager for the instance
func (instance *ads_organizationalperson) GetPropertyDS_manager() (value string, err error) {
	retValue, err := instance.GetProperty("DS_manager")
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

// SetDS_mhsORAddress sets the value of DS_mhsORAddress for the instance
func (instance *ads_organizationalperson) SetPropertyDS_mhsORAddress(value []string) (err error) {
	return instance.SetProperty("DS_mhsORAddress", (value))
}

// GetDS_mhsORAddress gets the value of DS_mhsORAddress for the instance
func (instance *ads_organizationalperson) GetPropertyDS_mhsORAddress() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_mhsORAddress")
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

// SetDS_middleName sets the value of DS_middleName for the instance
func (instance *ads_organizationalperson) SetPropertyDS_middleName(value string) (err error) {
	return instance.SetProperty("DS_middleName", (value))
}

// GetDS_middleName gets the value of DS_middleName for the instance
func (instance *ads_organizationalperson) GetPropertyDS_middleName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_middleName")
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

// SetDS_mobile sets the value of DS_mobile for the instance
func (instance *ads_organizationalperson) SetPropertyDS_mobile(value string) (err error) {
	return instance.SetProperty("DS_mobile", (value))
}

// GetDS_mobile gets the value of DS_mobile for the instance
func (instance *ads_organizationalperson) GetPropertyDS_mobile() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mobile")
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

// SetDS_msDS_AllowedToActOnBehalfOfOtherIdentity sets the value of DS_msDS_AllowedToActOnBehalfOfOtherIdentity for the instance
func (instance *ads_organizationalperson) SetPropertyDS_msDS_AllowedToActOnBehalfOfOtherIdentity(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDS_AllowedToActOnBehalfOfOtherIdentity", (value))
}

// GetDS_msDS_AllowedToActOnBehalfOfOtherIdentity gets the value of DS_msDS_AllowedToActOnBehalfOfOtherIdentity for the instance
func (instance *ads_organizationalperson) GetPropertyDS_msDS_AllowedToActOnBehalfOfOtherIdentity() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AllowedToActOnBehalfOfOtherIdentity")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(Uint8Array)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = Uint8Array(valuetmp)

	return
}

// SetDS_msDS_AllowedToDelegateTo sets the value of DS_msDS_AllowedToDelegateTo for the instance
func (instance *ads_organizationalperson) SetPropertyDS_msDS_AllowedToDelegateTo(value []string) (err error) {
	return instance.SetProperty("DS_msDS_AllowedToDelegateTo", (value))
}

// GetDS_msDS_AllowedToDelegateTo gets the value of DS_msDS_AllowedToDelegateTo for the instance
func (instance *ads_organizationalperson) GetPropertyDS_msDS_AllowedToDelegateTo() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AllowedToDelegateTo")
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

// SetDS_msDS_HABSeniorityIndex sets the value of DS_msDS_HABSeniorityIndex for the instance
func (instance *ads_organizationalperson) SetPropertyDS_msDS_HABSeniorityIndex(value int32) (err error) {
	return instance.SetProperty("DS_msDS_HABSeniorityIndex", (value))
}

// GetDS_msDS_HABSeniorityIndex gets the value of DS_msDS_HABSeniorityIndex for the instance
func (instance *ads_organizationalperson) GetPropertyDS_msDS_HABSeniorityIndex() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_HABSeniorityIndex")
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

// SetDS_msDS_PhoneticCompanyName sets the value of DS_msDS_PhoneticCompanyName for the instance
func (instance *ads_organizationalperson) SetPropertyDS_msDS_PhoneticCompanyName(value string) (err error) {
	return instance.SetProperty("DS_msDS_PhoneticCompanyName", (value))
}

// GetDS_msDS_PhoneticCompanyName gets the value of DS_msDS_PhoneticCompanyName for the instance
func (instance *ads_organizationalperson) GetPropertyDS_msDS_PhoneticCompanyName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_PhoneticCompanyName")
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

// SetDS_msDS_PhoneticDepartment sets the value of DS_msDS_PhoneticDepartment for the instance
func (instance *ads_organizationalperson) SetPropertyDS_msDS_PhoneticDepartment(value string) (err error) {
	return instance.SetProperty("DS_msDS_PhoneticDepartment", (value))
}

// GetDS_msDS_PhoneticDepartment gets the value of DS_msDS_PhoneticDepartment for the instance
func (instance *ads_organizationalperson) GetPropertyDS_msDS_PhoneticDepartment() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_PhoneticDepartment")
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

// SetDS_msDS_PhoneticDisplayName sets the value of DS_msDS_PhoneticDisplayName for the instance
func (instance *ads_organizationalperson) SetPropertyDS_msDS_PhoneticDisplayName(value string) (err error) {
	return instance.SetProperty("DS_msDS_PhoneticDisplayName", (value))
}

// GetDS_msDS_PhoneticDisplayName gets the value of DS_msDS_PhoneticDisplayName for the instance
func (instance *ads_organizationalperson) GetPropertyDS_msDS_PhoneticDisplayName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_PhoneticDisplayName")
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

// SetDS_msDS_PhoneticFirstName sets the value of DS_msDS_PhoneticFirstName for the instance
func (instance *ads_organizationalperson) SetPropertyDS_msDS_PhoneticFirstName(value string) (err error) {
	return instance.SetProperty("DS_msDS_PhoneticFirstName", (value))
}

// GetDS_msDS_PhoneticFirstName gets the value of DS_msDS_PhoneticFirstName for the instance
func (instance *ads_organizationalperson) GetPropertyDS_msDS_PhoneticFirstName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_PhoneticFirstName")
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

// SetDS_msDS_PhoneticLastName sets the value of DS_msDS_PhoneticLastName for the instance
func (instance *ads_organizationalperson) SetPropertyDS_msDS_PhoneticLastName(value string) (err error) {
	return instance.SetProperty("DS_msDS_PhoneticLastName", (value))
}

// GetDS_msDS_PhoneticLastName gets the value of DS_msDS_PhoneticLastName for the instance
func (instance *ads_organizationalperson) GetPropertyDS_msDS_PhoneticLastName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_PhoneticLastName")
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

// SetDS_msExchHouseIdentifier sets the value of DS_msExchHouseIdentifier for the instance
func (instance *ads_organizationalperson) SetPropertyDS_msExchHouseIdentifier(value string) (err error) {
	return instance.SetProperty("DS_msExchHouseIdentifier", (value))
}

// GetDS_msExchHouseIdentifier gets the value of DS_msExchHouseIdentifier for the instance
func (instance *ads_organizationalperson) GetPropertyDS_msExchHouseIdentifier() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msExchHouseIdentifier")
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
func (instance *ads_organizationalperson) SetPropertyDS_o(value []string) (err error) {
	return instance.SetProperty("DS_o", (value))
}

// GetDS_o gets the value of DS_o for the instance
func (instance *ads_organizationalperson) GetPropertyDS_o() (value []string, err error) {
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

// SetDS_otherFacsimileTelephoneNumber sets the value of DS_otherFacsimileTelephoneNumber for the instance
func (instance *ads_organizationalperson) SetPropertyDS_otherFacsimileTelephoneNumber(value []string) (err error) {
	return instance.SetProperty("DS_otherFacsimileTelephoneNumber", (value))
}

// GetDS_otherFacsimileTelephoneNumber gets the value of DS_otherFacsimileTelephoneNumber for the instance
func (instance *ads_organizationalperson) GetPropertyDS_otherFacsimileTelephoneNumber() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_otherFacsimileTelephoneNumber")
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

// SetDS_otherHomePhone sets the value of DS_otherHomePhone for the instance
func (instance *ads_organizationalperson) SetPropertyDS_otherHomePhone(value []string) (err error) {
	return instance.SetProperty("DS_otherHomePhone", (value))
}

// GetDS_otherHomePhone gets the value of DS_otherHomePhone for the instance
func (instance *ads_organizationalperson) GetPropertyDS_otherHomePhone() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_otherHomePhone")
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

// SetDS_otherIpPhone sets the value of DS_otherIpPhone for the instance
func (instance *ads_organizationalperson) SetPropertyDS_otherIpPhone(value []string) (err error) {
	return instance.SetProperty("DS_otherIpPhone", (value))
}

// GetDS_otherIpPhone gets the value of DS_otherIpPhone for the instance
func (instance *ads_organizationalperson) GetPropertyDS_otherIpPhone() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_otherIpPhone")
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

// SetDS_otherMailbox sets the value of DS_otherMailbox for the instance
func (instance *ads_organizationalperson) SetPropertyDS_otherMailbox(value []string) (err error) {
	return instance.SetProperty("DS_otherMailbox", (value))
}

// GetDS_otherMailbox gets the value of DS_otherMailbox for the instance
func (instance *ads_organizationalperson) GetPropertyDS_otherMailbox() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_otherMailbox")
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

// SetDS_otherMobile sets the value of DS_otherMobile for the instance
func (instance *ads_organizationalperson) SetPropertyDS_otherMobile(value []string) (err error) {
	return instance.SetProperty("DS_otherMobile", (value))
}

// GetDS_otherMobile gets the value of DS_otherMobile for the instance
func (instance *ads_organizationalperson) GetPropertyDS_otherMobile() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_otherMobile")
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

// SetDS_otherPager sets the value of DS_otherPager for the instance
func (instance *ads_organizationalperson) SetPropertyDS_otherPager(value []string) (err error) {
	return instance.SetProperty("DS_otherPager", (value))
}

// GetDS_otherPager gets the value of DS_otherPager for the instance
func (instance *ads_organizationalperson) GetPropertyDS_otherPager() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_otherPager")
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

// SetDS_otherTelephone sets the value of DS_otherTelephone for the instance
func (instance *ads_organizationalperson) SetPropertyDS_otherTelephone(value []string) (err error) {
	return instance.SetProperty("DS_otherTelephone", (value))
}

// GetDS_otherTelephone gets the value of DS_otherTelephone for the instance
func (instance *ads_organizationalperson) GetPropertyDS_otherTelephone() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_otherTelephone")
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
func (instance *ads_organizationalperson) SetPropertyDS_ou(value []string) (err error) {
	return instance.SetProperty("DS_ou", (value))
}

// GetDS_ou gets the value of DS_ou for the instance
func (instance *ads_organizationalperson) GetPropertyDS_ou() (value []string, err error) {
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

// SetDS_pager sets the value of DS_pager for the instance
func (instance *ads_organizationalperson) SetPropertyDS_pager(value string) (err error) {
	return instance.SetProperty("DS_pager", (value))
}

// GetDS_pager gets the value of DS_pager for the instance
func (instance *ads_organizationalperson) GetPropertyDS_pager() (value string, err error) {
	retValue, err := instance.GetProperty("DS_pager")
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

// SetDS_personalTitle sets the value of DS_personalTitle for the instance
func (instance *ads_organizationalperson) SetPropertyDS_personalTitle(value string) (err error) {
	return instance.SetProperty("DS_personalTitle", (value))
}

// GetDS_personalTitle gets the value of DS_personalTitle for the instance
func (instance *ads_organizationalperson) GetPropertyDS_personalTitle() (value string, err error) {
	retValue, err := instance.GetProperty("DS_personalTitle")
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

// SetDS_physicalDeliveryOfficeName sets the value of DS_physicalDeliveryOfficeName for the instance
func (instance *ads_organizationalperson) SetPropertyDS_physicalDeliveryOfficeName(value string) (err error) {
	return instance.SetProperty("DS_physicalDeliveryOfficeName", (value))
}

// GetDS_physicalDeliveryOfficeName gets the value of DS_physicalDeliveryOfficeName for the instance
func (instance *ads_organizationalperson) GetPropertyDS_physicalDeliveryOfficeName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_physicalDeliveryOfficeName")
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

// SetDS_postalAddress sets the value of DS_postalAddress for the instance
func (instance *ads_organizationalperson) SetPropertyDS_postalAddress(value []string) (err error) {
	return instance.SetProperty("DS_postalAddress", (value))
}

// GetDS_postalAddress gets the value of DS_postalAddress for the instance
func (instance *ads_organizationalperson) GetPropertyDS_postalAddress() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_postalAddress")
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

// SetDS_postalCode sets the value of DS_postalCode for the instance
func (instance *ads_organizationalperson) SetPropertyDS_postalCode(value string) (err error) {
	return instance.SetProperty("DS_postalCode", (value))
}

// GetDS_postalCode gets the value of DS_postalCode for the instance
func (instance *ads_organizationalperson) GetPropertyDS_postalCode() (value string, err error) {
	retValue, err := instance.GetProperty("DS_postalCode")
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

// SetDS_postOfficeBox sets the value of DS_postOfficeBox for the instance
func (instance *ads_organizationalperson) SetPropertyDS_postOfficeBox(value []string) (err error) {
	return instance.SetProperty("DS_postOfficeBox", (value))
}

// GetDS_postOfficeBox gets the value of DS_postOfficeBox for the instance
func (instance *ads_organizationalperson) GetPropertyDS_postOfficeBox() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_postOfficeBox")
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

// SetDS_preferredDeliveryMethod sets the value of DS_preferredDeliveryMethod for the instance
func (instance *ads_organizationalperson) SetPropertyDS_preferredDeliveryMethod(value []int32) (err error) {
	return instance.SetProperty("DS_preferredDeliveryMethod", (value))
}

// GetDS_preferredDeliveryMethod gets the value of DS_preferredDeliveryMethod for the instance
func (instance *ads_organizationalperson) GetPropertyDS_preferredDeliveryMethod() (value []int32, err error) {
	retValue, err := instance.GetProperty("DS_preferredDeliveryMethod")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, int32(valuetmp))
	}

	return
}

// SetDS_primaryInternationalISDNNumber sets the value of DS_primaryInternationalISDNNumber for the instance
func (instance *ads_organizationalperson) SetPropertyDS_primaryInternationalISDNNumber(value string) (err error) {
	return instance.SetProperty("DS_primaryInternationalISDNNumber", (value))
}

// GetDS_primaryInternationalISDNNumber gets the value of DS_primaryInternationalISDNNumber for the instance
func (instance *ads_organizationalperson) GetPropertyDS_primaryInternationalISDNNumber() (value string, err error) {
	retValue, err := instance.GetProperty("DS_primaryInternationalISDNNumber")
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

// SetDS_primaryTelexNumber sets the value of DS_primaryTelexNumber for the instance
func (instance *ads_organizationalperson) SetPropertyDS_primaryTelexNumber(value string) (err error) {
	return instance.SetProperty("DS_primaryTelexNumber", (value))
}

// GetDS_primaryTelexNumber gets the value of DS_primaryTelexNumber for the instance
func (instance *ads_organizationalperson) GetPropertyDS_primaryTelexNumber() (value string, err error) {
	retValue, err := instance.GetProperty("DS_primaryTelexNumber")
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

// SetDS_registeredAddress sets the value of DS_registeredAddress for the instance
func (instance *ads_organizationalperson) SetPropertyDS_registeredAddress(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_registeredAddress", (value))
}

// GetDS_registeredAddress gets the value of DS_registeredAddress for the instance
func (instance *ads_organizationalperson) GetPropertyDS_registeredAddress() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_registeredAddress")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(Uint8Array)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, Uint8Array(valuetmp))
	}

	return
}

// SetDS_st sets the value of DS_st for the instance
func (instance *ads_organizationalperson) SetPropertyDS_st(value string) (err error) {
	return instance.SetProperty("DS_st", (value))
}

// GetDS_st gets the value of DS_st for the instance
func (instance *ads_organizationalperson) GetPropertyDS_st() (value string, err error) {
	retValue, err := instance.GetProperty("DS_st")
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

// SetDS_street sets the value of DS_street for the instance
func (instance *ads_organizationalperson) SetPropertyDS_street(value string) (err error) {
	return instance.SetProperty("DS_street", (value))
}

// GetDS_street gets the value of DS_street for the instance
func (instance *ads_organizationalperson) GetPropertyDS_street() (value string, err error) {
	retValue, err := instance.GetProperty("DS_street")
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

// SetDS_streetAddress sets the value of DS_streetAddress for the instance
func (instance *ads_organizationalperson) SetPropertyDS_streetAddress(value string) (err error) {
	return instance.SetProperty("DS_streetAddress", (value))
}

// GetDS_streetAddress gets the value of DS_streetAddress for the instance
func (instance *ads_organizationalperson) GetPropertyDS_streetAddress() (value string, err error) {
	retValue, err := instance.GetProperty("DS_streetAddress")
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

// SetDS_teletexTerminalIdentifier sets the value of DS_teletexTerminalIdentifier for the instance
func (instance *ads_organizationalperson) SetPropertyDS_teletexTerminalIdentifier(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_teletexTerminalIdentifier", (value))
}

// GetDS_teletexTerminalIdentifier gets the value of DS_teletexTerminalIdentifier for the instance
func (instance *ads_organizationalperson) GetPropertyDS_teletexTerminalIdentifier() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_teletexTerminalIdentifier")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(Uint8Array)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, Uint8Array(valuetmp))
	}

	return
}

// SetDS_telexNumber sets the value of DS_telexNumber for the instance
func (instance *ads_organizationalperson) SetPropertyDS_telexNumber(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_telexNumber", (value))
}

// GetDS_telexNumber gets the value of DS_telexNumber for the instance
func (instance *ads_organizationalperson) GetPropertyDS_telexNumber() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_telexNumber")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(Uint8Array)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, Uint8Array(valuetmp))
	}

	return
}

// SetDS_thumbnailLogo sets the value of DS_thumbnailLogo for the instance
func (instance *ads_organizationalperson) SetPropertyDS_thumbnailLogo(value Uint8Array) (err error) {
	return instance.SetProperty("DS_thumbnailLogo", (value))
}

// GetDS_thumbnailLogo gets the value of DS_thumbnailLogo for the instance
func (instance *ads_organizationalperson) GetPropertyDS_thumbnailLogo() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_thumbnailLogo")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(Uint8Array)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = Uint8Array(valuetmp)

	return
}

// SetDS_thumbnailPhoto sets the value of DS_thumbnailPhoto for the instance
func (instance *ads_organizationalperson) SetPropertyDS_thumbnailPhoto(value Uint8Array) (err error) {
	return instance.SetProperty("DS_thumbnailPhoto", (value))
}

// GetDS_thumbnailPhoto gets the value of DS_thumbnailPhoto for the instance
func (instance *ads_organizationalperson) GetPropertyDS_thumbnailPhoto() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_thumbnailPhoto")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(Uint8Array)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = Uint8Array(valuetmp)

	return
}

// SetDS_title sets the value of DS_title for the instance
func (instance *ads_organizationalperson) SetPropertyDS_title(value string) (err error) {
	return instance.SetProperty("DS_title", (value))
}

// GetDS_title gets the value of DS_title for the instance
func (instance *ads_organizationalperson) GetPropertyDS_title() (value string, err error) {
	retValue, err := instance.GetProperty("DS_title")
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

// SetDS_x121Address sets the value of DS_x121Address for the instance
func (instance *ads_organizationalperson) SetPropertyDS_x121Address(value []string) (err error) {
	return instance.SetProperty("DS_x121Address", (value))
}

// GetDS_x121Address gets the value of DS_x121Address for the instance
func (instance *ads_organizationalperson) GetPropertyDS_x121Address() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_x121Address")
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
