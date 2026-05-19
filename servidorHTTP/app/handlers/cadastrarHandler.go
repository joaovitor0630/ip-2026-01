package handlers

import (
	"net/http"
	"servidorHTTP/app/utils"
)

// CadastrarHandler é responsável por processar os dados do formulário de cadastro de aluno
func CadastrarHandler(response http.ResponseWriter, request *http.Request) {
	// Verifica se o método da requisição é POST
	if request.Method != http.MethodPost {
		http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	// Obtém os valores enviados pelo formulário
	nome := request.FormValue("nome")
	email := request.FormValue("email")
	curso := request.FormValue("curso")
	matricula := request.FormValue("matricula")

	// Insere os dados do aluno no banco de dados
	err := utils.InserirAluno(nome, email, curso, matricula)
	if err != nil {
		http.Error(response, "Erro ao cadastrar o aluno no banco de dados", http.StatusInternalServerError)
		return
	}

	// Redireciona para a página inicial após o sucesso
	http.Redirect(response, request, "/index.html", http.StatusSeeOther)
}
