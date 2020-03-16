package main

import (
	"flag"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var bucket = flag.String("bucket", "my-bucket-name", "s3 bucket")
var objectKey = flag.String("key", "key/to/my/object.txt", "object key to fetch")
var objectVersion = flag.String("version-id", "HF7gxoC4Yo4z9PR5fQ2173JHOXYY7in1", "object version id to fetch")

func main() {
	flag.Parse()

	sess := session.Must(session.NewSession())
	svc := s3.New(sess)

	output, err := svc.GetObject(&s3.GetObjectInput{
		Bucket:    aws.String(*bucket),
		Key:       aws.String(*objectKey),
		VersionId: aws.String(*objectVersion),
	})

	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, output.Body)
	if err = output.Body.Close(); err != nil {
		panic(err)
	}
}
