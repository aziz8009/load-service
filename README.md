# Loan Service API

API ini adalah backend service untuk Loan Engine menggunakan Golang dengan arsitektur Clean Architecture dan MySQL sebagai database.

---

## Fitur

- Membuat loan baru (state: proposed)
- Approve loan (state: approved)
- Melakukan investasi pada loan (state: invested)
- Melakukan pencairan dana/disbursement (state: disbursed)
- Monitoring daftar loan dan detail loan
- Validasi alur state yang hanya maju
- Upload file PDF/JPEG untuk approval dan disbursement

---

## Teknologi

- Golang (Go)
- Gin Framework (HTTP REST API)
- GORM (ORM untuk MySQL)
- MySQL
- Clean Architecture (Layered: Entity, Repository, Usecase, Handler)
- Dependency Injection manual
- UUID untuk ID (optional)

---

## Struktur Folder

loan-service/
│
├── cmd/
│   └── main.go
│
├── config/
│   └── config.go
│
├── internal/
│   ├── domain/
│   │   ├── loan.go
│   │   └── repository.go
│   │
│   ├── usecase/
│   │   └── loan_usecase.go
│   │
│   ├── handler/
│   │   └── loan_handler.go
│   │
│   ├── repository/
│   │   └── loan_repository.go
│   │
│   └── model/
│       └── loan_model.go
│
├── pkg/
│   └── database/
│       └── mysql.go
│
├── .env
├── go.mod
└── go.sum


---

## Setup dan Instalasi

### 1. Clone repository

```bash
git clone https://github.com/aziz8009/load-service
cd loan-service