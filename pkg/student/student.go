package student

type Student struct {
	name 			string
	surname 		string
	id 				string
	admissionDate 	string
}

func (student Student) GetName() string {
	return student.name
}

func (student Student) GetSurname() string {
	return student.surname
}

func (student Student) GetId() string {
	return student.id
}

func (student Student) GetAdmissionDate() string {
	return student.admissionDate
}

func (student *Student) SetName(name string) {
	student.name = name
}

func (student *Student) SetSurname(surname string) {
	student.surname = surname
}

func (student *Student) SetId(id string) {
	student.id = id
}

func (student *Student) SetAdmissionDate(admissionDate string) {
	student.admissionDate = admissionDate
}

func (student Student) GetDetails() (string, string, string, string) {
	return student.name, student.surname, student.id, student.admissionDate
}