package dump

import (
	"log"
	"os"
	"time"

	"github.com/go-yaml/yaml"
)

func Save(saved FileDetails, path string) {
	/* Store the contents of stored.yaml file */
	file, err := os.ReadFile(path)
	if err != nil {
		log.Panicf("panic 1\n %s ", err.Error())
		return
	}

	/* Buffer to be filled */
	yamlFile := &Files{}

	yamlErr := yaml.Unmarshal(file, yamlFile)
	if yamlErr != nil {
		log.Panicf("panic 2\n %s", yamlErr.Error())
		return
	}

	/* Initializes the new file as a File struct */
	new := File{}
	new.Hash = saved.Hash
	new.URL = saved.URL
	new.Size = saved.Size
	new.FileName = saved.Name
	new.Date = time.Now()

	/* Append the new file to the existing storage */
	yamlFile.Stored = append(yamlFile.Stored, new)

	/* Serialize the array back to YAML */
	bytes, err := yaml.Marshal(yamlFile)
	if err != nil {
		log.Panicf("panic yamlmarshal\n %s", err.Error())
		return
	}

	/* Write the results back to the stored.yaml file */
	os.WriteFile(path, bytes, 0666)
}

func Edit() {

}

func Remove() {

}
