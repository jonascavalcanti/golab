package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func main() {
	barcode := "0006EGS8MITUYJ|000618L1QPXTWZ|0003SR7H3AHF8D|00047XZOXTBJDB|000153CHLNV06Q|0007YKA8V3AQFL|0008CU15NFAFIW|0009IDT15ULJEE|00020PY5Z0NETT"
	//barcode := "0001"
	//configuration2 := "0002f7c22e7904|000176a3a4d214|000205d29f4a4b"

	fmt.Println(ordered_barcode(barcode))
	//fmt.Println(ordered_barcode(configuration2))
}

func ordered_barcode(configuration string) []string {
	// Write your code here

	if len(configuration) < 4 || strings.Contains(configuration, "||") || strings.Contains(configuration, ">") {
		return []string{"Invalid configuration"}
	}

	configurations := strings.Split(configuration, "|")

	if len(configurations) < 0 {
		return []string{"Invalid configuration"}
	}

	indices := make(map[string]bool)
	for _, config := range configurations {
		ordinalIndex := config[:4]

		if _, exists := indices[ordinalIndex]; exists {
			return []string{"Invalid configuration"}
		}
		indices[ordinalIndex] = true
	}

	for i := 1; i <= len(configurations); i++ {
		index := fmt.Sprintf("%04d", i)

		if _, exists := indices[index]; !exists {
			return []string{"Invalid configuration"}
		}
		config := configurations[i-1]
		if !isValidConfiguration(config) {
			return []string{"Invalid configuration"}
		}
	}

	sort.Slice(configurations, func(i, j int) bool {
		return configurations[i][:4] < configurations[j][:4]
	})

	// Extract configuration values
	var result []string
	for _, config := range configurations {
		result = append(result, config[4:])
	}

	return result
}

func isValidConfiguration(config string) bool {
	validFormat := regexp.MustCompile(`^[0-9a-zA-Z]{10}$`)
	return validFormat.MatchString(config[4:])
}
