# WMI Code Generator
Tool to automatically generate wmi wrapper code for DotNet and GoLang based on a wmi library

## DotNet
Library : github.com/microsoft/wmi/dotnet/wmilib
Pre-Generated wrappers : github.com/microsoft/wmi/dotnet/generated

## GoLang

Library used for GoLang : github.com/microsoft/wmi/go/cim
Pre-Generated wrappers : github.com/microsoft/wmi/go/generated


## Usage
    wmi.exe
         -generatesource
                Description : Generate Source
                Example     : /generatesource
 
         -ns
                Description : Namespace Name
                Example     : /ns root/virtualization/v2
                Example     : /ns root/virtualization/v2,root/cimv2
 
         -recurse [Optional]
                Description : Recurse the namespace
                Example     : /recurse
 
         -format <Format Type> [Optional]
                 Values     : CS/GO
                Description : Format of the generated Source file
                Example     : /format CS
 
         -wmiclass [Optional]
                Description : WMI Class Name
                Example     : /wmiclass Msvm_ComputerSystem
                Optional
 
         -outdir <output directory path> [Optional]
                Description : Output Directory Path
                Example     : /outdir test
                Optional. Default : output
 
         -help [Optional]
                Description : Display this text
                Example     : /help
                Optional