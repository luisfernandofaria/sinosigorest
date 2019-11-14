package model

type LocalAcidente struct {
	ID        int64  `json:"id"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Endereco  string `json:"endereco"`
	Municipio string `json:"municipio"`
	Cep       string `json:"cep"`
}

func CriaNovoLocalAcidente(latitude string, longitude string, endereco string, municipio string, cep string) LocalAcidente {

	la := LocalAcidente{}

	la.Latitude = latitude
	la.Longitude = longitude
	la.Endereco = endereco
	la.Municipio = municipio
	la.Cep = cep

	return la
}