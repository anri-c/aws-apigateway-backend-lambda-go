package main

import (
        "github.com/aws/aws-lambda-go/events"
        "github.com/aws/aws-lambda-go/lambda"
        "context"
        "time"
        "encoding/json"
)

const (
        layout = "20060102"
        location = "Asia/Tokyo"
)

type Message struct {
        Msg string `json:"msg"`
}

func init() {
        loc, err := time.LoadLocation(location)
        if err != nil {
                loc = time.FixedZone(location, 9*60*60)
        }
        time.Local = loc
}

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
        day := time.Now()
        ua := request.RequestContext.Identity.UserAgent
        ip := request.RequestContext.Identity.SourceIP

        resource_path := request.RequestContext.ResourcePath

        var request_data map[string]string = map[string]string{}


        response_body := Message{
                Msg: "hello",
        }
        jsonBytes, err := json.Marshal(response_body)
        return events.APIGatewayProxyResponse{Body: string(jsonBytes), StatusCode: 200}, nil

}

func main () {
        lambda.Start(handleRequest)
}
