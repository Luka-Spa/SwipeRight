using Azure.Storage.Blobs;
using Azure.Storage.Blobs.Models;
using System;
using System.Collections.Generic;
using System.IO;
using System.Text;
using System.Threading.Tasks;

namespace image_processing_azure_functions.Services
{
  public class ImageProcessService : IProcessService
  {
    public async Task<string> UploadFile(string containerName, string fileName, byte[] file, string connectionString)
    {
      try
      {
        // Get a reference to a container named "sample-container" and then create it
        var container = new BlobContainerClient(connectionString, containerName);
        await container.CreateIfNotExistsAsync();
        await container.SetAccessPolicyAsync(PublicAccessType.Blob);

        // Get a reference to a blob named "sample-file" in a container named "sample-container"
        var blob = container.GetBlobClient(fileName);

        Stream str = new MemoryStream(file);
        // Upload local file
        await blob.UploadAsync(str);

        return blob.Uri.AbsoluteUri;
      }
      catch (Exception)
      {
        return string.Empty;
      }
    }
  }
}
