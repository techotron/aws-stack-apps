package main

import (
	"github.com/gin-gonic/gin"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws/session"
	// "github.com/aws/aws-sdk-go/aws/endpoints"

	"os"
	"fmt"
)


func listDdbTables() {
	// myCustomResolver := func(service, region string, optFns ...func(*endpoints.Options)) (endpoints.ResolvedEndpoint, error) {
	// 	if service == endpoints.DynamodbServiceID {
	// 		return endpoints.ResolvedEndpoint{
	// 			URL:           "s3.custom.endpoint.com",
	// 		}, nil
	// 	}
	
	// 	return endpoints.DefaultResolver().EndpointFor(service, region, optFns...)
	// }


	var aws_region = os.Getenv("AWS_REGION")

	// Create new AWS session using region from env var and default AWS Cred env vars
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(aws_region)},
	)
	
	svc := dynamodb.New(sess)

	input := &dynamodb.ListTablesInput{}

	fmt.Printf("Tables:\n")

	for {
		// Get the list of tables
		result, err := svc.ListTables(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case dynamodb.ErrCodeInternalServerError:
					fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
			return
		}

		for _, n := range result.TableNames {
			fmt.Println(*n)
		}

		// assign the last read tablename as the start for our next call to the ListTables function
		// the maximum number of table names returned in a call is 100 (default), which requires us to make
		// multiple calls to the ListTables function to retrieve all table names
		input.ExclusiveStartTableName = result.LastEvaluatedTableName

		if result.LastEvaluatedTableName == nil {
			break
		}
	}
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"health": "ok",
		})
	})

	r.GET("/listtables", func(c *gin.Context) {
		listDdbTables()
	})	

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}