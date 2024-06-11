package integration

import (
	"log"

	"github.com/Go11Group/Javokhir-A/homework/lesson36/internal/models"
	"github.com/Go11Group/Javokhir-A/homework/lesson36/internal/repositories"
	"github.com/jaswdr/faker"
)

func GenerateUserDataIntoDb(userRepo repositories.UserRepository) {
	fake := faker.New()

	for i := 0; i < 50; i++ {
		user := models.User{}
		user.Id = fake.UUID().V4()
		user.Name = fake.Person().Name()
		user.Password = fake.Internet().Password()
		user.Rank = fake.RandomFloat(3, 15, 11000)
		if err := userRepo.CreateUser(&user); err != nil {
			log.Fatal(err)
		}
	}
}
