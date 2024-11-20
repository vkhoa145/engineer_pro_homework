package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/vkhoa145/engineer_pro_homework/day3/application/models"
)

func OpenFile() *os.File {
	file, err := os.OpenFile("users.txt", os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}

	return file
}

func FindUser(key int, value string) *models.User {
	file := OpenFile()
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			fmt.Println("Empty file")
			return nil
		}

		parts := strings.Split(line, ",")
		if len(parts) > 0 && parts[key] == value {
			id, err := strconv.Atoi(parts[0])
			if err != nil {
				fmt.Println("Error when parsing:", err)
				return nil
			}

			return &models.User{
				ID:          id,
				Username:    parts[1],
				Password:    parts[2],
				UserProfile: parts[3],
			}
		}
	}

	return nil
}
