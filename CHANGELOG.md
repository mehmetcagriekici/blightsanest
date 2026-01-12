# Dev Logs

If you face any bugs problems or something not clear, please do reach me from mehmetcagriekici@gmail.com

## Initial Relase V1

1)  Created a subscription manager for client subscriber cancellation logic. Consumer cancel functions are hold for later instead of immediately used in the REPL preventing message deliveries.
2)  Implemented a new routing key system making sure subscribers match the publishers.
3)  Fixed the issues with data flow from the server to the clients by organizing subscriber callback functions.
4)  Separated client and server publishings by moving each to different exchanges.
5)  Implemented more API queries directly to the server side, allowing user to work on more specific crypto lists.
6)  Added DLX and manual acknowledgement logic to prevent unnecessary subscriber exits.
7)  Removed mutate command to manually update the current client list, instead all commands automatically updates the current client list with the result.
8)  Changed client rank, now accepting a wider range of coin fields.
9)  Simplified the CLI printing, making the output more readable.
10) Implemented a new command <set> to the client side, allowing users to config client state preferences beforehand.
11) Fixed some of the crypto algorithms and added a test for the pubsub logic
12) Implemented a new command <get> to the server side, allowing users to get crypto lists from the local database.
13) Implemented a new command <save> to the server, allowing users to save various crypto lists to the database with a custom id.
