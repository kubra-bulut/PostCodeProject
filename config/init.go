package config

import (
	"PostCodeProject/libs"
	"PostCodeProject/models"
	"archive/zip"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/xuri/excelize/v2"
	"gopkg.in/yaml.v3"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	LocalIPAddress = GetOutboundIP()
	Port           = 8760
	Cities         []models.City
	Counties       []models.County
	AllPostCodes   []models.PostCode
	Towns          []models.Town
)

// performans issue
func Init() {
	startTime := time.Now()
	filePath, _ := libs.GetAppDataPath("postcodes.yaml")
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = yaml.Unmarshal(data, &AllPostCodes)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	citiesMap := make(map[string]models.City)
	countiesMap := make(map[string]models.County)
	townsMap := make(map[string]models.Town)
	for _, postCode := range AllPostCodes {
		citiesMap[postCode.City] = models.City{
			Name: postCode.City,
		}
		countiesMap[fmt.Sprintf("%s.%s", postCode.City, postCode.County)] = models.County{
			City:   postCode.City,
			County: postCode.County,
		}

		townsMap[fmt.Sprintf("%s.%s.%s", postCode.City, postCode.County, postCode.Town)] = models.Town{
			City:   postCode.City,
			County: postCode.County,
			Town:   postCode.Town,
		}
	}
	for _, city := range citiesMap {
		Cities = append(Cities, city)
	}

	for _, county := range countiesMap {
		Counties = append(Counties, county)
	}

	for _, town := range townsMap {
		Towns = append(Towns, town)
	}

	fmt.Println("Time elapsed: ", time.Since(startTime))
}

//DownloadFile Downloads the file from the given url
func DownloadFile(filePath, url string) error {
	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//if err != nil {
	//	return err
	//}
	//fmt.Println(dir, filePath, url)
	// Create a Resty Client
	client := resty.New()
	fmt.Printf("Downloading %s, %s\n", url, filePath)
	// HTTP response gets saved into file, similar to curl -o flag
	if _, err := client.R().
		SetOutput(filePath).
		Get(url); err != nil {
		return err
	}

	fmt.Println("File downloaded")
	return nil //check this line later
}
func GetSourceToUnzip(source, destination string) error {

	// 1. Open the zip file
	reader, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer reader.Close()

	// 2. Get the absolute destination path
	destination, err = filepath.Abs(destination)
	if err != nil {
		return err
	}

	// 3. Iterate over zip files inside the archive and unzip each of them
	for _, f := range reader.File {

		err := UnzipFile(f, destination)
		if err != nil {
			return err
		}

	}

	return nil

}

// UnzipFile unzip the given zip file
func UnzipFile(f *zip.File, destination string) error {
	// 4. Check if file paths are not vulnerable to Zip Slip
	filePath := filepath.Join(destination, f.Name)
	if !strings.HasPrefix(filePath, filepath.Clean(destination)+string(os.PathSeparator)) {
		return fmt.Errorf("invalid file path: %s", filePath)
	}

	// 5. Create directory tree
	if f.FileInfo().IsDir() {
		if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
			return err
		}
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	// 6. Create a destination file for unzipped content
	destinationFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return err
	}

	defer destinationFile.Close()

	// 7. Unzip the content of a file and copy it to the destination file
	zippedFile, err := f.Open()
	if err != nil {
		return err
	}
	defer zippedFile.Close()

	if _, err := io.Copy(destinationFile, zippedFile); err != nil {
		return err
	}
	return nil
}

func ImportFromExcel(filepath string) {
	f, err := excelize.OpenFile(filepath)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	var postcodes []models.PostCode
	var allPostCodes []models.PostCode
	for i, row := range rows {
		if i == 0 {
			continue
		}
		postcode := models.PostCode{
			City:        strings.TrimSpace(row[0]),
			County:      strings.TrimSpace(row[1]),
			Town:        strings.TrimSpace(row[2]),
			District:    strings.TrimSpace(row[3]),
			Code:        strings.TrimSpace(row[4]),
			CountryCode: "TR",
		}
		allPostCodes = append(allPostCodes, postcode)
		postcodes = append(postcodes, postcode)
	}
	if len(postcodes) > 0 {
		fmt.Println("Remaining records inserted")
	}
	// tüm postcodeları kaydedelim
	filePath, _ := libs.GetAppDataPath("postcodes.yaml")
	data, err := yaml.Marshal(&allPostCodes)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = ioutil.WriteFile(filePath, data, 0644)
}

func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Println("GetOutboundIP", err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	return localAddr[0:idx]
}
