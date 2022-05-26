package utils

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gocarina/gocsv"
	"github.com/qube_cinema_code_challenge/model"
)

//Utility to read CSV
func readCsvFile(filePath string) *os.File {

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	return file
}

func Setup_Partner_Info(filepath string, partners *[]model.Partner_data) {

	file := readCsvFile(filepath)
	defer file.Close()

	if err := gocsv.UnmarshalFile(file, partners); err != nil { // Load clients from file
		log.Fatalf("Unmarshal problem : %s", err)
	}
}

func Setup_Delivery_Info(filepath string, delivery_info *[]model.Delivery_Data) {

	file := readCsvFile(filepath)
	defer file.Close()

	r := csv.NewReader(file)
	r.LazyQuotes = true
	records, err := r.ReadAll()
	if err != nil {
		log.Fatalf("Problem in Parsing CSV : %s", err)
	}

	for _, record := range records {

		delivery := model.Delivery_Data{}

		delivery.Delivery_ID = string(record[0])
		val, _ := strconv.Atoi(strings.TrimSpace((record[1])))
		delivery.Content_Size = val
		delivery.Theatre_ID = strings.TrimSpace(string(record[2]))

		*delivery_info = append(*delivery_info, delivery)
	}

}

func Get_Best_Possible_Delivery_Partner(partners []model.Partner_data, delivery_info []model.Delivery_Data) []model.Delivery_Result {

	delivery_result := []model.Delivery_Result{}

	//First loop over the Input file to find which Theatre we need to supply cntent to
	for _, delivery := range delivery_info {
		partner_cost_map := make(map[string]int)

		// Second loop to fing out best Partner (with lowest cost) for identified Theatre
		for _, partner := range partners {
			if strings.TrimSpace(partner.Theatre_ID) == strings.TrimSpace(delivery.Theatre_ID) {
				if content_Used_In_Range(delivery.Content_Size, &partner) {
					cost_occured := partner.Cost_PerGB * delivery.Content_Size
					partner_cost_map[partner.Partner_ID] = get_final_cost(&partner, cost_occured)
				}
			}
		}
		result := model.Delivery_Result{}
		if len(partner_cost_map) != 0 {
			partner_id, cost := get_Partner_With_Lowest_Cost(partner_cost_map)
			make_result_obj(&result, partner_id, delivery.Delivery_ID, cost, true)
		} else {
			make_result_obj(&result, "\"\"", delivery.Delivery_ID, 0, false)
		}

		//Append the final result in the delivery result slice.
		delivery_result = append(delivery_result, result)

	}

	return delivery_result

}

//Helper function to make final Delivery object.
func make_result_obj(obj *model.Delivery_Result, partner_id, devlivery_id string, cost int, possiblity bool) {

	if cost == 0 {
		obj.Cost = "\"\""
	} else {
		obj.Cost = strconv.Itoa(cost)
	}
	obj.Delivery_ID = devlivery_id
	obj.Partner_ID = partner_id
	obj.Possiblity = possiblity

}

func Write_Result_CSV(records []model.Delivery_Result) {

	file, err := os.Create("./output_file/delivery_partners_assigned.csv")
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(file)
	defer w.Flush()

	for _, record := range records {
		row := []string{record.Delivery_ID, strconv.FormatBool(record.Possiblity), record.Partner_ID, record.Cost}
		if err := w.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}

}
