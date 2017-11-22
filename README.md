Gerenciamento de Clientes
=========================
Esta é uma aplicação feita para cadastro e consulta de clientes.

Requisitos pré-instalados
-------------------------

 * [Golang](https://golang.org/) e biblioteca [mgo](https://labix.org/mgo)
 * [MongoDB](https://www.mongodb.com/)
 * Um navegador

Execução
--------
Para inserir automaticamente alguns clientes genéricos (não é necessário para
o funcionamento do programa) no banco de dados digite:

    go run makedb.go

No diretório raiz do repositório, digite:

    go build main.go
    ./main

Abra um navegador e acesse a URL indicada.

Creditos
--------
Feito por Victor do Prado
Email: victor.prasa@gmail.com

Licensa
-------
Este software é licensiado pela MIT License. (Veja License)
