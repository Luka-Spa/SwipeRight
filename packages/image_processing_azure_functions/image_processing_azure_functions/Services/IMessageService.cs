using System;
using System.Collections.Generic;
using System.Text;
using System.Threading.Tasks;

namespace image_processing_azure_functions.Services
{
  public interface IMessageService
  {
      Task SendMessageAsync<T>(string serviceBusMessage, string queueName, string connectionString);
  }
}
