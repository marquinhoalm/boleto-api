package santander

import (
	"errors"
	"strings"

	"github.com/PMoneda/flow"
	"github.com/mundipagg/boleto-api/config"
	"github.com/mundipagg/boleto-api/log"
	"github.com/mundipagg/boleto-api/models"
	"github.com/mundipagg/boleto-api/tmpl"
	"github.com/mundipagg/boleto-api/util"
	"github.com/mundipagg/boleto-api/validations"
)

type bankSantander struct {
	validate *models.Validator
	log      *log.Log
}

//New Create a new Santander Integration Instance
func New() bankSantander {
	b := bankSantander{
		validate: models.NewValidator(),
		log:      log.CreateLog(),
	}
	b.validate.Push(validations.ValidateAmount)
	b.validate.Push(validations.ValidateExpireDate)
	b.validate.Push(validations.ValidateBuyerDocumentNumber)
	b.validate.Push(validations.ValidateRecipientDocumentNumber)
	return b
}

//Log retorna a referencia do log
func (b bankSantander) Log() *log.Log {
	return b.log
}
func (b bankSantander) GetTicket(boleto *models.BoletoRequest) (string, error) {
	pipe := flow.NewFlow()
	url := config.Get().URLTicketSantander
	pipe.From("message://?source=inline", boleto, getRequestTicket(), tmpl.GetFuncMaps())
	pipe = pipe.To("logseq://?type=request&url="+url, b.log)
	pipe.To(url, map[string]string{"method": "POST", "insecureSkipVerify": "true"})
	pipe = pipe.To("logseq://?type=response&url="+url, b.log)
	ch := pipe.Choice()
	ch = ch.When(flow.Header("status").IsEqualTo("200"))
	ch = ch.To("transform://?format=xml", getTicketResponse(), `{{.returnCode}},{{.ticket}}`)
	ch = ch.Otherwise()
	ch = ch.To("logseq://?type=request&url="+url, b.log)
	ch = ch.To("set://?prop=body", errors.New("integration error"))
	switch t := pipe.GetBody().(type) {
	case string:
		items := pipe.GetBody().(string)
		parts := strings.Split(items, ",")
		returnCode, ticket := parts[0], parts[1]
		return ticket, checkError(returnCode)
	case error:
		return "", t
	}
	return "", nil
}

func (b bankSantander) RegisterBoleto(boleto *models.BoletoRequest) (models.BoletoResponse, error) {
	r := flow.NewFlow()
	serviceURL := config.Get().URLCiti
	from := getResponseSantander()
	to := getAPIResponseSantander()
	bod := r.From("message://?source=inline", boleto, getRequestSantander(), tmpl.GetFuncMaps())
	bod = bod.To("logseq://?type=request&url="+serviceURL, b.log)
	bod = bod.To(serviceURL, map[string]string{"method": "POST", "insecureSkipVerify": "true"})
	bod = bod.To("logseq://?type=response&url="+serviceURL, b.log)
	ch := bod.Choice()
	ch = ch.When(flow.Header("status").IsEqualTo("200"))
	ch = ch.To("transform://?format=xml", from, to, tmpl.GetFuncMaps())
	ch = ch.Otherwise()
	ch = ch.To("logseq://?type=response&url="+serviceURL, b.log).To("apierro://")
	switch t := bod.GetBody().(type) {
	case string:
		response := util.ParseJSON(t, new(models.BoletoResponse)).(*models.BoletoResponse)
		return *response, nil
	case models.BoletoResponse:
		return t, nil
	}
	return models.BoletoResponse{}, models.NewInternalServerError("Erro interno", "MP500")
}
func (b bankSantander) ProcessBoleto(boleto *models.BoletoRequest) (models.BoletoResponse, error) {
	errs := b.ValidateBoleto(boleto)
	if len(errs) > 0 {
		return models.BoletoResponse{Errors: errs}, nil
	}
	if ticket, err := b.GetTicket(boleto); err != nil {
		return models.BoletoResponse{Errors: errs}, err
	} else {
		boleto.Authentication.AuthorizationToken = ticket
	}
	return b.RegisterBoleto(boleto)
}

func (b bankSantander) ValidateBoleto(boleto *models.BoletoRequest) models.Errors {
	return models.Errors(b.validate.Assert(boleto))
}

//GetBankNumber retorna o codigo do banco
func (b bankSantander) GetBankNumber() models.BankNumber {
	return models.Citibank
}
