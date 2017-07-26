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
            <ticket>?</ticket>
            <!--Optional:-->
            <tpAmbiente>?</tpAmbiente>
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
                    <value>123456789</value>
                </entry>
                <entry>
                    <key>PAGADOR.TP-DOC</key>
                    <value>99</value>
                </entry>
                <entry>
                    <key>PAGADOR.NUM-DOC</key>
                    <value>999999999999999</value>
                </entry>
                <entry>
                    <key>PAGADOR.NOME</key>
                    <value>xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx</value>
                </entry>
                <entry>
                    <key>PAGADOR.ENDER</key>
                    <value>xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx</value>
                </entry>
                <entry>
                    <key>PAGADOR.BAIRRO</key>
                    <value>xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx</value>
                </entry>
                <entry>
                    <key>PAGADOR.CIDADE</key>
                    <value>xxxxxxxxxxxxxxxxxxxx</value>
                </entry>
                <entry>
                    <key>PAGADOR.UF</key>
                    <value>xx</value>
                </entry>
                <entry>
                    <key>PAGADOR.CEP</key>
                    <value>99999999</value>
                </entry>
                <entry>
                    <key>TITULO.NOSSO-NUMERO</key>
                    <value>9999999999999</value>
                </entry>
                <entry>
                    <key>TITULO.SEU-NUMERO </key>
                    <value>xxxxxxxxxxxxxxx</value>
                </entry>
                <entry>
                    <key>TITULO.DT-VENCTO</key>
                    <value>ddmmaaaa</value>
                </entry>
                <entry>
                    <key>TITULO.DT-EMISSAO</key>
                    <value>ddmmaaaa</value>
                </entry>
                <entry>
                    <key>TITULO.ESPECIE</key><
                    value>xx</value>
                </entry>
                <entry>
                    <key>TITULO.VL-NOMINAL</key>
                    <value>999999999999999</value>
                </entry>
                <entry>
                    <key>TITULO.PC-MULTA</key>
                    <value>99999</value>
                </entry>
                <entry>
                    <key>TITULO.QT-DIAS-MULTA</key>
                    <value>99</value>
                </entry>
                <entry>
                    <key>TITULO.PC-JURO</key>
                    <value>99999</value>
                </entry>
                <entry>
                    <key>TITULO.TP-DESC</key>
                    <value>9</value>
                </entry>
                <entry>
                    <key>TITULO.VL-DESC</key>
                    <value>999999999999999</value>
                </entry>
                <entry>
                    <key>TITULO.DT-LIMI-DESC</key>
                    <value>ddmmaaaa</value>
                </entry>
                <entry>
                    <key>TITULO.VL-ABATIMENTO</key>
                    <value>999999999999999</value>
                </entry>
                <entry>
                    <key>TITULO.TP-PROTESTO</key>
                    <value>9</value>
                </entry>
                <entry>
                    <key>TITULO.QT-DIAS-PROTESTO</key>
                    <value>9</value>
                </entry>
                <entry>
                    <key>TITULO.QT-DIAS-BAIXA</key>
                    <value>9</value>
                </entry>
                <entry>
                    <key>MENSAGEM</key>
                    <value>xxxxxxxxxxxxxxxxxxxxxxxxxxx</value>
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
                <retCode>0</retCode>
                <ticket>cfslcN5/EJuS3WSZIMnUp6P2pllnSra78ABGSocUKwpZd2TmKoKknBIWVjALtRC9bfa8CoKU7DBKD8dMhYWyv7i+VSGqnjGq8Lg99U1EzdrItIALgPnFm6LpsIFThCRZ</ticket>
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
