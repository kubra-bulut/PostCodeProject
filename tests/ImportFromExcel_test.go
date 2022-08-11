package tests

import (
	"PostCodeProject/config"
	"PostCodeProject/models"
	"fmt"
	"github.com/xuri/excelize/v2"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"testing"
)

func TestImportFromExcel(t *testing.T) {
	absPath, err := filepath.Abs("pk_20220413.xlsx")

	if err != nil {
		fmt.Println(err)
	}
	f, err := excelize.OpenFile(absPath)
	if err != nil {
		config.DownloadFile("pk_list.zip", "https://postakodu.ptt.gov.tr/Dosyalar/pk_list.zip")
		config.GetSourceToUnzip("pk_list.zip", "")
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
	filePath := "C:\\Users\\K\\GolandProjects\\PostCodeProject\\postcodes.yaml"
	data, err := yaml.Marshal(&allPostCodes)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = ioutil.WriteFile(filePath, data, 0644)
}

func TestOpenYamlFile(t *testing.T) {
	filePath := "C:\\Users\\K\\GolandProjects\\PostCodeProject\\postcodes.yaml"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	var allPostCodes []models.PostCode
	err = yaml.Unmarshal(data, &allPostCodes)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	for i, code := range allPostCodes {
		fmt.Println(i, code.Code)
	}
}

func TestFile(t *testing.T) {
	var postCode []models.PostCode
	//var cities []models.Name
	filePath := "C:\\Users\\K\\GolandProjects\\PostCodeProject\\postcodes.yaml"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = yaml.Unmarshal(data, &postCode)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

}
