package services

import (
	"mime/multipart"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
)

const MB = 1 << 20

type S3Handler struct{}

func (s *S3Handler) S3Instance() *session.Session {
	session, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_S3_REGION")),
	})

	if err != nil {
		// handle
	}

	return session
}

func (s *S3Handler) UploadFile(fileHeader *multipart.FileHeader) (string, error) {
	session := s.S3Instance()

	file, err := fileHeader.Open()
	if err != nil {
		panic(err)
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
