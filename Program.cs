using System;
using System.Collections.Generic;
using System.Collections.ObjectModel;

namespace powbot_launcher_v2
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("PowBot is starting...");

            HomeFolder.Create();
            string jreBinary = JRE.GetOrObtainJREBinary();
            string clientFile = Client.EnsureLatestClient();
            Shell.Execute(jreBinary, Client.GetDirectory(), true, new List<string> {"-jar", clientFile});
        }
    }
}
