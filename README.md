# Tools API

Esta API foi desenvolvida para o desafio backend da BossaBox. O desafio original pode ser encontrado [aqui](https://bossabox.notion.site/Back-end-0b2c45f1a00e4a849eefe3b1d57f23c6)

## Sumário
- [Tools API](#tools-api)
  - [Sumário](#sumário)
  - [Tecnologias Usadas](#tecnologias-usadas)
  - [Instalação e Execução](#instalação-e-execução)
    - [Instalação](#instalação)
    - [Execução](#execução)

<a id="tecnologias-usadas"></a>

## Tecnologias Usadas

As tecnologias usadas e os motivos de as escolherem neste projetos são:
- Linguagem: Go
  - Afinidade. 
- Web Framework: Gin
  - Simplificar criação e manutenção de servidor e rotas
- ORM: GORM
  - Pelo projeto utilizar ações mais simples no banco de dados(CRUD).
- Banco de Dados: Postgres 
  - Afinidade
- Descrição de API: API Blueprint
  - Senso de aventura. Tecnologia totalmente nova.

<a id="instalacao-e-execucao"></a>

## Instalação e Execução

Para esse projeto, vale destacar que **é necessário ter instalado na máquina**:
- Go (>1.21)
- Postgres

### Instalação

Para começar, crie um banco de dados no postgres usando as queries do projeto
- [queries](/db/db.sql)

### Execução

Para executar esse projeto, primeiro entre no aquivo de conexão com o banco e modifique a string de conexão(StringConnection) de acordo com o seu dados. 
- [arquivo de conexao](/db/conn.go)

Agora, execute o projeto rodando o comando: 

```bash
go run cmd/main.go
```

Pronto. 😄


