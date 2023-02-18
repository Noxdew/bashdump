package main

import (
	"context"
	"errors"
	"os"

	"github.com/mongodb/mongo-tools/common/log"
	"github.com/mongodb/mongo-tools/common/signals"
	"github.com/mongodb/mongo-tools/common/util"
	"github.com/mongodb/mongo-tools/mongorestore"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func listDatabases() []string {
	uri := os.Getenv("MONGO_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	if err := client.Connect(context.Background()); err != nil {
		panic(err)
	}

	result, err := client.ListDatabaseNames(context.Background(), bson.D{})
	if err != nil {
		panic(err)
	}

	return result
}

func restore(dir string) error {
	uri := os.Getenv("MONGO_URI")
	opts, err := mongorestore.ParseOptions([]string{"--uri", uri, dir}, "bashdump", "bashdump")

	if err != nil {
		log.Logv(log.Always, err.Error())
		log.Logvf(log.Always, util.ShortUsage("mongorestore"))
		return err
	}

	// print help or version info, if specified
	if opts.PrintHelp(false) {
		return errors.New("Help")
	}

	if opts.PrintVersion() {
		return errors.New("Version")
	}

	restore, err := mongorestore.New(opts)
	if err != nil {
		log.Logvf(log.Always, err.Error())
		return err
	}
	defer restore.Close()

	finishedChan := signals.HandleWithInterrupt(restore.HandleInterrupt)
	defer close(finishedChan)

	result := restore.Restore()
	if result.Err != nil {
		log.Logvf(log.Always, "Failed: %v", result.Err)
	}

	if restore.ToolOptions.WriteConcern.Acknowledged() {
		log.Logvf(log.Always, "%v document(s) restored successfully. %v document(s) failed to restore.", result.Successes, result.Failures)
	} else {
		log.Logvf(log.Always, "done")
	}

	return result.Err
}
