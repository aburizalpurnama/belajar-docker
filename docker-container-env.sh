docker pull mongo
docker container create --memory 500m --cpus 0.5 --name contohmongo --publish 27017:27017 --env MONGO_INITDB_ROOT_USERNAME=rizal --env MONGO_INITDB_ROOT_PASSWORD=rizal mongo:latest

docker container create --memory 500m --cpus 0.5 --name <container-name> --publish <host-port>:<container-port> --env <env-name>=<value> mongo:latest

docker container run rm --memory 500m --cpus 0.5 --name contohmongo --publish 27017:27017 --env MONGO_INITDB_ROOT_USERNAME=rizal --env MONGO_INITDB_ROOT_PASSWORD=rizal mongo:latest

- Bind Mounts

    docker container create  --memory 500m --cpus 0.5 --name mongodata --publish 27018:27017 --mount "type=bind,source=D:\rizal\Go-Projects\belajar-docker\data-db,destination=/data/db" --env MONGO_INITDB_ROOT_USERNAME=rizal --env MONGO_INITDB_ROOT_PASSWORD=rizal mongo:latest

- Bind Volume

    docker container create  --memory 500m --cpus 0.5 --name mongovolume --publish 27019:27017 --mount "type=volume,source=mongovolume,destination=/data/db" --env MONGO_INITDB_ROOT_USERNAME=rizal --env MONGO_INITDB_ROOT_PASSWORD=rizal mongo:latest

- Backup Volume

docker container create --name mongobackup --mount "type=bind,source=D:\rizal\Go-Projects\belajar-docker\backups,destination=/backup" --mount "type=volume,source=mongovolume,destination=/data" mongo:latest

Run & Remove
docker container run --rm --name ubuntubackup --mount "type=bind,source=D:\rizal\Go-Projects\belajar-docker\backups,destination=/backup" --mount "type=volume,source=mongovolume,destination=/data" ubuntu:latest tar cvf /backup/backup-mongo_2.tar.gz /data
	
Zip file

tar cvf /backup/backup-mongo_1.tar.gz /data


- Restore Volume

docker container run --rm --name ubuntubackup --mount "type=bind,source=D:\rizal\Go-Projects\belajar-docker\backups,destination=/backup" --mount "type=volume,source=mongorestore,destination=/data" ubuntu:latest bash -c "cd /data && tar xvf /backup/backup-mongo_1.tar.gz --strip 1"

    buat container baru

docker container create  --memory 500m --cpus 0.5 --name mongorestore --publish 27020:27017 --mount "type=volume,source=mongorestore,destination=/data/db" --env MONGO_INITDB_ROOT_USERNAME=rizal --env MONGO_INITDB_ROOT_PASSWORD=rizal mongo:latest