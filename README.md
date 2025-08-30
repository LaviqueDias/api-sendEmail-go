# 📧 API SendEmail (Go + Gin + SMTP)

Este projeto é uma **API em Go** que recebe dados de contato via JSON e envia e-mails usando qualquer servidor **SMTP**.  
A arquitetura segue **controller → service → model**, com validação de payloads, respostas padronizadas e envio de mensagens HTML.

---

## 🚀 Tecnologias
- [Go 1.22+](https://go.dev/)
- [Gin](https://github.com/gin-gonic/gin) — framework web
- [go-mail v2](https://github.com/go-mail/mail) — envio SMTP
- [Servidor SMTP](https://www.ibm.com/docs/pt-br/i/7.6.0?topic=information-smtp) (Zoho, Gmail, Outlook, etc.)

---

## ⚙️ Configuração

### Variáveis de ambiente
Crie um arquivo `.env` na raiz do projeto com suas credenciais SMTP:

```env
USER=seu-email@dominio.com
PASS=sua-app-password
SMTP_HOST=smtp.seuprovedor.com
SMTP_PORT=587
MAIL_FROM_NAME=Seu Nome ou Empresa
```

> ⚠️ O PASS precisa ser um App Password em vez da senha normal. A maioria dos provedores exige que você ative 2FA/MFA pra obter esse PASS.

## 📦 Instalação

Clone o reopsitório e instale as dependências:

```bash
git clone https://github.com/seu-usuario/api-sendEmail-go.git
cd api-sendEmail-go

go mod tidy
```

## ▶️ Executando 

```bash
go run main.go
```

> A API estará disponível em: *```http://localhost:8080```*

## 📡 Endpoints

```POST /api/email/send```

**Payload (Request)**

```json
{
    "name":"Name Test",
    "email":"emailtest@gmail.com",
    "message":"Message Test"
}
```
**Resposta de Sucesso (200 OK)**

```json
{
    "emailBody": "\n        <div style=\"font-family:Arial,Helvetica,sans-serif;font-size:14px;color:#222\">\n            <h2>Email sent via API</h2>\n            <p><b>Nome:</b> Name Test</p>\n            <p><b>Email:</b> emailtest@gmail.com</p>\n            <p><b>Mensagem:</b></p>\n            <p style=\"white-space:pre-wrap\">Message Test</p>\n            <br>\n            <h4>API repo: github.com/LaviqueDias/api-sendEmail-go</h4>\n        </div>",
    "message": "Email sent successfully"
}
```

**Resposta de Erro (exemplos)**
- ```400 Bad Request``` — payload inválido (campos obrigatórios faltando ou só espaços em branco)
- ```500 Internal Server Error``` — falha de autenticação ou envio via SMTP

## 📂 Estrutura Simplificada

```bash
api-sendEmail-go/
├── main.go                   # entrypoint
├── internal/
│   ├── configuration/        # validação, erros padrão
│   └── email/
│       ├── controller/       # recebe requests HTTP
│       ├── service/          # lógica de envio SMTP
│       └── model/            # structs (SMTPSender, etc.)
├── go.mod / go.sum
├── .env                      # variáveis de ambiente (não versionar!)

```

## 👨‍💻 Autor
Projeto desenvolvido por [Lavique Dias](github.com/LaviqueDias) 🚀