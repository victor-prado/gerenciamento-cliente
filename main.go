package main

import(
  "fmt"
  "log"
  "net/http"
	"html/template"
	"gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

//Struct para armazenar informações dos clientes
type Cliente struct {
  Nome string
  Email string
  Nasc string
  Doc string
  Tipo string
  End string
}

//Struct para armazenar o resultado de uma query de clientes
type Clientes struct {
  Resultado []Cliente
}

//Handler para a pagina inicial:
//Conversa com o banco de dados e com a página HTML
func IndexHandler(w http.ResponseWriter, r *http.Request) {
  url := "localhost"
  //inicia uma sessão no mongodb e trato um possivel erro
  sessao, err := mgo.Dial(url)
  if err != nil {
    log.Fatal(err)
  }
  //Abre a coleção cliente
  colecao := sessao.DB("empresa").C("cliente")


  //Extrair o texto da form submetida pelo usuario
  r.ParseForm()
	nome := r.PostFormValue("busca")
  //Realizar a query pelo texto inserido
  clientes := Clientes{}
  resultado := colecao.Find(bson.M{"nome": bson.RegEx{nome, ""}}).Iter()
  cliente := Cliente{}
  for resultado.Next(&cliente) {
    clientes.Resultado = append(clientes.Resultado, cliente)
  }
  //Fechar a sessão do mongodb
  sessao.Close()
  //Abrimos nosso template da pagina inicial
  t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, clientes)
}

func main() {
  fmt.Println("Listening at http://localhost:8082/")
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":8082", nil) //Servidor
}
