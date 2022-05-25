using System;
using System.Collections.Generic;
using System.Text;
using System.Threading.Tasks;

namespace image_processing_azure_functions.Services
{
  public interface IProcessService
  {
      Task<string> UploadFile(string containerName, string fileName, byte[] file, string connectionString);
    
  }
}
