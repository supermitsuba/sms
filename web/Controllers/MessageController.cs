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
        public IActionResult CreateMessage([FromBody]Message message)
        {
            if(message == null)
            {
                return BadRequest("Must contain a body.");
            }
            
            messageQueue.SendMessage(message);
            return Ok("Message will print shortly.");
        }
    }
}
