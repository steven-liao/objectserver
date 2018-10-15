package main

import (
	"log"

	minio "github.com/minio/minio-go"
)

func main() {
	endpoint := "play.minio.io:9000"
	accessKeyID := "Q3AM3UQ867SPQQA43P2F"
	secretAccessKey := "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
	useSSL := true

	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
		return
	}

	log.Printf("%#v\n", minioClient)

	bucketName := "mymusic"
	location := "us-east-1"

	err = minioClient.MakeBucket(bucketName, location)
	if err != nil {
		exists, err := minioClient.BucketExists(bucketName)
		if err == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	}
	log.Printf("Successfully created %s\n", bucketName)

	// Upload the zip file
	objectName := "master.zip"
	//	filePath := "/tmp/master.zip"
	//	contentType := "application/zip"

	// Upload the zip file with FPutObject
	//n, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	minioClient.RemoveObject(bucketName, objectName)
	n := 1
	//if err != nil {
	//	log.Fatalln(err)
	//}

	log.Printf("Successfully upload %s of size %d\n", objectName, n)
}
