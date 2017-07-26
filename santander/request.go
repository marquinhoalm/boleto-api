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

func getResponseSantander() string {
	return registerSantanderResponse
}

func getRequestSantander() string {
	return registerBoleto
}
