package handlers

import (
	"net/http"
	"servidorHTTP/app/utils"
	"text/template"
)

// ConsultarHandler é responsável por buscar e exibir os dados de um aluno
func ConsultarHandler(response http.ResponseWriter, request *http.Request) {
	// Verifica se o método da requisição é POST
	if request.Method != http.MethodPost {
		http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	// Obtém a matrícula enviada pelo formulário
	matricula := request.FormValue("matricula")

	// Busca o aluno no banco de dados pela matrícula
	aluno, err := utils.BuscarAlunoPorMatricula(matricula)
	if err != nil {
		http.Error(response, "Aluno não encontrado", http.StatusNotFound)
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
