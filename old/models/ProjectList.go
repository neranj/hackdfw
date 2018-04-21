package models

import (
	"fmt"
	"bytes"
	"encoding/json"
)

type Project struct {
	projectId      interface{} `json:"projectId"`
	creatorId      interface{} `json:"creatorid"`
	lat            interface{} `json:"lat"`
	long           interface{} `json:"long"`
	createdOn      interface{} `json:"createdOn"`
	updatedOn      interface{} `json:"updatedOn"`
	title          interface{} `json:"title"`
	discription    interface{} `json:"id"`
	address        interface{} `json:"discription"`
	status         interface{} `json:"status"`
	donePercentage interface{} `json:"donePercentage"`
}

func JSONMarshal(t interface{}) (interface{}, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return t, err
}


func GetProjectList() {
	connection := GetDb();
	var market_name string;
	var allProperty []map[string]string;
	queryString := `
					select * from prop

				  `

	stmt_new, err := connection.Query(queryString);
	defer stmt_new.Close()
	if err != nil {
		fmt.Println("ERROR GETTING MARKET DETAILS")
		panic(err)
	}

	for stmt_new.Next() {
		var projectLandingDetails Project
		err = stmt_new.Scan(&projectLandingDetails.projectId,&projectLandingDetails.creatorId,&projectLandingDetails.lat,&projectLandingDetails.long,
		&projectLandingDetails.createdOn,&projectLandingDetails.updatedOn,&projectLandingDetails.title,&projectLandingDetails.discription,&projectLandingDetails.address,&projectLandingDetails.status,&projectLandingDetails.donePercentage)

		b, err := JSONMarshal(projectLandingDetails)

		if err != nil {
			fmt.Println(err)
		}
		allProperty = append(allProperty, b)

	}

	projectList := map[string]string{
		"type": "projectList"  ,
		"list": allProperty,
	}
}