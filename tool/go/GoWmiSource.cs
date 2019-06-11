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

            AddReference("System");
            AddReference("System.Text");
            AddReference("System.Collections.Generic");
            AddReference("System.ComponentModel");
            Class = new GOWmiClass(this);
        }

        public GOWmiSource(ManagementClass wmiClass, WmiModule wModule) : base (wmiClass, wModule)
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
                    IFormat.GetDescriptionText(wmiClass.Qualifiers),
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
            Class = new GOWmiClass(wmiClass, this);
        }
        

        public override string GetSourceCode() 
        {
            StringBuilder sb = new StringBuilder();

            sb.AppendLine(CopyrightText);
            sb.AppendLine(HeaderComment);
            sb.AppendFormat(CultureInfo.InvariantCulture,
                "package {0}\n", Namespace);

            // Package imports
            sb.AppendLine("import (");
            foreach (var item in References)
            {
                sb.AppendLine(item.GetSourceCode());
            }
            sb.AppendLine(")");

            // Types
            sb.AppendLine(Class.GetSourceCode().Replace("\n", "\n\t"));

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

        protected override WmiSource GetWmiSource(string sourceName, WmiModule wModule)
        {
            return new GOWmiSource(sourceName, wModule);
        }

    }
}
