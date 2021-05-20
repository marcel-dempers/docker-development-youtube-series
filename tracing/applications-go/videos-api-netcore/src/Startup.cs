using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.HttpsPolicy;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using Microsoft.Extensions.Logging;
using Microsoft.OpenApi.Models;
using Jaeger;
using Jaeger.Reporters;
using Jaeger.Samplers;
using Jaeger.Senders;
using Jaeger.Senders.Thrift;
using OpenTracing;
using OpenTracing.Contrib.NetCore.Configuration;
using OpenTracing.Util;


namespace videos_api
{
    public class Startup
    {
        public Startup(IConfiguration configuration)
        {
            Configuration = configuration;
        }

        public IConfiguration Configuration { get; }

        // This method gets called by the runtime. Use this method to add services to the container.
        public void ConfigureServices(IServiceCollection services)
        { 
            services.AddLogging(loggingBuilder =>
            {
                loggingBuilder.AddConfiguration(Configuration.GetSection("Logging"));
                loggingBuilder.AddConsole();
                loggingBuilder.AddDebug();
            });

            services.AddSingleton<ITracer>(serviceProvider =>  
            {   
                ILoggerFactory loggerFactory = serviceProvider.GetRequiredService<ILoggerFactory>(); 
                Jaeger.Configuration.SenderConfiguration.DefaultSenderResolver = new SenderResolver(loggerFactory)
                  .RegisterSenderFactory<ThriftSenderFactory>();
                 
                var config = Jaeger.Configuration.FromEnv(loggerFactory);
                ITracer tracer = config.GetTracer();
                GlobalTracer.Register(tracer);  
                return tracer;  

                // var loggerFactory = serviceProvider.GetRequiredService<ILoggerFactory>();
                // var sampler = new ConstSampler(sample: true);

                // var tracer = new Tracer.Builder("videos-api")
                // .WithReporter(
                //     new RemoteReporter.Builder()
                //         .WithLoggerFactory(loggerFactory)
                //         .WithSender(
                //           new UdpSender("jaeger", 6831, 0))
                //         .Build())
                // .WithLoggerFactory(loggerFactory)
                // .WithSampler(sampler)
                // .Build();
                    
                // GlobalTracer.Register(tracer);
                // return tracer;

            });  

            services.AddOpenTracing();
           
            services.AddControllers();
            services.AddSwaggerGen(c =>
            {
                c.SwaggerDoc("v1", new OpenApiInfo { Title = "videos_api", Version = "v1" });
            });
        }

        // This method gets called by the runtime. Use this method to configure the HTTP request pipeline.
        public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
        {
            if (env.IsDevelopment())
            {
                app.UseDeveloperExceptionPage();
                app.UseSwagger();
                app.UseSwaggerUI(c => c.SwaggerEndpoint("/swagger/v1/swagger.json", "videos_api v1"));
            }

            app.UseHttpsRedirection();
            app.UseRouting();
            app.UseAuthorization();
            app.UseEndpoints(endpoints =>
            {
                endpoints.MapControllers();
            });
        }
    }
}
