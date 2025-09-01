using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Http;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using StackExchange.Redis;
using System;
using System.Text.Json;
using System.Threading.Tasks;

using System.Diagnostics;


var serviceName = "videos-api";
var builder = WebApplication.CreateBuilder(args);

builder.Services.AddControllers();

var app = builder.Build();

var environment = Environment.GetEnvironmentVariable("ENVIRONMENT");
var redisHost = Environment.GetEnvironmentVariable("REDIS_HOST");
var redisPort = Environment.GetEnvironmentVariable("REDIS_PORT");
var flaky = Environment.GetEnvironmentVariable("FLAKY");

var redis = ConnectionMultiplexer.Connect($"{redisHost}:{redisPort}");
var db = redis.GetDatabase();

var registeredActivity = new ActivitySource("ManualInstrumentations." + serviceName);

app.MapGet("/{id}", async (HttpContext context, string id) =>
{
    if (flaky == "true" && new Random().Next(90) < 30)
    {
        throw new Exception("flaky error occurred");
    }

    var videoData = await db.StringGetAsync(id);
    var response = videoData.IsNull ? "{}" : videoData.ToString();

    using (var myActivity = registeredActivity.StartActivity("do-some-processing"))
    {
      //sleep for x seconds - simulating some processing
      await Task.Delay(TimeSpan.FromSeconds(0.2));
    }

    if (environment == "DEBUG")
    {
        context.Response.Headers.Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE");
        context.Response.Headers.Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-MY-API-Version");
        context.Response.Headers.Add("Access-Control-Allow-Credentials", "true");
        context.Response.Headers.Add("Access-Control-Allow-Origin", "*");
    }

    await context.Response.WriteAsync(response);
});

app.Run();