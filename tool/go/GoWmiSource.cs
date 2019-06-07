// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;
using System.Globalization;
using System.IO;
using System.Management;
using System.Text;

namespace Microsoft.WmiCodeGen.GO
{
    public class GOWmiSource : WmiSource
    {
        public GOWmiSource(string sourceName, WmiModule wModule) : base(sourceName, wModule)
        {
            HeaderComment = string.Format(CultureInfo.InvariantCulture,
@"/*++
 * Copyright (c) Microsoft Corporation
 * 
 * Module Name:
 *  {0}.GO
 *  This is a public class. 
 * 
 * Abstract:
 *  {1}
 *  
 * Author:
 *      Auto Generated on {2} by MadhanM@Microsoft.com using WmiClassGen
 *      http://toolbox/WmiClassGen
 *      Source {3}
 * *********************************************/",
                                                Name,
                                                "This class doesn't exist. This is just a placeholder",
                                                DateTime.Now.ToShortDateString(),
                                                Parent.Parent.GONamespaceName
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
@"/*++
 * Copyright (c) Microsoft Corporation
 * 
 * Module Name:
 *  {0}.GO
 *  This is a public class. 
 * 
 * Abstract:
 *  {1}
 *  
 * Author:
 *      Auto Generated on {2} by MadhanM@Microsoft.com using WmiClassGen
 *      http://toolbox/WmiClassGen
 *      Source {3}
 * *********************************************/",
                                                Name,
                                                GOharpFormat.GetDescriptionText(wmiClass.Qualifiers),
                                                DateTime.Now.ToShortDateString(),
                                                Parent.Parent.GONamespaceName
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
            return new GOWmiReference(reference);
        }

        public void GenerateSources(string outdir)
        {
            string path = Path.Combine(outdir, Name + ".GO");

            File.WriteAllText(path, GetSourceCode());
            Logger.Info("Source {0}", path);
        }

        public void GenerateMakefile(string outputDir)
        {
            throw new NotImplementedException();
        }

        public static string GetHeaderCommentText(QualifierDataCollection qCollection)
        {

            object description = GOharpFormat.GetQualifierValue(qCollection, "description");
            return String.Format(CultureInfo.InvariantCulture,
@"
/// <summary>
/// {0}
/// </summary>",
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
