package postgres

import (
	"database/sql"
	"github.com/Go11Group/at_lesson/lesson28/model"
)

type StudentRepo struct {
	Db *sql.DB
}

func NewStudentRepo(db *sql.DB) *StudentRepo {
	return &StudentRepo{Db: db}
}

func (u *StudentRepo) GetAllStudents() ([]model.User, error) {
	rows, err := u.Db.Query(`select s.id, s.name, age, gender, c.name from student s
					left join course c on c.id = s.course_id `)
	if err != nil {
		return nil, err
	}

	var users []model.User
	var user model.User
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Gender, &user.Course)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *StudentRepo) GetByID(id string) (model.User, error) {
	var user model.User

	err := u.Db.QueryRow(`select s.id, s.name, age, gender, c.name from student s
					left join course c on c.id = s.course_id where s.id = $1`, id).
		Scan(&user.ID, &user.Name, &user.Age, &user.Gender, &user.Course)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func Create(db *sql.DB, user model.User) error {

	//uuid.NewString()

	//err := db.QueryRow(`select s.id, s.name, age, gender, c.name from student s
	//				left join course c on c.id = s.course_id where s.id = $1`, id).
	//	Scan(&user.ID, &user.Name, &user.Age, &user.Gender, &user.Course)
	//if err != nil {
	//	return err
	//}
	//
	return nil
}

func Update(db *sql.DB, user model.User) error {

	//err := db.QueryRow(`select s.id, s.name, age, gender, c.name from student s
	//				left join course c on c.id = s.course_id where s.id = $1`, id).
	//	Scan(&user.ID, &user.Name, &user.Age, &user.Gender, &user.Course)
	//if err != nil {
	//	return err
	//}

	return nil
}

func Delete(db *sql.DB, id string) error {

	//err := db.QueryRow(`select s.id, s.name, age, gender, c.name from student s
	//				left join course c on c.id = s.course_id where s.id = $1`, id).
	//	Scan(&user.ID, &user.Name, &user.Age, &user.Gender, &user.Course)
	//if err != nil {
	//	return err
	//}

	return nil
}
