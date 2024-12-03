CREATE TABLE produto (
    id UUID NOT NULL,
    nome TEXT NOT NULL,
    marca TEXT,
    quantidade INTEGER,
    unidade VARCHAR(25),
        CONSTRAINT pk_produto PRIMARY KEY (id)
);

COMMENT ON COLUMN produto.id IS 'Identificador único do produto';
COMMENT ON COLUMN produto.nome IS 'Nome do produto';
COMMENT ON COLUMN produto.marca IS 'Marca do produto';
COMMENT ON COLUMN produto.quantidade IS 'Quantidade do produto por unidade';
COMMENT ON COLUMN produto.unidade IS 'Unidade da quantidade do produto';

CREATE TABLE estoque (
    id UUID NOT NULL,
    id_produto UUID NOT NULL,
    quantidade INTEGER NOT NULL,
    custo DECIMAL(10,2) NOT NULL,
    data DATE NOT NULL,
        CONSTRAINT pk_estoque PRIMARY KEY (id),
        CONSTRAINT fk_estoque_produto FOREIGN KEY (id_produto) REFERENCES produto(id)
);

COMMENT ON COLUMN estoque.id IS 'Identificador único do estoque';
COMMENT ON COLUMN estoque.id_produto IS 'Identificador do produto adicionado ao estoque';
COMMENT ON COLUMN estoque.quantidade IS 'Quantidade de produto adicionada ao estoque';
COMMENT ON COLUMN estoque.custo IS 'Custo daquela quantidade de produto adicionada ao estoque';
COMMENT ON COLUMN estoque.data IS 'Data em que o produto foi adicionado ao estoque';

CREATE TABLE endereco (
    id UUID NOT NULL,
    endereco TEXT NOT NULL,
        CONSTRAINT pk_endereco PRIMARY KEY (id)
);

COMMENT ON COLUMN endereco.id IS 'Identificador único do endereço';
COMMENT ON COLUMN endereco.endereco IS 'Campo para armazenar o endereço';

CREATE TABLE cliente (
    id UUID NOT NULL,
    id_endereco_padrao UUID,
    nome TEXT NOT NULL,
    referencia TEXT,
    telefone VARCHAR(30),
    whatsapp VARCHAR(30),
    instagram VARCHAR(30),
        CONSTRAINT pk_cliente PRIMARY KEY (id),
        CONSTRAINT fk_cliente_endereco FOREIGN KEY (id_endereco_padrao) REFERENCES endereco(id)
);

COMMENT ON COLUMN cliente.id IS 'Identificador único do cliente';
COMMENT ON COLUMN cliente.id_endereco_padrao IS 'Identificador único do endereço padrão de entrega do cliente';
COMMENT ON COLUMN cliente.nome IS 'Campo para armazenar nome do cliente';
COMMENT ON COLUMN cliente.referencia IS 'Campo para armazenar referência do cliente';
COMMENT ON COLUMN cliente.telefone IS 'Campo para armazenar telefone do cliente';
COMMENT ON COLUMN cliente.whatsapp IS 'Campo para armazenar whatsapp do cliente';
COMMENT ON COLUMN cliente.instagram IS 'Campo para armazenar instagram do cliente';

CREATE TABLE venda (
    id UUID NOT NULL,
    id_cliente UUID NOT NULL,
    id_endereco UUID NOT NULL,
    custo DECIMAL(10,2) NOT NULL,
    entrega DECIMAL(10,2),
    valor DECIMAL(10,2) NOT NULL,
    pago BOOLEAN NOT NULL,
    data DATE NOT NULL,
        CONSTRAINT pk_venda PRIMARY KEY (id),
        CONSTRAINT fk_venda_cliente FOREIGN KEY (id_cliente) REFERENCES cliente(id),
        CONSTRAINT fk_venda_endereco FOREIGN KEY (id_endereco) REFERENCES endereco(id)
);

COMMENT ON COLUMN venda.id IS 'Identificador único da venda';
COMMENT ON COLUMN venda.id_cliente IS 'Identificador único do cliente para quem foi feita a venda';
COMMENT ON COLUMN venda.id_endereco IS 'Identificador único do endereço de entrega da venda';
COMMENT ON COLUMN venda.custo IS 'Campo para armazenar custo em reais (R$) da venda';
COMMENT ON COLUMN venda.entrega IS 'Campo para armazenar entrega em reais (R$) da venda';
COMMENT ON COLUMN venda.valor IS 'Campo para armazenar valor em reais (R$) da venda';
COMMENT ON COLUMN venda.data IS 'Campo para armazenar data em que a venda foi aberta';


CREATE TABLE venda_produto (
    id UUID NOT NULL,
    id_venda UUID NOT NULL,
    id_produto UUID NOT NULL,
    quantidade DECIMAL(10,2) NOT NULL,
        CONSTRAINT pk_venda_produto PRIMARY KEY (id),
        CONSTRAINT fk_venda_produto_venda FOREIGN KEY (id_venda) REFERENCES venda(id),
        CONSTRAINT fk_venda_produto_produto FOREIGN KEY (id_produto) REFERENCES produto(id)
);

COMMENT ON COLUMN venda_produto.id IS 'Identificador único da relação venda com o produto';
COMMENT ON COLUMN venda_produto.id_venda IS 'Identificador único do venda';
COMMENT ON COLUMN venda_produto.id_produto IS 'Identificador único do produto';
COMMENT ON COLUMN venda_produto.quantidade IS 'Campo para armazenar quantidade do produto nesta venda';