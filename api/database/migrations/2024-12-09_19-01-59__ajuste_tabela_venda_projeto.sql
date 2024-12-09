ALTER TABLE venda_produto DROP COLUMN quantidade;

ALTER TABLE venda_produto ADD COLUMN quantidade INTEGER NOT NULL;

COMMENT ON COLUMN venda_produto.quantidade IS 'Campo para armazenar quantidade do produto nesta venda';