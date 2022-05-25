using Confluent.Kafka;
using System;
using System.Collections.Generic;
using System.Text;
using System.Threading.Tasks;

namespace image_processing_azure_functions.Services
{
  class KafkaService : IMessageService
  {
    ProducerConfig config = new ProducerConfig
    {
      BootstrapServers = "<your-IP-port-pairs>",
      SslCaLocation = "/Path-to/cluster-ca-certificate.pem",
      SecurityProtocol = SecurityProtocol.SaslSsl,
      SaslMechanism = SaslMechanism.ScramSha256,
      SaslUsername = "ickafka",
      SaslPassword = "yourpassword",

    };
    public async Task SendMessageAsync<T>(string serviceBusMessage, string queueName, string connectionString)
    {
      using (var p = new ProducerBuilder<Null, string>(config).Build())
      {
        try
        {
          var dr = await p.ProduceAsync(queueName, new Message<Null, string> {Value = serviceBusMessage });
          Console.WriteLine($"Delivered '{dr.Value}' to '{dr.TopicPartitionOffset}'");
        }
        catch (ProduceException<Null, string> e)
        {
          Console.WriteLine($"Delivery failed: {e.Error.Reason}");
        }
      }
    }
  }
}
  }
}
