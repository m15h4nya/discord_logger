# discord_loger
Bot used to log message activities such as MessageDelete, MessageEdit, etc.

# Installation
`git clone https://github.com/m15h4nya/discord_logger.git`

# Starting
Firstly, edit your config for your guild

`docker build -t logger:latest`

`docker run -dp 8080:8080 --name logger logger:latest`

Now the docker container is running

Go to the `ip:8080/start` address in your browser to start the bot or `ip:8080/stop` to stop it 

test