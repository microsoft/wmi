// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_Printer struct
type MSFT_Printer struct {
	*CIM_ManagedSystemElement

	//
	BranchOfficeOfflineLogSizeMB uint32

	//
	Comment string

	//
	ComputerName string

	//
	Datatype string

	//
	DefaultJobPriority uint32

	//
	DeviceType uint32

	//
	DisableBranchOfficeLogging bool

	//
	DriverName string

	//
	JobCount uint32

	//
	KeepPrintedJobs bool

	//
	Location string

	//
	PermissionSDDL string

	//
	PortName string

	//
	PrinterStatus uint32

	//
	PrintProcessor string

	//
	Priority uint32

	//
	Published bool

	//
	RenderingMode uint32

	//
	SeparatorPageFile string

	//
	Shared bool

	//
	ShareName string

	//
	StartTime uint32

	//
	Type uint32

	//
	UntilTime uint32

	//
	WorkflowPolicy uint32
}

func NewMSFT_PrinterEx1(instance *cim.WmiInstance) (newInstance *MSFT_Printer, err error) {
	tmp, err := NewCIM_ManagedSystemElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_Printer{
		CIM_ManagedSystemElement: tmp,
	}
	return
}

func NewMSFT_PrinterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_Printer, err error) {
	tmp, err := NewCIM_ManagedSystemElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_Printer{
		CIM_ManagedSystemElement: tmp,
	}
	return
}

// SetBranchOfficeOfflineLogSizeMB sets the value of BranchOfficeOfflineLogSizeMB for the instance
func (instance *MSFT_Printer) SetPropertyBranchOfficeOfflineLogSizeMB(value uint32) (err error) {
	return instance.SetProperty("BranchOfficeOfflineLogSizeMB", value)
}

// GetBranchOfficeOfflineLogSizeMB gets the value of BranchOfficeOfflineLogSizeMB for the instance
func (instance *MSFT_Printer) GetPropertyBranchOfficeOfflineLogSizeMB() (value uint32, err error) {
	retValue, err := instance.GetProperty("BranchOfficeOfflineLogSizeMB")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetComment sets the value of Comment for the instance
func (instance *MSFT_Printer) SetPropertyComment(value string) (err error) {
	return instance.SetProperty("Comment", value)
}

// GetComment gets the value of Comment for the instance
func (instance *MSFT_Printer) GetPropertyComment() (value string, err error) {
	retValue, err := instance.GetProperty("Comment")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetComputerName sets the value of ComputerName for the instance
func (instance *MSFT_Printer) SetPropertyComputerName(value string) (err error) {
	return instance.SetProperty("ComputerName", value)
}

// GetComputerName gets the value of ComputerName for the instance
func (instance *MSFT_Printer) GetPropertyComputerName() (value string, err error) {
	retValue, err := instance.GetProperty("ComputerName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDatatype sets the value of Datatype for the instance
func (instance *MSFT_Printer) SetPropertyDatatype(value string) (err error) {
	return instance.SetProperty("Datatype", value)
}

// GetDatatype gets the value of Datatype for the instance
func (instance *MSFT_Printer) GetPropertyDatatype() (value string, err error) {
	retValue, err := instance.GetProperty("Datatype")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDefaultJobPriority sets the value of DefaultJobPriority for the instance
func (instance *MSFT_Printer) SetPropertyDefaultJobPriority(value uint32) (err error) {
	return instance.SetProperty("DefaultJobPriority", value)
}

// GetDefaultJobPriority gets the value of DefaultJobPriority for the instance
func (instance *MSFT_Printer) GetPropertyDefaultJobPriority() (value uint32, err error) {
	retValue, err := instance.GetProperty("DefaultJobPriority")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDeviceType sets the value of DeviceType for the instance
func (instance *MSFT_Printer) SetPropertyDeviceType(value uint32) (err error) {
	return instance.SetProperty("DeviceType", value)
}

// GetDeviceType gets the value of DeviceType for the instance
func (instance *MSFT_Printer) GetPropertyDeviceType() (value uint32, err error) {
	retValue, err := instance.GetProperty("DeviceType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisableBranchOfficeLogging sets the value of DisableBranchOfficeLogging for the instance
func (instance *MSFT_Printer) SetPropertyDisableBranchOfficeLogging(value bool) (err error) {
	return instance.SetProperty("DisableBranchOfficeLogging", value)
}

// GetDisableBranchOfficeLogging gets the value of DisableBranchOfficeLogging for the instance
func (instance *MSFT_Printer) GetPropertyDisableBranchOfficeLogging() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableBranchOfficeLogging")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDriverName sets the value of DriverName for the instance
func (instance *MSFT_Printer) SetPropertyDriverName(value string) (err error) {
	return instance.SetProperty("DriverName", value)
}

// GetDriverName gets the value of DriverName for the instance
func (instance *MSFT_Printer) GetPropertyDriverName() (value string, err error) {
	retValue, err := instance.GetProperty("DriverName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetJobCount sets the value of JobCount for the instance
func (instance *MSFT_Printer) SetPropertyJobCount(value uint32) (err error) {
	return instance.SetProperty("JobCount", value)
}

// GetJobCount gets the value of JobCount for the instance
func (instance *MSFT_Printer) GetPropertyJobCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("JobCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetKeepPrintedJobs sets the value of KeepPrintedJobs for the instance
func (instance *MSFT_Printer) SetPropertyKeepPrintedJobs(value bool) (err error) {
	return instance.SetProperty("KeepPrintedJobs", value)
}

// GetKeepPrintedJobs gets the value of KeepPrintedJobs for the instance
func (instance *MSFT_Printer) GetPropertyKeepPrintedJobs() (value bool, err error) {
	retValue, err := instance.GetProperty("KeepPrintedJobs")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocation sets the value of Location for the instance
func (instance *MSFT_Printer) SetPropertyLocation(value string) (err error) {
	return instance.SetProperty("Location", value)
}

// GetLocation gets the value of Location for the instance
func (instance *MSFT_Printer) GetPropertyLocation() (value string, err error) {
	retValue, err := instance.GetProperty("Location")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPermissionSDDL sets the value of PermissionSDDL for the instance
func (instance *MSFT_Printer) SetPropertyPermissionSDDL(value string) (err error) {
	return instance.SetProperty("PermissionSDDL", value)
}

// GetPermissionSDDL gets the value of PermissionSDDL for the instance
func (instance *MSFT_Printer) GetPropertyPermissionSDDL() (value string, err error) {
	retValue, err := instance.GetProperty("PermissionSDDL")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPortName sets the value of PortName for the instance
func (instance *MSFT_Printer) SetPropertyPortName(value string) (err error) {
	return instance.SetProperty("PortName", value)
}

// GetPortName gets the value of PortName for the instance
func (instance *MSFT_Printer) GetPropertyPortName() (value string, err error) {
	retValue, err := instance.GetProperty("PortName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPrinterStatus sets the value of PrinterStatus for the instance
func (instance *MSFT_Printer) SetPropertyPrinterStatus(value uint32) (err error) {
	return instance.SetProperty("PrinterStatus", value)
}

// GetPrinterStatus gets the value of PrinterStatus for the instance
func (instance *MSFT_Printer) GetPropertyPrinterStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("PrinterStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPrintProcessor sets the value of PrintProcessor for the instance
func (instance *MSFT_Printer) SetPropertyPrintProcessor(value string) (err error) {
	return instance.SetProperty("PrintProcessor", value)
}

// GetPrintProcessor gets the value of PrintProcessor for the instance
func (instance *MSFT_Printer) GetPropertyPrintProcessor() (value string, err error) {
	retValue, err := instance.GetProperty("PrintProcessor")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPriority sets the value of Priority for the instance
func (instance *MSFT_Printer) SetPropertyPriority(value uint32) (err error) {
	return instance.SetProperty("Priority", value)
}

// GetPriority gets the value of Priority for the instance
func (instance *MSFT_Printer) GetPropertyPriority() (value uint32, err error) {
	retValue, err := instance.GetProperty("Priority")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPublished sets the value of Published for the instance
func (instance *MSFT_Printer) SetPropertyPublished(value bool) (err error) {
	return instance.SetProperty("Published", value)
}

// GetPublished gets the value of Published for the instance
func (instance *MSFT_Printer) GetPropertyPublished() (value bool, err error) {
	retValue, err := instance.GetProperty("Published")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRenderingMode sets the value of RenderingMode for the instance
func (instance *MSFT_Printer) SetPropertyRenderingMode(value uint32) (err error) {
	return instance.SetProperty("RenderingMode", value)
}

// GetRenderingMode gets the value of RenderingMode for the instance
func (instance *MSFT_Printer) GetPropertyRenderingMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("RenderingMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSeparatorPageFile sets the value of SeparatorPageFile for the instance
func (instance *MSFT_Printer) SetPropertySeparatorPageFile(value string) (err error) {
	return instance.SetProperty("SeparatorPageFile", value)
}

// GetSeparatorPageFile gets the value of SeparatorPageFile for the instance
func (instance *MSFT_Printer) GetPropertySeparatorPageFile() (value string, err error) {
	retValue, err := instance.GetProperty("SeparatorPageFile")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetShared sets the value of Shared for the instance
func (instance *MSFT_Printer) SetPropertyShared(value bool) (err error) {
	return instance.SetProperty("Shared", value)
}

// GetShared gets the value of Shared for the instance
func (instance *MSFT_Printer) GetPropertyShared() (value bool, err error) {
	retValue, err := instance.GetProperty("Shared")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetShareName sets the value of ShareName for the instance
func (instance *MSFT_Printer) SetPropertyShareName(value string) (err error) {
	return instance.SetProperty("ShareName", value)
}

// GetShareName gets the value of ShareName for the instance
func (instance *MSFT_Printer) GetPropertyShareName() (value string, err error) {
	retValue, err := instance.GetProperty("ShareName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStartTime sets the value of StartTime for the instance
func (instance *MSFT_Printer) SetPropertyStartTime(value uint32) (err error) {
	return instance.SetProperty("StartTime", value)
}

// GetStartTime gets the value of StartTime for the instance
func (instance *MSFT_Printer) GetPropertyStartTime() (value uint32, err error) {
	retValue, err := instance.GetProperty("StartTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetType sets the value of Type for the instance
func (instance *MSFT_Printer) SetPropertyType(value uint32) (err error) {
	return instance.SetProperty("Type", value)
}

// GetType gets the value of Type for the instance
func (instance *MSFT_Printer) GetPropertyType() (value uint32, err error) {
	retValue, err := instance.GetProperty("Type")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUntilTime sets the value of UntilTime for the instance
func (instance *MSFT_Printer) SetPropertyUntilTime(value uint32) (err error) {
	return instance.SetProperty("UntilTime", value)
}

// GetUntilTime gets the value of UntilTime for the instance
func (instance *MSFT_Printer) GetPropertyUntilTime() (value uint32, err error) {
	retValue, err := instance.GetProperty("UntilTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWorkflowPolicy sets the value of WorkflowPolicy for the instance
func (instance *MSFT_Printer) SetPropertyWorkflowPolicy(value uint32) (err error) {
	return instance.SetProperty("WorkflowPolicy", value)
}

// GetWorkflowPolicy gets the value of WorkflowPolicy for the instance
func (instance *MSFT_Printer) GetPropertyWorkflowPolicy() (value uint32, err error) {
	retValue, err := instance.GetProperty("WorkflowPolicy")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="ConnectionName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Printer) AddConnection( /* IN */ ConnectionName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("AddConnection", ConnectionName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="BranchOfficeOfflineLogSizeMB" type="uint32 "></param>
// <param name="Comment" type="string "></param>
// <param name="ComputerName" type="string "></param>
// <param name="Datatype" type="string "></param>
// <param name="DisableBranchOfficeLogging" type="bool "></param>
// <param name="DriverName" type="string "></param>
// <param name="KeepPrintedJobs" type="bool "></param>
// <param name="Location" type="string "></param>
// <param name="Name" type="string "></param>
// <param name="PermissionSDDL" type="string "></param>
// <param name="PortName" type="string "></param>
// <param name="PrintProcessor" type="string "></param>
// <param name="Priority" type="uint32 "></param>
// <param name="Published" type="bool "></param>
// <param name="RenderingMode" type="uint32 "></param>
// <param name="SeparatorPageFile" type="string "></param>
// <param name="Shared" type="bool "></param>
// <param name="ShareName" type="string "></param>
// <param name="StartTime" type="uint32 "></param>
// <param name="UntilTime" type="uint32 "></param>
// <param name="WorkflowPolicy" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Printer) AddByExistingPort( /* IN */ Comment string,
	/* IN */ Datatype string,
	/* IN */ DriverName string,
	/* IN */ UntilTime uint32,
	/* IN */ KeepPrintedJobs bool,
	/* IN */ Location string,
	/* IN */ Name string,
	/* IN */ PermissionSDDL string,
	/* IN */ PortName string,
	/* IN */ PrintProcessor string,
	/* IN */ Priority uint32,
	/* IN */ Published bool,
	/* IN */ RenderingMode uint32,
	/* IN */ SeparatorPageFile string,
	/* IN */ ComputerName string,
	/* IN */ ShareName string,
	/* IN */ Shared bool,
	/* IN */ StartTime uint32,
	/* IN */ DisableBranchOfficeLogging bool,
	/* IN */ BranchOfficeOfflineLogSizeMB uint32,
	/* IN */ WorkflowPolicy uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("AddByExistingPort", Comment, Datatype, DriverName, UntilTime, KeepPrintedJobs, Location, Name, PermissionSDDL, PortName, PrintProcessor, Priority, Published, RenderingMode, SeparatorPageFile, ComputerName, ShareName, Shared, StartTime, DisableBranchOfficeLogging, BranchOfficeOfflineLogSizeMB, WorkflowPolicy)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="BranchOfficeOfflineLogSizeMB" type="uint32 "></param>
// <param name="Comment" type="string "></param>
// <param name="ComputerName" type="string "></param>
// <param name="Datatype" type="string "></param>
// <param name="DeviceURL" type="string "></param>
// <param name="DeviceUUID" type="string "></param>
// <param name="DisableBranchOfficeLogging" type="bool "></param>
// <param name="KeepPrintedJobs" type="bool "></param>
// <param name="Location" type="string "></param>
// <param name="Name" type="string "></param>
// <param name="PermissionSDDL" type="string "></param>
// <param name="PrintProcessor" type="string "></param>
// <param name="Priority" type="uint32 "></param>
// <param name="Published" type="bool "></param>
// <param name="RenderingMode" type="uint32 "></param>
// <param name="SeparatorPageFile" type="string "></param>
// <param name="Shared" type="bool "></param>
// <param name="ShareName" type="string "></param>
// <param name="StartTime" type="uint32 "></param>
// <param name="UntilTime" type="uint32 "></param>
// <param name="WorkflowPolicy" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Printer) AddByWsdPort( /* IN */ Comment string,
	/* IN */ Datatype string,
	/* IN */ DeviceURL string,
	/* IN */ UntilTime uint32,
	/* IN */ DeviceUUID string,
	/* IN */ KeepPrintedJobs bool,
	/* IN */ Location string,
	/* IN */ Name string,
	/* IN */ PermissionSDDL string,
	/* IN */ PrintProcessor string,
	/* IN */ Priority uint32,
	/* IN */ Published bool,
	/* IN */ RenderingMode uint32,
	/* IN */ SeparatorPageFile string,
	/* IN */ ComputerName string,
	/* IN */ ShareName string,
	/* IN */ Shared bool,
	/* IN */ StartTime uint32,
	/* IN */ DisableBranchOfficeLogging bool,
	/* IN */ BranchOfficeOfflineLogSizeMB uint32,
	/* IN */ WorkflowPolicy uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("AddByWsdPort", Comment, Datatype, DeviceURL, UntilTime, DeviceUUID, KeepPrintedJobs, Location, Name, PermissionSDDL, PrintProcessor, Priority, Published, RenderingMode, SeparatorPageFile, ComputerName, ShareName, Shared, StartTime, DisableBranchOfficeLogging, BranchOfficeOfflineLogSizeMB, WorkflowPolicy)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ComputerName" type="string "></param>
// <param name="Name" type="string "></param>
// <param name="NewName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Printer) RenameByName( /* IN */ Name string,
	/* IN */ NewName string,
	/* IN */ ComputerName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RenameByName", Name, NewName, ComputerName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="InputObject" type="MSFT_Printer "></param>
// <param name="NewName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Printer) RenameByObject( /* IN */ InputObject MSFT_Printer,
	/* IN */ NewName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RenameByObject", InputObject, NewName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
