package store

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gabrieldiaziv/wghidra/app/bo/iface"
)

const (
	cacheControl = "max-age=172800"
)

type storeS3 struct {
	sess *session.Session
	svc  *s3.S3
}

func NewStore(ctx context.Context) iface.Store {
	sess := session.Must(session.NewSession())
	svc := s3.New(sess)

	return &storeS3{
		sess: sess,
		svc:  svc,
	}

}
