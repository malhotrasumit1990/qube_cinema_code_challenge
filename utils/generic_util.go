package utils

import (
	"log"
	"strconv"
	"strings"

	"github.com/qube_cinema_code_challenge/model"
)

//Splits String on the basis of delimeter `-`
func split_string(content_range string) (string, string) {
	vals := strings.Split(content_range, "-")
	if len(vals) == 2 {
		return vals[0], vals[1]
	}
	log.Fatal("Content Range is not a valid value , should be like 0-100.")
	return "", ""
}

//Find out the correct range for the amount of content used
func content_Used_In_Range(content_used int, partner_info *model.Partner_data) bool {

	lower_val, upper_val := split_string(partner_info.Content_Size)
	if lower_val != "" && upper_val != "" {

		lv, _ := strconv.Atoi(strings.TrimSpace(lower_val))
		uv, _ := strconv.Atoi(strings.TrimSpace(upper_val))

		if content_used >= lv && content_used < uv {
			return true
		}

	}
	return false
}

//If cost_occurred is less than the minimum cost : This function will return the minimum cost payable
// Else the actual cost
func get_final_cost(partner_info *model.Partner_data, cost_occurred int) int {

	if cost_occurred < partner_info.Min_Cost {
		return partner_info.Min_Cost
	}
	return cost_occurred
}

//Get_Partner_With_Lowest_Cost Get Partner Id with Lowest Cost
func get_Partner_With_Lowest_Cost(partner_map_cost map[string]int) (string, int) {

	partners := make([]string, 0)
	cost := make([]int, 0)

	for k, v := range partner_map_cost {

		partners = append(partners, k)
		cost = append(cost, v)
	}

	min_cost := cost[0]
	partner := partners[0]

	for index, val := range cost {

		if val <= min_cost {
			min_cost = val
			partner = partners[index]

		}
	}

	return partner, min_cost
}
