# Rixen VM

Rixen (`rx`) è una CLI open-source pensata per sviluppatori che vogliono gestire macchine virtuali su macOS usando QEMU in modo semplice, pulito e potente.

## 🚀 Funzionalità principali

- Creazione VM da ISO o selezione automatica OS (`rx create`)
- Avvio VM con cartella condivisa (`rx start`)
- Download automatico ISO (Ubuntu, Fedora, Windows redirect)
- Rete con port forwarding SSH
- Struttura modularizzata e pronta per espansione

## 📦 Installazione

```bash
tar -xvzf rixen-vm-release.tar.gz
cd rixen-vm-clean
./install.sh
```

## 🧪 Comandi principali

```bash
rx os list
rx os versions ubuntu
rx create --name testvm --os ubuntu --cpu 2 --ram 2048 --disk 20
rx start testvm
```

## 🧹 Disinstallazione

```bash
./uninstall.sh
```

---

Sviluppato con ❤️ per sviluppatori.
