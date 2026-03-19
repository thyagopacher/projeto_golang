package services

import (
	"errors"
	"fmt"
	"projeto_go/internal/models"
	"projeto_go/internal/repositories"
)

type ProdutoService struct {
	repo *repositories.ProdutoRepository
}

func NewProdutoService(repo *repositories.ProdutoRepository) *ProdutoService {
	return &ProdutoService{repo: repo}
}

/**
* GET /Produtos
*/
func (s *ProdutoService) GetProdutos() ([]models.Produto, error) {
	return s.repo.GetAll()
}

/**
* GET /Produtos/:id
*/
func (s *ProdutoService) GetByID(id int) (models.Produto, error) {
	Produto, err := s.repo.GetByID(id)
	
	if err != nil {
		return models.Produto{}, err  // propaga erro do repo (ex: conexão, query falha)
	}

	if Produto.ID == 0 {
		return models.Produto{}, errors.New("usuário não encontrado")
	}

	return Produto, nil
}

/**
* POST /Produtos
*/
func (s *ProdutoService) CreateProduto(input models.Produto) (models.Produto, error) {

	if input.Nome == "" {
		return models.Produto{}, errors.New("nome é obrigatório")
	}

	criado, err := s.repo.Create(input)
	if err != nil {
		// Você pode melhorar o erro com wrap (opcional, mas recomendado)
		return models.Produto{}, fmt.Errorf("falha ao criar usuário: %w", err)
	}

	return criado, nil
}

/**
* PUT /Produtos
*/
func (s *ProdutoService) UpdateProduto(id int, input models.Produto) (models.Produto, error) {

	if input.Nome == "" {
		return models.Produto{}, errors.New("nome é obrigatório")
	}

	criado, err := s.repo.Update(id, input)
	if err != nil {
		// Você pode melhorar o erro com wrap (opcional, mas recomendado)
		return models.Produto{}, fmt.Errorf("falha ao criar usuário: %w", err)
	}

	return criado, nil
}

/**
* DELETE /Produtos/:id
*/
func (s *ProdutoService) DeleteProduto(id int) (bool, error) {

	ok, err := s.repo.Delete(id)

	if err != nil {
		// Você pode melhorar o erro com wrap (opcional, mas recomendado)
		return false, fmt.Errorf("falha ao excluir usuário: %w", err)
	}
	if !ok {
		return false, errors.New("usuário não encontrado")
	}

	return true, nil
}