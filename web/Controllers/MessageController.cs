using Microsoft.AspNetCore.Mvc;
using WebApplication.Filters;
using WebApplication.Models;
using WebApplication.Services;

namespace WebApplication.Controllers
{
    [ValidateModel]
    public class MessageController : Controller
    {
        private RabbitMQClient messageQueue;
        public MessageController(RabbitMQClient client)
        {
            this.messageQueue = client;
        }

        [HttpPost]
        [Route("api/message")]
        public IActionResult CreateMessage([FromBody]Message message, [FromQuery] bool priority)
        {
            if(message == null)
            {
                return BadRequest("Must contain a body.");
            }
            
            if(priority)
            {
                messageQueue.SendMessage(message, "priority");
            }
            else
            {
                messageQueue.SendMessage(message, "messages");
            }

            return Ok("Message will print shortly.");
        }

        [HttpGet]
        [Route("api/test")]
        public IActionResult GetTest()
        {
            return Ok("Working...");
        }
    }
}
