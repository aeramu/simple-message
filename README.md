# simple-message
Backend API server to send message implemented in golang. Simple-message can get all message on database,
send a message, and subscribe live message through websocket
## Getting Started
### Clone repo
```bash
git clone https://github.com/aeramu/simple-message
```
### Using Docker
This project included dockerfile, so you can use docker instead if it's easier.
```bash
sudo docker build -t message
sudo docker run -p 8080:8080 message
```
### Using Go
#### Prerequisites
Install Golang
```bash
sudo pacman -S go
```
#### Running
Just run it with Go. Golang will automatically install the dependency
```bash
go run ./cmd/message
```
### Websocket Client
There is a websocket client implemented in html in folder 
```
firefox ./cmd/client/index.html
```
Open web console on the browser to see the messages
## API
#### Port: 8080
#### GET /messages
Return list of messages
#### POST /messages
Request Body:
```
{
    "body":"some message"
}
```
#### Websocket /live
Websocket endpoint to subscribe message from server

