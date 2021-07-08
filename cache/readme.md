# Image Cache

Para criar uma cache dos downloads de módulos do go e ganhar tempo no processo de build de imagens durante a etapa 
de programação, olhe o texto da saída padrão durante o processo de build da imagem, sem cache. 

Ele estará cheio de linhas como esta:

```
2021/07/08 10:33:43 go: downloading github.com/hashicorp/memberlist v0.2.4
```

Copie e cole o texto em algum arquivo qualquer, depois use a função "find and replace" no seu editor, com as 
expressões regular abaixo:  

find: 
```goregexp
^.*?downloading\s(.*?)\s(.*)$
```

replace:
```goregexp
RUN go get -u $1@$2
```

Ao final do processo, as linhas serão transformadas como abaixo:

```shell
RUN go get -u github.com/hashicorp/memberlist@v0.2.4
```

Use estas linhas para criar um `Dockerfile` e em seguida, criar uma imagem de cache.

No caso desse código `cache:latest`, onde está imagem é uma imagem `golang:1.16-alpine` com todos os módulos go 
necessários para o projeto já baixados.

Agora abra o `Dockerfile` do projeto e na primeira etapa da imagem, em vez de usar 
`FROM golang:1.16-alpine AS builder`, use `FROM cache:latest AS builder` e ganhe tempo.



`build constraints exclude all`