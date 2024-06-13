// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_USBDevice struct
type CIM_USBDevice struct {
	*CIM_LogicalDevice

	// Indicates the USB class code.
	ClassCode uint8

	// CommandTimeout is configurable by management applications supporting USB Redirections. When the Redirection Service redirects a USBDevice command to a remote device, and the remote device does not respond before CommandTimout times out, the Redirection Service will emulate a media eject event and re-try the command and/or try to re-establish the connection to the remote device. The timeout is expressed using the interval format of the datetime type.
	CommandTimeout string

	// An array of USB 'alternate settings' for each interface in the currently selected configuration (indicated by the CurrentConfigValue property). This array has one entry for each interface in the configuration. If the property, CurrentConfigValue, is zero (indicating the Device is not configured), the array is undefined. To understand how to parse this octet string, refer to the USB Specification.
	CurrentAlternateSettings []uint8

	// Indicates the configuration currently selected for the Device. If this value is zero, the Device is unconfigured.
	CurrentConfigValue uint8

	// From the USB specification Device Descriptor, Device Release Number in Binary-Coded Decimal.
	DeviceReleaseNumber uint16

	// From the USB specification Device Descriptior, Manufacturer string.
	Manufacturer string

	// From the USB specification Device Descriptor, Maximum Packet size for the USB zero endpoint. Valid sizes are 8, 16, 32, 64.
	MaxPacketSize uint8

	// Number of device configurations that are defined for the Device.
	NumberOfConfigs uint8

	// From the USB specification Device Descriptor, Product String.
	Product string

	// From the USB specification Device Descriptor, Product ID assigned by manufacturer.
	ProductID uint16

	// Indicates the USB protocol code.
	ProtocolCode uint8

	// From the USB specification Device Descriptor, Serial Number String.
	SerialNumber string

	// Indicates the USB subclass code.
	SubclassCode uint8

	// Indicates the latest USB Version supported by the USB Device. The property is expressed as a Binary-Coded Decimal (BCD) where a decimal point is implied between the 2nd and 3rd digits. For example, a value of 0x201 indicates that version 2.01 is supported.
	USBVersion uint16

	// From the USB specification Device Descriptor, where 'bcdUSB' is the USB Specification Number, in Binary-Coded Decimal format, that the device complies with.
	USBVersionInBCD uint16

	// From the USB specification Device Descriptor, Vendor ID assigned by USB.org.
	VendorID uint16
}

func NewCIM_USBDeviceEx1(instance *cim.WmiInstance) (newInstance *CIM_USBDevice, err error) {
	tmp, err := NewCIM_LogicalDeviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_USBDevice{
		CIM_LogicalDevice: tmp,
	}
	return
}

func NewCIM_USBDeviceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_USBDevice, err error) {
	tmp, err := NewCIM_LogicalDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_USBDevice{
		CIM_LogicalDevice: tmp,
	}
	return
}

// SetClassCode sets the value of ClassCode for the instance
func (instance *CIM_USBDevice) SetPropertyClassCode(value uint8) (err error) {
	return instance.SetProperty("ClassCode", (value))
}

// GetClassCode gets the value of ClassCode for the instance
func (instance *CIM_USBDevice) GetPropertyClassCode() (value uint8, err error) {
	retValue, err := instance.GetProperty("ClassCode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetCommandTimeout sets the value of CommandTimeout for the instance
func (instance *CIM_USBDevice) SetPropertyCommandTimeout(value string) (err error) {
	return instance.SetProperty("CommandTimeout", (value))
}

// GetCommandTimeout gets the value of CommandTimeout for the instance
func (instance *CIM_USBDevice) GetPropertyCommandTimeout() (value string, err error) {
	retValue, err := instance.GetProperty("CommandTimeout")
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

// SetCurrentAlternateSettings sets the value of CurrentAlternateSettings for the instance
func (instance *CIM_USBDevice) SetPropertyCurrentAlternateSettings(value []uint8) (err error) {
	return instance.SetProperty("CurrentAlternateSettings", (value))
}

// GetCurrentAlternateSettings gets the value of CurrentAlternateSettings for the instance
func (instance *CIM_USBDevice) GetPropertyCurrentAlternateSettings() (value []uint8, err error) {
	retValue, err := instance.GetProperty("CurrentAlternateSettings")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetCurrentConfigValue sets the value of CurrentConfigValue for the instance
func (instance *CIM_USBDevice) SetPropertyCurrentConfigValue(value uint8) (err error) {
	return instance.SetProperty("CurrentConfigValue", (value))
}

// GetCurrentConfigValue gets the value of CurrentConfigValue for the instance
func (instance *CIM_USBDevice) GetPropertyCurrentConfigValue() (value uint8, err error) {
	retValue, err := instance.GetProperty("CurrentConfigValue")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetDeviceReleaseNumber sets the value of DeviceReleaseNumber for the instance
func (instance *CIM_USBDevice) SetPropertyDeviceReleaseNumber(value uint16) (err error) {
	return instance.SetProperty("DeviceReleaseNumber", (value))
}

// GetDeviceReleaseNumber gets the value of DeviceReleaseNumber for the instance
func (instance *CIM_USBDevice) GetPropertyDeviceReleaseNumber() (value uint16, err error) {
	retValue, err := instance.GetProperty("DeviceReleaseNumber")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetManufacturer sets the value of Manufacturer for the instance
func (instance *CIM_USBDevice) SetPropertyManufacturer(value string) (err error) {
	return instance.SetProperty("Manufacturer", (value))
}

// GetManufacturer gets the value of Manufacturer for the instance
func (instance *CIM_USBDevice) GetPropertyManufacturer() (value string, err error) {
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

// SetMaxPacketSize sets the value of MaxPacketSize for the instance
func (instance *CIM_USBDevice) SetPropertyMaxPacketSize(value uint8) (err error) {
	return instance.SetProperty("MaxPacketSize", (value))
}

// GetMaxPacketSize gets the value of MaxPacketSize for the instance
func (instance *CIM_USBDevice) GetPropertyMaxPacketSize() (value uint8, err error) {
	retValue, err := instance.GetProperty("MaxPacketSize")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetNumberOfConfigs sets the value of NumberOfConfigs for the instance
func (instance *CIM_USBDevice) SetPropertyNumberOfConfigs(value uint8) (err error) {
	return instance.SetProperty("NumberOfConfigs", (value))
}

// GetNumberOfConfigs gets the value of NumberOfConfigs for the instance
func (instance *CIM_USBDevice) GetPropertyNumberOfConfigs() (value uint8, err error) {
	retValue, err := instance.GetProperty("NumberOfConfigs")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetProduct sets the value of Product for the instance
func (instance *CIM_USBDevice) SetPropertyProduct(value string) (err error) {
	return instance.SetProperty("Product", (value))
}

// GetProduct gets the value of Product for the instance
func (instance *CIM_USBDevice) GetPropertyProduct() (value string, err error) {
	retValue, err := instance.GetProperty("Product")
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

// SetProductID sets the value of ProductID for the instance
func (instance *CIM_USBDevice) SetPropertyProductID(value uint16) (err error) {
	return instance.SetProperty("ProductID", (value))
}

// GetProductID gets the value of ProductID for the instance
func (instance *CIM_USBDevice) GetPropertyProductID() (value uint16, err error) {
	retValue, err := instance.GetProperty("ProductID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetProtocolCode sets the value of ProtocolCode for the instance
func (instance *CIM_USBDevice) SetPropertyProtocolCode(value uint8) (err error) {
	return instance.SetProperty("ProtocolCode", (value))
}

// GetProtocolCode gets the value of ProtocolCode for the instance
func (instance *CIM_USBDevice) GetPropertyProtocolCode() (value uint8, err error) {
	retValue, err := instance.GetProperty("ProtocolCode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetSerialNumber sets the value of SerialNumber for the instance
func (instance *CIM_USBDevice) SetPropertySerialNumber(value string) (err error) {
	return instance.SetProperty("SerialNumber", (value))
}

// GetSerialNumber gets the value of SerialNumber for the instance
func (instance *CIM_USBDevice) GetPropertySerialNumber() (value string, err error) {
	retValue, err := instance.GetProperty("SerialNumber")
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

// SetSubclassCode sets the value of SubclassCode for the instance
func (instance *CIM_USBDevice) SetPropertySubclassCode(value uint8) (err error) {
	return instance.SetProperty("SubclassCode", (value))
}

// GetSubclassCode gets the value of SubclassCode for the instance
func (instance *CIM_USBDevice) GetPropertySubclassCode() (value uint8, err error) {
	retValue, err := instance.GetProperty("SubclassCode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetUSBVersion sets the value of USBVersion for the instance
func (instance *CIM_USBDevice) SetPropertyUSBVersion(value uint16) (err error) {
	return instance.SetProperty("USBVersion", (value))
}

// GetUSBVersion gets the value of USBVersion for the instance
func (instance *CIM_USBDevice) GetPropertyUSBVersion() (value uint16, err error) {
	retValue, err := instance.GetProperty("USBVersion")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetUSBVersionInBCD sets the value of USBVersionInBCD for the instance
func (instance *CIM_USBDevice) SetPropertyUSBVersionInBCD(value uint16) (err error) {
	return instance.SetProperty("USBVersionInBCD", (value))
}

// GetUSBVersionInBCD gets the value of USBVersionInBCD for the instance
func (instance *CIM_USBDevice) GetPropertyUSBVersionInBCD() (value uint16, err error) {
	retValue, err := instance.GetProperty("USBVersionInBCD")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetVendorID sets the value of VendorID for the instance
func (instance *CIM_USBDevice) SetPropertyVendorID(value uint16) (err error) {
	return instance.SetProperty("VendorID", (value))
}

// GetVendorID gets the value of VendorID for the instance
func (instance *CIM_USBDevice) GetPropertyVendorID() (value uint16, err error) {
	retValue, err := instance.GetProperty("VendorID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// This method returns the USBDevice Descriptor as specified by the input parameters. Each parameter is briefly described here with more detail in its Qualifier list. RequestType is an input parameter that defines whether the request is for standard, class or vendor-specific information, as well as specifying the recipient. RequestValue is also an input parameter and defines the USB Descriptor Type and Index. RequestIndex is an input parameter which describes the language used to return a string Descriptor. RequestLength is both an input and output parameter. It specifies the length of the Descriptor that should be returned (on input) and what is actually returned in the Buffer parameter (on output). Buffer is an output parameter, containing the Descriptor data. The GetDescriptor method returns an integer value of 0 if the USB Descriptor is successfully returned, 1 if the request is not supported and any other number to indicate an error.
///In a subclass, the set of possible return codes could be specified, using a ValueMap qualifier on the method. The strings to which the ValueMap contents are 'translated' may also be specified in the subclass as a Values array qualifier.

// <param name="RequestIndex" type="uint16 ">RequestIndex defines the 2 byte Language ID code used by the USBDevice when returning string Descriptor data. The parameter is typically 0 for non-string Descriptors. Refer to the USB Specification for more information.</param>
// <param name="RequestLength" type="uint16 ">On input, RequestLength is the length (in octets) of the Descriptor that should be returned. If this value is less than the actual length of the Descriptor, only the requested length will be returned. If it is more than the actual length, the actual length is returned. On output, this parameter is the length, in octets, of the Buffer being returned. If the requested Descriptor does not exist, the contents of this parameter are undefined.</param>
// <param name="RequestType" type="uint8 ">RequestType is bit-mapped and identifies the type of Descriptor request and the recipient. The type of request may be 'standard', 'class' or 'vendor-specific'. The recipient may be 'device', 'interface', 'endpoint' or 'other'. Refer to the USB Specification for the appropriate values for each bit.</param>
// <param name="RequestValue" type="uint16 ">RequestValue contains the Descriptor Type in the high byte and the Descriptor Index (for example, index or offset into the Descriptor array) in the low byte. Refer to the USB Specification for more information.</param>

// <param name="Buffer" type="uint8 []">Buffer returns the requested Descriptor information. If the Descriptor does not exist, the contents of the Buffer are undefined.</param>
// <param name="RequestLength" type="uint16 ">On input, RequestLength is the length (in octets) of the Descriptor that should be returned. If this value is less than the actual length of the Descriptor, only the requested length will be returned. If it is more than the actual length, the actual length is returned. On output, this parameter is the length, in octets, of the Buffer being returned. If the requested Descriptor does not exist, the contents of this parameter are undefined.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_USBDevice) GetDescriptor( /* IN */ RequestType uint8,
	/* IN */ RequestValue uint16,
	/* IN */ RequestIndex uint16,
	/* IN/OUT */ RequestLength uint16,
	/* OUT */ Buffer []uint8) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetDescriptor", RequestType, RequestValue, RequestIndex)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
