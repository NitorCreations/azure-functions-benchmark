import { AzureFunction, Context, HttpRequest } from "@azure/functions";

const estimatePi = (n: number): number => {
  let total = 0;
  let totalIn = 0;
  for (let i = 0; i < n; i++) {
    const x = Math.random();
    const y = Math.random();
    if (x * x + y * y < 1) {
      totalIn++;
    }
    total++;
  }
  return (totalIn * 4.0) / total;
};

const httpTrigger: AzureFunction = async function (
  ctx: Context,
  req: HttpRequest
): Promise<any> {
  const t0 = performance.now();
  const n = parseInt(ctx.req.query["n"] ?? "100000");

  const pi = estimatePi(n);

  const t1 = performance.now();
  const d = t1 - t0;

  return {
    body: {
      duration: d,
      result: pi,
    },
  };
};

export default httpTrigger;
