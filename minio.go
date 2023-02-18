package main

import (
	"io"
	"os"
	"strings"

	"github.com/minio/minio-go"
	"github.com/mongodb/mongo-tools/common/log"
)

type bucket struct {
	client     *minio.Client
	bucketName string
}

func newMinio() *bucket {
	accessKey := os.Getenv("DO_ACCESS_KEY")
	secKey := os.Getenv("DO_SECRET_ACCESS_KEY")
	endpoint := os.Getenv("DO_SPACES_ENDPOINT")
	bucketName := os.Getenv("DO_SPACES_BUCKET")
	ssl := true

	// Initiate a client using DigitalOcean Spaces.
	client, err := minio.New(endpoint, accessKey, secKey, ssl)
	if err != nil {
		panic(err)
	}

	return &bucket{
		client:     client,
		bucketName: bucketName,
	}
}

func (b *bucket) upload(filePath, destination string) error {
	_, err := b.client.FPutObject(b.bucketName, destination, filePath, minio.PutObjectOptions{
		ContentType: "application/x-gzip",
	})
	return err
}

func (b *bucket) getLatestBackup(prefix string) ([]byte, error) {
	doneCh := make(chan struct{})
	defer close(doneCh)

	key := prefix + "/"
	// year
	log.Logvf(log.Always, "looking for latest year")
	key += b.getLatestKeyInPath(key)
	// month
	log.Logvf(log.Always, "looking for latest month")
	key += b.getLatestKeyInPath(key)
	// day
	log.Logvf(log.Always, "looking for latest day")
	key += b.getLatestKeyInPath(key)
	// backup
	log.Logvf(log.Always, "looking for latest backup")
	key += b.getLatestKeyInPath(key)

	log.Logvf(log.Always, "fetching backup")
	obj, err := b.client.GetObject(b.bucketName, key, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	log.Logvf(log.Always, "reading backup")
	buffer, err := io.ReadAll(obj)
	return buffer, err
}

func (b *bucket) getLatestKeyInPath(key string) string {
	doneCh := make(chan struct{})
	defer close(doneCh)

	latest := ""
	for message := range b.client.ListObjectsV2(b.bucketName, key, false, doneCh) {
		currentFile := strings.TrimPrefix(message.Key, key)
		if latest < currentFile {
			latest = currentFile
		}
	}
	return latest
}
