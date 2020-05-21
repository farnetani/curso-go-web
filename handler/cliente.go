package handler

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/jeffprestes/curso-go-web/lib/contx"
	"github.com/jeffprestes/curso-go-web/model"
	"github.com/jeffprestes/curso-go-web/repo"
)

func ListCliente(ctx *contx.Context) {
	clientes, err := repo.GetClientes()

	type Result struct {
		Data []model.Cliente
		Msg string
	}
	result := Result{
		Data : clientes,
		Msg : "ok",
	}

	if err != nil {
		log.Println("retornaParaListaClientes - Error: ", err.Error())
		ctx.NativeHTML(http.StatusInternalServerError, "erro")
		return
	}
	ctx.JSON(200, result)

	// ctx.Data["clientes"] = clientes
	// ctx.Resp.WriteHeader (http.StatusOK)

	// json.NewEncoder(ctx.Resp).Encode(clientes)
	// ctx.Resp.Write([]byte("teste"))
	// ctx.JSONWithoutEscape(200, clientes)
	// ctx.JSON(201, clientes)
	// ctx.XML(200, clientes)
}

//IndexCliente abre a pagina de gerenciamento de clientes
func IndexCliente(ctx *contx.Context) {
	clientes, err := repo.GetClientes()
	if err != nil {
		log.Println("retornaParaListaClientes - Error: ", err.Error())
		ctx.NativeHTML(http.StatusInternalServerError, "erro")
		return
	}
	ctx.Data["clientes"] = clientes
	ctx.NativeHTML(http.StatusOK, "clientes")
}

//AlteraCliente altera dados do cliente ou insere caso o ID não se encontrado na base de dados
func AlteraCliente(ctx *contx.Context, form model.Cliente) {
	_, err := repo.GetClientePeloID(form.ID)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println("AlteraCliente - repo.GetClientes - Error: ", err.Error())
			ctx.NativeHTML(http.StatusInternalServerError, "erro")
			return
		}
		err = nil
		_, err := repo.InsereCliente(form)
		if err != nil {
			log.Println("AlteraCliente - repo.InsereCliente - Error: ", err.Error())
			ctx.NativeHTML(http.StatusInternalServerError, "erro")
			return
		}
		ctx.Redirect("/")
		return
	}
	err = repo.AtualizaCliente(form)
	if err != nil {
		log.Println("AlteraCliente - repo.AtualizaCliente - Error: ", err.Error())
		ctx.NativeHTML(http.StatusInternalServerError, "erro")
		return
	}
	ctx.Redirect("/")
}
