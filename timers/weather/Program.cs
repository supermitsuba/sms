using System;
using System.Net.Http;
using System.Text;
using System.Threading.Tasks;
using Newtonsoft.Json;
using Newtonsoft.Json.Linq;
using Polly;
using Polly.Retry;

namespace ConsoleApplication
{
    public class Program
    {
        private static RetryPolicy policy = null;

        public static void Main(string[] args)
        {
            var message = RetryGetMessage(args[0], args[1]);
            message.Wait();
            
            dynamic results = JsonConvert.DeserializeObject(message.Result);
            double temperature = (results.main.temp * 9 / 5) - 459.67;
            string conditions = results.weather[0].main;
            

            var dateTimeMessage = string.Format("Date: {0:MM/dd/yyyy}Time: {0:hh:mm tt}", DateTime.Now);
            var task1 = RetryPostMessage(args[2], args[3], "{ \"duration\":\"10\", \"text\":\""+dateTimeMessage+"\" }");
            task1.Wait();

            var weatherMessage = string.Format("Now :    {0:0.0} F {1}", temperature, conditions);
            var task2 = RetryPostMessage(args[2], args[3], "{ \"duration\":\"10\", \"text\":\""+weatherMessage+"\" }");
            task2.Wait();
        }

        public static async Task<string> RetryGetMessage(string baseAddress, string relativeUrl)
        {
            return await Policy.Handle<Exception>()
                               .RetryAsync(3)
                               .ExecuteAsync<string>( 
                                   () => {
                                    return GetMessage(baseAddress, relativeUrl);
                                });
        }

        public static async Task<string> GetMessage(string baseAddress, string relativeUrl)
        {
            using(var client = new HttpClient())
            {
                client.BaseAddress = new Uri(baseAddress);
                var response = await client.GetAsync(relativeUrl);
                if(response.IsSuccessStatusCode)
                {
                    return await response.Content.ReadAsStringAsync();
                }
                else
                {
                    throw new HttpRequestException(response.StatusCode.ToString());
                }
            }
        }

        public static async Task<string> RetryPostMessage(string baseAddress, string relativeUrl, string message)
        {
            return await Policy.Handle<Exception>()
                               .RetryAsync(3)
                               .ExecuteAsync<string>( 
                                   () => {
                                    return PostMessage(baseAddress, relativeUrl, message);
                                });
        }

        public static async Task<string> PostMessage(string baseAddress, string relativeUrl, string message)
        {
            using(var client = new HttpClient())
            {
                client.BaseAddress = new Uri(baseAddress);
                var response = await client.PostAsync(relativeUrl, new StringContent(message, Encoding.UTF8, "application/json"));
                if(response.IsSuccessStatusCode)
                {
                    return await response.Content.ReadAsStringAsync();
                }
                else
                {
                    throw new HttpRequestException(response.StatusCode.ToString());
                }
            }
        }
    }
}
