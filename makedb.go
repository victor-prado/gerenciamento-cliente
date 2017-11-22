package main
/*Este código foi criado para inserir alguns clientes genéricos
no banco de dados para caso o usuário queira testar e se familiarizar
com a aplicação
*/

import(
  "fmt"
  "log"
  "gopkg.in/mgo.v2"
  //"gopkg.in/mgo.v2/bson"
)

type Cliente struct {
  Nome string
  Email string
  Nasc string
  Doc string
  Tipo string
  End string
}

func main() {
  //Abrindo uma sessão no mongoDB
  url := "localhost"
  sessao, err := mgo.Dial(url)
  if err != nil {
    log.Fatal(err)
  }
  colecao := sessao.DB("empresa").C("cliente") //coleção cliente
  //Criando clientes genéricos
  cliente1 := Cliente{
    Nome: "Alex Silva",
    Email: "alex.s@email.com",
    Nasc: "23/05/2001",
    Doc: "232.233.222-09",
    Tipo: "fisica",
    End: "Rua aparecida, 100, Sao Paulo"}

  cliente2 := Cliente{
    Nome: "Amanda Texeira",
    Email: "am.t@email.com",
    Nasc: "13/06/2003",
    Doc: "232.234.122-06",
    Tipo: "fisica",
    End: "Rua jd. Amelia, 403, Sao Paulo"}

  cliente3 := Cliente{
    Nome: "Atkins",
    Email: "atkins.p@email.com",
    Nasc: "21/02/1965",
    Doc: "111.555.232-04",
    Tipo: "juridica",
    End: "Rua Jac, 89, Sao Paulo"}

  colecao.Insert(cliente1)
  colecao.Insert(cliente2)
  colecao.Insert(cliente3)
  sessao.Close()
  fmt.Println("Clientes inseridos com sucesso!")
}
