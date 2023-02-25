package util

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Info struct {
	AwsS3Region    string
	AwsAccessKey   string
	AwsSecretKey   string
	AwsProfileName string
	BucketName     string
	S3Client       *s3.Client
}

// key를 활용해서 Client 생성
func (s *S3Info) SetS3ConfigByKey() error {
	creds := credentials.NewStaticCredentialsProvider(s.AwsAccessKey, s.AwsSecretKey, "")
	cfg, err := awsConfig.LoadDefaultConfig(context.TODO(), awsConfig.WithCredentialsProvider(creds),
		awsConfig.WithRegion(s.AwsS3Region),
	)
	if err != nil {
		log.Printf("error: %v", err)
		panic(err)

	}
	s.S3Client = s3.NewFromConfig(cfg)
	return nil
}

// 서버를 통해 파일을 받아왔을 때 사용
func (s *S3Info) UploadFile(file io.Reader, filename, preFix string) *manager.UploadOutput {
	uploader := manager.NewUploader(s.S3Client)
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return result
}

// 버킷 리스트 확인
func (s *S3Info) GetBucketList() {
	output, err := s.S3Client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		panic(err)
	}
	for _, bucket := range output.Buckets {
		fmt.Println(*bucket.Name)
	}
}

// 버킷안에 있는 Objects 확인할때 사용
func (s *S3Info) GetItems(prefix string) {
	paginator := s3.NewListObjectsV2Paginator(s.S3Client, &s3.ListObjectsV2Input{
		Bucket: &s.BucketName,
		Prefix: aws.String(prefix),
	})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Fatalln("error:", err)
		}
		for _, obj := range page.Contents {
			fmt.Println(aws.ToString(obj.Key))
		}
	}
}
