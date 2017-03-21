using System;
using System.Threading;
using Serial.Services;

namespace ConsoleApplication
{
    public class Program
    {
        public static void Main(string[] args)
        {
            var url = "amqp://guest:guest@localhost";
            var queueName = "messages";

            var s = new RabbitMQReceiver(url, queueName);

            s.MessageReceived += (sender, e) => {
                Console.WriteLine(e.MyMessage.Text);
                Thread.Sleep(e.MyMessage.Duration*1000);
            };

            s.Start();
            while(true) { Thread.Sleep(100); }
        }
    }
}
