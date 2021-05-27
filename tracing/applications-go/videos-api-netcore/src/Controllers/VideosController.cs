using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using StackExchange.Redis;
using System.Text.Json;
using System.Text.Json.Serialization;
using OpenTracing;
using OpenTracing.Propagation;

namespace videos_api.Controllers
{
    [ApiController]
    [Route("")]
    public class VideosController : ControllerBase
    {
        private readonly ITracer _tracer;
        private readonly string _redisHost;
        private readonly string _redisPort;
        private readonly ConnectionMultiplexer _redis; 
        private readonly ILogger<VideosController> _logger;
        private readonly JsonSerializerOptions _serializationOptions;

        public VideosController(ILogger<VideosController> logger, ITracer tracer)
        {
            _redisHost = Environment.GetEnvironmentVariable("REDIS_HOST");
            _redisPort = Environment.GetEnvironmentVariable("REDIS_PORT");
            _redis = ConnectionMultiplexer.Connect(_redisHost + ":" + _redisPort);
            _logger = logger;
            _tracer = tracer ?? throw new ArgumentNullException(nameof(tracer));
             
            _serializationOptions = new JsonSerializerOptions
            {
                PropertyNameCaseInsensitive = true
            };
        }

        [HttpGet]
        [Route("/{id}")]
        public Videos Get(string id)
        {    
            ISpanContext traceContext = _tracer.Extract(BuiltinFormats.HttpHeaders, new TextMapExtractAdapter(GetHeaders()));

            string videoContent;
            
            using (var scope = _tracer.BuildSpan("videos-api-net: redis-get")
                                .AsChildOf(traceContext)
                                .StartActive(true))
            {

              
              
              IDatabase db = _redis.GetDatabase();
              videoContent = db.StringGet(id);
            }

            var video = JsonSerializer.Deserialize<Videos>(videoContent,_serializationOptions);
            return video;

        }

        public Dictionary<string, string> GetHeaders()
        {
            var headers = new Dictionary<string, string>();
            foreach (var header in Request.Headers)
            {
                headers.Add(header.Key, header.Value);
            }

            return headers;
        }
    }
}
