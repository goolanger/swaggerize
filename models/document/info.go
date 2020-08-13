package document

type Info struct {
	Licence
	Contact
	Description, Title, Version string
	Terms string `yaml:"termsOfService"`
}
