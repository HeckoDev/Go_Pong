#!/bin/bash
# Script pour lancer Pong sous WSL sans audio

echo "🎮 Lancement de Pong Master sous WSL..."
echo "⚠️  L'audio est désactivé en mode WSL"
echo ""
echo "📋 Contrôles:"
echo "   Joueur 1: W (haut) / S (bas)"
echo "   Joueur 2: ↑ (haut) / ↓ (bas)"  
echo "   Entrée: Lancer la balle"
echo ""

# Configuration pour éviter les erreurs audio
export SDL_AUDIODRIVER=dummy
export ALSA_CARD=0

# Compilation avec proxy Go public
export GOPROXY="https://proxy.golang.org,direct"

# Build et lancement
if go build -o pong main.go 2>&1 | grep -v "ALSA lib"; then
    ./pong 2>&1 | grep -v "ALSA lib" | grep -v "XGB:"
else
    echo "❌ Erreur de compilation"
    exit 1
fi
