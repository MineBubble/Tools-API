# Requisitos

### 0 A API deve responder na porta **3000**

### 1 Deve haver uma rota para listar todas as ferramentas cadastradas
> GET /tools

### 2 Deve ser possível filtrar ferramentas utilizando uma busca por tag
> /tools?tag='tag'

### 3 Deve haver uma rota para cadastrar uma nova ferramenta
> POST /tools

> Content-Type: application/json

```json
    {
    "title": "hotel",
    "link": "https://github.com/typicode/hotel",
    "description": "Local app manager. Start apps within your browser, developer tool with local .localhost domain and https out of the box.",
    "tags":["node", "organizing", "webapps", "domain", "developer", "https", "proxy"]
}
```
Resposta: 
> Status: 201 Created

> Content-Type: application/json

```json
{
    "title": "hotel",
    "link": "https://github.com/typicode/hotel",
    "description": "Local app manager. Start apps within your browser, developer tool with local .localhost domain and https out of the box.",
    "tags":["node", "organizing", "webapps", "domain", "developer", "https", "proxy"],
    "id":5
}
```

### 4 O usuário deve poder remover uma ferramenta por ID

> DELETE /tools/:id

Resposta: 

> Status: 200 OK
