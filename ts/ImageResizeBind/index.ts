import { AzureFunction, Context, HttpRequest } from "@azure/functions";
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

  const src = Buffer.from(ctx.bindings.srcImage);
  const dst = await resize(src);
  ctx.bindings.dstImage = dst;

  const t1 = performance.now();
  const d = t1 - t0;

  return {
    body: {
      duration: d,
    },
  };
};

export default httpTrigger;
