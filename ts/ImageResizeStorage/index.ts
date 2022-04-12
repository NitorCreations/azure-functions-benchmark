import { AzureFunction, Context, HttpRequest } from "@azure/functions";
import { BlobServiceClient } from "@azure/storage-blob";
import sharp = require("sharp");

const resize = (data: Buffer): Promise<Buffer> => {
  return sharp(data)
    .resize(300, 200, { kernel: "nearest" })
    .jpeg({
      quality: 80,
    })
    .toBuffer();
};

const httpTrigger: AzureFunction = async (
  ctx: Context,
  req: HttpRequest
): Promise<any> => {
  const t0 = performance.now();

  const serviceClient = BlobServiceClient.fromConnectionString(
    process.env.StorageConnectionString
  );
  const containerClient = serviceClient.getContainerClient("images");
  const srcBlob = containerClient.getBlockBlobClient("src.jpeg");
  const dstBlob = containerClient.getBlockBlobClient("dst.jpeg");

  const src = await srcBlob.downloadToBuffer();
  const dst = await resize(src);
  await dstBlob.uploadData(dst);

  const t1 = performance.now();
  const d = t1 - t0;

  return {
    body: {
      duration: d,
    },
  };
};

export default httpTrigger;
