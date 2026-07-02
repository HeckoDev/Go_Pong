#!/bin/bash
# Configuration de l'environnement Go pour PongMaster
# Source ce fichier avec : source ./setup-env.sh

echo "🔧 Configuration de l'environnement Go..."

# Configure le proxy Go public
export GOPROXY="https://proxy.golang.org,direct"
export GONOPROXY=""
export GONOSUMDB=""

# Configure les variables Go
export CGO_ENABLED=1

echo "✅ Variables d'environnement configurées :"
echo "   GOPROXY=$GOPROXY"
echo "   CGO_ENABLED=$CGO_ENABLED"
echo ""
echo "💡 Pour rendre permanent, ajoutez ceci à votre ~/.zshrc ou ~/.bashrc :"
echo '   export GOPROXY="https://proxy.golang.org,direct"'
