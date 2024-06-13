package helpers

import (
	"fmt"
	"regexp"

	"github.com/Go11Group/Javokhir-A/exam-2/languge-learning-app/api/models"
)

func isUUID(s string) bool {
	fmt.Println(models.User{}.UserId)
	r := regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$`)
	return r.MatchString(s)
}
