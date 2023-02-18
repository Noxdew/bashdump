package main

import (
	"bytes"
	"errors"
	"os"
	"path"
	"time"

	"github.com/mongodb/mongo-tools/common/log"
)

const workdir = "dump"

func main() {
	prefix := os.Getenv("DO_SPACES_PREFIX")

	if len(os.Args) < 2 {
		panic(errors.New("usage: bashdump [dump|restore]"))
	}

	if os.Args[1] == "dump" {
		// Backup DB
		dump()

		// TAR GZ the resulting folder
		now := time.Now()
		fileName := "dump" + now.Format("2006-01-02-15-04-05") + "_" + prefix + ".tar.gz"

		log.Logvf(log.Always, "opening archive file %v", fileName)
		fileToWrite, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, os.FileMode(0755))
		if err != nil {
			panic(err)
		}

		log.Logvf(log.Always, "archiving file %v", fileName)
		if err := tar(workdir, fileToWrite); err != nil {
			panic(err)
		}

		// Upload to the bucket
		storageClient := newMinio()
		path := prefix + now.Format("/2006/01/02/") + fileName
		log.Logvf(log.Always, "uploading %v", path)
		if err := storageClient.upload(fileName, path); err != nil {
			panic(err)
		}
	} else if os.Args[1] == "restore" {
		storageClient := newMinio()
		log.Logvf(log.Always, "looking for latest backup")
		buffer, err := storageClient.getLatestBackup(prefix)
		if err != nil {
			panic(err)
		}

		// Untar backup
		log.Logvf(log.Always, "uncompressing archive")
		if err := untar(workdir, bytes.NewReader(buffer)); err != nil {
			panic(err)
		}

		// restore backup
		files, err := os.ReadDir(workdir)
		if err != nil {
			panic(err)
		}

		deleted := 0
		databases := listDatabases()
		for _, file := range files {
			if !file.IsDir() {
				continue
			}

			// delete folders where databases already exist
			if contains(databases, file.Name()) {
				os.RemoveAll(path.Join(workdir, file.Name()))
				deleted++
			}
		}

		if err := restore(workdir); err != nil {
			panic(err)
		}
	} else {
		panic(errors.New("usage: bashdump [dump|restore]"))
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
