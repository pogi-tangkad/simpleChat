# simpleChat
Simple webpage for chat with a little embedded CSS, Javascript, and a Go backend

I am trying to avoid external libraries so that the files can be placed on an air gapped
network for simple web based chatting.

My next task is to add a log file to store chat history. (possibly Python)

Once that is done, I want to create a database for storage and recall if the server needs restarted.
Since I don't want to use external libraries, I am probably going to use Python to interact with
sqlite3.

I would also like to build in accounts/authentication and sessions/rooms for that chat, but then it wouldn't be simple chat ;)
