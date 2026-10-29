// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Globalization;
using System.Management;
using System.Text;
using Microsoft.wmi.Common;

namespace Microsoft.wmi.GO
{
    public class GOWmiProperty : WmiProperty
    {
        public GOWmiProperty(PropertyData pData, GOWmiClass wClass) : base(pData, wClass)
        {
            ReferencedType = (Reference != null ? Reference.Alias + "." : "") + Type;
            IsEnum = (SType == typeof (Enum));
            HeaderComment = GOWmiSource.GetHeaderCommentText(pData.Qualifiers);
            Setter = GetPropertySetter(pData);
            Getter = GetPropertyGetter(pData);
            //Name = pData.Name;
            Name = FixGOPropertyName(pData.Name);
            //FixPropertyName(pData.Name, wmiClass);
            
        }

        public string ReferencedType { get; set; }
        public bool IsEnum { get; set; }

        public override string GetSourceCode()
        {
            StringBuilder sb = new StringBuilder();
            sb.AppendLine(HeaderComment);
            sb.AppendFormat(CultureInfo.InvariantCulture,
                "{1} {2}{0}", 
                ReferencedType, 
                FixPropertyName(GOWmiMethod.FixName(Name)), IsArray ? "[]" : ""
                );
            return sb.ToString();
        }

        public override string ToString()
        {
            return string.Format(CultureInfo.InvariantCulture, 
                "WmiProperty [Name:{0}] [Type:{1}]", Name, ReferencedType);
        }

        private string FixGOPropertyName(string name) {
            if (name == "type") return "_type";
            return name;
        }
        private string GetPropertySetter(PropertyData pData)
        {
            return String.Format(CultureInfo.InvariantCulture,
@"
// Set{1} sets the value of {1} for the instance
func (instance *{3}) SetProperty{1}(value {2}{0}) (err error) {{ 
    return instance.SetProperty(""{1}"", {4}(value))
}}", ReferencedType, pData.Name, IsArray ? "[]" : "", Parent.Name, IsEnum ? "" : "" );
        }

        private string GetPropertyGetter(PropertyData pData)
        {
            string nonarrayConversion = String.Format(CultureInfo.InvariantCulture, 
@"
    valuetmp, ok := retValue.({0}); 
    if !ok {{
        err = errors.Wrapf(errors.InvalidType, "" {0} is Invalid. Expected %s"", reflect.TypeOf(retValue))
        return  
    }}

    value = {1}(valuetmp)
", IsEnum ? "int32" : ReferencedType, ReferencedType);
            string arrayConversion = String.Format(CultureInfo.InvariantCulture, 
@"
    for _, interfaceValue := range retValue.([]interface{{}}) {{
        valuetmp, ok := interfaceValue.({0}); 
        if !ok {{
            err = errors.Wrapf(errors.InvalidType, "" {0} is Invalid. Expected %s"", reflect.TypeOf(interfaceValue))
            return  
        }}
        value = append(value, {1}(valuetmp))
    }}
", IsEnum ? "int32" : ReferencedType, ReferencedType);

            return String.Format(CultureInfo.InvariantCulture,
@"
// Get{1} gets the value of {1} for the instance
func (instance *{3}) GetProperty{1}()(value {2}{0}, err error) {{ 
    retValue, err := instance.GetProperty(""{1}"")
    if err != nil {{
        return
    }}
    if retValue == nil {{
        // Doesn't have any value. Return empty
        return
    }}
    {4}
    return
}}",
            ReferencedType, 
            pData.Name, 
            IsArray ? "[]" : "", 
            Parent.Name, 
            IsArray ? arrayConversion: nonarrayConversion);
        }
    }
}
