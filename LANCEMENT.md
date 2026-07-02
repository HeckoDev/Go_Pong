# 🎮 Guide de Lancement - Pong Master sur WSL

## ✅ Méthode Recommandée : Windows (.exe)

Le jeu fonctionne parfaitement sous Windows avec audio et graphismes :

### Depuis l'Explorateur Windows :
1. Ouvrez `\\wsl$\Ubuntu\home\alex\Sandbox\technologie\go\PongMaster`
2. Double-cliquez sur `pong.exe`

### Depuis PowerShell/CMD :
```powershell
cd \\wsl$\Ubuntu\home\alex\Sandbox\technologie\go\PongMaster
.\pong.exe
```

### Depuis WSL :
```bash
./pong.exe
```

---

## 🐧 Méthode Alternative : Linux (WSL)

⚠️ **Limitations** : L'audio peut ne pas fonctionner correctement sous WSL

### Option 1 : Script simplifié
```bash
./run-wsl.sh
```

### Option 2 : Makefile
```bash
export GOPROXY="https://proxy.golang.org,direct"
make run
```

### Option 3 : Compilation manuelle
```bash
export GOPROXY="https://proxy.golang.org,direct"
go build -o pong main.go
./pong
```

---

## 🎯 Contrôles du Jeu

- **Joueur 1** : W (haut) / S (bas)
- **Joueur 2** : ↑ (haut) / ↓ (bas)
- **Entrée** : Lancer la balle

---

## 🔧 Recompiler

### Pour Windows :
```bash
export GOPROXY="https://proxy.golang.org,direct"
make build-windows
```

### Pour Linux :
```bash
export GOPROXY="https://proxy.golang.org,direct"
make build-linux
```

### Pour toutes les plateformes :
```bash
export GOPROXY="https://proxy.golang.org,direct"
make build-all
```

---

## 🐛 Problèmes Courants

### Erreur "401 Unauthorized" lors du build
**Solution** : Configurer le proxy Go public
```bash
export GOPROXY="https://proxy.golang.org,direct"
```

### Erreur ALSA (audio) sous WSL
**Solution** : Utilisez `pong.exe` sous Windows, l'audio y fonctionne parfaitement.

### Erreur X11/Display
**Solution** : Vérifiez que WSLg est bien installé :
```bash
wsl.exe --version
```
Si WSLg n'est pas installé, mettez à jour WSL :
```powershell
wsl --update
```
