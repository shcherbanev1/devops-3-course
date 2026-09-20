# Используем Ubuntu 24.04 как базовый образ
FROM ubuntu:24.04

# Указываем автора образа
MAINTAINER Konstantin Shcherbanev

# Отключаем интерактивные вопросы при установке пакетов
ENV DEBIAN_FRONTEND=noninteractive

# Устанавливаем Nginx и PostgreSQL
RUN apt-get update \
    && apt-get install -y nginx postgresql \
    && rm -rf /var/lib/apt/lists/*

# Рабочая директория
WORKDIR /usr/share/nginx/html

# Добавляем главную HTML-страницу
ADD index.html .

# Копируем конфигурацию Nginx
COPY nginx.conf /etc/nginx/nginx.conf

# Том для хранения данных PostgreSQL
VOLUME ["/var/lib/postgresql"]

# Контейнер запускается от root
USER root

# Дока для открытых портов Nginx и PostgreSQL
EXPOSE 80 5432

# При запуске контейнера запускаем PostgreSQL,
# затем Nginx оставляем работать в foreground
CMD ["sh", "-c", "service postgresql start && exec nginx -g 'daemon off;'"]