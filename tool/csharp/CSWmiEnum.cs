// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System.Globalization;
using System.IO;
using System.Text;

namespace Microsoft.WmiCodeGen.CSharp
{
    public class CSWmiEnum : WmiEnum
    {
        public CSWmiEnum(string enumName, CSWmiSource wSource) : base(enumName, wSource)
        {
        }

        public override string GetSourceCode()
        {
            StringBuilder sbEnum = new StringBuilder();
            sbEnum.AppendFormat(CultureInfo.InvariantCulture, "\npublic enum {0}", Name);
            sbEnum.AppendLine("{");
            foreach (var item in EnumValues)
            {
                sbEnum.AppendFormat(CultureInfo.InvariantCulture, "\t{0} = {1},\n",
                    item.Key, item.Value);
            }
            sbEnum.AppendLine("}");
            return sbEnum.ToString();
        }

        public override void GenerateSource(string outdir)
        {
            string path = Path.Combine(outdir, Name + ".cs");
            StringBuilder sbEnum = new StringBuilder();
            sbEnum.AppendFormat(CultureInfo.InvariantCulture,
                   "namespace Microsoft.Test.Wmi.{0}\n", Namespace);
            sbEnum.AppendLine("{");
            sbEnum.AppendLine(GetSourceCode().Replace("\n", "\n\t"));
            sbEnum.AppendLine("}");
            File.WriteAllText(path, sbEnum.ToString());
            Logger.Info("Enum {0}", path);
        }

    }
}
