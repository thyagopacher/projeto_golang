package repositories

import "projeto_go/internal/models"

type UsuarioRepository struct {
	data []models.Usuario
}

// Construtor
func NewUsuarioRepository() *UsuarioRepository {
	return &UsuarioRepository{
		data: []models.Usuario{
			{ID: 1, Nome: "Thyago", Email: "thyago@email.com"},
			{ID: 2, Nome: "Maria", Email: "maria@email.com"},
		},
	}
}

// Listar todos
func (r *UsuarioRepository) GetAll() []models.Usuario {
	return r.data
}

// Buscar por ID
func (r *UsuarioRepository) GetByID(id int) *models.Usuario {
	for _, u := range r.data {
		if u.ID == id {
			return &u
		}
	}
	return nil
}

// Criar
func (r *UsuarioRepository) Create(usuario models.Usuario) models.Usuario {
	usuario.ID = len(r.data) + 1
	r.data = append(r.data, usuario)
	return usuario
}

// Atualizar
func (r *UsuarioRepository) Update(id int, usuario models.Usuario) *models.Usuario {
	for i, u := range r.data {
		if u.ID == id {
			usuario.ID = id
			r.data[i] = usuario
			return &usuario
		}
	}
	return nil
}

// Deletar
func (r *UsuarioRepository) Delete(id int) bool {
	for i, u := range r.data {
		if u.ID == id {
			r.data = append(r.data[:i], r.data[i+1:]...)
			return true
		}
	}
	return false
}