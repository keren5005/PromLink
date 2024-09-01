MATTERMOST PROTOCOL

Suppose that MATTERMOST server runs on ${SERVER_URL}


1) Obtain token
https://api.mattermost.com/#tag/authentication
Send post request to:

${SERVER_URL}/api/v4/users/login
with JSON body
{
    "login_id":"igor.gutnik@gmail.com",
    "password":"!QAZ2wsx"
}
On success, it will return code 200, and Header: Token ${TOKEN}

2)
https://api.mattermost.com/#tag/channels/operation/GetAllChannels
In order to send message we need to know channel ID. If we want to discover channel id by name, we can
GET ${SERVER_URL}/api/v4/channels
with Authorization: Bearer ${TOKEN}

3) Posting message
https://api.mattermost.com/#tag/posts/operation/CreatePost
In order to send a message to the specific channel:

POST ${SERVER_URL}/api/v4/posts
with Authorization: Bearer ${TOKEN}
and JSON body
{
    "channel_id" : ${CHANNEL_ID},
    "message" : "hello from postman"
}


