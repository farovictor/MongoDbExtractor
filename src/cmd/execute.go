package cmd

import (
	"os"

	mongo "github.com/farovictor/MongoDbDriver"
	constants "github.com/farovictor/MongoDbExtractor/src/constants"
	"github.com/farovictor/MongoDbExtractor/src/files"
	logger "github.com/farovictor/MongoDbExtractor/src/logging"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Execution logic for ping command
func pingExecute(cmd *cobra.Command, args []string) {
	if cmd.Flags().Lookup("conn-uri").Changed {
		ping, err := mongo.Ping(connUri)
		if err != nil {
			logger.ErrorLogger.Println("Error while pinging server:", err)
		}
		if ping {
			logger.InfoLogger.Println("Ping was successful")
		} else {
			logger.WarningLogger.Println("Ping wasnt successful. Check your connection string or network.")
		}
	}
}

// Execution logic for extract command
func extractMapping(cmd *cobra.Command, args []string) {
	if mapping != "" && mapping != constants.MappingDefault {
		logger.InfoLogger.Println("Mapping:", mapping)

		handler, disconnect := mongo.NewConnectionHandler(connUri, dbName, appName)
		defer disconnect()

		coll := handler.GetCollection(collectionName)
		logger.InfoLogger.Println("Collection retrieved")

		var filter bson.D
		if err := bson.UnmarshalExtJSON([]byte(query), true, &filter); err != nil {
			logger.ErrorLogger.Fatalln("Error handling query argument:", err)
		}
		logger.InfoLogger.Println("Filter retrieved", filter)

		options := options.FindOptions{
			BatchSize: &batchSize,
		}
		logger.InfoLogger.Println("Options set!")

		logger.InfoLogger.Println("Processing record")
		if err := handler.ExtractResults(mapping, outputFilePrefix, outputPath, files.DumpToJsonFile, coll, filter, &options); err != nil {
			logger.ErrorLogger.Fatalln(err)
		}
	} else {
		logger.ErrorLogger.Fatalln("Please set a valid mapping")
	}
}

// Execution logic for extract-batches command
func extractBatches(cmd *cobra.Command, args []string) {
	if mapping != "" && mapping != constants.MappingDefault {
		logger.InfoLogger.Println("Mapping:", mapping)

		handler, disconnect := mongo.NewConnectionHandler(connUri, dbName, appName)
		defer disconnect()

		coll := handler.GetCollection(collectionName)
		logger.InfoLogger.Println("Collection retrieved")

		var filter bson.D
		if err := bson.UnmarshalExtJSON([]byte(query), true, &filter); err != nil {
			logger.ErrorLogger.Fatalln("Error handling query argument:", err)
		}
		logger.InfoLogger.Println("Filter retrieved", filter)

		options := options.FindOptions{
			BatchSize: &batchSize,
		}
		logger.InfoLogger.Println("Options set!")

		logger.InfoLogger.Println("Processing record")
		if err := handler.StreamingResults(mapping, outputFilePrefix, outputPath, batchSize, numConcurrentFiles, files.DumpStreams, coll, filter, &options); err != nil {
			logger.ErrorLogger.Fatalln(err)
		}
		logger.InfoLogger.Println("record requested")
	} else {
		logger.ErrorLogger.Fatalln("Please set a valid mapping")
	}
}

// Execution logic for collxst command
func collExistsExecute(cmd *cobra.Command, args []string) {
	if !cmd.Flags().Lookup("app-name").Changed {
		os.Exit(1)
	}
	if cmd.Flags().Lookup("collection").Changed && cmd.Flags().Lookup("db-name").Changed && cmd.Flags().Lookup("conn-uri").Changed {
		handler, disconnect := mongo.NewConnectionHandler(connUri, dbName, appName)
		defer disconnect()
		logger.InfoLogger.Println(handler)
		exist := handler.CollectionExists(collectionName)
		logger.InfoLogger.Printf("Collection %s exists? %v", collectionName, exist)
	}
}
