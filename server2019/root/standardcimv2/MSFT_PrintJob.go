// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_PrintJob struct
type MSFT_PrintJob struct {
	CIM_ManagedSystemElement

	//
	ComputerName string

	//
	Datatype string

	//
	DocumentName string

	//
	Id uint32

	//
	JobStatus uint32

	//
	JobTime uint32

	//
	PagesPrinted uint32

	//
	Position uint32

	//
	PrinterName string

	//
	Priority uint32

	//
	Size uint32

	//
	SubmittedTime string

	//
	TotalPages uint32

	//
	UserName string
}

// SetComputerName sets the value of ComputerName for the instance
func (instance *MSFT_PrintJob) SetPropertyComputerName(value string) (err error) {
	return instance.SetProperty("ComputerName", value)
}

// GetComputerName gets the value of ComputerName for the instance
func (instance *MSFT_PrintJob) GetPropertyComputerName() (value string, err error) {
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
func (instance *MSFT_PrintJob) SetPropertyDatatype(value string) (err error) {
	return instance.SetProperty("Datatype", value)
}

// GetDatatype gets the value of Datatype for the instance
func (instance *MSFT_PrintJob) GetPropertyDatatype() (value string, err error) {
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

// SetDocumentName sets the value of DocumentName for the instance
func (instance *MSFT_PrintJob) SetPropertyDocumentName(value string) (err error) {
	return instance.SetProperty("DocumentName", value)
}

// GetDocumentName gets the value of DocumentName for the instance
func (instance *MSFT_PrintJob) GetPropertyDocumentName() (value string, err error) {
	retValue, err := instance.GetProperty("DocumentName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetId sets the value of Id for the instance
func (instance *MSFT_PrintJob) SetPropertyId(value uint32) (err error) {
	return instance.SetProperty("Id", value)
}

// GetId gets the value of Id for the instance
func (instance *MSFT_PrintJob) GetPropertyId() (value uint32, err error) {
	retValue, err := instance.GetProperty("Id")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetJobStatus sets the value of JobStatus for the instance
func (instance *MSFT_PrintJob) SetPropertyJobStatus(value uint32) (err error) {
	return instance.SetProperty("JobStatus", value)
}

// GetJobStatus gets the value of JobStatus for the instance
func (instance *MSFT_PrintJob) GetPropertyJobStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("JobStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetJobTime sets the value of JobTime for the instance
func (instance *MSFT_PrintJob) SetPropertyJobTime(value uint32) (err error) {
	return instance.SetProperty("JobTime", value)
}

// GetJobTime gets the value of JobTime for the instance
func (instance *MSFT_PrintJob) GetPropertyJobTime() (value uint32, err error) {
	retValue, err := instance.GetProperty("JobTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPagesPrinted sets the value of PagesPrinted for the instance
func (instance *MSFT_PrintJob) SetPropertyPagesPrinted(value uint32) (err error) {
	return instance.SetProperty("PagesPrinted", value)
}

// GetPagesPrinted gets the value of PagesPrinted for the instance
func (instance *MSFT_PrintJob) GetPropertyPagesPrinted() (value uint32, err error) {
	retValue, err := instance.GetProperty("PagesPrinted")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPosition sets the value of Position for the instance
func (instance *MSFT_PrintJob) SetPropertyPosition(value uint32) (err error) {
	return instance.SetProperty("Position", value)
}

// GetPosition gets the value of Position for the instance
func (instance *MSFT_PrintJob) GetPropertyPosition() (value uint32, err error) {
	retValue, err := instance.GetProperty("Position")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPrinterName sets the value of PrinterName for the instance
func (instance *MSFT_PrintJob) SetPropertyPrinterName(value string) (err error) {
	return instance.SetProperty("PrinterName", value)
}

// GetPrinterName gets the value of PrinterName for the instance
func (instance *MSFT_PrintJob) GetPropertyPrinterName() (value string, err error) {
	retValue, err := instance.GetProperty("PrinterName")
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
func (instance *MSFT_PrintJob) SetPropertyPriority(value uint32) (err error) {
	return instance.SetProperty("Priority", value)
}

// GetPriority gets the value of Priority for the instance
func (instance *MSFT_PrintJob) GetPropertyPriority() (value uint32, err error) {
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

// SetSize sets the value of Size for the instance
func (instance *MSFT_PrintJob) SetPropertySize(value uint32) (err error) {
	return instance.SetProperty("Size", value)
}

// GetSize gets the value of Size for the instance
func (instance *MSFT_PrintJob) GetPropertySize() (value uint32, err error) {
	retValue, err := instance.GetProperty("Size")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSubmittedTime sets the value of SubmittedTime for the instance
func (instance *MSFT_PrintJob) SetPropertySubmittedTime(value string) (err error) {
	return instance.SetProperty("SubmittedTime", value)
}

// GetSubmittedTime gets the value of SubmittedTime for the instance
func (instance *MSFT_PrintJob) GetPropertySubmittedTime() (value string, err error) {
	retValue, err := instance.GetProperty("SubmittedTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalPages sets the value of TotalPages for the instance
func (instance *MSFT_PrintJob) SetPropertyTotalPages(value uint32) (err error) {
	return instance.SetProperty("TotalPages", value)
}

// GetTotalPages gets the value of TotalPages for the instance
func (instance *MSFT_PrintJob) GetPropertyTotalPages() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalPages")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUserName sets the value of UserName for the instance
func (instance *MSFT_PrintJob) SetPropertyUserName(value string) (err error) {
	return instance.SetProperty("UserName", value)
}

// GetUserName gets the value of UserName for the instance
func (instance *MSFT_PrintJob) GetPropertyUserName() (value string, err error) {
	retValue, err := instance.GetProperty("UserName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="ComputerName" type="string "></param>
// <param name="ID" type="uint32 "></param>
// <param name="PrinterName" type="string "></param>

// <param name="cmdletOutput" type="MSFT_PrintJob []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PrintJob) GetByName( /* IN */ ComputerName string,
	/* IN */ ID uint32,
	/* IN */ PrinterName string,
	/* OUT */ cmdletOutput []MSFT_PrintJob) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetByName", ComputerName, ID, PrinterName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ID" type="uint32 "></param>
// <param name="PrinterObject" type="MSFT_Printer "></param>

// <param name="cmdletOutput" type="MSFT_PrintJob []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PrintJob) GetByObject( /* IN */ ID uint32,
	/* IN */ PrinterObject MSFT_Printer,
	/* OUT */ cmdletOutput []MSFT_PrintJob) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetByObject", ID, PrinterObject)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="InputObject" type="MSFT_PrintJob "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PrintJob) DeleteJobByObject( /* IN */ InputObject MSFT_PrintJob) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("DeleteJobByObject", InputObject)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ComputerName" type="string "></param>
// <param name="ID" type="uint32 "></param>
// <param name="PrinterName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PrintJob) DeleteJobById( /* IN */ ComputerName string,
	/* IN */ ID uint32,
	/* IN */ PrinterName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("DeleteJobById", ComputerName, ID, PrinterName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ID" type="uint32 "></param>
// <param name="PrinterObject" type="MSFT_Printer "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PrintJob) DeleteJobByPrinterObject( /* IN */ ID uint32,
	/* IN */ PrinterObject MSFT_Printer) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("DeleteJobByPrinterObject", ID, PrinterObject)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="InputObject" type="MSFT_PrintJob "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PrintJob) RestartJobByObject( /* IN */ InputObject MSFT_PrintJob) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RestartJobByObject", InputObject)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ComputerName" type="string "></param>
// <param name="ID" type="uint32 "></param>
// <param name="PrinterName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PrintJob) RestartJobById( /* IN */ ComputerName string,
	/* IN */ ID uint32,
	/* IN */ PrinterName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RestartJobById", ComputerName, ID, PrinterName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ID" type="uint32 "></param>
// <param name="PrinterObject" type="MSFT_Printer "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PrintJob) RestartJobByPrinterObject( /* IN */ ID uint32,
	/* IN */ PrinterObject MSFT_Printer) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RestartJobByPrinterObject", ID, PrinterObject)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="InputObject" type="MSFT_PrintJob "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PrintJob) ResumeJobByObject( /* IN */ InputObject MSFT_PrintJob) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ResumeJobByObject", InputObject)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ComputerName" type="string "></param>
// <param name="ID" type="uint32 "></param>
// <param name="PrinterName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PrintJob) ResumeJobById( /* IN */ ComputerName string,
	/* IN */ ID uint32,
	/* IN */ PrinterName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ResumeJobById", ComputerName, ID, PrinterName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ID" type="uint32 "></param>
// <param name="PrinterObject" type="MSFT_Printer "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PrintJob) ResumeJobByPrinterObject( /* IN */ ID uint32,
	/* IN */ PrinterObject MSFT_Printer) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ResumeJobByPrinterObject", ID, PrinterObject)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="InputObject" type="MSFT_PrintJob "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PrintJob) SuspendJobByObject( /* IN */ InputObject MSFT_PrintJob) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SuspendJobByObject", InputObject)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ComputerName" type="string "></param>
// <param name="ID" type="uint32 "></param>
// <param name="PrinterName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PrintJob) SuspendJobById( /* IN */ ComputerName string,
	/* IN */ ID uint32,
	/* IN */ PrinterName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SuspendJobById", ComputerName, ID, PrinterName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ID" type="uint32 "></param>
// <param name="PrinterObject" type="MSFT_Printer "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_PrintJob) SuspendJobByPrinterObject( /* IN */ ID uint32,
	/* IN */ PrinterObject MSFT_Printer) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SuspendJobByPrinterObject", ID, PrinterObject)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
