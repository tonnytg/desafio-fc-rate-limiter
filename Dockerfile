# Etapa de build
FROM golang:1.22.2-alpine as build
WORKDIR /app

# Copia os arquivos para o diretório de trabalho
COPY . .

# Compila o binário
RUN CGO_ENABLED=0 GOOS=linux go build -o rate-limit main.go

# Etapa final
FROM golang:1.22.2-alpine
WORKDIR /app

# Copia o binário compilado para a imagem final
COPY --from=build /app/rate-limit .

COPY .env .

# Certifica-se de que o binário tem permissões de execução
RUN ["chmod", "+x", "/app/rate-limit"]

# Define a porta exposta
EXPOSE 8080

# Define o ponto de entrada
ENTRYPOINT ["/app/rate-limit"]