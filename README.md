#  Simple User API - Deployment Bootcamp Batch 4

API sederhana untuk menampilkan dan menambahkan data users secara in-memory (tanpa database), menggunakan Golang dan Gin Gonic.  
Project ini dideploy ke VPS secara otomatis menggunakan GitHub Actions dan Docker.

---

## 📌 Public Endpoint

- GET: [`http://203.175.11.199:8004/api/users`]
- POST: [`http://203.175.11.199:8004/api/users`]

---

## 🧪 Cara Uji API (Postman/cURL)

### GET
```bash
curl http:/203.175.11.199:8004/api/users
```

### POST
```bash
curl -X POST -H "Content-Type: application/json" -d '{"name":"Andi"}' http://203.175.11.199:8004/api/users
```
## Cara Clone & Jalankan di Lokal
```bash
git clone https://github.com/feriyalbinyahya/go-deploy-vps-2.git
cd go-deploy-vps-2
go mod tidy
go run main.go
```
## Jalankan dengan Docker
docker build -t my-api .
docker run -d -p 8004:8000 my-api

## Step-by-Step Deployment ke VPS

### (Opsional) Setup Awal VPS
- SSH ke VPS:
```bash
ssh root@203.175.11.199
```
- Install Docker:
```bash
apt update && apt install docker.io -y
```
- Setup SSH Key agar GitHub Actions bisa connect:
    - Generate SSH key di lokal
    - Copy public key ke VPS (~/.ssh/authorized_keys)
    - Tambahkan private key ke GitHub Secrets

### Push kode ke GitHub
Setiap kali ada perubahan kode, commit dan push ke branch main di repository ini.

### GitHub Actions (.github/workflows/deploy.yml) berjalan otomatis
- Mentransfer file ke VPS
- Build Docker image
- Menjalankan container baru

### API otomatis tersedia di http://203.175.11.199:8004/api/users
Semua proses deployment terjadi otomatis tanpa SSH manual berkat CI/CD GitHub Actions + Docker


