import { serve } from "https://deno.land/std@0.50.0/http/server.ts";

for await (const req of serve({ port: 5000 })) {
  req.respond({ body: "Hello World! - from Deno 1.0.0\n" });
}