#!/bin/bash

echo "=== Configuración de la base de datos ==="

read -p "DB_USER (default: root): " DB_USER
DB_USER=${DB_USER:-root}

read -sp "DB_PASS: " DB_PASS
echo

read -p "DB_NAME (default: midb): " DB_NAME
DB_NAME=${DB_NAME:-midb}

echo ""
echo "Descargando MySQL..."
docker pull mysql:8

echo "Descargando API..."
docker pull ghcr.io/titojuanc/gintonic:latest

echo "Creando red..."
docker network create myred 2>/dev/null

echo "Levantando MySQL..."
docker run -d \
  --name db \
  --network myred \
  -e MYSQL_ROOT_PASSWORD=$DB_PASS \
  -e MYSQL_DATABASE=$DB_NAME \
  mysql:8

echo "Esperando que MySQL arranque..."
until docker exec db mysqladmin ping -h localhost --silent; do
  sleep 2
done
sleep 5

echo "Levantando API..."
docker run -d --name gintonic \
  -p 6769:6769 \
  --network myred \
  -e DB_USER=$DB_USER \
  -e DB_PASSWORD=$DB_PASS \
  -e DB_HOST=db \
  -e DB_NAME=$DB_NAME \
  ghcr.io/titojuanc/gintonic:latest

echo ""
echo "✓ Todo listo! API corriendo en http://localhost:6769"