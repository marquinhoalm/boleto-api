package santander

const registerBoleto = `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:impl="http://impl.webservice.ymb.app.bsbr.altec.com/">
   <soapenv:Header/>
   <soapenv:Body>
      <impl:registraTitulo>
         <!--Optional:-->
         <dto>
            <!--Optional:-->
            <dtNsu>?</dtNsu>
            <!--Optional:-->
            <estacao>?</estacao>
            <!--Optional:-->
            <nsu>?</nsu>
            <!--Optional:-->
            <ticket>{{unscape .Authentication.AuthorizationToken}}</ticket>
            <!--Optional:-->
            <tpAmbiente>T</tpAmbiente>
         </dto>
      </impl:registraTitulo>
   </soapenv:Body>
</soapenv:Envelope>
`

const registerSantanderResponse = `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
   <soapenv:Body>
      <dlwmin:registraTituloResponse xmlns:dlwmin="http://impl.webservice.ymb.app.bsbr.altec.com/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
         <return xmlns:ns2="http://impl.webservice.ymb.app.bsbr.altec.com/">
            <codcede/>
            <convenio>
               <codBanco/>
               <codConv/>
            </convenio>
            <descricaoErro>{{errorMessage}}</descricaoErro>
            <dtNsu>?</dtNsu>
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
               <cdBarra>{{barcodeNumber}}</cdBarra>
               <dtEmissao/>
               <dtEntr/>
               <dtLimiDesc/>
               <dtVencto/>
               <especie/>
               <linDig>{{digitableLine}}</linDig>
               <mensagem/>
               <nossoNumero>{{ourNumber}}</nossoNumero>
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
            <tpAmbiente>?</tpAmbiente>
         </return>
      </dlwmin:registraTituloResponse>
   </soapenv:Body>
</soapenv:Envelope>
`

const requestTicket = `

## SOAPAction:registraTitulo
## Content-Type:text/xml

<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:impl="http://impl.webservice.dl.app.bsbr.altec.com/">
   <soapenv:Header/>
   <soapenv:Body>
      <impl:create>
         <!--Optional:-->
         <TicketRequest>
            <dados>
                <entry>
                    <key>CONVENIO.COD-BANCO</key>
                    <value>0033</value>
                </entry>
                <entry>
                    <key>CONVENIO.COD-CONVENIO</key>
                    <value>{{.Agreement.AgreementNumber}}</value>
                </entry>
                <entry>
                    <key>PAGADOR.TP-DOC</key>
                    {{if eq .Buyer.Document.Type "CPF"}}
					 	<value>01</value>
                    {{else}}
					 	<value>02</value>
					{{end}}                    
                </entry>
                <entry>
                    <key>PAGADOR.NUM-DOC</key>
                    <value>{{.Buyer.Document.Number}}</value>
                </entry>
                <entry>
                    <key>PAGADOR.NOME</key>
                    <value>{{.Buyer.Name}}</value>
                </entry>
                <entry>
                    <key>PAGADOR.ENDER</key>
                    <value>{{.Buyer.Address.Street}}</value>
                </entry>
                <entry>
                    <key>PAGADOR.BAIRRO</key>
                    <value>{{.Buyer.Address.District}}</value>
                </entry>
                <entry>
                    <key>PAGADOR.CIDADE</key>
                    <value>{{.Buyer.Address.City}}</value>
                </entry>
                <entry>
                    <key>PAGADOR.UF</key>
                    <value>{{.Buyer.Address.StateCode}}</value>
                </entry>
                <entry>
                    <key>PAGADOR.CEP</key>
                    <value>{{.Buyer.Address.ZipCode}}</value>
                </entry>
                <entry>
                    <key>TITULO.NOSSO-NUMERO</key>
                    <value>{{.Title.OurNumber}}</value>
                </entry>
                <entry>
                    <key>TITULO.SEU-NUMERO </key>
                    <value>{{.Title.DocumentNumber}}</value>
                </entry>
                <entry>
                    <key>TITULO.DT-VENCTO</key>
                    <value>{{brdate .Title.ExpireDateTime}}</value>
                </entry>
                <entry>
                    <key>TITULO.DT-EMISSAO</key>
                    <value>{{today | brdate}}</value>
                </entry>
                <entry>
                    <key>TITULO.ESPECIE</key>
                    <value>99</value>
                </entry>
                <entry>
                    <key>TITULO.TP-DESC</key>
                    <value>0</value>
                </entry>               
                <entry>
                    <key>TITULO.TP-PROTESTO</key>
                    <value>0</value>
                </entry>
                <entry>
                    <key>TITULO.QT-DIAS-PROTESTO</key>
                    <value>0</value>
                </entry>
                <entry>
                    <key>TITULO.QT-DIAS-BAIXA</key>
                    <value>0</value>
                </entry>
                <entry>
                    <key>TITULO.VL-NOMINAL</key>
                    <value>{{.Title.AmountInCents}}</value>
                </entry>                
                <entry>
                    <key>MENSAGEM</key>
                    <value>{{.Title.Instructions}}</value>
                </entry>
            </dados>
            <expiracao>100</expiracao>
            <sistema>YMB</sistema>
         </TicketRequest>
      </impl:create>
   </soapenv:Body>
</soapenv:Envelope>
`

const ticketReponse = `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
    <soapenv:Body> 
        <dlwmin:createResponse xmlns:dlwmin="http://impl.webservice.dl.app.bsbr.altec.com/">
            <TicketResponse>
                <retCode>{{returnCode}}</retCode>
                <ticket>{{ticket}}</ticket>
            </TicketResponse>
        </dlwmin:createResponse>
    </soapenv:Body>
</soapenv:Envelope>
`

func getResponseSantander() string {
	return registerSantanderResponse
}

func getRequestSantander() string {
	return registerBoleto
}

func getRequestTicket() string {
	return requestTicket
}

func getTicketResponse() string {
	return ticketReponse
}
