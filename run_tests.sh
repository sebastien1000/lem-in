#!/bin/bash

# Script pour lancer automatiquement tous les fichiers de test dans le dossier tests/

# Obtenir le chemin absolu du répertoire du script
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Se déplacer vers le dossier cmd/app pour exécuter go run .
cd "$SCRIPT_DIR/cmd/app"

# Boucle sur tous les fichiers .txt dans le dossier tests/
for file in "$SCRIPT_DIR/tests"/*.txt; do
    # Extraire le nom du fichier sans le chemin
    filename=$(basename "$file")
    echo "Lancement de $filename"
    go run . "$filename"
    echo "Fin de $filename"
    echo "--------------------------------------------"
done