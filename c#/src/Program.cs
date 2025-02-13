using Microsoft.AspNetCore.Hosting;

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();
builder.Services.Configure<HostOptions>(builder.Configuration.GetSection("HostOptions"));

builder.WebHost.ConfigureKestrel(serverOptions =>
{
    serverOptions.ListenAnyIP(5000);
});

var app = builder.Build();

Console.WriteLine(Environment.ProcessorCount.ToString());

app.UseHttpsRedirection();

app.MapGet("/", async () =>
{
    await Task.Delay(40000); // Sleep for 40 seconds
    return "Hello World!";
});

app.Run();