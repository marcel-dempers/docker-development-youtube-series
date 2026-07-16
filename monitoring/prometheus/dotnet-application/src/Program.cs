using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Http;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using System;
using System.Text.Json;
using System.Threading.Tasks;
using System.Diagnostics;
using Prometheus;

var builder = WebApplication.CreateBuilder(args);
builder.WebHost.ConfigureKestrel(options =>
{
    options.ListenAnyIP(5000); // Bind to port 5000
});
builder.Services.AddControllers();

var app = builder.Build();

var ProcessedJobCount = Metrics
	.CreateCounter("dotnet_request_operations_total", "The total number of processed requests");

app.MapGet("/", async (HttpContext context) =>
{
  var sw = Stopwatch.StartNew();
  sw.Stop();

  ProcessedJobCount.Inc();
  var histogram =
  Metrics.CreateHistogram(
    "dotnet_request_duration_seconds",
    "Histogram for the duration in seconds.",
    new HistogramConfiguration
    {
        Buckets = Histogram.LinearBuckets(start: 1, width: 1, count: 5)
    });

  histogram.Observe(sw.Elapsed.TotalSeconds);

  var response = "hello world!";
  await context.Response.WriteAsync(response);
});

app.UseMetricServer();
app.Run();