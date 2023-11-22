package pkg

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	// "exchange-rate-api/db"
)

// check for the file in directory locations
// If file not found then create it
func checkFile(dir string) error {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		_, err := os.Create(dir)
		if err != nil {
			return err
		}
	}
	return nil
}

// save struct data into file.json, if file doesn't exists this create the file
// take struct data also file location and file names in 1 string as input
// directory ex: "../../assets/file.json"
func WriteJsonFile(directory string, data interface{}) (err error) {
	var byteData bytes.Buffer
	encodeData := json.NewEncoder(&byteData)
	encodeData.SetIndent("", "\t")
	encodeData.Encode(data)

	// json.MarshalIndent is used to make result pretty-formatted,
	// byteData, err := json.MarshalIndent(data, "", "\t")
	checkFile(directory)

	err = os.WriteFile(directory, byteData.Bytes(), 0644)
	if err != nil {
		log.Printf("file %s failed to write into json, err: \n%s\n", directory, err)
		return err
	}
	return nil
}
