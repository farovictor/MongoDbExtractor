package files

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"

	constants "github.com/farovictor/MongoDbExtractor/src/constants"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

// Simple dumper to write json files
func DumpToJsonFile(results []*bson.M, mapping string, filePrefix string, fileLocation string) error {
	// Defining final file name
	var fP string
	fileId, err := uuid.NewUUID()
	if filePrefix == constants.MappingDefault {
		if err != nil {
			return err
		}
		fP = fmt.Sprintf("%s_%s", mapping, fileId)
	} else {
		fP = fmt.Sprintf("%s_%s", filePrefix, fileId)
	}

	// Turn into json
	jsonData, err := json.Marshal(results)
	if err != nil {
		return err
	}

	outputFile := fmt.Sprintf("%s/%s.json", fileLocation, fP)

	if err = ioutil.WriteFile(outputFile, jsonData, 0644); err != nil {
		return err
	}

	return nil
}

// Concurrent batch dumper to write json files through a channel
// Params:
//
//	context: Context in which will run
//	dataSource: Channel that will pass the data
//	mapping: name for data contextualization, used in file name.
//	wg: Waiting Group manager
//	filePrefix: filePrefix name
//	fileLocation: output path where file will be saved
//
// Further reading about concurrency patterns.
// Concurrent design patterns: https://levelup.gitconnected.com/concurrency-design-patterns-in-golang-f0843f570689
// secondary reading: https://blog.devgenius.io/5-useful-concurrency-patterns-in-golang-8dc90ad1ea61
func DumpStreams(ctx context.Context, dataChannel <-chan []*bson.M, mapping string, wg *sync.WaitGroup, filePrefix string, fileLocation string) {
	defer wg.Done()
	for batch := range dataChannel {
		// TODO: Implement this properly
		DumpToJsonFile(batch, mapping, filePrefix, fileLocation)
	}
}
