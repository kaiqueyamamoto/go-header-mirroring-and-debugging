# Estágio de compilação
FROM golang:1.22-alpine AS builder

# Definir o diretório de trabalho dentro do container
WORKDIR /app

# Copiar os arquivos do módulo Go e baixar as dependências
COPY go.mod ./
#COPY go.sum ./
RUN go mod download

# Copiar o código fonte para o diretório de trabalho
COPY *.go ./

# Compilar o aplicativo para um binário estático
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .

# Estágio de execução
FROM alpine:latest  

# Definir o diretório de trabalho
WORKDIR /root/

# Copiar o binário do estágio de compilação para o estágio de execução
COPY --from=builder /app/myapp .

# Expor a porta 2020
EXPOSE 2020

# Comando para executar o aplicativo
CMD ["./myapp"]
