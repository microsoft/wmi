// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Globalization;
using System.IO;
using System.Management;
using System.Text;
using Microsoft.WmiCodeGen.Common;

namespace Microsoft.WmiCodeGen.GO
{
    public class GOWmiSource : WmiSource
    {
        public GOWmiSource(string sourceName, WmiModule wModule) : base(sourceName, wModule)
        {
            HeaderComment = string.Format(CultureInfo.InvariantCulture,
@"
// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.
//
// Author:
//      Auto Generated on {2} using github.com/microsoft/wmicodegen/tool
//      Source {3}
//////////////////////////////////////////////",
                    Name,
                    "This class doesn't exist. This is just a placeholder",
                    DateTime.Now.ToShortDateString(),
                    Parent.Parent.CSNamespaceName
                    );
            AddReference("github.com/microsoft/wmicodegen/go/cim");
            AddReference("github.com/microsoft/wmicodegen/go/wmi");
            Class = new GOWmiClass(this);
        }

        public GOWmiSource(ManagementClass wmiClass, WmiModule wModule) : base (wmiClass, wModule)
        {
            CopyrightText = string.Format(CultureInfo.InvariantCulture,
@"// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.
");

            ///
            HeaderComment = string.Format(CultureInfo.InvariantCulture,
@"// 
// Author:
//      Auto Generated on {2} using github.com/microsoft/wmicodegen/tool
//      Source {3}
//////////////////////////////////////////////",
                    Name,
                    IFormat.GetDescriptionText(wmiClass.Qualifiers),
                    DateTime.Now.ToShortDateString(),
                    Parent.Parent.CSNamespaceName
                    );

            AddReference("github.com/microsoft/wmicodegen/go/cim");
            AddReference("github.com/microsoft/wmicodegen/go/wmi");
            Class = new GOWmiClass(wmiClass, this);
        }
        

        public override string GetSourceCode() 
        {
            StringBuilder sb = new StringBuilder();

            sb.AppendLine(CopyrightText);
            sb.AppendLine(HeaderComment);
            sb.AppendFormat(CultureInfo.InvariantCulture,
                "package {0}\n", Parent.Name);

            // Package imports
            sb.AppendLine("import (");
            foreach (var item in References)
            {
                sb.AppendLine(item.GetSourceCode());
            }
            sb.AppendLine(")");

            // Types
            sb.AppendLine(Class.GetSourceCode());

            return sb.ToString();
        }
        protected override WmiReference GetWmiReference(string reference)
        {
            return new GOWmiReference(reference);
        }

        public override void GenerateSource(string outdir)
        {
            string path = Path.Combine(outdir, Name + ".go");

            File.WriteAllText(path, GetSourceCode());
            Logger.Info("Source {0}", path);
        }

        public void GenerateMakefile(string outputDir)
        {
            throw new NotImplementedException();
        }

        public static string GetHeaderCommentText(QualifierDataCollection qCollection)
        {
            object description = IFormat.GetQualifierValue(qCollection, "description");
            return String.Format(CultureInfo.InvariantCulture,
@"
// {0}",
               description != null ? description.ToString().Replace("\n", "\n///") : string.Empty);
        }

        protected override WmiEnum GetWmiEnum(string enumName, WmiSource wSource)
        {
            return new GOWmiEnum(enumName, wSource as GOWmiSource);
        }

        protected WmiSource GetWmiSource(string sourceName, WmiModule wModule)
        {
            return new GOWmiSource(sourceName, wModule);
        }
        //
        // Get the go type of the property
        //
        public override string GetType(PropertyData pData, out Type type)
        {
            type = ConvertCimTypeToSystemType(pData.Type);
            string typeName = type.Name;
            string moduleName = Parent.Parent.Name;
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
                                        Parent.Parent.AddSource(GetWmiSource(typeName, Parent.Parent.AddModule(moduleName)));
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
            else
            {
                typeName = typeName.ToLowerInvariant();
            }
            return FixTypeName(typeName);
            //return pData.IsArray ? typeName + "[]" : typeName;
        }

        private string FixTypeName(string name)
        {
            if (name.Equals("unint32", StringComparison.OrdinalIgnoreCase)) return "uint32";
            return name;
        }
    }
}
