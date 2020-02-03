// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Globalization;
using System.IO;
using System.Management;
using System.Text;
using Microsoft.wmi.Common;

namespace Microsoft.wmi.CSharp
{
    public class CSWmiSource : WmiSource
    {
        public CSWmiSource(string sourceName, WmiModule wModule) : base(sourceName, wModule)
        {
            HeaderComment = string.Format(CultureInfo.InvariantCulture,
@"/*++
 *       Copyright 2019 (c) Microsoft Corporation.
 *       Licensed under the MIT license.
 * Module Name:
 *  {0}.cs
 *  This is a public class. 
 * 
 * Abstract:
 *  {1}
 *  
 * Author:
 *      Auto Generated on {2} by github.com/microsoft/wmi
 *      Source {3}
 * *********************************************/",
                                                Name,
                                                "This class doesn't exist. This is just a placeholder",
                                                DateTime.Now.ToShortDateString(),
                                                Parent.Parent.CSNamespaceName
                                                );

            AddReference("System");
            AddReference("System.Text");
            AddReference("System.Collections.Generic");
            AddReference("System.ComponentModel");
            Class = new CSWmiClass(this);
        }

        public CSWmiSource(ManagementClass wmiClass, WmiModule wModule) : base (wmiClass, wModule)
        {
            HeaderComment = string.Format(CultureInfo.InvariantCulture,
@"/*++
 *       Copyright 2019 (c) Microsoft Corporation.
 *       Licensed under the MIT license.
 * 
 * Module Name:
 *  {0}.cs
 *  This is a public class. 
 * 
 * Abstract:
 *  {1}
 *  
 * Author:
 *      Auto Generated on {2} by github.com/microsoft/wmi
 *      Source {3}
 * *********************************************/",
                                                Name,
                                                CSharpFormat.GetDescriptionText(wmiClass.Qualifiers),
                                                DateTime.Now.ToShortDateString(),
                                                Parent.Parent.CSNamespaceName
                                                );

            AddReference("System");
            AddReference("System.Text");
            AddReference("System.Collections.Generic");
            AddReference("System.ComponentModel");
            AddReference("System.Threading");
            AddReference("System.Linq");
            AddReference("System.Globalization");
            AddReference("System.Runtime.InteropServices");
            Class = new CSWmiClass(wmiClass, this);
        }
        

        public override string GetSourceCode() 
        {
            StringBuilder sb = new StringBuilder();

            sb.AppendLine(CopyrightText);
            sb.AppendLine(HeaderComment);
            foreach (var item in References)
            {
                sb.AppendLine(item.GetSourceCode());
            }
            sb.AppendFormat(CultureInfo.InvariantCulture,
                "namespace Microsoft.Test.Wmi.{0}\n", Namespace);
            sb.AppendLine("{");
            sb.AppendLine(Class.GetSourceCode().Replace("\n", "\n\t"));
            sb.AppendLine("}");

            return sb.ToString();
        }
        protected override WmiReference GetWmiReference(string reference)
        {
            return new CSWmiReference(reference);
        }

        public override void GenerateSource(string outdir)
        {
            string path = Path.Combine(outdir, Name + ".cs");

            File.WriteAllText(path, GetSourceCode());
            Logger.Info("Source {0}", path);
        }

        public void GenerateMakefile(string outputDir)
        {
            throw new NotImplementedException();
        }

        public static string GetHeaderCommentText(QualifierDataCollection qCollection)
        {

            object description = CSharpFormat.GetQualifierValue(qCollection, "description");
            return String.Format(CultureInfo.InvariantCulture,
@"
/// <summary>
/// {0}
/// </summary>",
               description != null ? description.ToString().Replace("\n", "\n///") : string.Empty);
        }

        protected override WmiEnum GetWmiEnum(string enumName, WmiSource wSource)
        {
            return new CSWmiEnum(enumName, wSource as CSWmiSource);
        }

        protected WmiSource GetWmiSource(string sourceName, WmiModule wModule)
        {
            return new CSWmiSource(sourceName, wModule);
        }


        public override string GetType(PropertyData pData, out Type type)
        {
            type = ConvertCimTypeToSystemType(pData.Type);
            string typeName = type.Name;
            if (pData.Type == CimType.Object || pData.Type == CimType.Reference)
            {
                object typeValue;
                if (IFormat.TryGetQualifierValue(pData.Qualifiers, "CimType", out typeValue))
                {
                    string typeValueString = typeValue.ToString();
                    if (typeValueString.Contains(":"))
                    {
                        string[] typeValues = typeValueString.Split(new char[] { ':' });
                        if (typeValues.Length > 1)
                        {
                            // Add reference to this type
                            typeName = WmiClass.FixClassName(typeValues[1]);

                            if (!typeName.Equals("unint32", StringComparison.OrdinalIgnoreCase))
                            {
                                if (!Parent.Parent.HasSource(typeName) &&
                                    !CheckClass(typeName, Parent.Parent.Name))
                                {
                                    // Class not found in the current Namespace. Start searching from root namespace
                                    Logger.Debug("Class not found in the current Namespace." +
                                    " Start searching from root namespace - {0}",
                                        Namespace);
                                    string reference = GetReference(typeName, "root");
                                    if (!string.IsNullOrEmpty(reference))
                                    {
                                        AddReference(reference);
                                        Parent.Parent.AddReference(reference);
                                    }
                                    else
                                    {
                                        // No reference found anywhere. Create a Dummy Source
                                        Parent.Parent.AddSource(GetWmiSource(typeName, Parent.Parent.AddModule(CSWmiSource.GetModuleName(typeName))));
                                    }
                                }
                            }
                            else
                            {
                                type = typeof(UInt32);
                            }
                        }
                    }
                }
            }
            else if (pData.Type != CimType.Boolean && IFormat.HasQualifier(pData.Qualifiers, "values"))
            {
                WmiEnum wEnum = GetEnum(pData);
                if (wEnum != null)
                {
                    typeName = wEnum.Name;
                }
                else typeName = type.Name;
            }
            return FixTypeName(typeName);
            //return pData.IsArray ? typeName + "[]" : typeName;
        }

        private string FixTypeName(string name)
        {
            if (name.Equals("unint32", StringComparison.OrdinalIgnoreCase)) return typeof(UInt32).Name;
            return name;
        }

        public static string GetModuleName(string className)
        {
            if (className.StartsWith("__")) return "System";
            string[] classPrefixes = className.Split(new char[] { '_' });

            if (classPrefixes.Length > 1)
            {
                return classPrefixes[0];
            }
            return "Common";
        }
    }
}
