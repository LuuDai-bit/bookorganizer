package services

import (
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
)

const (
	MB                = 1 << 20
	PresignURLTimeout = 15 // minutes
)

type S3Handler struct{}

func (s *S3Handler) S3Instance() (*session.Session, error) {
	session, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_S3_REGION")),
	})

	if err != nil {
		return nil, err
	}

	return session, nil
}

func (s *S3Handler) UploadFile(fileHeader *multipart.FileHeader) (string, error) {
	session, err := s.S3Instance()
	if err != nil {
		return "", err
	}

	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}

	var fileSize int64 = fileHeader.Size
	fileBuffer := make([]byte, fileSize)
	if _, err = file.Read(fileBuffer); err != nil {
		return "", err
	}

	if _, err := file.Seek(0, 0); err != nil {
		return "", err
	}

	key := uuid.New().String()
	_, err = s3.New(session).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(os.Getenv("AWS_S3_BUCKET")),
		Key:                  aws.String(key),
		ACL:                  aws.String("private"),
		Body:                 file,
		ContentLength:        aws.Int64(fileSize),
		ContentType:          aws.String(http.DetectContentType(fileBuffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})

	return key, err
}

func (s *S3Handler) GeneratePresignUrl(key string) (string, error) {
	session, err := s.S3Instance()
	if err != nil {
		return "", err
	}

	req, _ := s3.New(session).GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("AWS_S3_BUCKET")),
		Key:    aws.String(key),
	})

	urlStr, err := req.Presign(PresignURLTimeout * time.Minute)

	if err != nil {
		return "", err
	}

	return urlStr, nil
}

func (s *S3Handler) DeleteObject(key string) error {
	session, err := s.S3Instance()
	if err != nil {
		return err
	}

	_, err = s3.New(session).DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(os.Getenv("AWS_S3_BUCKET")),
		Key:    aws.String(key),
	})

	return err
}
