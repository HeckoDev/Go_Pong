#!/bin/bash

# Script de lancement du jeu Pong
echo "🎮 Lancement du jeu Pong..."
echo "📋 Contrôles:"
echo "   Joueur 1: W (haut) / S (bas)"
echo "   Joueur 2: ↑ (haut) / ↓ (bas)"
echo "   Espace: Redémarrer la balle"
echo ""

# Compilation et lancement
go build && ./pong