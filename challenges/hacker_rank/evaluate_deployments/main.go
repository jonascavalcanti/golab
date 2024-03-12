package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	deployments := []string{
		`{"deployment_id": "d-12345678ab", "status": "Success"}`,
		`{"deployment_id": "d-12345678cd", "status": "ABCDE"}`,
		`{"deployment_id": "d-12345678cd", "status": "Fail"}`,
	}

	result := evaluate_deployments(deployments)
	fmt.Println(result) // Output: [2 1 1]
}

func evaluate_deployments(deployments []string) []int32 {
	success := 0
	fail := 0
	errorcount := 0

	for _, deployment := range deployments {
		var data map[string]interface{}
		if err := json.Unmarshal([]byte(deployment), &data); err != nil {
			errorcount++
			continue
		}

		deploy_id, error_id := data["deployment_id"].(string)
		status, error_status := data["status"].(string)

		fmt.Println(string(status))

		if len(deploy_id) > 0 && (!error_id || !error_status || len(deploy_id) != 12 || deploy_id[:2] != "d-" || !isValidString(deploy_id)) {
			errorcount++
			continue
		}

		if strings.TrimSpace(status) == "Success" {
			success++
		} else if strings.TrimSpace(status) == "Fail" {
			fail++
		} else {
			errorcount++
		}
	}

	return []int32{int32(success), int32(fail), int32(errorcount)}
}

func isValidString(s string) bool {
	validFormat := regexp.MustCompile(`^[0-9a-zA-Z]{10}$`)
	return validFormat.MatchString(s[:2])
}
