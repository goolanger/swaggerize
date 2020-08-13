package document

type Info struct {
	License
	Contact
	Description, Title, Version string
	Terms string `yaml:"termsOfService"`
}
