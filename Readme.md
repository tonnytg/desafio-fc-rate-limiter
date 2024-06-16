# Rate Limit

This code use Rate Limit with Tokens refiled each 1 second

## About

Rate Limit é uma forma de controlar as requisições de rede. Você consegue adotar várias estratégias para usar essa ferramenta conforme sua necessidade.<br />
Ela vem com o objetivo para impdir ataques de força bruta ou até evitar Web Scraping que pode causar lentidão no funcionamento da aplicação.

Hoje já existem várias ferramentas que auxiliam nesse suporte como Istio para o caso do Kubernetes ou a própria Cloud Flare.

Existem várias estratégias para adoção do Rate Limit.

- Token Bucket
- Leaky Bucket
- Rate Limit by IP
- Rate Limit by User
- Rate Limit by Request Type
- etc...

Cada estratégia tem seus prós e contras, mas o foco é atuar como Middlare na requisição catalogar e classificar a origem para contabilizar o tempo permitido.<br />
Caso violado as regras, o bloqueio da requisição deve acontecer, protegendo assim a aplicação por trás do Rate Limit.


*Fontes:*

- https://www.cloudflare.com/learning/bots/what-is-rate-limiting/
- https://istio.io/v1.11/docs/tasks/policy-enforcement/rate-limit/
- https://aws.amazon.com/blogs/architecture/rate-limiting-strategies-for-serverless-applications/
- https://cloud.google.com/architecture/infra-reliability-guide/traffic-load?hl=pt-br

### How this code works

Esse projeto tem uma struct que vai ajudara a configurar a quantidade de tokens.

    type RateLimiter struct {
        client            *redis.Client
        rateLimitIP       int
        rateLimitToken    int
        blockTime         time.Duration
        refillInterval    time.Duration
        tokensPerRefill   int
        maxTokensPerIP    int
        maxTokensPerToken int
    }


Utilizando o Gin todas as requisições são passadas para um Middlware

	r := gin.Default()
	r.Use(middleware.RateLimiterMiddleware(rateLimiter))


A liberação irá ser avaliada ou por Token ou por IP conforme arquivo `middleware/ratelimniter.go`

		if token != "" {
			allowed, err = rl.AllowToken(c.Request.Context(), token)
		} else {
			allowed, err = rl.AllowIP(c.Request.Context(), ip)
		}


De tempo em tempo será reposta a quantidade de Tokens conforme configurado o `.env`

	// If the current tokens are less than the limit, refill tokens
	if tokens < maxTokens {
		pipe.IncrBy(ctx, key, int64(rl.tokensPerRefill))
		tokens += rl.tokensPerRefill
		if tokens > maxTokens {
			tokens = maxTokens
		}
	}


## How to use

Makefile can help to up environment or run `docker-compose up --build -d`

    `make`

Run tests with `make test`


### Customization

You can change values in `.env`

        REDIS_ADDR=localhost:6379
        REDIS_PASSWORD=
        REDIS_DB=0
        RATE_LIMIT_IP=5
        RATE_LIMIT_TOKEN=4
        BLOCK_TIME=300
        REFILL_INTERVAL=1
        TOKENS_PER_REFILL=1
        MAX_TOKENS_PER_IP=5
        MAX_TOKENS_PER_TOKEN=10


## Test

To test local command:

    for i in `seq 1 10`;do curl -X GET http://localhost:8080; sleep 3; done

## Evidences

![test-rate-limit-gin-response.png](test-rate-limit-gin-response.png)