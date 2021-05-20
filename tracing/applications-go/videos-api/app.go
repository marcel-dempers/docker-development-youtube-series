package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/go-redis/redis/v8"
	"fmt"
	"context"
	"time"
	"strings"
	"os"
	"math/rand"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/opentracing/opentracing-go/ext"
)

//TODO : https://opentracing.io/guides/golang/quick-start/
// docker run --rm -it -p 6831:6831/udp -p 16686:16686 -p 14269:14269  --net tracing --name jaeger jaegertracing/all-in-one:latest

const serviceName = "videos-api"

var environment = os.Getenv("ENVIRONMENT")
var redis_host = os.Getenv("REDIS_HOST")
var redis_port = os.Getenv("REDIS_PORT")
var jaeger_host_port = os.Getenv("JAEGER_HOST_PORT")
var flaky = os.Getenv("FLAKY")
var delay = os.Getenv("DELAY")

var ctx = context.Background()
var rdb *redis.Client

func main() {
	
  cfg := &config.Configuration{
		ServiceName: serviceName,

		// "const" sampler is a binary sampling strategy: 0=never sample, 1=always sample.
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},

		// Log the emitted spans to stdout.
		Reporter: &config.ReporterConfig{
			LogSpans: true,
			LocalAgentHostPort: jaeger_host_port,
		},
	}

	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)
	
	router := httprouter.New()

	router.GET("/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params){
		
		spanCtx, _ := tracer.Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(r.Header),
		)

		span := tracer.StartSpan("videos-api: GET /id", ext.RPCServerOption(spanCtx))
		defer span.Finish()

		if flaky == "true" {
			if rand.Intn(90) < 30 {
				panic("flaky error occurred ")
		  } 
		}
		
		ctx := opentracing.ContextWithSpan(context.Background(), span)
		video := video(w,r,p, ctx)

		if strings.Contains(video, "jM36M39MA3I") && delay == "true" {
					time.Sleep(6 * time.Second)
		}

		cors(w)
		fmt.Fprintf(w, "%s", video)
	 
	})

	r := redis.NewClient(&redis.Options{
		Addr:     redis_host + ":" + redis_port,
		DB:       0,
	})
	rdb = r

	fmt.Println("Running...")
	log.Fatal(http.ListenAndServe(":10010", router))
}

func video(writer http.ResponseWriter, request *http.Request, p httprouter.Params, ctx context.Context)(response string){
	
	span, _ := opentracing.StartSpanFromContext(ctx, "videos-api: redis-get")
	defer span.Finish()
	id := p.ByName("id")
	
	videoData, err := rdb.Get(ctx, id).Result()
	if err == redis.Nil {

		span.Tracer().Inject(
			span.Context(),
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(request.Header),
		)
		
		return "{}"
		
		
	} else if err != nil {
		panic(err)
	} else {
	  
		span.Tracer().Inject(
			span.Context(),
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(request.Header),
		)

		return videoData
	}
}

type stop struct {
	error
}

func cors(writer http.ResponseWriter) () {
	if(environment == "DEBUG"){
		writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-MY-API-Version")
		writer.Header().Set("Access-Control-Allow-Credentials", "true")
		writer.Header().Set("Access-Control-Allow-Origin", "*")
	}
}

type videos struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Imageurl string `json:"imageurl"`
	Url string `json:"url"`

}
