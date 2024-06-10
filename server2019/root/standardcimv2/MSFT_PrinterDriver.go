// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_PrinterDriver struct
type MSFT_PrinterDriver struct {
	*CIM_ManagedSystemElement

	//
	ColorProfiles []string

	//
	ComputerName string

	//
	ConfigFile string

	//
	CoreDriverDependencies []string

	//
	DataFile string

	//
	Date string

	//
	DefaultDatatype string

	//
	DependentFiles []string

	//
	DriverVersion uint64

	//
	HardwareID string

	//
	HelpFile string

	//
	InfPath string

	//
	IsPackageAware bool

	//
	MajorVersion uint32

	//
	Manufacturer string

	//
	Monitor string

	//
	OEMUrl string

	//
	Path string

	//
	PreviousCompatibleNames []string

	//
	PrinterEnvironment string

	//
	PrintProcessor string

	//
	provider string

	//
	VendorSetup string
}

func NewMSFT_PrinterDriverEx1(instance *cim.WmiInstance) (newInstance *MSFT_PrinterDriver, err error) {
	tmp, err := NewCIM_ManagedSystemElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_PrinterDriver{
		CIM_ManagedSystemElement: tmp,
	}
	return
}

func NewMSFT_PrinterDriverEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_PrinterDriver, err error) {
	tmp, err := NewCIM_ManagedSystemElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_PrinterDriver{
		CIM_ManagedSystemElement: tmp,
	}
	return
}

// SetColorProfiles sets the value of ColorProfiles for the instance
func (instance *MSFT_PrinterDriver) SetPropertyColorProfiles(value []string) (err error) {
	return instance.SetProperty("ColorProfiles", (value))
}

// GetColorProfiles gets the value of ColorProfiles for the instance
func (instance *MSFT_PrinterDriver) GetPropertyColorProfiles() (value []string, err error) {
	retValue, err := instance.GetProperty("ColorProfiles")
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

// SetComputerName sets the value of ComputerName for the instance
func (instance *MSFT_PrinterDriver) SetPropertyComputerName(value string) (err error) {
	return instance.SetProperty("ComputerName", (value))
}

// GetComputerName gets the value of ComputerName for the instance
func (instance *MSFT_PrinterDriver) GetPropertyComputerName() (value string, err error) {
	retValue, err := instance.GetProperty("ComputerName")
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

// SetConfigFile sets the value of ConfigFile for the instance
func (instance *MSFT_PrinterDriver) SetPropertyConfigFile(value string) (err error) {
	return instance.SetProperty("ConfigFile", (value))
}

// GetConfigFile gets the value of ConfigFile for the instance
func (instance *MSFT_PrinterDriver) GetPropertyConfigFile() (value string, err error) {
	retValue, err := instance.GetProperty("ConfigFile")
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

// SetCoreDriverDependencies sets the value of CoreDriverDependencies for the instance
func (instance *MSFT_PrinterDriver) SetPropertyCoreDriverDependencies(value []string) (err error) {
	return instance.SetProperty("CoreDriverDependencies", (value))
}

// GetCoreDriverDependencies gets the value of CoreDriverDependencies for the instance
func (instance *MSFT_PrinterDriver) GetPropertyCoreDriverDependencies() (value []string, err error) {
	retValue, err := instance.GetProperty("CoreDriverDependencies")
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

// SetDataFile sets the value of DataFile for the instance
func (instance *MSFT_PrinterDriver) SetPropertyDataFile(value string) (err error) {
	return instance.SetProperty("DataFile", (value))
}

// GetDataFile gets the value of DataFile for the instance
func (instance *MSFT_PrinterDriver) GetPropertyDataFile() (value string, err error) {
	retValue, err := instance.GetProperty("DataFile")
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

// SetDate sets the value of Date for the instance
func (instance *MSFT_PrinterDriver) SetPropertyDate(value string) (err error) {
	return instance.SetProperty("Date", (value))
}

// GetDate gets the value of Date for the instance
func (instance *MSFT_PrinterDriver) GetPropertyDate() (value string, err error) {
	retValue, err := instance.GetProperty("Date")
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

// SetDefaultDatatype sets the value of DefaultDatatype for the instance
func (instance *MSFT_PrinterDriver) SetPropertyDefaultDatatype(value string) (err error) {
	return instance.SetProperty("DefaultDatatype", (value))
}

// GetDefaultDatatype gets the value of DefaultDatatype for the instance
func (instance *MSFT_PrinterDriver) GetPropertyDefaultDatatype() (value string, err error) {
	retValue, err := instance.GetProperty("DefaultDatatype")
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

// SetDependentFiles sets the value of DependentFiles for the instance
func (instance *MSFT_PrinterDriver) SetPropertyDependentFiles(value []string) (err error) {
	return instance.SetProperty("DependentFiles", (value))
}

// GetDependentFiles gets the value of DependentFiles for the instance
func (instance *MSFT_PrinterDriver) GetPropertyDependentFiles() (value []string, err error) {
	retValue, err := instance.GetProperty("DependentFiles")
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

// SetDriverVersion sets the value of DriverVersion for the instance
func (instance *MSFT_PrinterDriver) SetPropertyDriverVersion(value uint64) (err error) {
	return instance.SetProperty("DriverVersion", (value))
}

// GetDriverVersion gets the value of DriverVersion for the instance
func (instance *MSFT_PrinterDriver) GetPropertyDriverVersion() (value uint64, err error) {
	retValue, err := instance.GetProperty("DriverVersion")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetHardwareID sets the value of HardwareID for the instance
func (instance *MSFT_PrinterDriver) SetPropertyHardwareID(value string) (err error) {
	return instance.SetProperty("HardwareID", (value))
}

// GetHardwareID gets the value of HardwareID for the instance
func (instance *MSFT_PrinterDriver) GetPropertyHardwareID() (value string, err error) {
	retValue, err := instance.GetProperty("HardwareID")
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

// SetHelpFile sets the value of HelpFile for the instance
func (instance *MSFT_PrinterDriver) SetPropertyHelpFile(value string) (err error) {
	return instance.SetProperty("HelpFile", (value))
}

// GetHelpFile gets the value of HelpFile for the instance
func (instance *MSFT_PrinterDriver) GetPropertyHelpFile() (value string, err error) {
	retValue, err := instance.GetProperty("HelpFile")
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

// SetInfPath sets the value of InfPath for the instance
func (instance *MSFT_PrinterDriver) SetPropertyInfPath(value string) (err error) {
	return instance.SetProperty("InfPath", (value))
}

// GetInfPath gets the value of InfPath for the instance
func (instance *MSFT_PrinterDriver) GetPropertyInfPath() (value string, err error) {
	retValue, err := instance.GetProperty("InfPath")
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

// SetIsPackageAware sets the value of IsPackageAware for the instance
func (instance *MSFT_PrinterDriver) SetPropertyIsPackageAware(value bool) (err error) {
	return instance.SetProperty("IsPackageAware", (value))
}

// GetIsPackageAware gets the value of IsPackageAware for the instance
func (instance *MSFT_PrinterDriver) GetPropertyIsPackageAware() (value bool, err error) {
	retValue, err := instance.GetProperty("IsPackageAware")
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

// SetMajorVersion sets the value of MajorVersion for the instance
func (instance *MSFT_PrinterDriver) SetPropertyMajorVersion(value uint32) (err error) {
	return instance.SetProperty("MajorVersion", (value))
}

// GetMajorVersion gets the value of MajorVersion for the instance
func (instance *MSFT_PrinterDriver) GetPropertyMajorVersion() (value uint32, err error) {
	retValue, err := instance.GetProperty("MajorVersion")
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

// SetManufacturer sets the value of Manufacturer for the instance
func (instance *MSFT_PrinterDriver) SetPropertyManufacturer(value string) (err error) {
	return instance.SetProperty("Manufacturer", (value))
}

// GetManufacturer gets the value of Manufacturer for the instance
func (instance *MSFT_PrinterDriver) GetPropertyManufacturer() (value string, err error) {
	retValue, err := instance.GetProperty("Manufacturer")
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

// SetMonitor sets the value of Monitor for the instance
func (instance *MSFT_PrinterDriver) SetPropertyMonitor(value string) (err error) {
	return instance.SetProperty("Monitor", (value))
}

// GetMonitor gets the value of Monitor for the instance
func (instance *MSFT_PrinterDriver) GetPropertyMonitor() (value string, err error) {
	retValue, err := instance.GetProperty("Monitor")
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

// SetOEMUrl sets the value of OEMUrl for the instance
func (instance *MSFT_PrinterDriver) SetPropertyOEMUrl(value string) (err error) {
	return instance.SetProperty("OEMUrl", (value))
}

// GetOEMUrl gets the value of OEMUrl for the instance
func (instance *MSFT_PrinterDriver) GetPropertyOEMUrl() (value string, err error) {
	retValue, err := instance.GetProperty("OEMUrl")
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

// SetPath sets the value of Path for the instance
func (instance *MSFT_PrinterDriver) SetPropertyPath(value string) (err error) {
	return instance.SetProperty("Path", (value))
}

// GetPath gets the value of Path for the instance
func (instance *MSFT_PrinterDriver) GetPropertyPath() (value string, err error) {
	retValue, err := instance.GetProperty("Path")
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

// SetPreviousCompatibleNames sets the value of PreviousCompatibleNames for the instance
func (instance *MSFT_PrinterDriver) SetPropertyPreviousCompatibleNames(value []string) (err error) {
	return instance.SetProperty("PreviousCompatibleNames", (value))
}

// GetPreviousCompatibleNames gets the value of PreviousCompatibleNames for the instance
func (instance *MSFT_PrinterDriver) GetPropertyPreviousCompatibleNames() (value []string, err error) {
	retValue, err := instance.GetProperty("PreviousCompatibleNames")
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

// SetPrinterEnvironment sets the value of PrinterEnvironment for the instance
func (instance *MSFT_PrinterDriver) SetPropertyPrinterEnvironment(value string) (err error) {
	return instance.SetProperty("PrinterEnvironment", (value))
}

// GetPrinterEnvironment gets the value of PrinterEnvironment for the instance
func (instance *MSFT_PrinterDriver) GetPropertyPrinterEnvironment() (value string, err error) {
	retValue, err := instance.GetProperty("PrinterEnvironment")
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

// SetPrintProcessor sets the value of PrintProcessor for the instance
func (instance *MSFT_PrinterDriver) SetPropertyPrintProcessor(value string) (err error) {
	return instance.SetProperty("PrintProcessor", (value))
}

// GetPrintProcessor gets the value of PrintProcessor for the instance
func (instance *MSFT_PrinterDriver) GetPropertyPrintProcessor() (value string, err error) {
	retValue, err := instance.GetProperty("PrintProcessor")
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

// Setprovider sets the value of provider for the instance
func (instance *MSFT_PrinterDriver) SetPropertyprovider(value string) (err error) {
	return instance.SetProperty("provider", (value))
}

// Getprovider gets the value of provider for the instance
func (instance *MSFT_PrinterDriver) GetPropertyprovider() (value string, err error) {
	retValue, err := instance.GetProperty("provider")
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

// SetVendorSetup sets the value of VendorSetup for the instance
func (instance *MSFT_PrinterDriver) SetPropertyVendorSetup(value string) (err error) {
	return instance.SetProperty("VendorSetup", (value))
}

// GetVendorSetup gets the value of VendorSetup for the instance
func (instance *MSFT_PrinterDriver) GetPropertyVendorSetup() (value string, err error) {
	retValue, err := instance.GetProperty("VendorSetup")
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

// <param name="ComputerName" type="string "></param>
// <param name="InfPath" type="string "></param>
// <param name="Name" type="string "></param>
// <param name="PrinterEnvironment" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PrinterDriver) Add( /* IN */ Name string,
	/* IN */ InfPath string,
	/* IN */ PrinterEnvironment string,
	/* IN */ ComputerName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Add", Name, InfPath, PrinterEnvironment, ComputerName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
