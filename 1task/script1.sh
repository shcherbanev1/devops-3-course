#!/bin/bash

LOG_FILE="./script.log"
exec > >(tee -a "$LOG_FILE") 2>&1

echo "Script started"

BASE_DIR=""

while getopts "d:" opt; do
    case "$opt" in
        d)
            BASE_DIR="$OPTARG"
            ;;
        *)
            echo "Usage: $0 [-d path]"
            exit 1
            ;;
    esac
done

if [[ -z "$BASE_DIR" ]]; then
    read -rp "Enter path: " BASE_DIR
fi

mkdir -p "$BASE_DIR"

if getent group dev > /dev/null; then
    echo "Group dev already exists"
else
    groupadd dev
    echo "Group dev created"
fi

USERS=$(awk -F: '$3 >= 1000 && $1 != "nobody" {print $1}' /etc/passwd)

for USER in $USERS; do
    usermod -aG dev "$USER"
    echo "Added $USER to dev"
done

echo "%dev ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/dev
chmod 440 /etc/sudoers.d/dev

for USER in $USERS; do
    WORK_DIR="$BASE_DIR/${USER}_workdir"
    PRIMARY_GROUP=$(id -gn "$USER")

    mkdir -p "$WORK_DIR"
    chown "$USER:$PRIMARY_GROUP" "$WORK_DIR"
    chmod 660 "$WORK_DIR"

    setfacl -m g:dev:r "$WORK_DIR"

    echo "Created $WORK_DIR"
done

echo "Script finished"