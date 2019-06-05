// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System.Collections.Generic;
using System.Globalization;
using System.Management;

namespace Microsoft.WmiCodeGen
{
    public abstract class WmiEnum : IWmiBase
    {
        public WmiSource Parent { get; private set; }
        public WmiEnum(string enumName, WmiSource wSource)
        {
            Name = enumName;
            Parent = wSource;
            Namespace = wSource.Parent.Parent.CSNamespaceName;
        }

        public string Namespace { get; private set; }

        public void Add(string name, int value)
        {
            string nName = name;
            int i = 1;
            while (EnumValues.ContainsKey(nName))
            {
                nName = name + (i++).ToString();
            }

            EnumValues[nName] = value;
        }
        
        public Dictionary<string, int> m_enum = new Dictionary<string, int>();

        public Dictionary<string, int> EnumValues { get { return m_enum; } }
        public abstract void GenerateSource(string outdir);
        /// <summary>
        /// Compare WmiEnum objects
        /// </summary>
        /// <param name="inEnum"></param>
        /// <returns></returns>
        public bool Equals(WmiEnum inEnum)
        {
            if (Name != inEnum.Name) return false;

            if (EnumValues.Count != inEnum.EnumValues.Count) return false;

            if (!EnumValues.Equals(inEnum.EnumValues)) return false;

            foreach (KeyValuePair<string, int> kvp in EnumValues)
            {
                if (!inEnum.EnumValues.ContainsKey(kvp.Key)) return false;
                if (inEnum.EnumValues[kvp.Key] != kvp.Value) return false;
            }

            return true;
        }

        #region Static Methods
        public static string GetEnumName(PropertyData pData, WmiSource wSource)
        {
            string enumName = string.Format(CultureInfo.InvariantCulture,
                    "{1}_{0}", pData.Name, wSource.GetSuffix());
            if (enumName.Equals("0DebugTrace_Flags"))
            {
                enumName = "DebugTrace_Flags";
            }
            return enumName;
        }

        
        public static string FixEnumName(string input)
        {
            string output = input.Replace(" ", "_").Replace("/", "_").Replace("\\", "_").
                        Replace("(", "_").Replace(")", "_").Replace("-", "_").
                        Replace("&", "_").Replace(".", "_").Replace("'", "_").
                        Replace(":", "_").Replace("+", "plus").Replace("#", "_").
                        Replace(",", "_").Replace("\"", "_").Replace("*", "x").
                        Replace("%", "_").Replace("{", "_").Replace("}", "_").
                        Replace("[", "_").Replace("]", "_").Replace("^", "_").
                        Replace("=", "_").Replace("|", "_").Replace("<", "_").
                        Replace(">", "_").Replace("!", "_");
            int value; bool bValue;
            // Enum cannot contain names with Key words or intgers
            if (int.TryParse(output, out value) || bool.TryParse(output, out bValue) || int.TryParse(output.Substring(0, 1), out value))
            {
                output = "_" + output;
            }

            List<string> dTypeName = new List<string>{"string", "bool", "int", "double"};
            if (dTypeName.Contains(output)) 
            {
                output = "_" + output;
            }

            List<string> keyWords = new List<string> { "interface", "class", "public", "namespace" };
            if (keyWords.Contains(output))
            {
                output = "_" + output;
            }

            return output;
        }
        #endregion
    }
}
