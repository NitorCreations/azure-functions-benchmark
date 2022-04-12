import { AzureFunction, Context, HttpRequest } from "@azure/functions";

const fib = (seq: number): number => {
  let x = 0;
  let y = 1;
  for (let i = 0; i < seq; i++) {
    const temp = x;
    x = x + y;
    y = temp;
  }
  return y;
};

const httpTrigger: AzureFunction = async function (
  ctx: Context,
  req: HttpRequest
): Promise<any> {
  const t0 = performance.now();
  const seq = parseInt(ctx.req.query["seq"] ?? "30");

  const f = fib(seq);

  const t1 = performance.now();
  const d = t1 - t0;

  return {
    body: {
      duration: d,
      result: f,
    },
  };
};

export default httpTrigger;
