# travel-cost-api-go

Travel-cost-api é uma api que se conecta com mongodb com teste de integração.
Os dados utilizados nesse material de estudo foram retirados do [portal da transparência do Governo Brasileiro](http://www.portaltransparencia.gov.br/viagens/consulta?ordenarPor=de&direcao=desc)

### 📋 Pré-requisitos

Para rodar esta aplicação você precisará ter o docker instalado na sua máquina e um container com uma imagem mongodb iniciado

* Para instalar uma imagem docker basta seguir os seguintes passos:

```docker run -p "27017:27017" mongo:latest```

```docker container ps -a```

* Salve o containerID e cole no arquivo ./scripts/run_database.sh

* No diretório scripts rode o comando ```./runner.sh```

* 



### Installation

Use the package manager [pip](https://pip.pypa.io/en/stable/) to install foobar.

```bash
pip install foobar
```

## Usage

```python
import foobar

foobar.pluralize('word') # returns 'words'
foobar.pluralize('goose') # returns 'geese'
foobar.singularize('phenomena') # returns 'phenomenon'
```


## Licença
[MIT](https://choosealicense.com/licenses/mit/)