package store

import (
	"context"
	"io"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gabrieldiaziv/wghidra/app/system"
	"github.com/google/uuid"
)

func (s *storeS3) getObject(ctx context.Context, bucket string, key string) (io.ReadCloser, error) {
	res, err := s.svc.GetObjectWithContext(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

func (s *storeS3) putObject(ctx context.Context, bucket string, stream io.ReadCloser) (string, error) {
	key := uuid.New().String()

	_, err := s.svc.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Body:         nil,
		Bucket:       aws.String(system.Unwrap(bucket)),
		Key:          aws.String(key),
		CacheControl: aws.String(cacheControl),
	})

	if err != nil {
		log.Printf("Store.POST: %v", err)
		return "", err
	}

	return key, nil
}

func (s *storeS3) GetExe(ctx context.Context, id string) (io.ReadCloser, error) {
	return s.getObject(ctx, system.Unwrap(system.ENV.AWS.ExeBucket), id)
}
func (s *storeS3) PostExe(ctx context.Context, stream io.ReadCloser) (string, error) {
	return s.putObject(ctx, system.Unwrap(system.ENV.AWS.ExeBucket), stream)
}
func (s *storeS3) GetDecompiled(ctx context.Context, id string) (io.ReadCloser, error) {
	return s.getObject(ctx, system.Unwrap(system.ENV.AWS.DecBucket), id)
}
func (s *storeS3) PostDecompiled(ctx context.Context, stream io.ReadCloser) (string, error) {
	return s.putObject(ctx, system.Unwrap(system.ENV.AWS.DecBucket), stream)
}
