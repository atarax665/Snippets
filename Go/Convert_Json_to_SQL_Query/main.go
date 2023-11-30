package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type DeviceCluster struct {
	DeviceClusterID   string `json:"deviceClusterID"`
	DeviceClusterName string `json:"deviceClusterName"`
}

func main() {
	filePath := "your/json/file.json"
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	var devices map[string]DeviceCluster
	err = json.Unmarshal(jsonData, &devices)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	file, err := os.Create("output.sql")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	for _, device := range devices {
		sqlQuery := generateSQLQuery(device)
		_, err = file.WriteString(sqlQuery + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
		fmt.Println(sqlQuery)
	}
	defer file.Close()
}

func generateSQLQuery(device DeviceCluster) string {
	return fmt.Sprintf("INSERT INTO devices (device_cluster_id, device_cluster_name, active, created_by, updated_by, created_at, updated_at) VALUES ('%s', '%s', %t, '%s', '%s', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);",
		device.DeviceClusterID,
		device.DeviceClusterName,
		true,
		"SHORT_ID",
		"SHORT_ID",
	)
}
