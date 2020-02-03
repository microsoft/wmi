// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using System;

namespace Microsoft.wmi.Common
{
    public class Logger
    {
        private static void Log(string message, params object[] arguments)
        {
            Console.WriteLine(message, arguments);
        }
        public static void Debug(string message, params object[] arguments)
        {
            Log(message, arguments);
        }
        public static void Error(string message, params object[] arguments)
        {
            Log(message, arguments);
        }
        public static void Exception(Exception e, string message, params object[] arguments)
        {
            Log(e.Message);
            Log(e.StackTrace);
            Log(message, arguments);
            throw e;
        }
        public static void Info(string message, params object[] arguments)
        {
            Log(message, arguments);
        }
        public static void Warning(string message, params object[] arguments)
        {
            Log(message, arguments);
        }
    }
}
