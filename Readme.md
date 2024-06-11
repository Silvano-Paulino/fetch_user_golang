# Como executar o projecto

## Compile o projeto para criar o executável e utilize o novo comando:

1. Navegue até o diretório do projeto:
``` 
cd /path/to/your/project
``` 

2. Compile o código:
``` 
go build -o usercli
``` 

3. Verifique o executável:

- Você deverá ver um arquivo usercli ou usercli.exe no diretório do projeto.


## Para buscar um usuário use os seguintes comandos: 

- Buscar um usuário específico:
``` 
./usercli fetchuser --id 1  
```

- Buscar todos os usuários:
``` 
./usercli fetchallusers
``` 

- Buscar um usuário com atraso simulado:
``` 
./usercli fetchuserwithdelay --id 4
``` 
