```markdown
# 📌 MT5 CDN Project

Ce projet est une application **Go + React (Next.js)** permettant aux utilisateurs de gérer des fichiers et dossiers dans un **système de stockage** avec des fonctionnalités avancées telles que :
- **Upload et gestion des fichiers/dossiers**
- **Téléchargement de fichiers et dossiers compressés en ZIP**
- **Suppression des fichiers et dossiers**
- **Authentification et gestion des utilisateurs**
- **Tracking des requêtes avec Prometheus et visualisation avec Grafana** *(branche `add-grafana-and-prometheus`)*  
- **Load Balancing avec Nginx et plusieurs instances du backend**



### **1️⃣ Démarrer le Backend**
make start
```

### **3️⃣ Démarrer le Frontend**
```
cd front
npm install
npm run dev
```

---

## 📡 Liste des Endpoints et Tests avec `curl`

### **📌 Authentification**

#### 🔹 **Créer un utilisateur**
```
curl -X POST http://localhost/v1/users \
     -H "Cookie: token=<Token Jwt>" \
     -d '{"email": "test@gmail.com", "password": "Amine123"}'
```

#### 🔹 **Se connecter**
```
curl -X POST http://localhost/v1/login \
     -H "Cookie: token=<Token Jwt>" \
     -d '{"email": "test@gmail.com", "password": "Amine123"}' \
     -c cookies.txt
```

#### 🔹 **Récupérer l'utilisateur authentifié**
```
curl -X GET http://localhost/v1/me -b cookies.txt
```

---

### **📌 Gestion des fichiers & dossiers**

#### 🔹 **Uploader un fichier**
```
curl -X POST http://localhost/v1/upload \
     -H "Cookie: token=<Token Jwt>" \
     -F "files=@/chemin/vers/fichier.txt" \
     -F "folder=dossier"
```

#### 🔹 **Lister les fichiers d’un dossier**
```
curl -X GET "http://localhost/v1/files?folder=dossier" -b cookies.txt
```

#### 🔹 **Télécharger un fichier**
```
curl -X GET "http://localhost/v1/download?id=<FILE_ID>" 
     -H "Cookie: token=<Token Jwt>" \
      -b cookies.txt -o fichier_telecharge.txt
```

#### 🔹 **Télécharger un dossier compressé**
```
curl -X GET "http://localhost/v1/download-folder?folder=dossier" -b cookies.txt -o dossier.zip
```

#### 🔹 **Supprimer un fichier ou un dossier**
```
curl -X DELETE "http://localhost/v1/delete?path=dossier/fichier.txt" -b cookies.txt
```

---

### **📌 Monitoring avec Prometheus & Grafana**
> 📍 Disponible sur la branche `add-grafana-and-prometheus`

#### 🔹 **Vérifier les métriques Prometheus**
```sh
curl -X GET http://localhost/metrics
```

#### 🔹 **Accéder à Grafana**
- **URL:** `http://localhost:3030`
- **Identifiants par défaut**  
  - **User:** `admin`
  - **Password:** `admin`

> ⚠️ **Grafana est opérationnel uniquement sur la branche `add-grafana-and-prometheus`**

---

## 🏗️ Architecture du Projet

### **📂 Backend (`Go + Echo`)**
- **`internal/controllers/`** : Gère les requêtes HTTP (auth, fichiers, monitoring)
- **`internal/services/`** : Contient la logique métier
- **`internal/repositories/`** : Accès à la base de données
- **`internal/middlewares/`** : Middleware (auth, CORS, monitoring)
- **`pkg/server/`** : Configuration du serveur Echo

### **📂 Frontend (`Next.js + shadcn`)**
- **`/auth/`** : Gestion de l'authentification (login, signup)
- **`/dashboard/`** : Interface utilisateur pour gérer les fichiers
- **`/components/`** : Composants réutilisables (`Table`, `Button`, `Input`, etc.)

### **📂 Infrastructure (`Docker + Nginx`)**
- **Nginx** fait le load balancing entre plusieurs instances du serveur backend (`server-1`, `server-2`, `server-3`).
- **Prometheus** collecte les métriques et **Grafana** les affiche.

---

## ✅ Fonctionnalités Implémentées
✔️ **Authentification (inscription, connexion, session avec cookies)**  
✔️ **Upload et gestion des fichiers/dossiers**  
✔️ **Téléchargement de fichiers et de dossiers compressés en ZIP**  
✔️ **Suppression de fichiers et dossiers**  
✔️ **Affichage dynamique des fichiers et historique des téléchargements**  
✔️ **Monitoring avec Prometheus (métriques HTTP)** *(branch `add-grafana-and-prometheus`)*  
✔️ **Visualisation des métriques dans Grafana** *(branch `add-grafana-and-prometheus`)*  
✔️ **Reverse Proxy et Load Balancing avec Nginx**  



## 📝 Conclusion
Ce projet est une **solution complète de gestion de fichiers** avec **authentification, stockage, monitoring et visualisation des métriques**.  
Il est conçu pour être **scalable** grâce à **Docker, Nginx (Load Balancing) et Prometheus/Grafana**. 🚀
```

