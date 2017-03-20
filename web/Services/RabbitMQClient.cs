using System.Text;
using RabbitMQ.Client;
using WebApplication.Models;
using System;
using System.IO;

namespace WebApplication.Services
{
    public class RabbitMQClient : IDisposable
    {
        private readonly IConnection conn;
        private readonly string queueName;

        public RabbitMQClient(string url, string queueName)
        {
            ConnectionFactory factory = new ConnectionFactory();
            factory.AutomaticRecoveryEnabled = true;
            factory.Uri = url;
            this.conn = factory.CreateConnection();
            this.queueName = queueName;
        }

        public void SendMessage(Message message)
        {
            using (var channel = this.conn.CreateModel()) 
            {
                var messageString = Newtonsoft.Json.JsonConvert.SerializeObject(message);
                var messageBodyBytes = System.Text.Encoding.UTF8.GetBytes( messageString );

                channel.QueueDeclare(this.queueName, true, false, false, null); 
                var basicProperties = channel.CreateBasicProperties(); 
                basicProperties.DeliveryMode = 2; 
                
                channel.BasicPublish("", this.queueName, basicProperties, messageBodyBytes);
            }
        }

        #region IDisposable Support
        private bool disposedValue = false; // To detect redundant calls

        protected virtual void Dispose(bool disposing)
        {
            if (!disposedValue)
            {
                if (disposing)
                {
                    // TODO: dispose managed state (managed objects).
                    conn.Close();
                    conn.Dispose();
                }

                // TODO: free unmanaged resources (unmanaged objects) and override a finalizer below.
                // TODO: set large fields to null.

                disposedValue = true;
            }
        }

        // TODO: override a finalizer only if Dispose(bool disposing) above has code to free unmanaged resources.
        // ~RabbitMQClient() {
        //   // Do not change this code. Put cleanup code in Dispose(bool disposing) above.
        //   Dispose(false);
        // }

        // This code added to correctly implement the disposable pattern.
        public void Dispose()
        {
            // Do not change this code. Put cleanup code in Dispose(bool disposing) above.
            Dispose(true);
            // TODO: uncomment the following line if the finalizer is overridden above.
            // GC.SuppressFinalize(this);
        }
        #endregion
    }
}