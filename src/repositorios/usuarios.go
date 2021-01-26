package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

//Usuarios representa um repositorio de usuarios
type Usuarios struct {
	db *sql.DB
}

//NovoRepositorioDeUsuarios cria novo usuario
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

//Criar que insere de fato os dados do banco
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	var lastInsertID int
	sql := "INSERT INTO public.usuarios(nome, email, senha) VALUES($1,$2,$3) returning id;"

	erro := repositorio.db.QueryRow(sql, usuario.Nome, usuario.Email, usuario.Senha).Scan(&lastInsertID)
	if erro != nil {
		return 0, erro
	}

	return uint64(lastInsertID), nil
}

//Buscar tras os dados de um usuario pelo filtro informado
func (repositorio Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)

	sql := "SELECT id, nome, email, criadoem FROM public.usuarios WHERE nome LIKE $1 OR nick LIKE $2"

	linhas, erro := repositorio.db.Query(sql, nomeOuNick, nomeOuNick)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

//BuscarPorID tras apenas um usuario
func (repositorio Usuarios) BuscarPorID(ID int64) (modelos.Usuario, error) {
	sql := "SELECT id, nome, nick, email, criadoem FROM public.usuarios WHERE id = $1"

	linhas, erro := repositorio.db.Query(sql, ID)

	if erro != nil {
		return modelos.Usuario{}, erro
	}

	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

//Atualizar tras apenas um usuario
func (repositorio Usuarios) Atualizar(ID uint64, usuario modelos.Usuario) (uint64, error) {
	var lastUpdateID uint64
	sql := "UPDATE public.usuarios SET nome = $1, email = $3 WHERE ID = $4 returning id;"

	erro := repositorio.db.QueryRow(sql, usuario.Nome, usuario.Email, ID).Scan(&lastUpdateID)

	if erro != nil {
		return 0, erro
	}

	return lastUpdateID, nil
}

//Deletar deleta o usuario do banco de dados
func (repositorio Usuarios) Deletar(ID uint64) error {
	sql := "DELETE FROM public.usuarios WHERE ID = $1"

	_, erro := repositorio.db.Query(sql, ID)

	if erro != nil {
		return erro
	}

	return nil
}

//BuscarPorEmail busca no banco de dados
func (repositorio Usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	sql := "SELECT id, senha FROM public.usuarios WHERE email = $1"

	linhas, erro := repositorio.db.Query(sql, email)

	if erro != nil {
		return modelos.Usuario{}, erro
	}

	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Senha,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

//BuscarSenha busca a senha do user pela ID
func (repositorio Usuarios) BuscarSenha(ID uint64) (string, error) {
	sql := "SELECT senha FROM public.usuarios WHERE ID = $1"

	linhas, erro := repositorio.db.Query(sql, ID)

	if erro != nil {
		return "", erro
	}

	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.Senha,
		); erro != nil {
			return "", erro
		}
	}

	return usuario.Senha, nil
}

//AtualizarSenha altera a senha do um usuario
func (repositorio Usuarios) AtualizarSenha(ID uint64, senhaComHash string) (uint64, error) {
	var lastUpdateID uint64
	sql := "UPDATE public.usuarios SET senha = $1 WHERE ID = $2 returning id;"

	erro := repositorio.db.QueryRow(sql, senhaComHash, ID).Scan(&lastUpdateID)

	if erro != nil {
		return 0, erro
	}

	return lastUpdateID, nil
}
