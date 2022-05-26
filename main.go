package main

import (
	"fmt"

	"github.com/qube_cinema_code_challenge/model"
	"github.com/qube_cinema_code_challenge/utils"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var partners []model.Partner_data
	var delivery_info []model.Delivery_Data

	//Setup Inputs for the code to run
	utils.Setup_Partner_Info("./input_files/partners.csv", &partners)
	utils.Setup_Delivery_Info("./input_files/input.csv", &delivery_info)

	//Final Result with Delivery and Partner Information
	var delivery_result []model.Delivery_Result

	delivery_result = utils.Get_Best_Possible_Delivery_Partner(partners, delivery_info)

	utils.Write_Result_CSV(delivery_result)

	fmt.Println(delivery_result)

}
