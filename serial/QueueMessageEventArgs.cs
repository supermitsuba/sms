using System;

namespace Serial.Models
{
    public class QueueMessageEventArgs:EventArgs
    {
        public Message MyMessage {get;set;}
        public QueueMessageEventArgs(Message message)
        {
            MyMessage = message;
        }
    }
}