package srv

import (
	"log"

	"google.golang.org/grpc"

	influx "github.com/influxdata/influxdb/client/v2"

	"github.com/datravis/go-meetup/src/go/pipeline/pkg/db"
	"github.com/datravis/go-meetup/src/go/protogen"
)

type PipelineService struct {
	nerClient       protogen.NerServiceClient
	sentimentClient protogen.SentimentServiceClient
	influxClient    influx.Client
}

func NewServer() *PipelineService {
	nerConn, err := grpc.Dial("ner:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	sentimentConn, err := grpc.Dial("sentiment:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	influxClient, err := db.NewInfluxDBClient("http://influxdb:8086")
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	return &PipelineService{
		nerClient:       protogen.NewNerServiceClient(nerConn),
		sentimentClient: protogen.NewSentimentServiceClient(sentimentConn),
		influxClient:    influxClient,
	}
}
