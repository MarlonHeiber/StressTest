# StressTest

Programa simples em GO para realizar testes de carga em um serviço web.
O programa realizar requests HTTP para a URL especificada e distribuir os requests de acordo com o nível de concorrência definido.

Para executar somente são preciso 3 parâmetros: 
1. **--url:** URL do serviço a ser testado.
2. **--requests:** Número total de requests.
3. **--concurrency:** Número de chamadas simultâneas.

Após o términio do teste de carga o programa exibe um relatório com os seguintes dados:
- Tempo total gasto na execução
- Quantidade total de requests realizados.
- Quantidade de requests com status HTTP 200.
- Distribuição de outros códigos de status HTTP (como 404, 500, etc.).

## Como Rodar:

Execute no terminal o comando para gerar o build da imagem docker: `docker build -t stresstest .`

Em seguida execute um teste da seguinte forma:

`docker run --rm stresstest --url=http://google.com --requests=10 --concurrency=2`