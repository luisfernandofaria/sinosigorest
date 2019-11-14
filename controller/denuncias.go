package controller
 
import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "sinosigorest/database"
    "strconv"
    "text/template"
 
    "github.com/gorilla/mux"
    "sinosigorest/model"
)
 
var temp = template.Must(template.ParseGlob("templates/*.html"))
 
func Index(w http.ResponseWriter, r *http.Request) {
 
    todasAsDenuncias := model.BuscarDenuncias()
    temp.ExecuteTemplate(w, "Index", todasAsDenuncias)
}
 
func GetAll(w http.ResponseWriter, r *http.Request) {
 
    todasAsDenuncias := model.BuscarDenuncias()
    json.NewEncoder(w).Encode(todasAsDenuncias)
 
}
 
func PostNovaDenuncia(w http.ResponseWriter, r *http.Request) {
 
	var novoLocal model.LocalAcidente
 
    reqBody, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Fprintf(w, "Erro ao tentar inserir uma nova denúncia no banco de dados")
	}
	
	_ = json.NewDecoder(r.Body).Decode(&novoLocal)
	json.NewEncoder(w).Encode(novoLocal)
	fmt.Println(novoLocal.Endereco," <--- endereco")	
 
    var novaDenuncia model.Denuncia
 
	json.Unmarshal(reqBody, &novaDenuncia)
	json.NewEncoder(w).Encode(novaDenuncia)
	fmt.Println(novaDenuncia.Descricao," <--- descricao")	
	
    descricao := novaDenuncia.Descricao
    data := novaDenuncia.Data
    imagem := novaDenuncia.Imagem
    autordano := novaDenuncia.AutorDano
    emailusuario := novaDenuncia.EmailUsuario
    categoria := novaDenuncia.Categoria
    local := novoLocal.ID
 
    db := database.ConectarComBanco()
 
    insereDadosNoBanco, err := db.Prepare("insert into denuncia(descricao, data, imagem, autordano, emailusuario, categoria, local) values($1, $2, $3, $4, $5, $6, $7)")
    if err != nil {
        panic(err.Error())
    }
 
    _, err = insereDadosNoBanco.Exec(descricao, data, imagem, autordano, emailusuario, categoria, local)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()
 
}

//como poderia aproveitar a resposta ou a requisição de um método post/get e criar 2 objetos? ou como criar 2 objetos com uma única requisição?

func NovoPostTeste(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var novaDenuncia model.Denuncia
	_ = json.NewDecoder(r.Body).Decode(&novaDenuncia)
	json.NewEncoder(w).Encode(novaDenuncia)
	fmt.Println(novaDenuncia.EmailUsuario)
	fmt.Println(novaDenuncia.Categoria)
	fmt.Println(novaDenuncia.Descricao)
}

func PostNovoLocal(w http.ResponseWriter, r *http.Request) model.LocalAcidente {
	w.Header().Set("content-type", "application/json")
    var novoLocal model.LocalAcidente
 
    // reqBody, err := ioutil.ReadAll(r.Body)
    // if err != nil {
    //     fmt.Fprintf(w, "Erro ao tentar inserir um novo local no banco de dados")
    // }
    _ = json.NewDecoder(r.Body).Decode(&novoLocal)
    json.NewEncoder(w).Encode(novoLocal)
 
    latitude := novoLocal.Latitude
    longitude := novoLocal.Longitude
    endereco := novoLocal.Endereco
    municipio := novoLocal.Municipio
    cep := novoLocal.Cep
 
	fmt.Println(novoLocal.Municipio, "< -- municipio")

    db := database.ConectarComBanco()
    insereDadosNoBanco, err := db.Prepare("insert into localacidente(latitude, longitude, endereco, municipio, cep) values($1, $2, $3, $4, $5) returning id")
    if err != nil {
        panic(err.Error())
    }
 
    _, err = insereDadosNoBanco.Exec(latitude, longitude, endereco, municipio, cep)
    if err != nil {
        fmt.Println(err)
        return novoLocal //como retornar algo diferente aqui???
    }
    defer db.Close()
    return novoLocal
 
}
 
func GetDenunciaPorId(w http.ResponseWriter, r *http.Request) {
 
    todasAsDenuncias := model.BuscarDenuncias()
 
    denunciaID := mux.Vars(r)["id"]
 
    for _, denuncia := range todasAsDenuncias {
        if denuncia.ID == converterIdStringParaInt64(denunciaID) {
            json.NewEncoder(w).Encode(denuncia)
        }
    }
}
 
func converterIdStringParaInt64(id string) int64 {
 
    idConvertido, err := strconv.ParseInt(id, 10, 64)
    if err == nil {
        return idConvertido
    }
    return 0
}
 
func New(w http.ResponseWriter, r *http.Request) {
    temp.ExecuteTemplate(w, "New", nil)
}
 
func Insert(w http.ResponseWriter, r *http.Request) {
 
    if r.Method == "POST" {
 
        latitude := r.FormValue("latitude")
        longitude := r.FormValue("longitude")
        endereco := r.FormValue("endereco")
        municipio := r.FormValue("municipio")
        cep := r.FormValue("cep")
 
        var localacidente = model.CriaNovoLocalAcidente(latitude, longitude, endereco, municipio, cep)
 
        descricao := r.FormValue("descricao")
        data := r.FormValue("data")
        imagem := r.FormValue("imagem")
        autordano := r.FormValue("autordano")
        emailusuario := r.FormValue("emailusuario")
        categoria := r.FormValue("categoria")
 
        model.CriaNovaDenuncia(descricao, data, imagem, autordano, emailusuario, categoria, localacidente)
    }
    http.Redirect(w, r, "/", 301)
}
 


