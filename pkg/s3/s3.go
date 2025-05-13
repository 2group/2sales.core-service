package s3

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
	Client     *s3.Client
	BucketName string
	Endpoint   string
}

func NewS3Client(
	ctx context.Context,
	endpoint, bucket, accessKey, secretKey string,
) (*S3Client, error) {
	// Load base config (for region & creds)
	loaders := []func(*config.LoadOptions) error{
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(accessKey, secretKey, ""),
		),
	}

	// If a custom endpoint is passed, override
	if endpoint != "" {
		loaders = append(loaders, config.WithEndpointResolverWithOptions(
			aws.EndpointResolverWithOptionsFunc(
				func(service, region string, options ...interface{}) (aws.Endpoint, error) {
					return aws.Endpoint{
						URL:               endpoint,
						SigningRegion:     region,
						HostnameImmutable: true,
					}, nil
				},
			),
		))
	}

	awsCfg, err := config.LoadDefaultConfig(ctx, loaders...)
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS config: %w", err)
	}

	// Create S3 client with path‐style forcing (needed for MinIO/Ceph)
	client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	return &S3Client{
		Client:     client,
		BucketName: bucket,
		Endpoint:   endpoint,
	}, nil
}

// GeneratePresignedURL generates a presigned URL for uploading a file
func (s *S3Client) GeneratePresignedURL(
	ctx context.Context,
	objectKey string,
	expiration time.Duration,
) (string, error) {
	// 1) Create a presigner from your configured client
	presigner := s3.NewPresignClient(s.Client)

	// 2) Build the PutObjectInput
	input := &s3.PutObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(objectKey),
	}

	// 3) Generate the presigned URL
	resp, err := presigner.PresignPutObject(ctx, input, s3.WithPresignExpires(expiration))
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned PUT URL: %w", err)
	}

	return resp.URL, nil
}

// DownloadFile downloads a file from S3
func (s *S3Client) DownloadFile(objectKey string) ([]byte, error) {
	buf := manager.NewWriteAtBuffer([]byte{})
	downloader := manager.NewDownloader(s.Client)

	_, err := downloader.Download(context.TODO(), buf, &s3.GetObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to download file: %w", err)
	}

	return buf.Bytes(), nil
}

// DeleteFile deletes a file from S3
func (s *S3Client) DeleteFile(objectKey string) error {
	_, err := s.Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	return nil
}

// UploadFile uploads a file to S3 and returns its public URL
func (s *S3Client) UploadFile(objectKey string, filePath string) (string, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Create an uploader
	uploader := manager.NewUploader(s.Client)

	// Upload the file to S3
	_, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(objectKey),
		Body:   file,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	// Construct the S3 object URL
	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.BucketName, objectKey)
	return url, nil
}
