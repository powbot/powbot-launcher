using System;
using System.IO;
using System.Net;
using System.Security.Cryptography;

namespace powbot_launcher_v2
{
    class Client
    {

        public static string GetDirectory()
        {
            return Path.Combine(HomeFolder.GetDirectory(), "client");
        }

        private static string GetRemoteHash() 
        {
            using (var client = new WebClient())
            {
                return client.DownloadString("https://powbot.org/game/current_client").Replace("\n", "");
            }
        }

        private static void ObtainClient(string hash, string outputFile)
        {
            Console.WriteLine($"Downloading latest client: {hash}");
            using (var client = new WebClient())
            {
                client.DownloadFile($"https://powbot.org/game/{hash}.jar", outputFile);
            }
            
            if (!File.Exists(outputFile))
            {
                Console.WriteLine("Failed to download client");
                Environment.Exit(255);
            }
        }

        private static string ComputeSha1Hash(string file)
        {
            using (FileStream fs = File.OpenRead(file))
            {
                SHA1 sha = new SHA1Managed();
                return BitConverter.ToString(sha.ComputeHash(fs)).Replace("-", "").ToLower();
            }
        }

        public static string EnsureLatestClient()
        {

            if (!System.IO.Directory.Exists(GetDirectory())) {
                System.IO.Directory.CreateDirectory(GetDirectory());
            }

            string expectedHash = GetRemoteHash();
            string clientFile = System.IO.Path.Combine(GetDirectory(), "PowBot.jar");
            if (!File.Exists(clientFile)) {
                Console.WriteLine("No client file found, downloading...");
                ObtainClient(expectedHash, clientFile);
            }

            string actualHash = ComputeSha1Hash(clientFile);
            if (actualHash != expectedHash) {
                ObtainClient(expectedHash, clientFile);
            }

            Console.WriteLine($"Using client: {ComputeSha1Hash(clientFile)}");

            return clientFile;
        }
    }
}
