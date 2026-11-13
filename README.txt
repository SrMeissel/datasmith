How to run:
    docker compose up -d --build datasmith
find volume:
    docker volume inspect datasmith_ssh_keys
    echo "KEY" | sudo tee -a PATH 
