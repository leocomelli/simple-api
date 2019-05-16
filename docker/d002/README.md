# Desafio 2

1. Criar um Dockerfile chamado Dockerfile.multistage de acordo com as especificações abaixo:
    - stage 1
        - base em uma distribuição debian;
        - "instalar" o go 1.12.5 via tar.gz;
        - clonar https://github.com/leocomelli/simple-api;
        - compilar o simple-api;
    - stage 2
        - copiar o binário gerado no stage 1
        - iniciar automaticamente a api
 