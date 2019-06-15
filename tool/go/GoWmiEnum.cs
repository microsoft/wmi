// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System.Globalization;
using System.IO;
using System.Text;
using Microsoft.WmiCodeGen.Common;

namespace Microsoft.WmiCodeGen.GO
{
    public class GOWmiEnum : WmiEnum
    {
        public GOWmiEnum(string enumName, GOWmiSource wSource) : base(enumName, wSource)
        {
            CopyrightText = string.Format(CultureInfo.InvariantCulture,
@"// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on {1} using github.com/microsoft/wmicodegen/tool
//      Source {0}
//////////////////////////////////////////////",
                    Name,
                    System.DateTime.Now.ToShortDateString()
                    );
        }

        public override string GetSourceCode()
        {
            StringBuilder sbEnum = new StringBuilder();
            sbEnum.AppendFormat(CultureInfo.InvariantCulture, "\n// {0} \n", Name);
            sbEnum.AppendFormat(CultureInfo.InvariantCulture, "type {0} int\n", Name);
            sbEnum.AppendLine("const(");
            foreach (var item in EnumValues)
            {
                sbEnum.AppendFormat(CultureInfo.InvariantCulture, "\t// {0} enum\n",
                    item.Key);
                sbEnum.AppendFormat(CultureInfo.InvariantCulture, "\t{0} {2} = {1}\n",
                    item.Key, item.Value, Name);
            }
            sbEnum.AppendLine(")");
            return sbEnum.ToString();
        }

        public override void GenerateSource(string outdir)
        {
            string path = Path.Combine(outdir, Name + ".go");
            StringBuilder sbEnum = new StringBuilder();
            sbEnum.AppendLine(CopyrightText);
            sbEnum.AppendFormat(CultureInfo.InvariantCulture,
                   "package {0}\n", Parent.Parent.Name);
            sbEnum.AppendLine(GetSourceCode());
            File.WriteAllText(path, sbEnum.ToString());
            Logger.Info("Enum {0}", path);
        }

    }
}
