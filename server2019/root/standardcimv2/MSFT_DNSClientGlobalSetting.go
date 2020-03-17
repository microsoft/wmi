// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_DNSClientGlobalSetting struct
type MSFT_DNSClientGlobalSetting struct {
	CIM_DNSGeneralSettingData

	// 705
	DevolutionLevel uint32

	// 704
	SuffixSearchList []string

	// 703
	UseDevolution bool

	// 702
	UseSuffixSearchList bool
}

// SetDevolutionLevel sets the value of DevolutionLevel for the instance
func (instance *MSFT_DNSClientGlobalSetting) SetPropertyDevolutionLevel(value uint32) (err error) {
	return instance.SetProperty("DevolutionLevel", value)
}

// GetDevolutionLevel gets the value of DevolutionLevel for the instance
func (instance *MSFT_DNSClientGlobalSetting) GetPropertyDevolutionLevel() (value uint32, err error) {
	retValue, err := instance.GetProperty("DevolutionLevel")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSuffixSearchList sets the value of SuffixSearchList for the instance
func (instance *MSFT_DNSClientGlobalSetting) SetPropertySuffixSearchList(value []string) (err error) {
	return instance.SetProperty("SuffixSearchList", value)
}

// GetSuffixSearchList gets the value of SuffixSearchList for the instance
func (instance *MSFT_DNSClientGlobalSetting) GetPropertySuffixSearchList() (value []string, err error) {
	retValue, err := instance.GetProperty("SuffixSearchList")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUseDevolution sets the value of UseDevolution for the instance
func (instance *MSFT_DNSClientGlobalSetting) SetPropertyUseDevolution(value bool) (err error) {
	return instance.SetProperty("UseDevolution", value)
}

// GetUseDevolution gets the value of UseDevolution for the instance
func (instance *MSFT_DNSClientGlobalSetting) GetPropertyUseDevolution() (value bool, err error) {
	retValue, err := instance.GetProperty("UseDevolution")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUseSuffixSearchList sets the value of UseSuffixSearchList for the instance
func (instance *MSFT_DNSClientGlobalSetting) SetPropertyUseSuffixSearchList(value bool) (err error) {
	return instance.SetProperty("UseSuffixSearchList", value)
}

// GetUseSuffixSearchList gets the value of UseSuffixSearchList for the instance
func (instance *MSFT_DNSClientGlobalSetting) GetPropertyUseSuffixSearchList() (value bool, err error) {
	retValue, err := instance.GetProperty("UseSuffixSearchList")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
