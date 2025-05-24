#!/bin/bash

set -e

echo "🔧 Installazione Rixen CLI (rx)..."

# Verifica compilazione
if [ ! -f "rx" ]; then
  echo "❌ Eseguibile 'rx' non trovato nella directory corrente."
  echo "👉 Esegui prima 'go build' o assicurati che il file sia qui."
  exit 1
fi

# Rende eseguibile
chmod +x rx

# Copia globale
sudo cp rx /usr/local/bin/rx

# Crea cartelle di lavoro
mkdir -p ~/.rx/vms
mkdir -p ~/.rx/iso
mkdir -p ~/.rx/workspaces

echo "✅ Installazione completata!"
echo "🚀 Ora puoi usare 'rx' da qualsiasi terminale."
echo "Esempio:"
echo "    rx os list"
