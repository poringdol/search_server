package db

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/gocarina/gocsv"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"yandex-food/utils"
)

const dataPath = "/home/cddoma/csv/"
const preparedDataPath = "/home/cddoma/csv/prepared_data/"

func getCollection(client *mongo.Client) *mongo.Collection {
	return client.Database("yandex_food").Collection("clients")
}

// PrepareFiles preparing csv files - strip multiple spaces, convert all to lower case
func PrepareFiles() {
	fmt.Println("Preparing data files")
	os.Mkdir(preparedDataPath, os.ModePerm)

	files, _ := ioutil.ReadDir(dataPath)
	if len(files) == 0 {
		fmt.Println("no files")
	}

	wg := &sync.WaitGroup{}
	for _, file := range files {
		// пропустим папки
		if file.IsDir() {
			continue
		}
		// пропустим файлы с данными без адресов
		if strings.Count(file.Name(), "no-address") != 0 {
			continue
		}

		wg.Add(1)
		file := file

		go func() {
			defer wg.Done()

			f, _ := os.OpenFile(dataPath+file.Name(), os.O_RDWR|os.O_CREATE, os.ModePerm)

			var buffer bytes.Buffer
			fileScanner := bufio.NewScanner(f)

			for fileScanner.Scan() {
				stripedSpaces := strings.Join(strings.Fields(fileScanner.Text()), " ")
				buffer.WriteString(strings.ToLower(stripedSpaces))
				buffer.WriteString("\n")
			}
			if err := fileScanner.Err(); err != nil {
				utils.CheckError(err, fmt.Sprintf("read file: %s", file.Name()))
			}
			f.Close()

			fNew, err := os.OpenFile(preparedDataPath+file.Name(), os.O_RDWR|os.O_CREATE, os.ModePerm)
			utils.CheckError(err, fmt.Sprintf("open prepared file %s", fNew.Name()))
			fNAme := fNew.Name()
			fmt.Println(fNAme)
			fNew.WriteString(buffer.String())
			fNew.Close()
		}()
	}
	wg.Wait()
}

// SaveDataToDB saving data from csv-files to database
func SaveDataToDB(ctx context.Context, client *mongo.Client) {
	fmt.Println("Saving data to database")

	collection := getCollection(client)

	files, err := ioutil.ReadDir(preparedDataPath)
	utils.CheckError(err, "get file list")

	wg := &sync.WaitGroup{}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		// пропустим файлы с данными без адресов
		if strings.Count(file.Name(), "no-address") != 0 {
			continue
		}

		wg.Add(1)
		file := file
		var csvRecords []*utils.Record

		go func() {
			defer wg.Done()

			f, err := os.OpenFile(preparedDataPath+file.Name(), os.O_RDWR|os.O_CREATE, os.ModePerm)
			utils.CheckError(err, fmt.Sprintf("open file: %s", preparedDataPath+file.Name()))

			err = gocsv.UnmarshalFile(f, &csvRecords)
			utils.CheckError(err, "unmarshal file")

			bsonData := toBsonData(csvRecords)
			if len(bsonData) != 0 {
				_, err = collection.InsertMany(ctx, prepareSlice(bsonData))
				utils.CheckError(err, "insert records to db")
			}
			f.Close()
		}()
	}
	wg.Wait()
}
