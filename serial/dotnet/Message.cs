using System.ComponentModel.DataAnnotations;

namespace Serial.Models
{
    public class Message
    {
        [Range(1, 60, ErrorMessage="The number duration must be between 1 and 60")]
        public int Duration {get;set;}

        [StringLength(32, ErrorMessage="Text can only contain a maximum of 32 characters.")]
        public string Text {get;set;}
    }
}
