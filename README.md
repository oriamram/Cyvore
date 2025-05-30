Steps to get started

git clone https://github.com/oriamram/Cyvore.git

docker pull caffix/amass

go into backend and do go mod download and tidy

go into frontend and do npm ci

now because I chose not to use monorepo manager youll have to run all services seperately

each in its folder -
frontend: npm run dev
server: go run cmd/server/main.go
websocket: go run cmd/server/main.go

now all services should work and you can enter localhost:5173 and meet the login page
