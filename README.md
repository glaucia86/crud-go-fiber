# Aplicação CRUD com Golang + Fiber Web Framework + PostgreSQL & Azure

[![bit-go.png](https://i.postimg.cc/s204FQ2b/bit-go.png)](https://postimg.cc/GBGyGpGj)

Desenvolvimento de uma aplicação CRUD com Golang usando o Fiber Web Framework persistindo os dados localmente no PostgreSQL. E ao final, iremos realizar o deploy dessa API no **[Azure App Service](https://azure.microsoft.com/services/app-service/web/?WT.mc_id=crudgolangfiber-github-gllemos)**!

## Recursos Utilizados 🚀

* **[Visual Studio Code](https://code.visualstudio.com/?WT.mc_id=crudgolangfiber-github-gllemos)**
* **[Golang](https://golang.org/doc/install)**
* **[Extension Golang - Visual Studio Code](https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go&WT.mc_id=crudgolangfiber-github-gllemos)**
* **[Fiber](https://gofiber.io/)**
* **[Azure App Service Web Apps](https://docs.microsoft.com/azure/app-service/?WT.mc_id=crudgolangfiber-github-gllemos)**

## Executando Aplicação Localmente 🔥

1. Se você estiver usando a extensão do Golang no Windows, pode ser que, encontre alguns problemas na execução da aplicação. Para isso, execute o comando dentro da pasta raiz `api` o seguinte comando:

```bash
> gofmt -w .
```

> esse comando precisa, por enquanto, ser executado, devido a extensão do Golang do Visual Code, que não está realizando a formatação dos arquivos de maneira devida no Windows. (somente quando alteramos qualquer linha de código nos arquivos).

2. Agora, basta entrar na pasta do projeto `api` e executar o seguinte comando:

```bash
> go run server.go
```

3. Abre o postman e digite a URL da API: (GET) `localhost:3000/api`. Se aparecer a seguinte mensagem:

```bash
> Sejam bem-vindos(as) a API Golang + Fiber + PostGreSQL + Azure!
```

É porque a API está executando corretamente! 

## Links & Recursos Importantes ❗️

Como dito durante as lives e vídeos, sempre estaremos citando links e recursos importantes que direcionarão para: documentações, cursos gratuitos, livros e conteúdos relacionados a Golang. Abaixo, segue uma lista desses recursos, que são considerados leituras, cursos ou livros recomendados:

- ✅ **[Documentação Fiber Web Framework](https://docs.gofiber.io/)**
- ✅ **[Documentação Oficial - Golang](http://www.golangbr.org/doc/)**
- ✅ **[Azure for Golang Developers](https://docs.microsoft.com/azure/developer/go/?WT.mc_id=golangstudies-github-gllemos)**
- ✅ **[Aprenda Go com Testes](https://larien.gitbook.io/aprenda-go-com-testes/)**

## Tenho Dúvidas... O que Faço?! ❓

Caso tenham dúvidas aos códigos desenvolvidos durante a série de artigos, sintam-se a vontade em abrir uma **[ISSUE AQUI](https://github.com/glaucia86/crud-go-fiber/issues)**. Assim que possível, estarei respondendo as todas as dúvidas que tiverem!


