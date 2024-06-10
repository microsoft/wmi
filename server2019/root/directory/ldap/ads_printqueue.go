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

// ads_printqueue struct
type ads_printqueue struct {
	*ds_connectionpoint

	//
	DS_assetNumber string

	//
	DS_bytesPerMinute int32

	//
	DS_defaultPriority int32

	//
	DS_driverName string

	//
	DS_driverVersion int32

	//
	DS_location string

	//
	DS_operatingSystem string

	//
	DS_operatingSystemHotfix string

	//
	DS_operatingSystemServicePack string

	//
	DS_operatingSystemVersion string

	//
	DS_physicalLocationObject string

	//
	DS_portName []string

	//
	DS_printAttributes int32

	//
	DS_printBinNames []string

	//
	DS_printCollate bool

	//
	DS_printColor bool

	//
	DS_printDuplexSupported bool

	//
	DS_printEndTime int32

	//
	DS_printerName string

	//
	DS_printFormName string

	//
	DS_printKeepPrintedJobs bool

	//
	DS_printLanguage []string

	//
	DS_printMACAddress string

	//
	DS_printMaxCopies int32

	//
	DS_printMaxResolutionSupported int32

	//
	DS_printMaxXExtent int32

	//
	DS_printMaxYExtent int32

	//
	DS_printMediaReady []string

	//
	DS_printMediaSupported []string

	//
	DS_printMemory int32

	//
	DS_printMinXExtent int32

	//
	DS_printMinYExtent int32

	//
	DS_printNetworkAddress string

	//
	DS_printNotify string

	//
	DS_printNumberUp int32

	//
	DS_printOrientationsSupported []string

	//
	DS_printOwner string

	//
	DS_printPagesPerMinute int32

	//
	DS_printRate int32

	//
	DS_printRateUnit string

	//
	DS_printSeparatorFile string

	//
	DS_printShareName []string

	//
	DS_printSpooling string

	//
	DS_printStaplingSupported bool

	//
	DS_printStartTime int32

	//
	DS_printStatus string

	//
	DS_priority int32

	//
	DS_serverName string

	//
	DS_shortServerName string

	//
	DS_uNCName string

	//
	DS_versionNumber int32
}

func Newads_printqueueEx1(instance *cim.WmiInstance) (newInstance *ads_printqueue, err error) {
	tmp, err := Newds_connectionpointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_printqueue{
		ds_connectionpoint: tmp,
	}
	return
}

func Newads_printqueueEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_printqueue, err error) {
	tmp, err := Newds_connectionpointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_printqueue{
		ds_connectionpoint: tmp,
	}
	return
}

// SetDS_assetNumber sets the value of DS_assetNumber for the instance
func (instance *ads_printqueue) SetPropertyDS_assetNumber(value string) (err error) {
	return instance.SetProperty("DS_assetNumber", (value))
}

// GetDS_assetNumber gets the value of DS_assetNumber for the instance
func (instance *ads_printqueue) GetPropertyDS_assetNumber() (value string, err error) {
	retValue, err := instance.GetProperty("DS_assetNumber")
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

// SetDS_bytesPerMinute sets the value of DS_bytesPerMinute for the instance
func (instance *ads_printqueue) SetPropertyDS_bytesPerMinute(value int32) (err error) {
	return instance.SetProperty("DS_bytesPerMinute", (value))
}

// GetDS_bytesPerMinute gets the value of DS_bytesPerMinute for the instance
func (instance *ads_printqueue) GetPropertyDS_bytesPerMinute() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_bytesPerMinute")
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

// SetDS_defaultPriority sets the value of DS_defaultPriority for the instance
func (instance *ads_printqueue) SetPropertyDS_defaultPriority(value int32) (err error) {
	return instance.SetProperty("DS_defaultPriority", (value))
}

// GetDS_defaultPriority gets the value of DS_defaultPriority for the instance
func (instance *ads_printqueue) GetPropertyDS_defaultPriority() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_defaultPriority")
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

// SetDS_driverName sets the value of DS_driverName for the instance
func (instance *ads_printqueue) SetPropertyDS_driverName(value string) (err error) {
	return instance.SetProperty("DS_driverName", (value))
}

// GetDS_driverName gets the value of DS_driverName for the instance
func (instance *ads_printqueue) GetPropertyDS_driverName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_driverName")
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

// SetDS_driverVersion sets the value of DS_driverVersion for the instance
func (instance *ads_printqueue) SetPropertyDS_driverVersion(value int32) (err error) {
	return instance.SetProperty("DS_driverVersion", (value))
}

// GetDS_driverVersion gets the value of DS_driverVersion for the instance
func (instance *ads_printqueue) GetPropertyDS_driverVersion() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_driverVersion")
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

// SetDS_location sets the value of DS_location for the instance
func (instance *ads_printqueue) SetPropertyDS_location(value string) (err error) {
	return instance.SetProperty("DS_location", (value))
}

// GetDS_location gets the value of DS_location for the instance
func (instance *ads_printqueue) GetPropertyDS_location() (value string, err error) {
	retValue, err := instance.GetProperty("DS_location")
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

// SetDS_operatingSystem sets the value of DS_operatingSystem for the instance
func (instance *ads_printqueue) SetPropertyDS_operatingSystem(value string) (err error) {
	return instance.SetProperty("DS_operatingSystem", (value))
}

// GetDS_operatingSystem gets the value of DS_operatingSystem for the instance
func (instance *ads_printqueue) GetPropertyDS_operatingSystem() (value string, err error) {
	retValue, err := instance.GetProperty("DS_operatingSystem")
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

// SetDS_operatingSystemHotfix sets the value of DS_operatingSystemHotfix for the instance
func (instance *ads_printqueue) SetPropertyDS_operatingSystemHotfix(value string) (err error) {
	return instance.SetProperty("DS_operatingSystemHotfix", (value))
}

// GetDS_operatingSystemHotfix gets the value of DS_operatingSystemHotfix for the instance
func (instance *ads_printqueue) GetPropertyDS_operatingSystemHotfix() (value string, err error) {
	retValue, err := instance.GetProperty("DS_operatingSystemHotfix")
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

// SetDS_operatingSystemServicePack sets the value of DS_operatingSystemServicePack for the instance
func (instance *ads_printqueue) SetPropertyDS_operatingSystemServicePack(value string) (err error) {
	return instance.SetProperty("DS_operatingSystemServicePack", (value))
}

// GetDS_operatingSystemServicePack gets the value of DS_operatingSystemServicePack for the instance
func (instance *ads_printqueue) GetPropertyDS_operatingSystemServicePack() (value string, err error) {
	retValue, err := instance.GetProperty("DS_operatingSystemServicePack")
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

// SetDS_operatingSystemVersion sets the value of DS_operatingSystemVersion for the instance
func (instance *ads_printqueue) SetPropertyDS_operatingSystemVersion(value string) (err error) {
	return instance.SetProperty("DS_operatingSystemVersion", (value))
}

// GetDS_operatingSystemVersion gets the value of DS_operatingSystemVersion for the instance
func (instance *ads_printqueue) GetPropertyDS_operatingSystemVersion() (value string, err error) {
	retValue, err := instance.GetProperty("DS_operatingSystemVersion")
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

// SetDS_physicalLocationObject sets the value of DS_physicalLocationObject for the instance
func (instance *ads_printqueue) SetPropertyDS_physicalLocationObject(value string) (err error) {
	return instance.SetProperty("DS_physicalLocationObject", (value))
}

// GetDS_physicalLocationObject gets the value of DS_physicalLocationObject for the instance
func (instance *ads_printqueue) GetPropertyDS_physicalLocationObject() (value string, err error) {
	retValue, err := instance.GetProperty("DS_physicalLocationObject")
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

// SetDS_portName sets the value of DS_portName for the instance
func (instance *ads_printqueue) SetPropertyDS_portName(value []string) (err error) {
	return instance.SetProperty("DS_portName", (value))
}

// GetDS_portName gets the value of DS_portName for the instance
func (instance *ads_printqueue) GetPropertyDS_portName() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_portName")
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

// SetDS_printAttributes sets the value of DS_printAttributes for the instance
func (instance *ads_printqueue) SetPropertyDS_printAttributes(value int32) (err error) {
	return instance.SetProperty("DS_printAttributes", (value))
}

// GetDS_printAttributes gets the value of DS_printAttributes for the instance
func (instance *ads_printqueue) GetPropertyDS_printAttributes() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_printAttributes")
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

// SetDS_printBinNames sets the value of DS_printBinNames for the instance
func (instance *ads_printqueue) SetPropertyDS_printBinNames(value []string) (err error) {
	return instance.SetProperty("DS_printBinNames", (value))
}

// GetDS_printBinNames gets the value of DS_printBinNames for the instance
func (instance *ads_printqueue) GetPropertyDS_printBinNames() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_printBinNames")
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

// SetDS_printCollate sets the value of DS_printCollate for the instance
func (instance *ads_printqueue) SetPropertyDS_printCollate(value bool) (err error) {
	return instance.SetProperty("DS_printCollate", (value))
}

// GetDS_printCollate gets the value of DS_printCollate for the instance
func (instance *ads_printqueue) GetPropertyDS_printCollate() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_printCollate")
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

// SetDS_printColor sets the value of DS_printColor for the instance
func (instance *ads_printqueue) SetPropertyDS_printColor(value bool) (err error) {
	return instance.SetProperty("DS_printColor", (value))
}

// GetDS_printColor gets the value of DS_printColor for the instance
func (instance *ads_printqueue) GetPropertyDS_printColor() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_printColor")
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

// SetDS_printDuplexSupported sets the value of DS_printDuplexSupported for the instance
func (instance *ads_printqueue) SetPropertyDS_printDuplexSupported(value bool) (err error) {
	return instance.SetProperty("DS_printDuplexSupported", (value))
}

// GetDS_printDuplexSupported gets the value of DS_printDuplexSupported for the instance
func (instance *ads_printqueue) GetPropertyDS_printDuplexSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_printDuplexSupported")
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

// SetDS_printEndTime sets the value of DS_printEndTime for the instance
func (instance *ads_printqueue) SetPropertyDS_printEndTime(value int32) (err error) {
	return instance.SetProperty("DS_printEndTime", (value))
}

// GetDS_printEndTime gets the value of DS_printEndTime for the instance
func (instance *ads_printqueue) GetPropertyDS_printEndTime() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_printEndTime")
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

// SetDS_printerName sets the value of DS_printerName for the instance
func (instance *ads_printqueue) SetPropertyDS_printerName(value string) (err error) {
	return instance.SetProperty("DS_printerName", (value))
}

// GetDS_printerName gets the value of DS_printerName for the instance
func (instance *ads_printqueue) GetPropertyDS_printerName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_printerName")
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

// SetDS_printFormName sets the value of DS_printFormName for the instance
func (instance *ads_printqueue) SetPropertyDS_printFormName(value string) (err error) {
	return instance.SetProperty("DS_printFormName", (value))
}

// GetDS_printFormName gets the value of DS_printFormName for the instance
func (instance *ads_printqueue) GetPropertyDS_printFormName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_printFormName")
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

// SetDS_printKeepPrintedJobs sets the value of DS_printKeepPrintedJobs for the instance
func (instance *ads_printqueue) SetPropertyDS_printKeepPrintedJobs(value bool) (err error) {
	return instance.SetProperty("DS_printKeepPrintedJobs", (value))
}

// GetDS_printKeepPrintedJobs gets the value of DS_printKeepPrintedJobs for the instance
func (instance *ads_printqueue) GetPropertyDS_printKeepPrintedJobs() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_printKeepPrintedJobs")
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

// SetDS_printLanguage sets the value of DS_printLanguage for the instance
func (instance *ads_printqueue) SetPropertyDS_printLanguage(value []string) (err error) {
	return instance.SetProperty("DS_printLanguage", (value))
}

// GetDS_printLanguage gets the value of DS_printLanguage for the instance
func (instance *ads_printqueue) GetPropertyDS_printLanguage() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_printLanguage")
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

// SetDS_printMACAddress sets the value of DS_printMACAddress for the instance
func (instance *ads_printqueue) SetPropertyDS_printMACAddress(value string) (err error) {
	return instance.SetProperty("DS_printMACAddress", (value))
}

// GetDS_printMACAddress gets the value of DS_printMACAddress for the instance
func (instance *ads_printqueue) GetPropertyDS_printMACAddress() (value string, err error) {
	retValue, err := instance.GetProperty("DS_printMACAddress")
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

// SetDS_printMaxCopies sets the value of DS_printMaxCopies for the instance
func (instance *ads_printqueue) SetPropertyDS_printMaxCopies(value int32) (err error) {
	return instance.SetProperty("DS_printMaxCopies", (value))
}

// GetDS_printMaxCopies gets the value of DS_printMaxCopies for the instance
func (instance *ads_printqueue) GetPropertyDS_printMaxCopies() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_printMaxCopies")
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

// SetDS_printMaxResolutionSupported sets the value of DS_printMaxResolutionSupported for the instance
func (instance *ads_printqueue) SetPropertyDS_printMaxResolutionSupported(value int32) (err error) {
	return instance.SetProperty("DS_printMaxResolutionSupported", (value))
}

// GetDS_printMaxResolutionSupported gets the value of DS_printMaxResolutionSupported for the instance
func (instance *ads_printqueue) GetPropertyDS_printMaxResolutionSupported() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_printMaxResolutionSupported")
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

// SetDS_printMaxXExtent sets the value of DS_printMaxXExtent for the instance
func (instance *ads_printqueue) SetPropertyDS_printMaxXExtent(value int32) (err error) {
	return instance.SetProperty("DS_printMaxXExtent", (value))
}

// GetDS_printMaxXExtent gets the value of DS_printMaxXExtent for the instance
func (instance *ads_printqueue) GetPropertyDS_printMaxXExtent() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_printMaxXExtent")
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

// SetDS_printMaxYExtent sets the value of DS_printMaxYExtent for the instance
func (instance *ads_printqueue) SetPropertyDS_printMaxYExtent(value int32) (err error) {
	return instance.SetProperty("DS_printMaxYExtent", (value))
}

// GetDS_printMaxYExtent gets the value of DS_printMaxYExtent for the instance
func (instance *ads_printqueue) GetPropertyDS_printMaxYExtent() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_printMaxYExtent")
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

// SetDS_printMediaReady sets the value of DS_printMediaReady for the instance
func (instance *ads_printqueue) SetPropertyDS_printMediaReady(value []string) (err error) {
	return instance.SetProperty("DS_printMediaReady", (value))
}

// GetDS_printMediaReady gets the value of DS_printMediaReady for the instance
func (instance *ads_printqueue) GetPropertyDS_printMediaReady() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_printMediaReady")
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

// SetDS_printMediaSupported sets the value of DS_printMediaSupported for the instance
func (instance *ads_printqueue) SetPropertyDS_printMediaSupported(value []string) (err error) {
	return instance.SetProperty("DS_printMediaSupported", (value))
}

// GetDS_printMediaSupported gets the value of DS_printMediaSupported for the instance
func (instance *ads_printqueue) GetPropertyDS_printMediaSupported() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_printMediaSupported")
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

// SetDS_printMemory sets the value of DS_printMemory for the instance
func (instance *ads_printqueue) SetPropertyDS_printMemory(value int32) (err error) {
	return instance.SetProperty("DS_printMemory", (value))
}

// GetDS_printMemory gets the value of DS_printMemory for the instance
func (instance *ads_printqueue) GetPropertyDS_printMemory() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_printMemory")
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

// SetDS_printMinXExtent sets the value of DS_printMinXExtent for the instance
func (instance *ads_printqueue) SetPropertyDS_printMinXExtent(value int32) (err error) {
	return instance.SetProperty("DS_printMinXExtent", (value))
}

// GetDS_printMinXExtent gets the value of DS_printMinXExtent for the instance
func (instance *ads_printqueue) GetPropertyDS_printMinXExtent() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_printMinXExtent")
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

// SetDS_printMinYExtent sets the value of DS_printMinYExtent for the instance
func (instance *ads_printqueue) SetPropertyDS_printMinYExtent(value int32) (err error) {
	return instance.SetProperty("DS_printMinYExtent", (value))
}

// GetDS_printMinYExtent gets the value of DS_printMinYExtent for the instance
func (instance *ads_printqueue) GetPropertyDS_printMinYExtent() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_printMinYExtent")
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

// SetDS_printNetworkAddress sets the value of DS_printNetworkAddress for the instance
func (instance *ads_printqueue) SetPropertyDS_printNetworkAddress(value string) (err error) {
	return instance.SetProperty("DS_printNetworkAddress", (value))
}

// GetDS_printNetworkAddress gets the value of DS_printNetworkAddress for the instance
func (instance *ads_printqueue) GetPropertyDS_printNetworkAddress() (value string, err error) {
	retValue, err := instance.GetProperty("DS_printNetworkAddress")
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

// SetDS_printNotify sets the value of DS_printNotify for the instance
func (instance *ads_printqueue) SetPropertyDS_printNotify(value string) (err error) {
	return instance.SetProperty("DS_printNotify", (value))
}

// GetDS_printNotify gets the value of DS_printNotify for the instance
func (instance *ads_printqueue) GetPropertyDS_printNotify() (value string, err error) {
	retValue, err := instance.GetProperty("DS_printNotify")
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

// SetDS_printNumberUp sets the value of DS_printNumberUp for the instance
func (instance *ads_printqueue) SetPropertyDS_printNumberUp(value int32) (err error) {
	return instance.SetProperty("DS_printNumberUp", (value))
}

// GetDS_printNumberUp gets the value of DS_printNumberUp for the instance
func (instance *ads_printqueue) GetPropertyDS_printNumberUp() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_printNumberUp")
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

// SetDS_printOrientationsSupported sets the value of DS_printOrientationsSupported for the instance
func (instance *ads_printqueue) SetPropertyDS_printOrientationsSupported(value []string) (err error) {
	return instance.SetProperty("DS_printOrientationsSupported", (value))
}

// GetDS_printOrientationsSupported gets the value of DS_printOrientationsSupported for the instance
func (instance *ads_printqueue) GetPropertyDS_printOrientationsSupported() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_printOrientationsSupported")
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

// SetDS_printOwner sets the value of DS_printOwner for the instance
func (instance *ads_printqueue) SetPropertyDS_printOwner(value string) (err error) {
	return instance.SetProperty("DS_printOwner", (value))
}

// GetDS_printOwner gets the value of DS_printOwner for the instance
func (instance *ads_printqueue) GetPropertyDS_printOwner() (value string, err error) {
	retValue, err := instance.GetProperty("DS_printOwner")
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

// SetDS_printPagesPerMinute sets the value of DS_printPagesPerMinute for the instance
func (instance *ads_printqueue) SetPropertyDS_printPagesPerMinute(value int32) (err error) {
	return instance.SetProperty("DS_printPagesPerMinute", (value))
}

// GetDS_printPagesPerMinute gets the value of DS_printPagesPerMinute for the instance
func (instance *ads_printqueue) GetPropertyDS_printPagesPerMinute() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_printPagesPerMinute")
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

// SetDS_printRate sets the value of DS_printRate for the instance
func (instance *ads_printqueue) SetPropertyDS_printRate(value int32) (err error) {
	return instance.SetProperty("DS_printRate", (value))
}

// GetDS_printRate gets the value of DS_printRate for the instance
func (instance *ads_printqueue) GetPropertyDS_printRate() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_printRate")
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

// SetDS_printRateUnit sets the value of DS_printRateUnit for the instance
func (instance *ads_printqueue) SetPropertyDS_printRateUnit(value string) (err error) {
	return instance.SetProperty("DS_printRateUnit", (value))
}

// GetDS_printRateUnit gets the value of DS_printRateUnit for the instance
func (instance *ads_printqueue) GetPropertyDS_printRateUnit() (value string, err error) {
	retValue, err := instance.GetProperty("DS_printRateUnit")
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

// SetDS_printSeparatorFile sets the value of DS_printSeparatorFile for the instance
func (instance *ads_printqueue) SetPropertyDS_printSeparatorFile(value string) (err error) {
	return instance.SetProperty("DS_printSeparatorFile", (value))
}

// GetDS_printSeparatorFile gets the value of DS_printSeparatorFile for the instance
func (instance *ads_printqueue) GetPropertyDS_printSeparatorFile() (value string, err error) {
	retValue, err := instance.GetProperty("DS_printSeparatorFile")
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

// SetDS_printShareName sets the value of DS_printShareName for the instance
func (instance *ads_printqueue) SetPropertyDS_printShareName(value []string) (err error) {
	return instance.SetProperty("DS_printShareName", (value))
}

// GetDS_printShareName gets the value of DS_printShareName for the instance
func (instance *ads_printqueue) GetPropertyDS_printShareName() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_printShareName")
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

// SetDS_printSpooling sets the value of DS_printSpooling for the instance
func (instance *ads_printqueue) SetPropertyDS_printSpooling(value string) (err error) {
	return instance.SetProperty("DS_printSpooling", (value))
}

// GetDS_printSpooling gets the value of DS_printSpooling for the instance
func (instance *ads_printqueue) GetPropertyDS_printSpooling() (value string, err error) {
	retValue, err := instance.GetProperty("DS_printSpooling")
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

// SetDS_printStaplingSupported sets the value of DS_printStaplingSupported for the instance
func (instance *ads_printqueue) SetPropertyDS_printStaplingSupported(value bool) (err error) {
	return instance.SetProperty("DS_printStaplingSupported", (value))
}

// GetDS_printStaplingSupported gets the value of DS_printStaplingSupported for the instance
func (instance *ads_printqueue) GetPropertyDS_printStaplingSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_printStaplingSupported")
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

// SetDS_printStartTime sets the value of DS_printStartTime for the instance
func (instance *ads_printqueue) SetPropertyDS_printStartTime(value int32) (err error) {
	return instance.SetProperty("DS_printStartTime", (value))
}

// GetDS_printStartTime gets the value of DS_printStartTime for the instance
func (instance *ads_printqueue) GetPropertyDS_printStartTime() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_printStartTime")
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

// SetDS_printStatus sets the value of DS_printStatus for the instance
func (instance *ads_printqueue) SetPropertyDS_printStatus(value string) (err error) {
	return instance.SetProperty("DS_printStatus", (value))
}

// GetDS_printStatus gets the value of DS_printStatus for the instance
func (instance *ads_printqueue) GetPropertyDS_printStatus() (value string, err error) {
	retValue, err := instance.GetProperty("DS_printStatus")
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

// SetDS_priority sets the value of DS_priority for the instance
func (instance *ads_printqueue) SetPropertyDS_priority(value int32) (err error) {
	return instance.SetProperty("DS_priority", (value))
}

// GetDS_priority gets the value of DS_priority for the instance
func (instance *ads_printqueue) GetPropertyDS_priority() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_priority")
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

// SetDS_serverName sets the value of DS_serverName for the instance
func (instance *ads_printqueue) SetPropertyDS_serverName(value string) (err error) {
	return instance.SetProperty("DS_serverName", (value))
}

// GetDS_serverName gets the value of DS_serverName for the instance
func (instance *ads_printqueue) GetPropertyDS_serverName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_serverName")
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

// SetDS_shortServerName sets the value of DS_shortServerName for the instance
func (instance *ads_printqueue) SetPropertyDS_shortServerName(value string) (err error) {
	return instance.SetProperty("DS_shortServerName", (value))
}

// GetDS_shortServerName gets the value of DS_shortServerName for the instance
func (instance *ads_printqueue) GetPropertyDS_shortServerName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_shortServerName")
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

// SetDS_uNCName sets the value of DS_uNCName for the instance
func (instance *ads_printqueue) SetPropertyDS_uNCName(value string) (err error) {
	return instance.SetProperty("DS_uNCName", (value))
}

// GetDS_uNCName gets the value of DS_uNCName for the instance
func (instance *ads_printqueue) GetPropertyDS_uNCName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_uNCName")
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

// SetDS_versionNumber sets the value of DS_versionNumber for the instance
func (instance *ads_printqueue) SetPropertyDS_versionNumber(value int32) (err error) {
	return instance.SetProperty("DS_versionNumber", (value))
}

// GetDS_versionNumber gets the value of DS_versionNumber for the instance
func (instance *ads_printqueue) GetPropertyDS_versionNumber() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_versionNumber")
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
