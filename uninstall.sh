#!/bin/bash

echo "🧹 Rimozione Rixen CLI (rx)..."

# Rimuovi eseguibile
if [ -f "/usr/local/bin/rx" ]; then
    sudo rm /usr/local/bin/rx
    echo "✅ Eseguibile 'rx' rimosso da /usr/local/bin"
else
    echo "⚠️ 'rx' non trovato in /usr/local/bin"
fi

# Chiede se rimuovere le cartelle di lavoro
read -p "Vuoi anche rimuovere le cartelle di lavoro ~/.rx ? (y/N): " answer
if [[ "$answer" =~ ^[Yy]$ ]]; then
    rm -rf ~/.rx
    echo "✅ Cartelle ~/.rx rimosse"
else
    echo "ℹ️ Cartelle ~/.rx conservate"
fi

echo "✅ Disinstallazione completata."
