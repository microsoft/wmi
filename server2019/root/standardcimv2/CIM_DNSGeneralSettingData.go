// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// CIM_DNSGeneralSettingData struct
type CIM_DNSGeneralSettingData struct {
	CIM_IPAssignmentSettingData

	// 698
	AppendParentSuffixes bool

	// 697
	AppendPrimarySuffixes bool

	// 699
	DNSSuffixesToAppend []string
}

// SetAppendParentSuffixes sets the value of AppendParentSuffixes for the instance
func (instance *CIM_DNSGeneralSettingData) SetPropertyAppendParentSuffixes(value bool) (err error) {
	return instance.SetProperty("AppendParentSuffixes", value)
}

// GetAppendParentSuffixes gets the value of AppendParentSuffixes for the instance
func (instance *CIM_DNSGeneralSettingData) GetPropertyAppendParentSuffixes() (value bool, err error) {
	retValue, err := instance.GetProperty("AppendParentSuffixes")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAppendPrimarySuffixes sets the value of AppendPrimarySuffixes for the instance
func (instance *CIM_DNSGeneralSettingData) SetPropertyAppendPrimarySuffixes(value bool) (err error) {
	return instance.SetProperty("AppendPrimarySuffixes", value)
}

// GetAppendPrimarySuffixes gets the value of AppendPrimarySuffixes for the instance
func (instance *CIM_DNSGeneralSettingData) GetPropertyAppendPrimarySuffixes() (value bool, err error) {
	retValue, err := instance.GetProperty("AppendPrimarySuffixes")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDNSSuffixesToAppend sets the value of DNSSuffixesToAppend for the instance
func (instance *CIM_DNSGeneralSettingData) SetPropertyDNSSuffixesToAppend(value []string) (err error) {
	return instance.SetProperty("DNSSuffixesToAppend", value)
}

// GetDNSSuffixesToAppend gets the value of DNSSuffixesToAppend for the instance
func (instance *CIM_DNSGeneralSettingData) GetPropertyDNSSuffixesToAppend() (value []string, err error) {
	retValue, err := instance.GetProperty("DNSSuffixesToAppend")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
