## Dchat
Dchat app is web chatting application written in golang and vue js, with docker support
## Requirements
To run dchat you will need to install [Docker](https://www.docker.com/) (and [WSL2](https://docs.microsoft.com/en-us/windows/wsl/install) on windows)
## Installation
* first you need to download github repository using:
`git clone https://github.com/Gwestone/dchat`
* after you need to launch batch `restart-compose.bat` or shell `restart-compose.sh` scripts to automatically run docker compose or run in console commands `docker-compose build --pull` to build docker image and `docker-compose up` to lauch image
* after that you should go to address `localhost:80` and check the app