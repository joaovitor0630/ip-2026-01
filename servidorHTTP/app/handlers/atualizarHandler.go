package handlers

import (
	"fmt"
	"net/http"
	"servidorHTTP/app/utils"
	"text/template"
)

// AtualizarHandler é responsável por atualizar os dados de um aluno
func AtualizarHandler(response http.ResponseWriter, request *http.Request) {
	// Verifica se o método da requisição é POST
	if request.Method != http.MethodPost {
		http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	// Obtém os valores do formulário
	nome := request.FormValue("nome")
	email := request.FormValue("email")
	curso := request.FormValue("curso")
	novaMatricula := request.FormValue("novaMatricula")
	matriculaAtual := request.FormValue("matriculaAtual")

	// Cria um mapa para armazenar os campos a serem atualizados
	updates := make(map[string]string)

	if nome != "" {
		updates["nome"] = nome
	}
	if email != "" {
		updates["email"] = email
	}
	if curso != "" {
		updates["curso"] = curso
	}
	if novaMatricula != "" {
		updates["matricula"] = novaMatricula
	}

	// Verifica se há campos para atualizar
	if len(updates) == 0 {
		http.Error(response, "Nenhum campo para atualizar", http.StatusBadRequest)
		return
	}

	// Atualiza os campos informados no banco de dados
	err := utils.AtualizarAluno(matriculaAtual, updates)
	if err != nil {
		http.Error(response, "Erro ao atualizar os dados do aluno", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// Determina qual matrícula usar para buscar os dados atualizados
	var matriculaConsulta string
	if novaMatricula == "" {
		matriculaConsulta = matriculaAtual
	} else {
		matriculaConsulta = novaMatricula
	}

	// Busca os dados atualizados do aluno
	aluno, err := utils.BuscarAlunoPorMatricula(matriculaConsulta)
	if err != nil {
		http.Error(response, "Erro ao buscar informações do aluno", http.StatusInternalServerError)
		return
	}

	// Carrega o template de exibição do aluno
	tmpl, err := template.ParseFiles("static/aluno.html")
	if err != nil {
		http.Error(response, "Erro ao carregar o template", http.StatusInternalServerError)
		return
	}

	// Renderiza o template com os dados do aluno
	err = tmpl.Execute(response, aluno)
	if err != nil {
		http.Error(response, "Erro ao renderizar o template", http.StatusInternalServerError)
		return
	}
}
