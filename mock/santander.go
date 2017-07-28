package mock

import (
	"net/http/httputil"

	"strings"

	"errors"

	gin "gopkg.in/gin-gonic/gin.v1"
)

/*

0 Ticket validado ok
1 Erro, dados de entrada inválidos
2 Erro interno de criptografia
3 Erro, Ticket já utilizado anteriormente
4 Erro, Ticket gerado para outro sistema
5 Erro, Ticket expirado
6 Erro interno (dados)
7 Erro interno (timestamp)
*/
func getTicket(c *gin.Context) {
	const tok = `
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"> 
	<soapenv:Body> 
	<dlwmin:createResponse xmlns:dlwmin="http://impl.webservice.dl.app.bsbr.altec.com/"> 
	<TicketResponse> 
	<retCode>0</retCode> 
	<ticket>cfslcN5/EJuS3WSZIMnUp6P2pllnSra78ABGSocUKwpZd2TmKoKknBIWVjALtRC9bfa8CoKU7DBKD8dMhYWyv7i+VSGqnjGq8Lg99U1EzdrItIALgPnFm6LpsIFThCRZ</ticket> 
	</TicketResponse> 
	</dlwmin:createResponse> 
	</soapenv:Body> 
	</soapenv:Envelope>
	`

	const tokErr = `
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"> 
	<soapenv:Body> 
	<dlwmin:createResponse xmlns:dlwmin="http://impl.webservice.dl.app.bsbr.altec.com/"> 
	<TicketResponse> 
	<retCode>3</retCode> 
	<ticket></ticket> 
	</TicketResponse> 
	</dlwmin:createResponse> 
	</soapenv:Body> 
	</soapenv:Envelope>
	`
	data, _ := httputil.DumpRequest(c.Request, true)
	str := string(data)
	if strings.Contains(str, "<value>3</value>") {
		c.Data(200, "text/json", []byte(tokErr))
	} else if strings.Contains(str, "<value>500</value>") {
		c.AbortWithError(500, errors.New("internal error"))
	} else {
		c.Data(200, "text/json", []byte(tok))
	}

}

func registerBoletoSantander(c *gin.Context) {
	const tok = `
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
   <soapenv:Body>
      <dlwmin:registraTituloResponse xmlns:dlwmin="http://impl.webservice.ymb.app.bsbr.altec.com/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
         <return xmlns:ns2="http://impl.webservice.ymb.app.bsbr.altec.com/">
            <codcede/>
            <convenio>
               <codBanco/>
               <codConv/>
            </convenio>
            <descricaoErro></descricaoErro>
            <dtNsu>05082017</dtNsu>
            <estacao>?</estacao>
            <nsu>?</nsu>
            <pagador>
               <bairro/>
               <cep/>
               <cidade/>
               <ender/>
               <nome/>
               <numDoc/>
               <tpDoc/>
               <uf/>
            </pagador>
            <situacao>20</situacao>
            <titulo>
               <aceito>N</aceito>
               <cdBarra/>
               <dtEmissao/>
               <dtEntr/>
               <dtLimiDesc/>
               <dtVencto/>
               <especie/>
               <linDig/>
               <mensagem/>
               <nossoNumero/>
               <pcJuro/>
               <pcMulta/>
               <qtDiasBaixa/>
               <qtDiasMulta/>
               <qtDiasProtesto/>
               <seuNumero/>
               <tpDesc/>
               <tpProtesto/>
               <vlAbatimento/>
               <vlDesc/>
               <vlNominal/>
            </titulo>
            <tpAmbiente>T</tpAmbiente>
         </return>
      </dlwmin:registraTituloResponse>
   </soapenv:Body>
</soapenv:Envelope>
	`
	c.Data(200, "text/json", []byte(tok))
}
