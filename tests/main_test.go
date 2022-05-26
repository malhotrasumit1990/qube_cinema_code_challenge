package tests

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/qube_cinema_code_challenge/model"
	"github.com/qube_cinema_code_challenge/utils"
	"github.com/stretchr/testify/assert"
)

type TestCases struct {
	Partners        []model.Partner_data
	Deliveries      []model.Delivery_Data
	Delivery_Result []model.Delivery_Result
}

func setup_mock_result(filepath string) []model.Delivery_Result {

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal("Unable to read input file "+filepath, err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.LazyQuotes = true
	records, err := r.ReadAll()
	if err != nil {
		log.Fatalf("Problem in Parsing CSV : %s", err)
	}

	delivery_res := []model.Delivery_Result{}

	for _, record := range records {

		delivery := model.Delivery_Result{}

		delivery.Delivery_ID = string(record[0])
		val, _ := strconv.ParseBool(string(record[1]))
		delivery.Possiblity = val
		if !delivery.Possiblity {
			delivery.Partner_ID = string(record[2])
			delivery.Cost = string(record[3])
		} else {
			delivery.Partner_ID = strings.TrimSpace(string(record[2]))
			delivery.Cost = strings.TrimSpace(record[3])

		}
		delivery_res = append(delivery_res, delivery)
	}

	return delivery_res

}

func TestGet_Best_Possible_Delivery_Partner(t *testing.T) {

	var mock_partners []model.Partner_data
	var mock_delivery_info1 []model.Delivery_Data
	var mock_delivery_info2 []model.Delivery_Data
	var mock_delivery_info3 []model.Delivery_Data

	var mock_delivery_result1 []model.Delivery_Result
	var mock_delivery_result2 []model.Delivery_Result
	var mock_delivery_result3 []model.Delivery_Result

	//Setup Inputs for the code to run
	utils.Setup_Partner_Info("./test_partner_data.csv", &mock_partners)

	utils.Setup_Delivery_Info("./test_input1.csv", &mock_delivery_info1)
	utils.Setup_Delivery_Info("./test_input2.csv", &mock_delivery_info2)
	utils.Setup_Delivery_Info("./test_input3.csv", &mock_delivery_info3)

	mock_delivery_result1 = setup_mock_result("./test_output1.csv")
	mock_delivery_result2 = setup_mock_result("./test_output2.csv")
	mock_delivery_result3 = setup_mock_result("./test_output3.csv")

	cases := []TestCases{
		{Partners: mock_partners, Deliveries: mock_delivery_info1, Delivery_Result: mock_delivery_result1},
		{Partners: mock_partners, Deliveries: mock_delivery_info2, Delivery_Result: mock_delivery_result2},
		{Partners: mock_partners, Deliveries: mock_delivery_info3, Delivery_Result: mock_delivery_result3},
	}

	for _, tc := range cases {
		result := utils.Get_Best_Possible_Delivery_Partner(tc.Partners, tc.Deliveries)

		assert.Equal(t, len(tc.Delivery_Result), len(result))
		assert.Equal(t, strings.TrimSpace(tc.Delivery_Result[0].Partner_ID), result[0].Partner_ID)
		assert.Equal(t, strings.TrimSpace(tc.Delivery_Result[0].Cost), result[0].Cost)

	}
}
