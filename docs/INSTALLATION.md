# 🚀 Installation et Configuration

## 1. Installer Go

### Sur Linux (Debian/Ubuntu)
```bash
wget https://go.dev/dl/go1.23.0.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.23.0.linux-amd64.tar.gz
export PATH=/home/tpassin/go1.23/go/bin:/home/tpassin/go-workspace/bin:/home/tpassin/.nvm/versions/node/v22.17.1/bin:/home/tpassin/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/snap/bin:/home/tpassin/go/bin:/home/tpassin/go/bin:/usr/local/go/bin
```

### Sur macOS
```bash
brew install go
```

### Sur Windows
[Instructions détaillées...]

## 2. Cloner et configurer

```bash
git clone https://github.com/PassinThomas/task-tracker.git
cd task-tracker
go mod download
```

## 3. Compilation avec Makefile

```bash
make build      # Compile le binaire
make test       # Lance les tests
make run        # Build + exécute list
make clean      # Nettoie
```
