# Desafio 1

1. Criar um Dockerfile chamado Dockerfile.debian de acordo com as especificações abaixo:
    - base em uma distribuição debian;
    - instalar o pacote curl e suas dependências;
    - definir o diretório de trabalho para /usr/share/docker;
    - definir a variável de ambiente DEFAULT_PATH com o valor mypath;
    - baixar a última release da api https://github.com/leocomelli/simple-api;
    - iniciar a api automaticamente;
 
2. Com base no Dockerfile criado acima, adicionar...
    - um script de entrypoint que irá imprimir o hostname do container e, em seguida, iniciar a api;
 
3. Criar um Dockerfile chamado Dockerfile.alpine baseado nos dois itens anteriores (o gerenciador de pacotes é o apk). Analise qual a principal diferença.
 
4. Enviar as duas imagens para o ecr no repositório clarobr/simple-api, utilizaremos a seguinte convenção de nomes:

    ```
        <sobrenome-distribuição>, por exemplo:
            - clarobr/simple-api:comelli-debian
            - clarobr/simple-api:comelli-alpine
    ```
 
    O endereço do nosso registry é 719904057057.dkr.ecr.sa-east-1.amazonaws.com           
5. Fazer (e guardar) os comandos para...
    - construir as duas imagens
    - rodar container
    - rodar container (comando anterior +) expondo a porta
    - rodar container (comando anterior +) redefinindo a variável DEFAULT_PATH
    - rodar container (comando anterior +) redefinindo o diretório de trabalho
    - rodar container expondo a porta e passando o argumento anotherpath
    - rodar container em modo interativo ignorando o entrypoint
    - "tagear" e fazer push das duas imagens para o ecr

## Comandos

- construir as duas imagens

```
docker build -t clarobr/simple-api:comelli-debian -f Dockerfile.debian .
docker build -t clarobr/simple-api:comelli-alpine -f Dockerfile.alpine .
```

- rodar container

```
docker run -d clarobr/simple-api:comelli-debian
```

- rodar container (comando anterior +) expondo a porta

```
docker run -d -p 8081:8081 clarobr/simple-api:comelli-debian

curl http://localhost:8081/mypath
```    

- rodar container (comando anterior +) redefinindo a variável DEFAULT_PATH

```
docker run -d -p 8081:8081 --env DEFAULT_PATH=xxx clarobr/simple-api:comelli-debian

curl http://localhost:8081/xxx
```        

- rodar container (comando anterior +) redefinindo o diretório de trabalho

```
docker run -d -p 8081:8081 --env DEFAULT_PATH=xxx --workdir /tmp clarobr/simple-api:comelli-debian

curl http://localhost:8081/xxx
```            

- rodar container expondo a porta e passando o argumento anotherpath

```
docker run -d -p 8081:8081 clarobr/simple-api:comelli-debian anotherpath

curl http://localhost:8081/anotherpath
``` 

- rodar container em modo interativo ignorando o entrypoint

```
docker run -ti --entrypoint bash clarobr/simple-api:comelli-debian
```    

- "tagear" e fazer push das duas imagens para o ecr

```
$(aws ecr get-login --no-include-email --region sa-east-1)

docker tag clarobr/simple-api:comelli-debian 719904057057.dkr.ecr.sa-east-1.amazonaws.com/clarobr/simple-api:comelli-debian
docker tag clarobr/simple-api:comelli-alpine 719904057057.dkr.ecr.sa-east-1.amazonaws.com/clarobr/simple-api:comelli-alpine

docker push 719904057057.dkr.ecr.sa-east-1.amazonaws.com/clarobr/simple-api:comelli-debian
docker push 719904057057.dkr.ecr.sa-east-1.amazonaws.com/clarobr/simple-api:comelli-alpine
```        

