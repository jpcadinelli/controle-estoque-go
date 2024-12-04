package enum

var (
	ListaPermissoes = []string{
		PermissaoSistemaAdmin,

		PermissaoPermissaoCriar,
		PermissaoPermissaoVisualizar,
		PermissaoPermissaoListar,
		PermissaoPermissaoDropdown,
		PermissaoPermissaoAtualizar,
		PermissaoPermissaoDeletar,

		PermissaoUsuarioCriar,
		PermissaoUsuarioVisualizar,
		PermissaoUsuarioListar,
		PermissaoUsuarioDropdown,
		PermissaoUsuarioAtualizar,
		PermissaoUsuarioDeletar,

		PermissaoUsuarioAtribuirPermissao,
		PermissaoUsuarioRemoverPermissao,

		PermissaoProdutoCriar,
		PermissaoProdutoVisualizar,
		PermissaoProdutoListar,
		PermissaoProdutoDropdown,
		PermissaoProdutoAtualizar,
		PermissaoProdutoDeletar,

		PermissaoEstoqueCriar,
		PermissaoEstoqueAtualizar,

		PermissaoEnderecoCriar,
		PermissaoEnderecoAtualizar,
	}
)

const (
	PermissaoSistemaAdmin = "SISTEMA_ADMIN"

	PermissaoPermissaoCriar      = "PERMISSAO_CRIAR"
	PermissaoPermissaoVisualizar = "PERMISSAO_VISUALIZAR"
	PermissaoPermissaoListar     = "PERMISSAO_LISTAR"
	PermissaoPermissaoDropdown   = "PERMISSAO_DROPDOWN"
	PermissaoPermissaoAtualizar  = "PERMISSAO_ATUALIZAR"
	PermissaoPermissaoDeletar    = "PERMISSAO_DELETAR"

	PermissaoUsuarioCriar      = "USUARIO_CRIAR"
	PermissaoUsuarioVisualizar = "USUARIO_VISUALIZAR"
	PermissaoUsuarioListar     = "USUARIO_LISTAR"
	PermissaoUsuarioDropdown   = "USUARIO_DROPDOWN"
	PermissaoUsuarioAtualizar  = "USUARIO_ATUALIZAR"
	PermissaoUsuarioDeletar    = "USUARIO_DELETAR"

	PermissaoUsuarioAtribuirPermissao = "USUARIO_ATRIBUIR_PERMISSAO"
	PermissaoUsuarioRemoverPermissao  = "USUARIO_REMOVER_PERMISSAO"

	PermissaoProdutoCriar      = "PRODUTO_CRIAR"
	PermissaoProdutoVisualizar = "PRODUTO_VISUALIZAR"
	PermissaoProdutoListar     = "PRODUTO_LISTAR"
	PermissaoProdutoDropdown   = "PRODUTO_DROPDOWN"
	PermissaoProdutoAtualizar  = "PRODUTO_ATUALIZAR"
	PermissaoProdutoDeletar    = "PRODUTO_DELETAR"

	PermissaoEstoqueCriar     = "ESTOQUE_CRIAR"
	PermissaoEstoqueAtualizar = "ESTOQUE_ATUALIZAR"

	PermissaoEnderecoCriar     = "ENDERECO_CRIAR"
	PermissaoEnderecoAtualizar = "ENDERECO_ATUALIZAR"
)
